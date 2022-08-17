// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user_post_service/user_post_service.proto

package user_post_service

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

// UserPostServiceClient is the client API for UserPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPostServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetPostsForFeed(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetUserPosts(ctx context.Context, in *GetUserPostsRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetPostsForLoggedUserProfile(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateUserPost(ctx context.Context, in *CreateUserPostRequest, opts ...grpc.CallOption) (*CreateUserPostResponse, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Like(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Dislike(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type userPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPostServiceClient(cc grpc.ClientConnInterface) UserPostServiceClient {
	return &userPostServiceClient{cc}
}

func (c *userPostServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) GetPostsForFeed(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/GetPostsForFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) GetUserPosts(ctx context.Context, in *GetUserPostsRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/GetUserPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) GetPostsForLoggedUserProfile(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/GetPostsForLoggedUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) CreateUserPost(ctx context.Context, in *CreateUserPostRequest, opts ...grpc.CallOption) (*CreateUserPostResponse, error) {
	out := new(CreateUserPostResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/CreateUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) Like(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) Dislike(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_post_service.UserPostService/Dislike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPostServiceServer is the server API for UserPostService service.
// All implementations must embed UnimplementedUserPostServiceServer
// for forward compatibility
type UserPostServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetPostsForFeed(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetUserPosts(context.Context, *GetUserPostsRequest) (*GetAllResponse, error)
	GetPostsForLoggedUserProfile(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreateUserPost(context.Context, *CreateUserPostRequest) (*CreateUserPostResponse, error)
	AddComment(context.Context, *AddCommentRequest) (*GetResponse, error)
	Like(context.Context, *GetRequest) (*GetResponse, error)
	Dislike(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedUserPostServiceServer()
}

// UnimplementedUserPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserPostServiceServer struct {
}

func (UnimplementedUserPostServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserPostServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserPostServiceServer) GetPostsForFeed(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsForFeed not implemented")
}
func (UnimplementedUserPostServiceServer) GetUserPosts(context.Context, *GetUserPostsRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPosts not implemented")
}
func (UnimplementedUserPostServiceServer) GetPostsForLoggedUserProfile(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsForLoggedUserProfile not implemented")
}
func (UnimplementedUserPostServiceServer) CreateUserPost(context.Context, *CreateUserPostRequest) (*CreateUserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPost not implemented")
}
func (UnimplementedUserPostServiceServer) AddComment(context.Context, *AddCommentRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedUserPostServiceServer) Like(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedUserPostServiceServer) Dislike(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}
func (UnimplementedUserPostServiceServer) mustEmbedUnimplementedUserPostServiceServer() {}

// UnsafeUserPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPostServiceServer will
// result in compilation errors.
type UnsafeUserPostServiceServer interface {
	mustEmbedUnimplementedUserPostServiceServer()
}

func RegisterUserPostServiceServer(s grpc.ServiceRegistrar, srv UserPostServiceServer) {
	s.RegisterService(&UserPostService_ServiceDesc, srv)
}

func _UserPostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_GetPostsForFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).GetPostsForFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/GetPostsForFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).GetPostsForFeed(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_GetUserPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).GetUserPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/GetUserPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).GetUserPosts(ctx, req.(*GetUserPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_GetPostsForLoggedUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).GetPostsForLoggedUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/GetPostsForLoggedUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).GetPostsForLoggedUserProfile(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_CreateUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).CreateUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/CreateUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).CreateUserPost(ctx, req.(*CreateUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Like(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_post_service.UserPostService/Dislike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Dislike(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPostService_ServiceDesc is the grpc.ServiceDesc for UserPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_post_service.UserPostService",
	HandlerType: (*UserPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserPostService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserPostService_GetAll_Handler,
		},
		{
			MethodName: "GetPostsForFeed",
			Handler:    _UserPostService_GetPostsForFeed_Handler,
		},
		{
			MethodName: "GetUserPosts",
			Handler:    _UserPostService_GetUserPosts_Handler,
		},
		{
			MethodName: "GetPostsForLoggedUserProfile",
			Handler:    _UserPostService_GetPostsForLoggedUserProfile_Handler,
		},
		{
			MethodName: "CreateUserPost",
			Handler:    _UserPostService_CreateUserPost_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _UserPostService_AddComment_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _UserPostService_Like_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _UserPostService_Dislike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_post_service/user_post_service.proto",
}
