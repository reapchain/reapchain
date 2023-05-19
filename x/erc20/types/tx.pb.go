// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reapchain/erc20/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_reapchain_cosmos_sdk_types "github.com/reapchain/cosmos-sdk/types"
	types "github.com/reapchain/cosmos-sdk/types"
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

// MsgConvertCoin defines a Msg to convert a native Cosmos coin to a ERC20 token
type MsgConvertCoin struct {
	// Cosmos coin which denomination is registered in a token pair. The coin
	// amount defines the amount of coins to convert.
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin"`
	// recipient hex address to receive ERC20 token
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertCoin) Reset()         { *m = MsgConvertCoin{} }
func (m *MsgConvertCoin) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoin) ProtoMessage()    {}
func (*MsgConvertCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_166794c411253d11, []int{0}
}
func (m *MsgConvertCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoin.Merge(m, src)
}
func (m *MsgConvertCoin) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoin proto.InternalMessageInfo

func (m *MsgConvertCoin) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *MsgConvertCoin) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertCoin) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertCoinResponse returns no fields
type MsgConvertCoinResponse struct {
}

func (m *MsgConvertCoinResponse) Reset()         { *m = MsgConvertCoinResponse{} }
func (m *MsgConvertCoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoinResponse) ProtoMessage()    {}
func (*MsgConvertCoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_166794c411253d11, []int{1}
}
func (m *MsgConvertCoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoinResponse.Merge(m, src)
}
func (m *MsgConvertCoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoinResponse proto.InternalMessageInfo

// MsgConvertERC20 defines a Msg to convert a ERC20 token to a native Cosmos
// coin.
type MsgConvertERC20 struct {
	// ERC20 token contract address registered in a token pair
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// amount of ERC20 tokens to convert
	Amount github_com_reapchain_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/reapchain/cosmos-sdk/types.Int" json:"amount"`
	// bech32 address to receive native Cosmos coins
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender hex address from the owner of the given ERC20 tokens
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertERC20) Reset()         { *m = MsgConvertERC20{} }
func (m *MsgConvertERC20) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20) ProtoMessage()    {}
func (*MsgConvertERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_166794c411253d11, []int{2}
}
func (m *MsgConvertERC20) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20.Merge(m, src)
}
func (m *MsgConvertERC20) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20 proto.InternalMessageInfo

func (m *MsgConvertERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertERC20) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertERC20) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertERC20Response returns no fields
type MsgConvertERC20Response struct {
}

func (m *MsgConvertERC20Response) Reset()         { *m = MsgConvertERC20Response{} }
func (m *MsgConvertERC20Response) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20Response) ProtoMessage()    {}
func (*MsgConvertERC20Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_166794c411253d11, []int{3}
}
func (m *MsgConvertERC20Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20Response.Merge(m, src)
}
func (m *MsgConvertERC20Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertCoin)(nil), "reapchain.erc20.v1.MsgConvertCoin")
	proto.RegisterType((*MsgConvertCoinResponse)(nil), "reapchain.erc20.v1.MsgConvertCoinResponse")
	proto.RegisterType((*MsgConvertERC20)(nil), "reapchain.erc20.v1.MsgConvertERC20")
	proto.RegisterType((*MsgConvertERC20Response)(nil), "reapchain.erc20.v1.MsgConvertERC20Response")
}

func init() { proto.RegisterFile("reapchain/erc20/v1/tx.proto", fileDescriptor_166794c411253d11) }

