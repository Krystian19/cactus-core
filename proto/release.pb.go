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
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Poster        string `protobuf:"bytes,3,opt,name=poster,proto3" json:"poster,omitempty"`
	Background    string `protobuf:"bytes,4,opt,name=background,proto3" json:"background,omitempty"`
	ReleaseOrder  int64  `protobuf:"varint,5,opt,name=release_order,json=releaseOrder,proto3" json:"release_order,omitempty"`
	StartedAiring string `protobuf:"bytes,6,opt,name=started_airing,json=startedAiring,proto3" json:"started_airing,omitempty"`
	StoppedAiring string `protobuf:"bytes,7,opt,name=stopped_airing,json=stoppedAiring,proto3" json:"stopped_airing,omitempty"`
	CreatedAt     string `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relations
	AnimeId              int64    `protobuf:"varint,40,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	ReleaseTypeId        int64    `protobuf:"varint,41,opt,name=release_type_id,json=releaseTypeId,proto3" json:"release_type_id,omitempty"`
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

func (m *Release) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *Release) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *Release) GetReleaseOrder() int64 {
	if m != nil {
		return m.ReleaseOrder
	}
	return 0
}

func (m *Release) GetStartedAiring() string {
	if m != nil {
		return m.StartedAiring
	}
	return ""
}

func (m *Release) GetStoppedAiring() string {
	if m != nil {
		return m.StoppedAiring
	}
	return ""
}

func (m *Release) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Release) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Release) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

func (m *Release) GetReleaseTypeId() int64 {
	if m != nil {
		return m.ReleaseTypeId
	}
	return 0
}

type ReleaseQuery struct {
	AnimeId int64  `protobuf:"varint,1,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// List of Genre Ids
	Genres []int64 `protobuf:"varint,3,rep,packed,name=genres,proto3" json:"genres,omitempty"`
	// Pagination
	Offset               int64    `protobuf:"varint,30,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,31,opt,name=limit,proto3" json:"limit,omitempty"`
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

func (m *ReleaseQuery) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

func (m *ReleaseQuery) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReleaseQuery) GetGenres() []int64 {
	if m != nil {
		return m.Genres
	}
	return nil
}

func (m *ReleaseQuery) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReleaseQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReleaseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	Count                int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
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

func (m *ReleasesResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReleasesListResponse struct {
	Releases             []*Release `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReleasesListResponse) Reset()         { *m = ReleasesListResponse{} }
func (m *ReleasesListResponse) String() string { return proto.CompactTextString(m) }
func (*ReleasesListResponse) ProtoMessage()    {}
func (*ReleasesListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{6}
}

func (m *ReleasesListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasesListResponse.Unmarshal(m, b)
}
func (m *ReleasesListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasesListResponse.Marshal(b, m, deterministic)
}
func (m *ReleasesListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasesListResponse.Merge(m, src)
}
func (m *ReleasesListResponse) XXX_Size() int {
	return xxx_messageInfo_ReleasesListResponse.Size(m)
}
func (m *ReleasesListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasesListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasesListResponse proto.InternalMessageInfo

func (m *ReleasesListResponse) GetReleases() []*Release {
	if m != nil {
		return m.Releases
	}
	return nil
}

type ReleaseType struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseType) Reset()         { *m = ReleaseType{} }
func (m *ReleaseType) String() string { return proto.CompactTextString(m) }
func (*ReleaseType) ProtoMessage()    {}
func (*ReleaseType) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{7}
}

func (m *ReleaseType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseType.Unmarshal(m, b)
}
func (m *ReleaseType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseType.Marshal(b, m, deterministic)
}
func (m *ReleaseType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseType.Merge(m, src)
}
func (m *ReleaseType) XXX_Size() int {
	return xxx_messageInfo_ReleaseType.Size(m)
}
func (m *ReleaseType) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseType.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseType proto.InternalMessageInfo

func (m *ReleaseType) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseType) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type ReleaseTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseTypeRequest) Reset()         { *m = ReleaseTypeRequest{} }
func (m *ReleaseTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseTypeRequest) ProtoMessage()    {}
func (*ReleaseTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{8}
}

