// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: common/common.proto

package common

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
	MsgId_MSG_0                MsgId = 0
	MsgId_CREATE_ROOM          MsgId = 2001
	MsgId_JOIN_ROOM            MsgId = 2002
	MsgId_LEVE_ROOM            MsgId = 2003
	MsgId_DIS_ROOM             MsgId = 2004
	MsgId_MSG_CONNECT_DEV_Q    MsgId = 400 //来自各种协议连接
	MsgId_MSG_CONNECT_DEV_A    MsgId = 401
	MsgId_MSG_CONNECT_SERVER_Q MsgId = 402
	MsgId_MSG_CONNECT_SERVER_A MsgId = 403
	MsgId_MSG_SEND_DATA        MsgId = 404
	MsgId_MSG_CLOSE_SERVER     MsgId = 405
	MsgId_MSG_CLEAR_DEV_Q      MsgId = 406
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0:    "MSG_0",
		2001: "CREATE_ROOM",
		2002: "JOIN_ROOM",
		2003: "LEVE_ROOM",
		2004: "DIS_ROOM",
		400:  "MSG_CONNECT_DEV_Q",
		401:  "MSG_CONNECT_DEV_A",
		402:  "MSG_CONNECT_SERVER_Q",
		403:  "MSG_CONNECT_SERVER_A",
		404:  "MSG_SEND_DATA",
		405:  "MSG_CLOSE_SERVER",
		406:  "MSG_CLEAR_DEV_Q",
	}
	MsgId_value = map[string]int32{
		"MSG_0":                0,
		"CREATE_ROOM":          2001,
		"JOIN_ROOM":            2002,
		"LEVE_ROOM":            2003,
		"DIS_ROOM":             2004,
		"MSG_CONNECT_DEV_Q":    400,
		"MSG_CONNECT_DEV_A":    401,
		"MSG_CONNECT_SERVER_Q": 402,
		"MSG_CONNECT_SERVER_A": 403,
		"MSG_SEND_DATA":        404,
		"MSG_CLOSE_SERVER":     405,
		"MSG_CLEAR_DEV_Q":      406,
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
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type CommonMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	MsgId MsgId  `protobuf:"varint,3,opt,name=MsgId,proto3,enum=common.MsgId" json:"MsgId,omitempty"`
	Data  []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CommonMsg) Reset() {
	*x = CommonMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonMsg) ProtoMessage() {}

func (x *CommonMsg) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonMsg.ProtoReflect.Descriptor instead.
func (*CommonMsg) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonMsg) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonMsg) GetMsgId() MsgId {
	if x != nil {
		return x.MsgId
	}
	return MsgId_MSG_0
}

func (x *CommonMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x6a, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x23, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x52, 0x05,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xfa, 0x01, 0x0a, 0x05, 0x4d, 0x73,
	0x67, 0x49, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x53, 0x47, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd1, 0x0f,
	0x12, 0x0e, 0x0a, 0x09, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd2, 0x0f,
	0x12, 0x0e, 0x0a, 0x09, 0x4c, 0x45, 0x56, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd3, 0x0f,
	0x12, 0x0d, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xd4, 0x0f, 0x12,
	0x16, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x44,
	0x45, 0x56, 0x5f, 0x51, 0x10, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x5f, 0x41, 0x10, 0x91, 0x03, 0x12,
	0x19, 0x0a, 0x14, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x51, 0x10, 0x92, 0x03, 0x12, 0x19, 0x0a, 0x14, 0x4d, 0x53,
	0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x5f, 0x41, 0x10, 0x93, 0x03, 0x12, 0x12, 0x0a, 0x0d, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x45, 0x4e,
	0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x94, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x53, 0x47,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x95, 0x03,
	0x12, 0x14, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x44, 0x45,
	0x56, 0x5f, 0x51, 0x10, 0x96, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_common_proto_goTypes = []interface{}{
	(MsgId)(0),        // 0: common.MsgId
	(*CommonMsg)(nil), // 1: common.CommonMsg
}
var file_common_common_proto_depIdxs = []int32{
	0, // 0: common.CommonMsg.MsgId:type_name -> common.MsgId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonMsg); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
