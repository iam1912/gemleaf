// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: leaf.proto

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

type CreateLeafReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc  string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	MaxId uint64 `protobuf:"varint,3,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	Step  uint64 `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *CreateLeafReq) Reset() {
	*x = CreateLeafReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLeafReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLeafReq) ProtoMessage() {}

func (x *CreateLeafReq) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLeafReq.ProtoReflect.Descriptor instead.
func (*CreateLeafReq) Descriptor() ([]byte, []int) {
	return file_leaf_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLeafReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLeafReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateLeafReq) GetMaxId() uint64 {
	if x != nil {
		return x.MaxId
	}
	return 0
}

func (x *CreateLeafReq) GetStep() uint64 {
	if x != nil {
		return x.Step
	}
	return 0
}

type CreateLeafResq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateLeafResq) Reset() {
	*x = CreateLeafResq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLeafResq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLeafResq) ProtoMessage() {}

func (x *CreateLeafResq) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLeafResq.ProtoReflect.Descriptor instead.
func (*CreateLeafResq) Descriptor() ([]byte, []int) {
	return file_leaf_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLeafResq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLeafReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetLeafReq) Reset() {
	*x = GetLeafReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeafReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeafReq) ProtoMessage() {}

func (x *GetLeafReq) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeafReq.ProtoReflect.Descriptor instead.
func (*GetLeafReq) Descriptor() ([]byte, []int) {
	return file_leaf_proto_rawDescGZIP(), []int{2}
}

func (x *GetLeafReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetLeafResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetLeafResp) Reset() {
	*x = GetLeafResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeafResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeafResp) ProtoMessage() {}

func (x *GetLeafResp) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeafResp.ProtoReflect.Descriptor instead.
func (*GetLeafResp) Descriptor() ([]byte, []int) {
	return file_leaf_proto_rawDescGZIP(), []int{3}
}

func (x *GetLeafResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_leaf_proto protoreflect.FileDescriptor

var file_leaf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61,
	0x66, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x61,
	0x78, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x73, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x79,
	0x0a, 0x04, 0x4c, 0x65, 0x61, 0x66, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x66, 0x49, 0x44, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x61, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x66, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x61, 0x66, 0x52, 0x65, 0x73, 0x71, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaf_proto_rawDescOnce sync.Once
	file_leaf_proto_rawDescData = file_leaf_proto_rawDesc
)

func file_leaf_proto_rawDescGZIP() []byte {
	file_leaf_proto_rawDescOnce.Do(func() {
		file_leaf_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaf_proto_rawDescData)
	})
	return file_leaf_proto_rawDescData
}

var file_leaf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_leaf_proto_goTypes = []interface{}{
	(*CreateLeafReq)(nil),  // 0: proto.CreateLeafReq
	(*CreateLeafResq)(nil), // 1: proto.CreateLeafResq
	(*GetLeafReq)(nil),     // 2: proto.GetLeafReq
	(*GetLeafResp)(nil),    // 3: proto.GetLeafResp
}
var file_leaf_proto_depIdxs = []int32{
	2, // 0: proto.Leaf.GetLeafID:input_type -> proto.GetLeafReq
	0, // 1: proto.Leaf.CreateLeaf:input_type -> proto.CreateLeafReq
	3, // 2: proto.Leaf.GetLeafID:output_type -> proto.GetLeafResp
	1, // 3: proto.Leaf.CreateLeaf:output_type -> proto.CreateLeafResq
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_leaf_proto_init() }
func file_leaf_proto_init() {
	if File_leaf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leaf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLeafReq); i {
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
		file_leaf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLeafResq); i {
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
		file_leaf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeafReq); i {
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
		file_leaf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeafResp); i {
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
			RawDescriptor: file_leaf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_leaf_proto_goTypes,
		DependencyIndexes: file_leaf_proto_depIdxs,
		MessageInfos:      file_leaf_proto_msgTypes,
	}.Build()
	File_leaf_proto = out.File
	file_leaf_proto_rawDesc = nil
	file_leaf_proto_goTypes = nil
	file_leaf_proto_depIdxs = nil
}
