// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user_connection_service/user_connection_service.proto

package user_connection_service

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

// UserConnectionServiceClient is the client API for UserConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserConnectionServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	RegisterUserConnection(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Follow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	Unfollow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	AcceptConnectionRequest(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	DeclineConnectionRequest(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	GetConnectionsByUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*Connections, error)
	BlockUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	UnblockUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error)
	GetById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserConnection, error)
}

type userConnectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserConnectionServiceClient(cc grpc.ClientConnInterface) UserConnectionServiceClient {
	return &userConnectionServiceClient{cc}
}

func (c *userConnectionServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) RegisterUserConnection(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/RegisterUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) Follow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) Unfollow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/Unfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) AcceptConnectionRequest(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/AcceptConnectionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) DeclineConnectionRequest(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/DeclineConnectionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) GetConnectionsByUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*Connections, error) {
	out := new(Connections)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/GetConnectionsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) BlockUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) UnblockUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ConnectionsResponse, error) {
	out := new(ConnectionsResponse)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/UnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConnectionServiceClient) GetById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserConnection, error) {
	out := new(UserConnection)
	err := c.cc.Invoke(ctx, "/user_connection_service.UserConnectionService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserConnectionServiceServer is the server API for UserConnectionService service.
// All implementations must embed UnimplementedUserConnectionServiceServer
// for forward compatibility
type UserConnectionServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	RegisterUserConnection(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Follow(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	Unfollow(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	AcceptConnectionRequest(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	DeclineConnectionRequest(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	GetConnectionsByUser(context.Context, *UserIdRequest) (*Connections, error)
	BlockUser(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	UnblockUser(context.Context, *UserIdRequest) (*ConnectionsResponse, error)
	GetById(context.Context, *UserIdRequest) (*UserConnection, error)
	mustEmbedUnimplementedUserConnectionServiceServer()
}

// UnimplementedUserConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserConnectionServiceServer struct {
}

func (UnimplementedUserConnectionServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserConnectionServiceServer) RegisterUserConnection(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserConnection not implemented")
}
func (UnimplementedUserConnectionServiceServer) Follow(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUserConnectionServiceServer) Unfollow(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedUserConnectionServiceServer) AcceptConnectionRequest(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptConnectionRequest not implemented")
}
func (UnimplementedUserConnectionServiceServer) DeclineConnectionRequest(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineConnectionRequest not implemented")
}
func (UnimplementedUserConnectionServiceServer) GetConnectionsByUser(context.Context, *UserIdRequest) (*Connections, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionsByUser not implemented")
}
func (UnimplementedUserConnectionServiceServer) BlockUser(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedUserConnectionServiceServer) UnblockUser(context.Context, *UserIdRequest) (*ConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (UnimplementedUserConnectionServiceServer) GetById(context.Context, *UserIdRequest) (*UserConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserConnectionServiceServer) mustEmbedUnimplementedUserConnectionServiceServer() {}

// UnsafeUserConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserConnectionServiceServer will
// result in compilation errors.
type UnsafeUserConnectionServiceServer interface {
	mustEmbedUnimplementedUserConnectionServiceServer()
}

func RegisterUserConnectionServiceServer(s grpc.ServiceRegistrar, srv UserConnectionServiceServer) {
	s.RegisterService(&UserConnectionService_ServiceDesc, srv)
}

func _UserConnectionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_RegisterUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).RegisterUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/RegisterUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).RegisterUserConnection(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).Follow(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).Unfollow(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_AcceptConnectionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).AcceptConnectionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/AcceptConnectionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).AcceptConnectionRequest(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_DeclineConnectionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).DeclineConnectionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/DeclineConnectionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).DeclineConnectionRequest(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_GetConnectionsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).GetConnectionsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/GetConnectionsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).GetConnectionsByUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).BlockUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/UnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).UnblockUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConnectionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConnectionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_connection_service.UserConnectionService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConnectionServiceServer).GetById(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserConnectionService_ServiceDesc is the grpc.ServiceDesc for UserConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_connection_service.UserConnectionService",
	HandlerType: (*UserConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _UserConnectionService_GetAll_Handler,
		},
		{
			MethodName: "RegisterUserConnection",
			Handler:    _UserConnectionService_RegisterUserConnection_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UserConnectionService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _UserConnectionService_Unfollow_Handler,
		},
		{
			MethodName: "AcceptConnectionRequest",
			Handler:    _UserConnectionService_AcceptConnectionRequest_Handler,
		},
		{
			MethodName: "DeclineConnectionRequest",
			Handler:    _UserConnectionService_DeclineConnectionRequest_Handler,
		},
		{
			MethodName: "GetConnectionsByUser",
			Handler:    _UserConnectionService_GetConnectionsByUser_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _UserConnectionService_BlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _UserConnectionService_UnblockUser_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UserConnectionService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_connection_service/user_connection_service.proto",
}
