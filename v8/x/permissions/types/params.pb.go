// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reapchain/permissions/v1/params.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_reapchain_cosmos_sdk_types "github.com/reapchain/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	PodcWhitelistEnabled           bool                                      `protobuf:"varint,1,opt,name=podc_whitelist_enabled,json=podcWhitelistEnabled,proto3" json:"podc_whitelist_enabled,omitempty" yaml:"podc_whitelist_enabled"`
	GovMinInitialDepositEnabled    bool                                      `protobuf:"varint,2,opt,name=gov_min_initial_deposit_enabled,json=govMinInitialDepositEnabled,proto3" json:"gov_min_initial_deposit_enabled,omitempty" yaml:"gov_min_initial_deposit_enabled"`
	GovMinInitialDepositPercentage github_com_reapchain_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=gov_min_initial_deposit_percentage,json=govMinInitialDepositPercentage,proto3,customtype=github.com/reapchain/cosmos-sdk/types.Dec" json:"gov_min_initial_deposit_percentage"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c88ce75a0ed674b, []int{0}
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

func (m *Params) GetPodcWhitelistEnabled() bool {
	if m != nil {
		return m.PodcWhitelistEnabled
	}
	return false
}

func (m *Params) GetGovMinInitialDepositEnabled() bool {
	if m != nil {
		return m.GovMinInitialDepositEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "reapchain.permissions.v1.Params")
}

func init() {
	proto.RegisterFile("reapchain/permissions/v1/params.proto", fileDescriptor_1c88ce75a0ed674b)
}

var fileDescriptor_1c88ce75a0ed674b = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x73, 0x56, 0x8a, 0x66, 0x2c, 0x45, 0x8a, 0xe2, 0xa5, 0x04, 0x94, 0x0a, 0x9a, 0xa3,
	0x88, 0x20, 0x1d, 0x4b, 0x3b, 0x38, 0x88, 0xa5, 0x4b, 0xc1, 0x25, 0x5c, 0x93, 0x23, 0x3d, 0xcc,
	0xfd, 0x20, 0x77, 0x46, 0x8b, 0xf8, 0x3f, 0x38, 0xba, 0x08, 0xce, 0xfe, 0x25, 0x1d, 0x3b, 0x8a,
	0x43, 0x90, 0x76, 0xeb, 0xe8, 0x5f, 0x20, 0x4d, 0x68, 0x6b, 0x21, 0xd2, 0xed, 0xf1, 0xde, 0xe7,
	0xde, 0xf7, 0x03, 0xf7, 0xcc, 0xa3, 0x88, 0x60, 0xe9, 0x0d, 0x30, 0xe5, 0x48, 0x92, 0x88, 0x51,
	0xa5, 0xa8, 0xe0, 0x0a, 0xc5, 0x75, 0x24, 0x71, 0x84, 0x99, 0x72, 0x64, 0x24, 0xb4, 0x28, 0x55,
	0x96, 0x98, 0xf3, 0x07, 0x73, 0xe2, 0xfa, 0x7e, 0x39, 0x10, 0x81, 0x48, 0x21, 0x34, 0xaf, 0x32,
	0xde, 0xfe, 0x28, 0x98, 0xc5, 0x4e, 0xba, 0xa0, 0xf4, 0x64, 0xee, 0x49, 0xe1, 0x7b, 0xee, 0xc3,
	0x80, 0x6a, 0x12, 0x52, 0xa5, 0x5d, 0xc2, 0x71, 0x3f, 0x24, 0x7e, 0x05, 0x54, 0x41, 0x6d, 0xa7,
	0xd9, 0x9e, 0x25, 0x56, 0x35, 0x9f, 0x38, 0x15, 0x8c, 0x6a, 0xc2, 0xa4, 0x1e, 0xfe, 0x24, 0xd6,
	0xe1, 0x10, 0xb3, 0xb0, 0x61, 0xe7, 0x93, 0x76, 0xb7, 0x3c, 0x1f, 0xf4, 0x16, 0xfd, 0x76, 0xd6,
	0x2e, 0xbd, 0x01, 0xd3, 0x0a, 0x44, 0xec, 0x32, 0xca, 0x5d, 0xca, 0xa9, 0xa6, 0x38, 0x74, 0x7d,
	0x22, 0x85, 0xa2, 0x2b, 0x8d, 0xad, 0x54, 0xa3, 0x37, 0x4b, 0xac, 0x93, 0x0d, 0xe8, 0x9a, 0xcf,
	0x71, 0xe6, 0xb3, 0xe1, 0x89, 0xdd, 0x3d, 0x08, 0x44, 0x7c, 0x4d, 0xf9, 0x55, 0x36, 0x6f, 0x65,
	0xe3, 0x85, 0xdf, 0xb3, 0xf9, 0xef, 0x02, 0x49, 0x22, 0x8f, 0x70, 0x8d, 0x03, 0x52, 0x29, 0x54,
	0x41, 0x6d, 0xb7, 0x59, 0x1f, 0x25, 0x96, 0xf1, 0x35, 0xb7, 0xa4, 0x7a, 0x70, 0xdf, 0x77, 0x3c,
	0xc1, 0xd0, 0xea, 0xf7, 0x3c, 0xa1, 0x98, 0x50, 0x67, 0xca, 0xbf, 0x43, 0x7a, 0x28, 0x89, 0x72,
	0x5a, 0xc4, 0xeb, 0xc2, 0xbc, 0xec, 0xce, 0x72, 0x71, 0x63, 0xfb, 0xf5, 0xdd, 0x32, 0x9a, 0x37,
	0xa3, 0x09, 0x04, 0xe3, 0x09, 0x04, 0xdf, 0x13, 0x08, 0x5e, 0xa6, 0xd0, 0x18, 0x4f, 0xa1, 0xf1,
	0x39, 0x85, 0xc6, 0xed, 0x45, 0x6e, 0xd4, 0xaa, 0x8a, 0x2f, 0xd1, 0xe3, 0xda, 0xdd, 0xa4, 0xd1,
	0xfd, 0x62, 0x7a, 0x04, 0xe7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x1d, 0x4a, 0xff, 0x5d,
	0x02, 0x00, 0x00,
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
	{
		size := m.GovMinInitialDepositPercentage.Size()
		i -= size
		if _, err := m.GovMinInitialDepositPercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.GovMinInitialDepositEnabled {
		i--
		if m.GovMinInitialDepositEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.PodcWhitelistEnabled {
		i--
		if m.PodcWhitelistEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PodcWhitelistEnabled {
		n += 2
	}
	if m.GovMinInitialDepositEnabled {
		n += 2
	}
	l = m.GovMinInitialDepositPercentage.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field PodcWhitelistEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.PodcWhitelistEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovMinInitialDepositEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.GovMinInitialDepositEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovMinInitialDepositPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GovMinInitialDepositPercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)