func (m *ReleaseTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseTypeRequest.Unmarshal(m, b)
}
func (m *ReleaseTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseTypeRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseTypeRequest.Merge(m, src)
}
func (m *ReleaseTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseTypeRequest.Size(m)
}
func (m *ReleaseTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseTypeRequest proto.InternalMessageInfo

func (m *ReleaseTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReleaseTypeResponse struct {
	ReleaseType          *ReleaseType `protobuf:"bytes,1,opt,name=release_type,json=releaseType,proto3" json:"release_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReleaseTypeResponse) Reset()         { *m = ReleaseTypeResponse{} }
func (m *ReleaseTypeResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseTypeResponse) ProtoMessage()    {}
func (*ReleaseTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{9}
}

func (m *ReleaseTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseTypeResponse.Unmarshal(m, b)
}
func (m *ReleaseTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseTypeResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseTypeResponse.Merge(m, src)
}
func (m *ReleaseTypeResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseTypeResponse.Size(m)
}
func (m *ReleaseTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseTypeResponse proto.InternalMessageInfo

func (m *ReleaseTypeResponse) GetReleaseType() *ReleaseType {
	if m != nil {
		return m.ReleaseType
	}
	return nil
}

type ReleaseDescription struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Relations
	ReleaseId            int64    `protobuf:"varint,3,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	LanguageId           int64    `protobuf:"varint,4,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseDescription) Reset()         { *m = ReleaseDescription{} }
func (m *ReleaseDescription) String() string { return proto.CompactTextString(m) }
func (*ReleaseDescription) ProtoMessage()    {}
func (*ReleaseDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{10}
}

func (m *ReleaseDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseDescription.Unmarshal(m, b)
}
func (m *ReleaseDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseDescription.Marshal(b, m, deterministic)
}
func (m *ReleaseDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseDescription.Merge(m, src)
}
func (m *ReleaseDescription) XXX_Size() int {
	return xxx_messageInfo_ReleaseDescription.Size(m)
}
func (m *ReleaseDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseDescription proto.InternalMessageInfo

func (m *ReleaseDescription) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseDescription) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReleaseDescription) GetReleaseId() int64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

func (m *ReleaseDescription) GetLanguageId() int64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

type ReleaseDescriptionsRequest struct {
	ReleaseId            int64    `protobuf:"varint,1,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseDescriptionsRequest) Reset()         { *m = ReleaseDescriptionsRequest{} }
func (m *ReleaseDescriptionsRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseDescriptionsRequest) ProtoMessage()    {}
func (*ReleaseDescriptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{11}
}

func (m *ReleaseDescriptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseDescriptionsRequest.Unmarshal(m, b)
}
func (m *ReleaseDescriptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseDescriptionsRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseDescriptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseDescriptionsRequest.Merge(m, src)
}
func (m *ReleaseDescriptionsRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseDescriptionsRequest.Size(m)
}
func (m *ReleaseDescriptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseDescriptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseDescriptionsRequest proto.InternalMessageInfo

func (m *ReleaseDescriptionsRequest) GetReleaseId() int64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

type ReleaseDescriptionsResponse struct {
	ReleaseDescriptions  []*ReleaseDescription `protobuf:"bytes,1,rep,name=ReleaseDescriptions,proto3" json:"ReleaseDescriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ReleaseDescriptionsResponse) Reset()         { *m = ReleaseDescriptionsResponse{} }
func (m *ReleaseDescriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseDescriptionsResponse) ProtoMessage()    {}
func (*ReleaseDescriptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{12}
}

func (m *ReleaseDescriptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseDescriptionsResponse.Unmarshal(m, b)
}
func (m *ReleaseDescriptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseDescriptionsResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseDescriptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseDescriptionsResponse.Merge(m, src)
}
func (m *ReleaseDescriptionsResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseDescriptionsResponse.Size(m)
}
func (m *ReleaseDescriptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseDescriptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseDescriptionsResponse proto.InternalMessageInfo

func (m *ReleaseDescriptionsResponse) GetReleaseDescriptions() []*ReleaseDescription {
	if m != nil {
		return m.ReleaseDescriptions
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
	proto.RegisterType((*ReleasesListResponse)(nil), "proto.ReleasesListResponse")
	proto.RegisterType((*ReleaseType)(nil), "proto.ReleaseType")
	proto.RegisterType((*ReleaseTypeRequest)(nil), "proto.ReleaseTypeRequest")
	proto.RegisterType((*ReleaseTypeResponse)(nil), "proto.ReleaseTypeResponse")
	proto.RegisterType((*ReleaseDescription)(nil), "proto.ReleaseDescription")
	proto.RegisterType((*ReleaseDescriptionsRequest)(nil), "proto.ReleaseDescriptionsRequest")
	proto.RegisterType((*ReleaseDescriptionsResponse)(nil), "proto.ReleaseDescriptionsResponse")
}

func init() { proto.RegisterFile("proto/release.proto", fileDescriptor_987cf5695bafb180) }

var fileDescriptor_987cf5695bafb180 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x4e, 0x1b, 0x3d,
	0x10, 0xd5, 0x66, 0x09, 0x81, 0x49, 0x08, 0x9f, 0x1c, 0x3e, 0xba, 0x04, 0x15, 0x52, 0xf7, 0x47,
	0xa1, 0x17, 0x54, 0x02, 0x21, 0x55, 0xa2, 0x37, 0x54, 0xf4, 0x22, 0x2a, 0x52, 0x55, 0x97, 0xcb,
	0x4a, 0x68, 0x89, 0x4d, 0xe4, 0x36, 0x59, 0x2f, 0x5e, 0xa7, 0x52, 0x1e, 0xa0, 0x7d, 0x8d, 0x3e,
	0x58, 0x5f, 0xa6, 0x5a, 0x7b, 0xbc, 0xd9, 0x0d, 0x49, 0x55, 0xf5, 0x2a, 0x9a, 0x73, 0xce, 0x9c,
	0x3d, 0x1e, 0x8f, 0x03, 0x9d, 0x54, 0x2b, 0xa3, 0x5e, 0x69, 0x31, 0x16, 0x71, 0x26, 0x8e, 0x6d,
	0x45, 0xea, 0xf6, 0xa7, 0x4b, 0x1c, 0x37, 0x54, 0x93, 0x89, 0x4a, 0x1c, 0x45, 0x7f, 0xd5, 0xa0,
	0xc1, 0x9c, 0x98, 0xb4, 0xa1, 0x26, 0x79, 0x14, 0xf4, 0x82, 0x7e, 0xc8, 0x6a, 0x92, 0x93, 0x1d,
	0xa8, 0x1b, 0x69, 0xc6, 0x22, 0xaa, 0xf5, 0x82, 0xfe, 0x26, 0x73, 0x05, 0xd9, 0x85, 0xf5, 0x54,
	0x65, 0x46, 0xe8, 0x28, 0xb4, 0x30, 0x56, 0xe4, 0x00, 0xe0, 0x36, 0x1e, 0x7e, 0x1d, 0x69, 0x35,
	0x4d, 0x78, 0xb4, 0x66, 0xb9, 0x12, 0x42, 0x9e, 0xc2, 0x16, 0xa6, 0xba, 0x51, 0x9a, 0x0b, 0x1d,
	0xd5, 0xed, 0x87, 0x5a, 0x08, 0x7e, 0xc8, 0x31, 0xf2, 0x1c, 0xda, 0x99, 0x89, 0xb5, 0x11, 0xfc,
	0x26, 0x96, 0x5a, 0x26, 0xa3, 0x68, 0xdd, 0x1a, 0x6d, 0x21, 0x7a, 0x61, 0x41, 0x27, 0x53, 0x69,
	0x3a, 0x97, 0x35, 0xbc, 0xcc, 0xa2, 0x28, 0x7b, 0x0c, 0x30, 0xd4, 0x22, 0xb6, 0x6e, 0x26, 0xea,
	0x58, 0xc9, 0x26, 0x22, 0x17, 0x26, 0xa7, 0xa7, 0x29, 0xf7, 0xf4, 0x8e, 0xa3, 0x11, 0xb9, 0x30,
	0x64, 0x0f, 0x36, 0xe2, 0x44, 0x4e, 0xc4, 0x8d, 0xe4, 0x51, 0xdf, 0x66, 0x6d, 0xd8, 0x7a, 0xc0,
	0xc9, 0x0b, 0xd8, 0xf6, 0x67, 0x31, 0xb3, 0xd4, 0x2a, 0x8e, 0xac, 0xc2, 0x1f, 0xf1, 0x7a, 0x96,
	0x8a, 0x01, 0xa7, 0xdf, 0x03, 0x68, 0xe1, 0x74, 0x3f, 0x4e, 0x85, 0x9e, 0x55, 0x3c, 0x83, 0xaa,
	0xe7, 0xca, 0x69, 0x8f, 0x44, 0xa2, 0x45, 0x16, 0x85, 0xbd, 0xb0, 0x1f, 0x32, 0xac, 0x72, 0x5c,
	0xdd, 0xdd, 0x65, 0xc2, 0x44, 0x07, 0xd6, 0x06, 0xab, 0xdc, 0x65, 0x2c, 0x27, 0xd2, 0x44, 0x87,
	0x16, 0x76, 0x05, 0xed, 0x41, 0x1b, 0x63, 0x30, 0x71, 0x3f, 0x15, 0x99, 0x59, 0xbc, 0x6b, 0xfa,
	0x06, 0xb6, 0x51, 0x91, 0x79, 0xc9, 0x11, 0xd4, 0xef, 0xf3, 0xd0, 0x56, 0xd5, 0x3c, 0xe9, 0xb8,
	0x8d, 0x39, 0x2e, 0x9f, 0x87, 0x39, 0x05, 0x3d, 0x2f, 0xba, 0x99, 0xc8, 0x52, 0x95, 0x64, 0x82,
	0xf4, 0xa1, 0x81, 0xb3, 0xc0, 0xfe, 0x76, 0xb5, 0x9f, 0x79, 0x9a, 0x5e, 0xc3, 0x7f, 0xf3, 0x4f,
	0x63, 0xf7, 0x4b, 0xd8, 0x40, 0x3a, 0x8b, 0x82, 0x5e, 0xb8, 0xa4, 0xbd, 0xe0, 0xf3, 0x23, 0x0f,
	0xd5, 0x34, 0x31, 0x76, 0x70, 0x21, 0x73, 0x05, 0x7d, 0x0b, 0x3b, 0xde, 0xf5, 0x4a, 0x66, 0xe6,
	0x5f, 0x9c, 0xe9, 0x29, 0x34, 0xd9, 0xfc, 0x3e, 0xff, 0xee, 0x7d, 0xd0, 0x67, 0x40, 0x4a, 0x4d,
	0xab, 0xe6, 0x7d, 0x05, 0x9d, 0x8a, 0x0a, 0xd3, 0x9d, 0x41, 0xab, 0xbc, 0x58, 0x38, 0x3a, 0x52,
	0x4d, 0x68, 0x3b, 0x9a, 0xa5, 0x4d, 0xa3, 0x3f, 0x82, 0xe2, 0xa3, 0x97, 0x22, 0x1b, 0x6a, 0x99,
	0x1a, 0xa9, 0x92, 0x07, 0x81, 0x7b, 0xd0, 0xe4, 0x73, 0x1a, 0x63, 0x97, 0xa1, 0xfc, 0x49, 0xf8,
	0xef, 0x4b, 0x6e, 0x1f, 0x78, 0xc8, 0x36, 0x11, 0x19, 0x70, 0x72, 0x08, 0xcd, 0x71, 0x9c, 0x8c,
	0xa6, 0xf1, 0xc8, 0xf2, 0x6b, 0x96, 0x07, 0x0f, 0x0d, 0x38, 0x3d, 0x87, 0xee, 0xc3, 0x1c, 0xc5,
	0x46, 0x55, 0xdd, 0x83, 0x05, 0x77, 0xfa, 0x05, 0xf6, 0x97, 0x36, 0xe3, 0x6c, 0xde, 0x17, 0x23,
	0x2b, 0xd3, 0x78, 0x89, 0x7b, 0xd5, 0x11, 0x95, 0x14, 0x6c, 0x59, 0xd7, 0xc9, 0xcf, 0xb0, 0x78,
	0x12, 0x9f, 0x84, 0xfe, 0x26, 0x87, 0x82, 0xbc, 0x9e, 0xff, 0x13, 0xfe, 0xbf, 0xb0, 0x12, 0x2e,
	0x7f, 0x77, 0x77, 0x11, 0xc6, 0x64, 0xe7, 0xb0, 0xe1, 0x77, 0x8d, 0x2c, 0x68, 0xfc, 0xd9, 0xbb,
	0x8f, 0x1e, 0xe0, 0x45, 0x73, 0xdb, 0xfd, 0x5d, 0x15, 0x16, 0x2d, 0x94, 0xbe, 0x9b, 0xa4, 0x66,
	0xd6, 0xdd, 0x5f, 0x68, 0xac, 0x6c, 0xf3, 0x19, 0x6c, 0xb1, 0x38, 0xe1, 0x6a, 0xe2, 0x93, 0x57,
	0x7b, 0x57, 0x05, 0xbe, 0xac, 0x2e, 0xf6, 0xde, 0x92, 0xfd, 0xc2, 0xd8, 0xdd, 0x65, 0x14, 0xba,
	0x7c, 0x5e, 0x7a, 0x21, 0xe4, 0xc9, 0xca, 0xab, 0x28, 0x86, 0x41, 0xff, 0x24, 0x71, 0xee, 0xb7,
	0xeb, 0x56, 0x72, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x63, 0xec, 0xda, 0xa0, 0xd2, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
	Releases(ctx context.Context, in *ReleasesRequest, opts ...grpc.CallOption) (*ReleasesResponse, error)
	AiringReleases(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReleasesListResponse, error)
	RandomRelease(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReleaseResponse, error)
	ReleaseType(ctx context.Context, in *ReleaseTypeRequest, opts ...grpc.CallOption) (*ReleaseTypeResponse, error)
	ReleaseDescriptions(ctx context.Context, in *ReleaseDescriptionsRequest, opts ...grpc.CallOption) (*ReleaseDescriptionsResponse, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
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

func (c *releaseServiceClient) AiringReleases(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReleasesListResponse, error) {
	out := new(ReleasesListResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/AiringReleases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) RandomRelease(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/RandomRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) ReleaseType(ctx context.Context, in *ReleaseTypeRequest, opts ...grpc.CallOption) (*ReleaseTypeResponse, error) {
	out := new(ReleaseTypeResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/ReleaseType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) ReleaseDescriptions(ctx context.Context, in *ReleaseDescriptionsRequest, opts ...grpc.CallOption) (*ReleaseDescriptionsResponse, error) {
	out := new(ReleaseDescriptionsResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/ReleaseDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
type ReleaseServiceServer interface {
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	Releases(context.Context, *ReleasesRequest) (*ReleasesResponse, error)
	AiringReleases(context.Context, *Empty) (*ReleasesListResponse, error)
	RandomRelease(context.Context, *Empty) (*ReleaseResponse, error)
	ReleaseType(context.Context, *ReleaseTypeRequest) (*ReleaseTypeResponse, error)
	ReleaseDescriptions(context.Context, *ReleaseDescriptionsRequest) (*ReleaseDescriptionsResponse, error)
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
func (*UnimplementedReleaseServiceServer) AiringReleases(ctx context.Context, req *Empty) (*ReleasesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AiringReleases not implemented")
}
func (*UnimplementedReleaseServiceServer) RandomRelease(ctx context.Context, req *Empty) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomRelease not implemented")
}
func (*UnimplementedReleaseServiceServer) ReleaseType(ctx context.Context, req *ReleaseTypeRequest) (*ReleaseTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseType not implemented")
}
func (*UnimplementedReleaseServiceServer) ReleaseDescriptions(ctx context.Context, req *ReleaseDescriptionsRequest) (*ReleaseDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseDescriptions not implemented")
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

func _ReleaseService_AiringReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).AiringReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/AiringReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).AiringReleases(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_RandomRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).RandomRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/RandomRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).RandomRelease(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_ReleaseType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).ReleaseType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/ReleaseType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).ReleaseType(ctx, req.(*ReleaseTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_ReleaseDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).ReleaseDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/ReleaseDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).ReleaseDescriptions(ctx, req.(*ReleaseDescriptionsRequest))
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
		{
			MethodName: "AiringReleases",
			Handler:    _ReleaseService_AiringReleases_Handler,
		},
		{
			MethodName: "RandomRelease",
			Handler:    _ReleaseService_RandomRelease_Handler,
		},
		{
			MethodName: "ReleaseType",
			Handler:    _ReleaseService_ReleaseType_Handler,
		},
		{
			MethodName: "ReleaseDescriptions",
			Handler:    _ReleaseService_ReleaseDescriptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/release.proto",
}
