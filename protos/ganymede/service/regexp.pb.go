// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: regexp.proto

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

type RegexpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regexps []*Regexp `protobuf:"bytes,1,rep,name=regexps,proto3" json:"regexps,omitempty"`
}

func (x *RegexpResponse) Reset() {
	*x = RegexpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regexp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexpResponse) ProtoMessage() {}

func (x *RegexpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_regexp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexpResponse.ProtoReflect.Descriptor instead.
func (*RegexpResponse) Descriptor() ([]byte, []int) {
	return file_regexp_proto_rawDescGZIP(), []int{0}
}

func (x *RegexpResponse) GetRegexps() []*Regexp {
	if x != nil {
		return x.Regexps
	}
	return nil
}

type Regexp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rules string `protobuf:"bytes,3,opt,name=rules,proto3" json:"rules,omitempty"`
	Desc  string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Regexp) Reset() {
	*x = Regexp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regexp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regexp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regexp) ProtoMessage() {}

func (x *Regexp) ProtoReflect() protoreflect.Message {
	mi := &file_regexp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regexp.ProtoReflect.Descriptor instead.
func (*Regexp) Descriptor() ([]byte, []int) {
	return file_regexp_proto_rawDescGZIP(), []int{1}
}

func (x *Regexp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Regexp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Regexp) GetRules() string {
	if x != nil {
		return x.Rules
	}
	return ""
}

func (x *Regexp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regexp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_regexp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_regexp_proto_rawDescGZIP(), []int{2}
}

func (x *PongResponse) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

var File_regexp_proto protoreflect.FileDescriptor

var file_regexp_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6e,
	0x79, 0x6d, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x52, 0x07, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x70, 0x73, 0x22, 0x56, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x2e, 0x0a, 0x0c,
	0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x28, 0x5a, 0x26,
	0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61,
	0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regexp_proto_rawDescOnce sync.Once
	file_regexp_proto_rawDescData = file_regexp_proto_rawDesc
)

func file_regexp_proto_rawDescGZIP() []byte {
	file_regexp_proto_rawDescOnce.Do(func() {
		file_regexp_proto_rawDescData = protoimpl.X.CompressGZIP(file_regexp_proto_rawDescData)
	})
	return file_regexp_proto_rawDescData
}

var file_regexp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_regexp_proto_goTypes = []interface{}{
	(*RegexpResponse)(nil), // 0: ganymde.RegexpResponse
	(*Regexp)(nil),         // 1: ganymde.Regexp
	(*PongResponse)(nil),   // 2: ganymde.PongResponse
}
var file_regexp_proto_depIdxs = []int32{
	1, // 0: ganymde.RegexpResponse.regexps:type_name -> ganymde.Regexp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_regexp_proto_init() }
func file_regexp_proto_init() {
	if File_regexp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regexp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexpResponse); i {
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
		file_regexp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regexp); i {
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
		file_regexp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
			RawDescriptor: file_regexp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regexp_proto_goTypes,
		DependencyIndexes: file_regexp_proto_depIdxs,
		MessageInfos:      file_regexp_proto_msgTypes,
	}.Build()
	File_regexp_proto = out.File
	file_regexp_proto_rawDesc = nil
	file_regexp_proto_goTypes = nil
	file_regexp_proto_depIdxs = nil
}
