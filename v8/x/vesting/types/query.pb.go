// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reapchain/vesting/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryBalancesRequest is the request type for the Query/Balances RPC method.
type QueryBalancesRequest struct {
	// address of the clawback vesting account
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryBalancesRequest) Reset()         { *m = QueryBalancesRequest{} }
func (m *QueryBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalancesRequest) ProtoMessage()    {}
func (*QueryBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed258af479545424, []int{0}
}
func (m *QueryBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalancesRequest.Merge(m, src)
}
func (m *QueryBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalancesRequest proto.InternalMessageInfo

func (m *QueryBalancesRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryBalancesResponse is the response type for the Query/Balances RPC
// method.
type QueryBalancesResponse struct {
	// current amount of locked tokens
	Locked github_com_reapchain_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=locked,proto3,castrepeated=github.com/reapchain/cosmos-sdk/types.Coins" json:"locked"`
	// current amount of unvested tokens
	Unvested github_com_reapchain_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=unvested,proto3,castrepeated=github.com/reapchain/cosmos-sdk/types.Coins" json:"unvested"`
	// current amount of vested tokens
	Vested github_com_reapchain_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=vested,proto3,castrepeated=github.com/reapchain/cosmos-sdk/types.Coins" json:"vested"`
}

func (m *QueryBalancesResponse) Reset()         { *m = QueryBalancesResponse{} }
func (m *QueryBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalancesResponse) ProtoMessage()    {}
func (*QueryBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed258af479545424, []int{1}
}
func (m *QueryBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalancesResponse.Merge(m, src)
}
func (m *QueryBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalancesResponse proto.InternalMessageInfo

func (m *QueryBalancesResponse) GetLocked() github_com_reapchain_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *QueryBalancesResponse) GetUnvested() github_com_reapchain_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Unvested
	}
	return nil
}

func (m *QueryBalancesResponse) GetVested() github_com_reapchain_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Vested
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryBalancesRequest)(nil), "reapchain.vesting.v1.QueryBalancesRequest")
	proto.RegisterType((*QueryBalancesResponse)(nil), "reapchain.vesting.v1.QueryBalancesResponse")
}

func init() { proto.RegisterFile("reapchain/vesting/v1/query.proto", fileDescriptor_ed258af479545424) }

