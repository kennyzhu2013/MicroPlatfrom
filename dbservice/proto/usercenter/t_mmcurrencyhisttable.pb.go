// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mmcurrencyhisttable.proto

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

type TMmcurrencyHistTable struct {
	Seq                  string               `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Msisdn               string               `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Value                string               `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Opertype             string               `protobuf:"bytes,4,opt,name=opertype,proto3" json:"opertype,omitempty"`
	Operreason           string               `protobuf:"bytes,5,opt,name=operreason,proto3" json:"operreason,omitempty"`
	Opertime             *timestamp.Timestamp `protobuf:"bytes,6,opt,name=opertime,proto3" json:"opertime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMmcurrencyHistTable) Reset()         { *m = TMmcurrencyHistTable{} }
func (m *TMmcurrencyHistTable) String() string { return proto.CompactTextString(m) }
func (*TMmcurrencyHistTable) ProtoMessage()    {}
func (*TMmcurrencyHistTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mmcurrencyhisttable_f2581f92c38dace9, []int{0}
}
func (m *TMmcurrencyHistTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMmcurrencyHistTable.Unmarshal(m, b)
}
func (m *TMmcurrencyHistTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMmcurrencyHistTable.Marshal(b, m, deterministic)
}
func (dst *TMmcurrencyHistTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMmcurrencyHistTable.Merge(dst, src)
}
func (m *TMmcurrencyHistTable) XXX_Size() int {
	return xxx_messageInfo_TMmcurrencyHistTable.Size(m)
}
func (m *TMmcurrencyHistTable) XXX_DiscardUnknown() {
	xxx_messageInfo_TMmcurrencyHistTable.DiscardUnknown(m)
}

var xxx_messageInfo_TMmcurrencyHistTable proto.InternalMessageInfo

func (m *TMmcurrencyHistTable) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *TMmcurrencyHistTable) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TMmcurrencyHistTable) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TMmcurrencyHistTable) GetOpertype() string {
	if m != nil {
		return m.Opertype
	}
	return ""
}

func (m *TMmcurrencyHistTable) GetOperreason() string {
	if m != nil {
		return m.Operreason
	}
	return ""
}

func (m *TMmcurrencyHistTable) GetOpertime() *timestamp.Timestamp {
	if m != nil {
		return m.Opertime
	}
	return nil
}

func init() {
	proto.RegisterType((*TMmcurrencyHistTable)(nil), "usercenter.TMmcurrencyHistTable")
}

func init() {
	proto.RegisterFile("usercenter/t_mmcurrencyhisttable.proto", fileDescriptor_t_mmcurrencyhisttable_f2581f92c38dace9)
}

var fileDescriptor_t_mmcurrencyhisttable_f2581f92c38dace9 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x99, 0xaf, 0x5f, 0x87, 0x1a, 0x37, 0x12, 0x8a, 0x84, 0x59, 0x48, 0x71, 0x21, 0x5d,
	0x4d, 0x40, 0xc1, 0x8d, 0x3b, 0x57, 0x6e, 0xdc, 0x94, 0x59, 0xb9, 0x91, 0x64, 0x7a, 0x9b, 0x06,
	0x26, 0x7f, 0xcc, 0xbd, 0x11, 0xfa, 0x8a, 0x3e, 0x95, 0x4c, 0xda, 0x4e, 0xdd, 0x9d, 0x73, 0xee,
	0x2f, 0x39, 0x1c, 0xf6, 0x90, 0x11, 0x52, 0x0f, 0x9e, 0x20, 0x49, 0xfa, 0x74, 0xae, 0xcf, 0x29,
	0x81, 0xef, 0x0f, 0x7b, 0x8b, 0x44, 0x4a, 0x0f, 0xd0, 0xc6, 0x14, 0x28, 0x70, 0x76, 0xe1, 0x9a,
	0x17, 0x63, 0x69, 0x9f, 0x75, 0xdb, 0x07, 0x27, 0x4d, 0x18, 0x94, 0x37, 0xb2, 0x40, 0x3a, 0xef,
	0x64, 0xa4, 0x43, 0x04, 0x94, 0x64, 0x1d, 0x20, 0x29, 0x17, 0x2f, 0xea, 0xf8, 0xd1, 0xfd, 0x4f,
	0xc5, 0x96, 0xdd, 0xfb, 0xd4, 0xf3, 0x66, 0x91, 0xba, 0xb1, 0x87, 0xdf, 0xb0, 0x19, 0xc2, 0x97,
	0xa8, 0x56, 0xd5, 0xfa, 0x6a, 0x33, 0x4a, 0x7e, 0xcb, 0x6a, 0x87, 0x16, 0xb7, 0x5e, 0xfc, 0x2b,
	0xe1, 0xc9, 0xf1, 0x25, 0x9b, 0x7f, 0xab, 0x21, 0x83, 0x98, 0x95, 0xf8, 0x68, 0x78, 0xc3, 0x16,
	0x21, 0x42, 0x1a, 0xfb, 0xc5, 0xff, 0x72, 0x98, 0x3c, 0xbf, 0x63, 0x6c, 0xd4, 0x09, 0x14, 0x06,
	0x2f, 0xe6, 0xe5, 0xfa, 0x27, 0xe1, 0xcf, 0xa7, 0xb7, 0xd6, 0x81, 0xa8, 0x57, 0xd5, 0xfa, 0xfa,
	0xb1, 0x69, 0x4d, 0x08, 0xe6, 0x3c, 0x5f, 0xe7, 0x5d, 0xdb, 0x9d, 0x87, 0x6c, 0x26, 0xf6, 0x75,
	0xf1, 0x51, 0xbb, 0xb0, 0x85, 0x01, 0x75, 0x5d, 0xb8, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x74, 0x18, 0xab, 0x50, 0x01, 0x00, 0x00,
}