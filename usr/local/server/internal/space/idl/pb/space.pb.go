// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: idl/space.proto

package pb

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

type SpaceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// length > 6
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repos  string `protobuf:"bytes,2,opt,name=repos,proto3" json:"repos,omitempty"`
	Binary string `protobuf:"bytes,3,opt,name=binary,proto3" json:"binary,omitempty"`
}

func (x *SpaceCreateRequest) Reset() {
	*x = SpaceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceCreateRequest) ProtoMessage() {}

func (x *SpaceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceCreateRequest.ProtoReflect.Descriptor instead.
func (*SpaceCreateRequest) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpaceCreateRequest) GetRepos() string {
	if x != nil {
		return x.Repos
	}
	return ""
}

func (x *SpaceCreateRequest) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

type SpaceCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *SpaceError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *Space      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SpaceCreateReply) Reset() {
	*x = SpaceCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceCreateReply) ProtoMessage() {}

func (x *SpaceCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceCreateReply.ProtoReflect.Descriptor instead.
func (*SpaceCreateReply) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{1}
}

func (x *SpaceCreateReply) GetError() *SpaceError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SpaceCreateReply) GetData() *Space {
	if x != nil {
		return x.Data
	}
	return nil
}

type SpaceRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required
	SpaceID int64 `protobuf:"varint,1,opt,name=spaceID,proto3" json:"spaceID,omitempty"`
}

func (x *SpaceRemoveRequest) Reset() {
	*x = SpaceRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceRemoveRequest) ProtoMessage() {}

func (x *SpaceRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceRemoveRequest.ProtoReflect.Descriptor instead.
func (*SpaceRemoveRequest) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{2}
}

func (x *SpaceRemoveRequest) GetSpaceID() int64 {
	if x != nil {
		return x.SpaceID
	}
	return 0
}

type SpaceRemoveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *SpaceError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SpaceRemoveReply) Reset() {
	*x = SpaceRemoveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceRemoveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceRemoveReply) ProtoMessage() {}

func (x *SpaceRemoveReply) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceRemoveReply.ProtoReflect.Descriptor instead.
func (*SpaceRemoveReply) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{3}
}

func (x *SpaceRemoveReply) GetError() *SpaceError {
	if x != nil {
		return x.Error
	}
	return nil
}

type SpaceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required
	SpaceID int64 `protobuf:"varint,1,opt,name=spaceID,proto3" json:"spaceID,omitempty"`
}

func (x *SpaceInfoRequest) Reset() {
	*x = SpaceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceInfoRequest) ProtoMessage() {}

func (x *SpaceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceInfoRequest.ProtoReflect.Descriptor instead.
func (*SpaceInfoRequest) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{4}
}

func (x *SpaceInfoRequest) GetSpaceID() int64 {
	if x != nil {
		return x.SpaceID
	}
	return 0
}

type SpaceInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *SpaceError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *Space      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SpaceInfoReply) Reset() {
	*x = SpaceInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceInfoReply) ProtoMessage() {}

func (x *SpaceInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceInfoReply.ProtoReflect.Descriptor instead.
func (*SpaceInfoReply) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{5}
}

func (x *SpaceInfoReply) GetError() *SpaceError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SpaceInfoReply) GetData() *Space {
	if x != nil {
		return x.Data
	}
	return nil
}

type SpaceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default 1
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// default 10
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SpaceListRequest) Reset() {
	*x = SpaceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceListRequest) ProtoMessage() {}

func (x *SpaceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceListRequest.ProtoReflect.Descriptor instead.
func (*SpaceListRequest) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{6}
}

