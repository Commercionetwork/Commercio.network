// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/did/tx.proto

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

type MsgSetDid struct {
	Context              []string              `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"`
	ID                   string                `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	VerificationMethod   []*VerificationMethod `protobuf:"bytes,3,rep,name=verificationMethod,proto3" json:"verificationMethod,omitempty"`
	Service              []*ServiceNew         `protobuf:"bytes,4,rep,name=service,proto3" json:"service,omitempty"`
	Authentication       []*VerificationMethod `protobuf:"bytes,5,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []*VerificationMethod `protobuf:"bytes,6,rep,name=assertionMethod,proto3" json:"assertionMethod,omitempty"`
	CapabilityDelegation []*VerificationMethod `protobuf:"bytes,7,rep,name=capabilityDelegation,proto3" json:"capabilityDelegation,omitempty"`
	CapabilityInvocation []*VerificationMethod `protobuf:"bytes,8,rep,name=capabilityInvocation,proto3" json:"capabilityInvocation,omitempty"`
	KeyAgreement         []*VerificationMethod `protobuf:"bytes,9,rep,name=keyAgreement,proto3" json:"keyAgreement,omitempty"`
}

func (m *MsgSetDid) Reset()         { *m = MsgSetDid{} }
func (m *MsgSetDid) String() string { return proto.CompactTextString(m) }
func (*MsgSetDid) ProtoMessage()    {}
func (*MsgSetDid) Descriptor() ([]byte, []int) {
	return fileDescriptor_db840dd81353defe, []int{0}
}
func (m *MsgSetDid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDid.Merge(m, src)
}
func (m *MsgSetDid) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDid proto.InternalMessageInfo

func (m *MsgSetDid) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *MsgSetDid) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MsgSetDid) GetVerificationMethod() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethod
	}
	return nil
}

func (m *MsgSetDid) GetService() []*ServiceNew {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *MsgSetDid) GetAuthentication() []*VerificationMethod {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *MsgSetDid) GetAssertionMethod() []*VerificationMethod {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *MsgSetDid) GetCapabilityDelegation() []*VerificationMethod {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *MsgSetDid) GetCapabilityInvocation() []*VerificationMethod {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

func (m *MsgSetDid) GetKeyAgreement() []*VerificationMethod {
	if m != nil {
		return m.KeyAgreement
	}
	return nil
}

type MsgSetDidResponse struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *MsgSetDidResponse) Reset()         { *m = MsgSetDidResponse{} }
func (m *MsgSetDidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetDidResponse) ProtoMessage()    {}
func (*MsgSetDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db840dd81353defe, []int{1}
}
func (m *MsgSetDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDidResponse.Merge(m, src)
}
func (m *MsgSetDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDidResponse proto.InternalMessageInfo

func (m *MsgSetDidResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSetDid)(nil), "commercionetwork.commercionetwork.did.MsgSetDid")
	proto.RegisterType((*MsgSetDidResponse)(nil), "commercionetwork.commercionetwork.did.MsgSetDidResponse")
}

func init() { proto.RegisterFile("commercionetwork/did/tx.proto", fileDescriptor_db840dd81353defe) }

