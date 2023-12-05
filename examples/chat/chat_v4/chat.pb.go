// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: chat_v4/chat.proto

package chat_v4

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

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v4_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v4_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chat_v4_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Roomid   string `protobuf:"bytes,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v4_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v4_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_chat_v4_chat_proto_rawDescGZIP(), []int{1}
}

func (x *JoinRoomRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *JoinRoomRequest) GetRoomid() string {
	if x != nil {
		return x.Roomid
	}
	return ""
}

type JoinRoomReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Roomid string `protobuf:"bytes,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
	Uid    string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *JoinRoomReply) Reset() {
	*x = JoinRoomReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v4_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomReply) ProtoMessage() {}

func (x *JoinRoomReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v4_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomReply.ProtoReflect.Descriptor instead.
func (*JoinRoomReply) Descriptor() ([]byte, []int) {
	return file_chat_v4_chat_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRoomReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JoinRoomReply) GetRoomid() string {
	if x != nil {
		return x.Roomid
	}
	return ""
}

func (x *JoinRoomReply) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// Wraps multiple message actions.
type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*ChatRequest_JoinRoomRequest
	//	*ChatRequest_ChatMessage
	Action isChatRequest_Action `protobuf_oneof:"action"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v4_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v4_chat_proto_msgTypes[3]
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
	return file_chat_v4_chat_proto_rawDescGZIP(), []int{3}
}

func (m *ChatRequest) GetAction() isChatRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ChatRequest) GetJoinRoomRequest() *JoinRoomRequest {
	if x, ok := x.GetAction().(*ChatRequest_JoinRoomRequest); ok {
		return x.JoinRoomRequest
	}
	return nil
}

func (x *ChatRequest) GetChatMessage() *ChatMessage {
	if x, ok := x.GetAction().(*ChatRequest_ChatMessage); ok {
		return x.ChatMessage
	}
	return nil
}

type isChatRequest_Action interface {
	isChatRequest_Action()
}

type ChatRequest_JoinRoomRequest struct {
	JoinRoomRequest *JoinRoomRequest `protobuf:"bytes,1,opt,name=joinRoomRequest,proto3,oneof"`
}

type ChatRequest_ChatMessage struct {
	ChatMessage *ChatMessage `protobuf:"bytes,2,opt,name=chatMessage,proto3,oneof"`
}

func (*ChatRequest_JoinRoomRequest) isChatRequest_Action() {}

func (*ChatRequest_ChatMessage) isChatRequest_Action() {}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*ChatResponse_JoinRoomReply
	//	*ChatResponse_ChatMessage
	Action isChatResponse_Action `protobuf_oneof:"action"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v4_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v4_chat_proto_msgTypes[4]
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
	return file_chat_v4_chat_proto_rawDescGZIP(), []int{4}
}

func (m *ChatResponse) GetAction() isChatResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ChatResponse) GetJoinRoomReply() *JoinRoomReply {
	if x, ok := x.GetAction().(*ChatResponse_JoinRoomReply); ok {
		return x.JoinRoomReply
	}
	return nil
}

func (x *ChatResponse) GetChatMessage() *ChatMessage {
	if x, ok := x.GetAction().(*ChatResponse_ChatMessage); ok {
		return x.ChatMessage
	}
	return nil
}

type isChatResponse_Action interface {
	isChatResponse_Action()
}

type ChatResponse_JoinRoomReply struct {
	JoinRoomReply *JoinRoomReply `protobuf:"bytes,1,opt,name=joinRoomReply,proto3,oneof"`
}

type ChatResponse_ChatMessage struct {
	ChatMessage *ChatMessage `protobuf:"bytes,2,opt,name=chatMessage,proto3,oneof"`
}

func (*ChatResponse_JoinRoomReply) isChatResponse_Action() {}

func (*ChatResponse_ChatMessage) isChatResponse_Action() {}

var File_chat_v4_chat_proto protoreflect.FileDescriptor

var file_chat_v4_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x22, 0x43, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0d, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0f, 0x6a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x76, 0x34, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x83, 0x01, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x18,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x76, 0x34, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x76, 0x34, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v4_chat_proto_rawDescOnce sync.Once
	file_chat_v4_chat_proto_rawDescData = file_chat_v4_chat_proto_rawDesc
)

func file_chat_v4_chat_proto_rawDescGZIP() []byte {
	file_chat_v4_chat_proto_rawDescOnce.Do(func() {
		file_chat_v4_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v4_chat_proto_rawDescData)
	})
	return file_chat_v4_chat_proto_rawDescData
}

var file_chat_v4_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_v4_chat_proto_goTypes = []interface{}{
	(*ChatMessage)(nil),     // 0: chat_v4.ChatMessage
	(*JoinRoomRequest)(nil), // 1: chat_v4.JoinRoomRequest
	(*JoinRoomReply)(nil),   // 2: chat_v4.JoinRoomReply
	(*ChatRequest)(nil),     // 3: chat_v4.ChatRequest
	(*ChatResponse)(nil),    // 4: chat_v4.ChatResponse
}
var file_chat_v4_chat_proto_depIdxs = []int32{
	1, // 0: chat_v4.ChatRequest.joinRoomRequest:type_name -> chat_v4.JoinRoomRequest
	0, // 1: chat_v4.ChatRequest.chatMessage:type_name -> chat_v4.ChatMessage
	2, // 2: chat_v4.ChatResponse.joinRoomReply:type_name -> chat_v4.JoinRoomReply
	0, // 3: chat_v4.ChatResponse.chatMessage:type_name -> chat_v4.ChatMessage
	1, // 4: chat_v4.Chat.JoinRoom:input_type -> chat_v4.JoinRoomRequest
	0, // 5: chat_v4.Chat.SendMessage:input_type -> chat_v4.ChatMessage
	2, // 6: chat_v4.Chat.JoinRoom:output_type -> chat_v4.JoinRoomReply
	0, // 7: chat_v4.Chat.SendMessage:output_type -> chat_v4.ChatMessage
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chat_v4_chat_proto_init() }
func file_chat_v4_chat_proto_init() {
	if File_chat_v4_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v4_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_chat_v4_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRequest); i {
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
		file_chat_v4_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomReply); i {
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
		file_chat_v4_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_chat_v4_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_chat_v4_chat_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ChatRequest_JoinRoomRequest)(nil),
		(*ChatRequest_ChatMessage)(nil),
	}
	file_chat_v4_chat_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ChatResponse_JoinRoomReply)(nil),
		(*ChatResponse_ChatMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_v4_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_v4_chat_proto_goTypes,
		DependencyIndexes: file_chat_v4_chat_proto_depIdxs,
		MessageInfos:      file_chat_v4_chat_proto_msgTypes,
	}.Build()
	File_chat_v4_chat_proto = out.File
	file_chat_v4_chat_proto_rawDesc = nil
	file_chat_v4_chat_proto_goTypes = nil
	file_chat_v4_chat_proto_depIdxs = nil
}
