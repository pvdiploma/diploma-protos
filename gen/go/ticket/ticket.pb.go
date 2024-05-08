// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ticket/ticket.proto

package ticketv1

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

type AddTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventCategoryId int64  `protobuf:"varint,1,opt,name=event_category_id,json=eventCategoryId,proto3" json:"event_category_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surrname        string `protobuf:"bytes,3,opt,name=surrname,proto3" json:"surrname,omitempty"`
	Patronymic      string `protobuf:"bytes,4,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
	Discount        uint32 `protobuf:"varint,5,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *AddTicketRequest) Reset() {
	*x = AddTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTicketRequest) ProtoMessage() {}

func (x *AddTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTicketRequest.ProtoReflect.Descriptor instead.
func (*AddTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *AddTicketRequest) GetEventCategoryId() int64 {
	if x != nil {
		return x.EventCategoryId
	}
	return 0
}

func (x *AddTicketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTicketRequest) GetSurrname() string {
	if x != nil {
		return x.Surrname
	}
	return ""
}

func (x *AddTicketRequest) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *AddTicketRequest) GetDiscount() uint32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

type AddTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddTicketResponse) Reset() {
	*x = AddTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTicketResponse) ProtoMessage() {}

func (x *AddTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTicketResponse.ProtoReflect.Descriptor instead.
func (*AddTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *AddTicketResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTicketImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketImageRequest) Reset() {
	*x = GetTicketImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketImageRequest) ProtoMessage() {}

func (x *GetTicketImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketImageRequest.ProtoReflect.Descriptor instead.
func (*GetTicketImageRequest) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *GetTicketImageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTicketImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetTicketImageResponse) Reset() {
	*x = GetTicketImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketImageResponse) ProtoMessage() {}

func (x *GetTicketImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketImageResponse.ProtoReflect.Descriptor instead.
func (*GetTicketImageResponse) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *GetTicketImageResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetTicketInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketInfoRequest) Reset() {
	*x = GetTicketInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketInfoRequest) ProtoMessage() {}

func (x *GetTicketInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTicketInfoRequest) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *GetTicketInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// NOT IMPLEMENTED
type GetTicketInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketInfoResponse) Reset() {
	*x = GetTicketInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketInfoResponse) ProtoMessage() {}

func (x *GetTicketInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTicketInfoResponse) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *GetTicketInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTicketRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTicketResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ActivateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ActivateTicketRequest) Reset() {
	*x = ActivateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateTicketRequest) ProtoMessage() {}

func (x *ActivateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateTicketRequest.ProtoReflect.Descriptor instead.
func (*ActivateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *ActivateTicketRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ActivateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ActivateTicketResponse) Reset() {
	*x = ActivateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateTicketResponse) ProtoMessage() {}

func (x *ActivateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateTicketResponse.ProtoReflect.Descriptor instead.
func (*ActivateTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *ActivateTicketResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActivateTicketResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_ticket_ticket_proto protoreflect.FileDescriptor

var file_ticket_ticket_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0xaa, 0x01,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x72, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x72, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x40, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x8c, 0x03, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x70, 0x6f, 0x73, 0x6f, 0x6b, 0x68, 0x6f, 0x76, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_ticket_proto_rawDescOnce sync.Once
	file_ticket_ticket_proto_rawDescData = file_ticket_ticket_proto_rawDesc
)

func file_ticket_ticket_proto_rawDescGZIP() []byte {
	file_ticket_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_ticket_proto_rawDescData)
	})
	return file_ticket_ticket_proto_rawDescData
}

var file_ticket_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ticket_ticket_proto_goTypes = []interface{}{
	(*AddTicketRequest)(nil),       // 0: ticket.AddTicketRequest
	(*AddTicketResponse)(nil),      // 1: ticket.AddTicketResponse
	(*GetTicketImageRequest)(nil),  // 2: ticket.GetTicketImageRequest
	(*GetTicketImageResponse)(nil), // 3: ticket.GetTicketImageResponse
	(*GetTicketInfoRequest)(nil),   // 4: ticket.GetTicketInfoRequest
	(*GetTicketInfoResponse)(nil),  // 5: ticket.GetTicketInfoResponse
	(*DeleteTicketRequest)(nil),    // 6: ticket.DeleteTicketRequest
	(*DeleteTicketResponse)(nil),   // 7: ticket.DeleteTicketResponse
	(*ActivateTicketRequest)(nil),  // 8: ticket.ActivateTicketRequest
	(*ActivateTicketResponse)(nil), // 9: ticket.ActivateTicketResponse
}
var file_ticket_ticket_proto_depIdxs = []int32{
	0, // 0: ticket.TicketService.AddTicket:input_type -> ticket.AddTicketRequest
	2, // 1: ticket.TicketService.GetTicketImage:input_type -> ticket.GetTicketImageRequest
	4, // 2: ticket.TicketService.GetTicketInfo:input_type -> ticket.GetTicketInfoRequest
	6, // 3: ticket.TicketService.DeleteTicket:input_type -> ticket.DeleteTicketRequest
	8, // 4: ticket.TicketService.ActivateTicket:input_type -> ticket.ActivateTicketRequest
	1, // 5: ticket.TicketService.AddTicket:output_type -> ticket.AddTicketResponse
	3, // 6: ticket.TicketService.GetTicketImage:output_type -> ticket.GetTicketImageResponse
	5, // 7: ticket.TicketService.GetTicketInfo:output_type -> ticket.GetTicketInfoResponse
	7, // 8: ticket.TicketService.DeleteTicket:output_type -> ticket.DeleteTicketResponse
	9, // 9: ticket.TicketService.ActivateTicket:output_type -> ticket.ActivateTicketResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ticket_ticket_proto_init() }
func file_ticket_ticket_proto_init() {
	if File_ticket_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTicketRequest); i {
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
		file_ticket_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTicketResponse); i {
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
		file_ticket_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketImageRequest); i {
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
		file_ticket_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketImageResponse); i {
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
		file_ticket_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketInfoRequest); i {
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
		file_ticket_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketInfoResponse); i {
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
		file_ticket_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketRequest); i {
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
		file_ticket_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketResponse); i {
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
		file_ticket_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateTicketRequest); i {
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
		file_ticket_ticket_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateTicketResponse); i {
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
			RawDescriptor: file_ticket_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_ticket_proto_msgTypes,
	}.Build()
	File_ticket_ticket_proto = out.File
	file_ticket_ticket_proto_rawDesc = nil
	file_ticket_ticket_proto_goTypes = nil
	file_ticket_ticket_proto_depIdxs = nil
}
