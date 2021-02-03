// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: service.spotify/proto/v1/spotify.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Represents a user on the spotify service
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{1}
}

func (x *Artist) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists,proto3" json:"artists,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{2}
}

func (x *Track) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Track) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Track) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

type RedeemCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	DiscordId string `protobuf:"bytes,2,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
}

func (x *RedeemCodeRequest) Reset() {
	*x = RedeemCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedeemCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemCodeRequest) ProtoMessage() {}

func (x *RedeemCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemCodeRequest.ProtoReflect.Descriptor instead.
func (*RedeemCodeRequest) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{3}
}

func (x *RedeemCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RedeemCodeRequest) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

type RedeemCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RedeemCodeResponse) Reset() {
	*x = RedeemCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedeemCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemCodeResponse) ProtoMessage() {}

func (x *RedeemCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemCodeResponse.ProtoReflect.Descriptor instead.
func (*RedeemCodeResponse) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{4}
}

func (x *RedeemCodeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAuthURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordId string `protobuf:"bytes,1,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
}

func (x *GetAuthURLRequest) Reset() {
	*x = GetAuthURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthURLRequest) ProtoMessage() {}

func (x *GetAuthURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthURLRequest.ProtoReflect.Descriptor instead.
func (*GetAuthURLRequest) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthURLRequest) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

type GetAuthURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetAuthURLResponse) Reset() {
	*x = GetAuthURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthURLResponse) ProtoMessage() {}

func (x *GetAuthURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthURLResponse.ProtoReflect.Descriptor instead.
func (*GetAuthURLResponse) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetListeningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordId string `protobuf:"bytes,1,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
}

func (x *GetListeningRequest) Reset() {
	*x = GetListeningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListeningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListeningRequest) ProtoMessage() {}

func (x *GetListeningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListeningRequest.ProtoReflect.Descriptor instead.
func (*GetListeningRequest) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{7}
}

func (x *GetListeningRequest) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

type GetListeningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Track *Track `protobuf:"bytes,1,opt,name=track,proto3" json:"track,omitempty"`
}

func (x *GetListeningResponse) Reset() {
	*x = GetListeningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListeningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListeningResponse) ProtoMessage() {}

func (x *GetListeningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_spotify_proto_v1_spotify_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListeningResponse.ProtoReflect.Descriptor instead.
func (*GetListeningResponse) Descriptor() ([]byte, []int) {
	return file_service_spotify_proto_v1_spotify_proto_rawDescGZIP(), []int{8}
}

func (x *GetListeningResponse) GetTrack() *Track {
	if x != nil {
		return x.Track
	}
	return nil
}

var File_service_spotify_proto_v1_spotify_proto protoreflect.FileDescriptor

var file_service_spotify_proto_v1_spotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70,
	0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x4f, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a,
	0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x75, 0x74, 0x65,
	0x72, 0x70, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x22, 0x46,
	0x0a, 0x11, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x75, 0x74,
	0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x26,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x32, 0xa6, 0x02, 0x0a, 0x07, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x5b, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x25, 0x2e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65,
	0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x65,
	0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x2e, 0x65,
	0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x65, 0x75,
	0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2e, 0x73,
	0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x69, 0x64, 0x65, 0x79, 0x6e, 0x65, 0x74, 0x2f, 0x65, 0x75, 0x74, 0x65, 0x72, 0x70, 0x65, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x76, 0x31, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_spotify_proto_v1_spotify_proto_rawDescOnce sync.Once
	file_service_spotify_proto_v1_spotify_proto_rawDescData = file_service_spotify_proto_v1_spotify_proto_rawDesc
)

func file_service_spotify_proto_v1_spotify_proto_rawDescGZIP() []byte {
	file_service_spotify_proto_v1_spotify_proto_rawDescOnce.Do(func() {
		file_service_spotify_proto_v1_spotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_spotify_proto_v1_spotify_proto_rawDescData)
	})
	return file_service_spotify_proto_v1_spotify_proto_rawDescData
}

var file_service_spotify_proto_v1_spotify_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_spotify_proto_v1_spotify_proto_goTypes = []interface{}{
	(*User)(nil),                 // 0: euterpe.spotify.v1.User
	(*Artist)(nil),               // 1: euterpe.spotify.v1.Artist
	(*Track)(nil),                // 2: euterpe.spotify.v1.Track
	(*RedeemCodeRequest)(nil),    // 3: euterpe.spotify.v1.RedeemCodeRequest
	(*RedeemCodeResponse)(nil),   // 4: euterpe.spotify.v1.RedeemCodeResponse
	(*GetAuthURLRequest)(nil),    // 5: euterpe.spotify.v1.GetAuthURLRequest
	(*GetAuthURLResponse)(nil),   // 6: euterpe.spotify.v1.GetAuthURLResponse
	(*GetListeningRequest)(nil),  // 7: euterpe.spotify.v1.GetListeningRequest
	(*GetListeningResponse)(nil), // 8: euterpe.spotify.v1.GetListeningResponse
}
var file_service_spotify_proto_v1_spotify_proto_depIdxs = []int32{
	1, // 0: euterpe.spotify.v1.Track.artists:type_name -> euterpe.spotify.v1.Artist
	0, // 1: euterpe.spotify.v1.RedeemCodeResponse.user:type_name -> euterpe.spotify.v1.User
	2, // 2: euterpe.spotify.v1.GetListeningResponse.track:type_name -> euterpe.spotify.v1.Track
	3, // 3: euterpe.spotify.v1.Spotify.RedeemCode:input_type -> euterpe.spotify.v1.RedeemCodeRequest
	5, // 4: euterpe.spotify.v1.Spotify.GetAuthURL:input_type -> euterpe.spotify.v1.GetAuthURLRequest
	7, // 5: euterpe.spotify.v1.Spotify.GetListening:input_type -> euterpe.spotify.v1.GetListeningRequest
	4, // 6: euterpe.spotify.v1.Spotify.RedeemCode:output_type -> euterpe.spotify.v1.RedeemCodeResponse
	6, // 7: euterpe.spotify.v1.Spotify.GetAuthURL:output_type -> euterpe.spotify.v1.GetAuthURLResponse
	8, // 8: euterpe.spotify.v1.Spotify.GetListening:output_type -> euterpe.spotify.v1.GetListeningResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_spotify_proto_v1_spotify_proto_init() }
func file_service_spotify_proto_v1_spotify_proto_init() {
	if File_service_spotify_proto_v1_spotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_spotify_proto_v1_spotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedeemCodeRequest); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedeemCodeResponse); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthURLRequest); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthURLResponse); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListeningRequest); i {
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
		file_service_spotify_proto_v1_spotify_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListeningResponse); i {
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
			RawDescriptor: file_service_spotify_proto_v1_spotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_spotify_proto_v1_spotify_proto_goTypes,
		DependencyIndexes: file_service_spotify_proto_v1_spotify_proto_depIdxs,
		MessageInfos:      file_service_spotify_proto_v1_spotify_proto_msgTypes,
	}.Build()
	File_service_spotify_proto_v1_spotify_proto = out.File
	file_service_spotify_proto_v1_spotify_proto_rawDesc = nil
	file_service_spotify_proto_v1_spotify_proto_goTypes = nil
	file_service_spotify_proto_v1_spotify_proto_depIdxs = nil
}
