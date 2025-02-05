// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: test_agent.proto

package gen

import (
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BindSimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid string `protobuf:"bytes,1,opt,name=iccid,proto3" json:"iccid,omitempty"`
}

func (x *BindSimRequest) Reset() {
	*x = BindSimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindSimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindSimRequest) ProtoMessage() {}

func (x *BindSimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindSimRequest.ProtoReflect.Descriptor instead.
func (*BindSimRequest) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{0}
}

func (x *BindSimRequest) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

type BindSimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindSimResponse) Reset() {
	*x = BindSimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindSimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindSimResponse) ProtoMessage() {}

func (x *BindSimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindSimResponse.ProtoReflect.Descriptor instead.
func (*BindSimResponse) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{1}
}

type ActivateSimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid string `protobuf:"bytes,1,opt,name=iccid,proto3" json:"iccid,omitempty"`
}

func (x *ActivateSimRequest) Reset() {
	*x = ActivateSimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateSimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateSimRequest) ProtoMessage() {}

func (x *ActivateSimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateSimRequest.ProtoReflect.Descriptor instead.
func (*ActivateSimRequest) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ActivateSimRequest) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

type ActivateSimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActivateSimResponse) Reset() {
	*x = ActivateSimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateSimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateSimResponse) ProtoMessage() {}

func (x *ActivateSimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateSimResponse.ProtoReflect.Descriptor instead.
func (*ActivateSimResponse) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{3}
}

type GetSimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid string `protobuf:"bytes,1,opt,name=iccid,proto3" json:"iccid,omitempty"`
}

func (x *GetSimRequest) Reset() {
	*x = GetSimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimRequest) ProtoMessage() {}

func (x *GetSimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimRequest.ProtoReflect.Descriptor instead.
func (*GetSimRequest) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{4}
}

func (x *GetSimRequest) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

type GetSimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimInfo *SimInfo `protobuf:"bytes,1,opt,name=simInfo,json=sim,proto3" json:"simInfo,omitempty"`
}

func (x *GetSimResponse) Reset() {
	*x = GetSimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimResponse) ProtoMessage() {}

func (x *GetSimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimResponse.ProtoReflect.Descriptor instead.
func (*GetSimResponse) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{5}
}

func (x *GetSimResponse) GetSimInfo() *SimInfo {
	if x != nil {
		return x.SimInfo
	}
	return nil
}

type DeactivateSimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid string `protobuf:"bytes,1,opt,name=iccid,proto3" json:"iccid,omitempty"`
}

func (x *DeactivateSimRequest) Reset() {
	*x = DeactivateSimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateSimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateSimRequest) ProtoMessage() {}

func (x *DeactivateSimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateSimRequest.ProtoReflect.Descriptor instead.
func (*DeactivateSimRequest) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{6}
}

func (x *DeactivateSimRequest) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

type DeactivateSimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeactivateSimResponse) Reset() {
	*x = DeactivateSimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateSimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateSimResponse) ProtoMessage() {}

func (x *DeactivateSimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateSimResponse.ProtoReflect.Descriptor instead.
func (*DeactivateSimResponse) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{7}
}

type TerminateSimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid string `protobuf:"bytes,1,opt,name=iccid,proto3" json:"iccid,omitempty"`
}

func (x *TerminateSimRequest) Reset() {
	*x = TerminateSimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateSimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateSimRequest) ProtoMessage() {}

func (x *TerminateSimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateSimRequest.ProtoReflect.Descriptor instead.
func (*TerminateSimRequest) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{8}
}

func (x *TerminateSimRequest) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

type TerminateSimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateSimResponse) Reset() {
	*x = TerminateSimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateSimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateSimResponse) ProtoMessage() {}

func (x *TerminateSimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateSimResponse.ProtoReflect.Descriptor instead.
func (*TerminateSimResponse) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{9}
}

type SimInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iccid  string `protobuf:"bytes,1,opt,name=Iccid,json=iccid,proto3" json:"Iccid,omitempty"`
	Imsi   string `protobuf:"bytes,2,opt,name=Imsi,json=imsi,proto3" json:"Imsi,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
}

func (x *SimInfo) Reset() {
	*x = SimInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_agent_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimInfo) ProtoMessage() {}

func (x *SimInfo) ProtoReflect() protoreflect.Message {
	mi := &file_test_agent_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimInfo.ProtoReflect.Descriptor instead.
func (*SimInfo) Descriptor() ([]byte, []int) {
	return file_test_agent_proto_rawDescGZIP(), []int{10}
}

func (x *SimInfo) GetIccid() string {
	if x != nil {
		return x.Iccid
	}
	return ""
}

func (x *SimInfo) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *SimInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_test_agent_proto protoreflect.FileDescriptor

var file_test_agent_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f,
	0x42, 0x69, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2a, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x73,
	0x69, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x75,
	0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x69, 0x6d, 0x22, 0x2c, 0x0a, 0x14, 0x44, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2b, 0x0a, 0x13, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x69,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x22, 0x16,
	0x0a, 0x14, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x07, 0x53, 0x69, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x63, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x73, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xd8, 0x04, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64,
	0x53, 0x69, 0x6d, 0x12, 0x2e, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x12, 0x2d,
	0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a,
	0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x12, 0x32, 0x2e, 0x75,
	0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x12, 0x34, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x75,
	0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x0c, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6d, 0x12, 0x33, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x69,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6b, 0x61,
	0x6d, 0x61, 0x2f, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_agent_proto_rawDescOnce sync.Once
	file_test_agent_proto_rawDescData = file_test_agent_proto_rawDesc
)

func file_test_agent_proto_rawDescGZIP() []byte {
	file_test_agent_proto_rawDescOnce.Do(func() {
		file_test_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_agent_proto_rawDescData)
	})
	return file_test_agent_proto_rawDescData
}

var file_test_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_test_agent_proto_goTypes = []interface{}{
	(*BindSimRequest)(nil),        // 0: ukama.subscriber.test_agent.v1.BindSimRequest
	(*BindSimResponse)(nil),       // 1: ukama.subscriber.test_agent.v1.BindSimResponse
	(*ActivateSimRequest)(nil),    // 2: ukama.subscriber.test_agent.v1.ActivateSimRequest
	(*ActivateSimResponse)(nil),   // 3: ukama.subscriber.test_agent.v1.ActivateSimResponse
	(*GetSimRequest)(nil),         // 4: ukama.subscriber.test_agent.v1.GetSimRequest
	(*GetSimResponse)(nil),        // 5: ukama.subscriber.test_agent.v1.GetSimResponse
	(*DeactivateSimRequest)(nil),  // 6: ukama.subscriber.test_agent.v1.DeactivateSimRequest
	(*DeactivateSimResponse)(nil), // 7: ukama.subscriber.test_agent.v1.DeactivateSimResponse
	(*TerminateSimRequest)(nil),   // 8: ukama.subscriber.test_agent.v1.TerminateSimRequest
	(*TerminateSimResponse)(nil),  // 9: ukama.subscriber.test_agent.v1.TerminateSimResponse
	(*SimInfo)(nil),               // 10: ukama.subscriber.test_agent.v1.SimInfo
}
var file_test_agent_proto_depIdxs = []int32{
	10, // 0: ukama.subscriber.test_agent.v1.GetSimResponse.simInfo:type_name -> ukama.subscriber.test_agent.v1.SimInfo
	0,  // 1: ukama.subscriber.test_agent.v1.TestAgentService.BindSim:input_type -> ukama.subscriber.test_agent.v1.BindSimRequest
	4,  // 2: ukama.subscriber.test_agent.v1.TestAgentService.GetSim:input_type -> ukama.subscriber.test_agent.v1.GetSimRequest
	2,  // 3: ukama.subscriber.test_agent.v1.TestAgentService.ActivateSim:input_type -> ukama.subscriber.test_agent.v1.ActivateSimRequest
	6,  // 4: ukama.subscriber.test_agent.v1.TestAgentService.DeactivateSim:input_type -> ukama.subscriber.test_agent.v1.DeactivateSimRequest
	8,  // 5: ukama.subscriber.test_agent.v1.TestAgentService.TerminateSim:input_type -> ukama.subscriber.test_agent.v1.TerminateSimRequest
	1,  // 6: ukama.subscriber.test_agent.v1.TestAgentService.BindSim:output_type -> ukama.subscriber.test_agent.v1.BindSimResponse
	5,  // 7: ukama.subscriber.test_agent.v1.TestAgentService.GetSim:output_type -> ukama.subscriber.test_agent.v1.GetSimResponse
	3,  // 8: ukama.subscriber.test_agent.v1.TestAgentService.ActivateSim:output_type -> ukama.subscriber.test_agent.v1.ActivateSimResponse
	7,  // 9: ukama.subscriber.test_agent.v1.TestAgentService.DeactivateSim:output_type -> ukama.subscriber.test_agent.v1.DeactivateSimResponse
	9,  // 10: ukama.subscriber.test_agent.v1.TestAgentService.TerminateSim:output_type -> ukama.subscriber.test_agent.v1.TerminateSimResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_test_agent_proto_init() }
func file_test_agent_proto_init() {
	if File_test_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindSimRequest); i {
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
		file_test_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindSimResponse); i {
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
		file_test_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateSimRequest); i {
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
		file_test_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateSimResponse); i {
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
		file_test_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimRequest); i {
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
		file_test_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimResponse); i {
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
		file_test_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateSimRequest); i {
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
		file_test_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateSimResponse); i {
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
		file_test_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateSimRequest); i {
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
		file_test_agent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateSimResponse); i {
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
		file_test_agent_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimInfo); i {
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
			RawDescriptor: file_test_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_agent_proto_goTypes,
		DependencyIndexes: file_test_agent_proto_depIdxs,
		MessageInfos:      file_test_agent_proto_msgTypes,
	}.Build()
	File_test_agent_proto = out.File
	file_test_agent_proto_rawDesc = nil
	file_test_agent_proto_goTypes = nil
	file_test_agent_proto_depIdxs = nil
}
