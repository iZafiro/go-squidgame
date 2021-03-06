// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/leaderpb.proto

package leaderpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetPlayerStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int32 `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *GetPlayerStateRequest) Reset() {
	*x = GetPlayerStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerStateRequest) ProtoMessage() {}

func (x *GetPlayerStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerStateRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerStateRequest) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerStateRequest) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

type GetPlayerStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage      int32 `protobuf:"varint,1,opt,name=stage,proto3" json:"stage,omitempty"`
	Row        int32 `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	HasStarted bool  `protobuf:"varint,3,opt,name=hasStarted,proto3" json:"hasStarted,omitempty"`
	HasMoved   bool  `protobuf:"varint,4,opt,name=hasMoved,proto3" json:"hasMoved,omitempty"`
	HasLost    bool  `protobuf:"varint,5,opt,name=hasLost,proto3" json:"hasLost,omitempty"`
}

func (x *GetPlayerStateResponse) Reset() {
	*x = GetPlayerStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerStateResponse) ProtoMessage() {}

func (x *GetPlayerStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerStateResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerStateResponse) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerStateResponse) GetStage() int32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *GetPlayerStateResponse) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *GetPlayerStateResponse) GetHasStarted() bool {
	if x != nil {
		return x.HasStarted
	}
	return false
}

func (x *GetPlayerStateResponse) GetHasMoved() bool {
	if x != nil {
		return x.HasMoved
	}
	return false
}

func (x *GetPlayerStateResponse) GetHasLost() bool {
	if x != nil {
		return x.HasLost
	}
	return false
}

type SendPlayerMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int32 `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Move     int32 `protobuf:"varint,2,opt,name=move,proto3" json:"move,omitempty"`
}

func (x *SendPlayerMoveRequest) Reset() {
	*x = SendPlayerMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPlayerMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPlayerMoveRequest) ProtoMessage() {}

func (x *SendPlayerMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPlayerMoveRequest.ProtoReflect.Descriptor instead.
func (*SendPlayerMoveRequest) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{2}
}

func (x *SendPlayerMoveRequest) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *SendPlayerMoveRequest) GetMove() int32 {
	if x != nil {
		return x.Move
	}
	return 0
}

type SendPlayerMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SendPlayerMoveResponse) Reset() {
	*x = SendPlayerMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPlayerMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPlayerMoveResponse) ProtoMessage() {}

func (x *SendPlayerMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPlayerMoveResponse.ProtoReflect.Descriptor instead.
func (*SendPlayerMoveResponse) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{3}
}

func (x *SendPlayerMoveResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PlayerGetPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request int32 `protobuf:"varint,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *PlayerGetPoolRequest) Reset() {
	*x = PlayerGetPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerGetPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerGetPoolRequest) ProtoMessage() {}

func (x *PlayerGetPoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerGetPoolRequest.ProtoReflect.Descriptor instead.
func (*PlayerGetPoolRequest) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerGetPoolRequest) GetRequest() int32 {
	if x != nil {
		return x.Request
	}
	return 0
}

type PlayerGetPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool int32 `protobuf:"varint,1,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (x *PlayerGetPoolResponse) Reset() {
	*x = PlayerGetPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_leaderpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerGetPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerGetPoolResponse) ProtoMessage() {}

func (x *PlayerGetPoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_leaderpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerGetPoolResponse.ProtoReflect.Descriptor instead.
func (*PlayerGetPoolResponse) Descriptor() ([]byte, []int) {
	return file_api_leaderpb_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerGetPoolResponse) GetPool() int32 {
	if x != nil {
		return x.Pool
	}
	return 0
}

var File_api_leaderpb_proto protoreflect.FileDescriptor

var file_api_leaderpb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x22, 0x33,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x76,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x15,
	0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x15, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x32, 0x91, 0x02, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76,
	0x65, 0x12, 0x1f, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1e, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_leaderpb_proto_rawDescOnce sync.Once
	file_api_leaderpb_proto_rawDescData = file_api_leaderpb_proto_rawDesc
)

func file_api_leaderpb_proto_rawDescGZIP() []byte {
	file_api_leaderpb_proto_rawDescOnce.Do(func() {
		file_api_leaderpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_leaderpb_proto_rawDescData)
	})
	return file_api_leaderpb_proto_rawDescData
}

var file_api_leaderpb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_leaderpb_proto_goTypes = []interface{}{
	(*GetPlayerStateRequest)(nil),  // 0: leaderpb.GetPlayerStateRequest
	(*GetPlayerStateResponse)(nil), // 1: leaderpb.GetPlayerStateResponse
	(*SendPlayerMoveRequest)(nil),  // 2: leaderpb.SendPlayerMoveRequest
	(*SendPlayerMoveResponse)(nil), // 3: leaderpb.SendPlayerMoveResponse
	(*PlayerGetPoolRequest)(nil),   // 4: leaderpb.PlayerGetPoolRequest
	(*PlayerGetPoolResponse)(nil),  // 5: leaderpb.PlayerGetPoolResponse
}
var file_api_leaderpb_proto_depIdxs = []int32{
	0, // 0: leaderpb.LeaderService.GetPlayerState:input_type -> leaderpb.GetPlayerStateRequest
	2, // 1: leaderpb.LeaderService.SendPlayerMove:input_type -> leaderpb.SendPlayerMoveRequest
	4, // 2: leaderpb.LeaderService.PlayerGetPool:input_type -> leaderpb.PlayerGetPoolRequest
	1, // 3: leaderpb.LeaderService.GetPlayerState:output_type -> leaderpb.GetPlayerStateResponse
	3, // 4: leaderpb.LeaderService.SendPlayerMove:output_type -> leaderpb.SendPlayerMoveResponse
	5, // 5: leaderpb.LeaderService.PlayerGetPool:output_type -> leaderpb.PlayerGetPoolResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_leaderpb_proto_init() }
func file_api_leaderpb_proto_init() {
	if File_api_leaderpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_leaderpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerStateRequest); i {
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
		file_api_leaderpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerStateResponse); i {
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
		file_api_leaderpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPlayerMoveRequest); i {
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
		file_api_leaderpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPlayerMoveResponse); i {
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
		file_api_leaderpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerGetPoolRequest); i {
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
		file_api_leaderpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerGetPoolResponse); i {
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
			RawDescriptor: file_api_leaderpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_leaderpb_proto_goTypes,
		DependencyIndexes: file_api_leaderpb_proto_depIdxs,
		MessageInfos:      file_api_leaderpb_proto_msgTypes,
	}.Build()
	File_api_leaderpb_proto = out.File
	file_api_leaderpb_proto_rawDesc = nil
	file_api_leaderpb_proto_goTypes = nil
	file_api_leaderpb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LeaderServiceClient is the client API for LeaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeaderServiceClient interface {
	GetPlayerState(ctx context.Context, in *GetPlayerStateRequest, opts ...grpc.CallOption) (*GetPlayerStateResponse, error)
	SendPlayerMove(ctx context.Context, in *SendPlayerMoveRequest, opts ...grpc.CallOption) (*SendPlayerMoveResponse, error)
	PlayerGetPool(ctx context.Context, in *PlayerGetPoolRequest, opts ...grpc.CallOption) (*PlayerGetPoolResponse, error)
}

type leaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaderServiceClient(cc grpc.ClientConnInterface) LeaderServiceClient {
	return &leaderServiceClient{cc}
}

func (c *leaderServiceClient) GetPlayerState(ctx context.Context, in *GetPlayerStateRequest, opts ...grpc.CallOption) (*GetPlayerStateResponse, error) {
	out := new(GetPlayerStateResponse)
	err := c.cc.Invoke(ctx, "/leaderpb.LeaderService/GetPlayerState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderServiceClient) SendPlayerMove(ctx context.Context, in *SendPlayerMoveRequest, opts ...grpc.CallOption) (*SendPlayerMoveResponse, error) {
	out := new(SendPlayerMoveResponse)
	err := c.cc.Invoke(ctx, "/leaderpb.LeaderService/SendPlayerMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderServiceClient) PlayerGetPool(ctx context.Context, in *PlayerGetPoolRequest, opts ...grpc.CallOption) (*PlayerGetPoolResponse, error) {
	out := new(PlayerGetPoolResponse)
	err := c.cc.Invoke(ctx, "/leaderpb.LeaderService/PlayerGetPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaderServiceServer is the server API for LeaderService service.
type LeaderServiceServer interface {
	GetPlayerState(context.Context, *GetPlayerStateRequest) (*GetPlayerStateResponse, error)
	SendPlayerMove(context.Context, *SendPlayerMoveRequest) (*SendPlayerMoveResponse, error)
	PlayerGetPool(context.Context, *PlayerGetPoolRequest) (*PlayerGetPoolResponse, error)
}

// UnimplementedLeaderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLeaderServiceServer struct {
}

func (*UnimplementedLeaderServiceServer) GetPlayerState(context.Context, *GetPlayerStateRequest) (*GetPlayerStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerState not implemented")
}
func (*UnimplementedLeaderServiceServer) SendPlayerMove(context.Context, *SendPlayerMoveRequest) (*SendPlayerMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPlayerMove not implemented")
}
func (*UnimplementedLeaderServiceServer) PlayerGetPool(context.Context, *PlayerGetPoolRequest) (*PlayerGetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerGetPool not implemented")
}

func RegisterLeaderServiceServer(s *grpc.Server, srv LeaderServiceServer) {
	s.RegisterService(&_LeaderService_serviceDesc, srv)
}

func _LeaderService_GetPlayerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderServiceServer).GetPlayerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderpb.LeaderService/GetPlayerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderServiceServer).GetPlayerState(ctx, req.(*GetPlayerStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderService_SendPlayerMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPlayerMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderServiceServer).SendPlayerMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderpb.LeaderService/SendPlayerMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderServiceServer).SendPlayerMove(ctx, req.(*SendPlayerMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderService_PlayerGetPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderServiceServer).PlayerGetPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderpb.LeaderService/PlayerGetPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderServiceServer).PlayerGetPool(ctx, req.(*PlayerGetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeaderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "leaderpb.LeaderService",
	HandlerType: (*LeaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerState",
			Handler:    _LeaderService_GetPlayerState_Handler,
		},
		{
			MethodName: "SendPlayerMove",
			Handler:    _LeaderService_SendPlayerMove_Handler,
		},
		{
			MethodName: "PlayerGetPool",
			Handler:    _LeaderService_PlayerGetPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/leaderpb.proto",
}
