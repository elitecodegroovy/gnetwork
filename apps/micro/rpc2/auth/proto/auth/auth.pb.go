// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package mu_micro_book_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
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

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type AuthRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AuthRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AuthResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.auth.Error")
	proto.RegisterType((*AuthRequest)(nil), "mu.micro.book.srv.auth.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "mu.micro.book.srv.auth.AuthResponse")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x36, 0x6d, 0x9d, 0x0a, 0xc2, 0xa0, 0x25, 0x14, 0x84, 0x12, 0x3d, 0xf4, 0xb4,
	0x42, 0xf3, 0x0b, 0x0a, 0x7a, 0xf0, 0xa0, 0x87, 0x55, 0xf1, 0x26, 0xa4, 0xc9, 0x40, 0x43, 0x9a,
	0x6e, 0xdd, 0xd9, 0xed, 0xaf, 0xf4, 0x47, 0xc9, 0x4e, 0xa2, 0x78, 0xa8, 0xde, 0x7a, 0x09, 0xf3,
	0x31, 0x2f, 0xef, 0xed, 0xdb, 0x85, 0xcb, 0xad, 0x35, 0xce, 0xdc, 0xe6, 0xde, 0xad, 0xe4, 0xa3,
	0x84, 0x71, 0xdc, 0x78, 0xd5, 0x54, 0x85, 0x35, 0x6a, 0x69, 0x4c, 0xad, 0xd8, 0xee, 0x54, 0xd8,
	0xa6, 0x19, 0xc4, 0xf7, 0xd6, 0x1a, 0x8b, 0x08, 0xbd, 0xc2, 0x94, 0x94, 0x44, 0xd3, 0x68, 0x16,
	0x6b, 0x99, 0x71, 0x0c, 0xfd, 0x92, 0x5c, 0x5e, 0xad, 0x93, 0xe3, 0x69, 0x34, 0x3b, 0xd5, 0x1d,
	0xa5, 0x6f, 0x30, 0x5a, 0x78, 0xb7, 0xd2, 0xf4, 0xe1, 0x89, 0x5d, 0x90, 0x79, 0x26, 0xfb, 0x50,
	0xca, 0xcf, 0x3d, 0xdd, 0x11, 0x4e, 0x60, 0x18, 0xa6, 0xa7, 0xbc, 0xa1, 0xce, 0xe0, 0x87, 0xf1,
	0x02, 0x62, 0x67, 0x6a, 0xda, 0x24, 0x27, 0xb2, 0x68, 0x21, 0x65, 0x38, 0x6b, 0x8d, 0x79, 0x6b,
	0x36, 0x4c, 0x98, 0xc0, 0x80, 0x7d, 0x51, 0x10, 0xb3, 0x58, 0x0f, 0xf5, 0x37, 0x62, 0x06, 0x31,
	0x85, 0x73, 0x8b, 0xf1, 0x68, 0x7e, 0xa5, 0xf6, 0xf7, 0x53, 0x52, 0x4e, 0xb7, 0xda, 0xfd, 0xa1,
	0xf3, 0xcf, 0x08, 0x06, 0xcf, 0x64, 0x77, 0x55, 0x41, 0xf8, 0x0e, 0xe7, 0x8f, 0x79, 0x4d, 0x0b,
	0x09, 0x79, 0x09, 0x6b, 0xbc, 0xfe, 0xcb, 0xfa, 0xd7, 0x15, 0x4c, 0x6e, 0xfe, 0x17, 0xb5, 0x75,
	0xd2, 0x23, 0xcc, 0x01, 0xef, 0x68, 0xfd, 0xca, 0x64, 0x0f, 0x15, 0xb1, 0xec, 0xcb, 0x83, 0x67,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0xac, 0x43, 0xfa, 0x09, 0x02, 0x00, 0x00,
}
