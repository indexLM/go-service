// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: pkg/pb/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SaveUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	ClientId uint64 `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Sex      uint64 `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *SaveUserReq) Reset() {
	*x = SaveUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserReq) ProtoMessage() {}

func (x *SaveUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserReq.ProtoReflect.Descriptor instead.
func (*SaveUserReq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *SaveUserReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SaveUserReq) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *SaveUserReq) GetSex() uint64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

type SaveUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveUserRes) Reset() {
	*x = SaveUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserRes) ProtoMessage() {}

func (x *SaveUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserRes.ProtoReflect.Descriptor instead.
func (*SaveUserRes) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_user_proto_rawDescGZIP(), []int{1}
}

var File_pkg_pb_user_user_proto protoreflect.FileDescriptor

var file_pkg_pb_user_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x57,
	0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x32, 0x3a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_user_user_proto_rawDescOnce sync.Once
	file_pkg_pb_user_user_proto_rawDescData = file_pkg_pb_user_user_proto_rawDesc
)

func file_pkg_pb_user_user_proto_rawDescGZIP() []byte {
	file_pkg_pb_user_user_proto_rawDescOnce.Do(func() {
		file_pkg_pb_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_user_user_proto_rawDescData)
	})
	return file_pkg_pb_user_user_proto_rawDescData
}

var file_pkg_pb_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_user_user_proto_goTypes = []interface{}{
	(*SaveUserReq)(nil), // 0: user.SaveUserReq
	(*SaveUserRes)(nil), // 1: user.SaveUserRes
}
var file_pkg_pb_user_user_proto_depIdxs = []int32{
	0, // 0: user.User.SaveUser:input_type -> user.SaveUserReq
	1, // 1: user.User.SaveUser:output_type -> user.SaveUserRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_user_user_proto_init() }
func file_pkg_pb_user_user_proto_init() {
	if File_pkg_pb_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserReq); i {
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
		file_pkg_pb_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserRes); i {
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
			RawDescriptor: file_pkg_pb_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_user_user_proto_goTypes,
		DependencyIndexes: file_pkg_pb_user_user_proto_depIdxs,
		MessageInfos:      file_pkg_pb_user_user_proto_msgTypes,
	}.Build()
	File_pkg_pb_user_user_proto = out.File
	file_pkg_pb_user_user_proto_rawDesc = nil
	file_pkg_pb_user_user_proto_goTypes = nil
	file_pkg_pb_user_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	SaveUser(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserRes, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) SaveUser(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserRes, error) {
	out := new(SaveUserRes)
	err := c.cc.Invoke(ctx, "/user.User/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	SaveUser(context.Context, *SaveUserReq) (*SaveUserRes, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) SaveUser(context.Context, *SaveUserReq) (*SaveUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SaveUser(ctx, req.(*SaveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _User_SaveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/user/user.proto",
}
