// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/common/info.proto

package common

import (
	fmt "fmt"
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

// Peers Active Type
type PeerType int32

const (
	PeerType_UNKNOWN     PeerType = 0
	PeerType_HIGHWAY     PeerType = 1
	PeerType_MOTOR       PeerType = 2
	PeerType_VALIDATOR   PeerType = 3
	PeerType_THIRD_PARTY PeerType = 4
)

var PeerType_name = map[int32]string{
	0: "UNKNOWN",
	1: "HIGHWAY",
	2: "MOTOR",
	3: "VALIDATOR",
	4: "THIRD_PARTY",
}

var PeerType_value = map[string]int32{
	"UNKNOWN":     0,
	"HIGHWAY":     1,
	"MOTOR":       2,
	"VALIDATOR":   3,
	"THIRD_PARTY": 4,
}

func (x PeerType) String() string {
	return proto.EnumName(PeerType_name, int32(x))
}

func (PeerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_117f1cca4f9b8f25, []int{0}
}

// Account is used for storing all credentials and their locations to be encrypted.
type AccountInfo struct {
	// Address is the associated Sonr address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Credentials is a list of all credentials associated with the account.
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// Label is the label of the account.
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// Index is the index of the account.
	Index uint32 `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	// Balance is the balance of the account.
	Balance int32 `protobuf:"varint,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_117f1cca4f9b8f25, []int{0}
}
func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return m.Size()
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *AccountInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AccountInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AccountInfo) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// Basic Info Sent to Peers to Establish Connections
type PeerInfo struct {
	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PeerId    string   `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Multiaddr string   `protobuf:"bytes,4,opt,name=multiaddr,proto3" json:"multiaddr,omitempty"`
	Type      PeerType `protobuf:"varint,5,opt,name=type,proto3,enum=sonrhq.sonr.common.PeerType" json:"type,omitempty"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_117f1cca4f9b8f25, []int{1}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PeerInfo) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerInfo) GetMultiaddr() string {
	if m != nil {
		return m.Multiaddr
	}
	return ""
}

func (m *PeerInfo) GetType() PeerType {
	if m != nil {
		return m.Type
	}
	return PeerType_UNKNOWN
}

type WalletInfo struct {
	// Controller is the associated Sonr address.
	Controller string `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	// DiscoverPaths is a list of all known hardened coin type paths.
	DiscoveredPaths []string `protobuf:"bytes,2,rep,name=discovered_paths,json=discoveredPaths,proto3" json:"discovered_paths,omitempty"`
	// Algorithm is the algorithm of the wallet. CMP is the default.
	Algorithm string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// CreatedAt is the time the wallet was created.
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// LastUpdated is the last time the wallet was updated.
	LastUpdated int64 `protobuf:"varint,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (m *WalletInfo) Reset()         { *m = WalletInfo{} }
func (m *WalletInfo) String() string { return proto.CompactTextString(m) }
func (*WalletInfo) ProtoMessage()    {}
func (*WalletInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_117f1cca4f9b8f25, []int{2}
}
func (m *WalletInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletInfo.Merge(m, src)
}
func (m *WalletInfo) XXX_Size() int {
	return m.Size()
}
func (m *WalletInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WalletInfo proto.InternalMessageInfo

func (m *WalletInfo) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *WalletInfo) GetDiscoveredPaths() []string {
	if m != nil {
		return m.DiscoveredPaths
	}
	return nil
}

func (m *WalletInfo) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *WalletInfo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *WalletInfo) GetLastUpdated() int64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrhq.sonr.common.PeerType", PeerType_name, PeerType_value)
	proto.RegisterType((*AccountInfo)(nil), "sonrhq.sonr.common.AccountInfo")
	proto.RegisterType((*PeerInfo)(nil), "sonrhq.sonr.common.PeerInfo")
	proto.RegisterType((*WalletInfo)(nil), "sonrhq.sonr.common.WalletInfo")
}

func init() { proto.RegisterFile("sonr/common/info.proto", fileDescriptor_117f1cca4f9b8f25) }

