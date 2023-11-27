// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pb/chat.proto

package pb

import (
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int32  `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []string `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetWords() []string {
	if x != nil {
		return x.Words
	}
	return nil
}

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  int32              `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Name   string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Input  string             `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Wallet map[string]float32 `protobuf:"bytes,4,rep,name=wallet,proto3" json:"wallet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *ChatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *ChatRequest) GetWallet() map[string]float32 {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  int32              `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Answer []string           `protobuf:"bytes,2,rep,name=answer,proto3" json:"answer,omitempty"`
	Wallet map[string]float32 `protobuf:"bytes,3,rep,name=wallet,proto3" json:"wallet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatResponse) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *ChatResponse) GetAnswer() []string {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *ChatResponse) GetWallet() map[string]float32 {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type TimerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    int32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	LastTime int32 `protobuf:"varint,2,opt,name=last_time,json=lastTime,proto3" json:"last_time,omitempty"`
	NowTime  int32 `protobuf:"varint,3,opt,name=now_time,json=nowTime,proto3" json:"now_time,omitempty"`
}

func (x *TimerRequest) Reset() {
	*x = TimerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerRequest) ProtoMessage() {}

func (x *TimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerRequest.ProtoReflect.Descriptor instead.
func (*TimerRequest) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{4}
}

func (x *TimerRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *TimerRequest) GetLastTime() int32 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

func (x *TimerRequest) GetNowTime() int32 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

type TimerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  int32    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	IsExit bool     `protobuf:"varint,2,opt,name=is_exit,json=isExit,proto3" json:"is_exit,omitempty"`
	Reset_ bool     `protobuf:"varint,3,opt,name=reset,proto3" json:"reset,omitempty"`
	Answer []string `protobuf:"bytes,4,rep,name=answer,proto3" json:"answer,omitempty"`
}

func (x *TimerResponse) Reset() {
	*x = TimerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerResponse) ProtoMessage() {}

func (x *TimerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerResponse.ProtoReflect.Descriptor instead.
func (*TimerResponse) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{5}
}

func (x *TimerResponse) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *TimerResponse) GetIsExit() bool {
	if x != nil {
		return x.IsExit
	}
	return false
}

func (x *TimerResponse) GetReset_() bool {
	if x != nil {
		return x.Reset_
	}
	return false
}

func (x *TimerResponse) GetAnswer() []string {
	if x != nil {
		return x.Answer
	}
	return nil
}

var File_pb_chat_proto protoreflect.FileDescriptor

var file_pb_chat_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xad, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x77, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x65, 0x78, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x45, 0x78,
	0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x32, 0x41, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0x3c, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x34, 0x0a, 0x0d, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x3e, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0c, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_chat_proto_rawDescOnce sync.Once
	file_pb_chat_proto_rawDescData = file_pb_chat_proto_rawDesc
)

func file_pb_chat_proto_rawDescGZIP() []byte {
	file_pb_chat_proto_rawDescOnce.Do(func() {
		file_pb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_chat_proto_rawDescData)
	})
	return file_pb_chat_proto_rawDescData
}

var file_pb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_chat_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: pb.HelloRequest
	(*HelloResponse)(nil), // 1: pb.HelloResponse
	(*ChatRequest)(nil),   // 2: pb.ChatRequest
	(*ChatResponse)(nil),  // 3: pb.ChatResponse
	(*TimerRequest)(nil),  // 4: pb.TimerRequest
	(*TimerResponse)(nil), // 5: pb.TimerResponse
	nil,                   // 6: pb.ChatRequest.WalletEntry
	nil,                   // 7: pb.ChatResponse.WalletEntry
}
var file_pb_chat_proto_depIdxs = []int32{
	6, // 0: pb.ChatRequest.wallet:type_name -> pb.ChatRequest.WalletEntry
	7, // 1: pb.ChatResponse.wallet:type_name -> pb.ChatResponse.WalletEntry
	0, // 2: pb.Greet.SayHelloService:input_type -> pb.HelloRequest
	2, // 3: pb.Chat.AnswerService:input_type -> pb.ChatRequest
	4, // 4: pb.Timer.TimerService:input_type -> pb.TimerRequest
	1, // 5: pb.Greet.SayHelloService:output_type -> pb.HelloResponse
	3, // 6: pb.Chat.AnswerService:output_type -> pb.ChatResponse
	5, // 7: pb.Timer.TimerService:output_type -> pb.TimerResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_chat_proto_init() }
func file_pb_chat_proto_init() {
	if File_pb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_pb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_pb_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_pb_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
		file_pb_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerRequest); i {
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
		file_pb_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerResponse); i {
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
			RawDescriptor: file_pb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_pb_chat_proto_goTypes,
		DependencyIndexes: file_pb_chat_proto_depIdxs,
		MessageInfos:      file_pb_chat_proto_msgTypes,
	}.Build()
	File_pb_chat_proto = out.File
	file_pb_chat_proto_rawDesc = nil
	file_pb_chat_proto_goTypes = nil
	file_pb_chat_proto_depIdxs = nil
}
