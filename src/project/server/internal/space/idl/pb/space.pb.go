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

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
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

type SpaceCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *Space `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
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

func (x *SpaceCreateReply) GetError() *Error {
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

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *SpaceRemoveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SpaceRemoveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
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

func (x *SpaceRemoveReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SpaceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *SpaceInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SpaceInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *Space `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
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

func (x *SpaceInfoReply) GetError() *Error {
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

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
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

type SpaceListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  []*Space `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
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

func (x *SpaceListReply) GetError() *Error {
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

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
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

func (x *Space) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Space) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Space) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Space) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_space_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_idl_space_proto_rawDescGZIP(), []int{9}
}

func (x *Error) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_idl_space_proto protoreflect.FileDescriptor

var file_idl_space_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x28, 0x0a,
	0x12, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3b, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x22,
	0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x60, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x10, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x60, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2f, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa2,
	0x02, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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
	(*SpaceCreateRequest)(nil), // 0: demo.space.SpaceCreateRequest
	(*SpaceCreateReply)(nil),   // 1: demo.space.SpaceCreateReply
	(*SpaceRemoveRequest)(nil), // 2: demo.space.SpaceRemoveRequest
	(*SpaceRemoveReply)(nil),   // 3: demo.space.SpaceRemoveReply
	(*SpaceInfoRequest)(nil),   // 4: demo.space.SpaceInfoRequest
	(*SpaceInfoReply)(nil),     // 5: demo.space.SpaceInfoReply
	(*SpaceListRequest)(nil),   // 6: demo.space.SpaceListRequest
	(*SpaceListReply)(nil),     // 7: demo.space.SpaceListReply
	(*Space)(nil),              // 8: demo.space.Space
	(*Error)(nil),              // 9: demo.space.Error
}
var file_idl_space_proto_depIdxs = []int32{
	9,  // 0: demo.space.SpaceCreateReply.error:type_name -> demo.space.Error
	8,  // 1: demo.space.SpaceCreateReply.data:type_name -> demo.space.Space
	9,  // 2: demo.space.SpaceRemoveReply.error:type_name -> demo.space.Error
	9,  // 3: demo.space.SpaceInfoReply.error:type_name -> demo.space.Error
	8,  // 4: demo.space.SpaceInfoReply.data:type_name -> demo.space.Space
	9,  // 5: demo.space.SpaceListReply.error:type_name -> demo.space.Error
	8,  // 6: demo.space.SpaceListReply.data:type_name -> demo.space.Space
	0,  // 7: demo.space.SpaceService.Create:input_type -> demo.space.SpaceCreateRequest
	2,  // 8: demo.space.SpaceService.Remove:input_type -> demo.space.SpaceRemoveRequest
	4,  // 9: demo.space.SpaceService.Info:input_type -> demo.space.SpaceInfoRequest
	6,  // 10: demo.space.SpaceService.List:input_type -> demo.space.SpaceListRequest
	1,  // 11: demo.space.SpaceService.Create:output_type -> demo.space.SpaceCreateReply
	3,  // 12: demo.space.SpaceService.Remove:output_type -> demo.space.SpaceRemoveReply
	5,  // 13: demo.space.SpaceService.Info:output_type -> demo.space.SpaceInfoReply
	7,  // 14: demo.space.SpaceService.List:output_type -> demo.space.SpaceListReply
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
			switch v := v.(*Error); i {
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