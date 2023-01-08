// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: highway/vault/v1/models.proto

// Package Motor is used for defining a Motor node and its properties.

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/sonr-hq/sonr/pkg/common"
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
type Account struct {
	// Address is the associated Sonr address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Credentials is a list of all credentials associated with the account.
	Credentials map[string]*Credential `protobuf:"bytes,2,rep,name=credentials,proto3" json:"credentials,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99610c26b8ba83d, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetCredentials() map[string]*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

// Credential is used for storing a single credential.
type Credential struct {
	// Name is the name of the credential.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type is the type of the credential. (e.g. "biometric", "webauthn")
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Webauthn is for PassKey data of the credential.
	Credential []byte `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99610c26b8ba83d, []int{1}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Credential) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Credential) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

// Session is used for caching a session.
type Session struct {
	// Id is the session id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Challenge is the challenge used for the session.
	Challenge string `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	// RpId is the relying party id specified by the client.
	RpId string `protobuf:"bytes,3,opt,name=rp_id,json=rpId,proto3" json:"rp_id,omitempty"`
	// Origins is the list of origins specified by the client.
	RpOrigins []string `protobuf:"bytes,4,rep,name=rp_origins,json=rpOrigins,proto3" json:"rp_origins,omitempty"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_d99610c26b8ba83d, []int{2}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func (m *Session) GetRpId() string {
	if m != nil {
		return m.RpId
	}
	return ""
}

func (m *Session) GetRpOrigins() []string {
	if m != nil {
		return m.RpOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "sonrhq.highway.vault.v1.Account")
	proto.RegisterMapType((map[string]*Credential)(nil), "sonrhq.highway.vault.v1.Account.CredentialsEntry")
	proto.RegisterType((*Credential)(nil), "sonrhq.highway.vault.v1.Credential")
	proto.RegisterType((*Session)(nil), "sonrhq.highway.vault.v1.Session")
}

func init() { proto.RegisterFile("highway/vault/v1/models.proto", fileDescriptor_d99610c26b8ba83d) }

var fileDescriptor_d99610c26b8ba83d = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x23, 0x27, 0x5d, 0xc8, 0x9b, 0x31, 0x8a, 0x36, 0x98, 0x09, 0xad, 0x31, 0xd9, 0x25,
	0x97, 0x49, 0xa4, 0xbb, 0x6c, 0xbd, 0x6d, 0x63, 0x87, 0x9d, 0x06, 0x6e, 0x4f, 0xbb, 0x04, 0xd5,
	0x12, 0xb6, 0xa8, 0x2d, 0xa9, 0x92, 0xec, 0xe1, 0x6f, 0xb1, 0x8f, 0xb5, 0x63, 0x8f, 0xbb, 0x0c,
	0x46, 0xf2, 0x45, 0x8a, 0xff, 0x14, 0x97, 0x42, 0x4e, 0x7e, 0xfd, 0xe8, 0x79, 0x7e, 0xef, 0xab,
	0x3f, 0x70, 0x9e, 0xcb, 0x2c, 0xff, 0xc5, 0x1a, 0x5a, 0xb3, 0xaa, 0xf0, 0xb4, 0xde, 0xd2, 0x52,
	0x73, 0x51, 0x38, 0x62, 0xac, 0xf6, 0x1a, 0xbf, 0x75, 0x5a, 0xd9, 0xfc, 0x8e, 0x0c, 0x2e, 0xd2,
	0xb9, 0x48, 0xbd, 0x5d, 0xbd, 0x49, 0x75, 0x59, 0x6a, 0xd5, 0x06, 0x38, 0xf3, 0xac, 0xb7, 0xaf,
	0xce, 0x32, 0xad, 0xb3, 0x42, 0x50, 0x66, 0x24, 0x65, 0x4a, 0x69, 0xcf, 0xbc, 0xd4, 0x6a, 0x80,
	0xad, 0xff, 0x21, 0x98, 0x7f, 0x4e, 0x53, 0x5d, 0x29, 0x8f, 0x43, 0x98, 0x33, 0xce, 0xad, 0x70,
	0x2e, 0x44, 0x31, 0xda, 0x2c, 0x92, 0xc7, 0x5f, 0x7c, 0x05, 0xcb, 0xd4, 0x0a, 0x2e, 0x94, 0x97,
	0xac, 0x70, 0x61, 0x10, 0x4f, 0x37, 0xcb, 0x8b, 0x2d, 0x39, 0x32, 0x08, 0x19, 0x80, 0xe4, 0xeb,
	0x98, 0xf9, 0xa6, 0xbc, 0x6d, 0x92, 0xa7, 0x94, 0x55, 0x0a, 0xa7, 0xcf, 0x0d, 0xf8, 0x14, 0xa6,
	0xb7, 0xa2, 0x19, 0xda, 0xb7, 0x25, 0xfe, 0x04, 0x27, 0x35, 0x2b, 0x2a, 0x11, 0x06, 0x31, 0xda,
	0x2c, 0x2f, 0xde, 0x1d, 0x6d, 0x3a, 0xb2, 0x92, 0x3e, 0x71, 0x19, 0x7c, 0x44, 0xeb, 0x6b, 0x80,
	0x71, 0x01, 0x63, 0x98, 0x29, 0x56, 0x8a, 0x81, 0xdf, 0xd5, 0xad, 0xe6, 0x1b, 0xd3, 0xf3, 0x17,
	0x49, 0x57, 0xe3, 0x08, 0x60, 0x9c, 0x34, 0x9c, 0xc6, 0x68, 0xf3, 0x32, 0x79, 0xa2, 0xac, 0x6f,
	0x61, 0x7e, 0x25, 0x9c, 0x93, 0x5a, 0xe1, 0x57, 0x10, 0x48, 0x3e, 0x00, 0x03, 0xc9, 0xf1, 0x19,
	0x2c, 0xd2, 0x9c, 0x15, 0x85, 0x50, 0xd9, 0x23, 0x73, 0x14, 0xf0, 0x6b, 0x38, 0xb1, 0x66, 0x27,
	0x79, 0xc7, 0x5c, 0x24, 0x33, 0x6b, 0xbe, 0x73, 0x7c, 0x0e, 0x60, 0xcd, 0x4e, 0x5b, 0x99, 0x49,
	0xe5, 0xc2, 0x59, 0x3c, 0x6d, 0x33, 0xd6, 0xfc, 0xe8, 0x85, 0x2f, 0xd7, 0x7f, 0xf6, 0x11, 0xba,
	0xdf, 0x47, 0xe8, 0xff, 0x3e, 0x42, 0xbf, 0x0f, 0xd1, 0xe4, 0xfe, 0x10, 0x4d, 0xfe, 0x1e, 0xa2,
	0xc9, 0xcf, 0xcb, 0x4c, 0xfa, 0xbc, 0xba, 0x21, 0xa9, 0x2e, 0x69, 0x7b, 0x2c, 0xef, 0xf3, 0xbb,
	0xee, 0x4b, 0x7d, 0x2e, 0x2d, 0xdf, 0x19, 0x66, 0x7d, 0x43, 0xdb, 0x0d, 0x39, 0xfa, 0xfc, 0x49,
	0xdd, 0xbc, 0xe8, 0xee, 0xff, 0xc3, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x54, 0x8f, 0x57,
	0x6d, 0x02, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Credentials) > 0 {
		for k := range m.Credentials {
			v := m.Credentials[k]
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
			dAtA[i] = 0x12
		}
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

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Credential) > 0 {
		i -= len(m.Credential)
		copy(dAtA[i:], m.Credential)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Credential)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Type)))
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

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RpOrigins) > 0 {
		for iNdEx := len(m.RpOrigins) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RpOrigins[iNdEx])
			copy(dAtA[i:], m.RpOrigins[iNdEx])
			i = encodeVarintModels(dAtA, i, uint64(len(m.RpOrigins[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RpId) > 0 {
		i -= len(m.RpId)
		copy(dAtA[i:], m.RpId)
		i = encodeVarintModels(dAtA, i, uint64(len(m.RpId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Challenge) > 0 {
		i -= len(m.Challenge)
		copy(dAtA[i:], m.Challenge)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Challenge)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Id)))
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
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if len(m.Credentials) > 0 {
		for k, v := range m.Credentials {
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

func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Credential)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Challenge)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.RpId)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if len(m.RpOrigins) > 0 {
		for _, s := range m.RpOrigins {
			l = len(s)
			n += 1 + l + sovModels(uint64(l))
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
func (m *Account) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
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
			if m.Credentials == nil {
				m.Credentials = make(map[string]*Credential)
			}
			var mapkey string
			var mapvalue *Credential
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
					mapvalue = &Credential{}
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
			m.Credentials[mapkey] = mapvalue
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
func (m *Credential) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
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
			m.Credential = append(m.Credential[:0], dAtA[iNdEx:postIndex]...)
			if m.Credential == nil {
				m.Credential = []byte{}
			}
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
func (m *Session) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenge", wireType)
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
			m.Challenge = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpId", wireType)
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
			m.RpId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpOrigins", wireType)
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
			m.RpOrigins = append(m.RpOrigins, string(dAtA[iNdEx:postIndex]))
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
