// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	session.proto

It has these top-level messages:
	Empty
	Message
*/
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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Message)(nil), "pb.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Session service

type SessionClient interface {
	Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) Ping(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/pb.Session/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Session service

type SessionServer interface {
	Ping(context.Context, *Message) (*Message, error)
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Session/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).Ping(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Session_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}

func init() { proto.RegisterFile("session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x62, 0xe7,
	0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0x52, 0xe6, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x8d, 0x74, 0xb9, 0xd8, 0x83, 0x21, 0x46, 0x08, 0x29, 0x71, 0xb1,
	0x04, 0x64, 0xe6, 0xa5, 0x0b, 0x71, 0xeb, 0x15, 0x24, 0xe9, 0x41, 0x75, 0x4a, 0x21, 0x73, 0x94,
	0x18, 0x92, 0xd8, 0xc0, 0xf6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0x68, 0x52, 0xc3,
	0x78, 0x00, 0x00, 0x00,
}
