// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type ListPostsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsReq) Reset()         { *m = ListPostsReq{} }
func (m *ListPostsReq) String() string { return proto.CompactTextString(m) }
func (*ListPostsReq) ProtoMessage()    {}
func (*ListPostsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ListPostsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsReq.Unmarshal(m, b)
}
func (m *ListPostsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsReq.Marshal(b, m, deterministic)
}
func (m *ListPostsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsReq.Merge(m, src)
}
func (m *ListPostsReq) XXX_Size() int {
	return xxx_messageInfo_ListPostsReq.Size(m)
}
func (m *ListPostsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsReq proto.InternalMessageInfo

type ListPostsRes struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRes) Reset()         { *m = ListPostsRes{} }
func (m *ListPostsRes) String() string { return proto.CompactTextString(m) }
func (*ListPostsRes) ProtoMessage()    {}
func (*ListPostsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ListPostsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRes.Unmarshal(m, b)
}
func (m *ListPostsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRes.Marshal(b, m, deterministic)
}
func (m *ListPostsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRes.Merge(m, src)
}
func (m *ListPostsRes) XXX_Size() int {
	return xxx_messageInfo_ListPostsRes.Size(m)
}
func (m *ListPostsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRes proto.InternalMessageInfo

func (m *ListPostsRes) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type CreatePostReq struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostReq) Reset()         { *m = CreatePostReq{} }
func (m *CreatePostReq) String() string { return proto.CompactTextString(m) }
func (*CreatePostReq) ProtoMessage()    {}
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *CreatePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostReq.Unmarshal(m, b)
}
func (m *CreatePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostReq.Marshal(b, m, deterministic)
}
func (m *CreatePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostReq.Merge(m, src)
}
func (m *CreatePostReq) XXX_Size() int {
	return xxx_messageInfo_CreatePostReq.Size(m)
}
func (m *CreatePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostReq proto.InternalMessageInfo

func (m *CreatePostReq) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostRes struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRes) Reset()         { *m = CreatePostRes{} }
func (m *CreatePostRes) String() string { return proto.CompactTextString(m) }
func (*CreatePostRes) ProtoMessage()    {}
func (*CreatePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *CreatePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRes.Unmarshal(m, b)
}
func (m *CreatePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRes.Marshal(b, m, deterministic)
}
func (m *CreatePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRes.Merge(m, src)
}
func (m *CreatePostRes) XXX_Size() int {
	return xxx_messageInfo_CreatePostRes.Size(m)
}
func (m *CreatePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRes proto.InternalMessageInfo

func (m *CreatePostRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdatePostReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Post                 *Post    `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostReq) Reset()         { *m = UpdatePostReq{} }
func (m *UpdatePostReq) String() string { return proto.CompactTextString(m) }
func (*UpdatePostReq) ProtoMessage()    {}
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UpdatePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostReq.Unmarshal(m, b)
}
func (m *UpdatePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostReq.Marshal(b, m, deterministic)
}
func (m *UpdatePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostReq.Merge(m, src)
}
func (m *UpdatePostReq) XXX_Size() int {
	return xxx_messageInfo_UpdatePostReq.Size(m)
}
func (m *UpdatePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostReq proto.InternalMessageInfo

func (m *UpdatePostReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePostReq) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type UpdatePostRes struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRes) Reset()         { *m = UpdatePostRes{} }
func (m *UpdatePostRes) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRes) ProtoMessage()    {}
func (*UpdatePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *UpdatePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRes.Unmarshal(m, b)
}
func (m *UpdatePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRes.Marshal(b, m, deterministic)
}
func (m *UpdatePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRes.Merge(m, src)
}
func (m *UpdatePostRes) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRes.Size(m)
}
func (m *UpdatePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRes proto.InternalMessageInfo

func (m *UpdatePostRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeletePostReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostReq) Reset()         { *m = DeletePostReq{} }
func (m *DeletePostReq) String() string { return proto.CompactTextString(m) }
func (*DeletePostReq) ProtoMessage()    {}
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *DeletePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostReq.Unmarshal(m, b)
}
func (m *DeletePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostReq.Marshal(b, m, deterministic)
}
func (m *DeletePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostReq.Merge(m, src)
}
func (m *DeletePostReq) XXX_Size() int {
	return xxx_messageInfo_DeletePostReq.Size(m)
}
func (m *DeletePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostReq proto.InternalMessageInfo

func (m *DeletePostReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeletePostRes struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRes) Reset()         { *m = DeletePostRes{} }
func (m *DeletePostRes) String() string { return proto.CompactTextString(m) }
func (*DeletePostRes) ProtoMessage()    {}
func (*DeletePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *DeletePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRes.Unmarshal(m, b)
}
func (m *DeletePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRes.Marshal(b, m, deterministic)
}
func (m *DeletePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRes.Merge(m, src)
}
func (m *DeletePostRes) XXX_Size() int {
	return xxx_messageInfo_DeletePostRes.Size(m)
}
func (m *DeletePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRes proto.InternalMessageInfo

func (m *DeletePostRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Response struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{11}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "proto.Post")
	proto.RegisterType((*ListPostsReq)(nil), "proto.ListPostsReq")
	proto.RegisterType((*ListPostsRes)(nil), "proto.ListPostsRes")
	proto.RegisterType((*CreatePostReq)(nil), "proto.CreatePostReq")
	proto.RegisterType((*CreatePostRes)(nil), "proto.CreatePostRes")
	proto.RegisterType((*UpdatePostReq)(nil), "proto.UpdatePostReq")
	proto.RegisterType((*UpdatePostRes)(nil), "proto.UpdatePostRes")
	proto.RegisterType((*DeletePostReq)(nil), "proto.DeletePostReq")
	proto.RegisterType((*DeletePostRes)(nil), "proto.DeletePostRes")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*Error)(nil), "proto.Error")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xf9, 0x35, 0xed, 0xaf, 0xd3, 0xa6, 0x87, 0x51, 0x24, 0xe4, 0xd2, 0xba, 0xa7, 0x2a,
	0x58, 0xb4, 0x52, 0xf0, 0x28, 0xf8, 0xe7, 0xe4, 0x41, 0x56, 0xfc, 0x00, 0x6d, 0x33, 0x48, 0xa0,
	0xba, 0xe9, 0x4e, 0x2a, 0xf8, 0xb9, 0xfd, 0x02, 0xb2, 0x9b, 0x6d, 0x9b, 0x94, 0x18, 0x3c, 0xed,
	0xcc, 0x9b, 0x99, 0xf7, 0x66, 0x78, 0x0b, 0x21, 0x93, 0xfe, 0x4c, 0x97, 0x34, 0xc9, 0xb4, 0xca,
	0x15, 0x06, 0xf6, 0x11, 0xe7, 0xd0, 0x7a, 0x56, 0x9c, 0xe3, 0x00, 0xfc, 0x34, 0x89, 0xbc, 0x91,
	0x37, 0xee, 0x4a, 0x3f, 0x4d, 0x10, 0xa1, 0xb5, 0x50, 0xc9, 0x57, 0xe4, 0x5b, 0xc4, 0xc6, 0x62,
	0x00, 0xfd, 0xa7, 0x94, 0x73, 0xd3, 0xcf, 0x92, 0xd6, 0xe2, 0xaa, 0x92, 0x33, 0x9e, 0x42, 0x90,
	0x99, 0x38, 0xf2, 0x46, 0xff, 0xc6, 0xbd, 0x69, 0xaf, 0x50, 0x9a, 0x98, 0xba, 0x2c, 0x2a, 0xe2,
	0x12, 0xc2, 0x3b, 0x4d, 0xf3, 0x9c, 0x2c, 0x48, 0x6b, 0x1c, 0x42, 0xcb, 0x54, 0xac, 0xf2, 0xc1,
	0x88, 0x2d, 0x88, 0xb3, 0xea, 0x04, 0x63, 0x04, 0x9d, 0x77, 0x62, 0x9e, 0xbf, 0x91, 0x5b, 0x77,
	0x9b, 0x8a, 0x5b, 0x08, 0x5f, 0xb3, 0xa4, 0x44, 0x7e, 0x78, 0xd4, 0x56, 0xcc, 0x6f, 0x10, 0x2b,
	0x33, 0x34, 0x89, 0x0d, 0x21, 0xbc, 0xa7, 0x15, 0xfd, 0x2a, 0x66, 0xb8, 0xca, 0x0d, 0x4d, 0x5c,
	0x17, 0xd0, 0x91, 0xb4, 0xde, 0xd0, 0x1f, 0x7d, 0x78, 0x84, 0xff, 0x92, 0x38, 0x53, 0x1f, 0x4c,
	0x78, 0x02, 0x6d, 0x4d, 0xbc, 0x59, 0xe5, 0x6e, 0xc6, 0x65, 0x28, 0x20, 0x20, 0xad, 0x95, 0x76,
	0xb7, 0xf6, 0xdd, 0xad, 0x0f, 0x06, 0x93, 0x45, 0x49, 0xcc, 0x20, 0xb0, 0xb9, 0x11, 0x59, 0xaa,
	0x64, 0xbb, 0x96, 0x8d, 0xcb, 0xdb, 0xfa, 0x95, 0x6d, 0xa7, 0xdf, 0x1e, 0xf4, 0xad, 0xe7, 0x2f,
	0xc5, 0x87, 0xc2, 0x1b, 0x80, 0xbd, 0x45, 0x78, 0xec, 0xa4, 0x2a, 0x3e, 0xc7, 0x75, 0x28, 0xe3,
	0x0c, 0xba, 0x92, 0xe6, 0x89, 0x65, 0xc3, 0x23, 0xd7, 0x52, 0xfe, 0x63, 0x71, 0x0d, 0xc8, 0x46,
	0x70, 0x6f, 0xd3, 0x4e, 0xb0, 0xe2, 0x7d, 0x5c, 0x87, 0xda, 0xc9, 0xbd, 0x29, 0xbb, 0xc9, 0x8a,
	0x91, 0x71, 0x1d, 0xca, 0x8b, 0xb6, 0x05, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xb3,
	0x2e, 0xb2, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostsServiceClient is the client API for PostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostsServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRes, error)
	ReadPosts(ctx context.Context, in *ListPostsReq, opts ...grpc.CallOption) (*ListPostsRes, error)
	UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error)
	DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error)
}

type postsServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostsServiceClient(cc *grpc.ClientConn) PostsServiceClient {
	return &postsServiceClient{cc}
}

func (c *postsServiceClient) CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRes, error) {
	out := new(CreatePostRes)
	err := c.cc.Invoke(ctx, "/proto.PostsService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) ReadPosts(ctx context.Context, in *ListPostsReq, opts ...grpc.CallOption) (*ListPostsRes, error) {
	out := new(ListPostsRes)
	err := c.cc.Invoke(ctx, "/proto.PostsService/ReadPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error) {
	out := new(UpdatePostRes)
	err := c.cc.Invoke(ctx, "/proto.PostsService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error) {
	out := new(DeletePostRes)
	err := c.cc.Invoke(ctx, "/proto.PostsService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServiceServer is the server API for PostsService service.
type PostsServiceServer interface {
	CreatePost(context.Context, *CreatePostReq) (*CreatePostRes, error)
	ReadPosts(context.Context, *ListPostsReq) (*ListPostsRes, error)
	UpdatePost(context.Context, *UpdatePostReq) (*UpdatePostRes, error)
	DeletePost(context.Context, *DeletePostReq) (*DeletePostRes, error)
}

// UnimplementedPostsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostsServiceServer struct {
}

func (*UnimplementedPostsServiceServer) CreatePost(ctx context.Context, req *CreatePostReq) (*CreatePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedPostsServiceServer) ReadPosts(ctx context.Context, req *ListPostsReq) (*ListPostsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosts not implemented")
}
func (*UnimplementedPostsServiceServer) UpdatePost(ctx context.Context, req *UpdatePostReq) (*UpdatePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (*UnimplementedPostsServiceServer) DeletePost(ctx context.Context, req *DeletePostReq) (*DeletePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}

func RegisterPostsServiceServer(s *grpc.Server, srv PostsServiceServer) {
	s.RegisterService(&_PostsService_serviceDesc, srv)
}

func _PostsService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostsService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).CreatePost(ctx, req.(*CreatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_ReadPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).ReadPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostsService/ReadPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).ReadPosts(ctx, req.(*ListPostsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostsService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).UpdatePost(ctx, req.(*UpdatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostsService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).DeletePost(ctx, req.(*DeletePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostsService",
	HandlerType: (*PostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostsService_CreatePost_Handler,
		},
		{
			MethodName: "ReadPosts",
			Handler:    _PostsService_ReadPosts_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostsService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostsService_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
