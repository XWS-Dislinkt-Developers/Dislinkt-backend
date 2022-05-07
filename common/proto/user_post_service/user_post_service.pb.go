// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user_post_service/user_post_service.proto

package user_post_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//TODO: AddCommentToUserPostRequest{...}
type AddCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddComment *AddComment `protobuf:"bytes,1,opt,name=addComment,proto3" json:"addComment,omitempty"`
}

func (x *AddCommentRequest) Reset() {
	*x = AddCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentRequest) ProtoMessage() {}

func (x *AddCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentRequest.ProtoReflect.Descriptor instead.
func (*AddCommentRequest) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddCommentRequest) GetAddComment() *AddComment {
	if x != nil {
		return x.AddComment
	}
	return nil
}

type AddComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPost string `protobuf:"bytes,1,opt,name=idPost,proto3" json:"idPost,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AddComment) Reset() {
	*x = AddComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddComment) ProtoMessage() {}

func (x *AddComment) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddComment.ProtoReflect.Descriptor instead.
func (*AddComment) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddComment) GetIdPost() string {
	if x != nil {
		return x.IdPost
	}
	return ""
}

func (x *AddComment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPost *UserPost `protobuf:"bytes,1,opt,name=userPost,proto3" json:"userPost,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetUserPost() *UserPost {
	if x != nil {
		return x.UserPost
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{4}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPosts []*UserPost `protobuf:"bytes,1,rep,name=userPosts,proto3" json:"userPosts,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllResponse) GetUserPosts() []*UserPost {
	if x != nil {
		return x.UserPosts
	}
	return nil
}

type CreateUserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPost *UserPost `protobuf:"bytes,1,opt,name=userPost,proto3" json:"userPost,omitempty"`
}

func (x *CreateUserPostRequest) Reset() {
	*x = CreateUserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPostRequest) ProtoMessage() {}

func (x *CreateUserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPostRequest.ProtoReflect.Descriptor instead.
func (*CreateUserPostRequest) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserPostRequest) GetUserPost() *UserPost {
	if x != nil {
		return x.UserPost
	}
	return nil
}

type CreateUserPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPost *UserPost `protobuf:"bytes,1,opt,name=userPost,proto3" json:"userPost,omitempty"`
}

func (x *CreateUserPostResponse) Reset() {
	*x = CreateUserPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPostResponse) ProtoMessage() {}

func (x *CreateUserPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPostResponse.ProtoReflect.Descriptor instead.
func (*CreateUserPostResponse) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserPostResponse) GetUserPost() *UserPost {
	if x != nil {
		return x.UserPost
	}
	return nil
}

type UserPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Text      string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	ImagePath string                 `protobuf:"bytes,5,opt,name=imagePath,proto3" json:"imagePath,omitempty"`
	Reactions []*Reaction            `protobuf:"bytes,6,rep,name=reactions,proto3" json:"reactions,omitempty"`
	Comments  []*Comment             `protobuf:"bytes,7,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *UserPost) Reset() {
	*x = UserPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPost) ProtoMessage() {}

