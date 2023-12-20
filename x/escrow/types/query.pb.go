// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reapchain/escrow/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/reapchain/cosmos-sdk/types"
	query "github.com/reapchain/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryRegisteredDenomsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRegisteredDenomsRequest) Reset()         { *m = QueryRegisteredDenomsRequest{} }
func (m *QueryRegisteredDenomsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRegisteredDenomsRequest) ProtoMessage()    {}
func (*QueryRegisteredDenomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{0}
}
func (m *QueryRegisteredDenomsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRegisteredDenomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRegisteredDenomsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRegisteredDenomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRegisteredDenomsRequest.Merge(m, src)
}
func (m *QueryRegisteredDenomsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRegisteredDenomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRegisteredDenomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRegisteredDenomsRequest proto.InternalMessageInfo

func (m *QueryRegisteredDenomsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryRegisteredDenomsResponse struct {
	RegisteredDenoms []RegisteredDenom `protobuf:"bytes,1,rep,name=registered_denoms,json=registeredDenoms,proto3" json:"registered_denoms"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRegisteredDenomsResponse) Reset()         { *m = QueryRegisteredDenomsResponse{} }
func (m *QueryRegisteredDenomsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRegisteredDenomsResponse) ProtoMessage()    {}
func (*QueryRegisteredDenomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{1}
}
func (m *QueryRegisteredDenomsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRegisteredDenomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRegisteredDenomsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRegisteredDenomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRegisteredDenomsResponse.Merge(m, src)
}
func (m *QueryRegisteredDenomsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRegisteredDenomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRegisteredDenomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRegisteredDenomsResponse proto.InternalMessageInfo

func (m *QueryRegisteredDenomsResponse) GetRegisteredDenoms() []RegisteredDenom {
	if m != nil {
		return m.RegisteredDenoms
	}
	return nil
}

func (m *QueryRegisteredDenomsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{2}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{3}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryEscrowPoolBalanceRequest struct {
	// address is the address to query balances for.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryEscrowPoolBalanceRequest) Reset()         { *m = QueryEscrowPoolBalanceRequest{} }
func (m *QueryEscrowPoolBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowPoolBalanceRequest) ProtoMessage()    {}
func (*QueryEscrowPoolBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{4}
}
func (m *QueryEscrowPoolBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowPoolBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowPoolBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowPoolBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowPoolBalanceRequest.Merge(m, src)
}
func (m *QueryEscrowPoolBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowPoolBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowPoolBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowPoolBalanceRequest proto.InternalMessageInfo

type QueryEscrowPoolBalanceResponse struct {
	// balances is the balances of all the coins.
	EscrowPoolBalance types.Coin `protobuf:"bytes,1,opt,name=escrow_pool_balance,json=escrowPoolBalance,proto3" json:"escrow_pool_balance"`
}

func (m *QueryEscrowPoolBalanceResponse) Reset()         { *m = QueryEscrowPoolBalanceResponse{} }
func (m *QueryEscrowPoolBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowPoolBalanceResponse) ProtoMessage()    {}
func (*QueryEscrowPoolBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a076ad32d675ac, []int{5}
}
func (m *QueryEscrowPoolBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowPoolBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowPoolBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowPoolBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowPoolBalanceResponse.Merge(m, src)
}
func (m *QueryEscrowPoolBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowPoolBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowPoolBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowPoolBalanceResponse proto.InternalMessageInfo

func (m *QueryEscrowPoolBalanceResponse) GetEscrowPoolBalance() types.Coin {
	if m != nil {
		return m.EscrowPoolBalance
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*QueryRegisteredDenomsRequest)(nil), "reapchain.escrow.v1.QueryRegisteredDenomsRequest")
	proto.RegisterType((*QueryRegisteredDenomsResponse)(nil), "reapchain.escrow.v1.QueryRegisteredDenomsResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "reapchain.escrow.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "reapchain.escrow.v1.QueryParamsResponse")
	proto.RegisterType((*QueryEscrowPoolBalanceRequest)(nil), "reapchain.escrow.v1.QueryEscrowPoolBalanceRequest")
	proto.RegisterType((*QueryEscrowPoolBalanceResponse)(nil), "reapchain.escrow.v1.QueryEscrowPoolBalanceResponse")
}

func init() { proto.RegisterFile("reapchain/escrow/v1/query.proto", fileDescriptor_58a076ad32d675ac) }

var fileDescriptor_58a076ad32d675ac = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0x5b, 0x1a, 0xc1, 0x75, 0x69, 0x2e, 0x19, 0x4a, 0xda, 0x3a, 0xc1, 0xa0, 0x36, 0x62,
	0xb8, 0x6b, 0xd2, 0xa5, 0xb0, 0x20, 0x85, 0x7f, 0x12, 0x0b, 0x21, 0x0b, 0x12, 0x4b, 0x74, 0x71,
	0x0f, 0xd7, 0x52, 0x72, 0xcf, 0xf1, 0x39, 0x81, 0x0a, 0x21, 0x21, 0x26, 0x46, 0x24, 0xbe, 0x40,
	0x57, 0xbe, 0x00, 0x3b, 0x5b, 0xc7, 0x4a, 0x2c, 0x4c, 0x08, 0x25, 0x0c, 0x6c, 0x7c, 0x05, 0x94,
	0xbb, 0x73, 0x1a, 0xa7, 0x4e, 0xab, 0x6e, 0x97, 0xbb, 0xdf, 0x7b, 0xbf, 0x3f, 0xef, 0xc5, 0xa8,
	0x1c, 0x71, 0x16, 0x7a, 0x87, 0x2c, 0x10, 0x94, 0x4b, 0x2f, 0x82, 0x37, 0x74, 0x58, 0xa3, 0xfd,
	0x01, 0x8f, 0x8e, 0x48, 0x18, 0x41, 0x0c, 0xb8, 0x30, 0x05, 0x10, 0x0d, 0x20, 0xc3, 0x5a, 0xe9,
	0xae, 0x07, 0xb2, 0x07, 0x92, 0x76, 0x98, 0xe4, 0x1a, 0x4d, 0x87, 0xb5, 0x0e, 0x8f, 0x59, 0x8d,
	0x86, 0xcc, 0x0f, 0x04, 0x8b, 0x03, 0x10, 0xba, 0x41, 0xe9, 0x56, 0x16, 0x83, 0xcf, 0x05, 0x97,
	0x81, 0x34, 0x90, 0x4a, 0x16, 0xc4, 0xb0, 0x69, 0x84, 0x33, 0x4b, 0x98, 0x50, 0x79, 0x10, 0x24,
	0x24, 0x9b, 0x3e, 0x80, 0xdf, 0xe5, 0x94, 0x85, 0x01, 0x65, 0x42, 0x40, 0xac, 0x14, 0x24, 0xfd,
	0x8b, 0x3e, 0xf8, 0xa0, 0x8e, 0x74, 0x72, 0xd2, 0xb7, 0xee, 0x6b, 0xb4, 0xf9, 0x62, 0x22, 0xbd,
	0xc5, 0xfd, 0x40, 0xc6, 0x3c, 0xe2, 0x07, 0x8f, 0xb8, 0x80, 0x9e, 0x6c, 0xf1, 0xfe, 0x80, 0xcb,
	0x18, 0x3f, 0x41, 0xe8, 0xcc, 0xcc, 0xba, 0x5d, 0xb1, 0xab, 0xab, 0xf5, 0x6d, 0xa2, 0x85, 0x90,
	0x89, 0x10, 0xa2, 0x73, 0x32, 0x72, 0x48, 0x93, 0xf9, 0xdc, 0xd4, 0xb6, 0x66, 0x2a, 0xdd, 0xef,
	0x36, 0xda, 0x5a, 0x40, 0x24, 0x43, 0x10, 0x92, 0xe3, 0x97, 0x28, 0x1f, 0x4d, 0xdf, 0xda, 0x07,
	0xea, 0x71, 0xdd, 0xae, 0x2c, 0x57, 0x57, 0xeb, 0x77, 0x48, 0x46, 0xfe, 0x64, 0xae, 0x53, 0xe3,
	0xda, 0xc9, 0xaf, 0xb2, 0xd5, 0x5a, 0x8b, 0xe6, 0x08, 0xf0, 0xd3, 0x94, 0x85, 0x25, 0x65, 0x61,
	0xe7, 0x52, 0x0b, 0x5a, 0x55, 0xca, 0x43, 0x11, 0x61, 0x65, 0xa1, 0xc9, 0x22, 0x36, 0x4d, 0xc8,
	0x6d, 0xa2, 0x42, 0xea, 0xd6, 0xd8, 0xb9, 0x87, 0x72, 0xa1, 0xba, 0x31, 0xa1, 0x6d, 0x64, 0x7a,
	0xd0, 0x45, 0x46, 0xba, 0x29, 0x70, 0x1f, 0x98, 0xa8, 0x1e, 0x2b, 0x58, 0x13, 0xa0, 0xdb, 0x60,
	0x5d, 0x26, 0xbc, 0x24, 0x58, 0x5c, 0x44, 0x2b, 0x2a, 0x1f, 0xd5, 0xfa, 0x46, 0x4b, 0xff, 0xb8,
	0x7f, 0xfd, 0xd3, 0x71, 0xd9, 0xfa, 0x7b, 0x5c, 0xb6, 0xdc, 0x3e, 0x72, 0x16, 0x35, 0x30, 0xea,
	0x9e, 0xa3, 0x82, 0x16, 0xd1, 0x0e, 0x01, 0xba, 0xed, 0x8e, 0x7e, 0x36, 0x52, 0x6f, 0xa6, 0xc2,
	0x49, 0x62, 0x79, 0x08, 0x81, 0x30, 0x42, 0xf3, 0x7c, 0xbe, 0x71, 0xfd, 0xdf, 0x32, 0x5a, 0x51,
	0x9c, 0xf8, 0xab, 0x8d, 0xd6, 0xe6, 0x87, 0x8c, 0x6b, 0x99, 0xee, 0x2f, 0xda, 0xbc, 0x52, 0xfd,
	0x2a, 0x25, 0xda, 0x96, 0x4b, 0x3e, 0xfe, 0xf8, 0xf3, 0x65, 0xa9, 0x8a, 0xb7, 0x69, 0xd6, 0x9f,
	0xe9, 0xdc, 0x7a, 0xe1, 0x0f, 0x36, 0xca, 0xe9, 0x11, 0xe0, 0x9d, 0xc5, 0x74, 0xa9, 0x79, 0x97,
	0xaa, 0x97, 0x03, 0x8d, 0x9a, 0xdb, 0x4a, 0xcd, 0x16, 0xde, 0xc8, 0x54, 0xa3, 0x87, 0x8d, 0xbf,
	0xd9, 0x28, 0x7f, 0x6e, 0x4e, 0xf8, 0x02, 0xf3, 0x8b, 0xb6, 0xa2, 0xb4, 0x77, 0xa5, 0x1a, 0xa3,
	0x71, 0x5f, 0x69, 0xac, 0xe3, 0x5d, 0xba, 0xf8, 0xf3, 0x93, 0xda, 0x11, 0xfa, 0x4e, 0x65, 0xf7,
	0xbe, 0xf1, 0xec, 0x64, 0xe4, 0xd8, 0xa7, 0x23, 0xc7, 0xfe, 0x3d, 0x72, 0xec, 0xcf, 0x63, 0xc7,
	0x3a, 0x1d, 0x3b, 0xd6, 0xcf, 0xb1, 0x63, 0xbd, 0xda, 0xf5, 0x83, 0xf8, 0x70, 0xd0, 0x21, 0x1e,
	0xf4, 0x66, 0xba, 0x9e, 0x9d, 0x86, 0xfb, 0xf4, 0x6d, 0x42, 0x12, 0x1f, 0x85, 0x5c, 0x76, 0x72,
	0xea, 0x63, 0xb4, 0xf7, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x63, 0x76, 0x5d, 0xb7, 0x89, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	RegisteredDenoms(ctx context.Context, in *QueryRegisteredDenomsRequest, opts ...grpc.CallOption) (*QueryRegisteredDenomsResponse, error)
	// Params retrieves the escrow module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EscrowPool queries the available supply of a coin.
	EscrowPoolBalance(ctx context.Context, in *QueryEscrowPoolBalanceRequest, opts ...grpc.CallOption) (*QueryEscrowPoolBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) RegisteredDenoms(ctx context.Context, in *QueryRegisteredDenomsRequest, opts ...grpc.CallOption) (*QueryRegisteredDenomsResponse, error) {
	out := new(QueryRegisteredDenomsResponse)
	err := c.cc.Invoke(ctx, "/reapchain.escrow.v1.Query/RegisteredDenoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/reapchain.escrow.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EscrowPoolBalance(ctx context.Context, in *QueryEscrowPoolBalanceRequest, opts ...grpc.CallOption) (*QueryEscrowPoolBalanceResponse, error) {
	out := new(QueryEscrowPoolBalanceResponse)
	err := c.cc.Invoke(ctx, "/reapchain.escrow.v1.Query/EscrowPoolBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	RegisteredDenoms(context.Context, *QueryRegisteredDenomsRequest) (*QueryRegisteredDenomsResponse, error)
	// Params retrieves the escrow module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EscrowPool queries the available supply of a coin.
	EscrowPoolBalance(context.Context, *QueryEscrowPoolBalanceRequest) (*QueryEscrowPoolBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) RegisteredDenoms(ctx context.Context, req *QueryRegisteredDenomsRequest) (*QueryRegisteredDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisteredDenoms not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) EscrowPoolBalance(ctx context.Context, req *QueryEscrowPoolBalanceRequest) (*QueryEscrowPoolBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscrowPoolBalance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_RegisteredDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRegisteredDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegisteredDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.escrow.v1.Query/RegisteredDenoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegisteredDenoms(ctx, req.(*QueryRegisteredDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.escrow.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EscrowPoolBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEscrowPoolBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EscrowPoolBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.escrow.v1.Query/EscrowPoolBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EscrowPoolBalance(ctx, req.(*QueryEscrowPoolBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reapchain.escrow.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisteredDenoms",
			Handler:    _Query_RegisteredDenoms_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EscrowPoolBalance",
			Handler:    _Query_EscrowPoolBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reapchain/escrow/v1/query.proto",
}

func (m *QueryRegisteredDenomsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRegisteredDenomsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRegisteredDenomsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRegisteredDenomsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRegisteredDenomsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRegisteredDenomsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RegisteredDenoms) > 0 {
		for iNdEx := len(m.RegisteredDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryEscrowPoolBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowPoolBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowPoolBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEscrowPoolBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowPoolBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowPoolBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EscrowPoolBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRegisteredDenomsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRegisteredDenomsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RegisteredDenoms) > 0 {
		for _, e := range m.RegisteredDenoms {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryEscrowPoolBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEscrowPoolBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EscrowPoolBalance.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRegisteredDenomsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRegisteredDenomsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRegisteredDenomsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRegisteredDenomsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRegisteredDenomsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRegisteredDenomsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredDenoms = append(m.RegisteredDenoms, RegisteredDenom{})
			if err := m.RegisteredDenoms[len(m.RegisteredDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEscrowPoolBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryEscrowPoolBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowPoolBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEscrowPoolBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryEscrowPoolBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowPoolBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowPoolBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EscrowPoolBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)