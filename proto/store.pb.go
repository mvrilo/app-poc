// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: store.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_ACTIVE   Status = 0
	Status_INACTIVE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
	}
	Status_value = map[string]int32{
		"ACTIVE":   0,
		"INACTIVE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_store_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uri     string               `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Status  Status               `protobuf:"varint,4,opt,name=status,proto3,enum=store.Status" json:"status,omitempty"`
	Created *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Store) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ACTIVE
}

func (x *Store) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Store) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type Stores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stores []*Store `protobuf:"bytes,1,rep,name=stores,proto3" json:"stores,omitempty"`
}

func (x *Stores) Reset() {
	*x = Stores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stores) ProtoMessage() {}

func (x *Stores) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stores.ProtoReflect.Descriptor instead.
func (*Stores) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{1}
}

func (x *Stores) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{2}
}

func (x *FindRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uri  string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type ChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=store.Status" json:"status,omitempty"`
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeStatusRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ACTIVE
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=store.Status" json:"status,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{5}
}

func (x *ListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ACTIVE
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo string `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{6}
}

func (x *EchoResponse) GetEcho() string {
	if x != nil {
		return x.Echo
	}
	return ""
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22,
	0x4c, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x2a, 0x22, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x32,
	0x80, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData = file_store_proto_rawDesc
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_proto_rawDescData)
	})
	return file_store_proto_rawDescData
}

var file_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: store.Status
	(*Store)(nil),               // 1: store.Store
	(*Stores)(nil),              // 2: store.Stores
	(*FindRequest)(nil),         // 3: store.FindRequest
	(*CreateRequest)(nil),       // 4: store.CreateRequest
	(*ChangeStatusRequest)(nil), // 5: store.ChangeStatusRequest
	(*ListRequest)(nil),         // 6: store.ListRequest
	(*EchoResponse)(nil),        // 7: store.EchoResponse
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_store_proto_depIdxs = []int32{
	0,  // 0: store.Store.status:type_name -> store.Status
	8,  // 1: store.Store.created:type_name -> google.protobuf.Timestamp
	8,  // 2: store.Store.updated:type_name -> google.protobuf.Timestamp
	1,  // 3: store.Stores.stores:type_name -> store.Store
	0,  // 4: store.ChangeStatusRequest.status:type_name -> store.Status
	0,  // 5: store.ListRequest.status:type_name -> store.Status
	3,  // 6: store.StoreService.Find:input_type -> store.FindRequest
	6,  // 7: store.StoreService.List:input_type -> store.ListRequest
	4,  // 8: store.StoreService.Create:input_type -> store.CreateRequest
	5,  // 9: store.StoreService.ChangeStatus:input_type -> store.ChangeStatusRequest
	9,  // 10: store.StoreService.Echo:input_type -> google.protobuf.Empty
	1,  // 11: store.StoreService.Find:output_type -> store.Store
	2,  // 12: store.StoreService.List:output_type -> store.Stores
	1,  // 13: store.StoreService.Create:output_type -> store.Store
	1,  // 14: store.StoreService.ChangeStatus:output_type -> store.Store
	7,  // 15: store.StoreService.Echo:output_type -> store.EchoResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stores); i {
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
		file_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusRequest); i {
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
		file_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		EnumInfos:         file_store_proto_enumTypes,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Store, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Stores, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Store, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Store, error)
	Echo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*EchoResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/store.StoreService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Stores, error) {
	out := new(Stores)
	err := c.cc.Invoke(ctx, "/store.StoreService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/store.StoreService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/store.StoreService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Echo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/store.StoreService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	Find(context.Context, *FindRequest) (*Store, error)
	List(context.Context, *ListRequest) (*Stores, error)
	Create(context.Context, *CreateRequest) (*Store, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*Store, error)
	Echo(context.Context, *empty.Empty) (*EchoResponse, error)
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) Find(context.Context, *FindRequest) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedStoreServiceServer) List(context.Context, *ListRequest) (*Stores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStoreServiceServer) Create(context.Context, *CreateRequest) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedStoreServiceServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (*UnimplementedStoreServiceServer) Echo(context.Context, *empty.Empty) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Echo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _StoreService_Find_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StoreService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StoreService_Create_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _StoreService_ChangeStatus_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _StoreService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
