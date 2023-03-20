// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: core/project/v1/project.proto

package projectv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A project is the logical abstraction for a service, library, mobile app, etc.
// Adding this higher level primitive will enable a more generic way of
// addressing commonalities between various system types.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tier      string   `protobuf:"bytes,2,opt,name=tier,proto3" json:"tier,omitempty"`
	Owners    []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	Languages []string `protobuf:"bytes,4,rep,name=languages,proto3" json:"languages,omitempty"`
	// Allows an organization to populate any information they see fit that does not fit our schema
	Data         map[string]*structpb.Value `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dependencies *ProjectDependencies       `protobuf:"bytes,6,opt,name=dependencies,proto3" json:"dependencies,omitempty"`
	Oncall       *OnCall                    `protobuf:"bytes,7,opt,name=oncall,proto3" json:"oncall,omitempty"`
	LinkGroups   []*LinkGroup               `protobuf:"bytes,8,rep,name=link_groups,json=linkGroups,proto3" json:"link_groups,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *Project) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *Project) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *Project) GetData() map[string]*structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Project) GetDependencies() *ProjectDependencies {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Project) GetOncall() *OnCall {
	if x != nil {
		return x.Oncall
	}
	return nil
}

func (x *Project) GetLinkGroups() []*LinkGroup {
	if x != nil {
		return x.LinkGroups
	}
	return nil
}

// Dependencies could be either upstream or downstream of this project.
type ProjectDependencies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key in this map is the type_url of the dependency
	Upstreams   map[string]*Dependency `protobuf:"bytes,1,rep,name=upstreams,proto3" json:"upstreams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Downstreams map[string]*Dependency `protobuf:"bytes,2,rep,name=downstreams,proto3" json:"downstreams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProjectDependencies) Reset() {
	*x = ProjectDependencies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDependencies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDependencies) ProtoMessage() {}

func (x *ProjectDependencies) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDependencies.ProtoReflect.Descriptor instead.
func (*ProjectDependencies) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectDependencies) GetUpstreams() map[string]*Dependency {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

func (x *ProjectDependencies) GetDownstreams() map[string]*Dependency {
	if x != nil {
		return x.Downstreams
	}
	return nil
}

type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{2}
}

func (x *Dependency) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Oncall information that could come from several sources for a project
// From active alerts from source like pagerduty, to who is currently on call.
type OnCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagerduty *PagerDuty `protobuf:"bytes,1,opt,name=pagerduty,proto3" json:"pagerduty,omitempty"`
}

func (x *OnCall) Reset() {
	*x = OnCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnCall) ProtoMessage() {}

func (x *OnCall) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnCall.ProtoReflect.Descriptor instead.
func (*OnCall) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{3}
}

func (x *OnCall) GetPagerduty() *PagerDuty {
	if x != nil {
		return x.Pagerduty
	}
	return nil
}

