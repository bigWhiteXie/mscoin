// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: market.proto

package market

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

const (
	Market_FindSymbolThumbTrend_FullMethodName = "/market.Market/FindSymbolThumbTrend"
)

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*SymbolThumbRes, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) FindSymbolThumbTrend(ctx context.Context, in *MarketReq, opts ...grpc.CallOption) (*SymbolThumbRes, error) {
	out := new(SymbolThumbRes)
	err := c.cc.Invoke(ctx, Market_FindSymbolThumbTrend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	FindSymbolThumbTrend(context.Context, *MarketReq) (*SymbolThumbRes, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) FindSymbolThumbTrend(context.Context, *MarketReq) (*SymbolThumbRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSymbolThumbTrend not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_FindSymbolThumbTrend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).FindSymbolThumbTrend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_FindSymbolThumbTrend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).FindSymbolThumbTrend(ctx, req.(*MarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSymbolThumbTrend",
			Handler:    _Market_FindSymbolThumbTrend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market.proto",
}
