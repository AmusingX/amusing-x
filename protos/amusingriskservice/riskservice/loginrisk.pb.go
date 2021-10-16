// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: loginrisk.proto

package riskservice

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

type LoginRiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	StrategyType string `protobuf:"bytes,2,opt,name=strategy_type,proto3" json:"strategy_type,omitempty"`
	Phone        string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Action       string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *LoginRiskRequest) Reset() {
	*x = LoginRiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginrisk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRiskRequest) ProtoMessage() {}

func (x *LoginRiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loginrisk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRiskRequest.ProtoReflect.Descriptor instead.
func (*LoginRiskRequest) Descriptor() ([]byte, []int) {
	return file_loginrisk_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRiskRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LoginRiskRequest) GetStrategyType() string {
	if x != nil {
		return x.StrategyType
	}
	return ""
}

func (x *LoginRiskRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginRiskRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type LoginRiskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LoginRiskReply) Reset() {
	*x = LoginRiskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginrisk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRiskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRiskReply) ProtoMessage() {}

func (x *LoginRiskReply) ProtoReflect() protoreflect.Message {
	mi := &file_loginrisk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRiskReply.ProtoReflect.Descriptor instead.
func (*LoginRiskReply) Descriptor() ([]byte, []int) {
	return file_loginrisk_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRiskReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_loginrisk_proto protoreflect.FileDescriptor

var file_loginrisk_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x80,
	0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x28, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x61,
	0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x73,
	0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loginrisk_proto_rawDescOnce sync.Once
	file_loginrisk_proto_rawDescData = file_loginrisk_proto_rawDesc
)

func file_loginrisk_proto_rawDescGZIP() []byte {
	file_loginrisk_proto_rawDescOnce.Do(func() {
		file_loginrisk_proto_rawDescData = protoimpl.X.CompressGZIP(file_loginrisk_proto_rawDescData)
	})
	return file_loginrisk_proto_rawDescData
}

var file_loginrisk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loginrisk_proto_goTypes = []interface{}{
	(*LoginRiskRequest)(nil), // 0: riskservice.LoginRiskRequest
	(*LoginRiskReply)(nil),   // 1: riskservice.LoginRiskReply
}
var file_loginrisk_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loginrisk_proto_init() }
func file_loginrisk_proto_init() {
	if File_loginrisk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loginrisk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRiskRequest); i {
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
		file_loginrisk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRiskReply); i {
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
			RawDescriptor: file_loginrisk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loginrisk_proto_goTypes,
		DependencyIndexes: file_loginrisk_proto_depIdxs,
		MessageInfos:      file_loginrisk_proto_msgTypes,
	}.Build()
	File_loginrisk_proto = out.File
	file_loginrisk_proto_rawDesc = nil
	file_loginrisk_proto_goTypes = nil
	file_loginrisk_proto_depIdxs = nil
}
