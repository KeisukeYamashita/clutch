// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package feedbackv1

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

// FeedbackAPIClient is the client API for FeedbackAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackAPIClient interface {
	GetSurveys(ctx context.Context, in *GetSurveysRequest, opts ...grpc.CallOption) (*GetSurveysResponse, error)
	SubmitFeedback(ctx context.Context, in *SubmitFeedbackRequest, opts ...grpc.CallOption) (*SubmitFeedbackResponse, error)
}

type feedbackAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackAPIClient(cc grpc.ClientConnInterface) FeedbackAPIClient {
	return &feedbackAPIClient{cc}
}

func (c *feedbackAPIClient) GetSurveys(ctx context.Context, in *GetSurveysRequest, opts ...grpc.CallOption) (*GetSurveysResponse, error) {
	out := new(GetSurveysResponse)
	err := c.cc.Invoke(ctx, "/clutch.feedback.v1.FeedbackAPI/GetSurveys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackAPIClient) SubmitFeedback(ctx context.Context, in *SubmitFeedbackRequest, opts ...grpc.CallOption) (*SubmitFeedbackResponse, error) {
	out := new(SubmitFeedbackResponse)
	err := c.cc.Invoke(ctx, "/clutch.feedback.v1.FeedbackAPI/SubmitFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackAPIServer is the server API for FeedbackAPI service.
// All implementations should embed UnimplementedFeedbackAPIServer
// for forward compatibility
type FeedbackAPIServer interface {
	GetSurveys(context.Context, *GetSurveysRequest) (*GetSurveysResponse, error)
	SubmitFeedback(context.Context, *SubmitFeedbackRequest) (*SubmitFeedbackResponse, error)
}

// UnimplementedFeedbackAPIServer should be embedded to have forward compatible implementations.
type UnimplementedFeedbackAPIServer struct {
}

func (UnimplementedFeedbackAPIServer) GetSurveys(context.Context, *GetSurveysRequest) (*GetSurveysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSurveys not implemented")
}
func (UnimplementedFeedbackAPIServer) SubmitFeedback(context.Context, *SubmitFeedbackRequest) (*SubmitFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitFeedback not implemented")
}

// UnsafeFeedbackAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackAPIServer will
// result in compilation errors.
type UnsafeFeedbackAPIServer interface {
	mustEmbedUnimplementedFeedbackAPIServer()
}

func RegisterFeedbackAPIServer(s grpc.ServiceRegistrar, srv FeedbackAPIServer) {
	s.RegisterService(&FeedbackAPI_ServiceDesc, srv)
}

func _FeedbackAPI_GetSurveys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSurveysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackAPIServer).GetSurveys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.feedback.v1.FeedbackAPI/GetSurveys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackAPIServer).GetSurveys(ctx, req.(*GetSurveysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackAPI_SubmitFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackAPIServer).SubmitFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.feedback.v1.FeedbackAPI/SubmitFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackAPIServer).SubmitFeedback(ctx, req.(*SubmitFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbackAPI_ServiceDesc is the grpc.ServiceDesc for FeedbackAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbackAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.feedback.v1.FeedbackAPI",
	HandlerType: (*FeedbackAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSurveys",
			Handler:    _FeedbackAPI_GetSurveys_Handler,
		},
		{
			MethodName: "SubmitFeedback",
			Handler:    _FeedbackAPI_SubmitFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback/v1/feedback.proto",
}
