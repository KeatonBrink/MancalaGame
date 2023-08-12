// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mancala.proto

package protos

import (
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

type HandshakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *HandshakeRequest) Reset() {
	*x = HandshakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeRequest) ProtoMessage() {}

func (x *HandshakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeRequest.ProtoReflect.Descriptor instead.
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type HandshakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode              int32  `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage           string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Message                string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	ServerWebSocketAddress string `protobuf:"bytes,4,opt,name=serverWebSocketAddress,proto3" json:"serverWebSocketAddress,omitempty"`
}

func (x *HandshakeResponse) Reset() {
	*x = HandshakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResponse) ProtoMessage() {}

func (x *HandshakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResponse.ProtoReflect.Descriptor instead.
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *HandshakeResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *HandshakeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HandshakeResponse) GetServerWebSocketAddress() string {
	if x != nil {
		return x.ServerWebSocketAddress
	}
	return ""
}

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PitIndex int32  `protobuf:"varint,1,opt,name=pitIndex,proto3" json:"pitIndex,omitempty"`
	UserHash string `protobuf:"bytes,2,opt,name=userHash,proto3" json:"userHash,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{2}
}

func (x *MoveRequest) GetPitIndex() int32 {
	if x != nil {
		return x.PitIndex
	}
	return 0
}

func (x *MoveRequest) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Board   string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{3}
}

func (x *MoveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MoveResponse) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHash string `protobuf:"bytes,1,opt,name=userHash,proto3" json:"userHash,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Board   string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mancala_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mancala_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_mancala_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateResponse) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

var File_mancala_proto protoreflect.FileDescriptor

var file_mancala_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x22, 0x2e, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64,
	0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3e, 0x0a, 0x0c, 0x4d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x22, 0x40, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x32, 0xd3, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e,
	0x63, 0x61, 0x6c, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x47,
	0x61, 0x6d, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x19, 0x2e, 0x6d,
	0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c,
	0x61, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x12,
	0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x61,
	0x74, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x6e, 0x6b, 0x2f, 0x4d, 0x61, 0x6e, 0x63, 0x61, 0x6c, 0x61,
	0x47, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mancala_proto_rawDescOnce sync.Once
	file_mancala_proto_rawDescData = file_mancala_proto_rawDesc
)

func file_mancala_proto_rawDescGZIP() []byte {
	file_mancala_proto_rawDescOnce.Do(func() {
		file_mancala_proto_rawDescData = protoimpl.X.CompressGZIP(file_mancala_proto_rawDescData)
	})
	return file_mancala_proto_rawDescData
}

var file_mancala_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mancala_proto_goTypes = []interface{}{
	(*HandshakeRequest)(nil),  // 0: mancala.HandshakeRequest
	(*HandshakeResponse)(nil), // 1: mancala.HandshakeResponse
	(*MoveRequest)(nil),       // 2: mancala.MoveRequest
	(*MoveResponse)(nil),      // 3: mancala.MoveResponse
	(*UpdateRequest)(nil),     // 4: mancala.UpdateRequest
	(*UpdateResponse)(nil),    // 5: mancala.UpdateResponse
}
var file_mancala_proto_depIdxs = []int32{
	0, // 0: mancala.MancalaService.GameHandshake:input_type -> mancala.HandshakeRequest
	2, // 1: mancala.MancalaService.MakeMove:input_type -> mancala.MoveRequest
	4, // 2: mancala.MancalaService.RequestUpdate:input_type -> mancala.UpdateRequest
	1, // 3: mancala.MancalaService.GameHandshake:output_type -> mancala.HandshakeResponse
	3, // 4: mancala.MancalaService.MakeMove:output_type -> mancala.MoveResponse
	5, // 5: mancala.MancalaService.RequestUpdate:output_type -> mancala.UpdateResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mancala_proto_init() }
func file_mancala_proto_init() {
	if File_mancala_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mancala_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeRequest); i {
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
		file_mancala_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResponse); i {
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
		file_mancala_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveRequest); i {
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
		file_mancala_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveResponse); i {
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
		file_mancala_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_mancala_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_mancala_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mancala_proto_goTypes,
		DependencyIndexes: file_mancala_proto_depIdxs,
		MessageInfos:      file_mancala_proto_msgTypes,
	}.Build()
	File_mancala_proto = out.File
	file_mancala_proto_rawDesc = nil
	file_mancala_proto_goTypes = nil
	file_mancala_proto_depIdxs = nil
}
