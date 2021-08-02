package service

import (
	"context"
	"fmt"
	"math"
	log "myclush/logger"
	"myclush/pb"
	"myclush/utils"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func (p *putStreamServer) RunOoBwClient(ctx context.Context, req *pb.OoBwClientReq) (*pb.Replay, error) {
	// 获取本地ncid
	localNode := fmt.Sprintf("%s -> %s", utils.Hostname(), req.Server)
	empty := newReplay(false, "nil", localNode)
	localNcid, err := getLocalNcidWithContext(ctx)
	if err != nil {
		log.Error(err)
		return empty, err
	}
	log.Infof("Local Ncid is [%s]\n", localNcid)
	// 建立服务端连接
	addr := fmt.Sprintf("%s:%s", req.Server, req.Port)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	if err != nil {
		return newReplay(false, err.Error(), localNode), err
	}
	client := pb.NewRpcServiceClient(conn)
	// 获取服务端ncid
	serverNcid := ""
	replay, err := client.GetNcid(ctx, &pb.GG{HH: req.Server})
	if err != nil {
		log.Error(err)
		return empty, err
	}
	serverNcid = replay.Msg
	log.Infof("Server Ncid is [%s]\n", serverNcid)
	// 开启服务端
	ctxs, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, length, loop int32) {
		defer wg.Done()
		replay, err := client.RunOoBwServer(ctx, &pb.OoBwServerReq{
			Ncid:    localNcid,
			Buffer:  req.Buffer,
			Timeout: req.Timeout,
			Length:  length,
			Loop:    loop,
		})
		if err != nil {
			log.Error(err)
		} else {
			log.Debugf("node %s server pass %t\n", replay.Nodelist, replay.Pass)
		}
	}(ctxs, &wg, req.Length, req.Loop)
	// 启动定时器
	log.Debugf("Timer To %s\n", utils.FormatTime(req.Timer))
	// dateNanoUnix := utils.NewTimerAfterSeconds(req.Timer)
	// file test
	var cmdFile, args string
	if req.Length != 0 && req.Loop != 0 {
		cmdFile = "/usr/local/glex/examples/oo_bw_s_loop"
		args = fmt.Sprintf("8 %s 8 %s %d %d", serverNcid, req.Buffer, req.Length, req.Loop)
	} else {
		cmdFile = "/usr/local/glex/examples/oo_bw_s"
		args = fmt.Sprintf("8 %s 8 %s", serverNcid, req.Buffer)
	}
	if !utils.Isfile(cmdFile) {
		log.Errorf("%s not exist", cmdFile)
		return empty, fmt.Errorf("%s not exist", cmdFile)
	}
	execname := filepath.Base(cmdFile)
	// 启动定时器
	timer := utils.GenTikerWithTimer(req.Timer)
	<-timer
	log.Debugf("%s start on %s\n", execname, utils.LocalTime())
	// 启动客户端RunOoBwClient
	if req.Length == 0 || req.Loop == 0 {
		var cancel1 context.CancelFunc
		ctx, cancel1 = context.WithTimeout(ctx, utils.Timeout(int(req.Timeout)))
		log.Debugf("%s set timeout %d\n", execname, req.Timeout)
		defer cancel1()
	}
	command := fmt.Sprintf("%s %s", cmdFile, args)
	out, err := utils.ExecuteShellCmdWithContext(ctx, command)
	defer log.Debugf("%s finish on %s\n", execname, utils.LocalTime())
	if err != nil {
		log.Error(err)
		return empty, err
	}
	wg.Wait()
	return newReplay(true, string(out), localNode), nil
}

func RunOoBwClientService(ctx context.Context, server, node, buffer, port string, timeout int32, timer int64, results chan *pb.Replay, wg *sync.WaitGroup, oobwloop bool, length, loop int) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%s", node, port)
	if !oobwloop {
		sumTimeout := time.Second * (time.Duration(timeout) + time.Duration(math.Ceil(time.Until(time.Unix(0, timer)).Seconds())))
		log.Debugf("Total Timeout is %d\n", sumTimeout)
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, sumTimeout)
		defer cancel()
		length = 0
		loop = 0
	}
	log.Debugf("Length=%d, Loop=%d\n", length, loop)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		results <- newReplay(false, err.Error(), node)
		return
	}
	client := pb.NewRpcServiceClient(conn)
	log.Debugf("oo_bw test start on %s -> %s\n", node, server)
	replay, err := client.RunOoBwClient(ctx, &pb.OoBwClientReq{
		Server:  server,
		Buffer:  buffer,
		Timeout: timeout,
		Timer:   timer,
		Port:    port,
		Length:  int32(length),
		Loop:    int32(loop),
	})
	if err != nil {
		results <- newReplay(false, err.Error(), node)
		return
	}
	results <- replay
}

func ParseOoBwResult(result string) (float32, error) {
	var avg float32
	re := regexp.MustCompile(`.*\d+\s+\d+\s+[0-9\.]+\s+([0-9\.]+)\s+OK.*`)
	matches := re.FindAllStringSubmatch(result, -1)
	if len(matches) == 0 {
		return avg, fmt.Errorf("no matches")
	}
	cnts := len(matches)
	for _, match := range matches {
		write, err := strconv.ParseFloat(match[1], 32)
		if err != nil {
			return avg, err
		}
		avg += float32(write)
	}
	return avg / float32(cnts), nil
}