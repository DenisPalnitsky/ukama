// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: lookup.proto

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

// LookupServiceClient is the client API for LookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LookupServiceClient interface {
	// Orgs
	AddOrg(ctx context.Context, in *AddOrgRequest, opts ...grpc.CallOption) (*AddOrgResponse, error)
	UpdateOrg(ctx context.Context, in *UpdateOrgRequest, opts ...grpc.CallOption) (*UpdateOrgResponse, error)
	GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*GetOrgResponse, error)
	// For Node bootstarping
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// For other systems and debigging purpose
	AddNodeForOrg(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error)
	GetNodeForOrg(ctx context.Context, in *GetNodeForOrgRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	DeleteNodeForOrg(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// System
	GetSystemForOrg(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*GetSystemResponse, error)
	AddSystemForOrg(ctx context.Context, in *AddSystemRequest, opts ...grpc.CallOption) (*AddSystemResponse, error)
	UpdateSystemForOrg(ctx context.Context, in *UpdateSystemRequest, opts ...grpc.CallOption) (*UpdateSystemResponse, error)
	DeleteSystemForOrg(ctx context.Context, in *DeleteSystemRequest, opts ...grpc.CallOption) (*DeleteSystemResponse, error)
}

type lookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupServiceClient(cc grpc.ClientConnInterface) LookupServiceClient {
	return &lookupServiceClient{cc}
}

func (c *lookupServiceClient) AddOrg(ctx context.Context, in *AddOrgRequest, opts ...grpc.CallOption) (*AddOrgResponse, error) {
	out := new(AddOrgResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/AddOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) UpdateOrg(ctx context.Context, in *UpdateOrgRequest, opts ...grpc.CallOption) (*UpdateOrgResponse, error) {
	out := new(UpdateOrgResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/UpdateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) GetOrg(ctx context.Context, in *GetOrgRequest, opts ...grpc.CallOption) (*GetOrgResponse, error) {
	out := new(GetOrgResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) AddNodeForOrg(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error) {
	out := new(AddNodeResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/AddNodeForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) GetNodeForOrg(ctx context.Context, in *GetNodeForOrgRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/GetNodeForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) DeleteNodeForOrg(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/DeleteNodeForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) GetSystemForOrg(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*GetSystemResponse, error) {
	out := new(GetSystemResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/GetSystemForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) AddSystemForOrg(ctx context.Context, in *AddSystemRequest, opts ...grpc.CallOption) (*AddSystemResponse, error) {
	out := new(AddSystemResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/AddSystemForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) UpdateSystemForOrg(ctx context.Context, in *UpdateSystemRequest, opts ...grpc.CallOption) (*UpdateSystemResponse, error) {
	out := new(UpdateSystemResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/UpdateSystemForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookupServiceClient) DeleteSystemForOrg(ctx context.Context, in *DeleteSystemRequest, opts ...grpc.CallOption) (*DeleteSystemResponse, error) {
	out := new(DeleteSystemResponse)
	err := c.cc.Invoke(ctx, "/ukama.lookup.v1.LookupService/DeleteSystemForOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LookupServiceServer is the server API for LookupService service.
// All implementations must embed UnimplementedLookupServiceServer
// for forward compatibility
type LookupServiceServer interface {
	// Orgs
	AddOrg(context.Context, *AddOrgRequest) (*AddOrgResponse, error)
	UpdateOrg(context.Context, *UpdateOrgRequest) (*UpdateOrgResponse, error)
	GetOrg(context.Context, *GetOrgRequest) (*GetOrgResponse, error)
	// For Node bootstarping
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// For other systems and debigging purpose
	AddNodeForOrg(context.Context, *AddNodeRequest) (*AddNodeResponse, error)
	GetNodeForOrg(context.Context, *GetNodeForOrgRequest) (*GetNodeResponse, error)
	DeleteNodeForOrg(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	// System
	GetSystemForOrg(context.Context, *GetSystemRequest) (*GetSystemResponse, error)
	AddSystemForOrg(context.Context, *AddSystemRequest) (*AddSystemResponse, error)
	UpdateSystemForOrg(context.Context, *UpdateSystemRequest) (*UpdateSystemResponse, error)
	DeleteSystemForOrg(context.Context, *DeleteSystemRequest) (*DeleteSystemResponse, error)
	mustEmbedUnimplementedLookupServiceServer()
}

// UnimplementedLookupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLookupServiceServer struct {
}

func (UnimplementedLookupServiceServer) AddOrg(context.Context, *AddOrgRequest) (*AddOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrg not implemented")
}
func (UnimplementedLookupServiceServer) UpdateOrg(context.Context, *UpdateOrgRequest) (*UpdateOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrg not implemented")
}
func (UnimplementedLookupServiceServer) GetOrg(context.Context, *GetOrgRequest) (*GetOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedLookupServiceServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedLookupServiceServer) AddNodeForOrg(context.Context, *AddNodeRequest) (*AddNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNodeForOrg not implemented")
}
func (UnimplementedLookupServiceServer) GetNodeForOrg(context.Context, *GetNodeForOrgRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeForOrg not implemented")
}
func (UnimplementedLookupServiceServer) DeleteNodeForOrg(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeForOrg not implemented")
}
func (UnimplementedLookupServiceServer) GetSystemForOrg(context.Context, *GetSystemRequest) (*GetSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemForOrg not implemented")
}
func (UnimplementedLookupServiceServer) AddSystemForOrg(context.Context, *AddSystemRequest) (*AddSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSystemForOrg not implemented")
}
func (UnimplementedLookupServiceServer) UpdateSystemForOrg(context.Context, *UpdateSystemRequest) (*UpdateSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSystemForOrg not implemented")
}
func (UnimplementedLookupServiceServer) DeleteSystemForOrg(context.Context, *DeleteSystemRequest) (*DeleteSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSystemForOrg not implemented")
}
func (UnimplementedLookupServiceServer) mustEmbedUnimplementedLookupServiceServer() {}

// UnsafeLookupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LookupServiceServer will
// result in compilation errors.
type UnsafeLookupServiceServer interface {
	mustEmbedUnimplementedLookupServiceServer()
}

func RegisterLookupServiceServer(s grpc.ServiceRegistrar, srv LookupServiceServer) {
	s.RegisterService(&LookupService_ServiceDesc, srv)
}

func _LookupService_AddOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).AddOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/AddOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).AddOrg(ctx, req.(*AddOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_UpdateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).UpdateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/UpdateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).UpdateOrg(ctx, req.(*UpdateOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/GetOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).GetOrg(ctx, req.(*GetOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_AddNodeForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).AddNodeForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/AddNodeForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).AddNodeForOrg(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_GetNodeForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeForOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).GetNodeForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/GetNodeForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).GetNodeForOrg(ctx, req.(*GetNodeForOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_DeleteNodeForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).DeleteNodeForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/DeleteNodeForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).DeleteNodeForOrg(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_GetSystemForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).GetSystemForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/GetSystemForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).GetSystemForOrg(ctx, req.(*GetSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_AddSystemForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).AddSystemForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/AddSystemForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).AddSystemForOrg(ctx, req.(*AddSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_UpdateSystemForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).UpdateSystemForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/UpdateSystemForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).UpdateSystemForOrg(ctx, req.(*UpdateSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookupService_DeleteSystemForOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).DeleteSystemForOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ukama.lookup.v1.LookupService/DeleteSystemForOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).DeleteSystemForOrg(ctx, req.(*DeleteSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LookupService_ServiceDesc is the grpc.ServiceDesc for LookupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LookupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ukama.lookup.v1.LookupService",
	HandlerType: (*LookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrg",
			Handler:    _LookupService_AddOrg_Handler,
		},
		{
			MethodName: "UpdateOrg",
			Handler:    _LookupService_UpdateOrg_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _LookupService_GetOrg_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _LookupService_GetNode_Handler,
		},
		{
			MethodName: "AddNodeForOrg",
			Handler:    _LookupService_AddNodeForOrg_Handler,
		},
		{
			MethodName: "GetNodeForOrg",
			Handler:    _LookupService_GetNodeForOrg_Handler,
		},
		{
			MethodName: "DeleteNodeForOrg",
			Handler:    _LookupService_DeleteNodeForOrg_Handler,
		},
		{
			MethodName: "GetSystemForOrg",
			Handler:    _LookupService_GetSystemForOrg_Handler,
		},
		{
			MethodName: "AddSystemForOrg",
			Handler:    _LookupService_AddSystemForOrg_Handler,
		},
		{
			MethodName: "UpdateSystemForOrg",
			Handler:    _LookupService_UpdateSystemForOrg_Handler,
		},
		{
			MethodName: "DeleteSystemForOrg",
			Handler:    _LookupService_DeleteSystemForOrg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lookup.proto",
}
