// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/v1/offer.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Offer struct {
	ClassId    string    `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	NftId      string    `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Buyer      string    `protobuf:"bytes,3,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Price      uint64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Expiration time.Time `protobuf:"bytes,5,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0d5cfe801d650e, []int{0}
}
func (m *Offer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return m.Size()
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *Offer) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *Offer) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *Offer) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Offer) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

type OfferStoreRecord struct {
	ClassId    string                                        `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	NftId      string                                        `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Buyer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=buyer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"buyer,omitempty"`
	Price      uint64                                        `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Expiration time.Time                                     `protobuf:"bytes,5,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *OfferStoreRecord) Reset()         { *m = OfferStoreRecord{} }
func (m *OfferStoreRecord) String() string { return proto.CompactTextString(m) }
func (*OfferStoreRecord) ProtoMessage()    {}
func (*OfferStoreRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0d5cfe801d650e, []int{1}
}
func (m *OfferStoreRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OfferStoreRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OfferStoreRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OfferStoreRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfferStoreRecord.Merge(m, src)
}
func (m *OfferStoreRecord) XXX_Size() int {
	return m.Size()
}
func (m *OfferStoreRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_OfferStoreRecord.DiscardUnknown(m)
}

var xxx_messageInfo_OfferStoreRecord proto.InternalMessageInfo

func (m *OfferStoreRecord) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *OfferStoreRecord) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *OfferStoreRecord) GetBuyer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Buyer
	}
	return nil
}

func (m *OfferStoreRecord) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OfferStoreRecord) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Offer)(nil), "likechain.likenft.v1.Offer")
	proto.RegisterType((*OfferStoreRecord)(nil), "likechain.likenft.v1.OfferStoreRecord")
}

func init() { proto.RegisterFile("likechain/likenft/v1/offer.proto", fileDescriptor_ad0d5cfe801d650e) }

var fileDescriptor_ad0d5cfe801d650e = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x31, 0x6e, 0xc2, 0x30,
	0x18, 0x85, 0xe3, 0x96, 0x50, 0xea, 0x76, 0xa8, 0x22, 0x2a, 0xa5, 0x0c, 0x49, 0xc4, 0xc4, 0x82,
	0x2d, 0x8a, 0x7a, 0x00, 0x50, 0xa5, 0x8a, 0x09, 0x29, 0xed, 0xd4, 0xa5, 0x4a, 0x1c, 0x27, 0x58,
	0x90, 0x38, 0xb2, 0x0d, 0x82, 0x5b, 0x70, 0x8f, 0x5e, 0x84, 0x91, 0xb1, 0x13, 0xad, 0x60, 0xea,
	0x15, 0x3a, 0x55, 0x71, 0x00, 0x31, 0x77, 0xe8, 0x94, 0xf7, 0xfe, 0xff, 0xc5, 0xfe, 0x2c, 0x3d,
	0xe8, 0x4d, 0xd8, 0x98, 0x92, 0x51, 0xc0, 0x32, 0x5c, 0xa8, 0x2c, 0x56, 0x78, 0xd6, 0xc1, 0x3c,
	0x8e, 0xa9, 0x40, 0xb9, 0xe0, 0x8a, 0x5b, 0xf5, 0x63, 0x02, 0xed, 0x13, 0x68, 0xd6, 0x69, 0xd4,
	0x13, 0x9e, 0x70, 0x1d, 0xc0, 0x85, 0x2a, 0xb3, 0x0d, 0x37, 0xe1, 0x3c, 0x99, 0x50, 0xac, 0x5d,
	0x38, 0x8d, 0xb1, 0x62, 0x29, 0x95, 0x2a, 0x48, 0xf3, 0x32, 0xd0, 0x7c, 0x07, 0xd0, 0x1c, 0x16,
	0x87, 0x5b, 0x77, 0xb0, 0x46, 0x26, 0x81, 0x94, 0x6f, 0x2c, 0xb2, 0x81, 0x07, 0x5a, 0x97, 0xfe,
	0x85, 0xf6, 0x83, 0xc8, 0xba, 0x85, 0xd5, 0x2c, 0x56, 0xc5, 0xe2, 0x4c, 0x2f, 0xcc, 0x2c, 0x56,
	0x83, 0xc8, 0xaa, 0x43, 0x33, 0x9c, 0x2e, 0xa8, 0xb0, 0xcf, 0xcb, 0xa9, 0x36, 0xc5, 0x34, 0x17,
	0x8c, 0x50, 0xbb, 0xe2, 0x81, 0x56, 0xc5, 0x2f, 0x8d, 0xf5, 0x08, 0x21, 0x9d, 0xe7, 0x4c, 0x04,
	0x8a, 0xf1, 0xcc, 0x36, 0x3d, 0xd0, 0xba, 0xba, 0x6f, 0xa0, 0x92, 0x0e, 0x1d, 0xe8, 0xd0, 0xcb,
	0x81, 0xae, 0x5f, 0x5b, 0x6d, 0x5c, 0x63, 0xf9, 0xe9, 0x02, 0xff, 0xe4, 0xbf, 0xe6, 0x37, 0x80,
	0x37, 0x9a, 0xf6, 0x59, 0x71, 0x41, 0x7d, 0x4a, 0xb8, 0x88, 0xfe, 0x00, 0xfe, 0x74, 0x0a, 0x7e,
	0xdd, 0xef, 0xfc, 0x6c, 0xdc, 0x76, 0xc2, 0xd4, 0x68, 0x1a, 0x22, 0xc2, 0x53, 0x4c, 0xb8, 0x4c,
	0xb9, 0xdc, 0x7f, 0xda, 0x32, 0x1a, 0x63, 0xb5, 0xc8, 0xa9, 0x44, 0x3d, 0x42, 0x7a, 0x51, 0x24,
	0xa8, 0x94, 0xff, 0xf0, 0xd6, 0xfe, 0x70, 0xb5, 0x75, 0xc0, 0x7a, 0xeb, 0x80, 0xaf, 0xad, 0x03,
	0x96, 0x3b, 0xc7, 0x58, 0xef, 0x1c, 0xe3, 0x63, 0xe7, 0x18, 0xaf, 0x0f, 0x27, 0xac, 0xba, 0x0b,
	0x7c, 0x5f, 0x96, 0x42, 0xb4, 0xcb, 0xee, 0xcc, 0xba, 0x78, 0x7e, 0x2c, 0x90, 0xc6, 0x0f, 0xab,
	0xfa, 0xea, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x35, 0xa2, 0xc4, 0x62, 0x02, 0x00,
	0x00,
}

func (m *Offer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Offer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Offer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOffer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Price != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OfferStoreRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OfferStoreRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OfferStoreRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintOffer(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if m.Price != 0 {
		i = encodeVarintOffer(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintOffer(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovOffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Offer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovOffer(uint64(m.Price))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovOffer(uint64(l))
	return n
}

func (m *OfferStoreRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovOffer(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovOffer(uint64(m.Price))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovOffer(uint64(l))
	return n
}

func sovOffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOffer(x uint64) (n int) {
	return sovOffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Offer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffer
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
			return fmt.Errorf("proto: Offer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Offer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffer
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
func (m *OfferStoreRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffer
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
			return fmt.Errorf("proto: OfferStoreRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OfferStoreRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = append(m.Buyer[:0], dAtA[iNdEx:postIndex]...)
			if m.Buyer == nil {
				m.Buyer = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffer
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
				return ErrInvalidLengthOffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffer
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
func skipOffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOffer
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
					return 0, ErrIntOverflowOffer
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
					return 0, ErrIntOverflowOffer
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
				return 0, ErrInvalidLengthOffer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOffer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOffer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOffer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOffer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOffer = fmt.Errorf("proto: unexpected end of group")
)
