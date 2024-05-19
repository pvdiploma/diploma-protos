// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: deal/deal.proto

package dealsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DealsServiceClient is the client API for DealsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DealsServiceClient interface {
	OfferDeal(ctx context.Context, in *OfferDealRequest, opts ...grpc.CallOption) (*OfferDealResponse, error)
	AcceptDeal(ctx context.Context, in *AcceptDealRequest, opts ...grpc.CallOption) (*AcceptDealResponse, error)
	RejectDeal(ctx context.Context, in *RejectDealRequest, opts ...grpc.CallOption) (*RejectDealResponse, error)
	GetSentDeals(ctx context.Context, in *GetSentDealsRequest, opts ...grpc.CallOption) (*GetSentDealsResponse, error)
	GetProposedDeals(ctx context.Context, in *GetProposedDealsRequest, opts ...grpc.CallOption) (*GetProposedDealsResponse, error)
}

type dealsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDealsServiceClient(cc grpc.ClientConnInterface) DealsServiceClient {
	return &dealsServiceClient{cc}
}

func (c *dealsServiceClient) OfferDeal(ctx context.Context, in *OfferDealRequest, opts ...grpc.CallOption) (*OfferDealResponse, error) {
	out := new(OfferDealResponse)
	err := c.cc.Invoke(ctx, "/deals.DealsService/OfferDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealsServiceClient) AcceptDeal(ctx context.Context, in *AcceptDealRequest, opts ...grpc.CallOption) (*AcceptDealResponse, error) {
	out := new(AcceptDealResponse)
	err := c.cc.Invoke(ctx, "/deals.DealsService/AcceptDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealsServiceClient) RejectDeal(ctx context.Context, in *RejectDealRequest, opts ...grpc.CallOption) (*RejectDealResponse, error) {
	out := new(RejectDealResponse)
	err := c.cc.Invoke(ctx, "/deals.DealsService/RejectDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealsServiceClient) GetSentDeals(ctx context.Context, in *GetSentDealsRequest, opts ...grpc.CallOption) (*GetSentDealsResponse, error) {
	out := new(GetSentDealsResponse)
	err := c.cc.Invoke(ctx, "/deals.DealsService/GetSentDeals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealsServiceClient) GetProposedDeals(ctx context.Context, in *GetProposedDealsRequest, opts ...grpc.CallOption) (*GetProposedDealsResponse, error) {
	out := new(GetProposedDealsResponse)
	err := c.cc.Invoke(ctx, "/deals.DealsService/GetProposedDeals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DealsServiceServer is the server API for DealsService service.
// All implementations must embed UnimplementedDealsServiceServer
// for forward compatibility
type DealsServiceServer interface {
	OfferDeal(context.Context, *OfferDealRequest) (*OfferDealResponse, error)
	AcceptDeal(context.Context, *AcceptDealRequest) (*AcceptDealResponse, error)
	RejectDeal(context.Context, *RejectDealRequest) (*RejectDealResponse, error)
	GetSentDeals(context.Context, *GetSentDealsRequest) (*GetSentDealsResponse, error)
	GetProposedDeals(context.Context, *GetProposedDealsRequest) (*GetProposedDealsResponse, error)
	mustEmbedUnimplementedDealsServiceServer()
}

// UnimplementedDealsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDealsServiceServer struct {
}

func (UnimplementedDealsServiceServer) OfferDeal(context.Context, *OfferDealRequest) (*OfferDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfferDeal not implemented")
}
func (UnimplementedDealsServiceServer) AcceptDeal(context.Context, *AcceptDealRequest) (*AcceptDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDeal not implemented")
}
func (UnimplementedDealsServiceServer) RejectDeal(context.Context, *RejectDealRequest) (*RejectDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectDeal not implemented")
}
func (UnimplementedDealsServiceServer) GetSentDeals(context.Context, *GetSentDealsRequest) (*GetSentDealsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSentDeals not implemented")
}
func (UnimplementedDealsServiceServer) GetProposedDeals(context.Context, *GetProposedDealsRequest) (*GetProposedDealsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProposedDeals not implemented")
}
func (UnimplementedDealsServiceServer) mustEmbedUnimplementedDealsServiceServer() {}

// UnsafeDealsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DealsServiceServer will
// result in compilation errors.
type UnsafeDealsServiceServer interface {
	mustEmbedUnimplementedDealsServiceServer()
}

func RegisterDealsServiceServer(s grpc.ServiceRegistrar, srv DealsServiceServer) {
	s.RegisterService(&DealsService_ServiceDesc, srv)
}

func _DealsService_OfferDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfferDealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealsServiceServer).OfferDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deals.DealsService/OfferDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealsServiceServer).OfferDeal(ctx, req.(*OfferDealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealsService_AcceptDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptDealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealsServiceServer).AcceptDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deals.DealsService/AcceptDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealsServiceServer).AcceptDeal(ctx, req.(*AcceptDealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealsService_RejectDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectDealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealsServiceServer).RejectDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deals.DealsService/RejectDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealsServiceServer).RejectDeal(ctx, req.(*RejectDealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealsService_GetSentDeals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSentDealsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealsServiceServer).GetSentDeals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deals.DealsService/GetSentDeals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealsServiceServer).GetSentDeals(ctx, req.(*GetSentDealsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealsService_GetProposedDeals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProposedDealsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealsServiceServer).GetProposedDeals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deals.DealsService/GetProposedDeals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealsServiceServer).GetProposedDeals(ctx, req.(*GetProposedDealsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DealsService_ServiceDesc is the grpc.ServiceDesc for DealsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DealsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deals.DealsService",
	HandlerType: (*DealsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OfferDeal",
			Handler:    _DealsService_OfferDeal_Handler,
		},
		{
			MethodName: "AcceptDeal",
			Handler:    _DealsService_AcceptDeal_Handler,
		},
		{
			MethodName: "RejectDeal",
			Handler:    _DealsService_RejectDeal_Handler,
		},
		{
			MethodName: "GetSentDeals",
			Handler:    _DealsService_GetSentDeals_Handler,
		},
		{
			MethodName: "GetProposedDeals",
			Handler:    _DealsService_GetProposedDeals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deal/deal.proto",
}
