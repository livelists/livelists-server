// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contracts/channel/channel.proto

package channel

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

type ChannelStatus int32

const (
	ChannelStatus_Active   ChannelStatus = 0
	ChannelStatus_Archived ChannelStatus = 1
)

// Enum value maps for ChannelStatus.
var (
	ChannelStatus_name = map[int32]string{
		0: "Active",
		1: "Archived",
	}
	ChannelStatus_value = map[string]int32{
		"Active":   0,
		"Archived": 1,
	}
)

func (x ChannelStatus) Enum() *ChannelStatus {
	p := new(ChannelStatus)
	*p = x
	return p
}

func (x ChannelStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_contracts_channel_channel_proto_enumTypes[0].Descriptor()
}

func (ChannelStatus) Type() protoreflect.EnumType {
	return &file_contracts_channel_channel_proto_enumTypes[0]
}

func (x ChannelStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelStatus.Descriptor instead.
func (ChannelStatus) EnumDescriptor() ([]byte, []int) {
	return file_contracts_channel_channel_proto_rawDescGZIP(), []int{0}
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identification  string               `protobuf:"bytes,1,opt,name=identification,proto3" json:"identification,omitempty"`
	CreatedAt       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MaxParticipants int64                `protobuf:"varint,3,opt,name=maxParticipants,proto3" json:"maxParticipants,omitempty"`
	Status          ChannelStatus        `protobuf:"varint,4,opt,name=status,proto3,enum=livelists.ChannelStatus" json:"status,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_channel_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_channel_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_contracts_channel_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetIdentification() string {
	if x != nil {
		return x.Identification
	}
	return ""
}

func (x *Channel) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Channel) GetMaxParticipants() int64 {
	if x != nil {
		return x.MaxParticipants
	}
	return 0
}

func (x *Channel) GetStatus() ChannelStatus {
	if x != nil {
		return x.Status
	}
	return ChannelStatus_Active
}

type CreateChannelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identification  string `protobuf:"bytes,1,opt,name=identification,proto3" json:"identification,omitempty"`
	MaxParticipants int64  `protobuf:"varint,3,opt,name=maxParticipants,proto3" json:"maxParticipants,omitempty"`
}

func (x *CreateChannelReq) Reset() {
	*x = CreateChannelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_channel_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelReq) ProtoMessage() {}

func (x *CreateChannelReq) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_channel_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelReq.ProtoReflect.Descriptor instead.
func (*CreateChannelReq) Descriptor() ([]byte, []int) {
	return file_contracts_channel_channel_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChannelReq) GetIdentification() string {
	if x != nil {
		return x.Identification
	}
	return ""
}

func (x *CreateChannelReq) GetMaxParticipants() int64 {
	if x != nil {
		return x.MaxParticipants
	}
	return 0
}

var File_contracts_channel_channel_proto protoreflect.FileDescriptor

var file_contracts_channel_channel_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x69, 0x76, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6d,
	0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61,
	0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x2a, 0x29, 0x0a,
	0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x10, 0x01, 0x32, 0x52, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x13, 0x5a, 0x11,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contracts_channel_channel_proto_rawDescOnce sync.Once
	file_contracts_channel_channel_proto_rawDescData = file_contracts_channel_channel_proto_rawDesc
)

func file_contracts_channel_channel_proto_rawDescGZIP() []byte {
	file_contracts_channel_channel_proto_rawDescOnce.Do(func() {
		file_contracts_channel_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_contracts_channel_channel_proto_rawDescData)
	})
	return file_contracts_channel_channel_proto_rawDescData
}

var file_contracts_channel_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contracts_channel_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contracts_channel_channel_proto_goTypes = []interface{}{
	(ChannelStatus)(0),          // 0: livelists.ChannelStatus
	(*Channel)(nil),             // 1: livelists.Channel
	(*CreateChannelReq)(nil),    // 2: livelists.CreateChannelReq
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_contracts_channel_channel_proto_depIdxs = []int32{
	3, // 0: livelists.Channel.createdAt:type_name -> google.protobuf.Timestamp
	0, // 1: livelists.Channel.status:type_name -> livelists.ChannelStatus
	2, // 2: livelists.ChannelService.CreateChannel:input_type -> livelists.CreateChannelReq
	1, // 3: livelists.ChannelService.CreateChannel:output_type -> livelists.Channel
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contracts_channel_channel_proto_init() }
func file_contracts_channel_channel_proto_init() {
	if File_contracts_channel_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contracts_channel_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_contracts_channel_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelReq); i {
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
			RawDescriptor: file_contracts_channel_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contracts_channel_channel_proto_goTypes,
		DependencyIndexes: file_contracts_channel_channel_proto_depIdxs,
		EnumInfos:         file_contracts_channel_channel_proto_enumTypes,
		MessageInfos:      file_contracts_channel_channel_proto_msgTypes,
	}.Build()
	File_contracts_channel_channel_proto = out.File
	file_contracts_channel_channel_proto_rawDesc = nil
	file_contracts_channel_channel_proto_goTypes = nil
	file_contracts_channel_channel_proto_depIdxs = nil
}
