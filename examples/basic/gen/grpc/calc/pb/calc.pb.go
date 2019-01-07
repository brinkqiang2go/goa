// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	// Left operand
	A int32 `protobuf:"zigzag32,1,opt,name=a,proto3" json:"a,omitempty"`
	// Right operand
	B                    int32    `protobuf:"zigzag32,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_df3b5acbce7d6c22, []int{0}
}
func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (dst *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(dst, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Field                int32    `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_df3b5acbce7d6c22, []int{1}
}
func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (dst *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(dst, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetField() int32 {
	if m != nil {
		return m.Field
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "pb.AddResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	// Add implements add.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/pb.calc/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	// Add implements add.
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.calc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calc_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor_calc_df3b5acbce7d6c22) }

var fileDescriptor_calc_df3b5acbce7d6c22 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe0, 0xe2, 0x72, 0x4c,
	0x49, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54,
	0x60, 0xd4, 0x10, 0x0c, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x65,
	0x2e, 0x6e, 0xb0, 0xca, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb4, 0xcc,
	0xd4, 0x9c, 0x14, 0xa8, 0x72, 0x08, 0xc7, 0x48, 0x8f, 0x8b, 0x05, 0x64, 0x81, 0x90, 0x1a, 0x17,
	0xb3, 0x63, 0x4a, 0x8a, 0x10, 0x9f, 0x5e, 0x41, 0x92, 0x1e, 0xc2, 0x7c, 0x29, 0x7e, 0x38, 0x1f,
	0x62, 0x4a, 0x12, 0x1b, 0xd8, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xf4, 0xa6,
	0xc1, 0x97, 0x00, 0x00, 0x00,
}
