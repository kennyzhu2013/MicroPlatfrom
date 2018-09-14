// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_operatelog.proto

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

type TOperateLog struct {
	Operateid            string               `protobuf:"bytes,1,opt,name=operateid,proto3" json:"operateid,omitempty"`
	Operatedesc          string               `protobuf:"bytes,2,opt,name=operatedesc,proto3" json:"operatedesc,omitempty"`
	Operateresult        string               `protobuf:"bytes,3,opt,name=operateresult,proto3" json:"operateresult,omitempty"`
	Successcounts        string               `protobuf:"bytes,4,opt,name=successcounts,proto3" json:"successcounts,omitempty"`
	Operatetime          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=operatetime,proto3" json:"operatetime,omitempty"`
	Lasttime             string               `protobuf:"bytes,6,opt,name=lasttime,proto3" json:"lasttime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TOperateLog) Reset()         { *m = TOperateLog{} }
func (m *TOperateLog) String() string { return proto.CompactTextString(m) }
func (*TOperateLog) ProtoMessage()    {}
func (*TOperateLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_operatelog_7b2faa22c761e898, []int{0}
}
func (m *TOperateLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TOperateLog.Unmarshal(m, b)
}
func (m *TOperateLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TOperateLog.Marshal(b, m, deterministic)
}
func (dst *TOperateLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TOperateLog.Merge(dst, src)
}
func (m *TOperateLog) XXX_Size() int {
	return xxx_messageInfo_TOperateLog.Size(m)
}
func (m *TOperateLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TOperateLog.DiscardUnknown(m)
}

var xxx_messageInfo_TOperateLog proto.InternalMessageInfo

func (m *TOperateLog) GetOperateid() string {
	if m != nil {
		return m.Operateid
	}
	return ""
}

func (m *TOperateLog) GetOperatedesc() string {
	if m != nil {
		return m.Operatedesc
	}
	return ""
}

func (m *TOperateLog) GetOperateresult() string {
	if m != nil {
		return m.Operateresult
	}
	return ""
}

func (m *TOperateLog) GetSuccesscounts() string {
	if m != nil {
		return m.Successcounts
	}
	return ""
}

func (m *TOperateLog) GetOperatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Operatetime
	}
	return nil
}

func (m *TOperateLog) GetLasttime() string {
	if m != nil {
		return m.Lasttime
	}
	return ""
}

func init() {
	proto.RegisterType((*TOperateLog)(nil), "usercenter.TOperateLog")
}

func init() {
	proto.RegisterFile("usercenter/t_operatelog.proto", fileDescriptor_t_operatelog_7b2faa22c761e898)
}

var fileDescriptor_t_operatelog_7b2faa22c761e898 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0xc5, 0x71, 0xff, 0x98, 0x44, 0xa6, 0x8b, 0x26, 0x63, 0x5a, 0x30, 0xa5, 0x43, 0x26, 0x09,
	0xda, 0xb1, 0x9d, 0x3a, 0x17, 0x0a, 0x21, 0x53, 0x97, 0x62, 0xcb, 0x57, 0xd5, 0x20, 0xfb, 0x84,
	0xee, 0x34, 0xf4, 0x73, 0xf7, 0x0b, 0x84, 0xc8, 0x4e, 0x9c, 0x6c, 0x7a, 0xef, 0x7e, 0x7a, 0x77,
	0x4f, 0x3c, 0x44, 0x82, 0x60, 0x60, 0x64, 0x08, 0x9a, 0xbf, 0xd1, 0x43, 0x68, 0x18, 0x1c, 0x5a,
	0xe5, 0x03, 0x32, 0x4a, 0xb1, 0x8c, 0xab, 0x57, 0xdb, 0xf3, 0x6f, 0x6c, 0x95, 0xc1, 0x41, 0x5b,
	0x74, 0xcd, 0x68, 0x75, 0x82, 0xda, 0xf8, 0xa3, 0x3d, 0xff, 0x79, 0x20, 0xcd, 0xfd, 0x00, 0xc4,
	0xcd, 0xe0, 0x97, 0xd7, 0x14, 0xf4, 0xf8, 0x9f, 0x89, 0x62, 0xf7, 0x39, 0xc5, 0x7f, 0xa0, 0x95,
	0xf7, 0x62, 0x3d, 0x2f, 0xeb, 0xbb, 0x32, 0xab, 0xb3, 0xcd, 0x7a, 0xbb, 0x18, 0xb2, 0x16, 0xc5,
	0x2c, 0x3a, 0x20, 0x53, 0x5e, 0xa5, 0xf9, 0xb9, 0x25, 0x9f, 0xc4, 0xdd, 0x2c, 0x03, 0x50, 0x74,
	0x5c, 0x5e, 0x27, 0xe6, 0xd2, 0x3c, 0x50, 0x14, 0x8d, 0x01, 0x22, 0x83, 0x71, 0x64, 0x2a, 0x6f,
	0x26, 0xea, 0xc2, 0x94, 0x6f, 0xa7, 0x6d, 0x87, 0xab, 0xcb, 0xdb, 0x3a, 0xdb, 0x14, 0xcf, 0x95,
	0xb2, 0x88, 0xd6, 0x81, 0x3a, 0x76, 0x54, 0xbb, 0x63, 0xa5, 0xed, 0x39, 0x2e, 0x2b, 0xb1, 0x72,
	0x0d, 0x71, 0xfa, 0x9a, 0xa7, 0xf8, 0x93, 0x7e, 0x5f, 0x7d, 0xe5, 0x03, 0x76, 0xe0, 0xa8, 0xcd,
	0x53, 0xcc, 0xcb, 0x3e, 0x00, 0x00, 0xff, 0xff, 0x72, 0xe0, 0x82, 0x51, 0x70, 0x01, 0x00, 0x00,
}
