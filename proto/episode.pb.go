// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/episode.proto

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

type Episode struct {
	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Thumbnail    string `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	EpisodeOrder int64  `protobuf:"varint,3,opt,name=episode_order,json=episodeOrder,proto3" json:"episode_order,omitempty"`
	EpisodeCode  string `protobuf:"bytes,4,opt,name=episode_code,json=episodeCode,proto3" json:"episode_code,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relations
	ReleaseId            int64    `protobuf:"varint,7,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Episode) Reset()         { *m = Episode{} }
func (m *Episode) String() string { return proto.CompactTextString(m) }
func (*Episode) ProtoMessage()    {}
func (*Episode) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{0}
}

func (m *Episode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Episode.Unmarshal(m, b)
}
func (m *Episode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Episode.Marshal(b, m, deterministic)
}
func (m *Episode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Episode.Merge(m, src)
}
func (m *Episode) XXX_Size() int {
	return xxx_messageInfo_Episode.Size(m)
}
func (m *Episode) XXX_DiscardUnknown() {
	xxx_messageInfo_Episode.DiscardUnknown(m)
}

var xxx_messageInfo_Episode proto.InternalMessageInfo

func (m *Episode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Episode) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *Episode) GetEpisodeOrder() int64 {
	if m != nil {
		return m.EpisodeOrder
	}
	return 0
}

func (m *Episode) GetEpisodeCode() string {
	if m != nil {
		return m.EpisodeCode
	}
	return ""
}

func (m *Episode) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Episode) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Episode) GetReleaseId() int64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

type EpisodeQuery struct {
	ReleaseId int64 `protobuf:"varint,3,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	// Pagination
	Offset               int64    `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,11,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EpisodeQuery) Reset()         { *m = EpisodeQuery{} }
func (m *EpisodeQuery) String() string { return proto.CompactTextString(m) }
func (*EpisodeQuery) ProtoMessage()    {}
func (*EpisodeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{1}
}

func (m *EpisodeQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpisodeQuery.Unmarshal(m, b)
}
func (m *EpisodeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpisodeQuery.Marshal(b, m, deterministic)
}
func (m *EpisodeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpisodeQuery.Merge(m, src)
}
func (m *EpisodeQuery) XXX_Size() int {
	return xxx_messageInfo_EpisodeQuery.Size(m)
}
func (m *EpisodeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_EpisodeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_EpisodeQuery proto.InternalMessageInfo

func (m *EpisodeQuery) GetReleaseId() int64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

func (m *EpisodeQuery) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *EpisodeQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type EpisodeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EpisodeRequest) Reset()         { *m = EpisodeRequest{} }
func (m *EpisodeRequest) String() string { return proto.CompactTextString(m) }
func (*EpisodeRequest) ProtoMessage()    {}
func (*EpisodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{2}
}

func (m *EpisodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpisodeRequest.Unmarshal(m, b)
}
func (m *EpisodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpisodeRequest.Marshal(b, m, deterministic)
}
func (m *EpisodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpisodeRequest.Merge(m, src)
}
func (m *EpisodeRequest) XXX_Size() int {
	return xxx_messageInfo_EpisodeRequest.Size(m)
}
func (m *EpisodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EpisodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EpisodeRequest proto.InternalMessageInfo

func (m *EpisodeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EpisodesRequest struct {
	Query                *EpisodeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	OrderBy              *OrderBy      `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EpisodesRequest) Reset()         { *m = EpisodesRequest{} }
func (m *EpisodesRequest) String() string { return proto.CompactTextString(m) }
func (*EpisodesRequest) ProtoMessage()    {}
func (*EpisodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{3}
}

func (m *EpisodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpisodesRequest.Unmarshal(m, b)
}
func (m *EpisodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpisodesRequest.Marshal(b, m, deterministic)
}
func (m *EpisodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpisodesRequest.Merge(m, src)
}
func (m *EpisodesRequest) XXX_Size() int {
	return xxx_messageInfo_EpisodesRequest.Size(m)
}
func (m *EpisodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EpisodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EpisodesRequest proto.InternalMessageInfo

func (m *EpisodesRequest) GetQuery() *EpisodeQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *EpisodesRequest) GetOrderBy() *OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

