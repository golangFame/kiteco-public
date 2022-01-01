// Proto messages used by class_registration_test.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/util/class_registration_test.proto

package serving

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Config1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

func (x *Config1) Reset() {
	*x = Config1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config1) ProtoMessage() {}

func (x *Config1) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config1.ProtoReflect.Descriptor instead.
func (*Config1) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_util_class_registration_test_proto_rawDescGZIP(), []int{0}
}

func (x *Config1) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

type Config2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
}

func (x *Config2) Reset() {
	*x = Config2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config2) ProtoMessage() {}

func (x *Config2) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config2.ProtoReflect.Descriptor instead.
func (*Config2) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_util_class_registration_test_proto_rawDescGZIP(), []int{1}
}

func (x *Config2) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

type MessageWithAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyField *any.Any `protobuf:"bytes,1,opt,name=any_field,json=anyField,proto3" json:"any_field,omitempty"`
}

func (x *MessageWithAny) Reset() {
	*x = MessageWithAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithAny) ProtoMessage() {}

func (x *MessageWithAny) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_util_class_registration_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithAny.ProtoReflect.Descriptor instead.
func (*MessageWithAny) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_util_class_registration_test_proto_rawDescGZIP(), []int{2}
}

func (x *MessageWithAny) GetAnyField() *any.Any {
	if x != nil {
		return x.AnyField
	}
	return nil
}

var File_tensorflow_serving_util_class_registration_test_proto protoreflect.FileDescriptor

var file_tensorflow_serving_util_class_registration_test_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x31, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x22, 0x2c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0x43, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x41, 0x6e, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x61,
	0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_util_class_registration_test_proto_rawDescOnce sync.Once
	file_tensorflow_serving_util_class_registration_test_proto_rawDescData = file_tensorflow_serving_util_class_registration_test_proto_rawDesc
)

func file_tensorflow_serving_util_class_registration_test_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_util_class_registration_test_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_util_class_registration_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_util_class_registration_test_proto_rawDescData)
	})
	return file_tensorflow_serving_util_class_registration_test_proto_rawDescData
}

var file_tensorflow_serving_util_class_registration_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_util_class_registration_test_proto_goTypes = []interface{}{
	(*Config1)(nil),        // 0: tensorflow.serving.Config1
	(*Config2)(nil),        // 1: tensorflow.serving.Config2
	(*MessageWithAny)(nil), // 2: tensorflow.serving.MessageWithAny
	(*any.Any)(nil),        // 3: google.protobuf.Any
}
var file_tensorflow_serving_util_class_registration_test_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.MessageWithAny.any_field:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_util_class_registration_test_proto_init() }
func file_tensorflow_serving_util_class_registration_test_proto_init() {
	if File_tensorflow_serving_util_class_registration_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_util_class_registration_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config1); i {
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
		file_tensorflow_serving_util_class_registration_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config2); i {
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
		file_tensorflow_serving_util_class_registration_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithAny); i {
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
			RawDescriptor: file_tensorflow_serving_util_class_registration_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_util_class_registration_test_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_util_class_registration_test_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_util_class_registration_test_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_util_class_registration_test_proto = out.File
	file_tensorflow_serving_util_class_registration_test_proto_rawDesc = nil
	file_tensorflow_serving_util_class_registration_test_proto_goTypes = nil
	file_tensorflow_serving_util_class_registration_test_proto_depIdxs = nil
}
