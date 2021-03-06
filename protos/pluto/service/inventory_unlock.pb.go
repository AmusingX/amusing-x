// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: inventory_unlock.proto

package plutoservice

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

type InventoryUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Count int64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Obj   CacheObjType `protobuf:"varint,3,opt,name=obj,proto3,enum=plutoservice.CacheObjType" json:"obj,omitempty"`
}

func (x *InventoryUnlockRequest) Reset() {
	*x = InventoryUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_unlock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryUnlockRequest) ProtoMessage() {}

func (x *InventoryUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_unlock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryUnlockRequest.ProtoReflect.Descriptor instead.
func (*InventoryUnlockRequest) Descriptor() ([]byte, []int) {
	return file_inventory_unlock_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryUnlockRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InventoryUnlockRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *InventoryUnlockRequest) GetObj() CacheObjType {
	if x != nil {
		return x.Obj
	}
	return CacheObjType_Book
}

type InventoryUnlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed            bool   `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Message            string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RemainingInventory int64  `protobuf:"varint,3,opt,name=RemainingInventory,proto3" json:"RemainingInventory,omitempty"`
}

func (x *InventoryUnlockResponse) Reset() {
	*x = InventoryUnlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_unlock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryUnlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryUnlockResponse) ProtoMessage() {}

func (x *InventoryUnlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_unlock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryUnlockResponse.ProtoReflect.Descriptor instead.
func (*InventoryUnlockResponse) Descriptor() ([]byte, []int) {
	return file_inventory_unlock_proto_rawDescGZIP(), []int{1}
}

func (x *InventoryUnlockResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *InventoryUnlockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InventoryUnlockResponse) GetRemainingInventory() int64 {
	if x != nil {
		return x.RemainingInventory
	}
	return 0
}

var File_inventory_unlock_proto protoreflect.FileDescriptor

var file_inventory_unlock_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1a, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6f, 0x62, 0x6a,
	0x22, 0x7d, 0x0a, 0x17, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x52, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x42,
	0x22, 0x5a, 0x20, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_unlock_proto_rawDescOnce sync.Once
	file_inventory_unlock_proto_rawDescData = file_inventory_unlock_proto_rawDesc
)

func file_inventory_unlock_proto_rawDescGZIP() []byte {
	file_inventory_unlock_proto_rawDescOnce.Do(func() {
		file_inventory_unlock_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_unlock_proto_rawDescData)
	})
	return file_inventory_unlock_proto_rawDescData
}

var file_inventory_unlock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_unlock_proto_goTypes = []interface{}{
	(*InventoryUnlockRequest)(nil),  // 0: plutoservice.InventoryUnlockRequest
	(*InventoryUnlockResponse)(nil), // 1: plutoservice.InventoryUnlockResponse
	(CacheObjType)(0),               // 2: plutoservice.CacheObjType
}
var file_inventory_unlock_proto_depIdxs = []int32{
	2, // 0: plutoservice.InventoryUnlockRequest.obj:type_name -> plutoservice.CacheObjType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_unlock_proto_init() }
func file_inventory_unlock_proto_init() {
	if File_inventory_unlock_proto != nil {
		return
	}
	file_inventory_cache_init_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inventory_unlock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryUnlockRequest); i {
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
		file_inventory_unlock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryUnlockResponse); i {
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
			RawDescriptor: file_inventory_unlock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventory_unlock_proto_goTypes,
		DependencyIndexes: file_inventory_unlock_proto_depIdxs,
		MessageInfos:      file_inventory_unlock_proto_msgTypes,
	}.Build()
	File_inventory_unlock_proto = out.File
	file_inventory_unlock_proto_rawDesc = nil
	file_inventory_unlock_proto_goTypes = nil
	file_inventory_unlock_proto_depIdxs = nil
}
