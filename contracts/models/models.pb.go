// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contracts/wsMessages/models.proto

package models

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ParticipantStatus int32

const (
	ParticipantStatus_Active ParticipantStatus = 0
	ParticipantStatus_Banned ParticipantStatus = 1
)

// Enum value maps for ParticipantStatus.
var (
	ParticipantStatus_name = map[int32]string{
		0: "Active",
		1: "Banned",
	}
	ParticipantStatus_value = map[string]int32{
		"Active": 0,
		"Banned": 1,
	}
)

func (x ParticipantStatus) Enum() *ParticipantStatus {
	p := new(ParticipantStatus)
	*p = x
	return p
}

func (x ParticipantStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParticipantStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_contracts_wsMessages_models_proto_enumTypes[0].Descriptor()
}

func (ParticipantStatus) Type() protoreflect.EnumType {
	return &file_contracts_wsMessages_models_proto_enumTypes[0]
}

func (x ParticipantStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParticipantStatus.Descriptor instead.
func (ParticipantStatus) EnumDescriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{0}
}

type MessageType int32

const (
	MessageType_ParticipantCreated MessageType = 0
	MessageType_System             MessageType = 1
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "ParticipantCreated",
		1: "System",
	}
	MessageType_value = map[string]int32{
		"ParticipantCreated": 0,
		"System":             1,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_contracts_wsMessages_models_proto_enumTypes[1].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_contracts_wsMessages_models_proto_enumTypes[1]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{1}
}

type MessageSubType int32

const (
	MessageSubType_ParticipantJoined MessageSubType = 0
	MessageSubType_ChatCreated       MessageSubType = 1
	MessageSubType_TextMessage       MessageSubType = 2
)

// Enum value maps for MessageSubType.
var (
	MessageSubType_name = map[int32]string{
		0: "ParticipantJoined",
		1: "ChatCreated",
		2: "TextMessage",
	}
	MessageSubType_value = map[string]int32{
		"ParticipantJoined": 0,
		"ChatCreated":       1,
		"TextMessage":       2,
	}
)

func (x MessageSubType) Enum() *MessageSubType {
	p := new(MessageSubType)
	*p = x
	return p
}

func (x MessageSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_contracts_wsMessages_models_proto_enumTypes[2].Descriptor()
}

func (MessageSubType) Type() protoreflect.EnumType {
	return &file_contracts_wsMessages_models_proto_enumTypes[2]
}

func (x MessageSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageSubType.Descriptor instead.
func (MessageSubType) EnumDescriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{2}
}

type ParticipantShortInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string      `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	CustomData *CustomData `protobuf:"bytes,2,opt,name=customData,proto3,oneof" json:"customData,omitempty"`
}

func (x *ParticipantShortInfo) Reset() {
	*x = ParticipantShortInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_wsMessages_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantShortInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantShortInfo) ProtoMessage() {}

func (x *ParticipantShortInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_wsMessages_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantShortInfo.ProtoReflect.Descriptor instead.
func (*ParticipantShortInfo) Descriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{0}
}

func (x *ParticipantShortInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ParticipantShortInfo) GetCustomData() *CustomData {
	if x != nil {
		return x.CustomData
	}
	return nil
}

type ChannelParticipantGrants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendMessage  *bool `protobuf:"varint,1,opt,name=sendMessage,proto3,oneof" json:"sendMessage,omitempty"`
	ReadMessages *bool `protobuf:"varint,2,opt,name=readMessages,proto3,oneof" json:"readMessages,omitempty"`
	Admin        *bool `protobuf:"varint,3,opt,name=admin,proto3,oneof" json:"admin,omitempty"`
}

func (x *ChannelParticipantGrants) Reset() {
	*x = ChannelParticipantGrants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_wsMessages_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelParticipantGrants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelParticipantGrants) ProtoMessage() {}

func (x *ChannelParticipantGrants) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_wsMessages_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelParticipantGrants.ProtoReflect.Descriptor instead.
func (*ChannelParticipantGrants) Descriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelParticipantGrants) GetSendMessage() bool {
	if x != nil && x.SendMessage != nil {
		return *x.SendMessage
	}
	return false
}

func (x *ChannelParticipantGrants) GetReadMessages() bool {
	if x != nil && x.ReadMessages != nil {
		return *x.ReadMessages
	}
	return false
}

func (x *ChannelParticipantGrants) GetAdmin() bool {
	if x != nil && x.Admin != nil {
		return *x.Admin
	}
	return false
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string               `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Status     ParticipantStatus    `protobuf:"varint,3,opt,name=status,proto3,enum=models.ParticipantStatus" json:"status,omitempty"`
	ChannelId  string               `protobuf:"bytes,4,opt,name=channelId,proto3" json:"channelId,omitempty"`
	CustomData *CustomData          `protobuf:"bytes,5,opt,name=customData,proto3,oneof" json:"customData,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_wsMessages_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_wsMessages_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{2}
}

func (x *Participant) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Participant) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Participant) GetStatus() ParticipantStatus {
	if x != nil {
		return x.Status
	}
	return ParticipantStatus_Active
}

func (x *Participant) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Participant) GetCustomData() *CustomData {
	if x != nil {
		return x.CustomData
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender     *ParticipantShortInfo `protobuf:"bytes,2,opt,name=sender,proto3,oneof" json:"sender,omitempty"`
	Text       string                `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Type       MessageType           `protobuf:"varint,4,opt,name=type,proto3,enum=models.MessageType" json:"type,omitempty"`
	SubType    MessageSubType        `protobuf:"varint,5,opt,name=subType,proto3,enum=models.MessageSubType" json:"subType,omitempty"`
	LocalId    string                `protobuf:"bytes,6,opt,name=localId,proto3" json:"localId,omitempty"`
	CustomData *CustomData           `protobuf:"bytes,7,opt,name=customData,proto3,oneof" json:"customData,omitempty"`
	CreatedAt  *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_wsMessages_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_wsMessages_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetSender() *ParticipantShortInfo {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_ParticipantCreated
}