type EpisodeResponse struct {
	Episode              *Episode `protobuf:"bytes,1,opt,name=episode,proto3" json:"episode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EpisodeResponse) Reset()         { *m = EpisodeResponse{} }
func (m *EpisodeResponse) String() string { return proto.CompactTextString(m) }
func (*EpisodeResponse) ProtoMessage()    {}
func (*EpisodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{4}
}

func (m *EpisodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpisodeResponse.Unmarshal(m, b)
}
func (m *EpisodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpisodeResponse.Marshal(b, m, deterministic)
}
func (m *EpisodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpisodeResponse.Merge(m, src)
}
func (m *EpisodeResponse) XXX_Size() int {
	return xxx_messageInfo_EpisodeResponse.Size(m)
}
func (m *EpisodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EpisodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EpisodeResponse proto.InternalMessageInfo

func (m *EpisodeResponse) GetEpisode() *Episode {
	if m != nil {
		return m.Episode
	}
	return nil
}

type EpisodesResponse struct {
	Episodes             []*Episode `protobuf:"bytes,1,rep,name=episodes,proto3" json:"episodes,omitempty"`
	Count                uint64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EpisodesResponse) Reset()         { *m = EpisodesResponse{} }
func (m *EpisodesResponse) String() string { return proto.CompactTextString(m) }
func (*EpisodesResponse) ProtoMessage()    {}
func (*EpisodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_536a535dba699a81, []int{5}
}

func (m *EpisodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpisodesResponse.Unmarshal(m, b)
}
func (m *EpisodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpisodesResponse.Marshal(b, m, deterministic)
}
func (m *EpisodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpisodesResponse.Merge(m, src)
}
func (m *EpisodesResponse) XXX_Size() int {
	return xxx_messageInfo_EpisodesResponse.Size(m)
}
func (m *EpisodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EpisodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EpisodesResponse proto.InternalMessageInfo

func (m *EpisodesResponse) GetEpisodes() []*Episode {
	if m != nil {
		return m.Episodes
	}
	return nil
}

func (m *EpisodesResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Episode)(nil), "proto.Episode")
	proto.RegisterType((*EpisodeQuery)(nil), "proto.EpisodeQuery")
	proto.RegisterType((*EpisodeRequest)(nil), "proto.EpisodeRequest")
	proto.RegisterType((*EpisodesRequest)(nil), "proto.EpisodesRequest")
	proto.RegisterType((*EpisodeResponse)(nil), "proto.EpisodeResponse")
	proto.RegisterType((*EpisodesResponse)(nil), "proto.EpisodesResponse")
}

func init() { proto.RegisterFile("proto/episode.proto", fileDescriptor_536a535dba699a81) }

var fileDescriptor_536a535dba699a81 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x3a, 0x4e, 0xc6, 0xc5, 0x45, 0x5b, 0x28, 0xab, 0x08, 0x24, 0x63, 0x2e, 0x29,
	0x87, 0x22, 0x85, 0x0b, 0x52, 0x4f, 0x2d, 0x42, 0x82, 0x0b, 0x88, 0x85, 0x1b, 0x07, 0xcb, 0xf1,
	0x4e, 0x61, 0xa5, 0xd8, 0xeb, 0x7a, 0xd7, 0xa0, 0xfc, 0x04, 0x5f, 0xc8, 0xc7, 0x20, 0xef, 0x8e,
	0x13, 0x30, 0x9c, 0xe0, 0x94, 0xcc, 0x7b, 0x6f, 0xe6, 0x3d, 0x3f, 0x1b, 0x4e, 0x9b, 0x56, 0x5b,
	0xfd, 0x0c, 0x1b, 0x65, 0xb4, 0xc4, 0x0b, 0x37, 0xb1, 0xd0, 0xfd, 0x2c, 0x99, 0xe7, 0x4a, 0x5d,
	0x55, 0xba, 0xf6, 0x54, 0xf6, 0x23, 0x80, 0xe8, 0x95, 0x17, 0xb3, 0x04, 0x26, 0x4a, 0xf2, 0x20,
	0x0d, 0x56, 0x53, 0x31, 0x51, 0x92, 0x3d, 0x84, 0x85, 0xfd, 0xd2, 0x55, 0x9b, 0xba, 0x50, 0x5b,
	0x3e, 0x49, 0x83, 0xd5, 0x42, 0x1c, 0x00, 0xf6, 0x04, 0xee, 0x90, 0x4b, 0xae, 0x5b, 0x89, 0x2d,
	0x9f, 0xba, 0xc5, 0x63, 0x02, 0xdf, 0xf5, 0x18, 0x7b, 0x0c, 0xc3, 0x9c, 0x97, 0x5a, 0x22, 0x3f,
	0x72, 0x57, 0x62, 0xc2, 0x5e, 0xf6, 0xae, 0x8f, 0x00, 0xca, 0x16, 0x0b, 0x8b, 0x32, 0x2f, 0x2c,
	0x0f, 0xbd, 0x0d, 0x21, 0x57, 0xb6, 0xa7, 0xbb, 0x46, 0x0e, 0xf4, 0xcc, 0xd3, 0x84, 0x78, 0xba,
	0xc5, 0x2d, 0x16, 0x06, 0x73, 0x25, 0x79, 0xe4, 0x22, 0x2c, 0x08, 0x79, 0x23, 0xb3, 0x4f, 0x70,
	0x4c, 0x4f, 0xf7, 0xbe, 0xc3, 0x76, 0x37, 0x92, 0x4f, 0x47, 0x72, 0x76, 0x06, 0x33, 0x7d, 0x73,
	0x63, 0xd0, 0x72, 0x70, 0x14, 0x4d, 0xec, 0x1e, 0x84, 0x5b, 0x55, 0x29, 0xcb, 0x63, 0x07, 0xfb,
	0x21, 0x4b, 0x21, 0xa1, 0xe3, 0x02, 0x6f, 0x3b, 0x34, 0x76, 0xdc, 0x60, 0xf6, 0x19, 0x4e, 0x48,
	0x61, 0x06, 0xc9, 0x39, 0x84, 0xb7, 0x7d, 0x14, 0xa7, 0x8a, 0xd7, 0xa7, 0xfe, 0x3d, 0x5c, 0xfc,
	0x9a, 0x52, 0x78, 0x05, 0x3b, 0x87, 0xb9, 0x6b, 0x36, 0xdf, 0xec, 0x5c, 0xfd, 0xf1, 0x3a, 0x21,
	0xb5, 0x2b, 0xf7, 0x7a, 0x27, 0x22, 0xed, 0xff, 0x64, 0x97, 0x7b, 0x23, 0x81, 0xa6, 0xd1, 0xb5,
	0x41, 0xb6, 0x82, 0x88, 0x6a, 0x26, 0xab, 0xe4, 0x77, 0x2b, 0x31, 0xd0, 0xd9, 0x47, 0xb8, 0x7b,
	0x48, 0x49, 0xdb, 0x4f, 0x61, 0x4e, 0xb4, 0xe1, 0x41, 0x3a, 0xfd, 0xcb, 0xfa, 0x9e, 0xef, 0xdb,
	0x29, 0x75, 0x57, 0x5b, 0x17, 0xf2, 0x48, 0xf8, 0x61, 0xfd, 0x7d, 0xb2, 0xaf, 0xe7, 0x03, 0xb6,
	0x5f, 0x55, 0x89, 0xec, 0xc5, 0xe1, 0x5b, 0xbb, 0x3f, 0xba, 0xe6, 0xdb, 0x59, 0x9e, 0x8d, 0x61,
	0x8a, 0x73, 0x09, 0xf3, 0x21, 0x22, 0x1b, 0x69, 0x86, 0x66, 0x97, 0x0f, 0xfe, 0xc0, 0x69, 0xf9,
	0x1a, 0x4e, 0x5e, 0x6b, 0x6b, 0xd1, 0xd8, 0x7f, 0xbf, 0x71, 0x05, 0xc9, 0x5b, 0xfc, 0xf6, 0x3f,
	0x27, 0x36, 0x33, 0x87, 0x3f, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x5a, 0x8f, 0x40, 0xa3,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EpisodeServiceClient is the client API for EpisodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EpisodeServiceClient interface {
	Episode(ctx context.Context, in *EpisodeRequest, opts ...grpc.CallOption) (*EpisodeResponse, error)
	Episodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error)
	HottestEpisodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error)
	NewestEpisodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error)
}

type episodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEpisodeServiceClient(cc *grpc.ClientConn) EpisodeServiceClient {
	return &episodeServiceClient{cc}
}

func (c *episodeServiceClient) Episode(ctx context.Context, in *EpisodeRequest, opts ...grpc.CallOption) (*EpisodeResponse, error) {
	out := new(EpisodeResponse)
	err := c.cc.Invoke(ctx, "/proto.EpisodeService/Episode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) Episodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error) {
	out := new(EpisodesResponse)
	err := c.cc.Invoke(ctx, "/proto.EpisodeService/Episodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) HottestEpisodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error) {
	out := new(EpisodesResponse)
	err := c.cc.Invoke(ctx, "/proto.EpisodeService/HottestEpisodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *episodeServiceClient) NewestEpisodes(ctx context.Context, in *EpisodesRequest, opts ...grpc.CallOption) (*EpisodesResponse, error) {
	out := new(EpisodesResponse)
	err := c.cc.Invoke(ctx, "/proto.EpisodeService/NewestEpisodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EpisodeServiceServer is the server API for EpisodeService service.
type EpisodeServiceServer interface {
	Episode(context.Context, *EpisodeRequest) (*EpisodeResponse, error)
	Episodes(context.Context, *EpisodesRequest) (*EpisodesResponse, error)
	HottestEpisodes(context.Context, *EpisodesRequest) (*EpisodesResponse, error)
	NewestEpisodes(context.Context, *EpisodesRequest) (*EpisodesResponse, error)
}

// UnimplementedEpisodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEpisodeServiceServer struct {
}

func (*UnimplementedEpisodeServiceServer) Episode(ctx context.Context, req *EpisodeRequest) (*EpisodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Episode not implemented")
}
func (*UnimplementedEpisodeServiceServer) Episodes(ctx context.Context, req *EpisodesRequest) (*EpisodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Episodes not implemented")
}
func (*UnimplementedEpisodeServiceServer) HottestEpisodes(ctx context.Context, req *EpisodesRequest) (*EpisodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HottestEpisodes not implemented")
}
func (*UnimplementedEpisodeServiceServer) NewestEpisodes(ctx context.Context, req *EpisodesRequest) (*EpisodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewestEpisodes not implemented")
}

func RegisterEpisodeServiceServer(s *grpc.Server, srv EpisodeServiceServer) {
	s.RegisterService(&_EpisodeService_serviceDesc, srv)
}

func _EpisodeService_Episode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).Episode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EpisodeService/Episode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).Episode(ctx, req.(*EpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_Episodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).Episodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EpisodeService/Episodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).Episodes(ctx, req.(*EpisodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_HottestEpisodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).HottestEpisodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EpisodeService/HottestEpisodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).HottestEpisodes(ctx, req.(*EpisodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpisodeService_NewestEpisodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpisodeServiceServer).NewestEpisodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EpisodeService/NewestEpisodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpisodeServiceServer).NewestEpisodes(ctx, req.(*EpisodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EpisodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EpisodeService",
	HandlerType: (*EpisodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Episode",
			Handler:    _EpisodeService_Episode_Handler,
		},
		{
			MethodName: "Episodes",
			Handler:    _EpisodeService_Episodes_Handler,
		},
		{
			MethodName: "HottestEpisodes",
			Handler:    _EpisodeService_HottestEpisodes_Handler,
		},
		{
			MethodName: "NewestEpisodes",
			Handler:    _EpisodeService_NewestEpisodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/episode.proto",
}