var fileDescriptor_166794c411253d11 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xdb, 0xaa, 0x62, 0x1e, 0x62, 0xc8, 0x42, 0xa3, 0x0b, 0x28, 0x9b, 0x3a, 0x40, 0x1b,
	0x03, 0x7b, 0xe9, 0x2e, 0x5c, 0x69, 0xc5, 0xa1, 0x87, 0x5d, 0x72, 0xe4, 0x32, 0x39, 0x8e, 0x95,
	0x45, 0x50, 0xbf, 0xc8, 0xf6, 0xa2, 0xed, 0xca, 0x91, 0x03, 0x42, 0xe2, 0x27, 0xf0, 0x33, 0xf8,
	0x03, 0x3b, 0x4e, 0xe2, 0x82, 0x38, 0x4c, 0xa8, 0xe5, 0x87, 0xa0, 0x38, 0x59, 0x4a, 0x60, 0xd0,
	0xdd, 0x9e, 0xdf, 0xf7, 0xfc, 0xf9, 0xfb, 0xbe, 0x67, 0xfc, 0x40, 0x4b, 0x9e, 0x89, 0x63, 0x9e,
	0x2a, 0x26, 0xb5, 0x18, 0xee, 0xb3, 0x3c, 0x60, 0xf6, 0x94, 0x66, 0x1a, 0x2c, 0x10, 0x52, 0x83,
	0xd4, 0x81, 0x34, 0x0f, 0xbc, 0x87, 0x09, 0x40, 0xf2, 0x56, 0x32, 0x9e, 0xa5, 0x8c, 0x2b, 0x05,
	0x96, 0xdb, 0x14, 0x94, 0x29, 0x6f, 0x78, 0xf7, 0x12, 0x48, 0xc0, 0x95, 0xac, 0xa8, 0xaa, 0xae,
	0x2f, 0xc0, 0x4c, 0xc1, 0xb0, 0x88, 0x1b, 0xc9, 0xf2, 0x20, 0x92, 0x96, 0x07, 0x4c, 0x40, 0xaa,
	0x4a, 0x7c, 0x70, 0x86, 0xef, 0x1c, 0x9a, 0x64, 0x0c, 0x2a, 0x97, 0xda, 0x8e, 0x21, 0x55, 0xe4,
	0x00, 0x77, 0x0b, 0xbc, 0x8f, 0xb6, 0xd0, 0xce, 0xea, 0x70, 0x83, 0x96, 0x04, 0xb4, 0x20, 0xa0,
	0x15, 0x01, 0x2d, 0x06, 0x47, 0xdd, 0xf3, 0xcb, 0xcd, 0x56, 0xe8, 0x86, 0x89, 0x87, 0x6f, 0x69,
	0x29, 0x64, 0x9a, 0x4b, 0xdd, 0x6f, 0x6f, 0xa1, 0x9d, 0x95, 0xb0, 0x3e, 0x93, 0x75, 0xdc, 0x33,
	0x52, 0xc5, 0x52, 0xf7, 0x3b, 0x0e, 0xa9, 0x4e, 0x83, 0x3e, 0x5e, 0x6f, 0x3e, 0x1d, 0x4a, 0x93,
	0x81, 0x32, 0x72, 0xf0, 0x05, 0xe1, 0xb5, 0x05, 0xf4, 0x2a, 0x1c, 0x0f, 0xf7, 0xc9, 0x2e, 0xbe,
	0x2b, 0x40, 0x59, 0xcd, 0x85, 0x3d, 0xe2, 0x71, 0xac, 0xa5, 0x31, 0x4e, 0xe2, 0x4a, 0xb8, 0x76,
	0xd5, 0x7f, 0x59, 0xb6, 0xc9, 0x04, 0xf7, 0xf8, 0x14, 0x4e, 0x94, 0x2d, 0xa5, 0x8c, 0x82, 0x42,
	0xe8, 0xf7, 0xcb, 0xcd, 0xdd, 0x24, 0xb5, 0xc7, 0x27, 0x11, 0x15, 0x30, 0x65, 0x8b, 0xec, 0x4b,
	0x7f, 0xcf, 0x4d, 0xfc, 0x86, 0xd9, 0xb3, 0x4c, 0x1a, 0x3a, 0x51, 0x36, 0xac, 0x08, 0x1a, 0xbe,
	0x3a, 0xff, 0xf4, 0xd5, 0x6d, 0xf8, 0xda, 0xc0, 0xf7, 0xff, 0x10, 0x7f, 0x65, 0x6c, 0xf8, 0xb9,
	0x8d, 0x3b, 0x87, 0x26, 0x21, 0xef, 0x11, 0x5e, 0xfd, 0x3d, 0xf3, 0x01, 0xfd, 0x7b, 0xdd, 0xb4,
	0x19, 0x8e, 0xf7, 0x74, 0xf9, 0x4c, 0x1d, 0xe0, 0xde, 0xbb, 0xaf, 0x3f, 0x3f, 0xb5, 0x1f, 0x93,
	0x6d, 0x76, 0xed, 0x1f, 0x63, 0xa2, 0xbc, 0x73, 0xe4, 0x76, 0xf7, 0x01, 0xe1, 0xdb, 0x8d, 0xa8,
	0xb7, 0xff, 0xff, 0x92, 0x1b, 0xf2, 0xf6, 0x6e, 0x30, 0x54, 0xeb, 0x79, 0xe6, 0xf4, 0x3c, 0x21,
	0x8f, 0x96, 0xe8, 0x71, 0xbd, 0xd1, 0xe4, 0x7c, 0xe6, 0xa3, 0x8b, 0x99, 0x8f, 0x7e, 0xcc, 0x7c,
	0xf4, 0x71, 0xee, 0xb7, 0x2e, 0xe6, 0x7e, 0xeb, 0xdb, 0xdc, 0x6f, 0xbd, 0x66, 0xd7, 0x6e, 0x70,
	0x51, 0xe5, 0x2f, 0xd8, 0x69, 0x45, 0xec, 0x76, 0x19, 0xf5, 0xdc, 0x2f, 0x3f, 0xf8, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0x7a, 0xde, 0x11, 0x6c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error) {
	out := new(MsgConvertCoinResponse)
	err := c.cc.Invoke(ctx, "/reapchain.erc20.v1.Msg/ConvertCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error) {
	out := new(MsgConvertERC20Response)
	err := c.cc.Invoke(ctx, "/reapchain.erc20.v1.Msg/ConvertERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(context.Context, *MsgConvertCoin) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(context.Context, *MsgConvertERC20) (*MsgConvertERC20Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertCoin(ctx context.Context, req *MsgConvertCoin) (*MsgConvertCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCoin not implemented")
}
func (*UnimplementedMsgServer) ConvertERC20(ctx context.Context, req *MsgConvertERC20) (*MsgConvertERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC20 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.erc20.v1.Msg/ConvertCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertCoin(ctx, req.(*MsgConvertCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.erc20.v1.Msg/ConvertERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertERC20(ctx, req.(*MsgConvertERC20))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reapchain.erc20.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertCoin",
			Handler:    _Msg_ConvertCoin_Handler,
		},
		{
			MethodName: "ConvertERC20",
			Handler:    _Msg_ConvertERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reapchain/erc20/v1/tx.proto",
}

func (m *MsgConvertCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgConvertCoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertCoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConvertERC20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertERC20Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertCoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertCoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertERC20) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertERC20: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertERC20Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertERC20Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
