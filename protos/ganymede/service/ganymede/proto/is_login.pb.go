// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: ganymede/proto/is_login.proto

package ganymedeservice

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

type IsLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *IsLoginRequest) Reset() {
	*x = IsLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ganymede_proto_is_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLoginRequest) ProtoMessage() {}

func (x *IsLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ganymede_proto_is_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLoginRequest.ProtoReflect.Descriptor instead.
func (*IsLoginRequest) Descriptor() ([]byte, []int) {
	return file_ganymede_proto_is_login_proto_rawDescGZIP(), []int{0}
}

func (x *IsLoginRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type IsLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    bool      `protobuf:"varint,1,opt,name=Login,proto3" json:"Login,omitempty"`
	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *IsLoginResponse) Reset() {
	*x = IsLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ganymede_proto_is_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLoginResponse) ProtoMessage() {}

func (x *IsLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ganymede_proto_is_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLoginResponse.ProtoReflect.Descriptor instead.
func (*IsLoginResponse) Descriptor() ([]byte, []int) {
	return file_ganymede_proto_is_login_proto_rawDescGZIP(), []int{1}
}

func (x *IsLoginResponse) GetLogin() bool {
	if x != nil {
		return x.Login
	}
	return false
}

func (x *IsLoginResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_ganymede_proto_is_login_proto protoreflect.FileDescriptor

var file_ganymede_proto_is_login_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x1a, 0x20, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65,
	0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0e, 0x49, 0x73,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x0f, 0x49, 0x73,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x26, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x67, 0x61,
	0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ganymede_proto_is_login_proto_rawDescOnce sync.Once
	file_ganymede_proto_is_login_proto_rawDescData = file_ganymede_proto_is_login_proto_rawDesc
)

func file_ganymede_proto_is_login_proto_rawDescGZIP() []byte {
	file_ganymede_proto_is_login_proto_rawDescOnce.Do(func() {
		file_ganymede_proto_is_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_ganymede_proto_is_login_proto_rawDescData)
	})
	return file_ganymede_proto_is_login_proto_rawDescData
}

var file_ganymede_proto_is_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ganymede_proto_is_login_proto_goTypes = []interface{}{
	(*IsLoginRequest)(nil),  // 0: ganymde.IsLoginRequest
	(*IsLoginResponse)(nil), // 1: ganymde.IsLoginResponse
	(*UserInfo)(nil),        // 2: ganymde.UserInfo
}
var file_ganymede_proto_is_login_proto_depIdxs = []int32{
	2, // 0: ganymde.IsLoginResponse.user_info:type_name -> ganymde.UserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ganymede_proto_is_login_proto_init() }
func file_ganymede_proto_is_login_proto_init() {
	if File_ganymede_proto_is_login_proto != nil {
		return
	}
	file_ganymede_proto_oauth_login_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ganymede_proto_is_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLoginRequest); i {
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
		file_ganymede_proto_is_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLoginResponse); i {
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
			RawDescriptor: file_ganymede_proto_is_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ganymede_proto_is_login_proto_goTypes,
		DependencyIndexes: file_ganymede_proto_is_login_proto_depIdxs,
		MessageInfos:      file_ganymede_proto_is_login_proto_msgTypes,
	}.Build()
	File_ganymede_proto_is_login_proto = out.File
	file_ganymede_proto_is_login_proto_rawDesc = nil
	file_ganymede_proto_is_login_proto_goTypes = nil
	file_ganymede_proto_is_login_proto_depIdxs = nil
}