// PagerDuty specific information for a project
type PagerDuty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A project may have more than one PagerDuty service id
	// https://developer.pagerduty.com/api-reference/docs/CONCEPTS.md#services
	ServiceIds []string `protobuf:"bytes,1,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
}

func (x *PagerDuty) Reset() {
	*x = PagerDuty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagerDuty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagerDuty) ProtoMessage() {}

func (x *PagerDuty) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagerDuty.ProtoReflect.Descriptor instead.
func (*PagerDuty) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{4}
}

func (x *PagerDuty) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{5}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type LinkGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Links     []*Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	ImagePath string  `protobuf:"bytes,3,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"` // Can be an s3 location, path on the server, etc.
}

func (x *LinkGroup) Reset() {
	*x = LinkGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_project_v1_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkGroup) ProtoMessage() {}

func (x *LinkGroup) ProtoReflect() protoreflect.Message {
	mi := &file_core_project_v1_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkGroup.ProtoReflect.Descriptor instead.
func (*LinkGroup) Descriptor() ([]byte, []int) {
	return file_core_project_v1_project_proto_rawDescGZIP(), []int{6}
}

func (x *LinkGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LinkGroup) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *LinkGroup) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

var File_core_project_v1_project_proto protoreflect.FileDescriptor

var file_core_project_v1_project_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4,
	0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x6f, 0x6e, 0x63, 0x61,
	0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x6f, 0x6e, 0x63, 0x61, 0x6c, 0x6c,
	0x12, 0x42, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x1a, 0x4f, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x95, 0x03, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x58, 0x0a,
	0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x55, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x5e, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x60, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x62, 0x0a, 0x10, 0x44, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e, 0x0a,
	0x0a, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x49, 0x0a,
	0x06, 0x4f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x3f, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x64, 0x75, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x22, 0x2c, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x75, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x7b, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_project_v1_project_proto_rawDescOnce sync.Once
	file_core_project_v1_project_proto_rawDescData = file_core_project_v1_project_proto_rawDesc
)

func file_core_project_v1_project_proto_rawDescGZIP() []byte {
	file_core_project_v1_project_proto_rawDescOnce.Do(func() {
		file_core_project_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_project_v1_project_proto_rawDescData)
	})
	return file_core_project_v1_project_proto_rawDescData
}

var file_core_project_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_core_project_v1_project_proto_goTypes = []interface{}{
	(*Project)(nil),             // 0: clutch.core.project.v1.Project
	(*ProjectDependencies)(nil), // 1: clutch.core.project.v1.ProjectDependencies
	(*Dependency)(nil),          // 2: clutch.core.project.v1.Dependency
	(*OnCall)(nil),              // 3: clutch.core.project.v1.OnCall
	(*PagerDuty)(nil),           // 4: clutch.core.project.v1.PagerDuty
	(*Link)(nil),                // 5: clutch.core.project.v1.Link
	(*LinkGroup)(nil),           // 6: clutch.core.project.v1.LinkGroup
	nil,                         // 7: clutch.core.project.v1.Project.DataEntry
	nil,                         // 8: clutch.core.project.v1.ProjectDependencies.UpstreamsEntry
	nil,                         // 9: clutch.core.project.v1.ProjectDependencies.DownstreamsEntry
	(*structpb.Value)(nil),      // 10: google.protobuf.Value
}
var file_core_project_v1_project_proto_depIdxs = []int32{
	7,  // 0: clutch.core.project.v1.Project.data:type_name -> clutch.core.project.v1.Project.DataEntry
	1,  // 1: clutch.core.project.v1.Project.dependencies:type_name -> clutch.core.project.v1.ProjectDependencies
	3,  // 2: clutch.core.project.v1.Project.oncall:type_name -> clutch.core.project.v1.OnCall
	6,  // 3: clutch.core.project.v1.Project.link_groups:type_name -> clutch.core.project.v1.LinkGroup
	8,  // 4: clutch.core.project.v1.ProjectDependencies.upstreams:type_name -> clutch.core.project.v1.ProjectDependencies.UpstreamsEntry
	9,  // 5: clutch.core.project.v1.ProjectDependencies.downstreams:type_name -> clutch.core.project.v1.ProjectDependencies.DownstreamsEntry
	4,  // 6: clutch.core.project.v1.OnCall.pagerduty:type_name -> clutch.core.project.v1.PagerDuty
	5,  // 7: clutch.core.project.v1.LinkGroup.links:type_name -> clutch.core.project.v1.Link
	10, // 8: clutch.core.project.v1.Project.DataEntry.value:type_name -> google.protobuf.Value
	2,  // 9: clutch.core.project.v1.ProjectDependencies.UpstreamsEntry.value:type_name -> clutch.core.project.v1.Dependency
	2,  // 10: clutch.core.project.v1.ProjectDependencies.DownstreamsEntry.value:type_name -> clutch.core.project.v1.Dependency
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_core_project_v1_project_proto_init() }
func file_core_project_v1_project_proto_init() {
	if File_core_project_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_project_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_core_project_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDependencies); i {
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
		file_core_project_v1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
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
		file_core_project_v1_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnCall); i {
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
		file_core_project_v1_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagerDuty); i {
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
		file_core_project_v1_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_core_project_v1_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkGroup); i {
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
			RawDescriptor: file_core_project_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_project_v1_project_proto_goTypes,
		DependencyIndexes: file_core_project_v1_project_proto_depIdxs,
		MessageInfos:      file_core_project_v1_project_proto_msgTypes,
	}.Build()
	File_core_project_v1_project_proto = out.File
	file_core_project_v1_project_proto_rawDesc = nil
	file_core_project_v1_project_proto_goTypes = nil
	file_core_project_v1_project_proto_depIdxs = nil
}
