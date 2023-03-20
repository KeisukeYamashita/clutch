// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: shortlink/v1/shortlink.proto

package shortlinkv1

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
	ShortlinkAPI_Create_FullMethodName = "/clutch.shortlink.v1.ShortlinkAPI/Create"
	ShortlinkAPI_Get_FullMethodName    = "/clutch.shortlink.v1.ShortlinkAPI/Get"
)

// ShortlinkAPIClient is the client API for ShortlinkAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortlinkAPIClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type shortlinkAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewShortlinkAPIClient(cc grpc.ClientConnInterface) ShortlinkAPIClient {
	return &shortlinkAPIClient{cc}
}

func (c *shortlinkAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ShortlinkAPI_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortlinkAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ShortlinkAPI_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortlinkAPIServer is the server API for ShortlinkAPI service.
// All implementations should embed UnimplementedShortlinkAPIServer
// for forward compatibility
type ShortlinkAPIServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedShortlinkAPIServer should be embedded to have forward compatible implementations.
type UnimplementedShortlinkAPIServer struct {
}

func (UnimplementedShortlinkAPIServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShortlinkAPIServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeShortlinkAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortlinkAPIServer will
// result in compilation errors.
type UnsafeShortlinkAPIServer interface {
	mustEmbedUnimplementedShortlinkAPIServer()
}

func RegisterShortlinkAPIServer(s grpc.ServiceRegistrar, srv ShortlinkAPIServer) {
	s.RegisterService(&ShortlinkAPI_ServiceDesc, srv)
}

func _ShortlinkAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlinkAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortlinkAPI_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlinkAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortlinkAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlinkAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortlinkAPI_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlinkAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortlinkAPI_ServiceDesc is the grpc.ServiceDesc for ShortlinkAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortlinkAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.shortlink.v1.ShortlinkAPI",
	HandlerType: (*ShortlinkAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShortlinkAPI_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ShortlinkAPI_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortlink/v1/shortlink.proto",
}
