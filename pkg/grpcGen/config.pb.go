// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: grpcGen/config.proto

package grpcGen

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DnsServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DnsServer) Reset() {
	*x = DnsServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcGen_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsServer) ProtoMessage() {}

func (x *DnsServer) ProtoReflect() protoreflect.Message {
	mi := &file_grpcGen_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsServer.ProtoReflect.Descriptor instead.
func (*DnsServer) Descriptor() ([]byte, []int) {
	return file_grpcGen_config_proto_rawDescGZIP(), []int{0}
}

func (x *DnsServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DnsServerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []string `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *DnsServerList) Reset() {
	*x = DnsServerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcGen_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsServerList) ProtoMessage() {}

func (x *DnsServerList) ProtoReflect() protoreflect.Message {
	mi := &file_grpcGen_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsServerList.ProtoReflect.Descriptor instead.
func (*DnsServerList) Descriptor() ([]byte, []int) {
	return file_grpcGen_config_proto_rawDescGZIP(), []int{1}
}

func (x *DnsServerList) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

type Hostname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Hostname) Reset() {
	*x = Hostname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcGen_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hostname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hostname) ProtoMessage() {}

func (x *Hostname) ProtoReflect() protoreflect.Message {
	mi := &file_grpcGen_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hostname.ProtoReflect.Descriptor instead.
func (*Hostname) Descriptor() ([]byte, []int) {
	return file_grpcGen_config_proto_rawDescGZIP(), []int{2}
}

func (x *Hostname) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_grpcGen_config_proto protoreflect.FileDescriptor

var file_grpcGen_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x09, 0x44, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x08,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x82, 0x02, 0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x64,
	0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x44,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x32, 0x08, 0x2f, 0x64, 0x6e,
	0x73, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x52, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2e,
	0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x32, 0x0b, 0x2f, 0x64,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x32, 0xb1, 0x01, 0x0a, 0x0f, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12,
	0x09, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a,
	0x01, 0x2a, 0x1a, 0x09, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1c, 0x5a,
	0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpcGen_config_proto_rawDescOnce sync.Once
	file_grpcGen_config_proto_rawDescData = file_grpcGen_config_proto_rawDesc
)

func file_grpcGen_config_proto_rawDescGZIP() []byte {
	file_grpcGen_config_proto_rawDescOnce.Do(func() {
		file_grpcGen_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcGen_config_proto_rawDescData)
	})
	return file_grpcGen_config_proto_rawDescData
}

var file_grpcGen_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpcGen_config_proto_goTypes = []interface{}{
	(*DnsServer)(nil),     // 0: grpcGen.DnsServer
	(*DnsServerList)(nil), // 1: grpcGen.DnsServerList
	(*Hostname)(nil),      // 2: grpcGen.Hostname
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_grpcGen_config_proto_depIdxs = []int32{
	3, // 0: grpcGen.DnsService.GetServerList:input_type -> google.protobuf.Empty
	0, // 1: grpcGen.DnsService.AddServer:input_type -> grpcGen.DnsServer
	0, // 2: grpcGen.DnsService.RemoveServer:input_type -> grpcGen.DnsServer
	3, // 3: grpcGen.HostnameService.GetHostname:input_type -> google.protobuf.Empty
	2, // 4: grpcGen.HostnameService.UpdateHostname:input_type -> grpcGen.Hostname
	1, // 5: grpcGen.DnsService.GetServerList:output_type -> grpcGen.DnsServerList
	3, // 6: grpcGen.DnsService.AddServer:output_type -> google.protobuf.Empty
	3, // 7: grpcGen.DnsService.RemoveServer:output_type -> google.protobuf.Empty
	2, // 8: grpcGen.HostnameService.GetHostname:output_type -> grpcGen.Hostname
	3, // 9: grpcGen.HostnameService.UpdateHostname:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcGen_config_proto_init() }
func file_grpcGen_config_proto_init() {
	if File_grpcGen_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcGen_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsServer); i {
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
		file_grpcGen_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsServerList); i {
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
		file_grpcGen_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hostname); i {
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
			RawDescriptor: file_grpcGen_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpcGen_config_proto_goTypes,
		DependencyIndexes: file_grpcGen_config_proto_depIdxs,
		MessageInfos:      file_grpcGen_config_proto_msgTypes,
	}.Build()
	File_grpcGen_config_proto = out.File
	file_grpcGen_config_proto_rawDesc = nil
	file_grpcGen_config_proto_goTypes = nil
	file_grpcGen_config_proto_depIdxs = nil
}
