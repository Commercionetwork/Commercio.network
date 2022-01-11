// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/commerciomint/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	ConversionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=conversion_rate,json=conversionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"conversion_rate"`
	FreezePeriod   *time.Duration                         `protobuf:"bytes,4,opt,name=freeze_period,json=freezePeriod,proto3,stdduration" json:"freeze_period,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8860fd0377586071, []int{0}
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

func (m *Params) GetFreezePeriod() *time.Duration {
	if m != nil {
		return m.FreezePeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "commercionetwork.commercionetwork.commerciomint.Params")
}

func init() {
	proto.RegisterFile("commercionetwork/commerciomint/params.proto", fileDescriptor_8860fd0377586071)
}

var fileDescriptor_8860fd0377586071 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x27, 0x7c, 0xa5, 0xf0, 0x8d, 0x7f, 0x50, 0x5c, 0xd4, 0x2e, 0xd2, 0xe2, 0x42, 0x0a,
	0x62, 0x42, 0xf5, 0x0e, 0x86, 0x5e, 0x40, 0xe9, 0x46, 0x10, 0xa1, 0x64, 0xd2, 0xd3, 0x31, 0xd4,
	0x99, 0x33, 0x24, 0x69, 0xfd, 0xb9, 0x0a, 0x97, 0xde, 0x81, 0xb7, 0xd2, 0x65, 0x97, 0xe2, 0xa2,
	0xca, 0xcc, 0x8d, 0xc8, 0x64, 0x66, 0x14, 0x0b, 0x5d, 0x25, 0x79, 0x4f, 0x9e, 0x73, 0x1e, 0x38,
	0xfe, 0xb9, 0xc4, 0x38, 0x06, 0x2d, 0x15, 0x26, 0x60, 0x1f, 0x50, 0xcf, 0xf9, 0x4f, 0x10, 0xab,
	0xc4, 0xf2, 0x54, 0x68, 0x11, 0x1b, 0x96, 0x6a, 0xb4, 0xd8, 0xe2, 0xdb, 0x9f, 0xd9, 0xee, 0xa0,
	0xa0, 0x3b, 0xc7, 0x11, 0x46, 0xe8, 0x58, 0x5e, 0xdc, 0xca, 0x36, 0x1d, 0x2a, 0xd1, 0xc4, 0x68,
	0x78, 0x28, 0x0c, 0xf0, 0xe5, 0x20, 0x04, 0x2b, 0x06, 0x5c, 0xa2, 0x4a, 0xea, 0x7a, 0x84, 0x18,
	0xdd, 0x03, 0x77, 0xaf, 0x70, 0x31, 0xe3, 0xd3, 0x85, 0x16, 0x56, 0x61, 0x55, 0x3f, 0x7d, 0x23,
	0x7e, 0x73, 0xe4, 0xbc, 0x5a, 0xd7, 0xfe, 0x91, 0xc4, 0x64, 0x09, 0xda, 0x28, 0x4c, 0x26, 0x5a,
	0x58, 0x68, 0xff, 0xeb, 0x91, 0xfe, 0xff, 0x80, 0xad, 0x36, 0x5d, 0xef, 0x63, 0xd3, 0x3d, 0x8b,
	0x94, 0xbd, 0x5b, 0x84, 0x85, 0x17, 0xaf, 0xc6, 0x96, 0xc7, 0x85, 0x99, 0xce, 0xb9, 0x7d, 0x4a,
	0xc1, 0xb0, 0x21, 0xc8, 0xf1, 0xe1, 0x6f, 0x9b, 0xb1, 0xb0, 0xd0, 0x1a, 0xfa, 0x07, 0x33, 0x0d,
	0xf0, 0x0c, 0x93, 0x14, 0xb4, 0xc2, 0x69, 0xbb, 0xd1, 0x23, 0xfd, 0xbd, 0xcb, 0x13, 0x56, 0xba,
	0xb1, 0xda, 0x8d, 0x0d, 0x2b, 0xb7, 0xa0, 0xf1, 0xfa, 0xd9, 0x25, 0xe3, 0xfd, 0x92, 0x1a, 0x39,
	0x28, 0xb8, 0x5d, 0x65, 0x94, 0xac, 0x33, 0x4a, 0xbe, 0x32, 0x4a, 0x5e, 0x72, 0xea, 0xad, 0x73,
	0xea, 0xbd, 0xe7, 0xd4, 0xbb, 0x09, 0xfe, 0x78, 0xed, 0x5a, 0x41, 0x1d, 0x3c, 0x6e, 0x6d, 0xc5,
	0x79, 0x87, 0x4d, 0x27, 0x71, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x34, 0xff, 0xd1, 0x07, 0xc4,
	0x01, 0x00, 0x00,
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
	if m.FreezePeriod != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.FreezePeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FreezePeriod):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintParams(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.ConversionRate.Size()
		i -= size
		if _, err := m.ConversionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ConversionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.FreezePeriod != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FreezePeriod)
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConversionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreezePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FreezePeriod == nil {
				m.FreezePeriod = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.FreezePeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
