// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: vehicle.proto

package carserver

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

type ClimateState_CopActivationTemp int32

const (
	ClimateState_CopActivationTempUnspecified ClimateState_CopActivationTemp = 0
	ClimateState_CopActivationTempLow         ClimateState_CopActivationTemp = 1
	ClimateState_CopActivationTempMedium      ClimateState_CopActivationTemp = 2
	ClimateState_CopActivationTempHigh        ClimateState_CopActivationTemp = 3
)

// Enum value maps for ClimateState_CopActivationTemp.
var (
	ClimateState_CopActivationTemp_name = map[int32]string{
		0: "CopActivationTempUnspecified",
		1: "CopActivationTempLow",
		2: "CopActivationTempMedium",
		3: "CopActivationTempHigh",
	}
	ClimateState_CopActivationTemp_value = map[string]int32{
		"CopActivationTempUnspecified": 0,
		"CopActivationTempLow":         1,
		"CopActivationTempMedium":      2,
		"CopActivationTempHigh":        3,
	}
)

func (x ClimateState_CopActivationTemp) Enum() *ClimateState_CopActivationTemp {
	p := new(ClimateState_CopActivationTemp)
	*p = x
	return p
}

func (x ClimateState_CopActivationTemp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClimateState_CopActivationTemp) Descriptor() protoreflect.EnumDescriptor {
	return file_vehicle_proto_enumTypes[0].Descriptor()
}

func (ClimateState_CopActivationTemp) Type() protoreflect.EnumType {
	return &file_vehicle_proto_enumTypes[0]
}

func (x ClimateState_CopActivationTemp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClimateState_CopActivationTemp.Descriptor instead.
func (ClimateState_CopActivationTemp) EnumDescriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{1, 0}
}

type VehicleState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestMode *VehicleState_GuestMode `protobuf:"bytes,74,opt,name=guestMode,proto3" json:"guestMode,omitempty"`
}

func (x *VehicleState) Reset() {
	*x = VehicleState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleState) ProtoMessage() {}

func (x *VehicleState) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleState.ProtoReflect.Descriptor instead.
func (*VehicleState) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleState) GetGuestMode() *VehicleState_GuestMode {
	if x != nil {
		return x.GuestMode
	}
	return nil
}

type ClimateState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClimateState) Reset() {
	*x = ClimateState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClimateState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClimateState) ProtoMessage() {}

func (x *ClimateState) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClimateState.ProtoReflect.Descriptor instead.
func (*ClimateState) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{1}
}

type VehicleState_GuestMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestModeActive bool `protobuf:"varint,1,opt,name=GuestModeActive,proto3" json:"GuestModeActive,omitempty"`
}

func (x *VehicleState_GuestMode) Reset() {
	*x = VehicleState_GuestMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleState_GuestMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleState_GuestMode) ProtoMessage() {}

func (x *VehicleState_GuestMode) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleState_GuestMode.ProtoReflect.Descriptor instead.
func (*VehicleState_GuestMode) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VehicleState_GuestMode) GetGuestModeActive() bool {
	if x != nil {
		return x.GuestModeActive
	}
	return false
}

var File_vehicle_proto protoreflect.FileDescriptor

var file_vehicle_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x4a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x35, 0x0a, 0x09,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x6f,
	0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x55,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x43, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d,
	0x70, 0x4c, 0x6f, 0x77, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x6f, 0x70, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x48, 0x69, 0x67, 0x68, 0x10, 0x03, 0x42, 0x78,
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x73,
	0x6c, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_proto_rawDescOnce sync.Once
	file_vehicle_proto_rawDescData = file_vehicle_proto_rawDesc
)

func file_vehicle_proto_rawDescGZIP() []byte {
	file_vehicle_proto_rawDescOnce.Do(func() {
		file_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_proto_rawDescData)
	})
	return file_vehicle_proto_rawDescData
}

var file_vehicle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vehicle_proto_goTypes = []interface{}{
	(ClimateState_CopActivationTemp)(0), // 0: CarServer.ClimateState.CopActivationTemp
	(*VehicleState)(nil),                // 1: CarServer.VehicleState
	(*ClimateState)(nil),                // 2: CarServer.ClimateState
	(*VehicleState_GuestMode)(nil),      // 3: CarServer.VehicleState.GuestMode
}
var file_vehicle_proto_depIdxs = []int32{
	3, // 0: CarServer.VehicleState.guestMode:type_name -> CarServer.VehicleState.GuestMode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vehicle_proto_init() }
func file_vehicle_proto_init() {
	if File_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleState); i {
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
		file_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClimateState); i {
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
		file_vehicle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleState_GuestMode); i {
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
			RawDescriptor: file_vehicle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vehicle_proto_goTypes,
		DependencyIndexes: file_vehicle_proto_depIdxs,
		EnumInfos:         file_vehicle_proto_enumTypes,
		MessageInfos:      file_vehicle_proto_msgTypes,
	}.Build()
	File_vehicle_proto = out.File
	file_vehicle_proto_rawDesc = nil
	file_vehicle_proto_goTypes = nil
	file_vehicle_proto_depIdxs = nil
}
