// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reapchain/claims/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState define the claims module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// list of claim records with the corresponding airdrop recipient
	ClaimsRecords []ClaimsRecordAddress `protobuf:"bytes,2,rep,name=claims_records,json=claimsRecords,proto3" json:"claims_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6775c608d2cfd6a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimsRecords() []ClaimsRecordAddress {
	if m != nil {
		return m.ClaimsRecords
	}
	return nil
}

// Params defines the claims module's parameters.
type Params struct {
	// enable claiming process
	EnableClaims bool `protobuf:"varint,1,opt,name=enable_claims,json=enableClaims,proto3" json:"enable_claims,omitempty"`
	// timestamp of the airdrop start
	AirdropStartTime time.Time `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time"`
	// duration until decay of claimable tokens begin
	DurationUntilDecay time.Duration `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay"`
	// duration of the token claim decay period
	DurationOfDecay time.Duration `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay"`
	// denom of claimable coin
	ClaimsDenom string `protobuf:"bytes,5,opt,name=claims_denom,json=claimsDenom,proto3" json:"claims_denom,omitempty"`
	// list of authorized channel identifiers that can perform address
	// attestations via IBC.
	AuthorizedChannels []string `protobuf:"bytes,6,rep,name=authorized_channels,json=authorizedChannels,proto3" json:"authorized_channels,omitempty"`
	// list of channel identifiers from EVM compatible chains
	EVMChannels []string `protobuf:"bytes,7,rep,name=evm_channels,json=evmChannels,proto3" json:"evm_channels,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6775c608d2cfd6a, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableClaims() bool {
	if m != nil {
		return m.EnableClaims
	}
	return false
}

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimsDenom() string {
	if m != nil {
		return m.ClaimsDenom
	}
	return ""
}

func (m *Params) GetAuthorizedChannels() []string {
	if m != nil {
		return m.AuthorizedChannels
	}
	return nil
}

func (m *Params) GetEVMChannels() []string {
	if m != nil {
		return m.EVMChannels
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "evmos.claims.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "evmos.claims.v1.Params")
}

func init() { proto.RegisterFile("reapchain/claims/v1/genesis.proto", fileDescriptor_c6775c608d2cfd6a) }

var fileDescriptor_c6775c608d2cfd6a = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0xe3, 0xa6, 0x7f, 0xfe, 0x76, 0x9c, 0x12, 0x18, 0x2a, 0x61, 0xb2, 0x70, 0xdc, 0xc2,
	0x22, 0x2b, 0x9b, 0x06, 0x78, 0x00, 0xd2, 0x20, 0x24, 0x24, 0x54, 0x70, 0x29, 0x0b, 0x36, 0xd6,
	0xc4, 0xbe, 0x71, 0x2c, 0xd9, 0x1e, 0x6b, 0x66, 0x6c, 0x51, 0x9e, 0xa2, 0xcb, 0xbe, 0x06, 0x6f,
	0xd1, 0x65, 0x97, 0xac, 0x0a, 0x4a, 0x5e, 0x04, 0xcd, 0x8c, 0x87, 0xa0, 0x96, 0x05, 0xbb, 0xf1,
	0x9c, 0xef, 0x9c, 0x7b, 0x7d, 0x34, 0xe8, 0x80, 0x01, 0xa9, 0xe2, 0x25, 0xc9, 0xca, 0x20, 0xce,
	0x49, 0x56, 0xf0, 0xa0, 0x39, 0x0a, 0x52, 0x28, 0x81, 0x67, 0xdc, 0xaf, 0x18, 0x15, 0x14, 0x0f,
	0xa0, 0x29, 0x28, 0xf7, 0xb5, 0xec, 0x37, 0x47, 0xc3, 0xfd, 0x94, 0xa6, 0x54, 0x69, 0x81, 0x3c,
	0x69, 0x6c, 0xe8, 0xa6, 0x94, 0xa6, 0x39, 0x04, 0xea, 0x6b, 0x5e, 0x2f, 0x82, 0xa4, 0x66, 0x44,
	0x64, 0xb4, 0x6c, 0xf5, 0xd1, 0x6d, 0x5d, 0x64, 0x05, 0x70, 0x41, 0x8a, 0xaa, 0x05, 0xbc, 0xbf,
	0xad, 0xd2, 0x4e, 0x55, 0xc4, 0xe1, 0xa5, 0x85, 0xfa, 0x6f, 0xf4, 0x6e, 0xa7, 0x82, 0x08, 0xc0,
	0x2f, 0x51, 0xaf, 0x22, 0x8c, 0x14, 0xdc, 0xb1, 0x3c, 0x6b, 0x6c, 0x4f, 0x1e, 0xf9, 0xb7, 0x76,
	0xf5, 0xdf, 0x2b, 0x79, 0xba, 0x7d, 0x75, 0x33, 0xea, 0x84, 0x2d, 0x8c, 0x3f, 0xa0, 0x7b, 0x9a,
	0x88, 0x18, 0xc4, 0x94, 0x25, 0xdc, 0xd9, 0xf2, 0xba, 0x63, 0x7b, 0xf2, 0xf4, 0x8e, 0xfd, 0x58,
	0x9d, 0x42, 0x45, 0xbd, 0x4a, 0x12, 0x06, 0xdc, 0x64, 0xed, 0xc5, 0x7f, 0x48, 0xfc, 0xf0, 0x5b,
	0x17, 0xf5, 0xf4, 0x2c, 0xfc, 0x04, 0xed, 0x41, 0x49, 0xe6, 0x39, 0x44, 0x1a, 0x51, 0xbb, 0xed,
	0x84, 0x7d, 0x7d, 0xa9, 0x13, 0x71, 0x88, 0x30, 0xc9, 0x58, 0xc2, 0x68, 0x15, 0x71, 0x41, 0x98,
	0x88, 0x64, 0x1b, 0xce, 0x96, 0xfa, 0x8b, 0xa1, 0xaf, 0xab, 0xf2, 0x4d, 0x55, 0xfe, 0x47, 0x53,
	0xd5, 0x74, 0x47, 0x0e, 0xbf, 0xf8, 0x31, 0xb2, 0xc2, 0xfb, 0xad, 0xff, 0x54, 0xda, 0x25, 0x80,
	0xcf, 0xd0, 0xbe, 0xe9, 0x3c, 0xaa, 0x4b, 0x91, 0xe5, 0x51, 0x02, 0x31, 0x39, 0x77, 0xba, 0x2a,
	0xf5, 0xf1, 0x9d, 0xd4, 0x59, 0x0b, 0xeb, 0xd0, 0x4b, 0x19, 0x8a, 0x4d, 0xc0, 0x99, 0xf4, 0xcf,
	0xa4, 0x1d, 0x9f, 0xa0, 0x07, 0xbf, 0x63, 0xe9, 0xa2, 0xcd, 0xdc, 0xfe, 0xf7, 0xcc, 0x81, 0x71,
	0x9f, 0x2c, 0x74, 0xe0, 0x01, 0xea, 0xb7, 0xf5, 0x27, 0x50, 0xd2, 0xc2, 0xf9, 0xcf, 0xb3, 0xc6,
	0xbb, 0xa1, 0xad, 0xef, 0x66, 0xf2, 0x0a, 0x07, 0xe8, 0x21, 0xa9, 0xc5, 0x92, 0xb2, 0xec, 0x2b,
	0x24, 0x51, 0xbc, 0x24, 0x65, 0x09, 0x39, 0x77, 0x7a, 0x5e, 0x77, 0xbc, 0x1b, 0xe2, 0x8d, 0x74,
	0xdc, 0x2a, 0x78, 0x82, 0xfa, 0xd0, 0x14, 0x1b, 0xf2, 0x7f, 0x49, 0x4e, 0x07, 0xab, 0x9b, 0x91,
	0xfd, 0xfa, 0xd3, 0x3b, 0x83, 0x85, 0x36, 0x34, 0x85, 0xf9, 0x98, 0xbe, 0xbd, 0x5a, 0xb9, 0xd6,
	0xf5, 0xca, 0xb5, 0x7e, 0xae, 0x5c, 0xeb, 0x62, 0xed, 0x76, 0xae, 0xd7, 0x6e, 0xe7, 0xfb, 0xda,
	0xed, 0x7c, 0x7e, 0x96, 0x66, 0x62, 0x59, 0xcf, 0xfd, 0x98, 0x16, 0xc1, 0xe6, 0x55, 0x6e, 0x4e,
	0xcd, 0x8b, 0xe0, 0x8b, 0x79, 0xa4, 0xe2, 0xbc, 0x02, 0x3e, 0xef, 0xa9, 0x06, 0x9e, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x43, 0xae, 0xe3, 0x66, 0x50, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimsRecords) > 0 {
		for iNdEx := len(m.ClaimsRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimsRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EVMChannels) > 0 {
		for iNdEx := len(m.EVMChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EVMChannels[iNdEx])
			copy(dAtA[i:], m.EVMChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.EVMChannels[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AuthorizedChannels) > 0 {
		for iNdEx := len(m.AuthorizedChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedChannels[iNdEx])
			copy(dAtA[i:], m.AuthorizedChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AuthorizedChannels[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClaimsDenom) > 0 {
		i -= len(m.ClaimsDenom)
		copy(dAtA[i:], m.ClaimsDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClaimsDenom)))
		i--
		dAtA[i] = 0x2a
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if m.EnableClaims {
		i--
		if m.EnableClaims {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimsRecords) > 0 {
		for _, e := range m.ClaimsRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableClaims {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.ClaimsDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AuthorizedChannels) > 0 {
		for _, s := range m.AuthorizedChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EVMChannels) > 0 {
		for _, s := range m.EVMChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimsRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimsRecords = append(m.ClaimsRecords, ClaimsRecordAddress{})
			if err := m.ClaimsRecords[len(m.ClaimsRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableClaims", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableClaims = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimsDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimsDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedChannels = append(m.AuthorizedChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EVMChannels = append(m.EVMChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
