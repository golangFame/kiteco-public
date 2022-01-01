// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/apis/prediction_service.proto

package serving

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_tensorflow_serving_apis_prediction_service_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_prediction_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xfc, 0x03, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x79, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x07, 0x52, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x22,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_serving_apis_prediction_service_proto_goTypes = []interface{}{
	(*ClassificationRequest)(nil),    // 0: tensorflow.serving.ClassificationRequest
	(*RegressionRequest)(nil),        // 1: tensorflow.serving.RegressionRequest
	(*PredictRequest)(nil),           // 2: tensorflow.serving.PredictRequest
	(*MultiInferenceRequest)(nil),    // 3: tensorflow.serving.MultiInferenceRequest
	(*GetModelMetadataRequest)(nil),  // 4: tensorflow.serving.GetModelMetadataRequest
	(*ClassificationResponse)(nil),   // 5: tensorflow.serving.ClassificationResponse
	(*RegressionResponse)(nil),       // 6: tensorflow.serving.RegressionResponse
	(*PredictResponse)(nil),          // 7: tensorflow.serving.PredictResponse
	(*MultiInferenceResponse)(nil),   // 8: tensorflow.serving.MultiInferenceResponse
	(*GetModelMetadataResponse)(nil), // 9: tensorflow.serving.GetModelMetadataResponse
}
var file_tensorflow_serving_apis_prediction_service_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.PredictionService.Classify:input_type -> tensorflow.serving.ClassificationRequest
	1, // 1: tensorflow.serving.PredictionService.Regress:input_type -> tensorflow.serving.RegressionRequest
	2, // 2: tensorflow.serving.PredictionService.Predict:input_type -> tensorflow.serving.PredictRequest
	3, // 3: tensorflow.serving.PredictionService.MultiInference:input_type -> tensorflow.serving.MultiInferenceRequest
	4, // 4: tensorflow.serving.PredictionService.GetModelMetadata:input_type -> tensorflow.serving.GetModelMetadataRequest
	5, // 5: tensorflow.serving.PredictionService.Classify:output_type -> tensorflow.serving.ClassificationResponse
	6, // 6: tensorflow.serving.PredictionService.Regress:output_type -> tensorflow.serving.RegressionResponse
	7, // 7: tensorflow.serving.PredictionService.Predict:output_type -> tensorflow.serving.PredictResponse
	8, // 8: tensorflow.serving.PredictionService.MultiInference:output_type -> tensorflow.serving.MultiInferenceResponse
	9, // 9: tensorflow.serving.PredictionService.GetModelMetadata:output_type -> tensorflow.serving.GetModelMetadataResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_prediction_service_proto_init() }
func file_tensorflow_serving_apis_prediction_service_proto_init() {
	if File_tensorflow_serving_apis_prediction_service_proto != nil {
		return
	}
	file_tensorflow_serving_apis_classification_proto_init()
	file_tensorflow_serving_apis_get_model_metadata_proto_init()
	file_tensorflow_serving_apis_inference_proto_init()
	file_tensorflow_serving_apis_predict_proto_init()
	file_tensorflow_serving_apis_regression_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_prediction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_serving_apis_prediction_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_prediction_service_proto_depIdxs,
	}.Build()
	File_tensorflow_serving_apis_prediction_service_proto = out.File
	file_tensorflow_serving_apis_prediction_service_proto_rawDesc = nil
	file_tensorflow_serving_apis_prediction_service_proto_goTypes = nil
	file_tensorflow_serving_apis_prediction_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Classify.
	Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error)
	// Regress.
	Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error) {
	out := new(ClassificationResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Classify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error) {
	out := new(RegressionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Regress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error) {
	out := new(MultiInferenceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/MultiInference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error) {
	out := new(GetModelMetadataResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/GetModelMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Classify.
	Classify(context.Context, *ClassificationRequest) (*ClassificationResponse, error)
	// Regress.
	Regress(context.Context, *RegressionRequest) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(context.Context, *MultiInferenceRequest) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(context.Context, *GetModelMetadataRequest) (*GetModelMetadataResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Classify(context.Context, *ClassificationRequest) (*ClassificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classify not implemented")
}
func (*UnimplementedPredictionServiceServer) Regress(context.Context, *RegressionRequest) (*RegressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regress not implemented")
}
func (*UnimplementedPredictionServiceServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (*UnimplementedPredictionServiceServer) MultiInference(context.Context, *MultiInferenceRequest) (*MultiInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiInference not implemented")
}
func (*UnimplementedPredictionServiceServer) GetModelMetadata(context.Context, *GetModelMetadataRequest) (*GetModelMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelMetadata not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Classify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Classify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Classify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Classify(ctx, req.(*ClassificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Regress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Regress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Regress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Regress(ctx, req.(*RegressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_MultiInference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).MultiInference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/MultiInference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).MultiInference(ctx, req.(*MultiInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetModelMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/GetModelMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, req.(*GetModelMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classify",
			Handler:    _PredictionService_Classify_Handler,
		},
		{
			MethodName: "Regress",
			Handler:    _PredictionService_Regress_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "MultiInference",
			Handler:    _PredictionService_MultiInference_Handler,
		},
		{
			MethodName: "GetModelMetadata",
			Handler:    _PredictionService_GetModelMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/prediction_service.proto",
}
