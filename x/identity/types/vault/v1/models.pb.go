// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/vault/v1/models.proto

// Package Motor is used for defining a Motor node and its properties.

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	return fileDescriptor_d53fa4614f4f387b, []int{0}
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

type AccountConfig struct {
	// Name is the name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// multibase is the associated pubkey encoded in multibase.
	Multibase string `protobuf:"bytes,2,opt,name=multibase,proto3" json:"multibase,omitempty"`
	// PublicKey is the public key of the account.
	PublicKey []byte `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// CoinType is the coin type of the account. See https://github.com/satoshilabs/slips/blob/master/slip-0044.md for more information.
	CoinTypeIndex int32 `protobuf:"varint,4,opt,name=coin_type_index,json=coinTypeIndex,proto3" json:"coin_type_index,omitempty"`
	// CreatedAt is the time the account was created.
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Shares is a list of all shares associated with the account.
	Shares [][]byte `protobuf:"bytes,6,rep,name=shares,proto3" json:"shares,omitempty"`
}

func (m *AccountConfig) Reset()         { *m = AccountConfig{} }
func (m *AccountConfig) String() string { return proto.CompactTextString(m) }
func (*AccountConfig) ProtoMessage()    {}
func (*AccountConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d53fa4614f4f387b, []int{1}
}
func (m *AccountConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountConfig.Merge(m, src)
}
func (m *AccountConfig) XXX_Size() int {
	return m.Size()
}
func (m *AccountConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AccountConfig proto.InternalMessageInfo

func (m *AccountConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountConfig) GetMultibase() string {
	if m != nil {
		return m.Multibase
	}
	return ""
}

func (m *AccountConfig) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *AccountConfig) GetCoinTypeIndex() int32 {
	if m != nil {
		return m.CoinTypeIndex
	}
	return 0
}

func (m *AccountConfig) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AccountConfig) GetShares() [][]byte {
	if m != nil {
		return m.Shares
	}
	return nil
}

type WalletConfig struct {
	// Address is the associated blockchain address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// PublicKey is the public key of the wallet.
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Algorithm is the algorithm of the wallet.
	Algorithm string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// CID is the CID of the wallet.
	Cid string `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
	// Accounts is the map of accounts associated with the wallet.
	Accounts map[string]*AccountConfig `protobuf:"bytes,5,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *WalletConfig) Reset()         { *m = WalletConfig{} }
func (m *WalletConfig) String() string { return proto.CompactTextString(m) }
func (*WalletConfig) ProtoMessage()    {}
func (*WalletConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d53fa4614f4f387b, []int{2}
}
func (m *WalletConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletConfig.Merge(m, src)
}
func (m *WalletConfig) XXX_Size() int {
	return m.Size()
}
func (m *WalletConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WalletConfig proto.InternalMessageInfo

func (m *WalletConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WalletConfig) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *WalletConfig) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *WalletConfig) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *WalletConfig) GetAccounts() map[string]*AccountConfig {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountInfo)(nil), "sonrhq.protocol.vault.v1.AccountInfo")
	proto.RegisterType((*AccountConfig)(nil), "sonrhq.protocol.vault.v1.AccountConfig")
	proto.RegisterType((*WalletConfig)(nil), "sonrhq.protocol.vault.v1.WalletConfig")
	proto.RegisterMapType((map[string]*AccountConfig)(nil), "sonrhq.protocol.vault.v1.WalletConfig.AccountsEntry")
}

func init() { proto.RegisterFile("protocol/vault/v1/models.proto", fileDescriptor_d53fa4614f4f387b) }

var fileDescriptor_d53fa4614f4f387b = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xf6, 0x52, 0x88, 0xdb, 0x0a, 0x64, 0x21, 0x14, 0x9d, 0x7a, 0x51, 0xd5, 0x01,
	0x32, 0xc5, 0xdc, 0xc1, 0x80, 0x90, 0x18, 0x0e, 0xc4, 0x70, 0xb0, 0x20, 0x0b, 0x09, 0x89, 0xa5,
	0x72, 0x1c, 0x5f, 0x6a, 0x9d, 0x63, 0x87, 0xd8, 0x29, 0x97, 0x4f, 0xc0, 0xca, 0x37, 0xe1, 0x23,
	0xb0, 0x32, 0xde, 0xc8, 0x88, 0xda, 0x2f, 0x82, 0x12, 0xa7, 0x2d, 0x77, 0x52, 0xb7, 0xf7, 0xff,
	0x3f, 0xfb, 0xbd, 0x9f, 0xdf, 0x4b, 0x60, 0x58, 0x94, 0xda, 0x6a, 0xa6, 0x25, 0x5e, 0xd1, 0x4a,
	0x5a, 0xbc, 0x3a, 0xc5, 0xb9, 0x4e, 0xb9, 0x34, 0x71, 0x9b, 0x40, 0x81, 0xd1, 0xaa, 0x5c, 0x7e,
	0x8d, 0xb7, 0xc7, 0xe2, 0xf6, 0x58, 0xbc, 0x3a, 0x3d, 0x9e, 0x66, 0x5a, 0x67, 0x92, 0x63, 0x5a,
	0x08, 0x4c, 0x95, 0xd2, 0x96, 0x5a, 0xa1, 0x55, 0x77, 0x6f, 0xfe, 0x1d, 0xc0, 0xd1, 0x39, 0x63,
	0xba, 0x52, 0xf6, 0x42, 0x5d, 0x6a, 0x14, 0xc0, 0x7b, 0x34, 0x4d, 0x4b, 0x6e, 0x4c, 0x00, 0x66,
	0x20, 0xf2, 0xc9, 0x56, 0x36, 0x19, 0xc5, 0xed, 0x37, 0x5d, 0x5e, 0x05, 0x7d, 0x97, 0xe9, 0x24,
	0x7a, 0x04, 0x3d, 0x49, 0x13, 0x2e, 0x83, 0x41, 0xeb, 0x3b, 0xd1, 0xb8, 0x42, 0xa5, 0xfc, 0x3a,
	0x38, 0x9a, 0x81, 0x68, 0x42, 0x9c, 0x68, 0xaa, 0x24, 0x54, 0x52, 0xc5, 0x78, 0xe0, 0xcd, 0x40,
	0xe4, 0x91, 0xad, 0x9c, 0xff, 0x02, 0x70, 0xd2, 0x91, 0xbc, 0xd5, 0xea, 0x52, 0x64, 0x08, 0xc1,
	0x23, 0x45, 0x73, 0xde, 0x81, 0xb4, 0x31, 0x9a, 0x42, 0x3f, 0xaf, 0xa4, 0x15, 0x09, 0x35, 0xbc,
	0xe3, 0xd8, 0x1b, 0xe8, 0x04, 0xc2, 0xa2, 0x4a, 0xa4, 0x60, 0x8b, 0x2b, 0x5e, 0xb7, 0x38, 0x63,
	0xe2, 0x3b, 0xe7, 0x03, 0xaf, 0xd1, 0x13, 0xf8, 0x80, 0x69, 0xa1, 0x16, 0xb6, 0x2e, 0xf8, 0x62,
	0x0f, 0xe7, 0x91, 0x49, 0x63, 0x7f, 0xaa, 0x0b, 0x7e, 0xd1, 0x42, 0x9e, 0x40, 0xc8, 0x4a, 0x4e,
	0x2d, 0x4f, 0x17, 0xd4, 0xb6, 0x9c, 0x03, 0xe2, 0x77, 0xce, 0xb9, 0x45, 0x8f, 0xe1, 0xd0, 0x2c,
	0x69, 0xc9, 0x4d, 0x30, 0x9c, 0x0d, 0xa2, 0x31, 0xe9, 0xd4, 0xfc, 0x67, 0x1f, 0x8e, 0x3f, 0x53,
	0x29, 0xf9, 0xf6, 0x01, 0x87, 0x87, 0x79, 0x1b, 0xb4, 0x7f, 0x17, 0x74, 0x0a, 0x7d, 0x2a, 0x33,
	0x5d, 0x0a, 0xbb, 0xcc, 0xbb, 0xa9, 0xee, 0x0d, 0xf4, 0x10, 0x0e, 0x98, 0x48, 0x5b, 0x74, 0x9f,
	0x34, 0x21, 0xfa, 0x08, 0xef, 0x53, 0x37, 0x3a, 0x13, 0x78, 0xb3, 0x41, 0x34, 0x3a, 0x7b, 0x11,
	0x1f, 0xfa, 0x20, 0xe2, 0xff, 0x11, 0xe3, 0x6e, 0xe2, 0xe6, 0x9d, 0xb2, 0x65, 0x4d, 0x76, 0x55,
	0x8e, 0xd3, 0xdd, 0x32, 0x5c, 0xaa, 0x69, 0xda, 0xa0, 0xba, 0x77, 0x34, 0x21, 0x7a, 0x0d, 0xbd,
	0x15, 0x95, 0x95, 0x5b, 0xc3, 0xe8, 0xec, 0xe9, 0xe1, 0x8e, 0xb7, 0xd6, 0x4a, 0xdc, 0xad, 0x57,
	0xfd, 0x97, 0xe0, 0xcd, 0xfb, 0xdf, 0xeb, 0x10, 0xdc, 0xac, 0x43, 0xf0, 0x77, 0x1d, 0x82, 0x1f,
	0x9b, 0xb0, 0x77, 0xb3, 0x09, 0x7b, 0x7f, 0x36, 0x61, 0xef, 0xcb, 0xb3, 0x4c, 0xd8, 0x65, 0x95,
	0xc4, 0x4c, 0xe7, 0xd8, 0xd5, 0xc5, 0x4c, 0x97, 0x1c, 0x5f, 0x63, 0x91, 0x72, 0x65, 0x85, 0xad,
	0x71, 0xb3, 0x45, 0xb3, 0xfb, 0x1d, 0x92, 0x61, 0xdb, 0xf7, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x30, 0x3a, 0x01, 0x2a, 0x03, 0x00, 0x00,
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
		i = encodeVarintModels(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Shares[iNdEx])
			copy(dAtA[i:], m.Shares[iNdEx])
			i = encodeVarintModels(dAtA, i, uint64(len(m.Shares[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.CreatedAt != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x28
	}
	if m.CoinTypeIndex != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.CoinTypeIndex))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintModels(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Multibase) > 0 {
		i -= len(m.Multibase)
		copy(dAtA[i:], m.Multibase)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Multibase)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WalletConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for k := range m.Accounts {
			v := m.Accounts[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintModels(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintModels(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintModels(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Algorithm) > 0 {
		i -= len(m.Algorithm)
		copy(dAtA[i:], m.Algorithm)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Algorithm)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintModels(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
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
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovModels(uint64(m.Index))
	}
	if m.Balance != 0 {
		n += 1 + sovModels(uint64(m.Balance))
	}
	return n
}

func (m *AccountConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Multibase)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.CoinTypeIndex != 0 {
		n += 1 + sovModels(uint64(m.CoinTypeIndex))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovModels(uint64(m.CreatedAt))
	}
	if len(m.Shares) > 0 {
		for _, b := range m.Shares {
			l = len(b)
			n += 1 + l + sovModels(uint64(l))
		}
	}
	return n
}

func (m *WalletConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Algorithm)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if len(m.Accounts) > 0 {
		for k, v := range m.Accounts {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovModels(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovModels(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovModels(uint64(mapEntrySize))
		}
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
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
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
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
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
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
					return ErrIntOverflowModels
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
					return ErrIntOverflowModels
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
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *AccountConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: AccountConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multibase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multibase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinTypeIndex", wireType)
			}
			m.CoinTypeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinTypeIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, make([]byte, postIndex-iNdEx))
			copy(m.Shares[len(m.Shares)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *WalletConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: WalletConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Algorithm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Algorithm = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Accounts == nil {
				m.Accounts = make(map[string]*AccountConfig)
			}
			var mapkey string
			var mapvalue *AccountConfig
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowModels
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowModels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthModels
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthModels
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowModels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthModels
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthModels
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &AccountConfig{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipModels(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthModels
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Accounts[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)
