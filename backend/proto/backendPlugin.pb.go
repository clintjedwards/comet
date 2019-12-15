// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/proto/backendPlugin.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateMachineRequest_Size int32

const (
	CreateMachineRequest_NONE     CreateMachineRequest_Size = 0
	CreateMachineRequest_TINY     CreateMachineRequest_Size = 1
	CreateMachineRequest_SMALL    CreateMachineRequest_Size = 2
	CreateMachineRequest_MEDIUM   CreateMachineRequest_Size = 3
	CreateMachineRequest_LARGE    CreateMachineRequest_Size = 4
	CreateMachineRequest_ENORMOUS CreateMachineRequest_Size = 5
)

var CreateMachineRequest_Size_name = map[int32]string{
	0: "NONE",
	1: "TINY",
	2: "SMALL",
	3: "MEDIUM",
	4: "LARGE",
	5: "ENORMOUS",
}

var CreateMachineRequest_Size_value = map[string]int32{
	"NONE":     0,
	"TINY":     1,
	"SMALL":    2,
	"MEDIUM":   3,
	"LARGE":    4,
	"ENORMOUS": 5,
}

func (x CreateMachineRequest_Size) String() string {
	return proto.EnumName(CreateMachineRequest_Size_name, int32(x))
}

func (CreateMachineRequest_Size) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_16e734a482fe8dbd, []int{0, 0}
}

