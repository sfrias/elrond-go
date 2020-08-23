// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miniblockMetadata.proto

package fullHistory

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// MiniblockMetadata is used to store information about a history transaction
type MiniblockMetadata struct {
	SourceShardID                     uint32 `protobuf:"varint,1,opt,name=SourceShardID,proto3" json:"SourceShardID,omitempty"`
	DestinationShardID                uint32 `protobuf:"varint,2,opt,name=DestinationShardID,proto3" json:"DestinationShardID,omitempty"`
	Round                             uint64 `protobuf:"varint,3,opt,name=Round,proto3" json:"Round,omitempty"`
	HeaderNonce                       uint64 `protobuf:"varint,4,opt,name=HeaderNonce,proto3" json:"HeaderNonce,omitempty"`
	HeaderHash                        []byte `protobuf:"bytes,5,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	MiniblockHash                     []byte `protobuf:"bytes,6,opt,name=MiniblockHash,proto3" json:"MiniblockHash,omitempty"`
	Status                            []byte `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Epoch                             uint32 `protobuf:"varint,8,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	NotarizedAtSourceInMetaNonce      uint64 `protobuf:"varint,9,opt,name=NotarizedAtSourceInMetaNonce,proto3" json:"NotarizedAtSourceInMetaNonce,omitempty"`
	NotarizedAtDestinationInMetaNonce uint64 `protobuf:"varint,10,opt,name=NotarizedAtDestinationInMetaNonce,proto3" json:"NotarizedAtDestinationInMetaNonce,omitempty"`
	NotarizedAtSourceInMetaHash       []byte `protobuf:"bytes,11,opt,name=NotarizedAtSourceInMetaHash,proto3" json:"NotarizedAtSourceInMetaHash,omitempty"`
	NotarizedAtDestinationInMetaHash  []byte `protobuf:"bytes,12,opt,name=NotarizedAtDestinationInMetaHash,proto3" json:"NotarizedAtDestinationInMetaHash,omitempty"`
}

func (m *MiniblockMetadata) Reset()      { *m = MiniblockMetadata{} }
func (*MiniblockMetadata) ProtoMessage() {}
func (*MiniblockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd82f29831cbb1fe, []int{0}
}
func (m *MiniblockMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MiniblockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MiniblockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniblockMetadata.Merge(m, src)
}
func (m *MiniblockMetadata) XXX_Size() int {
	return m.Size()
}
func (m *MiniblockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniblockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MiniblockMetadata proto.InternalMessageInfo

func (m *MiniblockMetadata) GetSourceShardID() uint32 {
	if m != nil {
		return m.SourceShardID
	}
	return 0
}

func (m *MiniblockMetadata) GetDestinationShardID() uint32 {
	if m != nil {
		return m.DestinationShardID
	}
	return 0
}

func (m *MiniblockMetadata) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *MiniblockMetadata) GetHeaderNonce() uint64 {
	if m != nil {
		return m.HeaderNonce
	}
	return 0
}

func (m *MiniblockMetadata) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *MiniblockMetadata) GetMiniblockHash() []byte {
	if m != nil {
		return m.MiniblockHash
	}
	return nil
}

func (m *MiniblockMetadata) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MiniblockMetadata) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *MiniblockMetadata) GetNotarizedAtSourceInMetaNonce() uint64 {
	if m != nil {
		return m.NotarizedAtSourceInMetaNonce
	}
	return 0
}

func (m *MiniblockMetadata) GetNotarizedAtDestinationInMetaNonce() uint64 {
	if m != nil {
		return m.NotarizedAtDestinationInMetaNonce
	}
	return 0
}

func (m *MiniblockMetadata) GetNotarizedAtSourceInMetaHash() []byte {
	if m != nil {
		return m.NotarizedAtSourceInMetaHash
	}
	return nil
}

func (m *MiniblockMetadata) GetNotarizedAtDestinationInMetaHash() []byte {
	if m != nil {
		return m.NotarizedAtDestinationInMetaHash
	}
	return nil
}

func init() {
	proto.RegisterType((*MiniblockMetadata)(nil), "proto.MiniblockMetadata")
}

func init() { proto.RegisterFile("miniblockMetadata.proto", fileDescriptor_cd82f29831cbb1fe) }

var fileDescriptor_cd82f29831cbb1fe = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x3b, 0xdf, 0x07, 0xa8, 0x03, 0x2c, 0x9c, 0x18, 0x6d, 0xd4, 0xdc, 0x54, 0xe3, 0x82,
	0x8d, 0xb0, 0xf0, 0x05, 0x94, 0x40, 0x02, 0x46, 0x58, 0x94, 0x9d, 0xbb, 0xe9, 0x1f, 0xe8, 0x44,
	0xe8, 0x90, 0x76, 0xba, 0xd0, 0x95, 0x8f, 0xe0, 0x63, 0xf8, 0x02, 0xbe, 0x83, 0x4b, 0x96, 0x2c,
	0x65, 0xd8, 0xb8, 0xe4, 0x11, 0x0c, 0xb7, 0x6a, 0x8a, 0x51, 0x58, 0xb5, 0xe7, 0xdc, 0x5f, 0xcf,
	0x9c, 0x3b, 0x29, 0x3d, 0x18, 0x89, 0x50, 0x38, 0x43, 0xe9, 0xde, 0x75, 0x7c, 0xc5, 0x3d, 0xae,
	0x78, 0x75, 0x1c, 0x49, 0x25, 0x59, 0x1e, 0x1f, 0x87, 0xe7, 0x03, 0xa1, 0x82, 0xc4, 0xa9, 0xba,
	0x72, 0x54, 0x1b, 0xc8, 0x81, 0xac, 0xa1, 0xed, 0x24, 0x7d, 0x54, 0x28, 0xf0, 0x2d, 0xfd, 0xea,
	0xf4, 0x25, 0x47, 0x77, 0x3b, 0x3f, 0x13, 0xd9, 0x19, 0x2d, 0xf7, 0x64, 0x12, 0xb9, 0x7e, 0x2f,
	0xe0, 0x91, 0xd7, 0x6e, 0x98, 0xc4, 0x22, 0x95, 0xb2, 0xbd, 0x6a, 0xb2, 0x2a, 0x65, 0x0d, 0x3f,
	0x56, 0x22, 0xe4, 0x4a, 0xc8, 0xf0, 0x0b, 0xfd, 0x87, 0xe8, 0x2f, 0x13, 0xb6, 0x47, 0xf3, 0xb6,
	0x4c, 0x42, 0xcf, 0xfc, 0x6f, 0x91, 0x4a, 0xce, 0x4e, 0x05, 0xb3, 0x68, 0xb1, 0xe5, 0x73, 0xcf,
	0x8f, 0xba, 0x32, 0x74, 0x7d, 0x33, 0x87, 0xb3, 0xac, 0xc5, 0x80, 0xd2, 0x54, 0xb6, 0x78, 0x1c,
	0x98, 0x79, 0x8b, 0x54, 0x4a, 0x76, 0xc6, 0x59, 0xb6, 0xfd, 0x5e, 0x01, 0x91, 0x02, 0x22, 0xab,
	0x26, 0xdb, 0xa7, 0x85, 0x9e, 0xe2, 0x2a, 0x89, 0xcd, 0x2d, 0x1c, 0x7f, 0xaa, 0x65, 0xab, 0xe6,
	0x58, 0xba, 0x81, 0xb9, 0x8d, 0xc5, 0x53, 0xc1, 0xea, 0xf4, 0xb8, 0x2b, 0x15, 0x8f, 0xc4, 0x83,
	0xef, 0x5d, 0xa9, 0x74, 0xef, 0x76, 0xb8, 0xbc, 0xa0, 0xb4, 0xe6, 0x0e, 0xd6, 0x5c, 0xcb, 0xb0,
	0x1b, 0x7a, 0x92, 0x99, 0x67, 0x2e, 0x24, 0x1b, 0x44, 0x31, 0x68, 0x33, 0xc8, 0x2e, 0xe9, 0xd1,
	0x1f, 0xa7, 0xe1, 0xce, 0x45, 0x5c, 0x6a, 0x1d, 0xc2, 0xae, 0xa9, 0xb5, 0xee, 0x18, 0x8c, 0x29,
	0x61, 0xcc, 0x46, 0xae, 0xde, 0x9c, 0xcc, 0xc0, 0x98, 0xce, 0xc0, 0x58, 0xcc, 0x80, 0x3c, 0x6a,
	0x20, 0xcf, 0x1a, 0xc8, 0xab, 0x06, 0x32, 0xd1, 0x40, 0xa6, 0x1a, 0xc8, 0x9b, 0x06, 0xf2, 0xae,
	0xc1, 0x58, 0x68, 0x20, 0x4f, 0x73, 0x30, 0x26, 0x73, 0x30, 0xa6, 0x73, 0x30, 0x6e, 0x8b, 0xfd,
	0x64, 0x38, 0x6c, 0x89, 0x58, 0xc9, 0xe8, 0xde, 0x29, 0xe0, 0x5f, 0x78, 0xf1, 0x11, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0xfe, 0x60, 0x04, 0xd6, 0x02, 0x00, 0x00,
}

func (this *MiniblockMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MiniblockMetadata)
	if !ok {
		that2, ok := that.(MiniblockMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SourceShardID != that1.SourceShardID {
		return false
	}
	if this.DestinationShardID != that1.DestinationShardID {
		return false
	}
	if this.Round != that1.Round {
		return false
	}
	if this.HeaderNonce != that1.HeaderNonce {
		return false
	}
	if !bytes.Equal(this.HeaderHash, that1.HeaderHash) {
		return false
	}
	if !bytes.Equal(this.MiniblockHash, that1.MiniblockHash) {
		return false
	}
	if !bytes.Equal(this.Status, that1.Status) {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	if this.NotarizedAtSourceInMetaNonce != that1.NotarizedAtSourceInMetaNonce {
		return false
	}
	if this.NotarizedAtDestinationInMetaNonce != that1.NotarizedAtDestinationInMetaNonce {
		return false
	}
	if !bytes.Equal(this.NotarizedAtSourceInMetaHash, that1.NotarizedAtSourceInMetaHash) {
		return false
	}
	if !bytes.Equal(this.NotarizedAtDestinationInMetaHash, that1.NotarizedAtDestinationInMetaHash) {
		return false
	}
	return true
}
func (this *MiniblockMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 16)
	s = append(s, "&fullHistory.MiniblockMetadata{")
	s = append(s, "SourceShardID: "+fmt.Sprintf("%#v", this.SourceShardID)+",\n")
	s = append(s, "DestinationShardID: "+fmt.Sprintf("%#v", this.DestinationShardID)+",\n")
	s = append(s, "Round: "+fmt.Sprintf("%#v", this.Round)+",\n")
	s = append(s, "HeaderNonce: "+fmt.Sprintf("%#v", this.HeaderNonce)+",\n")
	s = append(s, "HeaderHash: "+fmt.Sprintf("%#v", this.HeaderHash)+",\n")
	s = append(s, "MiniblockHash: "+fmt.Sprintf("%#v", this.MiniblockHash)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "NotarizedAtSourceInMetaNonce: "+fmt.Sprintf("%#v", this.NotarizedAtSourceInMetaNonce)+",\n")
	s = append(s, "NotarizedAtDestinationInMetaNonce: "+fmt.Sprintf("%#v", this.NotarizedAtDestinationInMetaNonce)+",\n")
	s = append(s, "NotarizedAtSourceInMetaHash: "+fmt.Sprintf("%#v", this.NotarizedAtSourceInMetaHash)+",\n")
	s = append(s, "NotarizedAtDestinationInMetaHash: "+fmt.Sprintf("%#v", this.NotarizedAtDestinationInMetaHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMiniblockMetadata(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MiniblockMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MiniblockMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MiniblockMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NotarizedAtDestinationInMetaHash) > 0 {
		i -= len(m.NotarizedAtDestinationInMetaHash)
		copy(dAtA[i:], m.NotarizedAtDestinationInMetaHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.NotarizedAtDestinationInMetaHash)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.NotarizedAtSourceInMetaHash) > 0 {
		i -= len(m.NotarizedAtSourceInMetaHash)
		copy(dAtA[i:], m.NotarizedAtSourceInMetaHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.NotarizedAtSourceInMetaHash)))
		i--
		dAtA[i] = 0x5a
	}
	if m.NotarizedAtDestinationInMetaNonce != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.NotarizedAtDestinationInMetaNonce))
		i--
		dAtA[i] = 0x50
	}
	if m.NotarizedAtSourceInMetaNonce != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.NotarizedAtSourceInMetaNonce))
		i--
		dAtA[i] = 0x48
	}
	if m.Epoch != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MiniblockHash) > 0 {
		i -= len(m.MiniblockHash)
		copy(dAtA[i:], m.MiniblockHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.MiniblockHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.HeaderNonce != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.HeaderNonce))
		i--
		dAtA[i] = 0x20
	}
	if m.Round != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	if m.DestinationShardID != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.DestinationShardID))
		i--
		dAtA[i] = 0x10
	}
	if m.SourceShardID != 0 {
		i = encodeVarintMiniblockMetadata(dAtA, i, uint64(m.SourceShardID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMiniblockMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMiniblockMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MiniblockMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceShardID != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.SourceShardID))
	}
	if m.DestinationShardID != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.DestinationShardID))
	}
	if m.Round != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.Round))
	}
	if m.HeaderNonce != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.HeaderNonce))
	}
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	l = len(m.MiniblockHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.Epoch))
	}
	if m.NotarizedAtSourceInMetaNonce != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.NotarizedAtSourceInMetaNonce))
	}
	if m.NotarizedAtDestinationInMetaNonce != 0 {
		n += 1 + sovMiniblockMetadata(uint64(m.NotarizedAtDestinationInMetaNonce))
	}
	l = len(m.NotarizedAtSourceInMetaHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	l = len(m.NotarizedAtDestinationInMetaHash)
	if l > 0 {
		n += 1 + l + sovMiniblockMetadata(uint64(l))
	}
	return n
}

func sovMiniblockMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMiniblockMetadata(x uint64) (n int) {
	return sovMiniblockMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MiniblockMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MiniblockMetadata{`,
		`SourceShardID:` + fmt.Sprintf("%v", this.SourceShardID) + `,`,
		`DestinationShardID:` + fmt.Sprintf("%v", this.DestinationShardID) + `,`,
		`Round:` + fmt.Sprintf("%v", this.Round) + `,`,
		`HeaderNonce:` + fmt.Sprintf("%v", this.HeaderNonce) + `,`,
		`HeaderHash:` + fmt.Sprintf("%v", this.HeaderHash) + `,`,
		`MiniblockHash:` + fmt.Sprintf("%v", this.MiniblockHash) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`NotarizedAtSourceInMetaNonce:` + fmt.Sprintf("%v", this.NotarizedAtSourceInMetaNonce) + `,`,
		`NotarizedAtDestinationInMetaNonce:` + fmt.Sprintf("%v", this.NotarizedAtDestinationInMetaNonce) + `,`,
		`NotarizedAtSourceInMetaHash:` + fmt.Sprintf("%v", this.NotarizedAtSourceInMetaHash) + `,`,
		`NotarizedAtDestinationInMetaHash:` + fmt.Sprintf("%v", this.NotarizedAtDestinationInMetaHash) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMiniblockMetadata(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MiniblockMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMiniblockMetadata
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
			return fmt.Errorf("proto: MiniblockMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MiniblockMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceShardID", wireType)
			}
			m.SourceShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationShardID", wireType)
			}
			m.DestinationShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationShardID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderNonce", wireType)
			}
			m.HeaderNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeaderNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiniblockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MiniblockHash = append(m.MiniblockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MiniblockHash == nil {
				m.MiniblockHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = append(m.Status[:0], dAtA[iNdEx:postIndex]...)
			if m.Status == nil {
				m.Status = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotarizedAtSourceInMetaNonce", wireType)
			}
			m.NotarizedAtSourceInMetaNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotarizedAtSourceInMetaNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotarizedAtDestinationInMetaNonce", wireType)
			}
			m.NotarizedAtDestinationInMetaNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotarizedAtDestinationInMetaNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotarizedAtSourceInMetaHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NotarizedAtSourceInMetaHash = append(m.NotarizedAtSourceInMetaHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NotarizedAtSourceInMetaHash == nil {
				m.NotarizedAtSourceInMetaHash = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotarizedAtDestinationInMetaHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMiniblockMetadata
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
				return ErrInvalidLengthMiniblockMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NotarizedAtDestinationInMetaHash = append(m.NotarizedAtDestinationInMetaHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NotarizedAtDestinationInMetaHash == nil {
				m.NotarizedAtDestinationInMetaHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMiniblockMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMiniblockMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMiniblockMetadata
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
func skipMiniblockMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMiniblockMetadata
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
					return 0, ErrIntOverflowMiniblockMetadata
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
					return 0, ErrIntOverflowMiniblockMetadata
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
				return 0, ErrInvalidLengthMiniblockMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMiniblockMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMiniblockMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMiniblockMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMiniblockMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMiniblockMetadata = fmt.Errorf("proto: unexpected end of group")
)