// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/v1/nft_input.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type NFTInput struct {
	Uri      string    `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	UriHash  string    `protobuf:"bytes,2,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	Metadata JsonInput `protobuf:"bytes,3,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
}

func (m *NFTInput) Reset()         { *m = NFTInput{} }
func (m *NFTInput) String() string { return proto.CompactTextString(m) }
func (*NFTInput) ProtoMessage()    {}
func (*NFTInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdda3a1fff6751ff, []int{0}
}
func (m *NFTInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTInput.Merge(m, src)
}
func (m *NFTInput) XXX_Size() int {
	return m.Size()
}
func (m *NFTInput) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTInput.DiscardUnknown(m)
}

var xxx_messageInfo_NFTInput proto.InternalMessageInfo

func (m *NFTInput) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NFTInput) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func init() {
	proto.RegisterType((*NFTInput)(nil), "likechain.likenft.v1.NFTInput")
}

func init() {
	proto.RegisterFile("likechain/likenft/v1/nft_input.proto", fileDescriptor_cdda3a1fff6751ff)
}

var fileDescriptor_cdda3a1fff6751ff = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0xc9, 0xcc, 0x4e,
	0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb1, 0xf2, 0xd2, 0x4a, 0xf4, 0xcb, 0x0c, 0xf5, 0xf3,
	0xd2, 0x4a, 0xe2, 0x33, 0xf3, 0x0a, 0x4a, 0x4b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0xe0, 0xaa, 0xf4, 0xa0, 0xaa, 0xf4, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x0a,
	0xf4, 0x41, 0x2c, 0x88, 0x5a, 0xa5, 0x34, 0x2e, 0x0e, 0x3f, 0xb7, 0x10, 0x4f, 0x90, 0x6e, 0x21,
	0x01, 0x2e, 0xe6, 0xd2, 0xa2, 0x4c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48,
	0x92, 0x8b, 0xa3, 0xb4, 0x28, 0x33, 0x3e, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x09, 0x2c, 0xcc, 0x5e,
	0x5a, 0x94, 0xe9, 0x91, 0x58, 0x9c, 0x21, 0xa4, 0xcb, 0xc5, 0x91, 0x9b, 0x5a, 0x92, 0x98, 0x92,
	0x58, 0x92, 0x28, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0xe3, 0x24, 0x78, 0xe2, 0x9e, 0x3c, 0xc3, 0xad,
	0x7b, 0xf2, 0x9c, 0x5e, 0xc5, 0xf9, 0x79, 0x60, 0x13, 0x83, 0xe0, 0x4a, 0x9c, 0xfc, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x17, 0xec, 0xa9, 0xe4, 0x7c, 0xa8, 0xef, 0x40, 0x0c, 0x5d, 0x88, 0x67, 0xcb,
	0x8c, 0xf5, 0x2b, 0xe0, 0x3e, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xdf, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x11, 0xfa, 0x91, 0x17, 0x13, 0x01, 0x00, 0x00,
}

func (m *NFTInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Metadata.Size()
		i -= size
		if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNftInput(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintNftInput(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNftInput(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftInput(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftInput(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFTInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNftInput(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovNftInput(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovNftInput(uint64(l))
	return n
}

func sovNftInput(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftInput(x uint64) (n int) {
	return sovNftInput(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFTInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftInput
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
			return fmt.Errorf("proto: NFTInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftInput
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
				return ErrInvalidLengthNftInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftInput
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
				return ErrInvalidLengthNftInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftInput
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
				return ErrInvalidLengthNftInput
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftInput
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
func skipNftInput(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftInput
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
					return 0, ErrIntOverflowNftInput
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
					return 0, ErrIntOverflowNftInput
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
				return 0, ErrInvalidLengthNftInput
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftInput
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftInput
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftInput        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftInput          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftInput = fmt.Errorf("proto: unexpected end of group")
)
