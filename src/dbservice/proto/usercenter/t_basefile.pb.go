// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_basefile.proto

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

type TBasefile struct {
	Seq                  string               `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Datetime             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Msisdn               string               `protobuf:"bytes,3,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Mmcoin               string               `protobuf:"bytes,4,opt,name=mmcoin,proto3" json:"mmcoin,omitempty"`
	Sourcedtl            string               `protobuf:"bytes,5,opt,name=sourcedtl,proto3" json:"sourcedtl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TBasefile) Reset()         { *m = TBasefile{} }
func (m *TBasefile) String() string { return proto.CompactTextString(m) }
func (*TBasefile) ProtoMessage()    {}
func (*TBasefile) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_basefile_7ba6b399a511cdb3, []int{0}
}
func (m *TBasefile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TBasefile.Unmarshal(m, b)
}
func (m *TBasefile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TBasefile.Marshal(b, m, deterministic)
}
func (dst *TBasefile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TBasefile.Merge(dst, src)
}
func (m *TBasefile) XXX_Size() int {
	return xxx_messageInfo_TBasefile.Size(m)
}
func (m *TBasefile) XXX_DiscardUnknown() {
	xxx_messageInfo_TBasefile.DiscardUnknown(m)
}

var xxx_messageInfo_TBasefile proto.InternalMessageInfo

func (m *TBasefile) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *TBasefile) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

func (m *TBasefile) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TBasefile) GetMmcoin() string {
	if m != nil {
		return m.Mmcoin
	}
	return ""
}

func (m *TBasefile) GetSourcedtl() string {
	if m != nil {
		return m.Sourcedtl
	}
	return ""
}

func init() {
	proto.RegisterType((*TBasefile)(nil), "usercenter.TBasefile")
}

func init() {
	proto.RegisterFile("usercenter/t_basefile.proto", fileDescriptor_t_basefile_7ba6b399a511cdb3)
}

var fileDescriptor_t_basefile_7ba6b399a511cdb3 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0x15, 0x0a, 0x51, 0x63, 0x16, 0xe4, 0x01, 0x59, 0x85, 0xa1, 0x62, 0xea, 0x14, 0x4b,
	0x20, 0xb1, 0xb0, 0xf5, 0x11, 0xaa, 0x4e, 0x2c, 0x28, 0x4e, 0x2e, 0xc6, 0x52, 0xec, 0x33, 0xbe,
	0xf3, 0xc0, 0xf3, 0xf0, 0xa2, 0xa8, 0x26, 0x21, 0xdb, 0xff, 0xcb, 0x3f, 0x9f, 0x78, 0xc8, 0x04,
	0xa9, 0x87, 0xc0, 0x90, 0x34, 0x7f, 0x98, 0x8e, 0x60, 0x74, 0x13, 0xb4, 0x31, 0x21, 0xa3, 0x14,
	0x6b, 0xb9, 0x7b, 0xb3, 0x8e, 0x3f, 0xb3, 0x69, 0x7b, 0xf4, 0xda, 0xe2, 0xd4, 0x05, 0xab, 0xcb,
	0xc8, 0xe4, 0x51, 0x47, 0xfe, 0x8e, 0x40, 0x9a, 0x9d, 0x07, 0xe2, 0xce, 0xc7, 0x55, 0xfd, 0x81,
	0x9e, 0x7e, 0x2a, 0xd1, 0x9c, 0x8f, 0x33, 0x5c, 0xde, 0x89, 0x0d, 0xc1, 0x97, 0xaa, 0xf6, 0xd5,
	0xa1, 0x39, 0x5d, 0xa4, 0x7c, 0x15, 0xdb, 0xa1, 0x63, 0xb8, 0x3c, 0x53, 0x57, 0xfb, 0xea, 0x70,
	0xfb, 0xbc, 0x6b, 0x2d, 0xa2, 0x5d, 0x2e, 0x31, 0x79, 0x6c, 0xcf, 0x0b, 0xf3, 0xf4, 0xbf, 0x95,
	0xf7, 0xa2, 0xf6, 0xe4, 0x68, 0x08, 0x6a, 0x53, 0x60, 0xb3, 0x2b, 0xb9, 0xef, 0xd1, 0x05, 0x75,
	0x3d, 0xe7, 0xc5, 0xc9, 0x47, 0xd1, 0x10, 0xe6, 0xd4, 0xc3, 0xc0, 0x93, 0xba, 0x29, 0xd5, 0x1a,
	0x1c, 0xb7, 0xef, 0xb5, 0xc7, 0x01, 0x26, 0x32, 0x75, 0xf9, 0xf5, 0xe5, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x52, 0x79, 0x14, 0x1e, 0x01, 0x00, 0x00,
}