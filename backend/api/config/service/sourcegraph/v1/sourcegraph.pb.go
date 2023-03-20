// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: config/service/sourcegraph/v1/sourcegraph.proto

package sourcegraphv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TimeoutMs int64  `protobuf:"varint,3,opt,name=timeout_ms,json=timeoutMs,proto3" json:"timeout_ms,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_sourcegraph_v1_sourcegraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_sourcegraph_v1_sourcegraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Config) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Config) GetTimeoutMs() int64 {
	if x != nil {
		return x.TimeoutMs
	}
	return 0
}

var File_config_service_sourcegraph_v1_sourcegraph_proto protoreflect.FileDescriptor

var file_config_service_sourcegraph_v1_sourcegraph_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x24, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x63, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x4d, 0x73, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescOnce sync.Once
	file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescData = file_config_service_sourcegraph_v1_sourcegraph_proto_rawDesc
)

func file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescGZIP() []byte {
	file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescOnce.Do(func() {
		file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescData)
	})
	return file_config_service_sourcegraph_v1_sourcegraph_proto_rawDescData
}

var file_config_service_sourcegraph_v1_sourcegraph_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_service_sourcegraph_v1_sourcegraph_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: clutch.config.service.sourcegraph.v1.Config
}
var file_config_service_sourcegraph_v1_sourcegraph_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_service_sourcegraph_v1_sourcegraph_proto_init() }
func file_config_service_sourcegraph_v1_sourcegraph_proto_init() {
	if File_config_service_sourcegraph_v1_sourcegraph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_sourcegraph_v1_sourcegraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_config_service_sourcegraph_v1_sourcegraph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_sourcegraph_v1_sourcegraph_proto_goTypes,
		DependencyIndexes: file_config_service_sourcegraph_v1_sourcegraph_proto_depIdxs,
		MessageInfos:      file_config_service_sourcegraph_v1_sourcegraph_proto_msgTypes,
	}.Build()
	File_config_service_sourcegraph_v1_sourcegraph_proto = out.File
	file_config_service_sourcegraph_v1_sourcegraph_proto_rawDesc = nil
	file_config_service_sourcegraph_v1_sourcegraph_proto_goTypes = nil
	file_config_service_sourcegraph_v1_sourcegraph_proto_depIdxs = nil
}
