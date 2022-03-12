// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: private/private.proto

package private

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

type MsgId int32

const (
	MsgId_MSG_0       MsgId = 0
	MsgId_CREATE_ROOM MsgId = 2001
	MsgId_JOIN_ROOM   MsgId = 2002
	MsgId_LEVE_ROOM   MsgId = 2003
	MsgId_DIS_ROOM    MsgId = 2004
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0:    "MSG_0",
		2001: "CREATE_ROOM",
		2002: "JOIN_ROOM",
		2003: "LEVE_ROOM",
		2004: "DIS_ROOM",
	}
	MsgId_value = map[string]int32{
		"MSG_0":       0,
		"CREATE_ROOM": 2001,
		"JOIN_ROOM":   2002,
		"LEVE_ROOM":   2003,
		"DIS_ROOM":    2004,
	}
)

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}

func (x MsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_private_private_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_private_private_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{0}
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateType int32 `protobuf:"varint,1,opt,name=createType,proto3" json:"createType,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomRequest) GetCreateType() int32 {
	if x != nil {
		return x.CreateType
	}
	return 0
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID int32 `protobuf:"varint,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomResponse) GetRoomID() int32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID int32 `protobuf:"varint,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[2]
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
	return file_private_private_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRoomRequest) GetRoomID() int32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserAvtor string `protobuf:"bytes,3,opt,name=userAvtor,proto3" json:"userAvtor,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetUserAvtor() string {
	if x != nil {
		return x.UserAvtor
	}
	return ""
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRoomResponse) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type LeveRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID int32 `protobuf:"varint,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *LeveRoomRequest) Reset() {
	*x = LeveRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeveRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeveRoomRequest) ProtoMessage() {}

func (x *LeveRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeveRoomRequest.ProtoReflect.Descriptor instead.
func (*LeveRoomRequest) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{5}
}

func (x *LeveRoomRequest) GetRoomID() int32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type LeveRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *LeveRoomResponse) Reset() {
	*x = LeveRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeveRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeveRoomResponse) ProtoMessage() {}

func (x *LeveRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeveRoomResponse.ProtoReflect.Descriptor instead.
func (*LeveRoomResponse) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{6}
}

func (x *LeveRoomResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

type DisRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID int32 `protobuf:"varint,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *DisRoomRequest) Reset() {
	*x = DisRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisRoomRequest) ProtoMessage() {}

func (x *DisRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisRoomRequest.ProtoReflect.Descriptor instead.
func (*DisRoomRequest) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{7}
}

func (x *DisRoomRequest) GetRoomID() int32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type DisRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *DisRoomResponse) Reset() {
	*x = DisRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_private_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisRoomResponse) ProtoMessage() {}

func (x *DisRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_private_private_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisRoomResponse.ProtoReflect.Descriptor instead.
func (*DisRoomResponse) Descriptor() ([]byte, []int) {
	return file_private_private_proto_rawDescGZIP(), []int{8}
}

func (x *DisRoomResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_private_private_proto protoreflect.FileDescriptor

var file_private_private_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0x5c,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x74, 0x6f, 0x72, 0x22, 0x3b, 0x0a, 0x10,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x29, 0x0a, 0x0f, 0x4c, 0x65, 0x76,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x22, 0x24, 0x0a, 0x10, 0x4c, 0x65, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x73, 0x2a, 0x53, 0x0a, 0x05, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x53, 0x47, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd1, 0x0f, 0x12,
	0x0e, 0x0a, 0x09, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd2, 0x0f, 0x12,
	0x0e, 0x0a, 0x09, 0x4c, 0x45, 0x56, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd3, 0x0f, 0x12,
	0x0d, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd4, 0x0f, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_private_private_proto_rawDescOnce sync.Once
	file_private_private_proto_rawDescData = file_private_private_proto_rawDesc
)

func file_private_private_proto_rawDescGZIP() []byte {
	file_private_private_proto_rawDescOnce.Do(func() {
		file_private_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_private_proto_rawDescData)
	})
	return file_private_private_proto_rawDescData
}

var file_private_private_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_private_private_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_private_private_proto_goTypes = []interface{}{
	(MsgId)(0),                 // 0: private.MsgId
	(*CreateRoomRequest)(nil),  // 1: private.CreateRoomRequest
	(*CreateRoomResponse)(nil), // 2: private.CreateRoomResponse
	(*JoinRoomRequest)(nil),    // 3: private.JoinRoomRequest
	(*UserInfo)(nil),           // 4: private.UserInfo
	(*JoinRoomResponse)(nil),   // 5: private.JoinRoomResponse
	(*LeveRoomRequest)(nil),    // 6: private.LeveRoomRequest
	(*LeveRoomResponse)(nil),   // 7: private.LeveRoomResponse
	(*DisRoomRequest)(nil),     // 8: private.DisRoomRequest
	(*DisRoomResponse)(nil),    // 9: private.DisRoomResponse
}
var file_private_private_proto_depIdxs = []int32{
	4, // 0: private.JoinRoomResponse.users:type_name -> private.UserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_private_private_proto_init() }
func file_private_private_proto_init() {
	if File_private_private_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_private_private_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomResponse); i {
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
		file_private_private_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_private_private_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_private_private_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomResponse); i {
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
		file_private_private_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeveRoomRequest); i {
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
		file_private_private_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeveRoomResponse); i {
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
		file_private_private_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisRoomRequest); i {
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
		file_private_private_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisRoomResponse); i {
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
			RawDescriptor: file_private_private_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_private_proto_goTypes,
		DependencyIndexes: file_private_private_proto_depIdxs,
		EnumInfos:         file_private_private_proto_enumTypes,
		MessageInfos:      file_private_private_proto_msgTypes,
	}.Build()
	File_private_private_proto = out.File
	file_private_private_proto_rawDesc = nil
	file_private_private_proto_goTypes = nil
	file_private_private_proto_depIdxs = nil
}