var fileDescriptor_117f1cca4f9b8f25 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x45, 0x33, 0x71, 0xd2, 0xe2, 0x17, 0xda, 0x5a, 0x23, 0x04, 0x5e, 0x14, 0x2b, 0x84, 0x4d,
	0x60, 0x61, 0x23, 0x58, 0xb3, 0x30, 0xaa, 0x44, 0x2c, 0x20, 0x89, 0x46, 0x29, 0x51, 0xd9, 0x58,
	0x13, 0xcf, 0xb4, 0xb1, 0x18, 0xcf, 0x98, 0xf1, 0x04, 0xe8, 0x17, 0xb0, 0x65, 0xc7, 0xa7, 0xf0,
	0x0b, 0x2c, 0xbb, 0x64, 0x89, 0x92, 0x1f, 0x41, 0x63, 0xbb, 0x0a, 0x12, 0x2b, 0xfb, 0x9e, 0x67,
	0xf9, 0xdd, 0x7b, 0xf5, 0xe0, 0x7e, 0xa5, 0xa4, 0x8e, 0x32, 0x55, 0x14, 0x4a, 0x46, 0xb9, 0xbc,
	0x54, 0x61, 0xa9, 0x95, 0x51, 0x18, 0x5b, 0xbe, 0xfe, 0x14, 0xda, 0x47, 0xd8, 0x8c, 0x47, 0xdf,
	0x10, 0x0c, 0xe2, 0x2c, 0x53, 0x1b, 0x69, 0x12, 0x79, 0xa9, 0xb0, 0x0f, 0x87, 0x94, 0x31, 0xcd,
	0xab, 0xca, 0x47, 0x43, 0x34, 0x76, 0xc9, 0xad, 0xb4, 0x13, 0xc9, 0xcd, 0x17, 0xa5, 0x3f, 0xfa,
	0xdd, 0x66, 0xd2, 0x4a, 0x7c, 0x0f, 0xfa, 0x82, 0xae, 0xb8, 0xf0, 0x9d, 0x9a, 0x37, 0xc2, 0xd2,
	0x5c, 0x32, 0xfe, 0xd5, 0xef, 0x0d, 0xd1, 0xf8, 0x88, 0x34, 0xc2, 0xfe, 0x65, 0x45, 0x05, 0x95,
	0x19, 0xf7, 0xfb, 0x43, 0x34, 0xee, 0x93, 0x5b, 0x39, 0xfa, 0x81, 0xe0, 0xce, 0x9c, 0x73, 0x5d,
	0xdb, 0x38, 0x86, 0x6e, 0xce, 0x5a, 0x07, 0xdd, 0x9c, 0x61, 0x0c, 0x3d, 0x49, 0x0b, 0xde, 0x6e,
	0xae, 0xdf, 0xf1, 0x03, 0x38, 0x2c, 0x39, 0xd7, 0x69, 0xce, 0xda, 0xc5, 0x07, 0x56, 0x26, 0x0c,
	0x9f, 0x82, 0x5b, 0x6c, 0x84, 0xc9, 0xad, 0xf3, 0x7a, 0xbb, 0x4b, 0xf6, 0x00, 0x3f, 0x83, 0x9e,
	0xb9, 0x2e, 0x9b, 0xf5, 0xc7, 0xcf, 0x4f, 0xc3, 0xff, 0x4b, 0x09, 0xad, 0x8d, 0xc5, 0x75, 0xc9,
	0x49, 0xfd, 0xe5, 0xe8, 0x27, 0x02, 0x58, 0x52, 0x21, 0x78, 0x53, 0x51, 0x00, 0x90, 0x29, 0x69,
	0xb4, 0x12, 0x82, 0xeb, 0xd6, 0xe3, 0x3f, 0x04, 0x3f, 0x01, 0x8f, 0xe5, 0x55, 0xa6, 0x3e, 0x73,
	0xcd, 0x59, 0x5a, 0x52, 0xb3, 0xae, 0xfc, 0xee, 0xd0, 0x19, 0xbb, 0xe4, 0x64, 0xcf, 0xe7, 0x16,
	0x5b, 0xa7, 0x54, 0x5c, 0x29, 0x9d, 0x9b, 0x75, 0xd1, 0x86, 0xd8, 0x03, 0xfc, 0x10, 0x20, 0xd3,
	0x9c, 0x1a, 0xce, 0x52, 0x6a, 0xea, 0x20, 0x0e, 0x71, 0x5b, 0x12, 0x1b, 0xfc, 0x08, 0xee, 0x0a,
	0x5a, 0x99, 0x74, 0x53, 0x32, 0x4b, 0xea, 0x40, 0x0e, 0x19, 0x58, 0x76, 0xde, 0xa0, 0xa7, 0xb3,
	0xa6, 0x52, 0x9b, 0x05, 0x0f, 0xe0, 0xf0, 0x7c, 0xfa, 0x66, 0x3a, 0x5b, 0x4e, 0xbd, 0x8e, 0x15,
	0x93, 0xe4, 0xf5, 0x64, 0x19, 0x5f, 0x78, 0x08, 0xbb, 0xd0, 0x7f, 0x37, 0x5b, 0xcc, 0x88, 0xd7,
	0xc5, 0x47, 0xe0, 0xbe, 0x8f, 0xdf, 0x26, 0x67, 0xb1, 0x95, 0x0e, 0x3e, 0x81, 0xc1, 0x62, 0x92,
	0x90, 0xb3, 0x74, 0x1e, 0x93, 0xc5, 0x85, 0xd7, 0x7b, 0xf5, 0xf2, 0xd7, 0x36, 0x40, 0x37, 0xdb,
	0x00, 0xfd, 0xd9, 0x06, 0xe8, 0xfb, 0x2e, 0xe8, 0xdc, 0xec, 0x82, 0xce, 0xef, 0x5d, 0xd0, 0xf9,
	0xf0, 0xf8, 0x2a, 0x37, 0xeb, 0xcd, 0xca, 0x56, 0x18, 0x35, 0x95, 0x46, 0x99, 0xd2, 0x3c, 0xb2,
	0x0d, 0x56, 0xed, 0x31, 0xae, 0x0e, 0xea, 0x43, 0x7c, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0xc8, 0x83, 0x2a, 0xa2, 0x02, 0x00, 0x00,
}

