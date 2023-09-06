// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: witness/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgAnchoring struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	BlockHash   string `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockHeight string `protobuf:"bytes,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	ChainID     string `protobuf:"bytes,4,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Filter      string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (m *MsgAnchoring) Reset()         { *m = MsgAnchoring{} }
func (m *MsgAnchoring) String() string { return proto.CompactTextString(m) }
func (*MsgAnchoring) ProtoMessage()    {}
func (*MsgAnchoring) Descriptor() ([]byte, []int) {
	return fileDescriptor_39b323c2f7a86909, []int{0}
}
func (m *MsgAnchoring) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchoring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchoring.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchoring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchoring.Merge(m, src)
}
func (m *MsgAnchoring) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchoring) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchoring.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchoring proto.InternalMessageInfo

func (m *MsgAnchoring) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAnchoring) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *MsgAnchoring) GetBlockHeight() string {
	if m != nil {
		return m.BlockHeight
	}
	return ""
}

func (m *MsgAnchoring) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *MsgAnchoring) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type MsgAnchoringResponse struct {
}

func (m *MsgAnchoringResponse) Reset()         { *m = MsgAnchoringResponse{} }
func (m *MsgAnchoringResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAnchoringResponse) ProtoMessage()    {}
func (*MsgAnchoringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39b323c2f7a86909, []int{1}
}
func (m *MsgAnchoringResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchoringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchoringResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchoringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchoringResponse.Merge(m, src)
}
func (m *MsgAnchoringResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchoringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchoringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchoringResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAnchoring)(nil), "reapchain.witness.MsgAnchoring")
	proto.RegisterType((*MsgAnchoringResponse)(nil), "reapchain.witness.MsgAnchoringResponse")
}

func init() { proto.RegisterFile("witness/tx.proto", fileDescriptor_39b323c2f7a86909) }

var fileDescriptor_39b323c2f7a86909 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xcf, 0x2c, 0xc9,
	0x4b, 0x2d, 0x2e, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4a,
	0x4d, 0x2c, 0x48, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xca, 0x29, 0xcd, 0x62, 0xe4, 0xe2, 0xf1,
	0x2d, 0x4e, 0x77, 0xcc, 0x4b, 0xce, 0xc8, 0x2f, 0xca, 0xcc, 0x4b, 0x17, 0x92, 0xe0, 0x62, 0x4f,
	0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85,
	0x64, 0xb8, 0x38, 0x93, 0x72, 0xf2, 0x93, 0xb3, 0x3d, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc0, 0x72,
	0x08, 0x01, 0x21, 0x05, 0x2e, 0x6e, 0x08, 0x27, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x19, 0x2c,
	0x8f, 0x2c, 0x04, 0x36, 0x19, 0x64, 0xb7, 0xa7, 0x8b, 0x04, 0x0b, 0xd4, 0x64, 0x08, 0x57, 0x48,
	0x8c, 0x8b, 0x2d, 0x2d, 0x33, 0xa7, 0x24, 0xb5, 0x48, 0x82, 0x15, 0x2c, 0x01, 0xe5, 0x29, 0x89,
	0x71, 0x89, 0x20, 0xbb, 0x2d, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0x28, 0x86, 0x8b,
	0xd9, 0xb7, 0x38, 0x5d, 0x28, 0x94, 0x8b, 0x13, 0xe1, 0x6e, 0x79, 0x3d, 0x0c, 0xcf, 0xe9, 0x21,
	0x6b, 0x96, 0x52, 0x27, 0xa0, 0x00, 0x66, 0xba, 0x93, 0xf1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0x49, 0xc2, 0x4d, 0xd0, 0xaf, 0xd0, 0x87, 0x87, 0x6e, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x38, 0x84, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xb2, 0x7c, 0xb3,
	0x75, 0x01, 0x00, 0x00,
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
	Anchoring(ctx context.Context, in *MsgAnchoring, opts ...grpc.CallOption) (*MsgAnchoringResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Anchoring(ctx context.Context, in *MsgAnchoring, opts ...grpc.CallOption) (*MsgAnchoringResponse, error) {
	out := new(MsgAnchoringResponse)
	err := c.cc.Invoke(ctx, "/reapchain.witness.Msg/Anchoring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Anchoring(context.Context, *MsgAnchoring) (*MsgAnchoringResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Anchoring(ctx context.Context, req *MsgAnchoring) (*MsgAnchoringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Anchoring not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Anchoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAnchoring)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Anchoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reapchain.witness.Msg/Anchoring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Anchoring(ctx, req.(*MsgAnchoring))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reapchain.witness.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Anchoring",
			Handler:    _Msg_Anchoring_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "witness/tx.proto",
}

func (m *MsgAnchoring) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchoring) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchoring) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Filter) > 0 {
		i -= len(m.Filter)
		copy(dAtA[i:], m.Filter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Filter)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BlockHeight) > 0 {
		i -= len(m.BlockHeight)
		copy(dAtA[i:], m.BlockHeight)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BlockHeight)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnchoringResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchoringResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchoringResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAnchoring) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BlockHeight)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Filter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnchoringResponse) Size() (n int) {
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
func (m *MsgAnchoring) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAnchoring: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchoring: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
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
			m.BlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
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
			m.BlockHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
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
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
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
			m.Filter = string(dAtA[iNdEx:postIndex])
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
func (m *MsgAnchoringResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAnchoringResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchoringResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
