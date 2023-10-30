// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gravity/v1/pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/reapchain/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// IDSet represents a set of IDs
type IDSet struct {
	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *IDSet) Reset()         { *m = IDSet{} }
func (m *IDSet) String() string { return proto.CompactTextString(m) }
func (*IDSet) ProtoMessage()    {}
func (*IDSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d107f7cfc31f22, []int{0}
}
func (m *IDSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDSet.Merge(m, src)
}
func (m *IDSet) XXX_Size() int {
	return m.Size()
}
func (m *IDSet) XXX_DiscardUnknown() {
	xxx_messageInfo_IDSet.DiscardUnknown(m)
}

var xxx_messageInfo_IDSet proto.InternalMessageInfo

func (m *IDSet) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type BatchFees struct {
	Token     string                                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TotalFees github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_fees,json=totalFees,proto3,customtype=github.com/reapchain/cosmos-sdk/types.Int" json:"total_fees"`
	TxCount   uint64                                 `protobuf:"varint,3,opt,name=tx_count,json=txCount,proto3" json:"tx_count,omitempty"`
}

func (m *BatchFees) Reset()         { *m = BatchFees{} }
func (m *BatchFees) String() string { return proto.CompactTextString(m) }
func (*BatchFees) ProtoMessage()    {}
func (*BatchFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d107f7cfc31f22, []int{1}
}
func (m *BatchFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchFees.Merge(m, src)
}
func (m *BatchFees) XXX_Size() int {
	return m.Size()
}
func (m *BatchFees) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchFees.DiscardUnknown(m)
}

var xxx_messageInfo_BatchFees proto.InternalMessageInfo

func (m *BatchFees) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BatchFees) GetTxCount() uint64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

type EventWithdrawalReceived struct {
	BridgeContract string `protobuf:"bytes,1,opt,name=bridge_contract,json=bridgeContract,proto3" json:"bridge_contract,omitempty"`
	BridgeChainId  string `protobuf:"bytes,2,opt,name=bridge_chain_id,json=bridgeChainId,proto3" json:"bridge_chain_id,omitempty"`
	OutgoingTxId   string `protobuf:"bytes,3,opt,name=outgoing_tx_id,json=outgoingTxId,proto3" json:"outgoing_tx_id,omitempty"`
	Nonce          string `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *EventWithdrawalReceived) Reset()         { *m = EventWithdrawalReceived{} }
func (m *EventWithdrawalReceived) String() string { return proto.CompactTextString(m) }
func (*EventWithdrawalReceived) ProtoMessage()    {}
func (*EventWithdrawalReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d107f7cfc31f22, []int{2}
}
func (m *EventWithdrawalReceived) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdrawalReceived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdrawalReceived.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdrawalReceived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdrawalReceived.Merge(m, src)
}
func (m *EventWithdrawalReceived) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdrawalReceived) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdrawalReceived.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdrawalReceived proto.InternalMessageInfo

func (m *EventWithdrawalReceived) GetBridgeContract() string {
	if m != nil {
		return m.BridgeContract
	}
	return ""
}

func (m *EventWithdrawalReceived) GetBridgeChainId() string {
	if m != nil {
		return m.BridgeChainId
	}
	return ""
}

func (m *EventWithdrawalReceived) GetOutgoingTxId() string {
	if m != nil {
		return m.OutgoingTxId
	}
	return ""
}

func (m *EventWithdrawalReceived) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type EventWithdrawCanceled struct {
	Sender         string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TxId           string `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	BridgeContract string `protobuf:"bytes,3,opt,name=bridge_contract,json=bridgeContract,proto3" json:"bridge_contract,omitempty"`
	BridgeChainId  string `protobuf:"bytes,4,opt,name=bridge_chain_id,json=bridgeChainId,proto3" json:"bridge_chain_id,omitempty"`
}

func (m *EventWithdrawCanceled) Reset()         { *m = EventWithdrawCanceled{} }
func (m *EventWithdrawCanceled) String() string { return proto.CompactTextString(m) }
func (*EventWithdrawCanceled) ProtoMessage()    {}
func (*EventWithdrawCanceled) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d107f7cfc31f22, []int{3}
}
func (m *EventWithdrawCanceled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdrawCanceled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdrawCanceled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdrawCanceled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdrawCanceled.Merge(m, src)
}
func (m *EventWithdrawCanceled) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdrawCanceled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdrawCanceled.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdrawCanceled proto.InternalMessageInfo

