// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: user.proto

package pb

import (
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

// 注册相关
type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 登陆相关
type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 修改用户信息
type AlterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Bio           string                 `protobuf:"bytes,3,opt,name=Bio,proto3" json:"Bio,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Avatar        string                 `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterRequest) Reset() {
	*x = AlterRequest{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterRequest) ProtoMessage() {}

func (x *AlterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterRequest.ProtoReflect.Descriptor instead.
func (*AlterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *AlterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AlterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AlterRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *AlterRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *AlterRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type AlterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlterResponse) Reset() {
	*x = AlterResponse{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterResponse) ProtoMessage() {}

func (x *AlterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterResponse.ProtoReflect.Descriptor instead.
func (*AlterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *AlterResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AlterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 查看个人主页
type PersonalPageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonalPageRequest) Reset() {
	*x = PersonalPageRequest{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonalPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalPageRequest) ProtoMessage() {}

func (x *PersonalPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalPageRequest.ProtoReflect.Descriptor instead.
func (*PersonalPageRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *PersonalPageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Do struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,3,opt,name=IsPrivate,proto3" json:"IsPrivate,omitempty"`
	CreateAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Do) Reset() {
	*x = Do{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Do) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Do) ProtoMessage() {}

func (x *Do) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Do.ProtoReflect.Descriptor instead.
func (*Do) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *Do) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Do) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Do) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Do) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Do) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Avatar        string                 `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Bio           string                 `protobuf:"bytes,4,opt,name=Bio,proto3" json:"Bio,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Documents     []*Do                  `protobuf:"bytes,8,rep,name=Documents,proto3" json:"Documents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UserInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserInfo) GetDocuments() []*Do {
	if x != nil {
		return x.Documents
	}
	return nil
}

type PersonalPageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int64                  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	User          *UserInfo              `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonalPageResponse) Reset() {
	*x = PersonalPageResponse{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonalPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalPageResponse) ProtoMessage() {}

func (x *PersonalPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalPageResponse.ProtoReflect.Descriptor instead.
func (*PersonalPageResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *PersonalPageResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PersonalPageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PersonalPageResponse) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x35, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x35, 0x0a, 0x0d,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0x2d, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x02, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x8f, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42,
	0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21,
	0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x03, 0x2e, 0x44, 0x6f, 0x52, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x5b, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12,
	0x1d, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x32, 0xc4,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_proto_goTypes = []any{
	(*RegisterRequest)(nil),       // 0: RegisterRequest
	(*RegisterResponse)(nil),      // 1: RegisterResponse
	(*LoginRequest)(nil),          // 2: LoginRequest
	(*LoginResponse)(nil),         // 3: LoginResponse
	(*AlterRequest)(nil),          // 4: AlterRequest
	(*AlterResponse)(nil),         // 5: AlterResponse
	(*PersonalPageRequest)(nil),   // 6: PersonalPageRequest
	(*Do)(nil),                    // 7: Do
	(*UserInfo)(nil),              // 8: UserInfo
	(*PersonalPageResponse)(nil),  // 9: PersonalPageResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	10, // 0: Do.CreateAt:type_name -> google.protobuf.Timestamp
	10, // 1: Do.UpdatedAt:type_name -> google.protobuf.Timestamp
	10, // 2: UserInfo.CreatedAt:type_name -> google.protobuf.Timestamp
	10, // 3: UserInfo.UpdatedAt:type_name -> google.protobuf.Timestamp
	7,  // 4: UserInfo.Documents:type_name -> Do
	8,  // 5: PersonalPageResponse.User:type_name -> UserInfo
	0,  // 6: User.Register:input_type -> RegisterRequest
	2,  // 7: User.Login:input_type -> LoginRequest
	4,  // 8: User.Alter:input_type -> AlterRequest
	6,  // 9: User.PersonalPage:input_type -> PersonalPageRequest
	1,  // 10: User.Register:output_type -> RegisterResponse
	3,  // 11: User.Login:output_type -> LoginResponse
	5,  // 12: User.Alter:output_type -> AlterResponse
	9,  // 13: User.PersonalPage:output_type -> PersonalPageResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
