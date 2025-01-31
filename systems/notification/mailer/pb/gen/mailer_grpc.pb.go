// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: mailer.proto

package gen

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

// MailerServiceClient is the client API for MailerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerServiceClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
	GetEmailById(ctx context.Context, in *GetEmailByIdRequest, opts ...grpc.CallOption) (*GetEmailByIdResponse, error)
}

type mailerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerServiceClient(cc grpc.ClientConnInterface) MailerServiceClient {
	return &mailerServiceClient{cc}
}

func (c *mailerServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/ukama.notification.mailer.v1.MailerService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerServiceClient) GetEmailById(ctx context.Context, in *GetEmailByIdRequest, opts ...grpc.CallOption) (*GetEmailByIdResponse, error) {
	out := new(GetEmailByIdResponse)
	err := c.cc.Invoke(ctx, "/ukama.notification.mailer.v1.MailerService/GetEmailById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServiceServer is the server API for MailerService service.
// All implementations must embed UnimplementedMailerServiceServer
// for forward compatibility
type MailerServiceServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	GetEmailById(context.Context, *GetEmailByIdRequest) (*GetEmailByIdResponse, error)
	mustEmbedUnimplementedMailerServiceServer()
}

// UnimplementedMailerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailerServiceServer struct {
}

func (UnimplementedMailerServiceServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedMailerServiceServer) GetEmailById(context.Context, *GetEmailByIdRequest) (*GetEmailByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailById not implemented")
}
func (UnimplementedMailerServiceServer) mustEmbedUnimplementedMailerServiceServer() {}

// UnsafeMailerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServiceServer will
// result in compilation errors.
type UnsafeMailerServiceServer interface {
	mustEmbedUnimplementedMailerServiceServer()
}

func RegisterMailerServiceServer(s grpc.ServiceRegistrar, srv MailerServiceServer) {
	s.RegisterService(&MailerService_ServiceDesc, srv)
}

func _MailerService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.notification.mailer.v1.MailerService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailerService_GetEmailById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServiceServer).GetEmailById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.notification.mailer.v1.MailerService/GetEmailById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServiceServer).GetEmailById(ctx, req.(*GetEmailByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailerService_ServiceDesc is the grpc.ServiceDesc for MailerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ukama.notification.mailer.v1.MailerService",
	HandlerType: (*MailerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _MailerService_SendEmail_Handler,
		},
		{
			MethodName: "GetEmailById",
			Handler:    _MailerService_GetEmailById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer.proto",
}
