// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto-provisioner/service.proto

package proto_provisioner

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{0}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type BasicConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicConfig.ProtoReflect.Descriptor instead.
func (*BasicConfig) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{2}
}

func (x *BasicConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BasicConfig) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EnvConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigType:
	//
	//	*EnvConfig_BasicConfig
	ConfigType isEnvConfig_ConfigType `protobuf_oneof:"config_type"`
}

func (x *EnvConfig) Reset() {
	*x = EnvConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvConfig) ProtoMessage() {}

func (x *EnvConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvConfig.ProtoReflect.Descriptor instead.
func (*EnvConfig) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{3}
}

func (m *EnvConfig) GetConfigType() isEnvConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *EnvConfig) GetBasicConfig() *BasicConfig {
	if x, ok := x.GetConfigType().(*EnvConfig_BasicConfig); ok {
		return x.BasicConfig
	}
	return nil
}

type isEnvConfig_ConfigType interface {
	isEnvConfig_ConfigType()
}

type EnvConfig_BasicConfig struct {
	BasicConfig *BasicConfig `protobuf:"bytes,1,opt,name=basic_config,json=basicConfig,proto3,oneof"`
}

func (*EnvConfig_BasicConfig) isEnvConfig_ConfigType() {}

type TaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{4}
}

func (x *TaskCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskCreateRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type TaskCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TaskCreateResponse) Reset() {
	*x = TaskCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_provisioner_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateResponse) ProtoMessage() {}

func (x *TaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provisioner_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_provisioner_service_proto_rawDescGZIP(), []int{5}
}

func (x *TaskCreateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_provisioner_service_proto protoreflect.FileDescriptor

var file_proto_provisioner_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x33, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3d, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x61, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xb5, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_provisioner_service_proto_rawDescOnce sync.Once
	file_proto_provisioner_service_proto_rawDescData = file_proto_provisioner_service_proto_rawDesc
)

func file_proto_provisioner_service_proto_rawDescGZIP() []byte {
	file_proto_provisioner_service_proto_rawDescOnce.Do(func() {
		file_proto_provisioner_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_provisioner_service_proto_rawDescData)
	})
	return file_proto_provisioner_service_proto_rawDescData
}

var file_proto_provisioner_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_provisioner_service_proto_goTypes = []any{
	(*StatusRequest)(nil),      // 0: provisioner.StatusRequest
	(*StatusResponse)(nil),     // 1: provisioner.StatusResponse
	(*BasicConfig)(nil),        // 2: provisioner.BasicConfig
	(*EnvConfig)(nil),          // 3: provisioner.EnvConfig
	(*TaskCreateRequest)(nil),  // 4: provisioner.TaskCreateRequest
	(*TaskCreateResponse)(nil), // 5: provisioner.TaskCreateResponse
	(*anypb.Any)(nil),          // 6: google.protobuf.Any
}
var file_proto_provisioner_service_proto_depIdxs = []int32{
	2, // 0: provisioner.EnvConfig.basic_config:type_name -> provisioner.BasicConfig
	6, // 1: provisioner.TaskCreateRequest.data:type_name -> google.protobuf.Any
	0, // 2: provisioner.ProvisionerService.GetProvisionerStatus:input_type -> provisioner.StatusRequest
	4, // 3: provisioner.ProvisionerService.ExecuteTask:input_type -> provisioner.TaskCreateRequest
	1, // 4: provisioner.ProvisionerService.GetProvisionerStatus:output_type -> provisioner.StatusResponse
	5, // 5: provisioner.ProvisionerService.ExecuteTask:output_type -> provisioner.TaskCreateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_provisioner_service_proto_init() }
func file_proto_provisioner_service_proto_init() {
	if File_proto_provisioner_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_provisioner_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatusRequest); i {
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
		file_proto_provisioner_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatusResponse); i {
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
		file_proto_provisioner_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BasicConfig); i {
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
		file_proto_provisioner_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EnvConfig); i {
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
		file_proto_provisioner_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TaskCreateRequest); i {
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
		file_proto_provisioner_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TaskCreateResponse); i {
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
	file_proto_provisioner_service_proto_msgTypes[3].OneofWrappers = []any{
		(*EnvConfig_BasicConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_provisioner_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_provisioner_service_proto_goTypes,
		DependencyIndexes: file_proto_provisioner_service_proto_depIdxs,
		MessageInfos:      file_proto_provisioner_service_proto_msgTypes,
	}.Build()
	File_proto_provisioner_service_proto = out.File
	file_proto_provisioner_service_proto_rawDesc = nil
	file_proto_provisioner_service_proto_goTypes = nil
	file_proto_provisioner_service_proto_depIdxs = nil
}
