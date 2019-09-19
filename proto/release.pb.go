// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/release.proto

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

type Release struct {
	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Relations
	AnimeId              int64    `protobuf:"varint,3,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{0}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Release) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

type ReleaseQuery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AnimeId              int64    `protobuf:"varint,2,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseQuery) Reset()         { *m = ReleaseQuery{} }
func (m *ReleaseQuery) String() string { return proto.CompactTextString(m) }
func (*ReleaseQuery) ProtoMessage()    {}
func (*ReleaseQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{1}
}

func (m *ReleaseQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseQuery.Unmarshal(m, b)
}
func (m *ReleaseQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseQuery.Marshal(b, m, deterministic)
}
func (m *ReleaseQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseQuery.Merge(m, src)
}
func (m *ReleaseQuery) XXX_Size() int {
	return xxx_messageInfo_ReleaseQuery.Size(m)
}
func (m *ReleaseQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseQuery proto.InternalMessageInfo

func (m *ReleaseQuery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseQuery) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

type ReleaseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AnimeId              int64    `protobuf:"varint,2,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseRequest) Reset()         { *m = ReleaseRequest{} }
func (m *ReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseRequest) ProtoMessage()    {}
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{2}
}

func (m *ReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseRequest.Unmarshal(m, b)
}
func (m *ReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseRequest.Merge(m, src)
}
func (m *ReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseRequest.Size(m)
}
func (m *ReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseRequest proto.InternalMessageInfo

func (m *ReleaseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseRequest) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

type ReleasesRequest struct {
	Query                *ReleaseQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReleasesRequest) Reset()         { *m = ReleasesRequest{} }
func (m *ReleasesRequest) String() string { return proto.CompactTextString(m) }
func (*ReleasesRequest) ProtoMessage()    {}
func (*ReleasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{3}
}

func (m *ReleasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasesRequest.Unmarshal(m, b)
}
func (m *ReleasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasesRequest.Marshal(b, m, deterministic)
}
func (m *ReleasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasesRequest.Merge(m, src)
}
func (m *ReleasesRequest) XXX_Size() int {
	return xxx_messageInfo_ReleasesRequest.Size(m)
}
func (m *ReleasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasesRequest proto.InternalMessageInfo

func (m *ReleasesRequest) GetQuery() *ReleaseQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type ReleaseResponse struct {
	Release              *Release `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseResponse) Reset()         { *m = ReleaseResponse{} }
func (m *ReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseResponse) ProtoMessage()    {}
func (*ReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{4}
}

func (m *ReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseResponse.Unmarshal(m, b)
}
func (m *ReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseResponse.Merge(m, src)
}
func (m *ReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseResponse.Size(m)
}
func (m *ReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseResponse proto.InternalMessageInfo

func (m *ReleaseResponse) GetRelease() *Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type ReleasesResponse struct {
	Releases             []*Release `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReleasesResponse) Reset()         { *m = ReleasesResponse{} }
func (m *ReleasesResponse) String() string { return proto.CompactTextString(m) }
func (*ReleasesResponse) ProtoMessage()    {}
func (*ReleasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{5}
}

func (m *ReleasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasesResponse.Unmarshal(m, b)
}
func (m *ReleasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasesResponse.Marshal(b, m, deterministic)
}
func (m *ReleasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasesResponse.Merge(m, src)
}
func (m *ReleasesResponse) XXX_Size() int {
	return xxx_messageInfo_ReleasesResponse.Size(m)
}
func (m *ReleasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasesResponse proto.InternalMessageInfo

func (m *ReleasesResponse) GetReleases() []*Release {
	if m != nil {
		return m.Releases
	}
	return nil
}

func init() {
	proto.RegisterType((*Release)(nil), "proto.Release")
	proto.RegisterType((*ReleaseQuery)(nil), "proto.ReleaseQuery")
	proto.RegisterType((*ReleaseRequest)(nil), "proto.ReleaseRequest")
	proto.RegisterType((*ReleasesRequest)(nil), "proto.ReleasesRequest")
	proto.RegisterType((*ReleaseResponse)(nil), "proto.ReleaseResponse")
	proto.RegisterType((*ReleasesResponse)(nil), "proto.ReleasesResponse")
}

func init() { proto.RegisterFile("proto/release.proto", fileDescriptor_987cf5695bafb180) }

var fileDescriptor_987cf5695bafb180 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0xc6, 0x95, 0x56, 0xa5, 0xc5, 0xa0, 0x82, 0x32, 0xfe, 0x94, 0x9d, 0xaa, 0x9e, 0x0a, 0x87,
	0x21, 0x8d, 0x0b, 0xa8, 0x88, 0x3b, 0xdc, 0x08, 0x0f, 0x80, 0x06, 0xf1, 0x21, 0xd2, 0x58, 0xb7,
	0x24, 0x43, 0xe2, 0x09, 0x78, 0x6d, 0x44, 0xe2, 0x74, 0x5b, 0xcb, 0x85, 0x53, 0xe4, 0xcf, 0xfe,
	0xd9, 0xfe, 0x1c, 0x18, 0x2d, 0x75, 0x6b, 0xdb, 0x6b, 0x8d, 0x73, 0x9c, 0x19, 0x9c, 0xb8, 0x88,
	0x27, 0xee, 0xa9, 0x9e, 0x20, 0x15, 0x5e, 0xe7, 0x39, 0x44, 0x4a, 0x16, 0xac, 0x64, 0x75, 0x2c,
	0x22, 0x25, 0xf9, 0x09, 0x24, 0x56, 0xd9, 0x39, 0x16, 0x51, 0xc9, 0xea, 0x7d, 0xe1, 0x03, 0x7e,
	0x01, 0xd9, 0x6c, 0xa1, 0x3e, 0xf0, 0x55, 0xc9, 0x22, 0x76, 0xb5, 0xa9, 0x8b, 0x1f, 0x65, 0x75,
	0x07, 0x87, 0xd4, 0xeb, 0x79, 0x8d, 0xfa, 0x6b, 0xd0, 0x70, 0x1b, 0x8d, 0x76, 0xd1, 0x06, 0x72,
	0x42, 0x05, 0xae, 0xd6, 0x68, 0xec, 0x7f, 0xe0, 0x7b, 0x38, 0x22, 0xd8, 0x04, 0xfa, 0x12, 0x92,
	0xd5, 0xef, 0x0e, 0xae, 0xc1, 0xc1, 0x74, 0xe4, 0x4d, 0x4f, 0xb6, 0xd7, 0x13, 0xbe, 0xa2, 0x6a,
	0x3a, 0x5a, 0xa0, 0x59, 0xb6, 0x0b, 0x83, 0xbc, 0x86, 0x94, 0x8e, 0x45, 0x7c, 0xbe, 0xcb, 0x8b,
	0x90, 0xae, 0x1e, 0xe0, 0x78, 0x33, 0x9a, 0xe8, 0x2b, 0xc8, 0x28, 0x6d, 0x0a, 0x56, 0xc6, 0x7f,
	0xe0, 0x5d, 0x7e, 0xfa, 0xcd, 0x3a, 0xe3, 0x2f, 0xa8, 0x3f, 0xd5, 0x3b, 0xf2, 0xdb, 0xcd, 0x8f,
	0x9c, 0xf6, 0x38, 0x6f, 0x6e, 0x7c, 0xd6, 0x97, 0x69, 0x70, 0x03, 0x59, 0x58, 0x86, 0xf7, 0x6a,
	0xc2, 0x61, 0xc6, 0xe7, 0x03, 0xdd, 0xc3, 0x6f, 0x7b, 0x4e, 0xbf, 0xf9, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x46, 0xb4, 0xfd, 0x76, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
	Releases(ctx context.Context, in *ReleasesRequest, opts ...grpc.CallOption) (*ReleasesResponse, error)
}

type releaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewReleaseServiceClient(cc *grpc.ClientConn) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Releases(ctx context.Context, in *ReleasesRequest, opts ...grpc.CallOption) (*ReleasesResponse, error) {
	out := new(ReleasesResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/Releases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
type ReleaseServiceServer interface {
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	Releases(context.Context, *ReleasesRequest) (*ReleasesResponse, error)
}

// UnimplementedReleaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (*UnimplementedReleaseServiceServer) Release(ctx context.Context, req *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (*UnimplementedReleaseServiceServer) Releases(ctx context.Context, req *ReleasesRequest) (*ReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Releases not implemented")
}

func RegisterReleaseServiceServer(s *grpc.Server, srv ReleaseServiceServer) {
	s.RegisterService(&_ReleaseService_serviceDesc, srv)
}

func _ReleaseService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Releases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Releases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/Releases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Releases(ctx, req.(*ReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Release",
			Handler:    _ReleaseService_Release_Handler,
		},
		{
			MethodName: "Releases",
			Handler:    _ReleaseService_Releases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/release.proto",
}
