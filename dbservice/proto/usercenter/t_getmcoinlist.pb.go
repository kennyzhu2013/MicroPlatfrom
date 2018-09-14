// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_getmcoinlist.proto

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

type TGetmcoinlist struct {
	Seq                  string               `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Msisdn               string               `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Getdate              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=getdate,proto3" json:"getdate,omitempty"`
	Expiredate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expiredate,proto3" json:"expiredate,omitempty"`
	Status               string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Value                string               `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Balance              string               `protobuf:"bytes,7,opt,name=balance,proto3" json:"balance,omitempty"`
	Descr                string               `protobuf:"bytes,8,opt,name=descr,proto3" json:"descr,omitempty"`
	Mcoin                string               `protobuf:"bytes,9,opt,name=mcoin,proto3" json:"mcoin,omitempty"`
	Mcoinbalance         string               `protobuf:"bytes,10,opt,name=mcoinbalance,proto3" json:"mcoinbalance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TGetmcoinlist) Reset()         { *m = TGetmcoinlist{} }
func (m *TGetmcoinlist) String() string { return proto.CompactTextString(m) }
func (*TGetmcoinlist) ProtoMessage()    {}
func (*TGetmcoinlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_getmcoinlist_67be56ed7cd22a4b, []int{0}
}
func (m *TGetmcoinlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TGetmcoinlist.Unmarshal(m, b)
}
func (m *TGetmcoinlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TGetmcoinlist.Marshal(b, m, deterministic)
}
func (dst *TGetmcoinlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TGetmcoinlist.Merge(dst, src)
}
func (m *TGetmcoinlist) XXX_Size() int {
	return xxx_messageInfo_TGetmcoinlist.Size(m)
}
func (m *TGetmcoinlist) XXX_DiscardUnknown() {
	xxx_messageInfo_TGetmcoinlist.DiscardUnknown(m)
}

var xxx_messageInfo_TGetmcoinlist proto.InternalMessageInfo

func (m *TGetmcoinlist) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *TGetmcoinlist) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TGetmcoinlist) GetGetdate() *timestamp.Timestamp {
	if m != nil {
		return m.Getdate
	}
	return nil
}

func (m *TGetmcoinlist) GetExpiredate() *timestamp.Timestamp {
	if m != nil {
		return m.Expiredate
	}
	return nil
}

func (m *TGetmcoinlist) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TGetmcoinlist) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TGetmcoinlist) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *TGetmcoinlist) GetDescr() string {
	if m != nil {
		return m.Descr
	}
	return ""
}

func (m *TGetmcoinlist) GetMcoin() string {
	if m != nil {
		return m.Mcoin
	}
	return ""
}

func (m *TGetmcoinlist) GetMcoinbalance() string {
	if m != nil {
		return m.Mcoinbalance
	}
	return ""
}

func init() {
	proto.RegisterType((*TGetmcoinlist)(nil), "usercenter.TGetmcoinlist")
}

func init() {
	proto.RegisterFile("usercenter/t_getmcoinlist.proto", fileDescriptor_t_getmcoinlist_67be56ed7cd22a4b)
}

var fileDescriptor_t_getmcoinlist_67be56ed7cd22a4b = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0xd5, 0x96, 0x26, 0xed, 0x01, 0x12, 0xb2, 0x10, 0xb2, 0xba, 0x50, 0x75, 0xea, 0x94,
	0x48, 0xc0, 0x04, 0x1b, 0x0b, 0x7b, 0xd5, 0x89, 0x05, 0x39, 0xc9, 0x61, 0x22, 0xd9, 0x71, 0xc8,
	0x5d, 0x10, 0x3c, 0x1f, 0x2f, 0x86, 0x72, 0x26, 0x6a, 0x99, 0xd8, 0xee, 0xf7, 0xe5, 0xbb, 0x3f,
	0x31, 0x5c, 0xf7, 0x84, 0x5d, 0x89, 0x0d, 0x63, 0x97, 0xf3, 0x8b, 0x45, 0xf6, 0x65, 0xa8, 0x1b,
	0x57, 0x13, 0x67, 0x6d, 0x17, 0x38, 0x28, 0x38, 0x08, 0xab, 0x07, 0x5b, 0xf3, 0x5b, 0x5f, 0x64,
	0x65, 0xf0, 0xb9, 0x0d, 0xce, 0x34, 0x36, 0x17, 0xa9, 0xe8, 0x5f, 0xf3, 0x96, 0xbf, 0x5a, 0xa4,
	0x9c, 0x6b, 0x8f, 0xc4, 0xc6, 0xb7, 0x87, 0x2a, 0x0e, 0xda, 0x7c, 0x4f, 0xe1, 0x7c, 0xff, 0x74,
	0xb4, 0x40, 0x5d, 0xc0, 0x8c, 0xf0, 0x5d, 0x4f, 0xd6, 0x93, 0xed, 0x72, 0x37, 0x94, 0xea, 0x0a,
	0x12, 0x4f, 0x35, 0x55, 0x8d, 0x9e, 0x0a, 0xfc, 0x4d, 0xea, 0x0e, 0x52, 0x8b, 0x5c, 0x19, 0x46,
	0x3d, 0x5b, 0x4f, 0xb6, 0xa7, 0x37, 0xab, 0xcc, 0x86, 0x60, 0x1d, 0x66, 0xe3, 0xfe, 0x6c, 0x3f,
	0xae, 0xdb, 0x8d, 0xaa, 0xba, 0x07, 0xc0, 0xcf, 0xb6, 0xee, 0x50, 0x1a, 0x4f, 0xfe, 0x6d, 0x3c,
	0xb2, 0x87, 0x4b, 0x88, 0x0d, 0xf7, 0xa4, 0xe7, 0xf1, 0x92, 0x98, 0xd4, 0x25, 0xcc, 0x3f, 0x8c,
	0xeb, 0x51, 0x27, 0x82, 0x63, 0x50, 0x1a, 0xd2, 0xc2, 0x38, 0xd3, 0x94, 0xa8, 0x53, 0xe1, 0x63,
	0x1c, 0xfc, 0x0a, 0xa9, 0xec, 0xf4, 0x22, 0xfa, 0x12, 0x06, 0x2a, 0xcf, 0xa0, 0x97, 0x91, 0x4a,
	0x50, 0x1b, 0x38, 0x93, 0x62, 0x1c, 0x05, 0xf2, 0xf1, 0x0f, 0x7b, 0x5c, 0x3c, 0x27, 0x3e, 0x54,
	0xe8, 0xa8, 0x48, 0xe4, 0x0f, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x58, 0x90, 0xee,
	0xc2, 0x01, 0x00, 0x00,
}
