// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_iplist.proto

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

type TIpList struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip                   string               `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	SourceId             string               `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Descs                string               `protobuf:"bytes,5,opt,name=descs,proto3" json:"descs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TIpList) Reset()         { *m = TIpList{} }
func (m *TIpList) String() string { return proto.CompactTextString(m) }
func (*TIpList) ProtoMessage()    {}
func (*TIpList) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_iplist_a773c2cdfee7404c, []int{0}
}
func (m *TIpList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TIpList.Unmarshal(m, b)
}
func (m *TIpList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TIpList.Marshal(b, m, deterministic)
}
func (dst *TIpList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TIpList.Merge(dst, src)
}
func (m *TIpList) XXX_Size() int {
	return xxx_messageInfo_TIpList.Size(m)
}
func (m *TIpList) XXX_DiscardUnknown() {
	xxx_messageInfo_TIpList.DiscardUnknown(m)
}

var xxx_messageInfo_TIpList proto.InternalMessageInfo

func (m *TIpList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TIpList) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *TIpList) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *TIpList) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TIpList) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func init() {
	proto.RegisterType((*TIpList)(nil), "usercenter.TIpList")
}

func init() { proto.RegisterFile("usercenter/t_iplist.proto", fileDescriptor_t_iplist_a773c2cdfee7404c) }

var fileDescriptor_t_iplist_a773c2cdfee7404c = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x40, 0x4b, 0xeb, 0x4a, 0x0c, 0x16, 0x83, 0x29, 0x4b, 0xc5, 0xd4, 0xc9, 0x96,
	0x60, 0xec, 0x86, 0x58, 0x2a, 0x31, 0x55, 0x9d, 0x58, 0x22, 0xc7, 0x3e, 0x8c, 0xa5, 0xa4, 0x3e,
	0xf9, 0x2e, 0x03, 0x3f, 0x85, 0x7f, 0x8b, 0x64, 0x53, 0x65, 0xbb, 0xef, 0xdd, 0xbb, 0x77, 0x4f,
	0x3c, 0x4e, 0x04, 0xd9, 0xc1, 0x85, 0x21, 0x1b, 0xee, 0x22, 0x0e, 0x91, 0x58, 0x63, 0x4e, 0x9c,
	0xa4, 0x98, 0x57, 0xdb, 0x43, 0x88, 0xfc, 0x3d, 0xf5, 0xda, 0xa5, 0xd1, 0x84, 0x34, 0xd8, 0x4b,
	0x30, 0xc5, 0xd4, 0x4f, 0x5f, 0x06, 0xf9, 0x07, 0x81, 0x0c, 0xc7, 0x11, 0x88, 0xed, 0x88, 0xf3,
	0x54, 0x83, 0x9e, 0x7f, 0x1b, 0x71, 0x77, 0x3e, 0xe2, 0x47, 0x24, 0x96, 0xf7, 0xa2, 0x8d, 0x5e,
	0x35, 0xbb, 0x66, 0xbf, 0x3e, 0xb5, 0xd1, 0x17, 0x46, 0xd5, 0xfe, 0x33, 0xca, 0x27, 0xb1, 0xa6,
	0x34, 0x65, 0x07, 0x5d, 0xf4, 0xea, 0xa6, 0xc8, 0xab, 0x2a, 0x1c, 0xbd, 0x3c, 0x88, 0x8d, 0xcb,
	0x60, 0x19, 0x3a, 0x6f, 0x19, 0xd4, 0xed, 0xae, 0xd9, 0x6f, 0x5e, 0xb6, 0x3a, 0xa4, 0x14, 0x06,
	0xd0, 0xd7, 0x42, 0xfa, 0x7c, 0xfd, 0x7f, 0x12, 0xd5, 0xfe, 0x6e, 0x19, 0xe4, 0x83, 0x58, 0x78,
	0x20, 0x47, 0x6a, 0x51, 0x52, 0x2b, 0xbc, 0xad, 0x3e, 0x97, 0x63, 0xf2, 0x30, 0x50, 0xbf, 0x2c,
	0xf7, 0xaf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x07, 0x79, 0x7b, 0x12, 0x01, 0x00, 0x00,
}