func (x *Message) GetSubType() MessageSubType {
	if x != nil {
		return x.SubType
	}
	return MessageSubType_ParticipantJoined
}

func (x *Message) GetLocalId() string {
	if x != nil {
		return x.LocalId
	}
	return ""
}

func (x *Message) GetCustomData() *CustomData {
	if x != nil {
		return x.CustomData
	}
	return nil
}

func (x *Message) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CustomData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CustomData) Reset() {
	*x = CustomData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_wsMessages_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomData) ProtoMessage() {}

func (x *CustomData) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_wsMessages_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomData.ProtoReflect.Descriptor instead.
func (*CustomData) Descriptor() ([]byte, []int) {
	return file_contracts_wsMessages_models_proto_rawDescGZIP(), []int{4}
}

func (x *CustomData) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_contracts_wsMessages_models_proto protoreflect.FileDescriptor

var file_contracts_wsMessages_models_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x77, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0xb0, 0x01, 0x0a,
	0x18, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x80, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x22, 0xea, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x77, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2b, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x64, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x2a, 0x49, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x10, 0x02, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contracts_wsMessages_models_proto_rawDescOnce sync.Once
	file_contracts_wsMessages_models_proto_rawDescData = file_contracts_wsMessages_models_proto_rawDesc
)

func file_contracts_wsMessages_models_proto_rawDescGZIP() []byte {
	file_contracts_wsMessages_models_proto_rawDescOnce.Do(func() {
		file_contracts_wsMessages_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_contracts_wsMessages_models_proto_rawDescData)
	})
	return file_contracts_wsMessages_models_proto_rawDescData
}

var file_contracts_wsMessages_models_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_contracts_wsMessages_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_contracts_wsMessages_models_proto_goTypes = []interface{}{
	(ParticipantStatus)(0),           // 0: models.ParticipantStatus
	(MessageType)(0),                 // 1: models.MessageType
	(MessageSubType)(0),              // 2: models.MessageSubType
	(*ParticipantShortInfo)(nil),     // 3: models.ParticipantShortInfo
	(*ChannelParticipantGrants)(nil), // 4: models.ChannelParticipantGrants
	(*Participant)(nil),              // 5: models.Participant
	(*Message)(nil),                  // 6: models.Message
	(*CustomData)(nil),               // 7: models.CustomData
	nil,                              // 8: models.CustomData.DataEntry
	(*timestamp.Timestamp)(nil),      // 9: google.protobuf.Timestamp
}
var file_contracts_wsMessages_models_proto_depIdxs = []int32{
	7,  // 0: models.ParticipantShortInfo.customData:type_name -> models.CustomData
	9,  // 1: models.Participant.createdAt:type_name -> google.protobuf.Timestamp
	0,  // 2: models.Participant.status:type_name -> models.ParticipantStatus
	7,  // 3: models.Participant.customData:type_name -> models.CustomData
	3,  // 4: models.Message.sender:type_name -> models.ParticipantShortInfo
	1,  // 5: models.Message.type:type_name -> models.MessageType
	2,  // 6: models.Message.subType:type_name -> models.MessageSubType
	7,  // 7: models.Message.customData:type_name -> models.CustomData
	9,  // 8: models.Message.createdAt:type_name -> google.protobuf.Timestamp
	8,  // 9: models.CustomData.data:type_name -> models.CustomData.DataEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_contracts_wsMessages_models_proto_init() }
func file_contracts_wsMessages_models_proto_init() {
	if File_contracts_wsMessages_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contracts_wsMessages_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantShortInfo); i {
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
		file_contracts_wsMessages_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelParticipantGrants); i {
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
		file_contracts_wsMessages_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
		file_contracts_wsMessages_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_contracts_wsMessages_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomData); i {
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
	file_contracts_wsMessages_models_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_contracts_wsMessages_models_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_contracts_wsMessages_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_contracts_wsMessages_models_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contracts_wsMessages_models_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contracts_wsMessages_models_proto_goTypes,
		DependencyIndexes: file_contracts_wsMessages_models_proto_depIdxs,
		EnumInfos:         file_contracts_wsMessages_models_proto_enumTypes,
		MessageInfos:      file_contracts_wsMessages_models_proto_msgTypes,
	}.Build()
	File_contracts_wsMessages_models_proto = out.File
	file_contracts_wsMessages_models_proto_rawDesc = nil
	file_contracts_wsMessages_models_proto_goTypes = nil
	file_contracts_wsMessages_models_proto_depIdxs = nil
}
