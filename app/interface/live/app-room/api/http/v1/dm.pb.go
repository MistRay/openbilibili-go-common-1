// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dm.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SendDMReq struct {
	Roomid               int64    `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty" form:"cid" validate:"required"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" validate:"required"`
	Rnd                  string   `protobuf:"bytes,3,opt,name=rnd,proto3" json:"rnd,omitempty" form:"rnd" validate:"required"`
	Fontsize             int64    `protobuf:"varint,4,opt,name=fontsize,proto3" json:"fontsize,omitempty" form:"fontsize" validate:"required"`
	Mode                 int64    `protobuf:"varint,5,opt,name=mode,proto3" json:"mode,omitempty" form:"mode" validate:"gte=0"`
	Color                int64    `protobuf:"varint,6,opt,name=color,proto3" json:"color,omitempty" form:"color" validate:"required"`
	Bubble               int64    `protobuf:"varint,7,opt,name=bubble,proto3" json:"bubble,omitempty" form:"bubble"`
	Build                int64    `protobuf:"varint,8,opt,name=build,proto3" json:"build,omitempty" form:"build"`
	Anti                 string   `protobuf:"bytes,9,opt,name=anti,proto3" json:"anti,omitempty" form:"_anti"`
	Platform             string   `protobuf:"bytes,10,opt,name=platform,proto3" json:"platform,omitempty" form:"platform"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendDMReq) Reset()         { *m = SendDMReq{} }
func (m *SendDMReq) String() string { return proto.CompactTextString(m) }
func (*SendDMReq) ProtoMessage()    {}
func (*SendDMReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dm_0c140c5b50e05ef5, []int{0}
}
func (m *SendDMReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendDMReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendDMReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SendDMReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendDMReq.Merge(dst, src)
}
func (m *SendDMReq) XXX_Size() int {
	return m.Size()
}
func (m *SendDMReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendDMReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendDMReq proto.InternalMessageInfo

func (m *SendDMReq) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *SendDMReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SendDMReq) GetRnd() string {
	if m != nil {
		return m.Rnd
	}
	return ""
}

func (m *SendDMReq) GetFontsize() int64 {
	if m != nil {
		return m.Fontsize
	}
	return 0
}

func (m *SendDMReq) GetMode() int64 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *SendDMReq) GetColor() int64 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *SendDMReq) GetBubble() int64 {
	if m != nil {
		return m.Bubble
	}
	return 0
}

func (m *SendDMReq) GetBuild() int64 {
	if m != nil {
		return m.Build
	}
	return 0
}

func (m *SendDMReq) GetAnti() string {
	if m != nil {
		return m.Anti
	}
	return ""
}

func (m *SendDMReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type SendMsgResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgResp) Reset()         { *m = SendMsgResp{} }
func (m *SendMsgResp) String() string { return proto.CompactTextString(m) }
func (*SendMsgResp) ProtoMessage()    {}
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dm_0c140c5b50e05ef5, []int{1}
}
func (m *SendMsgResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMsgResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SendMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgResp.Merge(dst, src)
}
func (m *SendMsgResp) XXX_Size() int {
	return m.Size()
}
func (m *SendMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgResp proto.InternalMessageInfo

type HistoryReq struct {
	Roomid               int64    `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty" form:"room_id" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryReq) Reset()         { *m = HistoryReq{} }
