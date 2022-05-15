// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: europa/proto/sub_product.proto

package proto

import (
	proto "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
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

type SubProductDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubProductInfo *proto.SubProduct                `protobuf:"bytes,1,opt,name=subProductInfo,proto3" json:"subProductInfo,omitempty"`
	Attribute      []*proto.AttributeWithSubProduct `protobuf:"bytes,2,rep,name=Attribute,proto3" json:"Attribute,omitempty"`
}

func (x *SubProductDetail) Reset() {
	*x = SubProductDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_europa_proto_sub_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductDetail) ProtoMessage() {}

func (x *SubProductDetail) ProtoReflect() protoreflect.Message {
	mi := &file_europa_proto_sub_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductDetail.ProtoReflect.Descriptor instead.
func (*SubProductDetail) Descriptor() ([]byte, []int) {
	return file_europa_proto_sub_product_proto_rawDescGZIP(), []int{0}
}

func (x *SubProductDetail) GetSubProductInfo() *proto.SubProduct {
	if x != nil {
		return x.SubProductInfo
	}
	return nil
}

func (x *SubProductDetail) GetAttribute() []*proto.AttributeWithSubProduct {
	if x != nil {
		return x.Attribute
	}
	return nil
}

var File_europa_proto_sub_product_proto protoreflect.FileDescriptor

var file_europa_proto_sub_product_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a,
	0x10, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x61, 0x6e, 0x67,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x09, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x61, 0x6d, 0x75, 0x73,
	0x69, 0x6e, 0x67, 0x78, 0x2e, 0x66, 0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_europa_proto_sub_product_proto_rawDescOnce sync.Once
	file_europa_proto_sub_product_proto_rawDescData = file_europa_proto_sub_product_proto_rawDesc
)

func file_europa_proto_sub_product_proto_rawDescGZIP() []byte {
	file_europa_proto_sub_product_proto_rawDescOnce.Do(func() {
		file_europa_proto_sub_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_europa_proto_sub_product_proto_rawDescData)
	})
	return file_europa_proto_sub_product_proto_rawDescData
}

var file_europa_proto_sub_product_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_europa_proto_sub_product_proto_goTypes = []interface{}{
	(*SubProductDetail)(nil),              // 0: SubProductDetail
	(*proto.SubProduct)(nil),              // 1: panguservice.SubProduct
	(*proto.AttributeWithSubProduct)(nil), // 2: panguservice.AttributeWithSubProduct
}
var file_europa_proto_sub_product_proto_depIdxs = []int32{
	1, // 0: SubProductDetail.subProductInfo:type_name -> panguservice.SubProduct
	2, // 1: SubProductDetail.Attribute:type_name -> panguservice.AttributeWithSubProduct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_europa_proto_sub_product_proto_init() }
func file_europa_proto_sub_product_proto_init() {
	if File_europa_proto_sub_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_europa_proto_sub_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductDetail); i {
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
			RawDescriptor: file_europa_proto_sub_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_europa_proto_sub_product_proto_goTypes,
		DependencyIndexes: file_europa_proto_sub_product_proto_depIdxs,
		MessageInfos:      file_europa_proto_sub_product_proto_msgTypes,
	}.Build()
	File_europa_proto_sub_product_proto = out.File
	file_europa_proto_sub_product_proto_rawDesc = nil
	file_europa_proto_sub_product_proto_goTypes = nil
	file_europa_proto_sub_product_proto_depIdxs = nil
}
