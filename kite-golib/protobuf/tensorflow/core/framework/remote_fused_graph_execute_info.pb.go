// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

package framework

import (
	proto "github.com/golang/protobuf/proto"
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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph,proto3" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName,proto3" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName,proto3" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape,proto3" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape,proto3" json:"default_graph_output_tensor_shape,omitempty"`
}

func (x *RemoteFusedGraphExecuteInfo) Reset() {
	*x = RemoteFusedGraphExecuteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteFusedGraphExecuteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteFusedGraphExecuteInfo) ProtoMessage() {}

func (x *RemoteFusedGraphExecuteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteFusedGraphExecuteInfo.ProtoReflect.Descriptor instead.
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if x != nil {
		return x.RemoteGraph
	}
	return nil
}

func (x *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if x != nil {
		return x.GraphInputNodeName
	}
	return nil
}

func (x *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if x != nil {
		return x.GraphOutputNodeName
	}
	return nil
}

func (x *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if x != nil {
		return x.ExecutorName
	}
	return ""
}

func (x *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if x != nil {
		return x.SerializedExecutorParameters
	}
	return nil
}

func (x *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if x != nil {
		return x.DefaultGraphInputTensorShape
	}
	return nil
}

func (x *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if x != nil {
		return x.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtype DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
}

func (x *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*x = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}

func (x *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.ProtoReflect.Descriptor instead.
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if x != nil {
		return x.Dtype
	}
	return DataType_DT_INVALID
}

func (x *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

var File_tensorflow_core_framework_remote_fused_graph_execute_info_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x66, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x25, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x05, 0x0a, 0x1b, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x46, 0x75, 0x73, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x44, 0x65, 0x66, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x31, 0x0a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x44, 0x0a, 0x1e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x46, 0x75, 0x73, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x53, 0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x1c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x86, 0x01, 0x0a,
	0x21, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x75, 0x73, 0x65,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x1d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x53, 0x68, 0x61, 0x70, 0x65, 0x1a, 0x76, 0x0a, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a,
	0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x42, 0x80, 0x01,
	0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x20, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x46, 0x75, 0x73, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0xf8, 0x01, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescData = file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDesc
)

func file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescData)
	})
	return file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDescData
}

var file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_goTypes = []interface{}{
	(*RemoteFusedGraphExecuteInfo)(nil),                      // 0: tensorflow.RemoteFusedGraphExecuteInfo
	(*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), // 1: tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto
	(*GraphDef)(nil),         // 2: tensorflow.GraphDef
	(DataType)(0),            // 3: tensorflow.DataType
	(*TensorShapeProto)(nil), // 4: tensorflow.TensorShapeProto
}
var file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_depIdxs = []int32{
	2, // 0: tensorflow.RemoteFusedGraphExecuteInfo.remote_graph:type_name -> tensorflow.GraphDef
	1, // 1: tensorflow.RemoteFusedGraphExecuteInfo.default_graph_input_tensor_shape:type_name -> tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto
	1, // 2: tensorflow.RemoteFusedGraphExecuteInfo.default_graph_output_tensor_shape:type_name -> tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto
	3, // 3: tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto.dtype:type_name -> tensorflow.DataType
	4, // 4: tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto.shape:type_name -> tensorflow.TensorShapeProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_init() }
func file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_init() {
	if File_tensorflow_core_framework_remote_fused_graph_execute_info_proto != nil {
		return
	}
	file_tensorflow_core_framework_graph_proto_init()
	file_tensorflow_core_framework_tensor_shape_proto_init()
	file_tensorflow_core_framework_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteFusedGraphExecuteInfo); i {
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
		file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto); i {
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
			RawDescriptor: file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_remote_fused_graph_execute_info_proto = out.File
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_rawDesc = nil
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_goTypes = nil
	file_tensorflow_core_framework_remote_fused_graph_execute_info_proto_depIdxs = nil
}
