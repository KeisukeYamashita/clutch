// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: config/service/authz/v1/authz.proto

package authzv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/lyft/clutch/backend/api/api/v1"
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

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Principal_User
	//	*Principal_Group
	Type isPrincipal_Type `protobuf_oneof:"type"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authz_v1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authz_v1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_config_service_authz_v1_authz_proto_rawDescGZIP(), []int{0}
}

func (m *Principal) GetType() isPrincipal_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Principal) GetUser() string {
	if x, ok := x.GetType().(*Principal_User); ok {
		return x.User
	}
	return ""
}

func (x *Principal) GetGroup() string {
	if x, ok := x.GetType().(*Principal_Group); ok {
		return x.Group
	}
	return ""
}

type isPrincipal_Type interface {
	isPrincipal_Type()
}

type Principal_User struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type Principal_Group struct {
	Group string `protobuf:"bytes,2,opt,name=group,proto3,oneof"`
}

func (*Principal_User) isPrincipal_Type() {}

func (*Principal_Group) isPrincipal_Type() {}

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To         []string     `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Principals []*Principal `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authz_v1_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authz_v1_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_config_service_authz_v1_authz_proto_rawDescGZIP(), []int{1}
}

func (x *RoleBinding) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *RoleBinding) GetPrincipals() []*Principal {
	if x != nil {
		return x.Principals
	}
	return nil
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For logging purposes, give the policy a defined name.
	PolicyName string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// A list of acceptable action types. If left empty, all action types are accepted.
	ActionTypes []v1.ActionType `protobuf:"varint,2,rep,packed,name=action_types,json=actionTypes,proto3,enum=clutch.api.v1.ActionType" json:"action_types,omitempty"`
	// The full method in the format of a `/SERVICE/METHOD`. Wildcards are allowed, e.g. `*` or `/SERVICE/*`.
	// If left empty, all methods are accepted.
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	// Resource from the id annotation on proto objects. Wildcards are allowed.
	Resources []string `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authz_v1_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authz_v1_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_config_service_authz_v1_authz_proto_rawDescGZIP(), []int{2}
}

func (x *Policy) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *Policy) GetActionTypes() []v1.ActionType {
	if x != nil {
		return x.ActionTypes
	}
	return nil
}

func (x *Policy) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Policy) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName string    `protobuf:"bytes,1,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Policies []*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authz_v1_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authz_v1_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_config_service_authz_v1_authz_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *Role) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleBindings []*RoleBinding `protobuf:"bytes,1,rep,name=role_bindings,json=roleBindings,proto3" json:"role_bindings,omitempty"`
	Roles        []*Role        `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authz_v1_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authz_v1_authz_proto_msgTypes[4]
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
	return file_config_service_authz_v1_authz_proto_rawDescGZIP(), []int{4}
}

func (x *Config) GetRoleBindings() []*RoleBinding {
	if x != nil {
		return x.RoleBindings
	}
	return nil
}

func (x *Config) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_config_service_authz_v1_authz_proto protoreflect.FileDescriptor

var file_config_service_authz_v1_authz_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x12, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0b,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x89, 0x01, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08, 0x08,
	0x01, 0x22, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x5a, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92,
	0x01, 0x09, 0x08, 0x01, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0c,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x20, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x24, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0d, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_config_service_authz_v1_authz_proto_rawDescOnce sync.Once
	file_config_service_authz_v1_authz_proto_rawDescData = file_config_service_authz_v1_authz_proto_rawDesc
)

func file_config_service_authz_v1_authz_proto_rawDescGZIP() []byte {
	file_config_service_authz_v1_authz_proto_rawDescOnce.Do(func() {
		file_config_service_authz_v1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_authz_v1_authz_proto_rawDescData)
	})
	return file_config_service_authz_v1_authz_proto_rawDescData
}

var file_config_service_authz_v1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_service_authz_v1_authz_proto_goTypes = []interface{}{
	(*Principal)(nil),   // 0: clutch.config.service.authz.v1.Principal
	(*RoleBinding)(nil), // 1: clutch.config.service.authz.v1.RoleBinding
	(*Policy)(nil),      // 2: clutch.config.service.authz.v1.Policy
	(*Role)(nil),        // 3: clutch.config.service.authz.v1.Role
	(*Config)(nil),      // 4: clutch.config.service.authz.v1.Config
	(v1.ActionType)(0),  // 5: clutch.api.v1.ActionType
}
var file_config_service_authz_v1_authz_proto_depIdxs = []int32{
	0, // 0: clutch.config.service.authz.v1.RoleBinding.principals:type_name -> clutch.config.service.authz.v1.Principal
	5, // 1: clutch.config.service.authz.v1.Policy.action_types:type_name -> clutch.api.v1.ActionType
	2, // 2: clutch.config.service.authz.v1.Role.policies:type_name -> clutch.config.service.authz.v1.Policy
	1, // 3: clutch.config.service.authz.v1.Config.role_bindings:type_name -> clutch.config.service.authz.v1.RoleBinding
	3, // 4: clutch.config.service.authz.v1.Config.roles:type_name -> clutch.config.service.authz.v1.Role
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_service_authz_v1_authz_proto_init() }
func file_config_service_authz_v1_authz_proto_init() {
	if File_config_service_authz_v1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_authz_v1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
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
		file_config_service_authz_v1_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_config_service_authz_v1_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_config_service_authz_v1_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_config_service_authz_v1_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	file_config_service_authz_v1_authz_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Principal_User)(nil),
		(*Principal_Group)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_service_authz_v1_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_authz_v1_authz_proto_goTypes,
		DependencyIndexes: file_config_service_authz_v1_authz_proto_depIdxs,
		MessageInfos:      file_config_service_authz_v1_authz_proto_msgTypes,
	}.Build()
	File_config_service_authz_v1_authz_proto = out.File
	file_config_service_authz_v1_authz_proto_rawDesc = nil
	file_config_service_authz_v1_authz_proto_goTypes = nil
	file_config_service_authz_v1_authz_proto_depIdxs = nil
}