func (x *UserPost) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPost.ProtoReflect.Descriptor instead.
func (*UserPost) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserPost) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPost) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPost) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserPost) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UserPost) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *UserPost) GetReactions() []*Reaction {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *UserPost) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Liked    bool  `protobuf:"varint,2,opt,name=liked,proto3" json:"liked,omitempty"`
	Disliked bool  `protobuf:"varint,3,opt,name=disliked,proto3" json:"disliked,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *Reaction) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Reaction) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

func (x *Reaction) GetDisliked() bool {
	if x != nil {
		return x.Disliked
	}
	return false
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Text      string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_service_user_post_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_service_user_post_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_user_post_service_user_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *Comment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_user_post_service_user_post_service_proto protoreflect.FileDescriptor

var file_user_post_service_user_post_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22,
	0x50, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x51, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x36, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0x6f,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32,
	0xc9, 0x03, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x61, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x09, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x3a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x70, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a,
	0x0a, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x4e, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x57, 0x53, 0x2d, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x2d, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72,
	0x73, 0x2f, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_post_service_user_post_service_proto_rawDescOnce sync.Once
	file_user_post_service_user_post_service_proto_rawDescData = file_user_post_service_user_post_service_proto_rawDesc
)

func file_user_post_service_user_post_service_proto_rawDescGZIP() []byte {
	file_user_post_service_user_post_service_proto_rawDescOnce.Do(func() {
		file_user_post_service_user_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_post_service_user_post_service_proto_rawDescData)
	})
	return file_user_post_service_user_post_service_proto_rawDescData
}

var file_user_post_service_user_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_post_service_user_post_service_proto_goTypes = []interface{}{
	(*AddCommentRequest)(nil),      // 0: user_post_service.AddCommentRequest
	(*AddComment)(nil),             // 1: user_post_service.AddComment
	(*GetRequest)(nil),             // 2: user_post_service.GetRequest
	(*GetResponse)(nil),            // 3: user_post_service.GetResponse
	(*GetAllRequest)(nil),          // 4: user_post_service.GetAllRequest
	(*GetAllResponse)(nil),         // 5: user_post_service.GetAllResponse
	(*CreateUserPostRequest)(nil),  // 6: user_post_service.CreateUserPostRequest
	(*CreateUserPostResponse)(nil), // 7: user_post_service.CreateUserPostResponse
	(*UserPost)(nil),               // 8: user_post_service.UserPost
	(*Reaction)(nil),               // 9: user_post_service.Reaction
	(*Comment)(nil),                // 10: user_post_service.Comment
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_user_post_service_user_post_service_proto_depIdxs = []int32{
	1,  // 0: user_post_service.AddCommentRequest.addComment:type_name -> user_post_service.AddComment
	8,  // 1: user_post_service.GetResponse.userPost:type_name -> user_post_service.UserPost
	8,  // 2: user_post_service.GetAllResponse.userPosts:type_name -> user_post_service.UserPost
	8,  // 3: user_post_service.CreateUserPostRequest.userPost:type_name -> user_post_service.UserPost
	8,  // 4: user_post_service.CreateUserPostResponse.userPost:type_name -> user_post_service.UserPost
	11, // 5: user_post_service.UserPost.createdAt:type_name -> google.protobuf.Timestamp
	9,  // 6: user_post_service.UserPost.reactions:type_name -> user_post_service.Reaction
	10, // 7: user_post_service.UserPost.comments:type_name -> user_post_service.Comment
	11, // 8: user_post_service.Comment.createdAt:type_name -> google.protobuf.Timestamp
	2,  // 9: user_post_service.UserPostService.Get:input_type -> user_post_service.GetRequest
	4,  // 10: user_post_service.UserPostService.GetAll:input_type -> user_post_service.GetAllRequest
	6,  // 11: user_post_service.UserPostService.CreateUserPost:input_type -> user_post_service.CreateUserPostRequest
	0,  // 12: user_post_service.UserPostService.AddComment:input_type -> user_post_service.AddCommentRequest
	3,  // 13: user_post_service.UserPostService.Get:output_type -> user_post_service.GetResponse
	5,  // 14: user_post_service.UserPostService.GetAll:output_type -> user_post_service.GetAllResponse
	7,  // 15: user_post_service.UserPostService.CreateUserPost:output_type -> user_post_service.CreateUserPostResponse
	3,  // 16: user_post_service.UserPostService.AddComment:output_type -> user_post_service.GetResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_user_post_service_user_post_service_proto_init() }
func file_user_post_service_user_post_service_proto_init() {
	if File_user_post_service_user_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_post_service_user_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_post_service_user_post_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_post_service_user_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_post_service_user_post_service_proto_goTypes,
		DependencyIndexes: file_user_post_service_user_post_service_proto_depIdxs,
		MessageInfos:      file_user_post_service_user_post_service_proto_msgTypes,
	}.Build()
	File_user_post_service_user_post_service_proto = out.File
	file_user_post_service_user_post_service_proto_rawDesc = nil
	file_user_post_service_user_post_service_proto_goTypes = nil
	file_user_post_service_user_post_service_proto_depIdxs = nil
}
