// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: management/v1/severity.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Severity represents severity level of the check result or alert.
type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_EMERGENCY   Severity = 1
	Severity_SEVERITY_ALERT       Severity = 2
	Severity_SEVERITY_CRITICAL    Severity = 3
	Severity_SEVERITY_ERROR       Severity = 4
	Severity_SEVERITY_WARNING     Severity = 5
	Severity_SEVERITY_NOTICE      Severity = 6
	Severity_SEVERITY_INFO        Severity = 7
	Severity_SEVERITY_DEBUG       Severity = 8
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_EMERGENCY",
		2: "SEVERITY_ALERT",
		3: "SEVERITY_CRITICAL",
		4: "SEVERITY_ERROR",
		5: "SEVERITY_WARNING",
		6: "SEVERITY_NOTICE",
		7: "SEVERITY_INFO",
		8: "SEVERITY_DEBUG",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_EMERGENCY":   1,
		"SEVERITY_ALERT":       2,
		"SEVERITY_CRITICAL":    3,
		"SEVERITY_ERROR":       4,
		"SEVERITY_WARNING":     5,
		"SEVERITY_NOTICE":      6,
		"SEVERITY_INFO":        7,
		"SEVERITY_DEBUG":       8,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_management_v1_severity_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_management_v1_severity_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_management_v1_severity_proto_rawDescGZIP(), []int{0}
}

var File_management_v1_severity_proto protoreflect.FileDescriptor

var file_management_v1_severity_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2a, 0xcd, 0x01,
	0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45,
	0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49,
	0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x43, 0x45, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x08, 0x42, 0xae, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58,
	0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x19, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_severity_proto_rawDescOnce sync.Once
	file_management_v1_severity_proto_rawDescData = file_management_v1_severity_proto_rawDesc
)

func file_management_v1_severity_proto_rawDescGZIP() []byte {
	file_management_v1_severity_proto_rawDescOnce.Do(func() {
		file_management_v1_severity_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_severity_proto_rawDescData)
	})
	return file_management_v1_severity_proto_rawDescData
}

var (
	file_management_v1_severity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_management_v1_severity_proto_goTypes   = []any{
		(Severity)(0), // 0: management.v1.Severity
	}
)

var file_management_v1_severity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_v1_severity_proto_init() }
func file_management_v1_severity_proto_init() {
	if File_management_v1_severity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_severity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_severity_proto_goTypes,
		DependencyIndexes: file_management_v1_severity_proto_depIdxs,
		EnumInfos:         file_management_v1_severity_proto_enumTypes,
	}.Build()
	File_management_v1_severity_proto = out.File
	file_management_v1_severity_proto_rawDesc = nil
	file_management_v1_severity_proto_goTypes = nil
	file_management_v1_severity_proto_depIdxs = nil
}
