// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/vbr/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// this line is used by starport scaffolding # 3
type QueryGetBlockRewardsPoolFundsRequest struct {
}

func (m *QueryGetBlockRewardsPoolFundsRequest) Reset()         { *m = QueryGetBlockRewardsPoolFundsRequest{} }
func (m *QueryGetBlockRewardsPoolFundsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetBlockRewardsPoolFundsRequest) ProtoMessage()    {}
func (*QueryGetBlockRewardsPoolFundsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3130ad40f57f01ad, []int{0}
}
func (m *QueryGetBlockRewardsPoolFundsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetBlockRewardsPoolFundsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetBlockRewardsPoolFundsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetBlockRewardsPoolFundsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetBlockRewardsPoolFundsRequest.Merge(m, src)
}
func (m *QueryGetBlockRewardsPoolFundsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetBlockRewardsPoolFundsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetBlockRewardsPoolFundsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetBlockRewardsPoolFundsRequest proto.InternalMessageInfo

type QueryGetBlockRewardsPoolFundsResponse struct {
	Funds github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=funds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"funds" yaml:"funds"`
}

func (m *QueryGetBlockRewardsPoolFundsResponse) Reset()         { *m = QueryGetBlockRewardsPoolFundsResponse{} }
func (m *QueryGetBlockRewardsPoolFundsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetBlockRewardsPoolFundsResponse) ProtoMessage()    {}
func (*QueryGetBlockRewardsPoolFundsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3130ad40f57f01ad, []int{1}
}
func (m *QueryGetBlockRewardsPoolFundsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetBlockRewardsPoolFundsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetBlockRewardsPoolFundsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetBlockRewardsPoolFundsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetBlockRewardsPoolFundsResponse.Merge(m, src)
}
func (m *QueryGetBlockRewardsPoolFundsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetBlockRewardsPoolFundsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetBlockRewardsPoolFundsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetBlockRewardsPoolFundsResponse proto.InternalMessageInfo

func (m *QueryGetBlockRewardsPoolFundsResponse) GetFunds() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Funds
	}
	return nil
}

type QueryGetParamsRequest struct {
}

func (m *QueryGetParamsRequest) Reset()         { *m = QueryGetParamsRequest{} }
func (m *QueryGetParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetParamsRequest) ProtoMessage()    {}
func (*QueryGetParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3130ad40f57f01ad, []int{2}
}
func (m *QueryGetParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetParamsRequest.Merge(m, src)
}
func (m *QueryGetParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetParamsRequest proto.InternalMessageInfo

type QueryGetParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *QueryGetParamsResponse) Reset()         { *m = QueryGetParamsResponse{} }
func (m *QueryGetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetParamsResponse) ProtoMessage()    {}
func (*QueryGetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3130ad40f57f01ad, []int{3}
}
func (m *QueryGetParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetParamsResponse.Merge(m, src)
}
func (m *QueryGetParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetParamsResponse proto.InternalMessageInfo

func (m *QueryGetParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryGetBlockRewardsPoolFundsRequest)(nil), "commercionetwork.commercionetwork.vbr.QueryGetBlockRewardsPoolFundsRequest")
	proto.RegisterType((*QueryGetBlockRewardsPoolFundsResponse)(nil), "commercionetwork.commercionetwork.vbr.QueryGetBlockRewardsPoolFundsResponse")
	proto.RegisterType((*QueryGetParamsRequest)(nil), "commercionetwork.commercionetwork.vbr.QueryGetParamsRequest")
	proto.RegisterType((*QueryGetParamsResponse)(nil), "commercionetwork.commercionetwork.vbr.QueryGetParamsResponse")
}

func init() { proto.RegisterFile("commercionetwork/vbr/query.proto", fileDescriptor_3130ad40f57f01ad) }

var fileDescriptor_3130ad40f57f01ad = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6e, 0x13, 0x31,
	0x18, 0xc7, 0x63, 0xaa, 0x56, 0xc2, 0x85, 0xe5, 0x44, 0x21, 0x0a, 0xc1, 0x09, 0x47, 0x8b, 0x2a,
	0xa1, 0xda, 0x6a, 0x58, 0x10, 0x82, 0x25, 0x45, 0x30, 0xc0, 0x50, 0x6e, 0x44, 0x2c, 0xbe, 0x8b,
	0x39, 0x4e, 0xc9, 0xf9, 0xbb, 0xda, 0xce, 0x85, 0xac, 0x3c, 0x01, 0x12, 0x62, 0x67, 0x66, 0xe0,
	0x39, 0x3a, 0x56, 0x62, 0x41, 0x0c, 0x05, 0x25, 0x3c, 0x01, 0x4f, 0x80, 0xce, 0x76, 0x90, 0x9a,
	0x06, 0x08, 0xa8, 0xd3, 0x59, 0x9f, 0xed, 0xff, 0xf7, 0xf3, 0xff, 0xfb, 0x1f, 0x6e, 0x27, 0x90,
	0xe7, 0x42, 0x25, 0x19, 0x48, 0x61, 0x46, 0xa0, 0xfa, 0xac, 0x8c, 0x15, 0x3b, 0x18, 0x0a, 0x35,
	0xa6, 0x85, 0x02, 0x03, 0xc1, 0xd6, 0xfc, 0x09, 0x7a, 0xaa, 0x50, 0xc6, 0xaa, 0xd1, 0x4c, 0x01,
	0xd2, 0x81, 0x60, 0xbc, 0xc8, 0x18, 0x97, 0x12, 0x0c, 0x37, 0x19, 0x48, 0xed, 0x44, 0x1a, 0x97,
	0x52, 0x48, 0xc1, 0x2e, 0x59, 0xb5, 0xf2, 0x55, 0x92, 0x80, 0xce, 0x41, 0xb3, 0x98, 0x6b, 0xc1,
	0xca, 0xdd, 0x58, 0x18, 0xbe, 0xcb, 0x12, 0xc8, 0xa4, 0xdf, 0xbf, 0xbe, 0x10, 0xae, 0xe0, 0x8a,
	0xe7, 0x5e, 0x38, 0xbc, 0x89, 0x37, 0x9f, 0x56, 0xb0, 0x8f, 0x84, 0xe9, 0x0e, 0x20, 0xe9, 0x47,
	0x62, 0xc4, 0x55, 0x4f, 0xef, 0x03, 0x0c, 0x1e, 0x0e, 0x65, 0x4f, 0x47, 0xe2, 0x60, 0x28, 0xb4,
	0x09, 0xdf, 0x23, 0xbc, 0xf5, 0x97, 0x83, 0xba, 0x00, 0xa9, 0x45, 0x30, 0xc2, 0xab, 0x2f, 0xaa,
	0x42, 0x1d, 0xb5, 0x57, 0xb6, 0xd7, 0x3b, 0x4d, 0xea, 0x20, 0x69, 0x05, 0x49, 0x3d, 0x24, 0x7d,
	0x20, 0x92, 0x3d, 0xc8, 0x64, 0x77, 0xef, 0xf0, 0xb8, 0x55, 0xfb, 0x71, 0xdc, 0xba, 0x30, 0xe6,
	0xf9, 0xe0, 0x6e, 0x68, 0x2f, 0x86, 0x1f, 0xbe, 0xb6, 0x6e, 0xa5, 0x99, 0x79, 0x39, 0x8c, 0x2b,
	0xab, 0x98, 0x7f, 0xa4, 0xfb, 0xec, 0xe8, 0x5e, 0x9f, 0x99, 0x71, 0x21, 0xf4, 0x4c, 0x43, 0x47,
	0xae, 0x5f, 0x78, 0x05, 0x6f, 0xcc, 0x08, 0xf7, 0xed, 0x13, 0x67, 0xec, 0x25, 0xbe, 0x3c, 0xbf,
	0xe1, 0x59, 0x9f, 0xe3, 0x35, 0xe7, 0x46, 0x1d, 0xb5, 0xd1, 0xf6, 0x7a, 0x67, 0x87, 0x2e, 0x35,
	0x2c, 0xea, 0x64, 0xba, 0x1b, 0x9e, 0xfe, 0xa2, 0xa3, 0x77, 0x52, 0x61, 0xe4, 0x35, 0x3b, 0xef,
	0x56, 0xf0, 0xaa, 0x6d, 0x1c, 0x7c, 0x41, 0xb8, 0xfe, 0x3b, 0xe3, 0x82, 0xc7, 0x4b, 0x36, 0x5d,
	0x66, 0x4e, 0x8d, 0x27, 0x67, 0x23, 0xe6, 0xfc, 0x09, 0x6f, 0xbc, 0xfe, 0xf4, 0xfd, 0xed, 0xb9,
	0x6b, 0xc1, 0x55, 0xb6, 0x30, 0x49, 0xd6, 0xf7, 0xe0, 0x23, 0xc2, 0xe7, 0x7f, 0x59, 0x1b, 0xdc,
	0xfb, 0x47, 0x80, 0x13, 0xa3, 0x6a, 0xdc, 0xff, 0xcf, 0xdb, 0x9e, 0x77, 0xd3, 0xf2, 0x92, 0xa0,
	0xc9, 0xfe, 0x90, 0xfc, 0x6e, 0x74, 0x38, 0x21, 0xe8, 0x68, 0x42, 0xd0, 0xb7, 0x09, 0x41, 0x6f,
	0xa6, 0xa4, 0x76, 0x34, 0x25, 0xb5, 0xcf, 0x53, 0x52, 0x7b, 0x76, 0xe7, 0x44, 0xec, 0xe6, 0x14,
	0x4e, 0x15, 0x5e, 0x59, 0x51, 0x1b, 0xc6, 0x78, 0xcd, 0xfe, 0x4e, 0xb7, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x96, 0x9d, 0xfe, 0x52, 0x10, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// this line is used by starport scaffolding # 2
	GetBlockRewardsPoolFunds(ctx context.Context, in *QueryGetBlockRewardsPoolFundsRequest, opts ...grpc.CallOption) (*QueryGetBlockRewardsPoolFundsResponse, error)
	GetParams(ctx context.Context, in *QueryGetParamsRequest, opts ...grpc.CallOption) (*QueryGetParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetBlockRewardsPoolFunds(ctx context.Context, in *QueryGetBlockRewardsPoolFundsRequest, opts ...grpc.CallOption) (*QueryGetBlockRewardsPoolFundsResponse, error) {
	out := new(QueryGetBlockRewardsPoolFundsResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.vbr.Query/GetBlockRewardsPoolFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetParams(ctx context.Context, in *QueryGetParamsRequest, opts ...grpc.CallOption) (*QueryGetParamsResponse, error) {
	out := new(QueryGetParamsResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.vbr.Query/GetParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	GetBlockRewardsPoolFunds(context.Context, *QueryGetBlockRewardsPoolFundsRequest) (*QueryGetBlockRewardsPoolFundsResponse, error)
	GetParams(context.Context, *QueryGetParamsRequest) (*QueryGetParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetBlockRewardsPoolFunds(ctx context.Context, req *QueryGetBlockRewardsPoolFundsRequest) (*QueryGetBlockRewardsPoolFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockRewardsPoolFunds not implemented")
}
func (*UnimplementedQueryServer) GetParams(ctx context.Context, req *QueryGetParamsRequest) (*QueryGetParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParams not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetBlockRewardsPoolFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetBlockRewardsPoolFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetBlockRewardsPoolFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.vbr.Query/GetBlockRewardsPoolFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetBlockRewardsPoolFunds(ctx, req.(*QueryGetBlockRewardsPoolFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.vbr.Query/GetParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetParams(ctx, req.(*QueryGetParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commercionetwork.commercionetwork.vbr.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockRewardsPoolFunds",
			Handler:    _Query_GetBlockRewardsPoolFunds_Handler,
		},
		{
			MethodName: "GetParams",
			Handler:    _Query_GetParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commercionetwork/vbr/query.proto",
}

func (m *QueryGetBlockRewardsPoolFundsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetBlockRewardsPoolFundsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetBlockRewardsPoolFundsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetBlockRewardsPoolFundsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetBlockRewardsPoolFundsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetBlockRewardsPoolFundsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetBlockRewardsPoolFundsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetBlockRewardsPoolFundsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGetParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetBlockRewardsPoolFundsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetBlockRewardsPoolFundsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetBlockRewardsPoolFundsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetBlockRewardsPoolFundsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetBlockRewardsPoolFundsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetBlockRewardsPoolFundsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.DecCoin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