var fileDescriptor_ed258af479545424 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x3d, 0x6f, 0xe2, 0x40,
	0x10, 0xf5, 0x82, 0x8e, 0xe3, 0xf6, 0x3a, 0x8b, 0x93, 0x7c, 0xe8, 0x64, 0x10, 0x15, 0x02, 0xdd,
	0x2e, 0x86, 0x26, 0x35, 0x29, 0x53, 0x85, 0x32, 0xdd, 0xda, 0xde, 0x18, 0x07, 0xd8, 0x31, 0xde,
	0xb5, 0x15, 0x14, 0xa5, 0xc9, 0x2f, 0x88, 0x14, 0xa5, 0x4e, 0x9f, 0x5f, 0x42, 0x89, 0x44, 0x93,
	0x2a, 0x89, 0x20, 0x3f, 0x24, 0xf2, 0x07, 0x24, 0x8a, 0x5c, 0xa4, 0xa1, 0x1b, 0x7b, 0xdf, 0x7b,
	0x33, 0xef, 0xcd, 0xe0, 0x66, 0xc8, 0x59, 0xe0, 0x8c, 0x99, 0x2f, 0x68, 0xcc, 0xa5, 0xf2, 0x85,
	0x47, 0x63, 0x8b, 0xce, 0x23, 0x1e, 0x2e, 0x48, 0x10, 0x82, 0x02, 0xbd, 0xb6, 0x47, 0x90, 0x1c,
	0x41, 0x62, 0xab, 0x6e, 0x3a, 0x20, 0x67, 0x20, 0xa9, 0xcd, 0x24, 0xa7, 0xb1, 0x65, 0x73, 0xc5,
	0x2c, 0xea, 0x80, 0x2f, 0x32, 0x56, 0xfd, 0x9f, 0x07, 0xe0, 0x4d, 0x39, 0x65, 0x81, 0x4f, 0x99,
	0x10, 0xa0, 0x98, 0xf2, 0x41, 0xc8, 0xfc, 0xb5, 0xe6, 0x81, 0x07, 0x69, 0x49, 0x93, 0x2a, 0xfb,
	0xdb, 0xea, 0xe1, 0xda, 0x69, 0xd2, 0x78, 0xc8, 0xa6, 0x4c, 0x38, 0x5c, 0x8e, 0xf8, 0x3c, 0xe2,
	0x52, 0xe9, 0x06, 0xfe, 0xc9, 0x5c, 0x37, 0xe4, 0x52, 0x1a, 0xa8, 0x89, 0xda, 0xbf, 0x46, 0xbb,
	0xcf, 0xd6, 0xba, 0x84, 0xff, 0x7c, 0xa1, 0xc8, 0x00, 0x84, 0xe4, 0xfa, 0x39, 0xae, 0x4c, 0xc1,
	0x99, 0x70, 0xd7, 0x40, 0xcd, 0x72, 0xfb, 0x77, 0xff, 0x2f, 0xc9, 0x06, 0x26, 0xc9, 0xc0, 0x24,
	0x1f, 0x98, 0x1c, 0x83, 0x2f, 0x86, 0x83, 0xe5, 0x73, 0x43, 0x7b, 0x7c, 0x69, 0x74, 0x3d, 0x5f,
	0x8d, 0x23, 0x9b, 0x38, 0x30, 0xa3, 0x1f, 0xa9, 0x64, 0xb4, 0xff, 0xd2, 0x9d, 0x50, 0xb5, 0x08,
	0xb8, 0x4c, 0x39, 0x72, 0x94, 0xab, 0xeb, 0x17, 0xb8, 0x1a, 0x89, 0x24, 0x17, 0xee, 0x1a, 0xa5,
	0x83, 0x74, 0xda, 0xeb, 0x27, 0x9e, 0xf2, 0x4e, 0xe5, 0xc3, 0x78, 0xca, 0xd4, 0xfb, 0x0f, 0x08,
	0xff, 0x48, 0x53, 0xd5, 0xef, 0x11, 0xae, 0xee, 0xa2, 0xd5, 0x3b, 0xa4, 0xe8, 0x12, 0x48, 0xd1,
	0xca, 0xea, 0xdd, 0x6f, 0x61, 0xb3, 0x5d, 0xb5, 0x7a, 0x37, 0xeb, 0xb7, 0xbb, 0x52, 0x47, 0x6f,
	0xd3, 0xc2, 0x63, 0xb4, 0x73, 0x3c, 0xbd, 0xca, 0xd7, 0x7e, 0x3d, 0x3c, 0x59, 0x6e, 0x4c, 0xb4,
	0xda, 0x98, 0xe8, 0x75, 0x63, 0xa2, 0xdb, 0xad, 0xa9, 0xad, 0xb6, 0xa6, 0xf6, 0xb4, 0x35, 0xb5,
	0x33, 0xab, 0xd0, 0xf0, 0x27, 0xdd, 0x23, 0x7a, 0xb9, 0x17, 0x4f, 0xcd, 0xdb, 0x95, 0xf4, 0xfa,
	0x06, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xef, 0x06, 0xb4, 0x0b, 0x03, 0x00, 0x00,
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
	// Retrieves the unvested, vested and locked tokens for a vesting account
	Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error) {
	out := new(QueryBalancesResponse)
	err := c.cc.Invoke(ctx, "/reapchain.vesting.v1.Query/Balances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Retrieves the unvested, vested and locked tokens for a vesting account
	Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balances(ctx context.Context, req *QueryBalancesRequest) (*QueryBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balances not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.vesting.v1.Query/Balances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balances(ctx, req.(*QueryBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reapchain.vesting.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balances",
			Handler:    _Query_Balances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reapchain/vesting/v1/query.proto",
}

func (m *QueryBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vested) > 0 {
		for iNdEx := len(m.Vested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Unvested) > 0 {
		for iNdEx := len(m.Unvested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Unvested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Locked) > 0 {
		for iNdEx := len(m.Locked) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Locked[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Locked) > 0 {
		for _, e := range m.Locked {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Unvested) > 0 {
		for _, e := range m.Unvested {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Vested) > 0 {
		for _, e := range m.Vested {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalancesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryBalancesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
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
			m.Locked = append(m.Locked, types.Coin{})
			if err := m.Locked[len(m.Locked)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unvested", wireType)
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
			m.Unvested = append(m.Unvested, types.Coin{})
			if err := m.Unvested[len(m.Unvested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
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
			m.Vested = append(m.Vested, types.Coin{})
			if err := m.Vested[len(m.Vested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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