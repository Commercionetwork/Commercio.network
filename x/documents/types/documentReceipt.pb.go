// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/documents/documentReceipt.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type DocumentReceipt struct {
	UUID         string     `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Sender       string     `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient    string     `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TxHash       string     `protobuf:"bytes,4,opt,name=txHash,proto3" json:"txHash,omitempty"`
	DocumentUUID string     `protobuf:"bytes,5,opt,name=documentUUID,proto3" json:"documentUUID,omitempty"`
	Proof        string     `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
	Timestamp    *time.Time `protobuf:"bytes,7,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
}

func (m *DocumentReceipt) Reset()         { *m = DocumentReceipt{} }
func (m *DocumentReceipt) String() string { return proto.CompactTextString(m) }
func (*DocumentReceipt) ProtoMessage()    {}
func (*DocumentReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b279a7eb2363a8c, []int{0}
}
func (m *DocumentReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DocumentReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentReceipt.Merge(m, src)
}
func (m *DocumentReceipt) XXX_Size() int {
	return m.Size()
}
func (m *DocumentReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentReceipt proto.InternalMessageInfo

func (m *DocumentReceipt) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *DocumentReceipt) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *DocumentReceipt) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *DocumentReceipt) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *DocumentReceipt) GetDocumentUUID() string {
	if m != nil {
		return m.DocumentUUID
	}
	return ""
}

func (m *DocumentReceipt) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func (m *DocumentReceipt) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*DocumentReceipt)(nil), "commercionetwork.documents.DocumentReceipt")
}

func init() {
	proto.RegisterFile("commercionetwork/documents/documentReceipt.proto", fileDescriptor_6b279a7eb2363a8c)
}

var fileDescriptor_6b279a7eb2363a8c = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0x7a, 0x31,
	0x14, 0xc5, 0xe9, 0xff, 0x0f, 0x18, 0xaa, 0x89, 0x49, 0x43, 0x4c, 0xf3, 0x62, 0x0a, 0x61, 0x62,
	0xea, 0x33, 0xba, 0x33, 0x10, 0x06, 0x5d, 0x89, 0x24, 0xc6, 0x0d, 0xca, 0xe5, 0xd1, 0x68, 0xdf,
	0x6d, 0xfa, 0x4a, 0xc4, 0x6f, 0xc1, 0xc7, 0x72, 0x64, 0x74, 0xd3, 0xc0, 0xec, 0x77, 0x30, 0xb4,
	0x02, 0x22, 0xdb, 0x3d, 0xe7, 0xde, 0x5f, 0x4e, 0x4f, 0x4a, 0xaf, 0x14, 0x1a, 0x03, 0x4e, 0x69,
	0xcc, 0xc1, 0xbf, 0xa0, 0x7b, 0x4a, 0xc7, 0xa8, 0x66, 0x06, 0x72, 0x5f, 0xec, 0xa6, 0x3e, 0x28,
	0xd0, 0xd6, 0x4b, 0xeb, 0xd0, 0x23, 0x4b, 0xfe, 0x12, 0x72, 0x47, 0x24, 0xf5, 0x0c, 0x33, 0x0c,
	0x67, 0xe9, 0x66, 0x8a, 0x44, 0xd2, 0xc8, 0x10, 0xb3, 0x67, 0x48, 0x83, 0x1a, 0xcd, 0x26, 0xa9,
	0xd7, 0x06, 0x0a, 0x3f, 0x34, 0x36, 0x1e, 0xb4, 0xbe, 0x08, 0x3d, 0xef, 0x1d, 0x86, 0x31, 0x46,
	0xcb, 0x83, 0xc1, 0x5d, 0x8f, 0x93, 0x26, 0x69, 0xd7, 0xfa, 0x61, 0x66, 0x17, 0xb4, 0x5a, 0x40,
	0x3e, 0x06, 0xc7, 0xff, 0x05, 0xf7, 0x47, 0xb1, 0x4b, 0x5a, 0x73, 0xa0, 0xb4, 0xd5, 0x90, 0x7b,
	0xfe, 0x3f, 0xac, 0xf6, 0xc6, 0x86, 0xf2, 0xf3, 0xdb, 0x61, 0x31, 0xe5, 0xe5, 0x48, 0x45, 0xc5,
	0x5a, 0xf4, 0x6c, 0xfb, 0xf2, 0x90, 0x54, 0x09, 0xdb, 0x03, 0x8f, 0xd5, 0x69, 0xc5, 0x3a, 0xc4,
	0x09, 0xaf, 0x86, 0x65, 0x14, 0xac, 0x43, 0x6b, 0xbb, 0x0a, 0xfc, 0xa4, 0x49, 0xda, 0xa7, 0xd7,
	0x89, 0x8c, 0x25, 0xe5, 0xb6, 0xa4, 0xbc, 0xdf, 0x5e, 0x74, 0xcb, 0x8b, 0x8f, 0x06, 0xe9, 0xef,
	0x91, 0xee, 0xc3, 0xdb, 0x4a, 0x90, 0xe5, 0x4a, 0x90, 0xcf, 0x95, 0x20, 0x8b, 0xb5, 0x28, 0x2d,
	0xd7, 0xa2, 0xf4, 0xbe, 0x16, 0xa5, 0xc7, 0x4e, 0xa6, 0xfd, 0x74, 0x36, 0x92, 0x0a, 0x4d, 0x7a,
	0xf4, 0x33, 0x47, 0xc6, 0xfc, 0xd7, 0x67, 0xf9, 0x57, 0x0b, 0xc5, 0xa8, 0x1a, 0xe2, 0x6f, 0xbe,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xec, 0xae, 0x2f, 0xd7, 0x01, 0x00, 0x00,
}

func (m *DocumentReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DocumentReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Timestamp):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DocumentUUID) > 0 {
		i -= len(m.DocumentUUID)
		copy(dAtA[i:], m.DocumentUUID)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.DocumentUUID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintDocumentReceipt(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDocumentReceipt(dAtA []byte, offset int, v uint64) int {
	offset -= sovDocumentReceipt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DocumentReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	l = len(m.DocumentUUID)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	if m.Timestamp != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovDocumentReceipt(uint64(l))
	}
	return n
}

func sovDocumentReceipt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDocumentReceipt(x uint64) (n int) {
	return sovDocumentReceipt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DocumentReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocumentReceipt
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
			return fmt.Errorf("proto: DocumentReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DocumentUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DocumentUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocumentReceipt
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
				return ErrInvalidLengthDocumentReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDocumentReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocumentReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDocumentReceipt
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
func skipDocumentReceipt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocumentReceipt
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
					return 0, ErrIntOverflowDocumentReceipt
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
					return 0, ErrIntOverflowDocumentReceipt
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
				return 0, ErrInvalidLengthDocumentReceipt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDocumentReceipt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDocumentReceipt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDocumentReceipt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocumentReceipt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDocumentReceipt = fmt.Errorf("proto: unexpected end of group")
)
