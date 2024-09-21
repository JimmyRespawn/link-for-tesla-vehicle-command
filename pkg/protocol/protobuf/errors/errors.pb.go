// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: errors.proto

package errors

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

type GenericError_E int32

const (
	GenericError_E_GENERICERROR_NONE                       GenericError_E = 0
	GenericError_E_GENERICERROR_UNKNOWN                    GenericError_E = 1
	GenericError_E_GENERICERROR_CLOSURES_OPEN              GenericError_E = 2
	GenericError_E_GENERICERROR_ALREADY_ON                 GenericError_E = 3
	GenericError_E_GENERICERROR_DISABLED_FOR_USER_COMMAND  GenericError_E = 4
	GenericError_E_GENERICERROR_VEHICLE_NOT_IN_PARK        GenericError_E = 5
	GenericError_E_GENERICERROR_UNAUTHORIZED               GenericError_E = 6
	GenericError_E_GENERICERROR_NOT_ALLOWED_OVER_TRANSPORT GenericError_E = 7
)

// Enum value maps for GenericError_E.
var (
	GenericError_E_name = map[int32]string{
		0: "GENERICERROR_NONE",
		1: "GENERICERROR_UNKNOWN",
		2: "GENERICERROR_CLOSURES_OPEN",
		3: "GENERICERROR_ALREADY_ON",
		4: "GENERICERROR_DISABLED_FOR_USER_COMMAND",
		5: "GENERICERROR_VEHICLE_NOT_IN_PARK",
		6: "GENERICERROR_UNAUTHORIZED",
		7: "GENERICERROR_NOT_ALLOWED_OVER_TRANSPORT",
	}
	GenericError_E_value = map[string]int32{
		"GENERICERROR_NONE":                       0,
		"GENERICERROR_UNKNOWN":                    1,
		"GENERICERROR_CLOSURES_OPEN":              2,
		"GENERICERROR_ALREADY_ON":                 3,
		"GENERICERROR_DISABLED_FOR_USER_COMMAND":  4,
		"GENERICERROR_VEHICLE_NOT_IN_PARK":        5,
		"GENERICERROR_UNAUTHORIZED":               6,
		"GENERICERROR_NOT_ALLOWED_OVER_TRANSPORT": 7,
	}
)

func (x GenericError_E) Enum() *GenericError_E {
	p := new(GenericError_E)
	*p = x
	return p
}

func (x GenericError_E) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenericError_E) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_proto_enumTypes[0].Descriptor()
}

func (GenericError_E) Type() protoreflect.EnumType {
	return &file_errors_proto_enumTypes[0]
}

func (x GenericError_E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenericError_E.Descriptor instead.
func (GenericError_E) EnumDescriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{0}
}

type NominalError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenericError GenericError_E `protobuf:"varint,1,opt,name=genericError,proto3,enum=Errors.GenericError_E" json:"genericError,omitempty"`
}

func (x *NominalError) Reset() {
	*x = NominalError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NominalError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NominalError) ProtoMessage() {}

func (x *NominalError) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NominalError.ProtoReflect.Descriptor instead.
func (*NominalError) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{0}
}

func (x *NominalError) GetGenericError() GenericError_E {
	if x != nil {
		return x.GenericError
	}
	return GenericError_E_GENERICERROR_NONE
}

var File_errors_proto protoreflect.FileDescriptor

var file_errors_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x4e, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x45, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x2a, 0x9c, 0x02, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x45, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49,
	0x43, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x55, 0x52, 0x45, 0x53, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49,
	0x43, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x2a, 0x0a, 0x26, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x04, 0x12,
	0x24, 0x0a, 0x20, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x50,
	0x41, 0x52, 0x4b, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x2b, 0x0a, 0x27, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10,
	0x07, 0x42, 0x61, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x6c,
	0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_proto_rawDescOnce sync.Once
	file_errors_proto_rawDescData = file_errors_proto_rawDesc
)

func file_errors_proto_rawDescGZIP() []byte {
	file_errors_proto_rawDescOnce.Do(func() {
		file_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_proto_rawDescData)
	})
	return file_errors_proto_rawDescData
}

var file_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_proto_goTypes = []interface{}{
	(GenericError_E)(0),  // 0: Errors.GenericError_E
	(*NominalError)(nil), // 1: Errors.NominalError
}
var file_errors_proto_depIdxs = []int32{
	0, // 0: Errors.NominalError.genericError:type_name -> Errors.GenericError_E
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_proto_init() }
func file_errors_proto_init() {
	if File_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NominalError); i {
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
			RawDescriptor: file_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_proto_goTypes,
		DependencyIndexes: file_errors_proto_depIdxs,
		EnumInfos:         file_errors_proto_enumTypes,
		MessageInfos:      file_errors_proto_msgTypes,
	}.Build()
	File_errors_proto = out.File
	file_errors_proto_rawDesc = nil
	file_errors_proto_goTypes = nil
	file_errors_proto_depIdxs = nil
}
