// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: code.proto

package rpcs

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

type Code int32

const (
	Code_Success          Code = 0
	Code_OKStar           Code = 20001
	Code_OKUnStar         Code = 20002
	Code_ErrUserExist     Code = 40001
	Code_ErrEmailLimit    Code = 40002
	Code_ErrCaptchaNil    Code = 40003
	Code_ErrCaptchaErr    Code = 40004
	Code_ErrUserNotExist  Code = 40005
	Code_ErrPasswordErr   Code = 40006
	Code_ErrRecaptchaErr  Code = 40007
	Code_ErrTopicNotFound Code = 40008
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:     "Success",
		20001: "OKStar",
		20002: "OKUnStar",
		40001: "ErrUserExist",
		40002: "ErrEmailLimit",
		40003: "ErrCaptchaNil",
		40004: "ErrCaptchaErr",
		40005: "ErrUserNotExist",
		40006: "ErrPasswordErr",
		40007: "ErrRecaptchaErr",
		40008: "ErrTopicNotFound",
	}
	Code_value = map[string]int32{
		"Success":          0,
		"OKStar":           20001,
		"OKUnStar":         20002,
		"ErrUserExist":     40001,
		"ErrEmailLimit":    40002,
		"ErrCaptchaNil":    40003,
		"ErrCaptchaErr":    40004,
		"ErrUserNotExist":  40005,
		"ErrPasswordErr":   40006,
		"ErrRecaptchaErr":  40007,
		"ErrTopicNotFound": 40008,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_code_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_code_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{0}
}

var File_code_proto protoreflect.FileDescriptor

var file_code_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xe0, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x06, 0x4f, 0x4b, 0x53,
	0x74, 0x61, 0x72, 0x10, 0xa1, 0x9c, 0x01, 0x12, 0x0e, 0x0a, 0x08, 0x4f, 0x4b, 0x55, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x10, 0xa2, 0x9c, 0x01, 0x12, 0x12, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xc1, 0xb8, 0x02, 0x12, 0x13, 0x0a, 0x0d, 0x45,
	0x72, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0xc2, 0xb8, 0x02,
	0x12, 0x13, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4e, 0x69,
	0x6c, 0x10, 0xc3, 0xb8, 0x02, 0x12, 0x13, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x45, 0x72, 0x72, 0x10, 0xc4, 0xb8, 0x02, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x72,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xc5, 0xb8,
	0x02, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x45, 0x72, 0x72, 0x10, 0xc6, 0xb8, 0x02, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x52, 0x65,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x72, 0x72, 0x10, 0xc7, 0xb8, 0x02, 0x12, 0x16,
	0x0a, 0x10, 0x45, 0x72, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0xc8, 0xb8, 0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x7a, 0x65, 0x65, 0x36, 0x2f, 0x74, 0x72, 0x65, 0x65,
	0x68, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x73, 0x3b, 0x72, 0x70,
	0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_proto_rawDescOnce sync.Once
	file_code_proto_rawDescData = file_code_proto_rawDesc
)

func file_code_proto_rawDescGZIP() []byte {
	file_code_proto_rawDescOnce.Do(func() {
		file_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_proto_rawDescData)
	})
	return file_code_proto_rawDescData
}

var file_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_code_proto_goTypes = []interface{}{
	(Code)(0), // 0: proto.Code
}
var file_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_proto_init() }
func file_code_proto_init() {
	if File_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_code_proto_goTypes,
		DependencyIndexes: file_code_proto_depIdxs,
		EnumInfos:         file_code_proto_enumTypes,
	}.Build()
	File_code_proto = out.File
	file_code_proto_rawDesc = nil
	file_code_proto_goTypes = nil
	file_code_proto_depIdxs = nil
}
