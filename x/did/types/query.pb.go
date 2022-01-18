// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/did/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryResolveDidDocumentRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *QueryResolveDidDocumentRequest) Reset()         { *m = QueryResolveDidDocumentRequest{} }
func (m *QueryResolveDidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryResolveDidDocumentRequest) ProtoMessage()    {}
func (*QueryResolveDidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b34d99399b1471, []int{0}
}
func (m *QueryResolveDidDocumentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResolveDidDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResolveDidDocumentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResolveDidDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResolveDidDocumentRequest.Merge(m, src)
}
func (m *QueryResolveDidDocumentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryResolveDidDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResolveDidDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResolveDidDocumentRequest proto.InternalMessageInfo

func (m *QueryResolveDidDocumentRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type QueryResolveDidDocumentResponse struct {
	DidDocument *DidDocument `protobuf:"bytes,1,opt,name=DidDocument,proto3" json:"DidDocument,omitempty"`
}

func (m *QueryResolveDidDocumentResponse) Reset()         { *m = QueryResolveDidDocumentResponse{} }
func (m *QueryResolveDidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResolveDidDocumentResponse) ProtoMessage()    {}
func (*QueryResolveDidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b34d99399b1471, []int{1}
}
func (m *QueryResolveDidDocumentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResolveDidDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResolveDidDocumentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResolveDidDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResolveDidDocumentResponse.Merge(m, src)
}
func (m *QueryResolveDidDocumentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResolveDidDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResolveDidDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResolveDidDocumentResponse proto.InternalMessageInfo

func (m *QueryResolveDidDocumentResponse) GetDidDocument() *DidDocument {
	if m != nil {
		return m.DidDocument
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryResolveDidDocumentRequest)(nil), "commercionetwork.commercionetwork.did.QueryResolveDidDocumentRequest")
	proto.RegisterType((*QueryResolveDidDocumentResponse)(nil), "commercionetwork.commercionetwork.did.QueryResolveDidDocumentResponse")
}

func init() { proto.RegisterFile("commercionetwork/did/query.proto", fileDescriptor_82b34d99399b1471) }

var fileDescriptor_82b34d99399b1471 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x7b, 0x05, 0x45, 0xaf, 0xe0, 0x90, 0x49, 0x8a, 0x9c, 0xa5, 0x50, 0x15, 0xc1, 0x9c,
	0xad, 0x8b, 0xb3, 0x44, 0x21, 0xa3, 0xc1, 0xc9, 0xed, 0x92, 0x7b, 0xc4, 0xc3, 0xe6, 0x5e, 0x9a,
	0xbb, 0xb4, 0x16, 0x71, 0xf1, 0x13, 0x08, 0x7e, 0xa9, 0x8e, 0x45, 0x17, 0x47, 0x69, 0xfd, 0x20,
	0xd2, 0x44, 0xb0, 0xd8, 0x2a, 0x05, 0xd7, 0x3f, 0xbf, 0xf7, 0xbb, 0xff, 0xbb, 0x47, 0x1b, 0x11,
	0x26, 0x09, 0x64, 0x91, 0x42, 0x0d, 0x76, 0x80, 0xd9, 0x2d, 0x97, 0x4a, 0xf2, 0x5e, 0x0e, 0xd9,
	0xd0, 0x4d, 0x33, 0xb4, 0xe8, 0xb4, 0x7e, 0x12, 0xee, 0x42, 0x20, 0x95, 0xac, 0xef, 0xc4, 0x88,
	0x71, 0x17, 0xb8, 0x48, 0x15, 0x17, 0x5a, 0xa3, 0x15, 0x56, 0xa1, 0x36, 0xa5, 0xa4, 0x7e, 0x18,
	0xa1, 0x49, 0xd0, 0xf0, 0x50, 0x18, 0x28, 0xed, 0xbc, 0xdf, 0x0e, 0xc1, 0x8a, 0x36, 0x4f, 0x45,
	0xac, 0x74, 0x01, 0x7f, 0xb1, 0x7b, 0x4b, 0x2b, 0x49, 0x25, 0x3d, 0x8c, 0xf2, 0x04, 0xb4, 0x2d,
	0xb9, 0xe6, 0x31, 0x65, 0x97, 0x33, 0x53, 0x00, 0x06, 0xbb, 0x7d, 0xf0, 0xbe, 0x81, 0x00, 0x7a,
	0x39, 0x18, 0xeb, 0x6c, 0xd1, 0xaa, 0xef, 0x6d, 0x93, 0x06, 0x39, 0xd8, 0x0c, 0xaa, 0xbe, 0xd7,
	0x1c, 0xd0, 0xdd, 0x5f, 0x27, 0x4c, 0x8a, 0xda, 0x80, 0x73, 0x45, 0x6b, 0x73, 0x71, 0x31, 0x5b,
	0xeb, 0x74, 0xdc, 0x95, 0xfe, 0xc0, 0x9d, 0x17, 0xce, 0x6b, 0x3a, 0x2f, 0x84, 0xae, 0x15, 0x2f,
	0x3b, 0x23, 0x42, 0x37, 0x7c, 0x09, 0xda, 0x2a, 0x3b, 0x74, 0xce, 0x57, 0xf4, 0xfe, 0xbd, 0x66,
	0xfd, 0xe2, 0xbf, 0x9a, 0x72, 0xf7, 0xe6, 0xd1, 0xe3, 0xeb, 0xc7, 0x73, 0x75, 0xdf, 0x69, 0xf1,
	0xa5, 0x17, 0x50, 0x65, 0x6d, 0x05, 0x86, 0xdf, 0xfb, 0xde, 0xc3, 0x59, 0x30, 0x9a, 0x30, 0x32,
	0x9e, 0x30, 0xf2, 0x3e, 0x61, 0xe4, 0x69, 0xca, 0x2a, 0xe3, 0x29, 0xab, 0xbc, 0x4d, 0x59, 0xe5,
	0xfa, 0x34, 0x56, 0xf6, 0x26, 0x0f, 0x67, 0x2d, 0x16, 0x55, 0x0b, 0xc1, 0x5d, 0x61, 0xb7, 0xc3,
	0x14, 0x4c, 0xb8, 0x5e, 0x9c, 0xf6, 0xe4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x48, 0xc8, 0xdd,
	0x97, 0x02, 0x00, 0x00,
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
	// Queries a DidDocument by id.
	Identity(ctx context.Context, in *QueryResolveDidDocumentRequest, opts ...grpc.CallOption) (*QueryResolveDidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Identity(ctx context.Context, in *QueryResolveDidDocumentRequest, opts ...grpc.CallOption) (*QueryResolveDidDocumentResponse, error) {
	out := new(QueryResolveDidDocumentResponse)
	err := c.cc.Invoke(ctx, "/commercionetwork.commercionetwork.did.Query/Identity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a DidDocument by id.
	Identity(context.Context, *QueryResolveDidDocumentRequest) (*QueryResolveDidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Identity(ctx context.Context, req *QueryResolveDidDocumentRequest) (*QueryResolveDidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identity not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Identity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResolveDidDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Identity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commercionetwork.commercionetwork.did.Query/Identity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Identity(ctx, req.(*QueryResolveDidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commercionetwork.commercionetwork.did.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Identity",
			Handler:    _Query_Identity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commercionetwork/did/query.proto",
}

func (m *QueryResolveDidDocumentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResolveDidDocumentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResolveDidDocumentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryResolveDidDocumentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResolveDidDocumentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResolveDidDocumentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
func (m *QueryResolveDidDocumentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryResolveDidDocumentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DidDocument != nil {
		l = m.DidDocument.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryResolveDidDocumentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryResolveDidDocumentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResolveDidDocumentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
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
func (m *QueryResolveDidDocumentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryResolveDidDocumentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResolveDidDocumentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
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
			if m.DidDocument == nil {
				m.DidDocument = &DidDocument{}
			}
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
