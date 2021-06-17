// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/paillier/ciphertext.proto

package paillier

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_taurusgroup_cmp_ecdsa_proto "github.com/taurusgroup/cmp-ecdsa/proto"
	io "io"
	math "math"
	math_big "math/big"
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

type Ciphertext struct {
	C *math_big.Int `protobuf:"bytes,1,opt,name=c,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"c,omitempty"`
}

func (m *Ciphertext) Reset()         { *m = Ciphertext{} }
func (m *Ciphertext) String() string { return proto.CompactTextString(m) }
func (*Ciphertext) ProtoMessage()    {}
func (*Ciphertext) Descriptor() ([]byte, []int) {
	return fileDescriptor_70de784b5a50b250, []int{0}
}
func (m *Ciphertext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ciphertext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Ciphertext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ciphertext.Merge(m, src)
}
func (m *Ciphertext) XXX_Size() int {
	return m.Size()
}
func (m *Ciphertext) XXX_DiscardUnknown() {
	xxx_messageInfo_Ciphertext.DiscardUnknown(m)
}

var xxx_messageInfo_Ciphertext proto.InternalMessageInfo

func (m *Ciphertext) GetC() *math_big.Int {
	if m != nil {
		return m.C
	}
	return nil
}

func init() {
	proto.RegisterType((*Ciphertext)(nil), "paillier.Ciphertext")
}

func init() { proto.RegisterFile("pkg/paillier/ciphertext.proto", fileDescriptor_70de784b5a50b250) }

var fileDescriptor_70de784b5a50b250 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xc8, 0x4e, 0xd7,
	0x2f, 0x48, 0xcc, 0xcc, 0xc9, 0xc9, 0x4c, 0x2d, 0xd2, 0x4f, 0xce, 0x2c, 0xc8, 0x48, 0x2d, 0x2a,
	0x49, 0xad, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x49, 0x49, 0xe9, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83,
	0x15, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa8, 0x14, 0xcb, 0xc5, 0xe5,
	0x0c, 0x37, 0x4c, 0xc8, 0x9f, 0x8b, 0x31, 0x59, 0x82, 0x51, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0x71,
	0xd5, 0x7d, 0x79, 0xdb, 0xdc, 0xc4, 0x92, 0x0c, 0xfd, 0xa4, 0xcc, 0x74, 0x3d, 0xcf, 0xbc, 0x12,
	0x6b, 0x24, 0x83, 0x4b, 0x12, 0x4b, 0x8b, 0x4a, 0x8b, 0xd3, 0x8b, 0xf2, 0x4b, 0x0b, 0xf4, 0x93,
	0x73, 0x0b, 0x74, 0x53, 0x93, 0x53, 0x8a, 0x13, 0x21, 0x36, 0x81, 0x94, 0x3a, 0x27, 0x16, 0x97,
	0xa4, 0x16, 0x05, 0x31, 0x26, 0x3b, 0xf9, 0xaf, 0x78, 0x24, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x37, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0x12, 0x36, 0x16, 0xc9, 0xe7,
	0x49, 0x6c, 0x60, 0x4b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xfe, 0xe2, 0x84, 0x10,
	0x01, 0x00, 0x00,
}

func (m *Ciphertext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ciphertext) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ciphertext) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.C)
		i -= size
		if _, err := __caster.MarshalTo(m.C, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCiphertext(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCiphertext(dAtA []byte, offset int, v uint64) int {
	offset -= sovCiphertext(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ciphertext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.C)
		n += 1 + l + sovCiphertext(uint64(l))
	}
	return n
}

func sovCiphertext(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCiphertext(x uint64) (n int) {
	return sovCiphertext(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ciphertext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCiphertext
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
			return fmt.Errorf("proto: Ciphertext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ciphertext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCiphertext
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
				return ErrInvalidLengthCiphertext
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCiphertext
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.C = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCiphertext(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCiphertext
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
func skipCiphertext(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCiphertext
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
					return 0, ErrIntOverflowCiphertext
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
					return 0, ErrIntOverflowCiphertext
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
				return 0, ErrInvalidLengthCiphertext
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCiphertext
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCiphertext
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCiphertext        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCiphertext          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCiphertext = fmt.Errorf("proto: unexpected end of group")
)