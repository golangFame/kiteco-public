// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: tensorflow/core/protobuf/control_flow.proto

package protobuf

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

// Protocol buffer representing the values in ControlFlowContext.
type ValuesDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value names that have been seen in this context.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Value names referenced by but external to this context.
	ExternalValues map[string]string `protobuf:"bytes,2,rep,name=external_values,json=externalValues,proto3" json:"external_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ValuesDef) Reset() {
	*x = ValuesDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValuesDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValuesDef) ProtoMessage() {}

func (x *ValuesDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValuesDef.ProtoReflect.Descriptor instead.
func (*ValuesDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_control_flow_proto_rawDescGZIP(), []int{0}
}

func (x *ValuesDef) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ValuesDef) GetExternalValues() map[string]string {
	if x != nil {
		return x.ExternalValues
	}
	return nil
}

// Container for any kind of control flow context. Any other control flow
// contexts that are added below should also be added here.
type ControlFlowContextDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ctxt:
	//	*ControlFlowContextDef_CondCtxt
	//	*ControlFlowContextDef_WhileCtxt
	Ctxt isControlFlowContextDef_Ctxt `protobuf_oneof:"ctxt"`
}

func (x *ControlFlowContextDef) Reset() {
	*x = ControlFlowContextDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlFlowContextDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlFlowContextDef) ProtoMessage() {}

func (x *ControlFlowContextDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlFlowContextDef.ProtoReflect.Descriptor instead.
func (*ControlFlowContextDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_control_flow_proto_rawDescGZIP(), []int{1}
}

func (m *ControlFlowContextDef) GetCtxt() isControlFlowContextDef_Ctxt {
	if m != nil {
		return m.Ctxt
	}
	return nil
}

func (x *ControlFlowContextDef) GetCondCtxt() *CondContextDef {
	if x, ok := x.GetCtxt().(*ControlFlowContextDef_CondCtxt); ok {
		return x.CondCtxt
	}
	return nil
}

func (x *ControlFlowContextDef) GetWhileCtxt() *WhileContextDef {
	if x, ok := x.GetCtxt().(*ControlFlowContextDef_WhileCtxt); ok {
		return x.WhileCtxt
	}
	return nil
}

type isControlFlowContextDef_Ctxt interface {
	isControlFlowContextDef_Ctxt()
}

type ControlFlowContextDef_CondCtxt struct {
	CondCtxt *CondContextDef `protobuf:"bytes,1,opt,name=cond_ctxt,json=condCtxt,proto3,oneof"`
}

type ControlFlowContextDef_WhileCtxt struct {
	WhileCtxt *WhileContextDef `protobuf:"bytes,2,opt,name=while_ctxt,json=whileCtxt,proto3,oneof"`
}

func (*ControlFlowContextDef_CondCtxt) isControlFlowContextDef_Ctxt() {}

func (*ControlFlowContextDef_WhileCtxt) isControlFlowContextDef_Ctxt() {}

// Protocol buffer representing a CondContext object.
type CondContextDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// Name of the pred tensor.
	PredName string `protobuf:"bytes,2,opt,name=pred_name,json=predName,proto3" json:"pred_name,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,3,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Branch prediction. 0 or 1.
	Branch int32 `protobuf:"varint,4,opt,name=branch,proto3" json:"branch,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,5,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Contexts contained inside this context (e.g. nested conds).
	NestedContexts []*ControlFlowContextDef `protobuf:"bytes,6,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
}

func (x *CondContextDef) Reset() {
	*x = CondContextDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CondContextDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CondContextDef) ProtoMessage() {}