func (m *AccountInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Balance != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PeerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Multiaddr) > 0 {
		i -= len(m.Multiaddr)
		copy(dAtA[i:], m.Multiaddr)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Multiaddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WalletInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdated != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.LastUpdated))
		i--
		dAtA[i] = 0x28
	}
	if m.CreatedAt != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Algorithm) > 0 {
		i -= len(m.Algorithm)
		copy(dAtA[i:], m.Algorithm)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Algorithm)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DiscoveredPaths) > 0 {
		for iNdEx := len(m.DiscoveredPaths) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DiscoveredPaths[iNdEx])
			copy(dAtA[i:], m.DiscoveredPaths[iNdEx])
			i = encodeVarintInfo(dAtA, i, uint64(len(m.DiscoveredPaths[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovInfo(uint64(m.Index))
	}
	if m.Balance != 0 {
		n += 1 + sovInfo(uint64(m.Balance))
	}
	return n
}

func (m *PeerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Multiaddr)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovInfo(uint64(m.Type))
	}
	return n
}

func (m *WalletInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if len(m.DiscoveredPaths) > 0 {
		for _, s := range m.DiscoveredPaths {
			l = len(s)
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	l = len(m.Algorithm)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovInfo(uint64(m.CreatedAt))
	}
	if m.LastUpdated != 0 {
		n += 1 + sovInfo(uint64(m.LastUpdated))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: AccountInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *PeerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: PeerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiaddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multiaddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PeerType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *WalletInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: WalletInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscoveredPaths", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DiscoveredPaths = append(m.DiscoveredPaths, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Algorithm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Algorithm = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			m.LastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdated |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)
