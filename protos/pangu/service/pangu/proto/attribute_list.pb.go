// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: pangu/proto/attribute_list.proto

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

type AttributeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=Query,json=query,proto3" json:"Query,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=Page,json=page,proto3" json:"Page,omitempty"`
	Limit int64  `protobuf:"varint,3,opt,name=Limit,json=limit,proto3" json:"Limit,omitempty"`
	Sort  string `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *AttributeListRequest) Reset() {
	*x = AttributeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_attribute_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeListRequest) ProtoMessage() {}

func (x *AttributeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_attribute_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeListRequest.ProtoReflect.Descriptor instead.
func (*AttributeListRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_attribute_list_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *AttributeListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AttributeListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AttributeListRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type AttributeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64        `protobuf:"varint,1,opt,name=Page,json=page,proto3" json:"Page,omitempty"`
	Limit   int64        `protobuf:"varint,2,opt,name=Limit,json=limit,proto3" json:"Limit,omitempty"`
	Total   int64        `protobuf:"varint,3,opt,name=Total,json=total,proto3" json:"Total,omitempty"`
	HasNext bool         `protobuf:"varint,4,opt,name=hasNext,json=has_next,proto3" json:"hasNext,omitempty"`
	Data    []*Attribute `protobuf:"bytes,6,rep,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (x *AttributeListResponse) Reset() {
	*x = AttributeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_attribute_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeListResponse) ProtoMessage() {}

func (x *AttributeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_attribute_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeListResponse.ProtoReflect.Descriptor instead.
func (*AttributeListResponse) Descriptor() ([]byte, []int) {
	return file_pangu_proto_attribute_list_proto_rawDescGZIP(), []int{1}
}

func (x *AttributeListResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AttributeListResponse) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AttributeListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AttributeListResponse) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

func (x *AttributeListResponse) GetData() []*Attribute {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pangu_proto_attribute_list_proto protoreflect.FileDescriptor

var file_pangu_proto_attribute_list_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1b, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a,
	0x14, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x15, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x38, 0x5a, 0x36, 0x61,
	0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2e, 0x66, 0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73,
	0x69, 0x6e, 0x67, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x6e, 0x67,
	0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pangu_proto_attribute_list_proto_rawDescOnce sync.Once
	file_pangu_proto_attribute_list_proto_rawDescData = file_pangu_proto_attribute_list_proto_rawDesc
)

func file_pangu_proto_attribute_list_proto_rawDescGZIP() []byte {
	file_pangu_proto_attribute_list_proto_rawDescOnce.Do(func() {
		file_pangu_proto_attribute_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_attribute_list_proto_rawDescData)
	})
	return file_pangu_proto_attribute_list_proto_rawDescData
}

var file_pangu_proto_attribute_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pangu_proto_attribute_list_proto_goTypes = []interface{}{
	(*AttributeListRequest)(nil),  // 0: panguservice.AttributeListRequest
	(*AttributeListResponse)(nil), // 1: panguservice.AttributeListResponse
	(*Attribute)(nil),             // 2: panguservice.Attribute
}
var file_pangu_proto_attribute_list_proto_depIdxs = []int32{
	2, // 0: panguservice.AttributeListResponse.Data:type_name -> panguservice.Attribute
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pangu_proto_attribute_list_proto_init() }
func file_pangu_proto_attribute_list_proto_init() {
	if File_pangu_proto_attribute_list_proto != nil {
		return
	}
	file_pangu_proto_attribute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_attribute_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeListRequest); i {
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
		file_pangu_proto_attribute_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeListResponse); i {
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
			RawDescriptor: file_pangu_proto_attribute_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_attribute_list_proto_goTypes,
		DependencyIndexes: file_pangu_proto_attribute_list_proto_depIdxs,
		MessageInfos:      file_pangu_proto_attribute_list_proto_msgTypes,
	}.Build()
	File_pangu_proto_attribute_list_proto = out.File
	file_pangu_proto_attribute_list_proto_rawDesc = nil
	file_pangu_proto_attribute_list_proto_goTypes = nil
	file_pangu_proto_attribute_list_proto_depIdxs = nil
}