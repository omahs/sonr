// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/vault/v1/storage.proto

// Package Motor is used for defining a Motor node and its properties.

package v1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sonrhq/core/x/identity/types"
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

// RefreshSharesRequest is the request to refresh the keypair.
type RefreshSharesRequest struct {
	CredentialResponse string `protobuf:"bytes,1,opt,name=credential_response,json=credentialResponse,proto3" json:"credential_response,omitempty"`
	SessionId          string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (m *RefreshSharesRequest) Reset()         { *m = RefreshSharesRequest{} }
func (m *RefreshSharesRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshSharesRequest) ProtoMessage()    {}
func (*RefreshSharesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3799bf64fdcdea5, []int{0}
}
func (m *RefreshSharesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefreshSharesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefreshSharesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefreshSharesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshSharesRequest.Merge(m, src)
}
func (m *RefreshSharesRequest) XXX_Size() int {
	return m.Size()
}
func (m *RefreshSharesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshSharesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshSharesRequest proto.InternalMessageInfo

func (m *RefreshSharesRequest) GetCredentialResponse() string {
	if m != nil {
		return m.CredentialResponse
	}
	return ""
}

func (m *RefreshSharesRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// RefreshSharesResponse is the response to a Refresh request.
type RefreshSharesResponse struct {
	Id          []byte             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address     string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DidDocument *types.DidDocument `protobuf:"bytes,3,opt,name=did_document,json=didDocument,proto3" json:"did_document,omitempty"`
}

func (m *RefreshSharesResponse) Reset()         { *m = RefreshSharesResponse{} }
func (m *RefreshSharesResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshSharesResponse) ProtoMessage()    {}
func (*RefreshSharesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3799bf64fdcdea5, []int{1}
}
func (m *RefreshSharesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefreshSharesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefreshSharesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefreshSharesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshSharesResponse.Merge(m, src)
}
func (m *RefreshSharesResponse) XXX_Size() int {
	return m.Size()
}
func (m *RefreshSharesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshSharesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshSharesResponse proto.InternalMessageInfo

func (m *RefreshSharesResponse) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RefreshSharesResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RefreshSharesResponse) GetDidDocument() *types.DidDocument {
	if m != nil {
		return m.DidDocument
	}
	return nil
}

func init() {
	proto.RegisterType((*RefreshSharesRequest)(nil), "sonrhq.sonr.vault.v1.RefreshSharesRequest")
	proto.RegisterType((*RefreshSharesResponse)(nil), "sonrhq.sonr.vault.v1.RefreshSharesResponse")
}

func init() { proto.RegisterFile("sonr/vault/v1/storage.proto", fileDescriptor_f3799bf64fdcdea5) }

var fileDescriptor_f3799bf64fdcdea5 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x0b, 0xd3, 0x30,
	0x18, 0xc6, 0x97, 0x0a, 0xca, 0xb2, 0xe9, 0x21, 0x4e, 0x28, 0x9d, 0x96, 0x51, 0x54, 0xc6, 0x84,
	0x84, 0xcd, 0x9b, 0x17, 0x41, 0x06, 0xe2, 0xb5, 0x03, 0x0f, 0x5e, 0x46, 0xd6, 0xbc, 0x6b, 0x03,
	0x5b, 0xd2, 0x25, 0x69, 0x61, 0x57, 0x8f, 0x9e, 0x06, 0x7e, 0x02, 0xbf, 0x8d, 0xc7, 0x81, 0x17,
	0x8f, 0xb2, 0xf9, 0x41, 0xa4, 0x7f, 0xe6, 0x54, 0x76, 0xf0, 0x54, 0xd2, 0xe7, 0x79, 0xde, 0xb7,
	0xfd, 0x3d, 0xc1, 0x43, 0xab, 0x95, 0x61, 0x25, 0x2f, 0x36, 0x8e, 0x95, 0x53, 0x66, 0x9d, 0x36,
	0x3c, 0x05, 0x9a, 0x1b, 0xed, 0x34, 0x19, 0x54, 0x62, 0xb6, 0xa3, 0xd5, 0x83, 0xd6, 0x1e, 0x5a,
	0x4e, 0x83, 0x20, 0xd1, 0x06, 0x98, 0x14, 0xa0, 0x9c, 0x74, 0xfb, 0x2a, 0x25, 0xa4, 0x68, 0x12,
	0xc1, 0xe3, 0x54, 0xeb, 0x74, 0x03, 0x8c, 0xe7, 0x92, 0x71, 0xa5, 0xb4, 0xe3, 0x4e, 0x6a, 0x65,
	0x1b, 0x35, 0x5a, 0xe3, 0x41, 0x0c, 0x6b, 0x03, 0x36, 0x5b, 0x64, 0xdc, 0x80, 0x8d, 0x61, 0x57,
	0x80, 0x75, 0x84, 0xe1, 0x87, 0x89, 0x81, 0x7a, 0x20, 0xdf, 0x2c, 0x0d, 0xd8, 0x5c, 0x2b, 0x0b,
	0x3e, 0x1a, 0xa1, 0x71, 0x37, 0x26, 0x57, 0x29, 0x6e, 0x15, 0xf2, 0x04, 0x63, 0x0b, 0xd6, 0x4a,
	0xad, 0x96, 0x52, 0xf8, 0x5e, 0xed, 0xeb, 0xb6, 0x6f, 0xde, 0x89, 0xe8, 0x13, 0xc2, 0x8f, 0xfe,
	0x59, 0xd4, 0x06, 0x1f, 0x60, 0x4f, 0x8a, 0x7a, 0x70, 0x3f, 0xf6, 0xa4, 0x20, 0x3e, 0xbe, 0xc7,
	0x85, 0x30, 0x60, 0x6d, 0x3b, 0xe5, 0x72, 0x24, 0x6f, 0x71, 0x5f, 0x48, 0xb1, 0x14, 0x3a, 0x29,
	0xb6, 0xa0, 0x9c, 0x7f, 0x67, 0x84, 0xc6, 0xbd, 0xd9, 0x53, 0xda, 0x22, 0xa9, 0x18, 0xd0, 0x0b,
	0x03, 0x5a, 0x4e, 0xe9, 0x5c, 0x8a, 0x79, 0xeb, 0x8d, 0x7b, 0xe2, 0x7a, 0x98, 0x7d, 0x41, 0xb8,
	0xff, 0xbe, 0x62, 0xb7, 0x68, 0xd8, 0x92, 0x03, 0xc2, 0xf7, 0xff, 0xfa, 0x3a, 0x32, 0xa1, 0xb7,
	0x40, 0xd3, 0x5b, 0xac, 0x82, 0x17, 0xff, 0xe5, 0x6d, 0x7e, 0x37, 0x7a, 0xfe, 0xf1, 0xdb, 0xcf,
	0xcf, 0xde, 0xe8, 0x15, 0x9a, 0x44, 0x43, 0xf6, 0x47, 0xd3, 0x6d, 0xcd, 0xcc, 0x34, 0xa9, 0x37,
	0xaf, 0xbf, 0x9e, 0x42, 0x74, 0x3c, 0x85, 0xe8, 0xc7, 0x29, 0x44, 0x87, 0x73, 0xd8, 0x39, 0x9e,
	0xc3, 0xce, 0xf7, 0x73, 0xd8, 0xf9, 0xf0, 0x2c, 0x95, 0x2e, 0x2b, 0x56, 0x34, 0xd1, 0x5b, 0xd6,
	0x2c, 0x66, 0x75, 0xfd, 0x6e, 0x9f, 0x83, 0xfd, 0x7d, 0x6f, 0x56, 0x77, 0xeb, 0x82, 0x5f, 0xfe,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xdb, 0x76, 0x62, 0x4f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VaultStorageClient is the client API for VaultStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VaultStorageClient interface {
	// Refresh Shares
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	RefreshShares(ctx context.Context, in *RefreshSharesRequest, opts ...grpc.CallOption) (*RefreshSharesResponse, error)
}

type vaultStorageClient struct {
	cc grpc1.ClientConn
}

func NewVaultStorageClient(cc grpc1.ClientConn) VaultStorageClient {
	return &vaultStorageClient{cc}
}

func (c *vaultStorageClient) RefreshShares(ctx context.Context, in *RefreshSharesRequest, opts ...grpc.CallOption) (*RefreshSharesResponse, error) {
	out := new(RefreshSharesResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.sonr.vault.v1.VaultStorage/RefreshShares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VaultStorageServer is the server API for VaultStorage service.
type VaultStorageServer interface {
	// Refresh Shares
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	RefreshShares(context.Context, *RefreshSharesRequest) (*RefreshSharesResponse, error)
}

// UnimplementedVaultStorageServer can be embedded to have forward compatible implementations.
type UnimplementedVaultStorageServer struct {
}

func (*UnimplementedVaultStorageServer) RefreshShares(ctx context.Context, req *RefreshSharesRequest) (*RefreshSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshShares not implemented")
}

func RegisterVaultStorageServer(s grpc1.Server, srv VaultStorageServer) {
	s.RegisterService(&_VaultStorage_serviceDesc, srv)
}

func _VaultStorage_RefreshShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSharesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultStorageServer).RefreshShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.sonr.vault.v1.VaultStorage/RefreshShares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultStorageServer).RefreshShares(ctx, req.(*RefreshSharesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VaultStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonrhq.sonr.vault.v1.VaultStorage",
	HandlerType: (*VaultStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshShares",
			Handler:    _VaultStorage_RefreshShares_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sonr/vault/v1/storage.proto",
}

func (m *RefreshSharesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefreshSharesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefreshSharesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredentialResponse) > 0 {
		i -= len(m.CredentialResponse)
		copy(dAtA[i:], m.CredentialResponse)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.CredentialResponse)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RefreshSharesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefreshSharesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefreshSharesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DidDocument != nil {
		{
			size, err := m.DidDocument.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RefreshSharesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CredentialResponse)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *RefreshSharesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.DidDocument != nil {
		l = m.DidDocument.Size()
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RefreshSharesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: RefreshSharesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefreshSharesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialResponse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialResponse = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
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
func (m *RefreshSharesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: RefreshSharesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefreshSharesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DidDocument == nil {
				m.DidDocument = &types.DidDocument{}
			}
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
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
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
				return 0, ErrInvalidLengthStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStorage = fmt.Errorf("proto: unexpected end of group")
)
