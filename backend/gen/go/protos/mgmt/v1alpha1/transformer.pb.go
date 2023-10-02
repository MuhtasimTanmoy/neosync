// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: mgmt/v1alpha1/transformer.proto

package mgmtv1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTransformersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTransformersRequest) Reset() {
	*x = GetTransformersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformersRequest) ProtoMessage() {}

func (x *GetTransformersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformersRequest.ProtoReflect.Descriptor instead.
func (*GetTransformersRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_transformer_proto_rawDescGZIP(), []int{0}
}

type GetTransformersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transformers []*Transformers `protobuf:"bytes,1,rep,name=transformers,proto3" json:"transformers,omitempty"`
}

func (x *GetTransformersResponse) Reset() {
	*x = GetTransformersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformersResponse) ProtoMessage() {}

func (x *GetTransformersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformersResponse.ProtoReflect.Descriptor instead.
func (*GetTransformersResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_transformer_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransformersResponse) GetTransformers() []*Transformers {
	if x != nil {
		return x.Transformers
	}
	return nil
}

type Transformers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Config      *TransformerConfig `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Transformers) Reset() {
	*x = Transformers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transformers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transformers) ProtoMessage() {}

func (x *Transformers) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transformers.ProtoReflect.Descriptor instead.
func (*Transformers) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_transformer_proto_rawDescGZIP(), []int{2}
}

func (x *Transformers) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transformers) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transformers) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Transformers) GetConfig() *TransformerConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type TransformerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//
	//	*TransformerConfig_EmailConfig
	Config isTransformerConfig_Config `protobuf_oneof:"config"`
}

func (x *TransformerConfig) Reset() {
	*x = TransformerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformerConfig) ProtoMessage() {}

func (x *TransformerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformerConfig.ProtoReflect.Descriptor instead.
func (*TransformerConfig) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_transformer_proto_rawDescGZIP(), []int{3}
}

func (m *TransformerConfig) GetConfig() isTransformerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *TransformerConfig) GetEmailConfig() *EmailConfig {
	if x, ok := x.GetConfig().(*TransformerConfig_EmailConfig); ok {
		return x.EmailConfig
	}
	return nil
}

type isTransformerConfig_Config interface {
	isTransformerConfig_Config()
}

type TransformerConfig_EmailConfig struct {
	EmailConfig *EmailConfig `protobuf:"bytes,1,opt,name=email_config,json=emailConfig,proto3,oneof"`
}

func (*TransformerConfig_EmailConfig) isTransformerConfig_Config() {}

type EmailConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreserveDomain bool `protobuf:"varint,1,opt,name=preserve_domain,json=preserveDomain,proto3" json:"preserve_domain,omitempty"`
	PreserveLength bool `protobuf:"varint,2,opt,name=preserve_length,json=preserveLength,proto3" json:"preserve_length,omitempty"`
}

func (x *EmailConfig) Reset() {
	*x = EmailConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailConfig) ProtoMessage() {}

func (x *EmailConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_transformer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailConfig.ProtoReflect.Descriptor instead.
func (*EmailConfig) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_transformer_proto_rawDescGZIP(), []int{4}
}

func (x *EmailConfig) GetPreserveDomain() bool {
	if x != nil {
		return x.PreserveDomain
	}
	return false
}

func (x *EmailConfig) GetPreserveLength() bool {
	if x != nil {
		return x.PreserveLength
	}
	return false
}

var File_mgmt_v1alpha1_transformer_proto protoreflect.FileDescriptor

var file_mgmt_v1alpha1_transformer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73,
	0x22, 0x98, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x73, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x63, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3f, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x0d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x5f, 0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x32, 0x79, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcc, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x63, 0x6c, 0x65, 0x75, 0x73, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6e, 0x65, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x67, 0x6d, 0x74,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02,
	0x0d, 0x4d, 0x67, 0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x0d, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x19, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x67, 0x6d,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mgmt_v1alpha1_transformer_proto_rawDescOnce sync.Once
	file_mgmt_v1alpha1_transformer_proto_rawDescData = file_mgmt_v1alpha1_transformer_proto_rawDesc
)

func file_mgmt_v1alpha1_transformer_proto_rawDescGZIP() []byte {
	file_mgmt_v1alpha1_transformer_proto_rawDescOnce.Do(func() {
		file_mgmt_v1alpha1_transformer_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_v1alpha1_transformer_proto_rawDescData)
	})
	return file_mgmt_v1alpha1_transformer_proto_rawDescData
}

var file_mgmt_v1alpha1_transformer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mgmt_v1alpha1_transformer_proto_goTypes = []interface{}{
	(*GetTransformersRequest)(nil),  // 0: mgmt.v1alpha1.GetTransformersRequest
	(*GetTransformersResponse)(nil), // 1: mgmt.v1alpha1.GetTransformersResponse
	(*Transformers)(nil),            // 2: mgmt.v1alpha1.Transformers
	(*TransformerConfig)(nil),       // 3: mgmt.v1alpha1.TransformerConfig
	(*EmailConfig)(nil),             // 4: mgmt.v1alpha1.EmailConfig
}
var file_mgmt_v1alpha1_transformer_proto_depIdxs = []int32{
	2, // 0: mgmt.v1alpha1.GetTransformersResponse.transformers:type_name -> mgmt.v1alpha1.Transformers
	3, // 1: mgmt.v1alpha1.Transformers.config:type_name -> mgmt.v1alpha1.TransformerConfig
	4, // 2: mgmt.v1alpha1.TransformerConfig.email_config:type_name -> mgmt.v1alpha1.EmailConfig
	0, // 3: mgmt.v1alpha1.TransformersService.GetTransformers:input_type -> mgmt.v1alpha1.GetTransformersRequest
	1, // 4: mgmt.v1alpha1.TransformersService.GetTransformers:output_type -> mgmt.v1alpha1.GetTransformersResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mgmt_v1alpha1_transformer_proto_init() }
func file_mgmt_v1alpha1_transformer_proto_init() {
	if File_mgmt_v1alpha1_transformer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_v1alpha1_transformer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformersRequest); i {
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
		file_mgmt_v1alpha1_transformer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformersResponse); i {
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
		file_mgmt_v1alpha1_transformer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transformers); i {
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
		file_mgmt_v1alpha1_transformer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformerConfig); i {
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
		file_mgmt_v1alpha1_transformer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailConfig); i {
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
	file_mgmt_v1alpha1_transformer_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TransformerConfig_EmailConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mgmt_v1alpha1_transformer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mgmt_v1alpha1_transformer_proto_goTypes,
		DependencyIndexes: file_mgmt_v1alpha1_transformer_proto_depIdxs,
		MessageInfos:      file_mgmt_v1alpha1_transformer_proto_msgTypes,
	}.Build()
	File_mgmt_v1alpha1_transformer_proto = out.File
	file_mgmt_v1alpha1_transformer_proto_rawDesc = nil
	file_mgmt_v1alpha1_transformer_proto_goTypes = nil
	file_mgmt_v1alpha1_transformer_proto_depIdxs = nil
}
