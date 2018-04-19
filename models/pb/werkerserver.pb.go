// Code generated by protoc-gen-go. DO NOT EDIT.
// source: werkerserver.proto

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

type Request struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Request) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Response struct {
	OutputLine string `protobuf:"bytes,1,opt,name=outputLine" json:"outputLine,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Response) GetOutputLine() string {
	if m != nil {
		return m.OutputLine
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "models.Request")
	proto.RegisterType((*Response)(nil), "models.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Build service

type BuildClient interface {
	BuildInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_BuildInfoClient, error)
	KillHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_KillHashClient, error)
}

type buildClient struct {
	cc *grpc.ClientConn
}

func NewBuildClient(cc *grpc.ClientConn) BuildClient {
	return &buildClient{cc}
}

func (c *buildClient) BuildInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_BuildInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[0], c.cc, "/models.Build/BuildInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildBuildInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_BuildInfoClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type buildBuildInfoClient struct {
	grpc.ClientStream
}

func (x *buildBuildInfoClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildClient) KillHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_KillHashClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[1], c.cc, "/models.Build/KillHash", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildKillHashClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_KillHashClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type buildKillHashClient struct {
	grpc.ClientStream
}

func (x *buildKillHashClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Build service

type BuildServer interface {
	BuildInfo(*Request, Build_BuildInfoServer) error
	KillHash(*Request, Build_KillHashServer) error
}

func RegisterBuildServer(s *grpc.Server, srv BuildServer) {
	s.RegisterService(&_Build_serviceDesc, srv)
}

func _Build_BuildInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).BuildInfo(m, &buildBuildInfoServer{stream})
}

type Build_BuildInfoServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type buildBuildInfoServer struct {
	grpc.ServerStream
}

func (x *buildBuildInfoServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Build_KillHash_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).KillHash(m, &buildKillHashServer{stream})
}

type Build_KillHashServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type buildKillHashServer struct {
	grpc.ServerStream
}

func (x *buildKillHashServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Build_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Build",
	HandlerType: (*BuildServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BuildInfo",
			Handler:       _Build_BuildInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KillHash",
			Handler:       _Build_KillHash_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "werkerserver.proto",
}

func init() { proto.RegisterFile("werkerserver.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4f, 0x2d, 0xca,
	0x4e, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xcb, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x56, 0x92, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x95, 0xb4, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe4,
	0xb8, 0xb8, 0xf2, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0x7c, 0x32, 0xf3, 0x52, 0xa1, 0xaa, 0x90, 0x44,
	0x8c, 0xf2, 0xb8, 0x58, 0x9d, 0x4a, 0x33, 0x73, 0x52, 0x84, 0x8c, 0xb8, 0x38, 0xc1, 0x0c, 0xcf,
	0xbc, 0xb4, 0x7c, 0x21, 0x7e, 0x3d, 0x88, 0x4d, 0x7a, 0x50, 0x6b, 0xa4, 0x04, 0x10, 0x02, 0x10,
	0x83, 0x95, 0x18, 0x0c, 0x18, 0x85, 0x0c, 0xb9, 0x38, 0xbc, 0x33, 0x73, 0x72, 0x3c, 0x12, 0x8b,
	0x33, 0x88, 0xd4, 0xe2, 0xc4, 0x12, 0xc5, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0x8f, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x12, 0xfa, 0xad, 0xe9, 0xe5, 0x00, 0x00, 0x00,
}