// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: backend/protos/followships_microservice.proto

package grpc_generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FollowshipOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId int32 `protobuf:"varint,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	FolloweeId int32 `protobuf:"varint,2,opt,name=followee_id,json=followeeId,proto3" json:"followee_id,omitempty"`
}

func (x *FollowshipOperationRequest) Reset() {
	*x = FollowshipOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowshipOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowshipOperationRequest) ProtoMessage() {}

func (x *FollowshipOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowshipOperationRequest.ProtoReflect.Descriptor instead.
func (*FollowshipOperationRequest) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{0}
}

func (x *FollowshipOperationRequest) GetFollowerId() int32 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

func (x *FollowshipOperationRequest) GetFolloweeId() int32 {
	if x != nil {
		return x.FolloweeId
	}
	return 0
}

type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Offset   int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetFollowersRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFollowersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetFollowersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerIds []int32 `protobuf:"varint,1,rep,packed,name=follower_ids,json=followerIds,proto3" json:"follower_ids,omitempty"`
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowersResponse) GetFollowerIds() []int32 {
	if x != nil {
		return x.FollowerIds
	}
	return nil
}

type GetFollowingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Offset   int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetFollowingsRequest) Reset() {
	*x = GetFollowingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingsRequest) ProtoMessage() {}

func (x *GetFollowingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingsRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingsRequest) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowingsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFollowingsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetFollowingsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetFollowingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolloweeIds []int32 `protobuf:"varint,1,rep,packed,name=followee_ids,json=followeeIds,proto3" json:"followee_ids,omitempty"`
}

func (x *GetFollowingsResponse) Reset() {
	*x = GetFollowingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingsResponse) ProtoMessage() {}

func (x *GetFollowingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingsResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingsResponse) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetFollowingsResponse) GetFolloweeIds() []int32 {
	if x != nil {
		return x.FolloweeIds
	}
	return nil
}

type GetFollowshipCountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFollowshipCountsRequest) Reset() {
	*x = GetFollowshipCountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowshipCountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowshipCountsRequest) ProtoMessage() {}

func (x *GetFollowshipCountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowshipCountsRequest.ProtoReflect.Descriptor instead.
func (*GetFollowshipCountsRequest) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{5}
}

func (x *GetFollowshipCountsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFollowshipCountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerCount  int64 `protobuf:"varint,1,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	FollowingCount int64 `protobuf:"varint,2,opt,name=following_count,json=followingCount,proto3" json:"following_count,omitempty"`
}

func (x *GetFollowshipCountsResponse) Reset() {
	*x = GetFollowshipCountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowshipCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowshipCountsResponse) ProtoMessage() {}

func (x *GetFollowshipCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowshipCountsResponse.ProtoReflect.Descriptor instead.
func (*GetFollowshipCountsResponse) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{6}
}

func (x *GetFollowshipCountsResponse) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *GetFollowshipCountsResponse) GetFollowingCount() int64 {
	if x != nil {
		return x.FollowingCount
	}
	return 0
}

type DoesFollowshipExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId int32 `protobuf:"varint,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	FolloweeId int32 `protobuf:"varint,2,opt,name=followee_id,json=followeeId,proto3" json:"followee_id,omitempty"`
}

func (x *DoesFollowshipExistRequest) Reset() {
	*x = DoesFollowshipExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoesFollowshipExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoesFollowshipExistRequest) ProtoMessage() {}

func (x *DoesFollowshipExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoesFollowshipExistRequest.ProtoReflect.Descriptor instead.
func (*DoesFollowshipExistRequest) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{7}
}

func (x *DoesFollowshipExistRequest) GetFollowerId() int32 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

func (x *DoesFollowshipExistRequest) GetFolloweeId() int32 {
	if x != nil {
		return x.FolloweeId
	}
	return 0
}

type DoesFollowshipExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *DoesFollowshipExistResponse) Reset() {
	*x = DoesFollowshipExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_protos_followships_microservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoesFollowshipExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoesFollowshipExistResponse) ProtoMessage() {}

func (x *DoesFollowshipExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_protos_followships_microservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoesFollowshipExistResponse.ProtoReflect.Descriptor instead.
func (*DoesFollowshipExistResponse) Descriptor() ([]byte, []int) {
	return file_backend_protos_followships_microservice_proto_rawDescGZIP(), []int{8}
}

func (x *DoesFollowshipExistResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_backend_protos_followships_microservice_proto protoreflect.FileDescriptor

var file_backend_protos_followships_microservice_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x68, 0x69, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x39, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x64, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3a, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x6d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5e,
	0x0a, 0x1a, 0x44, 0x6f, 0x65, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x22, 0x35,
	0x0a, 0x1b, 0x44, 0x6f, 0x65, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xe9, 0x05, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x34,
	0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x68, 0x69, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x08,
	0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x34, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x82, 0x01, 0x0a, 0x13, 0x44, 0x6f, 0x65, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x34,
	0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x65, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x6f, 0x65, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x34, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73,
	0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_protos_followships_microservice_proto_rawDescOnce sync.Once
	file_backend_protos_followships_microservice_proto_rawDescData = file_backend_protos_followships_microservice_proto_rawDesc
)

func file_backend_protos_followships_microservice_proto_rawDescGZIP() []byte {
	file_backend_protos_followships_microservice_proto_rawDescOnce.Do(func() {
		file_backend_protos_followships_microservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_protos_followships_microservice_proto_rawDescData)
	})
	return file_backend_protos_followships_microservice_proto_rawDescData
}

var file_backend_protos_followships_microservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_backend_protos_followships_microservice_proto_goTypes = []interface{}{
	(*FollowshipOperationRequest)(nil),  // 0: followships_microservice.FollowshipOperationRequest
	(*GetFollowersRequest)(nil),         // 1: followships_microservice.GetFollowersRequest
	(*GetFollowersResponse)(nil),        // 2: followships_microservice.GetFollowersResponse
	(*GetFollowingsRequest)(nil),        // 3: followships_microservice.GetFollowingsRequest
	(*GetFollowingsResponse)(nil),       // 4: followships_microservice.GetFollowingsResponse
	(*GetFollowshipCountsRequest)(nil),  // 5: followships_microservice.GetFollowshipCountsRequest
	(*GetFollowshipCountsResponse)(nil), // 6: followships_microservice.GetFollowshipCountsResponse
	(*DoesFollowshipExistRequest)(nil),  // 7: followships_microservice.DoesFollowshipExistRequest
	(*DoesFollowshipExistResponse)(nil), // 8: followships_microservice.DoesFollowshipExistResponse
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_backend_protos_followships_microservice_proto_depIdxs = []int32{
	9, // 0: followships_microservice.FollowshipsService.Ping:input_type -> google.protobuf.Empty
	0, // 1: followships_microservice.FollowshipsService.Follow:input_type -> followships_microservice.FollowshipOperationRequest
	0, // 2: followships_microservice.FollowshipsService.Unfollow:input_type -> followships_microservice.FollowshipOperationRequest
	7, // 3: followships_microservice.FollowshipsService.DoesFollowshipExist:input_type -> followships_microservice.DoesFollowshipExistRequest
	1, // 4: followships_microservice.FollowshipsService.GetFollowers:input_type -> followships_microservice.GetFollowersRequest
	3, // 5: followships_microservice.FollowshipsService.GetFollowings:input_type -> followships_microservice.GetFollowingsRequest
	5, // 6: followships_microservice.FollowshipsService.GetFollowshipCounts:input_type -> followships_microservice.GetFollowshipCountsRequest
	9, // 7: followships_microservice.FollowshipsService.Ping:output_type -> google.protobuf.Empty
	9, // 8: followships_microservice.FollowshipsService.Follow:output_type -> google.protobuf.Empty
	9, // 9: followships_microservice.FollowshipsService.Unfollow:output_type -> google.protobuf.Empty
	8, // 10: followships_microservice.FollowshipsService.DoesFollowshipExist:output_type -> followships_microservice.DoesFollowshipExistResponse
	2, // 11: followships_microservice.FollowshipsService.GetFollowers:output_type -> followships_microservice.GetFollowersResponse
	4, // 12: followships_microservice.FollowshipsService.GetFollowings:output_type -> followships_microservice.GetFollowingsResponse
	6, // 13: followships_microservice.FollowshipsService.GetFollowshipCounts:output_type -> followships_microservice.GetFollowshipCountsResponse
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_protos_followships_microservice_proto_init() }
func file_backend_protos_followships_microservice_proto_init() {
	if File_backend_protos_followships_microservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_protos_followships_microservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowshipOperationRequest); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersRequest); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersResponse); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingsRequest); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingsResponse); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowshipCountsRequest); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowshipCountsResponse); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoesFollowshipExistRequest); i {
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
		file_backend_protos_followships_microservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoesFollowshipExistResponse); i {
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
			RawDescriptor: file_backend_protos_followships_microservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_protos_followships_microservice_proto_goTypes,
		DependencyIndexes: file_backend_protos_followships_microservice_proto_depIdxs,
		MessageInfos:      file_backend_protos_followships_microservice_proto_msgTypes,
	}.Build()
	File_backend_protos_followships_microservice_proto = out.File
	file_backend_protos_followships_microservice_proto_rawDesc = nil
	file_backend_protos_followships_microservice_proto_goTypes = nil
	file_backend_protos_followships_microservice_proto_depIdxs = nil
}