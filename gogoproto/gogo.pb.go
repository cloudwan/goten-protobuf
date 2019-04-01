// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goten/contrib/protobuf/gogoproto/gogo.proto

package gogoproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Enumdecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62024,
	Name:          "gogoproto.enumdecl",
	Tag:           "varint,62024,opt,name=enumdecl",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_TypedeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63030,
	Name:          "gogoproto.typedecl_all",
	Tag:           "varint,63030,opt,name=typedecl_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_EnumdeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63031,
	Name:          "gogoproto.enumdecl_all",
	Tag:           "varint,63031,opt,name=enumdecl_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoRegistration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63032,
	Name:          "gogoproto.goproto_registration",
	Tag:           "varint,63032,opt,name=goproto_registration",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_MessagenameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63033,
	Name:          "gogoproto.messagename_all",
	Tag:           "varint,63033,opt,name=messagename_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoSizecacheAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63034,
	Name:          "gogoproto.goproto_sizecache_all",
	Tag:           "varint,63034,opt,name=goproto_sizecache_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoUnkeyedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63035,
	Name:          "gogoproto.goproto_unkeyed_all",
	Tag:           "varint,63035,opt,name=goproto_unkeyed_all",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Typedecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64030,
	Name:          "gogoproto.typedecl",
	Tag:           "varint,64030,opt,name=typedecl",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Messagename = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64033,
	Name:          "gogoproto.messagename",
	Tag:           "varint,64033,opt,name=messagename",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoSizecache = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64034,
	Name:          "gogoproto.goproto_sizecache",
	Tag:           "varint,64034,opt,name=goproto_sizecache",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_GoprotoUnkeyed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64035,
	Name:          "gogoproto.goproto_unkeyed",
	Tag:           "varint,64035,opt,name=goproto_unkeyed",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Stdtime = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65010,
	Name:          "gogoproto.stdtime",
	Tag:           "varint,65010,opt,name=stdtime",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

var E_Stdduration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65011,
	Name:          "gogoproto.stdduration",
	Tag:           "varint,65011,opt,name=stdduration",
	Filename:      "goten/contrib/protobuf/gogoproto/gogo.proto",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_Enumdecl)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_TypedeclAll)
	proto.RegisterExtension(E_EnumdeclAll)
	proto.RegisterExtension(E_GoprotoRegistration)
	proto.RegisterExtension(E_MessagenameAll)
	proto.RegisterExtension(E_GoprotoSizecacheAll)
	proto.RegisterExtension(E_GoprotoUnkeyedAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Typedecl)
	proto.RegisterExtension(E_Messagename)
	proto.RegisterExtension(E_GoprotoSizecache)
	proto.RegisterExtension(E_GoprotoUnkeyed)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
	proto.RegisterExtension(E_Stdtime)
	proto.RegisterExtension(E_Stdduration)
}

func init() {
	proto.RegisterFile("goten/contrib/protobuf/gogoproto/gogo.proto", fileDescriptor_c9ba883560c81e19)
}

