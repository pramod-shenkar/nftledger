// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0--rc1
// source: service/request.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	model "nftledger/internal/adapter/inbound/grpc/generated/model"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RequestService_SaveRequest_FullMethodName   = "/service.RequestService/SaveRequest"
	RequestService_GetRequest_FullMethodName    = "/service.RequestService/GetRequest"
	RequestService_GetRequests_FullMethodName   = "/service.RequestService/GetRequests"
	RequestService_UpdateRequest_FullMethodName = "/service.RequestService/UpdateRequest"
	RequestService_DeleteRequest_FullMethodName = "/service.RequestService/DeleteRequest"
)

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	SaveRequest(ctx context.Context, in *model.SaveRequestRequest, opts ...grpc.CallOption) (*model.SaveRequestResponse, error)
	GetRequest(ctx context.Context, in *model.GetRequestRequest, opts ...grpc.CallOption) (*model.GetRequestResponse, error)
	GetRequests(ctx context.Context, in *model.GetRequestsRequest, opts ...grpc.CallOption) (*model.GetRequestsResponse, error)
	UpdateRequest(ctx context.Context, in *model.UpdateRequestRequest, opts ...grpc.CallOption) (*model.UpdateRequestResponse, error)
	DeleteRequest(ctx context.Context, in *model.DeleteRequestRequest, opts ...grpc.CallOption) (*model.DeleteRequestResponse, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) SaveRequest(ctx context.Context, in *model.SaveRequestRequest, opts ...grpc.CallOption) (*model.SaveRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.SaveRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_SaveRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequest(ctx context.Context, in *model.GetRequestRequest, opts ...grpc.CallOption) (*model.GetRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.GetRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_GetRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequests(ctx context.Context, in *model.GetRequestsRequest, opts ...grpc.CallOption) (*model.GetRequestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.GetRequestsResponse)
	err := c.cc.Invoke(ctx, RequestService_GetRequests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) UpdateRequest(ctx context.Context, in *model.UpdateRequestRequest, opts ...grpc.CallOption) (*model.UpdateRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.UpdateRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_UpdateRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) DeleteRequest(ctx context.Context, in *model.DeleteRequestRequest, opts ...grpc.CallOption) (*model.DeleteRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.DeleteRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_DeleteRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility.
type RequestServiceServer interface {
	SaveRequest(context.Context, *model.SaveRequestRequest) (*model.SaveRequestResponse, error)
	GetRequest(context.Context, *model.GetRequestRequest) (*model.GetRequestResponse, error)
	GetRequests(context.Context, *model.GetRequestsRequest) (*model.GetRequestsResponse, error)
	UpdateRequest(context.Context, *model.UpdateRequestRequest) (*model.UpdateRequestResponse, error)
	DeleteRequest(context.Context, *model.DeleteRequestRequest) (*model.DeleteRequestResponse, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRequestServiceServer struct{}

func (UnimplementedRequestServiceServer) SaveRequest(context.Context, *model.SaveRequestRequest) (*model.SaveRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRequest not implemented")
}
func (UnimplementedRequestServiceServer) GetRequest(context.Context, *model.GetRequestRequest) (*model.GetRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}
func (UnimplementedRequestServiceServer) GetRequests(context.Context, *model.GetRequestsRequest) (*model.GetRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequests not implemented")
}
func (UnimplementedRequestServiceServer) UpdateRequest(context.Context, *model.UpdateRequestRequest) (*model.UpdateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequest not implemented")
}
func (UnimplementedRequestServiceServer) DeleteRequest(context.Context, *model.DeleteRequestRequest) (*model.DeleteRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequest not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}
func (UnimplementedRequestServiceServer) testEmbeddedByValue()                        {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	// If the following call pancis, it indicates UnimplementedRequestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_SaveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.SaveRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).SaveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_SaveRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).SaveRequest(ctx, req.(*model.SaveRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_GetRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequest(ctx, req.(*model.GetRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_GetRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequests(ctx, req.(*model.GetRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_UpdateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.UpdateRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).UpdateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_UpdateRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).UpdateRequest(ctx, req.(*model.UpdateRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_DeleteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).DeleteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_DeleteRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).DeleteRequest(ctx, req.(*model.DeleteRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveRequest",
			Handler:    _RequestService_SaveRequest_Handler,
		},
		{
			MethodName: "GetRequest",
			Handler:    _RequestService_GetRequest_Handler,
		},
		{
			MethodName: "GetRequests",
			Handler:    _RequestService_GetRequests_Handler,
		},
		{
			MethodName: "UpdateRequest",
			Handler:    _RequestService_UpdateRequest_Handler,
		},
		{
			MethodName: "DeleteRequest",
			Handler:    _RequestService_DeleteRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/request.proto",
}
