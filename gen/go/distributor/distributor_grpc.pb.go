// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: distributor/distributor.proto

package distributorv1

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

// DistributorServiceClient is the client API for DistributorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistributorServiceClient interface {
	ReplenishBalance(ctx context.Context, in *ReplenishBalanceRequest, opts ...grpc.CallOption) (*ReplenishBalanceResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
}

type distributorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistributorServiceClient(cc grpc.ClientConnInterface) DistributorServiceClient {
	return &distributorServiceClient{cc}
}

func (c *distributorServiceClient) ReplenishBalance(ctx context.Context, in *ReplenishBalanceRequest, opts ...grpc.CallOption) (*ReplenishBalanceResponse, error) {
	out := new(ReplenishBalanceResponse)
	err := c.cc.Invoke(ctx, "/distributor.DistributorService/ReplenishBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributorServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/distributor.DistributorService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributorServiceServer is the server API for DistributorService service.
// All implementations must embed UnimplementedDistributorServiceServer
// for forward compatibility
type DistributorServiceServer interface {
	ReplenishBalance(context.Context, *ReplenishBalanceRequest) (*ReplenishBalanceResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	mustEmbedUnimplementedDistributorServiceServer()
}

// UnimplementedDistributorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDistributorServiceServer struct {
}

func (UnimplementedDistributorServiceServer) ReplenishBalance(context.Context, *ReplenishBalanceRequest) (*ReplenishBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplenishBalance not implemented")
}
func (UnimplementedDistributorServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedDistributorServiceServer) mustEmbedUnimplementedDistributorServiceServer() {}

// UnsafeDistributorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistributorServiceServer will
// result in compilation errors.
type UnsafeDistributorServiceServer interface {
	mustEmbedUnimplementedDistributorServiceServer()
}

func RegisterDistributorServiceServer(s grpc.ServiceRegistrar, srv DistributorServiceServer) {
	s.RegisterService(&DistributorService_ServiceDesc, srv)
}

func _DistributorService_ReplenishBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplenishBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServiceServer).ReplenishBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributor.DistributorService/ReplenishBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServiceServer).ReplenishBalance(ctx, req.(*ReplenishBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributorService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributor.DistributorService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DistributorService_ServiceDesc is the grpc.ServiceDesc for DistributorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistributorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "distributor.DistributorService",
	HandlerType: (*DistributorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplenishBalance",
			Handler:    _DistributorService_ReplenishBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _DistributorService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor/distributor.proto",
}
