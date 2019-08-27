// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/book.proto

package book

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

type GetBookTitle struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookTitle) Reset()         { *m = GetBookTitle{} }
func (m *GetBookTitle) String() string { return proto.CompactTextString(m) }
func (*GetBookTitle) ProtoMessage()    {}
func (*GetBookTitle) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff6d7530e9d9699, []int{0}
}

func (m *GetBookTitle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookTitle.Unmarshal(m, b)
}
func (m *GetBookTitle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookTitle.Marshal(b, m, deterministic)
}
func (m *GetBookTitle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookTitle.Merge(m, src)
}
func (m *GetBookTitle) XXX_Size() int {
	return xxx_messageInfo_GetBookTitle.Size(m)
}
func (m *GetBookTitle) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookTitle.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookTitle proto.InternalMessageInfo

func (m *GetBookTitle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type BookResponse struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookResponse) Reset()         { *m = BookResponse{} }
func (m *BookResponse) String() string { return proto.CompactTextString(m) }
func (*BookResponse) ProtoMessage()    {}
func (*BookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff6d7530e9d9699, []int{1}
}

func (m *BookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookResponse.Unmarshal(m, b)
}
func (m *BookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookResponse.Marshal(b, m, deterministic)
}
func (m *BookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookResponse.Merge(m, src)
}
func (m *BookResponse) XXX_Size() int {
	return xxx_messageInfo_BookResponse.Size(m)
}
func (m *BookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookResponse proto.InternalMessageInfo

func (m *BookResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBookTitle)(nil), "GetBookTitle")
	proto.RegisterType((*BookResponse)(nil), "BookResponse")
}

func init() { proto.RegisterFile("proto/book.proto", fileDescriptor_eff6d7530e9d9699) }

var fileDescriptor_eff6d7530e9d9699 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xca, 0xcf, 0xcf, 0xd6, 0x03, 0x33, 0x95, 0x54, 0xb8, 0x78, 0xdc, 0x53, 0x4b,
	0x9c, 0xf2, 0xf3, 0xb3, 0x43, 0x32, 0x4b, 0x72, 0x52, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x40, 0x0c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc9, 0x86, 0x8b, 0x07, 0xa4, 0x24, 0x28,
	0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x18, 0x87, 0x2a, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c,
	0xfc, 0x22, 0x09, 0x26, 0xb0, 0x30, 0x94, 0x67, 0x64, 0xc8, 0xc5, 0x02, 0xd2, 0x2d, 0xa4, 0xc9,
	0xc5, 0x0e, 0xb5, 0x4b, 0x88, 0x57, 0x0f, 0xd9, 0x56, 0x29, 0x5e, 0x3d, 0x64, 0xe3, 0x95, 0x18,
	0x92, 0xd8, 0xc0, 0xae, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xa4, 0x72, 0x93, 0xb1,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookClient interface {
	GetBook(ctx context.Context, in *GetBookTitle, opts ...grpc.CallOption) (*BookResponse, error)
}

type bookClient struct {
	cc *grpc.ClientConn
}

func NewBookClient(cc *grpc.ClientConn) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetBook(ctx context.Context, in *GetBookTitle, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/Book/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
type BookServer interface {
	GetBook(context.Context, *GetBookTitle) (*BookResponse, error)
}

// UnimplementedBookServer can be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (*UnimplementedBookServer) GetBook(ctx context.Context, req *GetBookTitle) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}

func RegisterBookServer(s *grpc.Server, srv BookServer) {
	s.RegisterService(&_Book_serviceDesc, srv)
}

func _Book_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Book/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBook(ctx, req.(*GetBookTitle))
	}
	return interceptor(ctx, in, info, handler)
}

var _Book_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _Book_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book.proto",
}