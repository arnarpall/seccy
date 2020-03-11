// Code generated by protoc-gen-go. DO NOT EDIT.
// source: seccy.proto

package seccy

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

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b152d0ceac800389, []int{0}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b152d0ceac800389, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b152d0ceac800389, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b152d0ceac800389, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*SetRequest)(nil), "seccy.SetRequest")
	proto.RegisterType((*Empty)(nil), "seccy.Empty")
	proto.RegisterType((*GetRequest)(nil), "seccy.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "seccy.GetResponse")
}

func init() {
	proto.RegisterFile("seccy.proto", fileDescriptor_b152d0ceac800389)
}

var fileDescriptor_b152d0ceac800389 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4e, 0x4d, 0x4e,
	0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x4c, 0xb8, 0xb8, 0x82,
	0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c,
	0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0xa3, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52,
	0xa9, 0x24, 0xc7, 0xc5, 0xe5, 0x8e, 0x47, 0xbb, 0x92, 0x32, 0x17, 0x37, 0x58, 0xbe, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x15, 0x61, 0x1a, 0x23, 0x92, 0x69, 0x46, 0xb1, 0x5c, 0xac, 0xc1, 0x20, 0xc7,
	0x08, 0xa9, 0x71, 0x31, 0x07, 0xa7, 0x96, 0x08, 0x09, 0xea, 0x41, 0x1c, 0x8a, 0x70, 0x98, 0x14,
	0x0f, 0x54, 0x08, 0x6c, 0xab, 0x90, 0x0e, 0x17, 0xb3, 0x3b, 0x92, 0x3a, 0x84, 0x0b, 0xa4, 0x84,
	0x90, 0x85, 0x20, 0x96, 0x26, 0xb1, 0x81, 0x3d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0xd7, 0x16, 0x9c, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SeccyClient is the client API for Seccy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SeccyClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type seccyClient struct {
	cc grpc.ClientConnInterface
}

func NewSeccyClient(cc grpc.ClientConnInterface) SeccyClient {
	return &seccyClient{cc}
}

func (c *seccyClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/seccy.Seccy/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seccyClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/seccy.Seccy/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeccyServer is the server API for Seccy service.
type SeccyServer interface {
	Set(context.Context, *SetRequest) (*Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedSeccyServer can be embedded to have forward compatible implementations.
type UnimplementedSeccyServer struct {
}

func (*UnimplementedSeccyServer) Set(ctx context.Context, req *SetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedSeccyServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterSeccyServer(s *grpc.Server, srv SeccyServer) {
	s.RegisterService(&_Seccy_serviceDesc, srv)
}

func _Seccy_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeccyServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seccy.Seccy/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeccyServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seccy_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeccyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seccy.Seccy/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeccyServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Seccy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seccy.Seccy",
	HandlerType: (*SeccyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Seccy_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Seccy_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seccy.proto",
}