func (m *EventWithdrawCanceled) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventWithdrawCanceled) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *EventWithdrawCanceled) GetBridgeContract() string {
	if m != nil {
		return m.BridgeContract
	}
	return ""
}

func (m *EventWithdrawCanceled) GetBridgeChainId() string {
	if m != nil {
		return m.BridgeChainId
	}
	return ""
}

func init() {
	proto.RegisterType((*IDSet)(nil), "gravity.v1.IDSet")
	proto.RegisterType((*BatchFees)(nil), "gravity.v1.BatchFees")
	proto.RegisterType((*EventWithdrawalReceived)(nil), "gravity.v1.EventWithdrawalReceived")
	proto.RegisterType((*EventWithdrawCanceled)(nil), "gravity.v1.EventWithdrawCanceled")
}

func init() { proto.RegisterFile("gravity/v1/pool.proto", fileDescriptor_18d107f7cfc31f22) }

var fileDescriptor_18d107f7cfc31f22 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x62, 0xa7, 0x90, 0x15, 0x14, 0xb4, 0xb4, 0xe0, 0x72, 0x70, 0xa3, 0x08, 0x95, 0x5c,
	0x6a, 0xab, 0xe2, 0x01, 0x90, 0x12, 0x7e, 0x94, 0x03, 0x17, 0x83, 0x84, 0xe0, 0x62, 0x39, 0xbb,
	0x83, 0xb3, 0xaa, 0xb3, 0x13, 0xd9, 0x13, 0xe3, 0x3e, 0x03, 0x17, 0x2e, 0xbc, 0x02, 0xcf, 0xd2,
	0x63, 0x8f, 0x88, 0x43, 0x85, 0x92, 0x17, 0x41, 0xde, 0x75, 0xc5, 0x8f, 0x7a, 0xe0, 0xe4, 0xf9,
	0x3e, 0x7f, 0x3b, 0xf3, 0xcd, 0x0f, 0xdf, 0xcf, 0xcb, 0xac, 0xd6, 0x74, 0x16, 0xd7, 0x27, 0xf1,
	0x0a, 0xb1, 0x88, 0x56, 0x25, 0x12, 0x0a, 0xde, 0xd1, 0x51, 0x7d, 0xf2, 0x68, 0x2f, 0xc7, 0x1c,
	0x2d, 0x1d, 0xb7, 0x91, 0x53, 0x8c, 0x0e, 0x78, 0x7f, 0xf6, 0xfc, 0x0d, 0x90, 0xb8, 0xc7, 0x3d,
	0xad, 0xaa, 0x80, 0x0d, 0xbd, 0xb1, 0x9f, 0xb4, 0xe1, 0xe8, 0x33, 0xe3, 0x83, 0x49, 0x46, 0x72,
	0xf1, 0x12, 0xa0, 0x12, 0x7b, 0xbc, 0x4f, 0x78, 0x0a, 0x26, 0x60, 0x43, 0x36, 0x1e, 0x24, 0x0e,
	0x88, 0xd7, 0x9c, 0x13, 0x52, 0x56, 0xa4, 0x1f, 0x01, 0xaa, 0xe0, 0x46, 0xfb, 0x6b, 0x12, 0x9d,
	0x5f, 0x1e, 0xf6, 0x7e, 0x5c, 0x1e, 0x1e, 0xe5, 0x9a, 0x16, 0xeb, 0x79, 0x24, 0x71, 0x19, 0x4b,
	0xac, 0x96, 0x58, 0x75, 0x9f, 0xe3, 0x4a, 0x9d, 0xc6, 0x74, 0xb6, 0x82, 0x2a, 0x9a, 0x19, 0x4a,
	0x06, 0x36, 0x83, 0x2d, 0x72, 0xc0, 0x6f, 0x51, 0x93, 0x4a, 0x5c, 0x1b, 0x0a, 0xbc, 0x21, 0x1b,
	0xfb, 0xc9, 0x4d, 0x6a, 0xa6, 0x2d, 0x1c, 0x7d, 0x63, 0xfc, 0xe1, 0x8b, 0x1a, 0x0c, 0xbd, 0xd3,
	0xb4, 0x50, 0x65, 0xf6, 0x29, 0x2b, 0x12, 0x90, 0xa0, 0x6b, 0x50, 0xe2, 0x09, 0xbf, 0x3b, 0x2f,
	0xb5, 0xca, 0x21, 0x95, 0x68, 0xa8, 0xcc, 0x24, 0x75, 0x2e, 0x77, 0x1d, 0x3d, 0xed, 0x58, 0x71,
	0xf4, 0x5b, 0xb8, 0xc8, 0xb4, 0x49, 0xb5, 0x72, 0x9e, 0x93, 0x3b, 0x9d, 0xb0, 0x65, 0x67, 0x4a,
	0x3c, 0xe6, 0xbb, 0xb8, 0xa6, 0x1c, 0xb5, 0xc9, 0x53, 0x6a, 0x5a, 0x99, 0x67, 0x65, 0xb7, 0xaf,
	0xd8, 0xb7, 0xcd, 0x4c, 0xb5, 0x23, 0x31, 0x68, 0x24, 0x04, 0xbe, 0x1b, 0x89, 0x05, 0xa3, 0xaf,
	0x8c, 0xef, 0xff, 0x65, 0x74, 0x9a, 0x19, 0x09, 0x05, 0x28, 0xf1, 0x80, 0xef, 0x54, 0x60, 0x14,
	0x94, 0x9d, 0xbb, 0x0e, 0x89, 0xfb, 0xbc, 0xef, 0x8a, 0x38, 0x2f, 0x3e, 0xb5, 0xc9, 0xaf, 0xe9,
	0xc9, 0xfb, 0xdf, 0x9e, 0xfc, 0x6b, 0x7a, 0x9a, 0xbc, 0x3f, 0xdf, 0x84, 0xec, 0x62, 0x13, 0xb2,
	0x9f, 0x9b, 0x90, 0x7d, 0xd9, 0x86, 0xbd, 0x8b, 0x6d, 0xd8, 0xfb, 0xbe, 0x0d, 0x7b, 0x1f, 0x9e,
	0xfd, 0xb1, 0xa8, 0x57, 0xee, 0x60, 0x8e, 0x27, 0xf6, 0xed, 0xbf, 0x70, 0x89, 0x6a, 0x5d, 0x40,
	0xdc, 0xc4, 0x57, 0xe7, 0x66, 0xb7, 0x38, 0xdf, 0xb1, 0xb7, 0xf4, 0xf4, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0x6e, 0x11, 0xe9, 0x86, 0x02, 0x00, 0x00,
}

