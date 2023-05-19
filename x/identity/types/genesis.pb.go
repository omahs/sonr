// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/identity/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the identity module's genesis state.
type GenesisState struct {
	Params               Params                     `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DidDocuments         []Identification           `protobuf:"bytes,2,rep,name=did_documents,json=didDocuments,proto3" json:"did_documents"`
	Relationships        []VerificationRelationship `protobuf:"bytes,3,rep,name=relationships,proto3" json:"relationships"`
	ClaimableWalletList  []ClaimableWallet          `protobuf:"bytes,4,rep,name=claimable_wallet_list,json=claimableWalletList,proto3" json:"claimable_wallet_list"`
	ClaimableWalletCount uint64                     `protobuf:"varint,5,opt,name=claimable_wallet_count,json=claimableWalletCount,proto3" json:"claimable_wallet_count,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee3e6e2aad889c, []int{0}
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

func (m *GenesisState) GetDidDocuments() []Identification {
	if m != nil {
		return m.DidDocuments
	}
	return nil
}

func (m *GenesisState) GetRelationships() []VerificationRelationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

func (m *GenesisState) GetClaimableWalletList() []ClaimableWallet {
	if m != nil {
		return m.ClaimableWalletList
	}
	return nil
}

func (m *GenesisState) GetClaimableWalletCount() uint64 {
	if m != nil {
		return m.ClaimableWalletCount
	}
	return 0
}

// Params defines the parameters for the module.
type Params struct {
	AccountDidMethodName    string   `protobuf:"bytes,1,opt,name=account_did_method_name,json=accountDidMethodName,proto3" json:"account_did_method_name,omitempty"`
	AccountDidMethodContext string   `protobuf:"bytes,2,opt,name=account_did_method_context,json=accountDidMethodContext,proto3" json:"account_did_method_context,omitempty"`
	AcccountDiscoveryReward int64    `protobuf:"varint,3,opt,name=acccount_discovery_reward,json=acccountDiscoveryReward,proto3" json:"acccount_discovery_reward,omitempty"`
	DidBaseContext          string   `protobuf:"bytes,4,opt,name=did_base_context,json=didBaseContext,proto3" json:"did_base_context,omitempty"`
	MaximumIdentityAliases  int32    `protobuf:"varint,5,opt,name=maximum_identity_aliases,json=maximumIdentityAliases,proto3" json:"maximum_identity_aliases,omitempty"`
	SupportedDidMethods     []string `protobuf:"bytes,6,rep,name=supported_did_methods,json=supportedDidMethods,proto3" json:"supported_did_methods,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee3e6e2aad889c, []int{1}
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

func (m *Params) GetAccountDidMethodName() string {
	if m != nil {
		return m.AccountDidMethodName
	}
	return ""
}

func (m *Params) GetAccountDidMethodContext() string {
	if m != nil {
		return m.AccountDidMethodContext
	}
	return ""
}

func (m *Params) GetAcccountDiscoveryReward() int64 {
	if m != nil {
		return m.AcccountDiscoveryReward
	}
	return 0
}

func (m *Params) GetDidBaseContext() string {
	if m != nil {
		return m.DidBaseContext
	}
	return ""
}

func (m *Params) GetMaximumIdentityAliases() int32 {
	if m != nil {
		return m.MaximumIdentityAliases
	}
	return 0
}

func (m *Params) GetSupportedDidMethods() []string {
	if m != nil {
		return m.SupportedDidMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sonrhq.core.identity.GenesisState")
	proto.RegisterType((*Params)(nil), "sonrhq.core.identity.Params")
}

func init() { proto.RegisterFile("core/identity/genesis.proto", fileDescriptor_22ee3e6e2aad889c) }

var fileDescriptor_22ee3e6e2aad889c = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0xb3, 0xdd, 0x34, 0xd0, 0x69, 0x2b, 0x32, 0x4d, 0x9b, 0x35, 0xca, 0x1a, 0x8a, 0xe2,
	0x9e, 0x36, 0x10, 0x15, 0xa4, 0x9e, 0x9a, 0x14, 0xa4, 0xe0, 0x3f, 0x56, 0x50, 0xe8, 0x65, 0x99,
	0xcc, 0xbc, 0x26, 0x03, 0x3b, 0x3b, 0xeb, 0xcc, 0xc4, 0x26, 0x5f, 0x42, 0xfc, 0x58, 0x3d, 0xe6,
	0xe8, 0x49, 0x24, 0xf9, 0x22, 0x92, 0xd9, 0x49, 0x6c, 0xea, 0xde, 0x96, 0xfc, 0xde, 0xdf, 0xf3,
	0x64, 0xde, 0x61, 0xd0, 0x43, 0x2a, 0x15, 0x74, 0x39, 0x83, 0xdc, 0x70, 0x33, 0xeb, 0x8e, 0x20,
	0x07, 0xcd, 0x75, 0x5c, 0x28, 0x69, 0x24, 0x6e, 0x6a, 0x99, 0xab, 0xf1, 0xb7, 0x78, 0x35, 0x13,
	0xaf, 0x67, 0xda, 0xcd, 0x91, 0x1c, 0x49, 0x3b, 0xd0, 0x5d, 0x7d, 0x95, 0xb3, 0xed, 0xd6, 0x76,
	0x10, 0xe3, 0xcc, 0x81, 0xf6, 0x36, 0xa0, 0x19, 0xe1, 0xc2, 0x15, 0x9c, 0xfe, 0xf0, 0xd1, 0xc1,
	0x9b, 0xb2, 0xf2, 0x93, 0x21, 0x06, 0xf0, 0x19, 0x6a, 0x14, 0x44, 0x11, 0xa1, 0x03, 0xaf, 0xe3,
	0x45, 0xfb, 0xbd, 0x47, 0x71, 0xd5, 0x5f, 0x88, 0x3f, 0xda, 0x99, 0x7e, 0xfd, 0xe6, 0xf7, 0xe3,
	0x5a, 0xe2, 0x0c, 0xfc, 0x01, 0x1d, 0x32, 0xce, 0x52, 0x26, 0xe9, 0x44, 0x40, 0x6e, 0x74, 0xb0,
	0xd3, 0xf1, 0xa3, 0xfd, 0xde, 0x93, 0xea, 0x88, 0x4b, 0xfb, 0xf1, 0x95, 0x53, 0x62, 0xb8, 0xcc,
	0x5d, 0xd4, 0x01, 0xe3, 0xec, 0x62, 0xed, 0xe3, 0x2b, 0x74, 0xa8, 0x20, 0xb3, 0x5c, 0x8f, 0x79,
	0xa1, 0x03, 0xdf, 0x06, 0xc6, 0xd5, 0x81, 0x9f, 0x41, 0x6d, 0xe2, 0x92, 0x5b, 0x9a, 0x8b, 0xde,
	0x8e, 0xc2, 0x29, 0x3a, 0xb6, 0x9b, 0x20, 0xc3, 0x0c, 0xd2, 0x6b, 0x92, 0x65, 0x60, 0xd2, 0x8c,
	0x6b, 0x13, 0xd4, 0x6d, 0xc7, 0xd3, 0xea, 0x8e, 0xc1, 0x5a, 0xf9, 0x62, 0x0d, 0x17, 0x7d, 0x44,
	0xb7, 0x7f, 0x7e, 0xcb, 0xb5, 0xc1, 0x2f, 0xd0, 0xc9, 0x7f, 0x05, 0x54, 0x4e, 0x72, 0x13, 0xec,
	0x76, 0xbc, 0xa8, 0x9e, 0x34, 0xef, 0x48, 0x83, 0x15, 0x3b, 0x9d, 0xef, 0xa0, 0x46, 0xb9, 0x5c,
	0xfc, 0x12, 0xb5, 0x08, 0xb5, 0x46, 0xba, 0x5a, 0xab, 0x00, 0x33, 0x96, 0x2c, 0xcd, 0x89, 0x00,
	0x7b, 0x37, 0x7b, 0x49, 0xd3, 0xe1, 0x0b, 0xce, 0xde, 0x59, 0xf8, 0x9e, 0x08, 0xc0, 0xaf, 0x51,
	0xbb, 0x42, 0xa3, 0x32, 0x37, 0x30, 0x35, 0xc1, 0x8e, 0x35, 0x5b, 0x77, 0xcd, 0x41, 0x89, 0xf1,
	0x19, 0x7a, 0x40, 0xe8, 0xc6, 0xd6, 0x54, 0x7e, 0x07, 0x35, 0x4b, 0x15, 0x5c, 0x13, 0xc5, 0x02,
	0xbf, 0xe3, 0x45, 0xbe, 0x75, 0x9d, 0xec, 0x78, 0x62, 0x31, 0x8e, 0xd0, 0xfd, 0x55, 0xe1, 0x90,
	0x68, 0xd8, 0xd4, 0xd5, 0x6d, 0xdd, 0x3d, 0xc6, 0x59, 0x9f, 0x68, 0x58, 0xb7, 0xbc, 0x42, 0x81,
	0x20, 0x53, 0x2e, 0x26, 0x22, 0x5d, 0x6f, 0x36, 0x25, 0x19, 0x27, 0x1a, 0xb4, 0x5d, 0xce, 0x6e,
	0x72, 0xe2, 0xf8, 0xa5, 0xc3, 0xe7, 0x25, 0xc5, 0x3d, 0x74, 0xac, 0x27, 0x45, 0x21, 0x95, 0x01,
	0x76, 0xeb, 0x78, 0x3a, 0x68, 0x74, 0xfc, 0x68, 0x2f, 0x39, 0xda, 0xc0, 0xcd, 0xc9, 0x74, 0xff,
	0xfc, 0x66, 0x11, 0x7a, 0xf3, 0x45, 0xe8, 0xfd, 0x59, 0x84, 0xde, 0xcf, 0x65, 0x58, 0x9b, 0x2f,
	0xc3, 0xda, 0xaf, 0x65, 0x58, 0xbb, 0x7a, 0x36, 0xe2, 0x66, 0x3c, 0x19, 0xc6, 0x54, 0x8a, 0x6e,
	0x79, 0xdd, 0x5d, 0xfb, 0x56, 0xa6, 0xff, 0x5e, 0x8b, 0x99, 0x15, 0xa0, 0x87, 0x0d, 0xfb, 0x5a,
	0x9e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x19, 0xb6, 0x4b, 0xad, 0x03, 0x00, 0x00,
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
	if m.ClaimableWalletCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ClaimableWalletCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ClaimableWalletList) > 0 {
		for iNdEx := len(m.ClaimableWalletList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimableWalletList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DidDocuments) > 0 {
		for iNdEx := len(m.DidDocuments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidDocuments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SupportedDidMethods) > 0 {
		for iNdEx := len(m.SupportedDidMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedDidMethods[iNdEx])
			copy(dAtA[i:], m.SupportedDidMethods[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.SupportedDidMethods[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MaximumIdentityAliases != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaximumIdentityAliases))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DidBaseContext) > 0 {
		i -= len(m.DidBaseContext)
		copy(dAtA[i:], m.DidBaseContext)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DidBaseContext)))
		i--
		dAtA[i] = 0x22
	}
	if m.AcccountDiscoveryReward != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AcccountDiscoveryReward))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AccountDidMethodContext) > 0 {
		i -= len(m.AccountDidMethodContext)
		copy(dAtA[i:], m.AccountDidMethodContext)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AccountDidMethodContext)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountDidMethodName) > 0 {
		i -= len(m.AccountDidMethodName)
		copy(dAtA[i:], m.AccountDidMethodName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AccountDidMethodName)))
		i--
		dAtA[i] = 0xa
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
	if len(m.DidDocuments) > 0 {
		for _, e := range m.DidDocuments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClaimableWalletList) > 0 {
		for _, e := range m.ClaimableWalletList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ClaimableWalletCount != 0 {
		n += 1 + sovGenesis(uint64(m.ClaimableWalletCount))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountDidMethodName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AccountDidMethodContext)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.AcccountDiscoveryReward != 0 {
		n += 1 + sovGenesis(uint64(m.AcccountDiscoveryReward))
	}
	l = len(m.DidBaseContext)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaximumIdentityAliases != 0 {
		n += 1 + sovGenesis(uint64(m.MaximumIdentityAliases))
	}
	if len(m.SupportedDidMethods) > 0 {
		for _, s := range m.SupportedDidMethods {
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
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocuments", wireType)
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
			m.DidDocuments = append(m.DidDocuments, Identification{})
			if err := m.DidDocuments[len(m.DidDocuments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
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
			m.Relationships = append(m.Relationships, VerificationRelationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableWalletList", wireType)
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
			m.ClaimableWalletList = append(m.ClaimableWalletList, ClaimableWallet{})
			if err := m.ClaimableWalletList[len(m.ClaimableWalletList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableWalletCount", wireType)
			}
			m.ClaimableWalletCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimableWalletCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDidMethodName", wireType)
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
			m.AccountDidMethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDidMethodContext", wireType)
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
			m.AccountDidMethodContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcccountDiscoveryReward", wireType)
			}
			m.AcccountDiscoveryReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AcccountDiscoveryReward |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidBaseContext", wireType)
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
			m.DidBaseContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumIdentityAliases", wireType)
			}
			m.MaximumIdentityAliases = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumIdentityAliases |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedDidMethods", wireType)
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
			m.SupportedDidMethods = append(m.SupportedDidMethods, string(dAtA[iNdEx:postIndex]))
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
