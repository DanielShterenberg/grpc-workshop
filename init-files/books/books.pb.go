// Code generated by protoc-gen-go.
// source: books/books.proto
// DO NOT EDIT!

/*
Package books is a generated protocol buffer package.

It is generated from these files:
	books/books.proto

It has these top-level messages:
	Empty
*/
package books

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

func init() {
	proto.RegisterType((*Empty)(nil), "books.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	List(context.Context, *Empty) (*Empty, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BookService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books/books.proto",
}

func init() { proto.RegisterFile("books/books.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xca, 0xcf, 0xcf,
	0x2e, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x12, 0x3b,
	0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x91, 0x31, 0x17, 0xb7, 0x53, 0x7e, 0x7e, 0x76, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x0a, 0x17, 0x8b, 0x4f, 0x66, 0x71, 0x89, 0x10, 0x8f,
	0x1e, 0x44, 0x13, 0x58, 0x91, 0x14, 0x0a, 0x4f, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0x96, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x67, 0xe4, 0xb5, 0x60, 0x00, 0x00, 0x00,
}