// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc/sys/sys.proto

package sys

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

// user start---------------------
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	CurrentAuthority string `protobuf:"bytes,2,opt,name=currentAuthority,proto3" json:"currentAuthority,omitempty"`
	Id               int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	UserName         string `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
	AccessToken      string `protobuf:"bytes,5,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	AccessExpire     int64  `protobuf:"varint,6,opt,name=AccessExpire,proto3" json:"AccessExpire,omitempty"`
	RefreshAfter     int64  `protobuf:"varint,7,opt,name=RefreshAfter,proto3" json:"RefreshAfter,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LoginResp) GetCurrentAuthority() string {
	if x != nil {
		return x.CurrentAuthority
	}
	return ""
}

func (x *LoginResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginResp) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *LoginResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

var File_rpc_sys_sys_proto protoreflect.FileDescriptor

var file_rpc_sys_sys_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x79, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xe5, 0x01, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x32, 0x2d, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_sys_proto_rawDescOnce sync.Once
	file_rpc_sys_sys_proto_rawDescData = file_rpc_sys_sys_proto_rawDesc
)

func file_rpc_sys_sys_proto_rawDescGZIP() []byte {
	file_rpc_sys_sys_proto_rawDescOnce.Do(func() {
		file_rpc_sys_sys_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_sys_proto_rawDescData)
	})
	return file_rpc_sys_sys_proto_rawDescData
}

var file_rpc_sys_sys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_sys_sys_proto_goTypes = []interface{}{
	(*LoginReq)(nil),  // 0: sys.LoginReq
	(*LoginResp)(nil), // 1: sys.loginResp
}
var file_rpc_sys_sys_proto_depIdxs = []int32{
	0, // 0: sys.Sys.Login:input_type -> sys.LoginReq
	1, // 1: sys.Sys.Login:output_type -> sys.loginResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_sys_sys_proto_init() }
func file_rpc_sys_sys_proto_init() {
	if File_rpc_sys_sys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_sys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_rpc_sys_sys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
			RawDescriptor: file_rpc_sys_sys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sys_sys_proto_goTypes,
		DependencyIndexes: file_rpc_sys_sys_proto_depIdxs,
		MessageInfos:      file_rpc_sys_sys_proto_msgTypes,
	}.Build()
	File_rpc_sys_sys_proto = out.File
	file_rpc_sys_sys_proto_rawDesc = nil
	file_rpc_sys_sys_proto_goTypes = nil
	file_rpc_sys_sys_proto_depIdxs = nil
}
