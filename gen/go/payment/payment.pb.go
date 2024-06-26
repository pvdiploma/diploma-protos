// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payment/payment.proto

package paymentv1

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

type TicketCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventCategoryId int64  `protobuf:"varint,1,opt,name=event_category_id,json=eventCategoryId,proto3" json:"event_category_id,omitempty"`
	Amount          uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TicketCategory) Reset() {
	*x = TicketCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCategory) ProtoMessage() {}

func (x *TicketCategory) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCategory.ProtoReflect.Descriptor instead.
func (*TicketCategory) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *TicketCategory) GetEventCategoryId() int64 {
	if x != nil {
		return x.EventCategoryId
	}
	return 0
}

func (x *TicketCategory) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PurchaseTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WidgetId   int64             `protobuf:"varint,1,opt,name=widget_id,json=widgetId,proto3" json:"widget_id,omitempty"` // must contains info 'bout event id and distributor id
	Name       string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string            `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Patronymic string            `protobuf:"bytes,5,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
	Email      string            `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone      string            `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Categories []*TicketCategory `protobuf:"bytes,8,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *PurchaseTicketRequest) Reset() {
	*x = PurchaseTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketRequest) ProtoMessage() {}

func (x *PurchaseTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseTicketRequest.ProtoReflect.Descriptor instead.
func (*PurchaseTicketRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseTicketRequest) GetWidgetId() int64 {
	if x != nil {
		return x.WidgetId
	}
	return 0
}

func (x *PurchaseTicketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PurchaseTicketRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *PurchaseTicketRequest) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *PurchaseTicketRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PurchaseTicketRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PurchaseTicketRequest) GetCategories() []*TicketCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

type PurchaseTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *PurchaseTicketResponse) Reset() {
	*x = PurchaseTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketResponse) ProtoMessage() {}

func (x *PurchaseTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseTicketResponse.ProtoReflect.Descriptor instead.
func (*PurchaseTicketResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseTicketResponse) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type SendTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId int64 `protobuf:"varint,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *SendTicketRequest) Reset() {
	*x = SendTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTicketRequest) ProtoMessage() {}

func (x *SendTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTicketRequest.ProtoReflect.Descriptor instead.
func (*SendTicketRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *SendTicketRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

type SendTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendTicketResponse) Reset() {
	*x = SendTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTicketResponse) ProtoMessage() {}

func (x *SendTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTicketResponse.ProtoReflect.Descriptor instead.
func (*SendTicketResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *SendTicketResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_payment_payment_proto protoreflect.FileDescriptor

var file_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x54, 0x0a, 0x0e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x28, 0x0a, 0x16, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x12,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xaa, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1f, 0x5a, 0x1d, 0x70, 0x6f, 0x73, 0x6f, 0x6b, 0x68, 0x6f, 0x76, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_payment_proto_rawDescOnce sync.Once
	file_payment_payment_proto_rawDescData = file_payment_payment_proto_rawDesc
)

func file_payment_payment_proto_rawDescGZIP() []byte {
	file_payment_payment_proto_rawDescOnce.Do(func() {
		file_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_payment_proto_rawDescData)
	})
	return file_payment_payment_proto_rawDescData
}

var file_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_payment_payment_proto_goTypes = []interface{}{
	(*TicketCategory)(nil),         // 0: payment.TicketCategory
	(*PurchaseTicketRequest)(nil),  // 1: payment.PurchaseTicketRequest
	(*PurchaseTicketResponse)(nil), // 2: payment.PurchaseTicketResponse
	(*SendTicketRequest)(nil),      // 3: payment.SendTicketRequest
	(*SendTicketResponse)(nil),     // 4: payment.SendTicketResponse
}
var file_payment_payment_proto_depIdxs = []int32{
	0, // 0: payment.PurchaseTicketRequest.categories:type_name -> payment.TicketCategory
	1, // 1: payment.PaymentService.PurchaseTicket:input_type -> payment.PurchaseTicketRequest
	3, // 2: payment.PaymentService.SendTicket:input_type -> payment.SendTicketRequest
	2, // 3: payment.PaymentService.PurchaseTicket:output_type -> payment.PurchaseTicketResponse
	4, // 4: payment.PaymentService.SendTicket:output_type -> payment.SendTicketResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_payment_proto_init() }
func file_payment_payment_proto_init() {
	if File_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCategory); i {
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
		file_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseTicketRequest); i {
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
		file_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseTicketResponse); i {
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
		file_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTicketRequest); i {
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
		file_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTicketResponse); i {
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
			RawDescriptor: file_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_payment_proto_goTypes,
		DependencyIndexes: file_payment_payment_proto_depIdxs,
		MessageInfos:      file_payment_payment_proto_msgTypes,
	}.Build()
	File_payment_payment_proto = out.File
	file_payment_payment_proto_rawDesc = nil
	file_payment_payment_proto_goTypes = nil
	file_payment_payment_proto_depIdxs = nil
}