func (x *CondContextDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CondContextDef.ProtoReflect.Descriptor instead.
func (*CondContextDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_control_flow_proto_rawDescGZIP(), []int{2}
}

func (x *CondContextDef) GetContextName() string {
	if x != nil {
		return x.ContextName
	}
	return ""
}

func (x *CondContextDef) GetPredName() string {
	if x != nil {
		return x.PredName
	}
	return ""
}

func (x *CondContextDef) GetPivotName() string {
	if x != nil {
		return x.PivotName
	}
	return ""
}

func (x *CondContextDef) GetBranch() int32 {
	if x != nil {
		return x.Branch
	}
	return 0
}

func (x *CondContextDef) GetValuesDef() *ValuesDef {
	if x != nil {
		return x.ValuesDef
	}
	return nil
}

func (x *CondContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if x != nil {
		return x.NestedContexts
	}
	return nil
}

// Protocol buffer representing a WhileContext object.
type WhileContextDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// The number of iterations allowed to run in parallel.
	ParallelIterations int32 `protobuf:"varint,2,opt,name=parallel_iterations,json=parallelIterations,proto3" json:"parallel_iterations,omitempty"`
	// Whether backprop is enabled for this while loop.
	BackProp bool `protobuf:"varint,3,opt,name=back_prop,json=backProp,proto3" json:"back_prop,omitempty"`
	// Whether GPU-CPU memory swap is enabled for this loop.
	SwapMemory bool `protobuf:"varint,4,opt,name=swap_memory,json=swapMemory,proto3" json:"swap_memory,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,5,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Name of the pivot_for_pred tensor.
	PivotForPredName string `protobuf:"bytes,6,opt,name=pivot_for_pred_name,json=pivotForPredName,proto3" json:"pivot_for_pred_name,omitempty"`
	// Name of the pivot_for_body tensor.
	PivotForBodyName string `protobuf:"bytes,7,opt,name=pivot_for_body_name,json=pivotForBodyName,proto3" json:"pivot_for_body_name,omitempty"`
	// List of names for exit tensors.
	LoopExitNames []string `protobuf:"bytes,8,rep,name=loop_exit_names,json=loopExitNames,proto3" json:"loop_exit_names,omitempty"`
	// List of names for enter tensors.
	LoopEnterNames []string `protobuf:"bytes,10,rep,name=loop_enter_names,json=loopEnterNames,proto3" json:"loop_enter_names,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,9,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Optional name of the maximum_iterations tensor.
	MaximumIterationsName string `protobuf:"bytes,11,opt,name=maximum_iterations_name,json=maximumIterationsName,proto3" json:"maximum_iterations_name,omitempty"`
	// Contexts contained inside this context (e.g. nested whiles).
	NestedContexts []*ControlFlowContextDef `protobuf:"bytes,12,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
}

func (x *WhileContextDef) Reset() {
	*x = WhileContextDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhileContextDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhileContextDef) ProtoMessage() {}

func (x *WhileContextDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_control_flow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhileContextDef.ProtoReflect.Descriptor instead.
func (*WhileContextDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_control_flow_proto_rawDescGZIP(), []int{3}
}

func (x *WhileContextDef) GetContextName() string {
	if x != nil {
		return x.ContextName
	}
	return ""
}

func (x *WhileContextDef) GetParallelIterations() int32 {
	if x != nil {
		return x.ParallelIterations
	}
	return 0
}

func (x *WhileContextDef) GetBackProp() bool {
	if x != nil {
		return x.BackProp
	}
	return false
}

func (x *WhileContextDef) GetSwapMemory() bool {
	if x != nil {
		return x.SwapMemory
	}
	return false
}

func (x *WhileContextDef) GetPivotName() string {
	if x != nil {
		return x.PivotName
	}
	return ""
}

func (x *WhileContextDef) GetPivotForPredName() string {
	if x != nil {
		return x.PivotForPredName
	}
	return ""
}

func (x *WhileContextDef) GetPivotForBodyName() string {
	if x != nil {
		return x.PivotForBodyName
	}
	return ""
}

func (x *WhileContextDef) GetLoopExitNames() []string {
	if x != nil {
		return x.LoopExitNames
	}
	return nil
}

func (x *WhileContextDef) GetLoopEnterNames() []string {
	if x != nil {
		return x.LoopEnterNames
	}
	return nil
}

func (x *WhileContextDef) GetValuesDef() *ValuesDef {
	if x != nil {
		return x.ValuesDef
	}
	return nil
}

func (x *WhileContextDef) GetMaximumIterationsName() string {
	if x != nil {
		return x.MaximumIterationsName
	}
	return ""
}

func (x *WhileContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if x != nil {
		return x.NestedContexts
	}
	return nil
}

var File_tensorflow_core_protobuf_control_flow_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_control_flow_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x52, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x66,
	0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x63, 0x74, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x66, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x64, 0x43, 0x74, 0x78, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x77,
	0x68, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x74, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x68, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x66, 0x48, 0x00, 0x52, 0x09,
	0x77, 0x68, 0x69, 0x6c, 0x65, 0x43, 0x74, 0x78, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x63, 0x74, 0x78,
	0x74, 0x22, 0x89, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x44, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x44, 0x65,
	0x66, 0x12, 0x4a, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x46,
	0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x66, 0x52, 0x0e, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x22, 0xac, 0x04,
	0x0a, 0x0f, 0x57, 0x68, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65,
	0x66, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c,
	0x5f, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x72,
	0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72,
	0x6f, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x70, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x70, 0x69, 0x76, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x6f, 0x70, 0x45,
	0x78, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x6f, 0x70,
	0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x6f, 0x70, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x66,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x52, 0x09, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4a, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x46, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x66, 0x52, 0x0e, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x42, 0x70, 0x0a, 0x18,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_control_flow_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_control_flow_proto_rawDescData = file_tensorflow_core_protobuf_control_flow_proto_rawDesc
)

func file_tensorflow_core_protobuf_control_flow_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_control_flow_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_control_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_control_flow_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_control_flow_proto_rawDescData
}

var file_tensorflow_core_protobuf_control_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_core_protobuf_control_flow_proto_goTypes = []interface{}{
	(*ValuesDef)(nil),             // 0: tensorflow.ValuesDef
	(*ControlFlowContextDef)(nil), // 1: tensorflow.ControlFlowContextDef
	(*CondContextDef)(nil),        // 2: tensorflow.CondContextDef
	(*WhileContextDef)(nil),       // 3: tensorflow.WhileContextDef
	nil,                           // 4: tensorflow.ValuesDef.ExternalValuesEntry
}
var file_tensorflow_core_protobuf_control_flow_proto_depIdxs = []int32{
	4, // 0: tensorflow.ValuesDef.external_values:type_name -> tensorflow.ValuesDef.ExternalValuesEntry
	2, // 1: tensorflow.ControlFlowContextDef.cond_ctxt:type_name -> tensorflow.CondContextDef
	3, // 2: tensorflow.ControlFlowContextDef.while_ctxt:type_name -> tensorflow.WhileContextDef
	0, // 3: tensorflow.CondContextDef.values_def:type_name -> tensorflow.ValuesDef
	1, // 4: tensorflow.CondContextDef.nested_contexts:type_name -> tensorflow.ControlFlowContextDef
	0, // 5: tensorflow.WhileContextDef.values_def:type_name -> tensorflow.ValuesDef
	1, // 6: tensorflow.WhileContextDef.nested_contexts:type_name -> tensorflow.ControlFlowContextDef
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_control_flow_proto_init() }
func file_tensorflow_core_protobuf_control_flow_proto_init() {
	if File_tensorflow_core_protobuf_control_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_control_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValuesDef); i {
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
		file_tensorflow_core_protobuf_control_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlFlowContextDef); i {
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
		file_tensorflow_core_protobuf_control_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CondContextDef); i {
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
		file_tensorflow_core_protobuf_control_flow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhileContextDef); i {
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
	file_tensorflow_core_protobuf_control_flow_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ControlFlowContextDef_CondCtxt)(nil),
		(*ControlFlowContextDef_WhileCtxt)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_control_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_control_flow_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_control_flow_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_control_flow_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_control_flow_proto = out.File
	file_tensorflow_core_protobuf_control_flow_proto_rawDesc = nil
	file_tensorflow_core_protobuf_control_flow_proto_goTypes = nil
	file_tensorflow_core_protobuf_control_flow_proto_depIdxs = nil
}
