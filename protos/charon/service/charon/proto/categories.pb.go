// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: charon/proto/categories.proto

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

type SearchFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []int64  `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Desc []string `protobuf:"bytes,3,rep,name=desc,proto3" json:"desc,omitempty"`
}

func (x *SearchFilter) Reset() {
	*x = SearchFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charon_proto_categories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilter) ProtoMessage() {}

func (x *SearchFilter) ProtoReflect() protoreflect.Message {
	mi := &file_charon_proto_categories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilter.ProtoReflect.Descriptor instead.
func (*SearchFilter) Descriptor() ([]byte, []int) {
	return file_charon_proto_categories_proto_rawDescGZIP(), []int{0}
}

func (x *SearchFilter) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchFilter) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *SearchFilter) GetDesc() []string {
	if x != nil {
		return x.Desc
	}
	return nil
}

var File_charon_proto_categories_proto protoreflect.FileDescriptor

var file_charon_proto_categories_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x46,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x42, 0x24, 0x5a, 0x22, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e,
	0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x63,
	0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_charon_proto_categories_proto_rawDescOnce sync.Once
	file_charon_proto_categories_proto_rawDescData = file_charon_proto_categories_proto_rawDesc
)

func file_charon_proto_categories_proto_rawDescGZIP() []byte {
	file_charon_proto_categories_proto_rawDescOnce.Do(func() {
		file_charon_proto_categories_proto_rawDescData = protoimpl.X.CompressGZIP(file_charon_proto_categories_proto_rawDescData)
	})
	return file_charon_proto_categories_proto_rawDescData
}

var file_charon_proto_categories_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_charon_proto_categories_proto_goTypes = []interface{}{
	(*SearchFilter)(nil), // 0: charonservice.SearchFilter
}
var file_charon_proto_categories_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_charon_proto_categories_proto_init() }
func file_charon_proto_categories_proto_init() {
	if File_charon_proto_categories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_charon_proto_categories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilter); i {
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
			RawDescriptor: file_charon_proto_categories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_charon_proto_categories_proto_goTypes,
		DependencyIndexes: file_charon_proto_categories_proto_depIdxs,
		MessageInfos:      file_charon_proto_categories_proto_msgTypes,
	}.Build()
	File_charon_proto_categories_proto = out.File
	file_charon_proto_categories_proto_rawDesc = nil
	file_charon_proto_categories_proto_goTypes = nil
	file_charon_proto_categories_proto_depIdxs = nil
}
