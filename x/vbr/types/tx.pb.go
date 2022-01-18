// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/vbr/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgIncrementBlockRewardsPool struct {
	Funder string                                   `protobuf:"bytes,1,opt,name=funder,proto3" json:"funder,omitempty" yaml:"funder"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgIncrementBlockRewardsPool) Reset()         { *m = MsgIncrementBlockRewardsPool{} }
func (m *MsgIncrementBlockRewardsPool) String() string { return proto.CompactTextString(m) }
func (*MsgIncrementBlockRewardsPool) ProtoMessage()    {}
func (*MsgIncrementBlockRewardsPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec5dd532f6ad8352, []int{0}
}
func (m *MsgIncrementBlockRewardsPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncrementBlockRewardsPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncrementBlockRewardsPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncrementBlockRewardsPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncrementBlockRewardsPool.Merge(m, src)
}
func (m *MsgIncrementBlockRewardsPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncrementBlockRewardsPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncrementBlockRewardsPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncrementBlockRewardsPool proto.InternalMessageInfo

func (m *MsgIncrementBlockRewardsPool) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

func (m *MsgIncrementBlockRewardsPool) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgIncrementBlockRewardsPoolResponse struct {
}

func (m *MsgIncrementBlockRewardsPoolResponse) Reset()         { *m = MsgIncrementBlockRewardsPoolResponse{} }
func (m *MsgIncrementBlockRewardsPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIncrementBlockRewardsPoolResponse) ProtoMessage()    {}
func (*MsgIncrementBlockRewardsPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec5dd532f6ad8352, []int{1}
}
func (m *MsgIncrementBlockRewardsPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncrementBlockRewardsPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncrementBlockRewardsPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncrementBlockRewardsPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncrementBlockRewardsPoolResponse.Merge(m, src)
}
func (m *MsgIncrementBlockRewardsPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncrementBlockRewardsPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncrementBlockRewardsPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncrementBlockRewardsPoolResponse proto.InternalMessageInfo

type MsgSetParams struct {
	Government           string                                 `protobuf:"bytes,1,opt,name=Government,proto3" json:"Government,omitempty" yaml:"government"`
	DistrEpochIdentifier string                                 `protobuf:"bytes,2,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
	EarnRate             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=earn_rate,json=earnRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"earn_rate" yaml:"earn_rate"`
}

func (m *MsgSetParams) Reset()         { *m = MsgSetParams{} }
func (m *MsgSetParams) String() string { return proto.CompactTextString(m) }
func (*MsgSetParams) ProtoMessage()    {}
func (*MsgSetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec5dd532f6ad8352, []int{2}
}
func (m *MsgSetParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParams.Merge(m, src)
}
func (m *MsgSetParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParams proto.InternalMessageInfo

func (m *MsgSetParams) GetGovernment() string {
	if m != nil {
		return m.Government
	}
	return ""
}

func (m *MsgSetParams) GetDistrEpochIdentifier() string {
	if m != nil {
		return m.DistrEpochIdentifier
	}
	return ""
}

type MsgSetParamsResponse struct {
}

func (m *MsgSetParamsResponse) Reset()         { *m = MsgSetParamsResponse{} }
func (m *MsgSetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetParamsResponse) ProtoMessage()    {}
func (*MsgSetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec5dd532f6ad8352, []int{3}
}
func (m *MsgSetParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParamsResponse.Merge(m, src)
}
func (m *MsgSetParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIncrementBlockRewardsPool)(nil), "commercionetwork.commercionetwork.vbr.MsgIncrementBlockRewardsPool")
	proto.RegisterType((*MsgIncrementBlockRewardsPoolResponse)(nil), "commercionetwork.commercionetwork.vbr.MsgIncrementBlockRewardsPoolResponse")
	proto.RegisterType((*MsgSetParams)(nil), "commercionetwork.commercionetwork.vbr.MsgSetParams")
	proto.RegisterType((*MsgSetParamsResponse)(nil), "commercionetwork.commercionetwork.vbr.MsgSetParamsResponse")
}

func init() { proto.RegisterFile("commercionetwork/vbr/tx.proto", fileDescriptor_ec5dd532f6ad8352) }

var fileDescriptor_ec5dd532f6ad8352 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x6a, 0x13, 0x5f,
	0x14, 0xce, 0x24, 0x10, 0x7e, 0xb9, 0x3f, 0x05, 0x3b, 0xc4, 0x92, 0x06, 0x3b, 0x53, 0x07, 0x2d,
	0x71, 0xe1, 0x0c, 0x6d, 0x11, 0x44, 0x57, 0x4e, 0x15, 0x29, 0x12, 0x28, 0xe3, 0x42, 0x70, 0x13,
	0xee, 0xcc, 0x9c, 0x4e, 0x87, 0x64, 0xee, 0x09, 0xf7, 0xde, 0xa4, 0xed, 0xc2, 0x77, 0xf0, 0x39,
	0x04, 0x77, 0xee, 0xdd, 0x76, 0xd9, 0xa5, 0xb8, 0x18, 0x25, 0x79, 0x83, 0xbc, 0x80, 0x72, 0xe7,
	0x4e, 0xd2, 0xda, 0x6a, 0x09, 0xe2, 0xea, 0x1e, 0xce, 0x39, 0xdf, 0xf7, 0x9d, 0x3f, 0xf7, 0x90,
	0xf5, 0x08, 0xb3, 0x0c, 0x78, 0x94, 0x22, 0x03, 0x79, 0x84, 0xbc, 0xef, 0x8d, 0x43, 0xee, 0xc9,
	0x63, 0x77, 0xc8, 0x51, 0xa2, 0x79, 0xff, 0x72, 0xd8, 0xbd, 0xe2, 0x18, 0x87, 0xbc, 0xdd, 0x4c,
	0x30, 0xc1, 0x02, 0xe1, 0x29, 0x4b, 0x83, 0xdb, 0x56, 0x84, 0x22, 0x43, 0xe1, 0x85, 0x54, 0x80,
	0x37, 0xde, 0x0a, 0x41, 0xd2, 0x2d, 0x2f, 0xc2, 0x94, 0xe9, 0xb8, 0xf3, 0xd9, 0x20, 0x77, 0xba,
	0x22, 0xd9, 0x63, 0x11, 0x87, 0x0c, 0x98, 0xf4, 0x07, 0x18, 0xf5, 0x03, 0x38, 0xa2, 0x3c, 0x16,
	0xfb, 0x88, 0x03, 0xf3, 0x01, 0xa9, 0x1f, 0x8c, 0x58, 0x0c, 0xbc, 0x65, 0x6c, 0x18, 0x9d, 0x86,
	0xbf, 0x32, 0xcb, 0xed, 0x9b, 0x27, 0x34, 0x1b, 0x3c, 0x71, 0xb4, 0xdf, 0x09, 0xca, 0x04, 0x53,
	0x92, 0x3a, 0xcd, 0x70, 0xc4, 0x64, 0xab, 0xba, 0x51, 0xeb, 0xfc, 0xbf, 0xbd, 0xe6, 0x6a, 0x71,
	0x57, 0x89, 0xbb, 0xa5, 0xb8, 0xbb, 0x8b, 0x29, 0xf3, 0x9f, 0x9d, 0xe6, 0x76, 0xe5, 0x9c, 0x49,
	0xc3, 0x9c, 0x0f, 0xdf, 0xec, 0x4e, 0x92, 0xca, 0xc3, 0x51, 0xa8, 0x5a, 0xf4, 0xca, 0xd2, 0xf5,
	0xf3, 0x50, 0xc4, 0x7d, 0x4f, 0x9e, 0x0c, 0x41, 0x14, 0x0c, 0x22, 0x28, 0xb5, 0x9c, 0x4d, 0x72,
	0xef, 0xba, 0x06, 0x02, 0x10, 0x43, 0x64, 0x02, 0x9c, 0x1f, 0x06, 0xb9, 0xd1, 0x15, 0xc9, 0x6b,
	0x90, 0xfb, 0x94, 0xd3, 0x4c, 0x98, 0x8f, 0x08, 0x79, 0x89, 0x63, 0xe0, 0x4c, 0xc1, 0xca, 0xee,
	0x6e, 0xcf, 0x72, 0x7b, 0x45, 0xd7, 0x94, 0x2c, 0x62, 0x4e, 0x70, 0x21, 0xd1, 0x7c, 0x43, 0x56,
	0xe3, 0x54, 0x48, 0xde, 0x83, 0x21, 0x46, 0x87, 0xbd, 0x34, 0x06, 0x26, 0xd3, 0x83, 0x14, 0x78,
	0xab, 0x5a, 0x50, 0xdc, 0x9d, 0xe5, 0xf6, 0xba, 0xa6, 0xf8, 0x7d, 0x9e, 0x13, 0x34, 0x8b, 0xc0,
	0x0b, 0xe5, 0xdf, 0x5b, 0xb8, 0xcd, 0x1e, 0x69, 0x00, 0xe5, 0xac, 0xc7, 0xa9, 0x84, 0x56, 0xad,
	0xe0, 0xf2, 0xd5, 0x98, 0xbe, 0xe6, 0xf6, 0xe6, 0x12, 0x53, 0x79, 0x0e, 0xd1, 0x2c, 0xb7, 0x6f,
	0x69, 0xe5, 0x05, 0x91, 0x13, 0xfc, 0xa7, 0xec, 0x40, 0x99, 0xab, 0xa4, 0x79, 0x71, 0x00, 0xf3,
	0xc9, 0x6c, 0x7f, 0xaa, 0x92, 0x5a, 0x57, 0x24, 0xe6, 0x47, 0x83, 0xac, 0xfd, 0xf9, 0x23, 0xec,
	0xba, 0x4b, 0xfd, 0x43, 0xf7, 0xba, 0x65, 0xb4, 0x5f, 0xfd, 0x03, 0x92, 0x79, 0xdd, 0xe6, 0x3b,
	0xd2, 0x38, 0xdf, 0xe6, 0xce, 0xf2, 0xcc, 0x0b, 0x50, 0xfb, 0xe9, 0x5f, 0x80, 0xe6, 0xf2, 0x7e,
	0x70, 0x3a, 0xb1, 0x8c, 0xb3, 0x89, 0x65, 0x7c, 0x9f, 0x58, 0xc6, 0xfb, 0xa9, 0x55, 0x39, 0x9b,
	0x5a, 0x95, 0x2f, 0x53, 0xab, 0xf2, 0xf6, 0xf1, 0x2f, 0xeb, 0xba, 0x74, 0xdb, 0x57, 0x1c, 0xc7,
	0xfa, 0xdc, 0xd5, 0x12, 0xc3, 0x7a, 0x71, 0x95, 0x3b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x57,
	0xff, 0x67, 0x6a, 0x13, 0x04, 0x00, 0x00,
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
	IncrementBlockRewardsPool(ctx context.Context, in *MsgIncrementBlockRewardsPool, opts ...grpc.CallOption) (*MsgIncrementBlockRewardsPoolResponse, error)
	SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IncrementBlockRewardsPool(ctx context.Context, in *MsgIncrementBlockRewardsPool, opts ...grpc.CallOption) (*MsgIncrementBlockRewardsPoolResponse, error) {
	out := new(MsgIncrementBlockRewardsPoolResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.vbr.Msg/IncrementBlockRewardsPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error) {
	out := new(MsgSetParamsResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.vbr.Msg/SetParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	IncrementBlockRewardsPool(context.Context, *MsgIncrementBlockRewardsPool) (*MsgIncrementBlockRewardsPoolResponse, error)
	SetParams(context.Context, *MsgSetParams) (*MsgSetParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IncrementBlockRewardsPool(ctx context.Context, req *MsgIncrementBlockRewardsPool) (*MsgIncrementBlockRewardsPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementBlockRewardsPool not implemented")
}
func (*UnimplementedMsgServer) SetParams(ctx context.Context, req *MsgSetParams) (*MsgSetParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IncrementBlockRewardsPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIncrementBlockRewardsPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IncrementBlockRewardsPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.vbr.Msg/IncrementBlockRewardsPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IncrementBlockRewardsPool(ctx, req.(*MsgIncrementBlockRewardsPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.vbr.Msg/SetParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetParams(ctx, req.(*MsgSetParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commercionetwork.commercionetwork.vbr.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementBlockRewardsPool",
			Handler:    _Msg_IncrementBlockRewardsPool_Handler,
		},
		{
			MethodName: "SetParams",
			Handler:    _Msg_SetParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commercionetwork/vbr/tx.proto",
}

func (m *MsgIncrementBlockRewardsPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncrementBlockRewardsPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncrementBlockRewardsPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIncrementBlockRewardsPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncrementBlockRewardsPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncrementBlockRewardsPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EarnRate.Size()
		i -= size
		if _, err := m.EarnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Government) > 0 {
		i -= len(m.Government)
		copy(dAtA[i:], m.Government)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Government)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgIncrementBlockRewardsPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgIncrementBlockRewardsPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Government)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.EarnRate.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetParamsResponse) Size() (n int) {
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
func (m *MsgIncrementBlockRewardsPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIncrementBlockRewardsPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncrementBlockRewardsPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
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
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgIncrementBlockRewardsPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIncrementBlockRewardsPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncrementBlockRewardsPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgSetParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Government", wireType)
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
			m.Government = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
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
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarnRate", wireType)
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
			if err := m.EarnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgSetParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
