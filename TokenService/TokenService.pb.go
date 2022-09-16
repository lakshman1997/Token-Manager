// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: TokenService/TokenService.proto

package TokenService

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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Low          uint64 `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	Mid          uint64 `protobuf:"varint,4,opt,name=mid,proto3" json:"mid,omitempty"`
	High         uint64 `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	PartialValue uint64 `protobuf:"varint,6,opt,name=partialValue,proto3" json:"partialValue,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{3}
}

func (x *Domain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Domain) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Domain) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Domain) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Domain) GetPartialValue() uint64 {
	if x != nil {
		return x.PartialValue
	}
	return 0
}

type ResponseForWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnPartialValue uint64 `protobuf:"varint,1,opt,name=returnPartialValue,proto3" json:"returnPartialValue,omitempty"`
	Errored            string `protobuf:"bytes,2,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *ResponseForWrite) Reset() {
	*x = ResponseForWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseForWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseForWrite) ProtoMessage() {}

func (x *ResponseForWrite) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseForWrite.ProtoReflect.Descriptor instead.
func (*ResponseForWrite) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseForWrite) GetReturnPartialValue() uint64 {
	if x != nil {
		return x.ReturnPartialValue
	}
	return 0
}

func (x *ResponseForWrite) GetErrored() string {
	if x != nil {
		return x.Errored
	}
	return ""
}

type ResponseForRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnFinalValue uint64 `protobuf:"varint,1,opt,name=returnFinalValue,proto3" json:"returnFinalValue,omitempty"`
	Errored          string `protobuf:"bytes,2,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *ResponseForRead) Reset() {
	*x = ResponseForRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseForRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseForRead) ProtoMessage() {}

func (x *ResponseForRead) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseForRead.ProtoReflect.Descriptor instead.
func (*ResponseForRead) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseForRead) GetReturnFinalValue() uint64 {
	if x != nil {
		return x.ReturnFinalValue
	}
	return 0
}

func (x *ResponseForRead) GetErrored() string {
	if x != nil {
		return x.Errored
	}
	return ""
}

type ReplicationDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Low          uint64 `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	Mid          uint64 `protobuf:"varint,4,opt,name=mid,proto3" json:"mid,omitempty"`
	High         uint64 `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	PartialValue uint64 `protobuf:"varint,6,opt,name=partialValue,proto3" json:"partialValue,omitempty"`
	FinalValue   uint64 `protobuf:"varint,7,opt,name=finalValue,proto3" json:"finalValue,omitempty"`
}

func (x *ReplicationDomain) Reset() {
	*x = ReplicationDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationDomain) ProtoMessage() {}

func (x *ReplicationDomain) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationDomain.ProtoReflect.Descriptor instead.
func (*ReplicationDomain) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{6}
}

func (x *ReplicationDomain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReplicationDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReplicationDomain) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *ReplicationDomain) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *ReplicationDomain) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *ReplicationDomain) GetPartialValue() uint64 {
	if x != nil {
		return x.PartialValue
	}
	return 0
}

func (x *ReplicationDomain) GetFinalValue() uint64 {
	if x != nil {
		return x.FinalValue
	}
	return 0
}

type ReplicationResponseForWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartialValue uint64 `protobuf:"varint,1,opt,name=partialValue,proto3" json:"partialValue,omitempty"`
}

func (x *ReplicationResponseForWrite) Reset() {
	*x = ReplicationResponseForWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationResponseForWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponseForWrite) ProtoMessage() {}

func (x *ReplicationResponseForWrite) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponseForWrite.ProtoReflect.Descriptor instead.
func (*ReplicationResponseForWrite) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{7}
}

func (x *ReplicationResponseForWrite) GetPartialValue() uint64 {
	if x != nil {
		return x.PartialValue
	}
	return 0
}

type ReplicationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body       string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	FinalValue uint64 `protobuf:"varint,2,opt,name=finalValue,proto3" json:"finalValue,omitempty"`
}

func (x *ReplicationID) Reset() {
	*x = ReplicationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationID) ProtoMessage() {}

func (x *ReplicationID) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationID.ProtoReflect.Descriptor instead.
func (*ReplicationID) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{8}
}

func (x *ReplicationID) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ReplicationID) GetFinalValue() uint64 {
	if x != nil {
		return x.FinalValue
	}
	return 0
}

type ReplicationResponseForRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalValue       uint64 `protobuf:"varint,1,opt,name=finalValue,proto3" json:"finalValue,omitempty"`
	ErrorWhileReturn string `protobuf:"bytes,2,opt,name=errorWhileReturn,proto3" json:"errorWhileReturn,omitempty"`
}

func (x *ReplicationResponseForRead) Reset() {
	*x = ReplicationResponseForRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_TokenService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationResponseForRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponseForRead) ProtoMessage() {}

func (x *ReplicationResponseForRead) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_TokenService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponseForRead.ProtoReflect.Descriptor instead.
func (*ReplicationResponseForRead) Descriptor() ([]byte, []int) {
	return file_TokenService_TokenService_proto_rawDescGZIP(), []int{9}
}

func (x *ReplicationResponseForRead) GetFinalValue() uint64 {
	if x != nil {
		return x.FinalValue
	}
	return 0
}

func (x *ReplicationResponseForRead) GetErrorWhileReturn() string {
	if x != nil {
		return x.ErrorWhileReturn
	}
	return ""
}

var File_TokenService_TokenService_proto protoreflect.FileDescriptor

var file_TokenService_TokenService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x19, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x18, 0x0a, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5c, 0x0a,
	0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x0f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2a,
	0x0a, 0x10, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x65, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x6f, 0x77,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x1b, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x46, 0x6f, 0x72, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a,
	0x0d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x68, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x57, 0x68, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x57, 0x68, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x32, 0xb3, 0x03, 0x0a,
	0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x52, 0x55, 0x44, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x1a, 0x1e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x1a, 0x1d, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64, 0x22, 0x00, 0x12, 0x63,
	0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x29, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x28, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x32, 0x2f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenService_TokenService_proto_rawDescOnce sync.Once
	file_TokenService_TokenService_proto_rawDescData = file_TokenService_TokenService_proto_rawDesc
)

func file_TokenService_TokenService_proto_rawDescGZIP() []byte {
	file_TokenService_TokenService_proto_rawDescOnce.Do(func() {
		file_TokenService_TokenService_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenService_TokenService_proto_rawDescData)
	})
	return file_TokenService_TokenService_proto_rawDescData
}

var file_TokenService_TokenService_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_TokenService_TokenService_proto_goTypes = []interface{}{
	(*Msg)(nil),                         // 0: TokenService.Msg
	(*ID)(nil),                          // 1: TokenService.ID
	(*Status)(nil),                      // 2: TokenService.Status
	(*Domain)(nil),                      // 3: TokenService.Domain
	(*ResponseForWrite)(nil),            // 4: TokenService.responseForWrite
	(*ResponseForRead)(nil),             // 5: TokenService.responseForRead
	(*ReplicationDomain)(nil),           // 6: TokenService.ReplicationDomain
	(*ReplicationResponseForWrite)(nil), // 7: TokenService.ReplicationResponseForWrite
	(*ReplicationID)(nil),               // 8: TokenService.ReplicationID
	(*ReplicationResponseForRead)(nil),  // 9: TokenService.ReplicationResponseForRead
}
var file_TokenService_TokenService_proto_depIdxs = []int32{
	1, // 0: TokenService.TokenCRUD.Create:input_type -> TokenService.ID
	1, // 1: TokenService.TokenCRUD.Delete:input_type -> TokenService.ID
	3, // 2: TokenService.TokenCRUD.Write:input_type -> TokenService.Domain
	1, // 3: TokenService.TokenCRUD.Read:input_type -> TokenService.ID
	6, // 4: TokenService.TokenCRUD.ReplicationForWrite:input_type -> TokenService.ReplicationDomain
	8, // 5: TokenService.TokenCRUD.ReplicationForRead:input_type -> TokenService.ReplicationID
	2, // 6: TokenService.TokenCRUD.Create:output_type -> TokenService.Status
	2, // 7: TokenService.TokenCRUD.Delete:output_type -> TokenService.Status
	4, // 8: TokenService.TokenCRUD.Write:output_type -> TokenService.responseForWrite
	5, // 9: TokenService.TokenCRUD.Read:output_type -> TokenService.responseForRead
	7, // 10: TokenService.TokenCRUD.ReplicationForWrite:output_type -> TokenService.ReplicationResponseForWrite
	9, // 11: TokenService.TokenCRUD.ReplicationForRead:output_type -> TokenService.ReplicationResponseForRead
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TokenService_TokenService_proto_init() }
func file_TokenService_TokenService_proto_init() {
	if File_TokenService_TokenService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TokenService_TokenService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_TokenService_TokenService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_TokenService_TokenService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_TokenService_TokenService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_TokenService_TokenService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseForWrite); i {
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
		file_TokenService_TokenService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseForRead); i {
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
		file_TokenService_TokenService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationDomain); i {
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
		file_TokenService_TokenService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationResponseForWrite); i {
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
		file_TokenService_TokenService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationID); i {
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
		file_TokenService_TokenService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationResponseForRead); i {
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
			RawDescriptor: file_TokenService_TokenService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TokenService_TokenService_proto_goTypes,
		DependencyIndexes: file_TokenService_TokenService_proto_depIdxs,
		MessageInfos:      file_TokenService_TokenService_proto_msgTypes,
	}.Build()
	File_TokenService_TokenService_proto = out.File
	file_TokenService_TokenService_proto_rawDesc = nil
	file_TokenService_TokenService_proto_goTypes = nil
	file_TokenService_TokenService_proto_depIdxs = nil
}