func (x *SpaceListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SpaceListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SpaceListRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SpaceListRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SpaceListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *SpaceError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  []*Space    `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Total int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SpaceListReply) Reset() {
	*x = SpaceListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceListReply) ProtoMessage() {}

func (x *SpaceListReply) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceListReply.ProtoReflect.Descriptor instead.
func (*SpaceListReply) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{7}
}

func (x *SpaceListReply) GetError() *SpaceError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SpaceListReply) GetData() []*Space {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SpaceListReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceID int64  `protobuf:"varint,1,opt,name=spaceID,proto3" json:"spaceID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Repos   string `protobuf:"bytes,3,opt,name=repos,proto3" json:"repos,omitempty"`
	Binary  string `protobuf:"bytes,4,opt,name=binary,proto3" json:"binary,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{8}
}

func (x *Space) GetSpaceID() int64 {
	if x != nil {
		return x.SpaceID
	}
	return 0
}

func (x *Space) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Space) GetRepos() string {
	if x != nil {
		return x.Repos
	}
	return ""
}

func (x *Space) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

type SpaceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SpaceError) Reset() {
	*x = SpaceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceError) ProtoMessage() {}

func (x *SpaceError) ProtoReflect() protoreflect.Message {
	mi := &file_idl_space_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceError.ProtoReflect.Descriptor instead.
func (*SpaceError) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{9}
}

func (x *SpaceError) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SpaceError) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_idl_space_proto protoreflect.FileDescriptor

var file_idl_space_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x12, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x22, 0x5d, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2e, 0x0a, 0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x22,
	0x3b, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x10,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x22, 0x5b, 0x0a, 0x0e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x71, 0x0a,
	0x0e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x63, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x70, 0x61, 0x63, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xfa, 0x01, 0x0a, 0x0c,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_space_proto_rawDescOnce sync.Once
	file_idl_space_proto_rawDescData = file_idl_space_proto_rawDesc
)

func file_idl_space_proto_rawDescGZIP() []byte {
	file_idl_space_proto_rawDescOnce.Do(func() {
		file_idl_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_space_proto_rawDescData)
	})
	return file_idl_space_proto_rawDescData
}

var file_idl_space_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_idl_space_proto_goTypes = []interface{}{
	(*SpaceCreateRequest)(nil), // 0: space.SpaceCreateRequest
	(*SpaceCreateReply)(nil),   // 1: space.SpaceCreateReply
	(*SpaceRemoveRequest)(nil), // 2: space.SpaceRemoveRequest
	(*SpaceRemoveReply)(nil),   // 3: space.SpaceRemoveReply
	(*SpaceInfoRequest)(nil),   // 4: space.SpaceInfoRequest
	(*SpaceInfoReply)(nil),     // 5: space.SpaceInfoReply
	(*SpaceListRequest)(nil),   // 6: space.SpaceListRequest
	(*SpaceListReply)(nil),     // 7: space.SpaceListReply
	(*Space)(nil),              // 8: space.Space
	(*SpaceError)(nil),         // 9: space.SpaceError
}
var file_idl_space_proto_depIdxs = []int32{
	9,  // 0: space.SpaceCreateReply.error:type_name -> space.SpaceError
	8,  // 1: space.SpaceCreateReply.data:type_name -> space.Space
	9,  // 2: space.SpaceRemoveReply.error:type_name -> space.SpaceError
	9,  // 3: space.SpaceInfoReply.error:type_name -> space.SpaceError
	8,  // 4: space.SpaceInfoReply.data:type_name -> space.Space
	9,  // 5: space.SpaceListReply.error:type_name -> space.SpaceError
	8,  // 6: space.SpaceListReply.data:type_name -> space.Space
	0,  // 7: space.SpaceService.Create:input_type -> space.SpaceCreateRequest
	2,  // 8: space.SpaceService.Remove:input_type -> space.SpaceRemoveRequest
	4,  // 9: space.SpaceService.Info:input_type -> space.SpaceInfoRequest
	6,  // 10: space.SpaceService.List:input_type -> space.SpaceListRequest
	1,  // 11: space.SpaceService.Create:output_type -> space.SpaceCreateReply
	3,  // 12: space.SpaceService.Remove:output_type -> space.SpaceRemoveReply
	5,  // 13: space.SpaceService.Info:output_type -> space.SpaceInfoReply
	7,  // 14: space.SpaceService.List:output_type -> space.SpaceListReply
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_idl_space_proto_init() }
func file_idl_space_proto_init() {
	if File_idl_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceCreateRequest); i {
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
		file_idl_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceCreateReply); i {
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
		file_idl_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceRemoveRequest); i {
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
		file_idl_space_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceRemoveReply); i {
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
		file_idl_space_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceInfoRequest); i {
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
		file_idl_space_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceInfoReply); i {
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
		file_idl_space_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceListRequest); i {
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
		file_idl_space_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceListReply); i {
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
		file_idl_space_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Space); i {
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
		file_idl_space_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceError); i {
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
			RawDescriptor: file_idl_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_space_proto_goTypes,
		DependencyIndexes: file_idl_space_proto_depIdxs,
		MessageInfos:      file_idl_space_proto_msgTypes,
	}.Build()
	File_idl_space_proto = out.File
	file_idl_space_proto_rawDesc = nil
	file_idl_space_proto_goTypes = nil
	file_idl_space_proto_depIdxs = nil
}
