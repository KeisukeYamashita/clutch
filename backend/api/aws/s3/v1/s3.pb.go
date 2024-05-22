// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: aws/s3/v1/s3.proto

package s3v1

import (
	_ "github.com/lyft/clutch/backend/api/api/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region  string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_s3_v1_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_aws_s3_v1_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_aws_s3_v1_s3_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bucket) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Bucket) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type AccessPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bucket          string                 `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Alias           string                 `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	CreationDate    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	BucketAccountId string                 `protobuf:"bytes,5,opt,name=bucket_account_id,json=bucketAccountId,proto3" json:"bucket_account_id,omitempty"`
	AccessPointArn  string                 `protobuf:"bytes,6,opt,name=access_point_arn,json=accessPointArn,proto3" json:"access_point_arn,omitempty"`
	Region          string                 `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	Account         string                 `protobuf:"bytes,8,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccessPoint) Reset() {
	*x = AccessPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_s3_v1_s3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPoint) ProtoMessage() {}

func (x *AccessPoint) ProtoReflect() protoreflect.Message {
	mi := &file_aws_s3_v1_s3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPoint.ProtoReflect.Descriptor instead.
func (*AccessPoint) Descriptor() ([]byte, []int) {
	return file_aws_s3_v1_s3_proto_rawDescGZIP(), []int{1}
}

func (x *AccessPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessPoint) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *AccessPoint) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *AccessPoint) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *AccessPoint) GetBucketAccountId() string {
	if x != nil {
		return x.BucketAccountId
	}
	return ""
}

func (x *AccessPoint) GetAccessPointArn() string {
	if x != nil {
		return x.AccessPointArn
	}
	return ""
}

func (x *AccessPoint) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AccessPoint) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_aws_s3_v1_s3_proto protoreflect.FileDescriptor

var file_aws_s3_v1_s3_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x77, 0x73, 0x2f, 0x73, 0x33, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x33, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x73, 0x33, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x3a, 0x3a, 0xb2, 0xe1, 0x1c, 0x36, 0x0a, 0x34, 0x0a, 0x17, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x19, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x7b,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x22, 0xd9,
	0x02, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x72,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3f, 0xb2, 0xe1, 0x1c, 0x3b, 0x0a,
	0x39, 0x0a, 0x1c, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x19, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x77, 0x73, 0x2f, 0x73, 0x33, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x33, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aws_s3_v1_s3_proto_rawDescOnce sync.Once
	file_aws_s3_v1_s3_proto_rawDescData = file_aws_s3_v1_s3_proto_rawDesc
)

func file_aws_s3_v1_s3_proto_rawDescGZIP() []byte {
	file_aws_s3_v1_s3_proto_rawDescOnce.Do(func() {
		file_aws_s3_v1_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_aws_s3_v1_s3_proto_rawDescData)
	})
	return file_aws_s3_v1_s3_proto_rawDescData
}

var file_aws_s3_v1_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aws_s3_v1_s3_proto_goTypes = []interface{}{
	(*Bucket)(nil),                // 0: clutch.aws.s3.v1.Bucket
	(*AccessPoint)(nil),           // 1: clutch.aws.s3.v1.AccessPoint
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_aws_s3_v1_s3_proto_depIdxs = []int32{
	2, // 0: clutch.aws.s3.v1.AccessPoint.creation_date:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aws_s3_v1_s3_proto_init() }
func file_aws_s3_v1_s3_proto_init() {
	if File_aws_s3_v1_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aws_s3_v1_s3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
		file_aws_s3_v1_s3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPoint); i {
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
			RawDescriptor: file_aws_s3_v1_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aws_s3_v1_s3_proto_goTypes,
		DependencyIndexes: file_aws_s3_v1_s3_proto_depIdxs,
		MessageInfos:      file_aws_s3_v1_s3_proto_msgTypes,
	}.Build()
	File_aws_s3_v1_s3_proto = out.File
	file_aws_s3_v1_s3_proto_rawDesc = nil
	file_aws_s3_v1_s3_proto_goTypes = nil
	file_aws_s3_v1_s3_proto_depIdxs = nil
}
