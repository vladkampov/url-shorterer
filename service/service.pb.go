// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type FullURLObject struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Visited              int32    `protobuf:"varint,3,opt,name=visited,proto3" json:"visited,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FullURLObject) Reset()         { *m = FullURLObject{} }
func (m *FullURLObject) String() string { return proto.CompactTextString(m) }
func (*FullURLObject) ProtoMessage()    {}
func (*FullURLObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *FullURLObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FullURLObject.Unmarshal(m, b)
}
func (m *FullURLObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FullURLObject.Marshal(b, m, deterministic)
}
func (m *FullURLObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FullURLObject.Merge(m, src)
}
func (m *FullURLObject) XXX_Size() int {
	return xxx_messageInfo_FullURLObject.Size(m)
}
func (m *FullURLObject) XXX_DiscardUnknown() {
	xxx_messageInfo_FullURLObject.DiscardUnknown(m)
}

var xxx_messageInfo_FullURLObject proto.InternalMessageInfo

func (m *FullURLObject) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *FullURLObject) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *FullURLObject) GetVisited() int32 {
	if m != nil {
		return m.Visited
	}
	return 0
}

func (m *FullURLObject) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type URLRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URLRequest) Reset()         { *m = URLRequest{} }
func (m *URLRequest) String() string { return proto.CompactTextString(m) }
func (*URLRequest) ProtoMessage()    {}
func (*URLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *URLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URLRequest.Unmarshal(m, b)
}
func (m *URLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URLRequest.Marshal(b, m, deterministic)
}
func (m *URLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URLRequest.Merge(m, src)
}
func (m *URLRequest) XXX_Size() int {
	return xxx_messageInfo_URLRequest.Size(m)
}
func (m *URLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_URLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_URLRequest proto.InternalMessageInfo

func (m *URLRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *URLRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type HashedURLReply struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashedURLReply) Reset()         { *m = HashedURLReply{} }
func (m *HashedURLReply) String() string { return proto.CompactTextString(m) }
func (*HashedURLReply) ProtoMessage()    {}
func (*HashedURLReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *HashedURLReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashedURLReply.Unmarshal(m, b)
}
func (m *HashedURLReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashedURLReply.Marshal(b, m, deterministic)
}
func (m *HashedURLReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashedURLReply.Merge(m, src)
}
func (m *HashedURLReply) XXX_Size() int {
	return xxx_messageInfo_HashedURLReply.Size(m)
}
func (m *HashedURLReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HashedURLReply.DiscardUnknown(m)
}

var xxx_messageInfo_HashedURLReply proto.InternalMessageInfo

func (m *HashedURLReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UserIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (m *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(m, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ArrayURLsReply struct {
	Urls                 []*FullURLObject `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArrayURLsReply) Reset()         { *m = ArrayURLsReply{} }
func (m *ArrayURLsReply) String() string { return proto.CompactTextString(m) }
func (*ArrayURLsReply) ProtoMessage()    {}
func (*ArrayURLsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ArrayURLsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayURLsReply.Unmarshal(m, b)
}
func (m *ArrayURLsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayURLsReply.Marshal(b, m, deterministic)
}
func (m *ArrayURLsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayURLsReply.Merge(m, src)
}
func (m *ArrayURLsReply) XXX_Size() int {
	return xxx_messageInfo_ArrayURLsReply.Size(m)
}
func (m *ArrayURLsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayURLsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayURLsReply proto.InternalMessageInfo

func (m *ArrayURLsReply) GetUrls() []*FullURLObject {
	if m != nil {
		return m.Urls
	}
	return nil
}

type HashedUrlRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashedUrlRequest) Reset()         { *m = HashedUrlRequest{} }
func (m *HashedUrlRequest) String() string { return proto.CompactTextString(m) }
func (*HashedUrlRequest) ProtoMessage()    {}
func (*HashedUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *HashedUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashedUrlRequest.Unmarshal(m, b)
}
func (m *HashedUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashedUrlRequest.Marshal(b, m, deterministic)
}
func (m *HashedUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashedUrlRequest.Merge(m, src)
}
func (m *HashedUrlRequest) XXX_Size() int {
	return xxx_messageInfo_HashedUrlRequest.Size(m)
}
func (m *HashedUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HashedUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HashedUrlRequest proto.InternalMessageInfo

func (m *HashedUrlRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *HashedUrlRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type URLReply struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Visited              int32    `protobuf:"varint,2,opt,name=visited,proto3" json:"visited,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URLReply) Reset()         { *m = URLReply{} }
func (m *URLReply) String() string { return proto.CompactTextString(m) }
func (*URLReply) ProtoMessage()    {}
func (*URLReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *URLReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URLReply.Unmarshal(m, b)
}
func (m *URLReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URLReply.Marshal(b, m, deterministic)
}
func (m *URLReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URLReply.Merge(m, src)
}
func (m *URLReply) XXX_Size() int {
	return xxx_messageInfo_URLReply.Size(m)
}
func (m *URLReply) XXX_DiscardUnknown() {
	xxx_messageInfo_URLReply.DiscardUnknown(m)
}

var xxx_messageInfo_URLReply proto.InternalMessageInfo

func (m *URLReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *URLReply) GetVisited() int32 {
	if m != nil {
		return m.Visited
	}
	return 0
}

func init() {
	proto.RegisterType((*FullURLObject)(nil), "FullURLObject")
	proto.RegisterType((*URLRequest)(nil), "URLRequest")
	proto.RegisterType((*HashedURLReply)(nil), "HashedURLReply")
	proto.RegisterType((*UserIdRequest)(nil), "UserIdRequest")
	proto.RegisterType((*ArrayURLsReply)(nil), "ArrayURLsReply")
	proto.RegisterType((*HashedUrlRequest)(nil), "HashedUrlRequest")
	proto.RegisterType((*URLReply)(nil), "URLReply")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x29, 0x20, 0xb8, 0x43, 0x58, 0xb0, 0x17, 0x37, 0x9c, 0x48, 0x4f, 0xeb, 0xa5, 0x07,
	0x34, 0x7a, 0x34, 0x5e, 0x44, 0x13, 0x8c, 0x49, 0x4d, 0xcf, 0x06, 0xd8, 0x49, 0x76, 0x4d, 0x23,
	0x38, 0xed, 0x92, 0xec, 0xbf, 0xf0, 0x27, 0x1b, 0x2a, 0x1b, 0xb7, 0x2a, 0xb7, 0x99, 0x76, 0xde,
	0x9b, 0xd7, 0x2f, 0x85, 0xa1, 0x45, 0xda, 0x15, 0x6b, 0x94, 0x5b, 0xda, 0xb8, 0x8d, 0xc8, 0x61,
	0x78, 0x5f, 0x1a, 0xa3, 0xd5, 0xe2, 0x79, 0xf5, 0x86, 0x6b, 0xc7, 0xc7, 0xd0, 0x29, 0xc9, 0x24,
	0x6c, 0xca, 0xd2, 0x48, 0xed, 0x4b, 0xce, 0xa1, 0x9b, 0x2f, 0x6d, 0x9e, 0xb4, 0xfd, 0x91, 0xaf,
	0x79, 0x02, 0xfd, 0x5d, 0x61, 0x0b, 0x87, 0x59, 0xd2, 0x99, 0xb2, 0xf4, 0x44, 0xd5, 0x2d, 0x3f,
	0x87, 0x7e, 0x69, 0x91, 0x5e, 0x8b, 0x2c, 0xe9, 0x7a, 0x41, 0x6f, 0xdf, 0x3e, 0x66, 0xe2, 0x06,
	0x40, 0xab, 0x85, 0xc2, 0x8f, 0x12, 0xed, 0x7f, 0x6b, 0x1a, 0xc2, 0x76, 0x20, 0x14, 0x10, 0x3f,
	0x2c, 0x6d, 0x8e, 0x99, 0x97, 0x6f, 0x4d, 0xf5, 0x57, 0x2c, 0x52, 0x18, 0x6a, 0x3f, 0x5d, 0xfb,
	0x37, 0xdc, 0x58, 0xe0, 0x76, 0x05, 0xf1, 0x1d, 0xd1, 0xb2, 0xd2, 0x6a, 0x61, 0xbf, 0xdd, 0x04,
	0x74, 0x4b, 0x32, 0x36, 0x61, 0xd3, 0x4e, 0x3a, 0x98, 0xc5, 0x32, 0xe0, 0xa1, 0xfc, 0x9d, 0xb8,
	0x85, 0xf1, 0x21, 0x03, 0x99, 0x7a, 0x45, 0xcd, 0x85, 0x35, 0xb8, 0x1c, 0x7d, 0xc4, 0x35, 0x9c,
	0x1e, 0x8f, 0xdf, 0xc4, 0xd9, 0x0e, 0x70, 0xce, 0x3e, 0x19, 0x44, 0x2f, 0xf9, 0x86, 0x1c, 0xbe,
	0x23, 0xf1, 0x0b, 0xe8, 0x1f, 0x1a, 0x3e, 0x90, 0x3f, 0x34, 0x27, 0x23, 0x19, 0x12, 0x12, 0x2d,
	0x9e, 0x42, 0x6f, 0x8e, 0x4e, 0x93, 0xe1, 0x67, 0xf2, 0x77, 0xf4, 0x49, 0x24, 0x1b, 0x93, 0x12,
	0xa2, 0x39, 0xba, 0xa7, 0x4a, 0x93, 0xb1, 0x3c, 0x96, 0x01, 0xc7, 0xc9, 0x48, 0x86, 0xb4, 0x44,
	0x6b, 0xd5, 0xf3, 0x3f, 0xe7, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xf8, 0x86, 0x93, 0x4a,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortenerClient interface {
	Shorten(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*HashedURLReply, error)
	GetUrl(ctx context.Context, in *HashedUrlRequest, opts ...grpc.CallOption) (*URLReply, error)
	GetMyUrls(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ArrayURLsReply, error)
}

type shortenerClient struct {
	cc *grpc.ClientConn
}

func NewShortenerClient(cc *grpc.ClientConn) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) Shorten(ctx context.Context, in *URLRequest, opts ...grpc.CallOption) (*HashedURLReply, error) {
	out := new(HashedURLReply)
	err := c.cc.Invoke(ctx, "/Shortener/Shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetUrl(ctx context.Context, in *HashedUrlRequest, opts ...grpc.CallOption) (*URLReply, error) {
	out := new(URLReply)
	err := c.cc.Invoke(ctx, "/Shortener/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetMyUrls(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*ArrayURLsReply, error) {
	out := new(ArrayURLsReply)
	err := c.cc.Invoke(ctx, "/Shortener/GetMyUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
type ShortenerServer interface {
	Shorten(context.Context, *URLRequest) (*HashedURLReply, error)
	GetUrl(context.Context, *HashedUrlRequest) (*URLReply, error)
	GetMyUrls(context.Context, *UserIdRequest) (*ArrayURLsReply, error)
}

func RegisterShortenerServer(s *grpc.Server, srv ShortenerServer) {
	s.RegisterService(&_Shortener_serviceDesc, srv)
}

func _Shortener_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Shorten(ctx, req.(*URLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashedUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetUrl(ctx, req.(*HashedUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetMyUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetMyUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortener/GetMyUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetMyUrls(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shortener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shorten",
			Handler:    _Shortener_Shorten_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _Shortener_GetUrl_Handler,
		},
		{
			MethodName: "GetMyUrls",
			Handler:    _Shortener_GetMyUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
