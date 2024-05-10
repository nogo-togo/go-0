// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: game/game.proto

package game

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

// The request message containing the user's name.
type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type GameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GameReply) Reset() {
	*x = GameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameReply) ProtoMessage() {}

func (x *GameReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameReply.ProtoReflect.Descriptor instead.
func (*GameReply) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_game_game_proto protoreflect.FileDescriptor

var file_game_game_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x21, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x25, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x40, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x16, 0x42, 0x09, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_game_proto_rawDescOnce sync.Once
	file_game_game_proto_rawDescData = file_game_game_proto_rawDesc
)

func file_game_game_proto_rawDescGZIP() []byte {
	file_game_game_proto_rawDescOnce.Do(func() {
		file_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_game_proto_rawDescData)
	})
	return file_game_game_proto_rawDescData
}

var file_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_game_game_proto_goTypes = []interface{}{
	(*GameRequest)(nil), // 0: helloworld.GameRequest
	(*GameReply)(nil),   // 1: helloworld.GameReply
}
var file_game_game_proto_depIdxs = []int32{
	0, // 0: helloworld.Game.Play:input_type -> helloworld.GameRequest
	1, // 1: helloworld.Game.Play:output_type -> helloworld.GameReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_game_proto_init() }
func file_game_game_proto_init() {
	if File_game_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_game_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameReply); i {
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
			RawDescriptor: file_game_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_game_proto_goTypes,
		DependencyIndexes: file_game_game_proto_depIdxs,
		MessageInfos:      file_game_game_proto_msgTypes,
	}.Build()
	File_game_game_proto = out.File
	file_game_game_proto_rawDesc = nil
	file_game_game_proto_goTypes = nil
	file_game_game_proto_depIdxs = nil
}
