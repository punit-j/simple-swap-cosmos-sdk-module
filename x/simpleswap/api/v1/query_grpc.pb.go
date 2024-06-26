// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/simpleswap/v1/query.proto

package simpleswapv1

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
	Query_Params_FullMethodName            = "/cosmos.simpleswap.v1.Query/Params"
	Query_Pool_FullMethodName              = "/cosmos.simpleswap.v1.Query/Pool"
	Query_LiquidityProvider_FullMethodName = "/cosmos.simpleswap.v1.Query/LiquidityProvider"
	Query_CoinReserve_FullMethodName       = "/cosmos.simpleswap.v1.Query/CoinReserve"
	Query_CoinReserves_FullMethodName      = "/cosmos.simpleswap.v1.Query/CoinReserves"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// GetPool returns the pool information.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// GetLiquidityProvider returns the liquidity provider information.
	LiquidityProvider(ctx context.Context, in *QueryLiquidityProviderRequest, opts ...grpc.CallOption) (*QueryLiquidityProviderResponse, error)
	// GetCoinReserve returns the coin reserve information for a given coin denom.
	CoinReserve(ctx context.Context, in *QueryCoinReserveRequest, opts ...grpc.CallOption) (*QueryCoinReserveResponse, error)
	// GetCoinReserves returns the coin reserves information for all coin denoms.
	CoinReserves(ctx context.Context, in *QueryCoinReservesRequest, opts ...grpc.CallOption) (*QueryCoinReservesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_Pool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidityProvider(ctx context.Context, in *QueryLiquidityProviderRequest, opts ...grpc.CallOption) (*QueryLiquidityProviderResponse, error) {
	out := new(QueryLiquidityProviderResponse)
	err := c.cc.Invoke(ctx, Query_LiquidityProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CoinReserve(ctx context.Context, in *QueryCoinReserveRequest, opts ...grpc.CallOption) (*QueryCoinReserveResponse, error) {
	out := new(QueryCoinReserveResponse)
	err := c.cc.Invoke(ctx, Query_CoinReserve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CoinReserves(ctx context.Context, in *QueryCoinReservesRequest, opts ...grpc.CallOption) (*QueryCoinReservesResponse, error) {
	out := new(QueryCoinReservesResponse)
	err := c.cc.Invoke(ctx, Query_CoinReserves_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params returns the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// GetPool returns the pool information.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// GetLiquidityProvider returns the liquidity provider information.
	LiquidityProvider(context.Context, *QueryLiquidityProviderRequest) (*QueryLiquidityProviderResponse, error)
	// GetCoinReserve returns the coin reserve information for a given coin denom.
	CoinReserve(context.Context, *QueryCoinReserveRequest) (*QueryCoinReserveResponse, error)
	// GetCoinReserves returns the coin reserves information for all coin denoms.
	CoinReserves(context.Context, *QueryCoinReservesRequest) (*QueryCoinReservesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) LiquidityProvider(context.Context, *QueryLiquidityProviderRequest) (*QueryLiquidityProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityProvider not implemented")
}
func (UnimplementedQueryServer) CoinReserve(context.Context, *QueryCoinReserveRequest) (*QueryCoinReserveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinReserve not implemented")
}
func (UnimplementedQueryServer) CoinReserves(context.Context, *QueryCoinReservesRequest) (*QueryCoinReservesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinReserves not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidityProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidityProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidityProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_LiquidityProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidityProvider(ctx, req.(*QueryLiquidityProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CoinReserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoinReserveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinReserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CoinReserve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinReserve(ctx, req.(*QueryCoinReserveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CoinReserves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoinReservesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinReserves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CoinReserves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinReserves(ctx, req.(*QueryCoinReservesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.simpleswap.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "LiquidityProvider",
			Handler:    _Query_LiquidityProvider_Handler,
		},
		{
			MethodName: "CoinReserve",
			Handler:    _Query_CoinReserve_Handler,
		},
		{
			MethodName: "CoinReserves",
			Handler:    _Query_CoinReserves_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/simpleswap/v1/query.proto",
}
