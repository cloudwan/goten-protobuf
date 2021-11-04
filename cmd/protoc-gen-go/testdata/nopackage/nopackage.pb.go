// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/nopackage/nopackage.proto

package nopackage

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

type Enum int32

const (
	Enum_ZERO Enum = 0
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ZERO",
	}
	Enum_value = map[string]int32{
		"ZERO": 0,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum(num)
	return nil
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField *string `protobuf:"bytes,1,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	EnumField   *Enum   `protobuf:"varint,2,opt,name=enum_field,json=enumField,enum=Enum,def=0" json:"enum_field,omitempty"`
}

// Default values for Message fields.
const (
	Default_Message_EnumField = Enum_ZERO
)

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_msgTypes[0]
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
	return file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *Message) GetEnumField() Enum {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return Default_Message_EnumField
}

var File_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6e, 0x6f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x3a,
	0x04, 0x5a, 0x45, 0x52, 0x4f, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x2a, 0x10, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00,
}

var (
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescData = file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_goTypes = []interface{}{
	(Enum)(0),       // 0: Enum
	(*Message)(nil), // 1: Message
}
var file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_depIdxs = []int32{
	0, // 0: Message.enum_field:type_name -> Enum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_init() }
func file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_init() {
	if File_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto = out.File
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_nopackage_nopackage_proto_depIdxs = nil
}