var fileDescriptor_db840dd81353defe = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xbf, 0x6e, 0xe2, 0x40,
	0x10, 0x87, 0x31, 0x3e, 0xe0, 0xbc, 0x77, 0xe2, 0x74, 0xab, 0x2b, 0x2c, 0xa4, 0xb3, 0x10, 0x51,
	0x22, 0x2a, 0x93, 0x90, 0x86, 0x94, 0x89, 0xdc, 0xa0, 0x88, 0x14, 0x46, 0x4a, 0x11, 0x29, 0x85,
	0xd9, 0x9d, 0x98, 0x15, 0xd8, 0x6b, 0x79, 0x97, 0x7f, 0xca, 0x4b, 0xe4, 0xb1, 0x52, 0x52, 0xa6,
	0x8c, 0xe0, 0x0d, 0xf2, 0x04, 0x11, 0xd8, 0x26, 0x01, 0x53, 0x10, 0x39, 0xa5, 0xc7, 0x9a, 0xef,
	0x1b, 0x8f, 0x7e, 0x63, 0xf4, 0x9f, 0x70, 0xcf, 0x83, 0x90, 0x30, 0xee, 0x83, 0x9c, 0xf0, 0x70,
	0xd0, 0xa0, 0x8c, 0x36, 0xe4, 0xd4, 0x0c, 0x42, 0x2e, 0x39, 0x3e, 0xde, 0x7d, 0x6d, 0xa6, 0x0a,
	0x94, 0xd1, 0xca, 0xc9, 0x5e, 0x0a, 0x65, 0xd4, 0xe2, 0x64, 0xe4, 0x81, 0x2f, 0x23, 0x5c, 0xed,
	0xad, 0x80, 0xb4, 0x8e, 0x70, 0xbb, 0x20, 0x2d, 0x46, 0xb1, 0x8e, 0x4a, 0x84, 0xfb, 0x12, 0xa6,
	0x52, 0x57, 0xaa, 0x6a, 0x5d, 0xb3, 0x93, 0x47, 0x5c, 0x46, 0xf9, 0xb6, 0xa5, 0xe7, 0xab, 0x4a,
	0x5d, 0xb3, 0xf3, 0x6d, 0x0b, 0x33, 0x84, 0xc7, 0x10, 0xb2, 0x07, 0x46, 0x1c, 0xc9, 0xb8, 0xdf,
	0x01, 0xd9, 0xe7, 0x54, 0x57, 0xab, 0x6a, 0xfd, 0x57, 0xf3, 0xc2, 0x3c, 0x68, 0x46, 0xf3, 0x36,
	0x05, 0xb0, 0xf7, 0x40, 0xf1, 0x35, 0x2a, 0x09, 0x08, 0xc7, 0x8c, 0x80, 0xfe, 0x63, 0xcd, 0x3f,
	0x3b, 0x90, 0xdf, 0x8d, 0xba, 0x6e, 0x60, 0x62, 0x27, 0x04, 0xec, 0xa0, 0xb2, 0x33, 0x92, 0x7d,
	0xf0, 0x65, 0x2c, 0xd1, 0x0b, 0x59, 0x67, 0xde, 0x01, 0x62, 0x82, 0xfe, 0x38, 0x42, 0x40, 0xf8,
	0x69, 0x2f, 0xc5, 0xac, 0x8e, 0x5d, 0x22, 0xf6, 0xd0, 0x3f, 0xe2, 0x04, 0x4e, 0x8f, 0x0d, 0x99,
	0x9c, 0x59, 0x30, 0x04, 0x37, 0xfa, 0x9a, 0x52, 0x56, 0xd3, 0x5e, 0xec, 0xb6, 0xae, 0xed, 0x8f,
	0x79, 0xbc, 0xbc, 0x9f, 0xdf, 0xa8, 0xfb, 0xc0, 0xe2, 0x7b, 0xf4, 0x7b, 0x00, 0xb3, 0x4b, 0x37,
	0x04, 0x58, 0x65, 0x55, 0xd7, 0xb2, 0x6a, 0xb6, 0x70, 0xb5, 0x23, 0xf4, 0x77, 0x93, 0x79, 0x1b,
	0x44, 0xc0, 0x7d, 0x01, 0x71, 0xc2, 0x95, 0x24, 0xe1, 0xcd, 0x47, 0xa4, 0x76, 0x84, 0x8b, 0x25,
	0x2a, 0xc6, 0xc7, 0x71, 0x7a, 0xa0, 0x7e, 0x83, 0xae, 0xb4, 0xbe, 0xda, 0x91, 0x0c, 0x73, 0x65,
	0x3f, 0x2f, 0x0c, 0x65, 0xbe, 0x30, 0x94, 0xd7, 0x85, 0xa1, 0x3c, 0x2d, 0x8d, 0xdc, 0x7c, 0x69,
	0xe4, 0x5e, 0x96, 0x46, 0xee, 0xae, 0xe5, 0x32, 0xd9, 0x1f, 0xf5, 0x56, 0xa0, 0x46, 0xea, 0xc6,
	0x53, 0x85, 0x69, 0xf4, 0xf3, 0x98, 0x05, 0x20, 0x7a, 0xc5, 0xf5, 0xc5, 0x9f, 0xbf, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0xf0, 0xf7, 0xc2, 0x61, 0x04, 0x00, 0x00,
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
	// this line is used by starport scaffolding # proto/tx/rpc
	SetDid(ctx context.Context, in *MsgSetDid, opts ...grpc.CallOption) (*MsgSetDidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetDid(ctx context.Context, in *MsgSetDid, opts ...grpc.CallOption) (*MsgSetDidResponse, error) {
	out := new(MsgSetDidResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.did.Msg/SetDid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SetDid(context.Context, *MsgSetDid) (*MsgSetDidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetDid(ctx context.Context, req *MsgSetDid) (*MsgSetDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.did.Msg/SetDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDid(ctx, req.(*MsgSetDid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commercionetwork.commercionetwork.did.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDid",
			Handler:    _Msg_SetDid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commercionetwork/did/tx.proto",
}

func (m *MsgSetDid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyAgreement) > 0 {
		for iNdEx := len(m.KeyAgreement) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyAgreement[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CapabilityInvocation[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CapabilityDelegation[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssertionMethod[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authentication[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Service) > 0 {
		for iNdEx := len(m.Service) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Service[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.VerificationMethod) > 0 {
		for iNdEx := len(m.VerificationMethod) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethod[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgSetDid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.VerificationMethod) > 0 {
		for _, e := range m.VerificationMethod {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Service) > 0 {
		for _, e := range m.Service {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, e := range m.Authentication {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, e := range m.AssertionMethod {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, e := range m.CapabilityDelegation {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, e := range m.CapabilityInvocation {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.KeyAgreement) > 0 {
		for _, e := range m.KeyAgreement {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSetDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetDid) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetDid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
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
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethod = append(m.VerificationMethod, &VerificationMethod{})
			if err := m.VerificationMethod[len(m.VerificationMethod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = append(m.Service, &ServiceNew{})
			if err := m.Service[len(m.Service)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, &VerificationMethod{})
			if err := m.Authentication[len(m.Authentication)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, &VerificationMethod{})
			if err := m.AssertionMethod[len(m.AssertionMethod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, &VerificationMethod{})
			if err := m.CapabilityDelegation[len(m.CapabilityDelegation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, &VerificationMethod{})
			if err := m.CapabilityInvocation[len(m.CapabilityInvocation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyAgreement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyAgreement = append(m.KeyAgreement, &VerificationMethod{})
			if err := m.KeyAgreement[len(m.KeyAgreement)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSetDidResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			m.ID = string(dAtA[iNdEx:postIndex])
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
