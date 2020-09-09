// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: proto/user/user.proto

package userpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TotalTransaction int64  `protobuf:"varint,3,opt,name=totalTransaction,proto3" json:"totalTransaction,omitempty"`
	TotalRevenue     int64  `protobuf:"varint,4,opt,name=totalRevenue,proto3" json:"totalRevenue,omitempty"`
}

func (x *UserBrief) Reset() {
	*x = UserBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBrief) ProtoMessage() {}

func (x *UserBrief) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBrief.ProtoReflect.Descriptor instead.
func (*UserBrief) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserBrief) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserBrief) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserBrief) GetTotalTransaction() int64 {
	if x != nil {
		return x.TotalTransaction
	}
	return 0
}

func (x *UserBrief) GetTotalRevenue() int64 {
	if x != nil {
		return x.TotalRevenue
	}
	return 0
}

type GetUserBriefByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserBriefByIDRequest) Reset() {
	*x = GetUserBriefByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBriefByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBriefByIDRequest) ProtoMessage() {}

func (x *GetUserBriefByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBriefByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserBriefByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserBriefByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserBriefByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBrief *UserBrief `protobuf:"bytes,1,opt,name=userBrief,proto3" json:"userBrief,omitempty"`
}

func (x *GetUserBriefByIDResponse) Reset() {
	*x = GetUserBriefByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBriefByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBriefByIDResponse) ProtoMessage() {}

func (x *GetUserBriefByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBriefByIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserBriefByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserBriefByIDResponse) GetUserBrief() *UserBrief {
	if x != nil {
		return x.UserBrief
	}
	return nil
}

var File_proto_user_user_proto protoreflect.FileDescriptor

var file_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x7f, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x22, 0x29,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x32, 0x62, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x72,
	0x69, 0x65, 0x66, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData = file_proto_user_user_proto_rawDesc
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_user_proto_rawDescData)
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_user_user_proto_goTypes = []interface{}{
	(*UserBrief)(nil),                // 0: user.UserBrief
	(*GetUserBriefByIDRequest)(nil),  // 1: user.GetUserBriefByIDRequest
	(*GetUserBriefByIDResponse)(nil), // 2: user.GetUserBriefByIDResponse
}
var file_proto_user_user_proto_depIdxs = []int32{
	0, // 0: user.GetUserBriefByIDResponse.userBrief:type_name -> user.UserBrief
	1, // 1: user.UserService.GetUserBriefByID:input_type -> user.GetUserBriefByIDRequest
	2, // 2: user.UserService.GetUserBriefByID:output_type -> user.GetUserBriefByIDResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBrief); i {
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
		file_proto_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBriefByIDRequest); i {
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
		file_proto_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBriefByIDResponse); i {
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
			RawDescriptor: file_proto_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_rawDesc = nil
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserBriefByID(ctx context.Context, in *GetUserBriefByIDRequest, opts ...grpc.CallOption) (*GetUserBriefByIDResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserBriefByID(ctx context.Context, in *GetUserBriefByIDRequest, opts ...grpc.CallOption) (*GetUserBriefByIDResponse, error) {
	out := new(GetUserBriefByIDResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserBriefByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserBriefByID(context.Context, *GetUserBriefByIDRequest) (*GetUserBriefByIDResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserBriefByID(context.Context, *GetUserBriefByIDRequest) (*GetUserBriefByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBriefByID not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserBriefByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBriefByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserBriefByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserBriefByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserBriefByID(ctx, req.(*GetUserBriefByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserBriefByID",
			Handler:    _UserService_GetUserBriefByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
