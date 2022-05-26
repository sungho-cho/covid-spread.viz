// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: covid_data.proto

package proto

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

// CovidDataClient is the client API for CovidData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CovidDataClient interface {
	GetActiveCases(ctx context.Context, in *GetActiveCasesRequest, opts ...grpc.CallOption) (*GetActiveCasesResponse, error)
	GetDateData(ctx context.Context, in *GetDateDataRequest, opts ...grpc.CallOption) (*GetDateDataResponse, error)
}

type covidDataClient struct {
	cc grpc.ClientConnInterface
}

func NewCovidDataClient(cc grpc.ClientConnInterface) CovidDataClient {
	return &covidDataClient{cc}
}

func (c *covidDataClient) GetActiveCases(ctx context.Context, in *GetActiveCasesRequest, opts ...grpc.CallOption) (*GetActiveCasesResponse, error) {
	out := new(GetActiveCasesResponse)
	err := c.cc.Invoke(ctx, "/proto.CovidData/GetActiveCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *covidDataClient) GetDateData(ctx context.Context, in *GetDateDataRequest, opts ...grpc.CallOption) (*GetDateDataResponse, error) {
	out := new(GetDateDataResponse)
	err := c.cc.Invoke(ctx, "/proto.CovidData/GetDateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CovidDataServer is the server API for CovidData service.
// All implementations must embed UnimplementedCovidDataServer
// for forward compatibility
type CovidDataServer interface {
	GetActiveCases(context.Context, *GetActiveCasesRequest) (*GetActiveCasesResponse, error)
	GetDateData(context.Context, *GetDateDataRequest) (*GetDateDataResponse, error)
	mustEmbedUnimplementedCovidDataServer()
}

// UnimplementedCovidDataServer must be embedded to have forward compatible implementations.
type UnimplementedCovidDataServer struct {
}

func (UnimplementedCovidDataServer) GetActiveCases(context.Context, *GetActiveCasesRequest) (*GetActiveCasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveCases not implemented")
}
func (UnimplementedCovidDataServer) GetDateData(context.Context, *GetDateDataRequest) (*GetDateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDateData not implemented")
}
func (UnimplementedCovidDataServer) mustEmbedUnimplementedCovidDataServer() {}

// UnsafeCovidDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CovidDataServer will
// result in compilation errors.
type UnsafeCovidDataServer interface {
	mustEmbedUnimplementedCovidDataServer()
}

func RegisterCovidDataServer(s grpc.ServiceRegistrar, srv CovidDataServer) {
	s.RegisterService(&CovidData_ServiceDesc, srv)
}

func _CovidData_GetActiveCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CovidDataServer).GetActiveCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CovidData/GetActiveCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CovidDataServer).GetActiveCases(ctx, req.(*GetActiveCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CovidData_GetDateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CovidDataServer).GetDateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CovidData/GetDateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CovidDataServer).GetDateData(ctx, req.(*GetDateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CovidData_ServiceDesc is the grpc.ServiceDesc for CovidData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CovidData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CovidData",
	HandlerType: (*CovidDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActiveCases",
			Handler:    _CovidData_GetActiveCases_Handler,
		},
		{
			MethodName: "GetDateData",
			Handler:    _CovidData_GetDateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "covid_data.proto",
}
