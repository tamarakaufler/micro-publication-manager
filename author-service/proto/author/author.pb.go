// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/author/author.proto

package author

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

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Books                []*Book  `protobuf:"bytes,4,rep,name=books" json:"books,omitempty"`
	CountryId            string   `protobuf:"bytes,5,opt,name=country_id,json=countryId" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_author_8c8bb50646d09e05, []int{0}
}
func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (dst *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(dst, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Author) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Author) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

func (m *Author) GetCountryId() string {
	if m != nil {
		return m.CountryId
	}
	return ""
}

type Book struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	Language             string   `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_author_8c8bb50646d09e05, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Book) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Author               *Author  `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_author_8c8bb50646d09e05, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*Author)(nil), "author.Author")
	proto.RegisterType((*Book)(nil), "author.Book")
	proto.RegisterType((*Response)(nil), "author.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthorService service

type AuthorServiceClient interface {
	CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Response, error)
}

type authorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorServiceClient(cc *grpc.ClientConn) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/author.AuthorService/CreateAuthor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorService service

type AuthorServiceServer interface {
	CreateAuthor(context.Context, *Author) (*Response, error)
}

func RegisterAuthorServiceServer(s *grpc.Server, srv AuthorServiceServer) {
	s.RegisterService(&_AuthorService_serviceDesc, srv)
}

func _AuthorService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author.AuthorService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "author.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _AuthorService_CreateAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/author/author.proto",
}

func init() { proto.RegisterFile("proto/author/author.proto", fileDescriptor_author_8c8bb50646d09e05) }

var fileDescriptor_author_8c8bb50646d09e05 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4e, 0x83, 0x30,
	0x14, 0x86, 0x65, 0x03, 0x84, 0xc3, 0x5c, 0x4c, 0xe3, 0x45, 0xdd, 0x62, 0x42, 0xb8, 0x30, 0xbb,
	0x9a, 0x09, 0x3e, 0x81, 0xee, 0x6a, 0x89, 0xf1, 0x02, 0x1f, 0x60, 0xe9, 0x68, 0x9d, 0xcd, 0x18,
	0x25, 0xa5, 0x98, 0xf8, 0x04, 0x3e, 0x82, 0xaf, 0x6b, 0xe8, 0x69, 0x93, 0x19, 0xaf, 0xe0, 0xff,
	0x7e, 0xfa, 0xf3, 0xf7, 0x1c, 0xb8, 0xed, 0xb4, 0x32, 0xea, 0x81, 0x0d, 0xe6, 0x43, 0x69, 0xf7,
	0x58, 0x5b, 0x46, 0x62, 0x54, 0xc5, 0x4f, 0x00, 0xf1, 0x93, 0x7d, 0x25, 0x73, 0x98, 0x48, 0x4e,
	0x83, 0x3c, 0x58, 0xa5, 0xd5, 0x44, 0x72, 0x72, 0x07, 0xf0, 0x2e, 0x75, 0x6f, 0x76, 0x2d, 0x3b,
	0x09, 0x3a, 0xb1, 0x3c, 0xb5, 0xe4, 0x95, 0x9d, 0x04, 0x59, 0x42, 0xda, 0x30, 0xef, 0x4e, 0xad,
	0x9b, 0x8c, 0xc0, 0x9a, 0x05, 0x44, 0x7b, 0xa5, 0x8e, 0x3d, 0x0d, 0xf3, 0xe9, 0x2a, 0x2b, 0x67,
	0x6b, 0xf7, 0xf3, 0x67, 0xa5, 0x8e, 0x15, 0x5a, 0x63, 0x7e, 0xad, 0x86, 0xd6, 0xe8, 0xaf, 0x9d,
	0xe4, 0x34, 0xc2, 0x7c, 0x47, 0xb6, 0xbc, 0xf8, 0x0e, 0x20, 0x1c, 0x3f, 0xff, 0xd7, 0x6b, 0x09,
	0x29, 0xa6, 0x8d, 0xc7, 0xb0, 0x56, 0x82, 0x60, 0xcb, 0xc9, 0x02, 0x92, 0x86, 0xb5, 0x87, 0x81,
	0x1d, 0xce, 0x4a, 0xa1, 0x26, 0x37, 0x10, 0x19, 0x69, 0x1a, 0x41, 0x43, 0x6b, 0xa0, 0x20, 0x39,
	0x64, 0x5c, 0xf4, 0xb5, 0x96, 0x9d, 0x91, 0xaa, 0x75, 0x3d, 0xce, 0x51, 0xf1, 0x02, 0x49, 0x25,
	0xfa, 0x4e, 0xb5, 0xbd, 0x20, 0x14, 0x2e, 0x6b, 0x2d, 0x98, 0x11, 0xd8, 0x28, 0xa9, 0xbc, 0x24,
	0xf7, 0xe0, 0x66, 0x6a, 0x3b, 0x65, 0xe5, 0xdc, 0xdf, 0x19, 0xc7, 0x5b, 0x39, 0xb7, 0xdc, 0xc0,
	0x15, 0x92, 0x37, 0xa1, 0x3f, 0x65, 0x2d, 0x48, 0x09, 0xb3, 0x8d, 0xcd, 0xf0, 0x7b, 0xf8, 0x7b,
	0x70, 0x71, 0xed, 0xb5, 0x2f, 0x51, 0x5c, 0xec, 0x63, 0xbb, 0xc5, 0xc7, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x1f, 0x07, 0xe9, 0xe2, 0x01, 0x00, 0x00,
}
