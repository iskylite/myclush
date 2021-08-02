// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: stream.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// putStream request
type PutStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Md5      string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Body     []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Sn       string `protobuf:"bytes,5,opt,name=sn,proto3" json:"sn,omitempty"`
	Port     string `protobuf:"bytes,6,opt,name=port,proto3" json:"port,omitempty"`
	Nodelist string `protobuf:"bytes,7,opt,name=nodelist,proto3" json:"nodelist,omitempty"`
	Width    int32  `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	Uid      uint32 `protobuf:"varint,9,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid      uint32 `protobuf:"varint,10,opt,name=gid,proto3" json:"gid,omitempty"`
	Filemod  uint32 `protobuf:"varint,11,opt,name=filemod,proto3" json:"filemod,omitempty"`
	Modtime  int64  `protobuf:"varint,12,opt,name=modtime,proto3" json:"modtime,omitempty"`
}

func (x *PutStreamReq) Reset() {
	*x = PutStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutStreamReq) ProtoMessage() {}

func (x *PutStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutStreamReq.ProtoReflect.Descriptor instead.
func (*PutStreamReq) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *PutStreamReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutStreamReq) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *PutStreamReq) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *PutStreamReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *PutStreamReq) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *PutStreamReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *PutStreamReq) GetNodelist() string {
	if x != nil {
		return x.Nodelist
	}
	return ""
}

func (x *PutStreamReq) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *PutStreamReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PutStreamReq) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *PutStreamReq) GetFilemod() uint32 {
	if x != nil {
		return x.Filemod
	}
	return 0
}

func (x *PutStreamReq) GetModtime() int64 {
	if x != nil {
		return x.Modtime
	}
	return 0
}

// node struct
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width int32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Value string  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Node) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Node) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// putStream/runcmd response
type PutStreamResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replay []*Replay `protobuf:"bytes,1,rep,name=replay,proto3" json:"replay,omitempty"`
}

func (x *PutStreamResp) Reset() {
	*x = PutStreamResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutStreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutStreamResp) ProtoMessage() {}

func (x *PutStreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutStreamResp.ProtoReflect.Descriptor instead.
func (*PutStreamResp) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *PutStreamResp) GetReplay() []*Replay {
	if x != nil {
		return x.Replay
	}
	return nil
}

type Replay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass     bool   `protobuf:"varint,1,opt,name=pass,proto3" json:"pass,omitempty"`
	Nodelist string `protobuf:"bytes,2,opt,name=nodelist,proto3" json:"nodelist,omitempty"`
	Msg      string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Replay) Reset() {
	*x = Replay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replay) ProtoMessage() {}

func (x *Replay) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replay.ProtoReflect.Descriptor instead.
func (*Replay) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *Replay) GetPass() bool {
	if x != nil {
		return x.Pass
	}
	return false
}

func (x *Replay) GetNodelist() string {
	if x != nil {
		return x.Nodelist
	}
	return ""
}

func (x *Replay) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// runcmd request
type CmdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd      string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Timeout  int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Nodelist string `protobuf:"bytes,3,opt,name=nodelist,proto3" json:"nodelist,omitempty"`
	Width    int32  `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Port     string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *CmdReq) Reset() {
	*x = CmdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdReq) ProtoMessage() {}

func (x *CmdReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdReq.ProtoReflect.Descriptor instead.
func (*CmdReq) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{4}
}

func (x *CmdReq) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *CmdReq) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *CmdReq) GetNodelist() string {
	if x != nil {
		return x.Nodelist
	}
	return ""
}

func (x *CmdReq) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *CmdReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type GG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HH string `protobuf:"bytes,1,opt,name=HH,proto3" json:"HH,omitempty"`
}

func (x *GG) Reset() {
	*x = GG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GG) ProtoMessage() {}

func (x *GG) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GG.ProtoReflect.Descriptor instead.
func (*GG) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{5}
}

func (x *GG) GetHH() string {
	if x != nil {
		return x.HH
	}
	return ""
}

// oo_bw_server 测试
type OoBwServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ncid    string `protobuf:"bytes,1,opt,name=ncid,proto3" json:"ncid,omitempty"`
	Buffer  string `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Length  int32  `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	Loop    int32  `protobuf:"varint,5,opt,name=loop,proto3" json:"loop,omitempty"`
}

func (x *OoBwServerReq) Reset() {
	*x = OoBwServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OoBwServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OoBwServerReq) ProtoMessage() {}

func (x *OoBwServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OoBwServerReq.ProtoReflect.Descriptor instead.
func (*OoBwServerReq) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{6}
}

func (x *OoBwServerReq) GetNcid() string {
	if x != nil {
		return x.Ncid
	}
	return ""
}

func (x *OoBwServerReq) GetBuffer() string {
	if x != nil {
		return x.Buffer
	}
	return ""
}

func (x *OoBwServerReq) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *OoBwServerReq) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *OoBwServerReq) GetLoop() int32 {
	if x != nil {
		return x.Loop
	}
	return 0
}

// oo_bw_client 测试
type OoBwClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server  string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Buffer  string `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Timer   int64  `protobuf:"varint,4,opt,name=timer,proto3" json:"timer,omitempty"`
	Port    string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Length  int32  `protobuf:"varint,6,opt,name=length,proto3" json:"length,omitempty"`
	Loop    int32  `protobuf:"varint,7,opt,name=loop,proto3" json:"loop,omitempty"`
}

