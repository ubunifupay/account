// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package greeting

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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Created              string   `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Account) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

type AccountRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{1}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

type AccountResponse struct {
	Account              []*Account `protobuf:"bytes,1,rep,name=Account,proto3" json:"Account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{2}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetAccount() []*Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "greeting.Account")
	proto.RegisterType((*AccountRequest)(nil), "greeting.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "greeting.AccountResponse")
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_477cbf5ae5b46edf) }

var fileDescriptor_477cbf5ae5b46edf = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4a, 0xc4, 0x30,
	0x14, 0xa6, 0x19, 0x70, 0x9c, 0x14, 0x46, 0x27, 0x6e, 0xe2, 0xe0, 0xa2, 0x64, 0x35, 0x20, 0xb4,
	0x50, 0xf7, 0x82, 0x47, 0xb0, 0xe2, 0x01, 0x62, 0xf2, 0x08, 0x01, 0xc9, 0x8b, 0x49, 0xea, 0x01,
	0xbc, 0x82, 0x47, 0xf3, 0x0a, 0x1e, 0x44, 0x4c, 0x4c, 0x15, 0x9c, 0xe5, 0x7b, 0xdf, 0x2f, 0x1f,
	0xbd, 0xf0, 0x01, 0x13, 0x0e, 0x52, 0x29, 0x9c, 0x5d, 0xea, 0xf3, 0xc5, 0x4e, 0x4d, 0x00, 0x48,
	0xd6, 0x99, 0xfd, 0x95, 0x41, 0x34, 0xcf, 0x30, 0x48, 0x6f, 0x07, 0xe9, 0x1c, 0x26, 0x99, 0x2c,
	0xba, 0x58, 0x78, 0xe2, 0x91, 0xae, 0xef, 0x8a, 0x90, 0x6d, 0x29, 0xb1, 0x9a, 0x37, 0x5d, 0x73,
	0xd8, 0x4c, 0xc4, 0x6a, 0xd6, 0xd1, 0x56, 0x43, 0x54, 0xc1, 0xfa, 0x6f, 0x01, 0x27, 0x19, 0xf8,
	0xfb, 0x62, 0x9c, 0xae, 0x55, 0x00, 0x99, 0x40, 0xf3, 0x55, 0x46, 0xeb, 0x29, 0xce, 0xe9, 0xf6,
	0xc7, 0x76, 0x82, 0x97, 0x19, 0x62, 0x12, 0xb7, 0xf4, 0x6c, 0xf9, 0x44, 0x8f, 0x2e, 0x02, 0xbb,
	0x5e, 0xb2, 0x79, 0xd3, 0xad, 0x0e, 0xed, 0xb8, 0xeb, 0x6b, 0xeb, 0xbe, 0x72, 0x2b, 0x63, 0x54,
	0x8b, 0xe3, 0x03, 0x84, 0x57, 0xab, 0x80, 0xdd, 0xff, 0x56, 0xe7, 0xff, 0x85, 0x25, 0x76, 0x7f,
	0x79, 0x04, 0x29, 0xf1, 0x62, 0xf7, 0xf6, 0xf1, 0xf9, 0x4e, 0x5a, 0xb6, 0xa9, 0xd3, 0xc5, 0xa7,
	0x93, 0x3c, 0xca, 0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xe6, 0x31, 0x27, 0x53, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/greeting.AccountService/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Account(context.Context, *AccountRequest) (*AccountResponse, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) Account(ctx context.Context, req *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeting.AccountService/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Account(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeting.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _AccountService_Account_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account.proto",
}
