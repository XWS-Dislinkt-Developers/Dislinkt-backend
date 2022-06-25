// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user_service/user_service.proto

package user_service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	ConfirmAccount(ctx context.Context, in *ConfirmAccountRequest, opts ...grpc.CallOption) (*ConfirmAccountResponse, error)
	UpdatePersonalData(ctx context.Context, in *UpdatePersonalDataRequest, opts ...grpc.CallOption) (*UpdatePersonalDataResponse, error)
	UpdateUserWorkEducation(ctx context.Context, in *UpdateUserWAERequest, opts ...grpc.CallOption) (*UpdateUserWAEResponse, error)
	UpdateUserSkillsInterests(ctx context.Context, in *UpdateUserSAIRequest, opts ...grpc.CallOption) (*UpdateUserSAIResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmAccount(ctx context.Context, in *ConfirmAccountRequest, opts ...grpc.CallOption) (*ConfirmAccountResponse, error) {
	out := new(ConfirmAccountResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/ConfirmAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePersonalData(ctx context.Context, in *UpdatePersonalDataRequest, opts ...grpc.CallOption) (*UpdatePersonalDataResponse, error) {
	out := new(UpdatePersonalDataResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/UpdatePersonalData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserWorkEducation(ctx context.Context, in *UpdateUserWAERequest, opts ...grpc.CallOption) (*UpdateUserWAEResponse, error) {
	out := new(UpdateUserWAEResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/UpdateUserWorkEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserSkillsInterests(ctx context.Context, in *UpdateUserSAIRequest, opts ...grpc.CallOption) (*UpdateUserSAIResponse, error) {
	out := new(UpdateUserSAIResponse)
	err := c.cc.Invoke(ctx, "/user_service.UserService/UpdateUserSkillsInterests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	ConfirmAccount(context.Context, *ConfirmAccountRequest) (*ConfirmAccountResponse, error)
	UpdatePersonalData(context.Context, *UpdatePersonalDataRequest) (*UpdatePersonalDataResponse, error)
	UpdateUserWorkEducation(context.Context, *UpdateUserWAERequest) (*UpdateUserWAEResponse, error)
	UpdateUserSkillsInterests(context.Context, *UpdateUserSAIRequest) (*UpdateUserSAIResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ConfirmAccount(context.Context, *ConfirmAccountRequest) (*ConfirmAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmAccount not implemented")
}
func (UnimplementedUserServiceServer) UpdatePersonalData(context.Context, *UpdatePersonalDataRequest) (*UpdatePersonalDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalData not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserWorkEducation(context.Context, *UpdateUserWAERequest) (*UpdateUserWAEResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserWorkEducation not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserSkillsInterests(context.Context, *UpdateUserSAIRequest) (*UpdateUserSAIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSkillsInterests not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/ConfirmAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmAccount(ctx, req.(*ConfirmAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePersonalData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePersonalData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/UpdatePersonalData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePersonalData(ctx, req.(*UpdatePersonalDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserWorkEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserWAERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserWorkEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/UpdateUserWorkEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserWorkEducation(ctx, req.(*UpdateUserWAERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserSkillsInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSAIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserSkillsInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/UpdateUserSkillsInterests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserSkillsInterests(ctx, req.(*UpdateUserSAIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ConfirmAccount",
			Handler:    _UserService_ConfirmAccount_Handler,
		},
		{
			MethodName: "UpdatePersonalData",
			Handler:    _UserService_UpdatePersonalData_Handler,
		},
		{
			MethodName: "UpdateUserWorkEducation",
			Handler:    _UserService_UpdateUserWorkEducation_Handler,
		},
		{
			MethodName: "UpdateUserSkillsInterests",
			Handler:    _UserService_UpdateUserSkillsInterests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service/user_service.proto",
}
