// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lava/epochstorage/epoch_details.proto

package types

import (
	fmt "fmt"
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

type EpochDetails struct {
	StartBlock    uint64   `protobuf:"varint,1,opt,name=startBlock,proto3" json:"startBlock,omitempty"`
	EarliestStart uint64   `protobuf:"varint,2,opt,name=earliestStart,proto3" json:"earliestStart,omitempty"`
	DeletedEpochs []uint64 `protobuf:"varint,3,rep,packed,name=deletedEpochs,proto3" json:"deletedEpochs,omitempty"`
}

func (m *EpochDetails) Reset()         { *m = EpochDetails{} }
func (m *EpochDetails) String() string { return proto.CompactTextString(m) }
func (*EpochDetails) ProtoMessage()    {}
func (*EpochDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_963498f65c9d4c4e, []int{0}
}
func (m *EpochDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochDetails.Merge(m, src)
}
func (m *EpochDetails) XXX_Size() int {
	return m.Size()
}
func (m *EpochDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EpochDetails proto.InternalMessageInfo

func (m *EpochDetails) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *EpochDetails) GetEarliestStart() uint64 {
	if m != nil {
		return m.EarliestStart
	}
	return 0
}

func (m *EpochDetails) GetDeletedEpochs() []uint64 {
	if m != nil {
		return m.DeletedEpochs
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochDetails)(nil), "lavanet.lava.epochstorage.EpochDetails")
}

func init() {
	proto.RegisterFile("lava/epochstorage/epoch_details.proto", fileDescriptor_963498f65c9d4c4e)
}

var fileDescriptor_963498f65c9d4c4e = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x49, 0x2c, 0x4b,
	0xd4, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0x28, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x85, 0x70, 0xe2,
	0x53, 0x52, 0x4b, 0x12, 0x33, 0x73, 0x8a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x41,
	0xca, 0xf2, 0x52, 0x4b, 0xf4, 0x40, 0xb4, 0x1e, 0xb2, 0x72, 0xa5, 0x2a, 0x2e, 0x1e, 0x57, 0x10,
	0xdf, 0x05, 0xa2, 0x41, 0x48, 0x8e, 0x8b, 0xab, 0xb8, 0x24, 0xb1, 0xa8, 0xc4, 0x29, 0x27, 0x3f,
	0x39, 0x5b, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x08, 0x49, 0x44, 0x48, 0x85, 0x8b, 0x37, 0x35,
	0xb1, 0x28, 0x27, 0x33, 0xb5, 0xb8, 0x24, 0x18, 0x24, 0x2a, 0xc1, 0x04, 0x56, 0x82, 0x2a, 0x08,
	0x52, 0x95, 0x92, 0x9a, 0x93, 0x5a, 0x92, 0x9a, 0x02, 0x36, 0xbc, 0x58, 0x82, 0x59, 0x81, 0x19,
	0xa4, 0x0a, 0x45, 0xd0, 0xc9, 0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0x6e, 0x07, 0xd3,
	0xfa, 0x15, 0xa8, 0x9e, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xd2, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x17, 0xf0, 0x63, 0x0e, 0x01, 0x00, 0x00,
}

func (m *EpochDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeletedEpochs) > 0 {
		dAtA2 := make([]byte, len(m.DeletedEpochs)*10)
		var j1 int
		for _, num := range m.DeletedEpochs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEpochDetails(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if m.EarliestStart != 0 {
		i = encodeVarintEpochDetails(dAtA, i, uint64(m.EarliestStart))
		i--
		dAtA[i] = 0x10
	}
	if m.StartBlock != 0 {
		i = encodeVarintEpochDetails(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochDetails(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochDetails(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartBlock != 0 {
		n += 1 + sovEpochDetails(uint64(m.StartBlock))
	}
	if m.EarliestStart != 0 {
		n += 1 + sovEpochDetails(uint64(m.EarliestStart))
	}
	if len(m.DeletedEpochs) > 0 {
		l = 0
		for _, e := range m.DeletedEpochs {
			l += sovEpochDetails(uint64(e))
		}
		n += 1 + sovEpochDetails(uint64(l)) + l
	}
	return n
}

func sovEpochDetails(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochDetails(x uint64) (n int) {
	return sovEpochDetails(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochDetails
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
			return fmt.Errorf("proto: EpochDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochDetails
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarliestStart", wireType)
			}
			m.EarliestStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochDetails
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EarliestStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEpochDetails
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DeletedEpochs = append(m.DeletedEpochs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEpochDetails
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEpochDetails
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEpochDetails
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.DeletedEpochs) == 0 {
					m.DeletedEpochs = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEpochDetails
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DeletedEpochs = append(m.DeletedEpochs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedEpochs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpochDetails(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochDetails
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
func skipEpochDetails(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochDetails
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
					return 0, ErrIntOverflowEpochDetails
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
					return 0, ErrIntOverflowEpochDetails
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
				return 0, ErrInvalidLengthEpochDetails
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochDetails
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochDetails
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochDetails        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochDetails          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochDetails = fmt.Errorf("proto: unexpected end of group")
)