var fileDescriptor_c9ba883560c81e19 = []byte{
	// 1337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0x5b, 0x6f, 0xdc, 0x44,
	0x14, 0x80, 0x85, 0x68, 0xd5, 0xec, 0x24, 0x69, 0x9a, 0x4d, 0x08, 0xa5, 0x02, 0x51, 0xde, 0x90,
	0x90, 0xb2, 0x4f, 0x15, 0xea, 0x44, 0x55, 0x94, 0x46, 0x69, 0x54, 0x44, 0x4a, 0x48, 0x9b, 0x72,
	0x93, 0x58, 0x79, 0xbd, 0x13, 0xc7, 0xd4, 0xf6, 0x6c, 0xed, 0x71, 0x69, 0xfa, 0x86, 0xca, 0x45,
	0x08, 0x71, 0x47, 0x82, 0x96, 0xb6, 0x50, 0x10, 0xf7, 0x6b, 0xb9, 0x5f, 0x5e, 0x78, 0xe1, 0xf2,
	0xca, 0x7f, 0xe0, 0x05, 0x08, 0xf7, 0xbc, 0xe5, 0x05, 0x1d, 0xfb, 0x1c, 0xef, 0xd8, 0x59, 0x69,
	0x66, 0xdf, 0xbc, 0x9b, 0xf9, 0xbe, 0x1d, 0x9f, 0xe3, 0x39, 0xe7, 0xc4, 0xec, 0x16, 0x4f, 0x2a,
	0x11, 0x35, 0x5c, 0x19, 0xa9, 0xd8, 0x6f, 0x35, 0x3a, 0xb1, 0x54, 0xb2, 0x95, 0xae, 0x34, 0x3c,
	0xe9, 0xc9, 0xec, 0x43, 0x76, 0x35, 0x99, 0x5d, 0xd6, 0x6b, 0xc5, 0xb7, 0x7b, 0xf6, 0x7a, 0x52,
	0x7a, 0x81, 0xe8, 0x02, 0x6d, 0x91, 0xb8, 0xb1, 0xdf, 0x51, 0x32, 0xce, 0x17, 0xf3, 0x23, 0x6c,
	0x0c, 0x17, 0x37, 0x45, 0x94, 0x86, 0xcd, 0x4e, 0x2c, 0x56, 0xfc, 0xd3, 0xf5, 0xeb, 0x27, 0x73,
	0x72, 0x92, 0xc8, 0xc9, 0xb9, 0x28, 0x0d, 0xef, 0xe8, 0x28, 0x5f, 0x46, 0xc9, 0xee, 0x2b, 0xbf,
	0x5c, 0xbd, 0xf7, 0xaa, 0x9b, 0x07, 0x96, 0x46, 0x11, 0x85, 0xbf, 0x2d, 0x66, 0x20, 0x5f, 0x62,
	0xd7, 0x94, 0x7c, 0x89, 0x8a, 0xfd, 0xc8, 0x13, 0xb1, 0xc1, 0xf8, 0x3d, 0x1a, 0xc7, 0x34, 0xe3,
	0x51, 0x44, 0xf9, 0x2c, 0x1b, 0xee, 0xc7, 0xf5, 0x03, 0xba, 0x86, 0x84, 0x2e, 0x99, 0x67, 0x23,
	0x99, 0xc4, 0x4d, 0x13, 0x25, 0xc3, 0xc8, 0x09, 0x85, 0x41, 0xf3, 0x63, 0xa6, 0xa9, 0x2d, 0xed,
	0x04, 0x6c, 0xb6, 0xa0, 0x38, 0x67, 0x03, 0xf0, 0x4d, 0x5b, 0xb8, 0x81, 0xc1, 0xf0, 0x13, 0x6e,
	0xa4, 0x58, 0xcf, 0x8f, 0xb3, 0x71, 0xb8, 0x3e, 0xe5, 0x04, 0xa9, 0xd0, 0x77, 0x72, 0x53, 0x4f,
	0xcf, 0x71, 0x58, 0x46, 0xb2, 0x9f, 0xcf, 0x6e, 0xcb, 0xb6, 0x33, 0x56, 0x08, 0xb4, 0x3d, 0x69,
	0x59, 0xf4, 0x84, 0x52, 0x22, 0x4e, 0x9a, 0x4e, 0xd0, 0x6b, 0x7b, 0x87, 0xfc, 0xa0, 0x30, 0x9e,
	0x5b, 0x2f, 0x67, 0x71, 0x3e, 0x27, 0x67, 0x82, 0x80, 0x2f, 0xb3, 0x6b, 0x7b, 0x3c, 0x15, 0x16,
	0xce, 0xf3, 0xe8, 0x1c, 0xdf, 0xf2, 0x64, 0x80, 0x76, 0x91, 0xd1, 0xf7, 0x45, 0x2e, 0x2d, 0x9c,
	0x2f, 0xa3, 0xb3, 0x8e, 0x2c, 0xa5, 0x14, 0x8c, 0xb7, 0xb1, 0xd1, 0x53, 0x22, 0x6e, 0xc9, 0x44,
	0x34, 0xc5, 0xc9, 0xd4, 0x09, 0x2c, 0x74, 0x17, 0x50, 0x37, 0x82, 0xe0, 0x1c, 0x70, 0xe0, 0xda,
	0xcf, 0x06, 0x56, 0x1c, 0x57, 0x58, 0x28, 0x2e, 0xa2, 0x62, 0x07, 0xac, 0x07, 0x74, 0x86, 0x0d,
	0x79, 0x32, 0xbf, 0x25, 0x0b, 0xfc, 0x12, 0xe2, 0x83, 0xc4, 0xa0, 0xa2, 0x23, 0x3b, 0x69, 0xe0,
	0x28, 0x9b, 0x1d, 0xbc, 0x42, 0x0a, 0x62, 0x50, 0xd1, 0x47, 0x58, 0x5f, 0x25, 0x45, 0xa2, 0xc5,
	0x73, 0x9a, 0x0d, 0xca, 0x28, 0x58, 0x93, 0x91, 0xcd, 0x26, 0x2e, 0xa3, 0x81, 0x21, 0x02, 0x82,
	0x29, 0x56, 0xb3, 0x4d, 0xc4, 0x1b, 0xeb, 0x74, 0x3c, 0x28, 0x03, 0xf3, 0x6c, 0x84, 0x0a, 0x94,
	0x2f, 0x23, 0x0b, 0xc5, 0x9b, 0xa8, 0xd8, 0xa9, 0x61, 0x78, 0x1b, 0x4a, 0x24, 0xca, 0x13, 0x36,
	0x92, 0xb7, 0xe8, 0x36, 0x10, 0xc1, 0x50, 0xb6, 0x44, 0xe4, 0xae, 0xda, 0x19, 0xde, 0xa6, 0x50,
	0x12, 0x03, 0x8a, 0x59, 0x36, 0x1c, 0x3a, 0x71, 0xb2, 0xea, 0x04, 0x56, 0xe9, 0x78, 0x07, 0x1d,
	0x43, 0x05, 0x84, 0x11, 0x49, 0xa3, 0x7e, 0x34, 0xef, 0x52, 0x44, 0x34, 0x0c, 0x8f, 0x5e, 0xa2,
	0x9c, 0x56, 0x20, 0x9a, 0xfd, 0xd8, 0xde, 0xa3, 0xa3, 0x97, 0xb3, 0x0b, 0xba, 0x71, 0x8a, 0xd5,
	0x12, 0xff, 0x8c, 0x95, 0xe6, 0x7d, 0xca, 0x74, 0x06, 0x00, 0x7c, 0x0f, 0xbb, 0xae, 0x67, 0x9b,
	0xb0, 0x90, 0x7d, 0x80, 0xb2, 0x89, 0x1e, 0xad, 0x02, 0x4b, 0x42, 0xbf, 0xca, 0x0f, 0xa9, 0x24,
	0x88, 0x8a, 0x6b, 0x91, 0x8d, 0xa7, 0x51, 0xe2, 0xac, 0xf4, 0x17, 0xb5, 0x8f, 0x28, 0x6a, 0x39,
	0x5b, 0x8a, 0xda, 0x31, 0x36, 0x81, 0xc6, 0xfe, 0xf2, 0xfa, 0x31, 0x15, 0xd6, 0x9c, 0x5e, 0x2e,
	0x67, 0xf7, 0x3e, 0xb6, 0xa7, 0x08, 0xe7, 0x69, 0x25, 0xa2, 0x04, 0x98, 0x66, 0xe8, 0x74, 0x2c,
	0xcc, 0x57, 0xd0, 0x4c, 0x15, 0x7f, 0xae, 0x10, 0x2c, 0x38, 0x1d, 0x90, 0xdf, 0xcd, 0x76, 0x93,
	0x3c, 0x8d, 0x62, 0xe1, 0x4a, 0x2f, 0xf2, 0xcf, 0x88, 0xb6, 0x85, 0xfa, 0x93, 0x4a, 0xaa, 0x96,
	0x35, 0x1c, 0xcc, 0x87, 0xd9, 0xae, 0x62, 0x56, 0x69, 0xfa, 0x61, 0x47, 0xc6, 0xca, 0x60, 0xfc,
	0x94, 0x32, 0x55, 0x70, 0x87, 0x33, 0x8c, 0xcf, 0xb1, 0x9d, 0xd9, 0x47, 0xdb, 0x47, 0xf2, 0x33,
	0x14, 0x0d, 0x77, 0x29, 0x2c, 0x1c, 0xae, 0x0c, 0x3b, 0x4e, 0x6c, 0x53, 0xff, 0x3e, 0xa7, 0xc2,
	0x81, 0x08, 0x16, 0x0e, 0xb5, 0xd6, 0x11, 0xd0, 0xed, 0x2d, 0x0c, 0x5f, 0x50, 0xe1, 0x20, 0x06,
	0x15, 0x34, 0x30, 0x58, 0x28, 0xbe, 0x24, 0x05, 0x31, 0xa0, 0xb8, 0xb3, 0xdb, 0x68, 0x63, 0xe1,
	0xf9, 0x89, 0x8a, 0x1d, 0x58, 0x6d, 0x50, 0x7d, 0xb5, 0x5e, 0x1e, 0xc2, 0x96, 0x34, 0x14, 0x2a,
	0x51, 0x28, 0x92, 0xc4, 0xf1, 0x04, 0x4c, 0x1c, 0x16, 0x1b, 0xfb, 0x9a, 0x2a, 0x91, 0x86, 0xc1,
	0xde, 0xb4, 0x09, 0x11, 0xc2, 0xee, 0x3a, 0xee, 0xaa, 0x8d, 0xee, 0x9b, 0xca, 0xe6, 0x8e, 0x12,
	0x0b, 0x4e, 0x6d, 0xfe, 0x49, 0xa3, 0x13, 0x62, 0xcd, 0xea, 0xe9, 0xfc, 0xb6, 0x32, 0xff, 0x2c,
	0xe7, 0x64, 0x5e, 0x43, 0x46, 0x2a, 0xf3, 0x54, 0xfd, 0xc6, 0x2d, 0xae, 0x85, 0xfc, 0xbe, 0x48,
	0xf7, 0xd0, 0x06, 0xde, 0x6f, 0x79, 0x9c, 0xe2, 0xb7, 0xc3, 0x43, 0x5e, 0x1e, 0x7a, 0xcc, 0xb2,
	0xb3, 0x1b, 0xc5, 0x73, 0x5e, 0x9a, 0x79, 0xf8, 0x21, 0x36, 0x5c, 0x1a, 0x78, 0xcc, 0xaa, 0x87,
	0x51, 0x35, 0xa4, 0xcf, 0x3b, 0x7c, 0x1f, 0xdb, 0x06, 0xc3, 0x8b, 0x19, 0x7f, 0x04, 0xf1, 0x6c,
	0x39, 0x3f, 0xc0, 0x06, 0x68, 0x68, 0x31, 0xa3, 0x8f, 0x22, 0x5a, 0x20, 0x80, 0xd3, 0xc0, 0x62,
	0xc6, 0x1f, 0x23, 0x9c, 0x10, 0xc0, 0xed, 0x43, 0xf8, 0xdd, 0x13, 0xdb, 0xb0, 0xe9, 0x50, 0xec,
	0xa6, 0xd8, 0x0e, 0x9c, 0x54, 0xcc, 0xf4, 0xe3, 0xf8, 0xe3, 0x44, 0xf0, 0x5b, 0xd9, 0x76, 0xcb,
	0x80, 0x3f, 0x89, 0x68, 0xbe, 0x9e, 0xcf, 0xb2, 0x41, 0x6d, 0x3a, 0x31, 0xe3, 0x4f, 0x21, 0xae,
	0x53, 0xb0, 0x75, 0x9c, 0x4e, 0xcc, 0x82, 0xa7, 0x69, 0xeb, 0x48, 0x40, 0xd8, 0x68, 0x30, 0x31,
	0xd3, 0xcf, 0x50, 0xd4, 0x09, 0xe1, 0xd3, 0xac, 0x56, 0x34, 0x1b, 0x33, 0xff, 0x2c, 0xf2, 0x5d,
	0x06, 0x22, 0xa0, 0x35, 0x3b, 0xb3, 0xe2, 0x39, 0x8a, 0x80, 0x46, 0xc1, 0x31, 0xaa, 0x0e, 0x30,
	0x66, 0xd3, 0xf3, 0x74, 0x8c, 0x2a, 0xf3, 0x0b, 0x64, 0x33, 0xab, 0xf9, 0x66, 0xc5, 0x0b, 0x94,
	0xcd, 0x6c, 0x3d, 0x6c, 0xa3, 0x3a, 0x11, 0x98, 0x1d, 0x2f, 0xd2, 0x36, 0x2a, 0x03, 0x01, 0x5f,
	0x64, 0xf5, 0xad, 0xd3, 0x80, 0xd9, 0xf7, 0x12, 0xfa, 0x46, 0xb7, 0x0c, 0x03, 0xfc, 0x2e, 0x36,
	0xd1, 0x7b, 0x12, 0x30, 0x5b, 0xcf, 0x6d, 0x54, 0xfe, 0x77, 0xd3, 0x07, 0x01, 0x7e, 0xac, 0xdb,
	0x52, 0xf4, 0x29, 0xc0, 0xac, 0x3d, 0xbf, 0x51, 0x2e, 0xdc, 0xfa, 0x10, 0xc0, 0x67, 0x18, 0xeb,
	0x36, 0x60, 0xb3, 0xeb, 0x02, 0xba, 0x34, 0x08, 0x8e, 0x06, 0xf6, 0x5f, 0x33, 0x7f, 0x91, 0x8e,
	0x06, 0x12, 0x70, 0x34, 0xa8, 0xf5, 0x9a, 0xe9, 0x4b, 0x74, 0x34, 0x08, 0x81, 0x27, 0x5b, 0xeb,
	0x6e, 0x66, 0xc3, 0x65, 0x7a, 0xb2, 0x35, 0x8a, 0x1f, 0x61, 0xa3, 0x5b, 0x1a, 0xa2, 0x59, 0xf5,
	0x1a, 0xaa, 0x76, 0x55, 0xfb, 0xa1, 0xde, 0xbc, 0xb0, 0x19, 0x9a, 0x6d, 0xaf, 0x57, 0x9a, 0x17,
	0xf6, 0x42, 0x3e, 0xc5, 0x06, 0xa2, 0x34, 0x08, 0xe0, 0xf0, 0xd4, 0x6f, 0xe8, 0xd1, 0x4d, 0x45,
	0xd0, 0x26, 0xc5, 0xaf, 0x9b, 0x18, 0x1d, 0x02, 0xf8, 0x3e, 0xb6, 0x5d, 0x84, 0x2d, 0xd1, 0x36,
	0x91, 0xbf, 0x6d, 0x52, 0xc1, 0x84, 0xd5, 0x7c, 0x9a, 0xb1, 0xfc, 0xd5, 0x08, 0x84, 0xd9, 0xc4,
	0xfe, 0xbe, 0x99, 0xbf, 0xa5, 0xd1, 0x90, 0xae, 0x20, 0x4b, 0x8a, 0x41, 0xb0, 0x5e, 0x16, 0x64,
	0x19, 0xd9, 0xcf, 0x76, 0x3c, 0x90, 0xc8, 0x48, 0x39, 0x9e, 0x89, 0xfe, 0x03, 0x69, 0x5a, 0x0f,
	0x01, 0x0b, 0x65, 0x2c, 0x94, 0xe3, 0x25, 0x26, 0xf6, 0x4f, 0x64, 0x0b, 0x00, 0x60, 0xd7, 0x49,
	0x94, 0xcd, 0x7d, 0xff, 0x45, 0x30, 0x01, 0xb0, 0x69, 0xb8, 0x3e, 0x21, 0xd6, 0x4c, 0xec, 0xdf,
	0xb4, 0x69, 0x5c, 0xcf, 0x0f, 0xb0, 0x1a, 0x5c, 0x66, 0x6f, 0x95, 0x4c, 0xf0, 0x3f, 0x08, 0x77,
	0x09, 0xf8, 0xe5, 0x44, 0xb5, 0x95, 0x6f, 0x0e, 0xf6, 0xbf, 0x98, 0x69, 0x5a, 0xcf, 0x67, 0xd8,
	0x60, 0xa2, 0xda, 0xed, 0x14, 0xe7, 0x53, 0x03, 0xfe, 0xdf, 0x66, 0xf1, 0xca, 0xa2, 0x60, 0x0e,
	0xde, 0xcf, 0xc6, 0x5c, 0x19, 0x56, 0xc1, 0x83, 0x6c, 0x5e, 0xce, 0xcb, 0xc5, 0xac, 0x4c, 0xdc,
	0x3b, 0xe5, 0xf9, 0x6a, 0x35, 0x6d, 0x4d, 0xba, 0x32, 0x6c, 0xb8, 0x81, 0x4c, 0xdb, 0x0f, 0x3a,
	0x51, 0x23, 0xf0, 0x4f, 0xa6, 0x7e, 0xbb, 0x61, 0x7a, 0xb3, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe0, 0x82, 0xa0, 0xbd, 0x7c, 0x15, 0x00, 0x00,
}