func (m *HistoryReq) String() string { return proto.CompactTextString(m) }
func (*HistoryReq) ProtoMessage()    {}
func (*HistoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dm_0c140c5b50e05ef5, []int{2}
}
func (m *HistoryReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HistoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryReq.Merge(dst, src)
}
func (m *HistoryReq) XXX_Size() int {
	return m.Size()
}
func (m *HistoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryReq proto.InternalMessageInfo

func (m *HistoryReq) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type HistoryResp struct {
	Room                 []string `protobuf:"bytes,1,rep,name=Room" json:"Room,omitempty"`
	Admin                []string `protobuf:"bytes,2,rep,name=Admin" json:"Admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryResp) Reset()         { *m = HistoryResp{} }
func (m *HistoryResp) String() string { return proto.CompactTextString(m) }
func (*HistoryResp) ProtoMessage()    {}
func (*HistoryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dm_0c140c5b50e05ef5, []int{3}
}
func (m *HistoryResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HistoryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryResp.Merge(dst, src)
}
func (m *HistoryResp) XXX_Size() int {
	return m.Size()
}
func (m *HistoryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryResp.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryResp proto.InternalMessageInfo

func (m *HistoryResp) GetRoom() []string {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *HistoryResp) GetAdmin() []string {
	if m != nil {
		return m.Admin
	}
	return nil
}

func init() {
	proto.RegisterType((*SendDMReq)(nil), "live.approom.v1.SendDMReq")
	proto.RegisterType((*SendMsgResp)(nil), "live.approom.v1.SendMsgResp")
	proto.RegisterType((*HistoryReq)(nil), "live.approom.v1.HistoryReq")
	proto.RegisterType((*HistoryResp)(nil), "live.approom.v1.HistoryResp")
}
func (m *SendDMReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendDMReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Roomid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Roomid))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDm(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Rnd) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDm(dAtA, i, uint64(len(m.Rnd)))
		i += copy(dAtA[i:], m.Rnd)
	}
	if m.Fontsize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Fontsize))
	}
	if m.Mode != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Mode))
	}
	if m.Color != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Color))
	}
	if m.Bubble != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Bubble))
	}
	if m.Build != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Build))
	}
	if len(m.Anti) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintDm(dAtA, i, uint64(len(m.Anti)))
		i += copy(dAtA[i:], m.Anti)
	}
	if len(m.Platform) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintDm(dAtA, i, uint64(len(m.Platform)))
		i += copy(dAtA[i:], m.Platform)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SendMsgResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMsgResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HistoryReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Roomid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDm(dAtA, i, uint64(m.Roomid))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HistoryResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Room) > 0 {
		for _, s := range m.Room {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Admin) > 0 {
		for _, s := range m.Admin {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SendDMReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Roomid != 0 {
		n += 1 + sovDm(uint64(m.Roomid))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovDm(uint64(l))
	}
	l = len(m.Rnd)
	if l > 0 {
		n += 1 + l + sovDm(uint64(l))
	}
	if m.Fontsize != 0 {
		n += 1 + sovDm(uint64(m.Fontsize))
	}
	if m.Mode != 0 {
		n += 1 + sovDm(uint64(m.Mode))
	}
	if m.Color != 0 {
		n += 1 + sovDm(uint64(m.Color))
	}
	if m.Bubble != 0 {
		n += 1 + sovDm(uint64(m.Bubble))
	}
	if m.Build != 0 {
		n += 1 + sovDm(uint64(m.Build))
	}
	l = len(m.Anti)
	if l > 0 {
		n += 1 + l + sovDm(uint64(l))
	}
	l = len(m.Platform)
	if l > 0 {
		n += 1 + l + sovDm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendMsgResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HistoryReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Roomid != 0 {
		n += 1 + sovDm(uint64(m.Roomid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HistoryResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Room) > 0 {
		for _, s := range m.Room {
			l = len(s)
			n += 1 + l + sovDm(uint64(l))
		}
	}
	if len(m.Admin) > 0 {
		for _, s := range m.Admin {
			l = len(s)
			n += 1 + l + sovDm(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDm(x uint64) (n int) {
	return sovDm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendDMReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendDMReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendDMReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rnd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rnd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fontsize", wireType)
			}
			m.Fontsize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fontsize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Color", wireType)
			}
			m.Color = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Color |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bubble", wireType)
			}
			m.Bubble = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bubble |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			m.Build = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Build |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anti", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Anti = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SendMsgResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendMsgResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMsgResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HistoryReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistoryReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roomid", wireType)
			}
			m.Roomid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Roomid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HistoryResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistoryResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Room", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Room = append(m.Room, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDm
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDm
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
					return 0, ErrIntOverflowDm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDm
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDm
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDm(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDm   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dm.proto", fileDescriptor_dm_0c140c5b50e05ef5) }

var fileDescriptor_dm_0c140c5b50e05ef5 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x95, 0x36, 0xed, 0xda, 0x37, 0xa6, 0x0d, 0xc3, 0xc1, 0x1a, 0x53, 0x13, 0x99, 0x69,
	0x94, 0x03, 0x29, 0x83, 0x49, 0x48, 0x95, 0x40, 0x62, 0x9b, 0x34, 0x0e, 0xf4, 0x12, 0x6e, 0x5c,
	0xa6, 0xa4, 0x76, 0x83, 0xa5, 0x38, 0xce, 0x12, 0xa7, 0x12, 0xfc, 0x0c, 0xc4, 0x8f, 0xe2, 0xc8,
	0x2f, 0x88, 0x50, 0x7f, 0x42, 0x7e, 0x01, 0xb2, 0x93, 0x96, 0x52, 0x82, 0x38, 0xd5, 0x7e, 0xef,
	0xfb, 0x9e, 0x9f, 0xdd, 0x17, 0x18, 0x50, 0xe1, 0xa5, 0x99, 0x54, 0x12, 0x1d, 0xc6, 0x7c, 0xc9,
	0xbc, 0x20, 0x4d, 0x33, 0x29, 0x85, 0xb7, 0x3c, 0x3f, 0x7e, 0x16, 0x71, 0xf5, 0xa9, 0x08, 0xbd,
	0xb9, 0x14, 0x93, 0x48, 0x46, 0x72, 0x62, 0xb8, 0xb0, 0x58, 0x98, 0x9d, 0xd9, 0x98, 0x55, 0xed,
	0x93, 0x6f, 0x36, 0x0c, 0x3f, 0xb0, 0x84, 0x5e, 0xcf, 0x7c, 0x76, 0x87, 0xa6, 0xd0, 0xd7, 0x75,
	0x38, 0xc5, 0x96, 0x6b, 0x8d, 0xbb, 0x97, 0xa4, 0x2a, 0x9d, 0xd1, 0x42, 0x66, 0x62, 0x4a, 0xe6,
	0x9c, 0x12, 0x77, 0x19, 0xc4, 0x9c, 0x06, 0x8a, 0x4d, 0x49, 0xc6, 0xee, 0x0a, 0x9e, 0x31, 0x4a,
	0xfc, 0xc6, 0x40, 0x17, 0xd0, 0x15, 0x79, 0x84, 0x3b, 0xae, 0x35, 0x1e, 0x6e, 0x8b, 0x22, 0x8f,
	0xda, 0x45, 0x8d, 0x6b, 0x2b, 0x4b, 0x28, 0xee, 0xee, 0x5a, 0x59, 0xf2, 0x8f, 0xe3, 0x34, 0x8e,
	0xae, 0x60, 0xb0, 0x90, 0x89, 0xca, 0xf9, 0x17, 0x86, 0x6d, 0xd3, 0xe9, 0x93, 0xaa, 0x74, 0x1e,
	0xd7, 0xea, 0x3a, 0xd3, 0xee, 0x6f, 0x44, 0x74, 0x01, 0xb6, 0x90, 0x94, 0xe1, 0x9e, 0x29, 0xe0,
	0x56, 0xa5, 0x73, 0xd2, 0x74, 0x2c, 0xe9, 0x1f, 0x72, 0xa4, 0xd8, 0xeb, 0xe7, 0xc4, 0x37, 0x34,
	0x9a, 0x42, 0x6f, 0x2e, 0x63, 0x99, 0xe1, 0xbe, 0xd1, 0x4e, 0xab, 0xd2, 0x71, 0x9b, 0x17, 0xd2,
	0xe1, 0xf6, 0x43, 0x6b, 0x05, 0x3d, 0x85, 0x7e, 0x58, 0x84, 0x61, 0xcc, 0xf0, 0x9e, 0x91, 0xef,
	0x57, 0xa5, 0x73, 0x50, 0xcb, 0x75, 0x9c, 0xf8, 0x0d, 0x80, 0xce, 0xa0, 0x17, 0x16, 0x3c, 0xa6,
	0x78, 0x60, 0xc8, 0xa3, 0xaa, 0x74, 0xee, 0xad, 0x49, 0x1e, 0xeb, 0x92, 0xe6, 0x17, 0x9d, 0x82,
	0x1d, 0x24, 0x8a, 0xe3, 0xa1, 0x79, 0xc0, 0x2d, 0xec, 0x56, 0x87, 0x89, 0x6f, 0xb2, 0x68, 0x02,
	0x83, 0x34, 0x0e, 0x94, 0xce, 0x60, 0x30, 0xe4, 0x83, 0xaa, 0x74, 0x0e, 0x6b, 0x72, 0x9d, 0x21,
	0xfe, 0x06, 0x22, 0x07, 0xb0, 0xaf, 0xa7, 0x62, 0x96, 0x47, 0x3e, 0xcb, 0x53, 0xf2, 0x1e, 0xe0,
	0x1d, 0xcf, 0x95, 0xcc, 0x3e, 0xeb, 0x29, 0x79, 0xb3, 0x33, 0x25, 0x67, 0x55, 0xe9, 0x90, 0xe6,
	0x6f, 0x93, 0x52, 0xdc, 0xfe, 0x67, 0x52, 0xc8, 0x2b, 0xd8, 0xdf, 0x54, 0xcb, 0x53, 0x84, 0xc0,
	0xf6, 0xa5, 0x14, 0xd8, 0x72, 0xbb, 0xe3, 0xa1, 0x6f, 0xd6, 0xe8, 0x21, 0xf4, 0xde, 0x52, 0xc1,
	0x13, 0xdc, 0x31, 0xc1, 0x7a, 0xf3, 0xe2, 0xab, 0x05, 0x9d, 0xeb, 0x19, 0xba, 0x82, 0xbd, 0xa6,
	0x39, 0x74, 0xec, 0xed, 0xcc, 0xbf, 0xb7, 0x19, 0xe6, 0xe3, 0x93, 0xd6, 0x5c, 0x73, 0x25, 0x74,
	0x03, 0x70, 0xc3, 0x54, 0xd3, 0x07, 0x7a, 0xf4, 0x17, 0xfb, 0xfb, 0xbe, 0x2d, 0x85, 0xb6, 0xda,
	0xbf, 0x3c, 0xfa, 0xbe, 0x1a, 0x59, 0x3f, 0x56, 0x23, 0xeb, 0xe7, 0x6a, 0x64, 0x7d, 0xec, 0x2c,
	0xcf, 0xc3, 0xbe, 0xf9, 0xb4, 0x5e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc3, 0x62, 0xe4,
	0xa6, 0x03, 0x00, 0x00,
}