func (m *IDSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintPool(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BatchFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxCount != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.TxCount))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TotalFees.Size()
		i -= size
		if _, err := m.TotalFees.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdrawalReceived) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdrawalReceived) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdrawalReceived) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OutgoingTxId) > 0 {
		i -= len(m.OutgoingTxId)
		copy(dAtA[i:], m.OutgoingTxId)
		i = encodeVarintPool(dAtA, i, uint64(len(m.OutgoingTxId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BridgeChainId) > 0 {
		i -= len(m.BridgeChainId)
		copy(dAtA[i:], m.BridgeChainId)
		i = encodeVarintPool(dAtA, i, uint64(len(m.BridgeChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BridgeContract) > 0 {
		i -= len(m.BridgeContract)
		copy(dAtA[i:], m.BridgeContract)
		i = encodeVarintPool(dAtA, i, uint64(len(m.BridgeContract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdrawCanceled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdrawCanceled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdrawCanceled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BridgeChainId) > 0 {
		i -= len(m.BridgeChainId)
		copy(dAtA[i:], m.BridgeChainId)
		i = encodeVarintPool(dAtA, i, uint64(len(m.BridgeChainId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BridgeContract) > 0 {
		i -= len(m.BridgeContract)
		copy(dAtA[i:], m.BridgeContract)
		i = encodeVarintPool(dAtA, i, uint64(len(m.BridgeContract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintPool(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IDSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovPool(uint64(e))
		}
		n += 1 + sovPool(uint64(l)) + l
	}
	return n
}

func (m *BatchFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.TotalFees.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.TxCount != 0 {
		n += 1 + sovPool(uint64(m.TxCount))
	}
	return n
}

func (m *EventWithdrawalReceived) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BridgeContract)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.BridgeChainId)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.OutgoingTxId)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func (m *EventWithdrawCanceled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.BridgeContract)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.BridgeChainId)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IDSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: IDSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPool
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPool
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *BatchFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: BatchFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFees", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxCount", wireType)
			}
			m.TxCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *EventWithdrawalReceived) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: EventWithdrawalReceived: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdrawalReceived: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutgoingTxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutgoingTxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *EventWithdrawCanceled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: EventWithdrawCanceled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdrawCanceled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
