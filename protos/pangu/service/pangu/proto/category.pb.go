// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: pangu/proto/category.proto

package proto

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

type CategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       bool      `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	CategoryInfo *Category `protobuf:"bytes,2,opt,name=CategoryInfo,json=category,proto3" json:"CategoryInfo,omitempty"`
	Data         *Category `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_pangu_proto_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CategoryResponse) GetCategoryInfo() *Category {
	if x != nil {
		return x.CategoryInfo
	}
	return nil
}

func (x *CategoryResponse) GetData() *Category {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pangu_proto_category_proto protoreflect.FileDescriptor

var file_pangu_proto_category_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61,
	0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x70, 0x61, 0x6e, 0x67,
	0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e,
	0x01, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x38, 0x5a, 0x36, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2e, 0x66, 0x69, 0x74, 0x2f,
	0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61,
	0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pangu_proto_category_proto_rawDescOnce sync.Once
	file_pangu_proto_category_proto_rawDescData = file_pangu_proto_category_proto_rawDesc
)

func file_pangu_proto_category_proto_rawDescGZIP() []byte {
	file_pangu_proto_category_proto_rawDescOnce.Do(func() {
		file_pangu_proto_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_category_proto_rawDescData)
	})
	return file_pangu_proto_category_proto_rawDescData
}

var file_pangu_proto_category_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pangu_proto_category_proto_goTypes = []interface{}{
	(*CategoryRequest)(nil),  // 0: panguservice.CategoryRequest
	(*CategoryResponse)(nil), // 1: panguservice.CategoryResponse
	(*Category)(nil),         // 2: panguservice.Category
}
var file_pangu_proto_category_proto_depIdxs = []int32{
	2, // 0: panguservice.CategoryResponse.CategoryInfo:type_name -> panguservice.Category
	2, // 1: panguservice.CategoryResponse.Data:type_name -> panguservice.Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pangu_proto_category_proto_init() }
func file_pangu_proto_category_proto_init() {
	if File_pangu_proto_category_proto != nil {
		return
	}
	file_pangu_proto_category_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryRequest); i {
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
		file_pangu_proto_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
			RawDescriptor: file_pangu_proto_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_category_proto_goTypes,
		DependencyIndexes: file_pangu_proto_category_proto_depIdxs,
		MessageInfos:      file_pangu_proto_category_proto_msgTypes,
	}.Build()
	File_pangu_proto_category_proto = out.File
	file_pangu_proto_category_proto_rawDesc = nil
	file_pangu_proto_category_proto_goTypes = nil
	file_pangu_proto_category_proto_depIdxs = nil
}
