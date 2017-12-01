// Code generated by protoc-gen-go.
// source: books/books.proto
// DO NOT EDIT!

/*
Package books is a generated protocol buffer package.

It is generated from these files:
	books/books.proto

It has these top-level messages:
	Empty
	Book
	BookList
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

type Book struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Book) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type BookList struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *BookList) Reset()                    { *m = BookList{} }
func (m *BookList) String() string            { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()               {}
func (*BookList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BookList) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "books.Empty")
	proto.RegisterType((*Book)(nil), "books.Book")
	proto.RegisterType((*BookList)(nil), "books.BookList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	// New function
	Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := grpc.Invoke(ctx, "/books.BookService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Insert(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/books.BookService/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	List(context.Context, *Empty) (*BookList, error)
	// New function
	Insert(context.Context, *Book) (*Empty, error)
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

func _BookService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.BookService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Insert(ctx, req.(*Book))
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
		{
			MethodName: "Insert",
			Handler:    _BookService_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books/books.proto",
}

func init() { proto.RegisterFile("books/books.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xca, 0xcf, 0xcf,
	0x2e, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x12, 0x3b,
	0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x92, 0x0b, 0x17, 0x8b, 0x53, 0x7e, 0x7e, 0xb6, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0x90,
	0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84,
	0x23, 0x24, 0xc6, 0xc5, 0x96, 0x58, 0x5a, 0x92, 0x91, 0x5f, 0x24, 0xc1, 0x0c, 0x16, 0x86, 0xf2,
	0x94, 0x74, 0xb9, 0x38, 0x40, 0xa6, 0xf8, 0x64, 0x16, 0x97, 0x08, 0x29, 0x72, 0x41, 0xec, 0x90,
	0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd6, 0x83, 0x58, 0x0f, 0x92, 0x0f, 0x82, 0xc8, 0x18,
	0xc5, 0x72, 0x71, 0x83, 0xb8, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xea, 0x5c, 0x2c,
	0x60, 0x9d, 0x3c, 0x50, 0xa5, 0x60, 0x97, 0x49, 0xf1, 0x23, 0x69, 0x04, 0x49, 0x2b, 0x31, 0x08,
	0xa9, 0x72, 0xb1, 0x79, 0xe6, 0x15, 0xa7, 0x16, 0x95, 0x08, 0x21, 0x9b, 0x2a, 0x85, 0xa2, 0x4f,
	0x89, 0x21, 0x89, 0x0d, 0xec, 0x55, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe7, 0xd0,
	0xea, 0xff, 0x00, 0x00, 0x00,
}