func (x *OoBwClientReq) Reset() {
	*x = OoBwClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OoBwClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OoBwClientReq) ProtoMessage() {}

func (x *OoBwClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OoBwClientReq.ProtoReflect.Descriptor instead.
func (*OoBwClientReq) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{7}
}

func (x *OoBwClientReq) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *OoBwClientReq) GetBuffer() string {
	if x != nil {
		return x.Buffer
	}
	return ""
}

func (x *OoBwClientReq) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *OoBwClientReq) GetTimer() int64 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *OoBwClientReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *OoBwClientReq) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *OoBwClientReq) GetLoop() int32 {
	if x != nil {
		return x.Loop
	}
	return 0
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x92, 0x02, 0x0a, 0x0c, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x67, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x6f, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x6f, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x6f, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x6f, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x50,
	0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x22, 0x4a, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7a, 0x0a, 0x06,
	0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x14, 0x0a, 0x02, 0x47, 0x47, 0x12, 0x0e,
	0x0a, 0x02, 0x48, 0x48, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x48, 0x48, 0x22, 0x81,
	0x01, 0x0a, 0x0d, 0x4f, 0x6f, 0x42, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x63, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f,
	0x6f, 0x70, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x4f, 0x6f, 0x42, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6c, 0x6f, 0x6f, 0x70, 0x32, 0x80, 0x02, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x43, 0x6d,
	0x64, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x47,
	0x1a, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x47, 0x12, 0x1d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e,
	0x63, 0x69, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x47, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4f, 0x6f,
	0x42, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x6f,
	0x42, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4f, 0x6f,
	0x42, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x6f,
	0x42, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stream_proto_goTypes = []interface{}{
	(*PutStreamReq)(nil),  // 0: pb.PutStreamReq
	(*Node)(nil),          // 1: pb.Node
	(*PutStreamResp)(nil), // 2: pb.PutStreamResp
	(*Replay)(nil),        // 3: pb.Replay
	(*CmdReq)(nil),        // 4: pb.CmdReq
	(*GG)(nil),            // 5: pb.GG
	(*OoBwServerReq)(nil), // 6: pb.OoBwServerReq
	(*OoBwClientReq)(nil), // 7: pb.OoBwClientReq
}
var file_stream_proto_depIdxs = []int32{
	1, // 0: pb.Node.nodes:type_name -> pb.Node
	3, // 1: pb.PutStreamResp.replay:type_name -> pb.Replay
	0, // 2: pb.RpcService.PutStream:input_type -> pb.PutStreamReq
	4, // 3: pb.RpcService.RunCmd:input_type -> pb.CmdReq
	5, // 4: pb.RpcService.Ping:input_type -> pb.GG
	5, // 5: pb.RpcService.GetNcid:input_type -> pb.GG
	6, // 6: pb.RpcService.RunOoBwServer:input_type -> pb.OoBwServerReq
	7, // 7: pb.RpcService.RunOoBwClient:input_type -> pb.OoBwClientReq
	2, // 8: pb.RpcService.PutStream:output_type -> pb.PutStreamResp
	2, // 9: pb.RpcService.RunCmd:output_type -> pb.PutStreamResp
	5, // 10: pb.RpcService.Ping:output_type -> pb.GG
	3, // 11: pb.RpcService.GetNcid:output_type -> pb.Replay
	3, // 12: pb.RpcService.RunOoBwServer:output_type -> pb.Replay
	3, // 13: pb.RpcService.RunOoBwClient:output_type -> pb.Replay
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutStreamReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutStreamResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replay); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GG); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OoBwServerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OoBwClientReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServiceClient interface {
	PutStream(ctx context.Context, opts ...grpc.CallOption) (RpcService_PutStreamClient, error)
	RunCmd(ctx context.Context, in *CmdReq, opts ...grpc.CallOption) (*PutStreamResp, error)
	Ping(ctx context.Context, in *GG, opts ...grpc.CallOption) (*GG, error)
	GetNcid(ctx context.Context, in *GG, opts ...grpc.CallOption) (*Replay, error)
	RunOoBwServer(ctx context.Context, in *OoBwServerReq, opts ...grpc.CallOption) (*Replay, error)
	RunOoBwClient(ctx context.Context, in *OoBwClientReq, opts ...grpc.CallOption) (*Replay, error)
}

type rpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcServiceClient(cc grpc.ClientConnInterface) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (RpcService_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcService_serviceDesc.Streams[0], "/pb.RpcService/PutStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcServicePutStreamClient{stream}
	return x, nil
}

type RpcService_PutStreamClient interface {
	Send(*PutStreamReq) error
	CloseAndRecv() (*PutStreamResp, error)
	grpc.ClientStream
}

type rpcServicePutStreamClient struct {
	grpc.ClientStream
}

func (x *rpcServicePutStreamClient) Send(m *PutStreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcServicePutStreamClient) CloseAndRecv() (*PutStreamResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutStreamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcServiceClient) RunCmd(ctx context.Context, in *CmdReq, opts ...grpc.CallOption) (*PutStreamResp, error) {
	out := new(PutStreamResp)
	err := c.cc.Invoke(ctx, "/pb.RpcService/RunCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) Ping(ctx context.Context, in *GG, opts ...grpc.CallOption) (*GG, error) {
	out := new(GG)
	err := c.cc.Invoke(ctx, "/pb.RpcService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) GetNcid(ctx context.Context, in *GG, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/pb.RpcService/GetNcid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) RunOoBwServer(ctx context.Context, in *OoBwServerReq, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/pb.RpcService/RunOoBwServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) RunOoBwClient(ctx context.Context, in *OoBwClientReq, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/pb.RpcService/RunOoBwClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServiceServer is the server API for RpcService service.
type RpcServiceServer interface {
	PutStream(RpcService_PutStreamServer) error
	RunCmd(context.Context, *CmdReq) (*PutStreamResp, error)
	Ping(context.Context, *GG) (*GG, error)
	GetNcid(context.Context, *GG) (*Replay, error)
	RunOoBwServer(context.Context, *OoBwServerReq) (*Replay, error)
	RunOoBwClient(context.Context, *OoBwClientReq) (*Replay, error)
}

// UnimplementedRpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServiceServer struct {
}

func (*UnimplementedRpcServiceServer) PutStream(RpcService_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}
func (*UnimplementedRpcServiceServer) RunCmd(context.Context, *CmdReq) (*PutStreamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCmd not implemented")
}
func (*UnimplementedRpcServiceServer) Ping(context.Context, *GG) (*GG, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedRpcServiceServer) GetNcid(context.Context, *GG) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNcid not implemented")
}
func (*UnimplementedRpcServiceServer) RunOoBwServer(context.Context, *OoBwServerReq) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunOoBwServer not implemented")
}
func (*UnimplementedRpcServiceServer) RunOoBwClient(context.Context, *OoBwClientReq) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunOoBwClient not implemented")
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServiceServer).PutStream(&rpcServicePutStreamServer{stream})
}

type RpcService_PutStreamServer interface {
	SendAndClose(*PutStreamResp) error
	Recv() (*PutStreamReq, error)
	grpc.ServerStream
}

type rpcServicePutStreamServer struct {
	grpc.ServerStream
}

func (x *rpcServicePutStreamServer) SendAndClose(m *PutStreamResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcServicePutStreamServer) Recv() (*PutStreamReq, error) {
	m := new(PutStreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RpcService_RunCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).RunCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RpcService/RunCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).RunCmd(ctx, req.(*CmdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RpcService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).Ping(ctx, req.(*GG))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_GetNcid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).GetNcid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RpcService/GetNcid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).GetNcid(ctx, req.(*GG))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_RunOoBwServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OoBwServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).RunOoBwServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RpcService/RunOoBwServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).RunOoBwServer(ctx, req.(*OoBwServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_RunOoBwClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OoBwClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).RunOoBwClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RpcService/RunOoBwClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).RunOoBwClient(ctx, req.(*OoBwClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCmd",
			Handler:    _RpcService_RunCmd_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _RpcService_Ping_Handler,
		},
		{
			MethodName: "GetNcid",
			Handler:    _RpcService_GetNcid_Handler,
		},
		{
			MethodName: "RunOoBwServer",
			Handler:    _RpcService_RunOoBwServer_Handler,
		},
		{
			MethodName: "RunOoBwClient",
			Handler:    _RpcService_RunOoBwClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutStream",
			Handler:       _RpcService_PutStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}