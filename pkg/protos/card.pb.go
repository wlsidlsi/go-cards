// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pkg/protos/card.proto

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Suit string `protobuf:"bytes,2,opt,name=suit,proto3" json:"suit,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_pkg_protos_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Card) GetSuit() string {
	if x != nil {
		return x.Suit
	}
	return ""
}

var File_pkg_protos_card_proto protoreflect.FileDescriptor

var file_pkg_protos_card_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22,
	0x2a, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x75, 0x69, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_protos_card_proto_rawDescOnce sync.Once
	file_pkg_protos_card_proto_rawDescData = file_pkg_protos_card_proto_rawDesc
)

func file_pkg_protos_card_proto_rawDescGZIP() []byte {
	file_pkg_protos_card_proto_rawDescOnce.Do(func() {
		file_pkg_protos_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_card_proto_rawDescData)
	})
	return file_pkg_protos_card_proto_rawDescData
}

var file_pkg_protos_card_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_protos_card_proto_goTypes = []any{
	(*Card)(nil), // 0: protos.Card
}
var file_pkg_protos_card_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_protos_card_proto_init() }
func file_pkg_protos_card_proto_init() {
	if File_pkg_protos_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_card_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Card); i {
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
			RawDescriptor: file_pkg_protos_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protos_card_proto_goTypes,
		DependencyIndexes: file_pkg_protos_card_proto_depIdxs,
		MessageInfos:      file_pkg_protos_card_proto_msgTypes,
	}.Build()
	File_pkg_protos_card_proto = out.File
	file_pkg_protos_card_proto_rawDesc = nil
	file_pkg_protos_card_proto_goTypes = nil
	file_pkg_protos_card_proto_depIdxs = nil
}
