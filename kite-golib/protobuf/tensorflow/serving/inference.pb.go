// This file contains messages for various machine learning inferences
// such as regression and classification.
//
// In many applications more than one type of inference is desired for a single
// input.  For example, given meteorologic data an application may want to
// perform a classification to determine if we should expect rain, snow or sun
// and also perform a regression to predict the temperature.
// Sharing the single input data between two inference tasks can be accomplished
// using MultiInferenceRequest and MultiInferenceResponse.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/apis/inference.proto

package serving

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

// Inference request such as classification, regression, etc...
type InferenceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	// All ModelSpecs in a MultiInferenceRequest must access the same model name.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Signature's method_name. Should be one of the method names defined in
	// third_party/tensorflow/python/saved_model/signature_constants.py.
	// e.g. "tensorflow/serving/classify".
	MethodName string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (x *InferenceTask) Reset() {
	*x = InferenceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceTask) ProtoMessage() {}

func (x *InferenceTask) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceTask.ProtoReflect.Descriptor instead.
func (*InferenceTask) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_inference_proto_rawDescGZIP(), []int{0}
}

func (x *InferenceTask) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *InferenceTask) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

// Inference result, matches the type of request or is an error.
type InferenceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Types that are assignable to Result:
	//	*InferenceResult_ClassificationResult
	//	*InferenceResult_RegressionResult
	Result isInferenceResult_Result `protobuf_oneof:"result"`
}

func (x *InferenceResult) Reset() {
	*x = InferenceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResult) ProtoMessage() {}

func (x *InferenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResult.ProtoReflect.Descriptor instead.
func (*InferenceResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_inference_proto_rawDescGZIP(), []int{1}
}

func (x *InferenceResult) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (m *InferenceResult) GetResult() isInferenceResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *InferenceResult) GetClassificationResult() *ClassificationResult {
	if x, ok := x.GetResult().(*InferenceResult_ClassificationResult); ok {
		return x.ClassificationResult
	}
	return nil
}

func (x *InferenceResult) GetRegressionResult() *RegressionResult {
	if x, ok := x.GetResult().(*InferenceResult_RegressionResult); ok {
		return x.RegressionResult
	}
	return nil
}

type isInferenceResult_Result interface {
	isInferenceResult_Result()
}

type InferenceResult_ClassificationResult struct {
	ClassificationResult *ClassificationResult `protobuf:"bytes,2,opt,name=classification_result,json=classificationResult,proto3,oneof"`
}

type InferenceResult_RegressionResult struct {
	RegressionResult *RegressionResult `protobuf:"bytes,3,opt,name=regression_result,json=regressionResult,proto3,oneof"`
}

func (*InferenceResult_ClassificationResult) isInferenceResult_Result() {}

func (*InferenceResult_RegressionResult) isInferenceResult_Result() {}

// Inference request containing one or more requests.
type MultiInferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Inference tasks.
	Tasks []*InferenceTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// Input data.
	Input *Input `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *MultiInferenceRequest) Reset() {
	*x = MultiInferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiInferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiInferenceRequest) ProtoMessage() {}

func (x *MultiInferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiInferenceRequest.ProtoReflect.Descriptor instead.
func (*MultiInferenceRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_inference_proto_rawDescGZIP(), []int{2}
}

func (x *MultiInferenceRequest) GetTasks() []*InferenceTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *MultiInferenceRequest) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

// Inference request containing one or more responses.
type MultiInferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of results; one for each InferenceTask in the request, returned in the
	// same order as the request.
	Results []*InferenceResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MultiInferenceResponse) Reset() {
	*x = MultiInferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiInferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiInferenceResponse) ProtoMessage() {}

func (x *MultiInferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_inference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiInferenceResponse.ProtoReflect.Descriptor instead.
func (*MultiInferenceResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_inference_proto_rawDescGZIP(), []int{3}
}

func (x *MultiInferenceResponse) GetResults() []*InferenceResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_tensorflow_serving_apis_inference_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_inference_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x2c, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6e, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x8f, 0x02, 0x0a, 0x0f, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x5f, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x53, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x57, 0x0a, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x03,
	0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_inference_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_inference_proto_rawDescData = file_tensorflow_serving_apis_inference_proto_rawDesc
)

func file_tensorflow_serving_apis_inference_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_inference_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_inference_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_inference_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_inference_proto_rawDescData
}

var file_tensorflow_serving_apis_inference_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_serving_apis_inference_proto_goTypes = []interface{}{
	(*InferenceTask)(nil),          // 0: tensorflow.serving.InferenceTask
	(*InferenceResult)(nil),        // 1: tensorflow.serving.InferenceResult
	(*MultiInferenceRequest)(nil),  // 2: tensorflow.serving.MultiInferenceRequest
	(*MultiInferenceResponse)(nil), // 3: tensorflow.serving.MultiInferenceResponse
	(*ModelSpec)(nil),              // 4: tensorflow.serving.ModelSpec
	(*ClassificationResult)(nil),   // 5: tensorflow.serving.ClassificationResult
	(*RegressionResult)(nil),       // 6: tensorflow.serving.RegressionResult
	(*Input)(nil),                  // 7: tensorflow.serving.Input
}
var file_tensorflow_serving_apis_inference_proto_depIdxs = []int32{
	4, // 0: tensorflow.serving.InferenceTask.model_spec:type_name -> tensorflow.serving.ModelSpec
	4, // 1: tensorflow.serving.InferenceResult.model_spec:type_name -> tensorflow.serving.ModelSpec
	5, // 2: tensorflow.serving.InferenceResult.classification_result:type_name -> tensorflow.serving.ClassificationResult
	6, // 3: tensorflow.serving.InferenceResult.regression_result:type_name -> tensorflow.serving.RegressionResult
	0, // 4: tensorflow.serving.MultiInferenceRequest.tasks:type_name -> tensorflow.serving.InferenceTask
	7, // 5: tensorflow.serving.MultiInferenceRequest.input:type_name -> tensorflow.serving.Input
	1, // 6: tensorflow.serving.MultiInferenceResponse.results:type_name -> tensorflow.serving.InferenceResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_inference_proto_init() }
func file_tensorflow_serving_apis_inference_proto_init() {
	if File_tensorflow_serving_apis_inference_proto != nil {
		return
	}
	file_tensorflow_serving_apis_classification_proto_init()
	file_tensorflow_serving_apis_input_proto_init()
	file_tensorflow_serving_apis_model_proto_init()
	file_tensorflow_serving_apis_regression_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_inference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceTask); i {
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
		file_tensorflow_serving_apis_inference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResult); i {
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
		file_tensorflow_serving_apis_inference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiInferenceRequest); i {
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
		file_tensorflow_serving_apis_inference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiInferenceResponse); i {
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
	file_tensorflow_serving_apis_inference_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InferenceResult_ClassificationResult)(nil),
		(*InferenceResult_RegressionResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_inference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_inference_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_inference_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_inference_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_inference_proto = out.File
	file_tensorflow_serving_apis_inference_proto_rawDesc = nil
	file_tensorflow_serving_apis_inference_proto_goTypes = nil
	file_tensorflow_serving_apis_inference_proto_depIdxs = nil
}
