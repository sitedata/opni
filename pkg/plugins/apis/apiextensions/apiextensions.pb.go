// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	ragù          v0.2.3
// source: pkg/plugins/apis/apiextensions/apiextensions.proto

package apiextensions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CertConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ca       string `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	CaData   string `protobuf:"bytes,2,opt,name=caData,proto3" json:"caData,omitempty"`
	Cert     string `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	CertData string `protobuf:"bytes,4,opt,name=certData,proto3" json:"certData,omitempty"`
	Key      string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	KeyData  string `protobuf:"bytes,6,opt,name=keyData,proto3" json:"keyData,omitempty"`
}

func (x *CertConfig) Reset() {
	*x = CertConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertConfig) ProtoMessage() {}

func (x *CertConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertConfig.ProtoReflect.Descriptor instead.
func (*CertConfig) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{0}
}

func (x *CertConfig) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

func (x *CertConfig) GetCaData() string {
	if x != nil {
		return x.CaData
	}
	return ""
}

func (x *CertConfig) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *CertConfig) GetCertData() string {
	if x != nil {
		return x.CertData
	}
	return ""
}

func (x *CertConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CertConfig) GetKeyData() string {
	if x != nil {
		return x.KeyData
	}
	return ""
}

type GatewayAPIExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAddr     string   `protobuf:"bytes,1,opt,name=HttpAddr,proto3" json:"HttpAddr,omitempty"`
	PathPrefixes []string `protobuf:"bytes,2,rep,name=pathPrefixes,proto3" json:"pathPrefixes,omitempty"`
}

func (x *GatewayAPIExtensionConfig) Reset() {
	*x = GatewayAPIExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayAPIExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayAPIExtensionConfig) ProtoMessage() {}

func (x *GatewayAPIExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayAPIExtensionConfig.ProtoReflect.Descriptor instead.
func (*GatewayAPIExtensionConfig) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayAPIExtensionConfig) GetHttpAddr() string {
	if x != nil {
		return x.HttpAddr
	}
	return ""
}

func (x *GatewayAPIExtensionConfig) GetPathPrefixes() []string {
	if x != nil {
		return x.PathPrefixes
	}
	return nil
}

var File_pkg_plugins_apis_apiextensions_apiextensions_proto protoreflect.FileDescriptor

var file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a,
	0x0a, 0x43, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0c, 0x0a, 0x02, 0x63,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x10, 0x0a, 0x06, 0x63, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x0e, 0x0a, 0x04, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x12, 0x0a, 0x08, 0x63,
	0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12,
	0x0d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x11,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x00, 0x3a, 0x00, 0x22, 0x49, 0x0a, 0x19, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x50,
	0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x08, 0x48, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x00, 0x12, 0x16, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00, 0x32, 0x6f,
	0x0a, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x00, 0x32,
	0x6f, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x28,
	0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x00,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce sync.Once
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc
)

func file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescGZIP() []byte {
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescOnce.Do(func() {
		file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData)
	})
	return file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDescData
}

var file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = []interface{}{
	(*CertConfig)(nil),                          // 0: apiextensions.CertConfig
	(*GatewayAPIExtensionConfig)(nil),           // 1: apiextensions.GatewayAPIExtensionConfig
	(*emptypb.Empty)(nil),                       // 2: google.protobuf.Empty
	(*descriptorpb.ServiceDescriptorProto)(nil), // 3: google.protobuf.ServiceDescriptorProto
}
var file_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = []int32{
	2, // 0: apiextensions.ManagementAPIExtension.Descriptor:input_type -> google.protobuf.Empty
	0, // 1: apiextensions.GatewayAPIExtension.Configure:input_type -> apiextensions.CertConfig
	3, // 2: apiextensions.ManagementAPIExtension.Descriptor:output_type -> google.protobuf.ServiceDescriptorProto
	1, // 3: apiextensions.GatewayAPIExtension.Configure:output_type -> apiextensions.GatewayAPIExtensionConfig
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_plugins_apis_apiextensions_apiextensions_proto_init() }
func file_pkg_plugins_apis_apiextensions_apiextensions_proto_init() {
	if File_pkg_plugins_apis_apiextensions_apiextensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertConfig); i {
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
		file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayAPIExtensionConfig); i {
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
			RawDescriptor: file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes,
		DependencyIndexes: file_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs,
		MessageInfos:      file_pkg_plugins_apis_apiextensions_apiextensions_proto_msgTypes,
	}.Build()
	File_pkg_plugins_apis_apiextensions_apiextensions_proto = out.File
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_rawDesc = nil
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_goTypes = nil
	file_pkg_plugins_apis_apiextensions_apiextensions_proto_depIdxs = nil
}
