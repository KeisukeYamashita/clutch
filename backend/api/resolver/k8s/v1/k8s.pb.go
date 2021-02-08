// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: resolver/k8s/v1/k8s.proto

package k8sv1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/clutch/backend/api/resolver/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PodID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *PodID) Reset() {
	*x = PodID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodID) ProtoMessage() {}

func (x *PodID) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodID.ProtoReflect.Descriptor instead.
func (*PodID) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{0}
}

func (x *PodID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodID) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *PodID) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type IPAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *IPAddress) Reset() {
	*x = IPAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAddress) ProtoMessage() {}

func (x *IPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAddress.ProtoReflect.Descriptor instead.
func (*IPAddress) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{1}
}

func (x *IPAddress) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type HPAName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *HPAName) Reset() {
	*x = HPAName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HPAName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HPAName) ProtoMessage() {}

func (x *HPAName) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HPAName.ProtoReflect.Descriptor instead.
func (*HPAName) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{2}
}

func (x *HPAName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HPAName) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *HPAName) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{3}
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *Deployment) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type StatefulSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *StatefulSet) Reset() {
	*x = StatefulSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSet) ProtoMessage() {}

func (x *StatefulSet) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSet.ProtoReflect.Descriptor instead.
func (*StatefulSet) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{4}
}

func (x *StatefulSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatefulSet) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *StatefulSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{5}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *Service) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CronJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CronJob) Reset() {
	*x = CronJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJob) ProtoMessage() {}

func (x *CronJob) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJob.ProtoReflect.Descriptor instead.
func (*CronJob) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{6}
}

func (x *CronJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronJob) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *CronJob) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ConfigMap) Reset() {
	*x = ConfigMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMap) ProtoMessage() {}

