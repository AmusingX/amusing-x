// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: category_delete.proto

package charonservice

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

type CategoryDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CategoryDeleteRequest) Reset() {
	*x = CategoryDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDeleteRequest) ProtoMessage() {}

func (x *CategoryDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDeleteRequest.ProtoReflect.Descriptor instead.
func (*CategoryDeleteRequest) Descriptor() ([]byte, []int) {
	return file_category_delete_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryDeleteRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CategoryDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *CategoryDeleteResponse) Reset() {
	*x = CategoryDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDeleteResponse) ProtoMessage() {}

func (x *CategoryDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDeleteResponse.ProtoReflect.Descriptor instead.
func (*CategoryDeleteResponse) Descriptor() ([]byte, []int) {
	return file_category_delete_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryDeleteResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

var File_category_delete_proto protoreflect.FileDescriptor

var file_category_delete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x42, 0x24, 0x5a, 0x22, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_category_delete_proto_rawDescOnce sync.Once
	file_category_delete_proto_rawDescData = file_category_delete_proto_rawDesc
)

func file_category_delete_proto_rawDescGZIP() []byte {
	file_category_delete_proto_rawDescOnce.Do(func() {
		file_category_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_delete_proto_rawDescData)
	})
	return file_category_delete_proto_rawDescData
}

var file_category_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_category_delete_proto_goTypes = []interface{}{
	(*CategoryDeleteRequest)(nil),  // 0: charonservice.CategoryDeleteRequest
	(*CategoryDeleteResponse)(nil), // 1: charonservice.CategoryDeleteResponse
}
var file_category_delete_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_category_delete_proto_init() }
func file_category_delete_proto_init() {
	if File_category_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDeleteRequest); i {
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
		file_category_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDeleteResponse); i {
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
			RawDescriptor: file_category_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_category_delete_proto_goTypes,
		DependencyIndexes: file_category_delete_proto_depIdxs,
		MessageInfos:      file_category_delete_proto_msgTypes,
	}.Build()
	File_category_delete_proto = out.File
	file_category_delete_proto_rawDesc = nil
	file_category_delete_proto_goTypes = nil
	file_category_delete_proto_depIdxs = nil
}
