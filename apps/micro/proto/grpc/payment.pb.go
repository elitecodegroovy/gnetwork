// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc/payment.proto

package payment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PayReq struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayReq) Reset()         { *m = PayReq{} }
func (m *PayReq) String() string { return proto.CompactTextString(m) }
func (*PayReq) ProtoMessage()    {}
func (*PayReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78f6f8d5031611e, []int{0}
}

func (m *PayReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayReq.Unmarshal(m, b)
}
func (m *PayReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayReq.Marshal(b, m, deterministic)
}
func (m *PayReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayReq.Merge(m, src)
}
func (m *PayReq) XXX_Size() int {
	return xxx_messageInfo_PayReq.Size(m)
}
func (m *PayReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayReq proto.InternalMessageInfo

func (m *PayReq) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PayReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type PayResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Success              string   `protobuf:"bytes,2,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayResp) Reset()         { *m = PayResp{} }
func (m *PayResp) String() string { return proto.CompactTextString(m) }
func (*PayResp) ProtoMessage()    {}
func (*PayResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78f6f8d5031611e, []int{1}
}

func (m *PayResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayResp.Unmarshal(m, b)
}
func (m *PayResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayResp.Marshal(b, m, deterministic)
}
func (m *PayResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayResp.Merge(m, src)
}
func (m *PayResp) XXX_Size() int {
	return xxx_messageInfo_PayResp.Size(m)
}
func (m *PayResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayResp.DiscardUnknown(m)
}

var xxx_messageInfo_PayResp proto.InternalMessageInfo

func (m *PayResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PayResp) GetSuccess() string {
	if m != nil {
		return m.Success
	}
	return ""
}

func (m *PayResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PayReq)(nil), "PayReq")
	proto.RegisterType((*PayResp)(nil), "PayResp")
}

func init() { proto.RegisterFile("proto/grpc/payment.proto", fileDescriptor_d78f6f8d5031611e) }

var fileDescriptor_d78f6f8d5031611e = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x03,
	0x0b, 0x49, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6,
	0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x64, 0x95, 0x6c, 0xb8, 0xd8, 0x02,
	0x12, 0x2b, 0x83, 0x52, 0x0b, 0x85, 0x24, 0xb8, 0xd8, 0xf3, 0x8b, 0x52, 0x52, 0x8b, 0x3c, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x90, 0x4c, 0x62, 0x72, 0x72, 0x7e, 0x69,
	0x5e, 0x89, 0x04, 0x13, 0x44, 0x06, 0xca, 0x55, 0xf2, 0xe4, 0x62, 0x07, 0xeb, 0x2e, 0x2e, 0x10,
	0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x05, 0xeb, 0x65, 0x0d, 0x02, 0xb3, 0x41, 0x1a, 0x8b,
	0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x61, 0x1a, 0xa1, 0x5c, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2,
	0x74, 0x09, 0x66, 0xb0, 0x28, 0x88, 0x69, 0xe4, 0xc8, 0xc5, 0x17, 0x00, 0x71, 0x77, 0x70, 0x6a,
	0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3e, 0x17, 0x73, 0x40, 0x62, 0xa5, 0x10, 0xbb, 0x1e, 0xc4,
	0x81, 0x52, 0x1c, 0x7a, 0x50, 0xbb, 0x94, 0x44, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0xaf, 0xc4,
	0xa5, 0x0f, 0x76, 0x22, 0xc8, 0xc3, 0x56, 0x8c, 0x5a, 0x49, 0x6c, 0x60, 0x2f, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x83, 0xe2, 0xd3, 0x38, 0x0c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error)
}

type paymentServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error) {
	out := new(PayResp)
	err := c.cc.Invoke(ctx, "/PaymentService/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	Pay(context.Context, *PayReq) (*PayResp, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) Pay(ctx context.Context, req *PayReq) (*PayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentService/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Pay(ctx, req.(*PayReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _PaymentService_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc/payment.proto",
}