type CreateMachineRequest struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size                 CreateMachineRequest_Size `protobuf:"varint,3,opt,name=size,proto3,enum=proto.CreateMachineRequest_Size" json:"size,omitempty"`
	Metadata             map[string]string         `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateMachineRequest) Reset()         { *m = CreateMachineRequest{} }
func (m *CreateMachineRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMachineRequest) ProtoMessage()    {}
func (*CreateMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e734a482fe8dbd, []int{0}
}

func (m *CreateMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineRequest.Unmarshal(m, b)
}
func (m *CreateMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineRequest.Marshal(b, m, deterministic)
}
func (m *CreateMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineRequest.Merge(m, src)
}
func (m *CreateMachineRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMachineRequest.Size(m)
}
func (m *CreateMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineRequest proto.InternalMessageInfo

func (m *CreateMachineRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMachineRequest) GetSize() CreateMachineRequest_Size {
	if m != nil {
		return m.Size
	}
	return CreateMachineRequest_NONE
}

func (m *CreateMachineRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateMachineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMachineResponse) Reset()         { *m = CreateMachineResponse{} }
func (m *CreateMachineResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMachineResponse) ProtoMessage()    {}
func (*CreateMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e734a482fe8dbd, []int{1}
}

func (m *CreateMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineResponse.Unmarshal(m, b)
}
func (m *CreateMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineResponse.Marshal(b, m, deterministic)
}
func (m *CreateMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineResponse.Merge(m, src)
}
func (m *CreateMachineResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMachineResponse.Size(m)
}
func (m *CreateMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineResponse proto.InternalMessageInfo

type GetPluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoRequest) Reset()         { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()    {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e734a482fe8dbd, []int{2}
}

func (m *GetPluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoRequest.Unmarshal(m, b)
}
func (m *GetPluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoRequest.Merge(m, src)
}
func (m *GetPluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoRequest.Size(m)
}
func (m *GetPluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoRequest proto.InternalMessageInfo

type GetPluginInfoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoResponse) Reset()         { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16e734a482fe8dbd, []int{3}
}

func (m *GetPluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoResponse.Unmarshal(m, b)
}
func (m *GetPluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoResponse.Merge(m, src)
}
func (m *GetPluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoResponse.Size(m)
}
func (m *GetPluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("proto.CreateMachineRequest_Size", CreateMachineRequest_Size_name, CreateMachineRequest_Size_value)
	proto.RegisterType((*CreateMachineRequest)(nil), "proto.CreateMachineRequest")
	proto.RegisterMapType((map[string]string)(nil), "proto.CreateMachineRequest.MetadataEntry")
	proto.RegisterType((*CreateMachineResponse)(nil), "proto.CreateMachineResponse")
	proto.RegisterType((*GetPluginInfoRequest)(nil), "proto.GetPluginInfoRequest")
	proto.RegisterType((*GetPluginInfoResponse)(nil), "proto.GetPluginInfoResponse")
}

func init() { proto.RegisterFile("backend/proto/backendPlugin.proto", fileDescriptor_16e734a482fe8dbd) }

var fileDescriptor_16e734a482fe8dbd = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4e, 0xfa, 0x50,
	0x10, 0xc5, 0xff, 0xfd, 0x22, 0x30, 0x7f, 0x21, 0xcd, 0x04, 0xb5, 0x41, 0x17, 0xb5, 0x2b, 0xdc,
	0x40, 0x82, 0x2e, 0x8c, 0xae, 0x50, 0x1b, 0x02, 0xa1, 0xc5, 0x14, 0x59, 0xb8, 0xbc, 0xd0, 0x51,
	0x1b, 0xe0, 0x16, 0xe9, 0xc5, 0x04, 0x5e, 0xc8, 0xd7, 0xf0, 0xd1, 0x4c, 0x3f, 0x24, 0x42, 0x1a,
	0x56, 0x9d, 0x39, 0x33, 0xe7, 0xd7, 0xd3, 0x29, 0x5c, 0x8c, 0xd9, 0x64, 0x4a, 0xdc, 0x6f, 0x2e,
	0x96, 0xa1, 0x08, 0x9b, 0x59, 0xf7, 0x34, 0x5b, 0xbd, 0x05, 0xbc, 0x91, 0x68, 0xa8, 0x25, 0x0f,
	0xeb, 0x5b, 0x86, 0xea, 0xc3, 0x92, 0x98, 0x20, 0x87, 0x4d, 0xde, 0x03, 0x4e, 0x1e, 0x7d, 0xac,
	0x28, 0x12, 0x58, 0x01, 0x39, 0xf0, 0x0d, 0xc9, 0x94, 0xea, 0x25, 0x4f, 0x0e, 0x7c, 0x44, 0x50,
	0x39, 0x9b, 0x93, 0x21, 0x27, 0x4a, 0x52, 0xe3, 0x35, 0xa8, 0x51, 0xb0, 0x21, 0x43, 0x31, 0xa5,
	0x7a, 0xa5, 0x65, 0xa6, 0xe4, 0x46, 0x1e, 0xae, 0x31, 0x0c, 0x36, 0xe4, 0x25, 0xdb, 0x68, 0x43,
	0x71, 0x4e, 0x82, 0xf9, 0x4c, 0x30, 0x43, 0x35, 0x95, 0xfa, 0xff, 0xd6, 0xe5, 0x21, 0xa7, 0x93,
	0xed, 0xda, 0x5c, 0x2c, 0xd7, 0xde, 0xd6, 0x5a, 0xbb, 0x83, 0xf2, 0xce, 0x08, 0x75, 0x50, 0xa6,
	0xb4, 0xce, 0x22, 0xc7, 0x25, 0x56, 0x41, 0xfb, 0x64, 0xb3, 0xd5, 0x6f, 0xe8, 0xb4, 0xb9, 0x95,
	0x6f, 0x24, 0xab, 0x07, 0x6a, 0x9c, 0x08, 0x8b, 0xa0, 0xba, 0x03, 0xd7, 0xd6, 0xff, 0xc5, 0xd5,
	0x73, 0xd7, 0x7d, 0xd1, 0x25, 0x2c, 0x81, 0x36, 0x74, 0xda, 0xfd, 0xbe, 0x2e, 0x23, 0x40, 0xc1,
	0xb1, 0x1f, 0xbb, 0x23, 0x47, 0x57, 0x62, 0xb9, 0xdf, 0xf6, 0x3a, 0xb6, 0xae, 0xe2, 0x11, 0x14,
	0x6d, 0x77, 0xe0, 0x39, 0x83, 0xd1, 0x50, 0xd7, 0xac, 0x53, 0x38, 0xde, 0x0b, 0x1e, 0x2d, 0x42,
	0x1e, 0x91, 0x75, 0x02, 0xd5, 0x0e, 0x89, 0xf4, 0xea, 0x5d, 0xfe, 0x1a, 0x66, 0x5f, 0x14, 0x1b,
	0xf6, 0xf4, 0xd4, 0xd0, 0xfa, 0x92, 0xa0, 0x7c, 0xff, 0xf7, 0x5f, 0x61, 0x0f, 0xca, 0x3b, 0x6c,
	0x3c, 0x3b, 0x70, 0xaa, 0xda, 0x79, 0xfe, 0x30, 0xa5, 0xc7, 0xac, 0x9d, 0xd7, 0x6e, 0x59, 0x79,
	0x21, 0xb7, 0xac, 0xdc, 0xa4, 0xe3, 0x42, 0x32, 0xbc, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xec, 0x81, 0x18, 0x69, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendPluginClient is the client API for BackendPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendPluginClient interface {
	CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error)
	GetPluginInfo(ctx context.Context, in *GetPluginInfoRequest, opts ...grpc.CallOption) (*GetPluginInfoResponse, error)
}

type backendPluginClient struct {
	cc *grpc.ClientConn
}

func NewBackendPluginClient(cc *grpc.ClientConn) BackendPluginClient {
	return &backendPluginClient{cc}
}

func (c *backendPluginClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error) {
	out := new(CreateMachineResponse)
	err := c.cc.Invoke(ctx, "/proto.BackendPlugin/CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendPluginClient) GetPluginInfo(ctx context.Context, in *GetPluginInfoRequest, opts ...grpc.CallOption) (*GetPluginInfoResponse, error) {
	out := new(GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.BackendPlugin/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendPluginServer is the server API for BackendPlugin service.
type BackendPluginServer interface {
	CreateMachine(context.Context, *CreateMachineRequest) (*CreateMachineResponse, error)
	GetPluginInfo(context.Context, *GetPluginInfoRequest) (*GetPluginInfoResponse, error)
}

// UnimplementedBackendPluginServer can be embedded to have forward compatible implementations.
type UnimplementedBackendPluginServer struct {
}

func (*UnimplementedBackendPluginServer) CreateMachine(ctx context.Context, req *CreateMachineRequest) (*CreateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}
func (*UnimplementedBackendPluginServer) GetPluginInfo(ctx context.Context, req *GetPluginInfoRequest) (*GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterBackendPluginServer(s *grpc.Server, srv BackendPluginServer) {
	s.RegisterService(&_BackendPlugin_serviceDesc, srv)
}

func _BackendPlugin_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendPluginServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BackendPlugin/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendPluginServer).CreateMachine(ctx, req.(*CreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendPlugin_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendPluginServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BackendPlugin/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendPluginServer).GetPluginInfo(ctx, req.(*GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackendPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BackendPlugin",
	HandlerType: (*BackendPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMachine",
			Handler:    _BackendPlugin_CreateMachine_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _BackendPlugin_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/proto/backendPlugin.proto",
}
