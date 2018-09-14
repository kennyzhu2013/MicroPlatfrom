// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_usersubsstatus.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TUsersubsStatus struct {
	Userid               string               `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Status               string               `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Validtime1           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=validtime1,proto3" json:"validtime1,omitempty"`
	Validtime2           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=validtime2,proto3" json:"validtime2,omitempty"`
	Mbiquto              string               `protobuf:"bytes,5,opt,name=mbiquto,proto3" json:"mbiquto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TUsersubsStatus) Reset()         { *m = TUsersubsStatus{} }
func (m *TUsersubsStatus) String() string { return proto.CompactTextString(m) }
func (*TUsersubsStatus) ProtoMessage()    {}
func (*TUsersubsStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_usersubsstatus_d428d02daeee83e6, []int{0}
}
func (m *TUsersubsStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TUsersubsStatus.Unmarshal(m, b)
}
func (m *TUsersubsStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TUsersubsStatus.Marshal(b, m, deterministic)
}
func (dst *TUsersubsStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TUsersubsStatus.Merge(dst, src)
}
func (m *TUsersubsStatus) XXX_Size() int {
	return xxx_messageInfo_TUsersubsStatus.Size(m)
}
func (m *TUsersubsStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TUsersubsStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TUsersubsStatus proto.InternalMessageInfo

func (m *TUsersubsStatus) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *TUsersubsStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TUsersubsStatus) GetValidtime1() *timestamp.Timestamp {
	if m != nil {
		return m.Validtime1
	}
	return nil
}

func (m *TUsersubsStatus) GetValidtime2() *timestamp.Timestamp {
	if m != nil {
		return m.Validtime2
	}
	return nil
}

func (m *TUsersubsStatus) GetMbiquto() string {
	if m != nil {
		return m.Mbiquto
	}
	return ""
}

func init() {
	proto.RegisterType((*TUsersubsStatus)(nil), "usercenter.TUsersubsStatus")
}

func init() {
	proto.RegisterFile("usercenter/t_usersubsstatus.proto", fileDescriptor_t_usersubsstatus_d428d02daeee83e6)
}

var fileDescriptor_t_usersubsstatus_d428d02daeee83e6 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8e, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0x75, 0xfc, 0x1c, 0x60, 0x0a, 0xa4, 0x2b, 0x90, 0x95, 0x2a, 0x50, 0xa5, 0xb2, 0xc5,
	0xd1, 0x41, 0xc7, 0x23, 0x84, 0xd0, 0xd0, 0x20, 0x3b, 0x67, 0x8c, 0x25, 0x3b, 0x6b, 0x6e, 0xd7,
	0x48, 0x3c, 0x27, 0x2f, 0x84, 0x6c, 0xe7, 0x74, 0x50, 0xa6, 0xdb, 0x59, 0xcd, 0x7c, 0x33, 0xec,
	0x26, 0xa1, 0x19, 0xb7, 0x66, 0x47, 0x66, 0x94, 0xf4, 0x96, 0x05, 0x26, 0x8d, 0x48, 0x8a, 0x12,
	0x8a, 0x38, 0x02, 0x41, 0xc7, 0x66, 0xcb, 0xe2, 0xd1, 0x3a, 0xfa, 0x48, 0x5a, 0x6c, 0x21, 0x48,
	0x0b, 0x5e, 0xed, 0xac, 0x2c, 0x26, 0x9d, 0xde, 0x65, 0xa4, 0xef, 0x68, 0x50, 0x92, 0x0b, 0x06,
	0x49, 0x85, 0x38, 0x5f, 0x15, 0x74, 0xfb, 0xd3, 0xb0, 0xab, 0xcd, 0xcb, 0xbe, 0xe2, 0xb9, 0x54,
	0x74, 0xd7, 0xac, 0xcd, 0x78, 0x37, 0xf0, 0x66, 0xd9, 0xac, 0x2e, 0xd6, 0x7b, 0x95, 0xff, 0x75,
	0x04, 0x3f, 0xaa, 0xff, 0xaa, 0xba, 0x07, 0xc6, 0xbe, 0x94, 0x77, 0x43, 0x66, 0xdf, 0xf1, 0xe3,
	0x65, 0xb3, 0xba, 0xec, 0x17, 0xc2, 0x02, 0x58, 0x6f, 0xc4, 0x34, 0x45, 0x6c, 0xa6, 0xe6, 0xf5,
	0x1f, 0xf7, 0xbf, 0x6c, 0xcf, 0x4f, 0x0e, 0xc8, 0xf6, 0x1d, 0x67, 0x67, 0x41, 0xbb, 0xcf, 0x44,
	0xc0, 0x4f, 0xcb, 0xa0, 0x49, 0x3e, 0x9d, 0xbf, 0xb6, 0x01, 0x06, 0xe3, 0x51, 0xb7, 0x85, 0x71,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0d, 0x24, 0x8d, 0x54, 0x01, 0x00, 0x00,
}