func (x *ConfigMap) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMap.ProtoReflect.Descriptor instead.
func (*ConfigMap) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigMap) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigMap) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *ConfigMap) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Clientset string `protobuf:"bytes,2,opt,name=clientset,proto3" json:"clientset,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_k8s_v1_k8s_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_resolver_k8s_v1_k8s_proto_rawDescGZIP(), []int{8}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetClientset() string {
	if x != nil {
		return x.Clientset
	}
	return ""
}

func (x *Job) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_resolver_k8s_v1_k8s_proto protoreflect.FileDescriptor

var file_resolver_k8s_v1_k8s_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x2f,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xea, 0x9f,
	0x1d, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x0d, 0x0a, 0x0b, 0x6d, 0x79,
	0x2d, 0x70, 0x6f, 0x64, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x74, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x21, 0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x10,
	0xea, 0x9f, 0x1d, 0x0c, 0x0a, 0x06, 0x70, 0x6f, 0x64, 0x20, 0x49, 0x44, 0x1a, 0x02, 0x08, 0x01,
	0x22, 0x60, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x0a, 0x49, 0x50, 0x20, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x10, 0x01, 0x1a, 0x0a, 0x0a, 0x08, 0x31, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2e,
	0x31, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x14, 0xea, 0x9f,
	0x1d, 0x10, 0x0a, 0x0a, 0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x02,
	0x08, 0x01, 0x22, 0xc9, 0x01, 0x0a, 0x07, 0x48, 0x50, 0x41, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xea, 0x9f,
	0x1d, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x0d, 0x0a, 0x0b, 0x6d, 0x79,
	0x2d, 0x68, 0x70, 0x61, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x74, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x21, 0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e,
	0xea, 0x9f, 0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x22, 0xd3,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xea, 0x9f, 0x1d,
	0x1e, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x14, 0x0a, 0x12, 0x6d, 0x79, 0x2d,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e, 0xea, 0x9f, 0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x02, 0x08, 0x01, 0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75,
	0x6c, 0x53, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x23, 0xea, 0x9f, 0x1d, 0x1f, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01,
	0x1a, 0x15, 0x0a, 0x13, 0x6d, 0x79, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73,
	0x65, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65,
	0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21,
	0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10,
	0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e, 0xea, 0x9f,
	0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x22, 0xcd, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xea, 0x9f, 0x1d, 0x1b, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x10, 0x01, 0x1a, 0x11, 0x0a, 0x0f, 0x6d, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65,
	0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21,
	0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10,
	0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e, 0xea, 0x9f,
	0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x22, 0xcd, 0x01, 0x0a,
	0x07, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x33, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xea, 0x9f, 0x1d, 0x1b, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x10, 0x01, 0x1a, 0x11, 0x0a, 0x0f, 0x6d, 0x79, 0x2d, 0x63, 0x72, 0x6f, 0x6e, 0x6a,
	0x6f, 0x62, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65,
	0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21,
	0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10,
	0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e, 0xea, 0x9f,
	0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x22, 0xd2, 0x01, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xea, 0x9f, 0x1d, 0x1e, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x14, 0x0a, 0x12, 0x6d, 0x79, 0x2d, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2d, 0x6d, 0x61, 0x70, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xea, 0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x10, 0x01, 0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x65, 0x74, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74,
	0x12, 0x3f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x21, 0xea, 0x9f, 0x1d, 0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x3a, 0x0e, 0xea, 0x9f, 0x1d, 0x0a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08,
	0x01, 0x22, 0xc5, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xea, 0x9f, 0x1d, 0x17, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x10, 0x01, 0x1a, 0x0d, 0x0a, 0x0b, 0x6d, 0x79, 0x2d, 0x6a, 0x6f, 0x62, 0x2d,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xea,
	0x9f, 0x1d, 0x1a, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x10, 0x01,
	0x22, 0x0b, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x52, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xea, 0x9f, 0x1d,
	0x1d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10, 0x01, 0x1a, 0x0e,
	0x0a, 0x0c, 0x6d, 0x79, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3a, 0x0e, 0xea, 0x9f, 0x1d, 0x0a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x6b, 0x38, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resolver_k8s_v1_k8s_proto_rawDescOnce sync.Once
	file_resolver_k8s_v1_k8s_proto_rawDescData = file_resolver_k8s_v1_k8s_proto_rawDesc
)

func file_resolver_k8s_v1_k8s_proto_rawDescGZIP() []byte {
	file_resolver_k8s_v1_k8s_proto_rawDescOnce.Do(func() {
		file_resolver_k8s_v1_k8s_proto_rawDescData = protoimpl.X.CompressGZIP(file_resolver_k8s_v1_k8s_proto_rawDescData)
	})
	return file_resolver_k8s_v1_k8s_proto_rawDescData
}

var file_resolver_k8s_v1_k8s_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_resolver_k8s_v1_k8s_proto_goTypes = []interface{}{
	(*PodID)(nil),       // 0: clutch.resolver.k8s.v1.PodID
	(*IPAddress)(nil),   // 1: clutch.resolver.k8s.v1.IPAddress
	(*HPAName)(nil),     // 2: clutch.resolver.k8s.v1.HPAName
	(*Deployment)(nil),  // 3: clutch.resolver.k8s.v1.Deployment
	(*StatefulSet)(nil), // 4: clutch.resolver.k8s.v1.StatefulSet
	(*Service)(nil),     // 5: clutch.resolver.k8s.v1.Service
	(*CronJob)(nil),     // 6: clutch.resolver.k8s.v1.CronJob
	(*ConfigMap)(nil),   // 7: clutch.resolver.k8s.v1.ConfigMap
	(*Job)(nil),         // 8: clutch.resolver.k8s.v1.Job
}
var file_resolver_k8s_v1_k8s_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resolver_k8s_v1_k8s_proto_init() }
func file_resolver_k8s_v1_k8s_proto_init() {
	if File_resolver_k8s_v1_k8s_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resolver_k8s_v1_k8s_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodID); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAddress); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HPAName); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSet); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJob); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMap); i {
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
		file_resolver_k8s_v1_k8s_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
			RawDescriptor: file_resolver_k8s_v1_k8s_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resolver_k8s_v1_k8s_proto_goTypes,
		DependencyIndexes: file_resolver_k8s_v1_k8s_proto_depIdxs,
		MessageInfos:      file_resolver_k8s_v1_k8s_proto_msgTypes,
	}.Build()
	File_resolver_k8s_v1_k8s_proto = out.File
	file_resolver_k8s_v1_k8s_proto_rawDesc = nil
	file_resolver_k8s_v1_k8s_proto_goTypes = nil
	file_resolver_k8s_v1_k8s_proto_depIdxs = nil
}
