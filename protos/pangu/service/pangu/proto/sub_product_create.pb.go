// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: pangu/proto/sub_product_create.proto

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

type SubProductCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64   `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Desc        string  `protobuf:"bytes,3,opt,name=Desc,json=desc,proto3" json:"Desc,omitempty"`
	ProductId   int64   `protobuf:"varint,4,opt,name=ProductId,json=product_id,proto3" json:"ProductId,omitempty"`
	Currency    string  `protobuf:"bytes,5,opt,name=Currency,json=currency,proto3" json:"Currency,omitempty"`
	Price       int64   `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Stock       int64   `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	AttributeId []int64 `protobuf:"varint,8,rep,packed,name=AttributeId,json=attribute_id,proto3" json:"AttributeId,omitempty"`
}

func (x *SubProductCreateRequest) Reset() {
	*x = SubProductCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_sub_product_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductCreateRequest) ProtoMessage() {}

func (x *SubProductCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_sub_product_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductCreateRequest.ProtoReflect.Descriptor instead.
func (*SubProductCreateRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_sub_product_create_proto_rawDescGZIP(), []int{0}
}

func (x *SubProductCreateRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *SubProductCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubProductCreateRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SubProductCreateRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SubProductCreateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SubProductCreateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SubProductCreateRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *SubProductCreateRequest) GetAttributeId() []int64 {
	if x != nil {
		return x.AttributeId
	}
	return nil
}

var File_pangu_proto_sub_product_create_proto protoreflect.FileDescriptor

var file_pangu_proto_sub_product_create_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x21, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x42, 0x38, 0x5a, 0x36, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2e, 0x66,
	0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pangu_proto_sub_product_create_proto_rawDescOnce sync.Once
	file_pangu_proto_sub_product_create_proto_rawDescData = file_pangu_proto_sub_product_create_proto_rawDesc
)

func file_pangu_proto_sub_product_create_proto_rawDescGZIP() []byte {
	file_pangu_proto_sub_product_create_proto_rawDescOnce.Do(func() {
		file_pangu_proto_sub_product_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_sub_product_create_proto_rawDescData)
	})
	return file_pangu_proto_sub_product_create_proto_rawDescData
}

var file_pangu_proto_sub_product_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pangu_proto_sub_product_create_proto_goTypes = []interface{}{
	(*SubProductCreateRequest)(nil), // 0: panguservice.SubProductCreateRequest
}
var file_pangu_proto_sub_product_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pangu_proto_sub_product_create_proto_init() }
func file_pangu_proto_sub_product_create_proto_init() {
	if File_pangu_proto_sub_product_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_sub_product_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductCreateRequest); i {
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
			RawDescriptor: file_pangu_proto_sub_product_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_sub_product_create_proto_goTypes,
		DependencyIndexes: file_pangu_proto_sub_product_create_proto_depIdxs,
		MessageInfos:      file_pangu_proto_sub_product_create_proto_msgTypes,
	}.Build()
	File_pangu_proto_sub_product_create_proto = out.File
	file_pangu_proto_sub_product_create_proto_rawDesc = nil
	file_pangu_proto_sub_product_create_proto_goTypes = nil
	file_pangu_proto_sub_product_create_proto_depIdxs = nil
}
