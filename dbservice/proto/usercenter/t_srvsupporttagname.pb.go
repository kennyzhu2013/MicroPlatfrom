// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_srvsupporttagname.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TSrvSupportTagName struct {
	Tagid                string   `protobuf:"bytes,1,opt,name=tagid,proto3" json:"tagid,omitempty"`
	Tagname              string   `protobuf:"bytes,2,opt,name=tagname,proto3" json:"tagname,omitempty"`
	Tagtypename          string   `protobuf:"bytes,3,opt,name=tagtypename,proto3" json:"tagtypename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TSrvSupportTagName) Reset()         { *m = TSrvSupportTagName{} }
func (m *TSrvSupportTagName) String() string { return proto.CompactTextString(m) }
func (*TSrvSupportTagName) ProtoMessage()    {}
func (*TSrvSupportTagName) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_srvsupporttagname_78627f659ce07a32, []int{0}
}
func (m *TSrvSupportTagName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSrvSupportTagName.Unmarshal(m, b)
}
func (m *TSrvSupportTagName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSrvSupportTagName.Marshal(b, m, deterministic)
}
func (dst *TSrvSupportTagName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSrvSupportTagName.Merge(dst, src)
}
func (m *TSrvSupportTagName) XXX_Size() int {
	return xxx_messageInfo_TSrvSupportTagName.Size(m)
}
func (m *TSrvSupportTagName) XXX_DiscardUnknown() {
	xxx_messageInfo_TSrvSupportTagName.DiscardUnknown(m)
}

var xxx_messageInfo_TSrvSupportTagName proto.InternalMessageInfo

func (m *TSrvSupportTagName) GetTagid() string {
	if m != nil {
		return m.Tagid
	}
	return ""
}

func (m *TSrvSupportTagName) GetTagname() string {
	if m != nil {
		return m.Tagname
	}
	return ""
}

func (m *TSrvSupportTagName) GetTagtypename() string {
	if m != nil {
		return m.Tagtypename
	}
	return ""
}

func init() {
	proto.RegisterType((*TSrvSupportTagName)(nil), "usercenter.TSrvSupportTagName")
}

func init() {
	proto.RegisterFile("usercenter/t_srvsupporttagname.proto", fileDescriptor_t_srvsupporttagname_78627f659ce07a32)
}

var fileDescriptor_t_srvsupporttagname_78627f659ce07a32 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xbf, 0x0e, 0x82, 0x30,
	0x10, 0xc6, 0x83, 0x46, 0xd4, 0xba, 0x35, 0x0e, 0xc4, 0x89, 0x18, 0x07, 0x27, 0x3a, 0x38, 0xba,
	0xf9, 0x00, 0x0e, 0xc2, 0xe4, 0x62, 0x0e, 0x38, 0x2a, 0x09, 0xa5, 0x4d, 0x7b, 0x25, 0xf1, 0xed,
	0x8d, 0x45, 0x83, 0xdb, 0x7d, 0x7f, 0x72, 0xbf, 0x8f, 0x1d, 0xbc, 0x43, 0x5b, 0x61, 0x4f, 0x68,
	0x05, 0x3d, 0x9c, 0x1d, 0x9c, 0x37, 0x46, 0x5b, 0x22, 0x90, 0x3d, 0x28, 0xcc, 0x8c, 0xd5, 0xa4,
	0x39, 0x9b, 0x5a, 0xbb, 0xb3, 0x6c, 0xe9, 0xe9, 0xcb, 0xac, 0xd2, 0x4a, 0x48, 0xdd, 0x41, 0x2f,
	0x45, 0x28, 0x95, 0xbe, 0x11, 0x86, 0x5e, 0x06, 0x9d, 0xa0, 0x56, 0xa1, 0x23, 0x50, 0x66, 0xba,
	0xc6, 0x47, 0xfb, 0x86, 0xf1, 0x22, 0xb7, 0x43, 0x3e, 0x42, 0x0a, 0x90, 0x57, 0x50, 0xc8, 0xb7,
	0x6c, 0x41, 0x20, 0xdb, 0x3a, 0x89, 0xd2, 0xe8, 0xb8, 0xbe, 0x8d, 0x82, 0x27, 0x6c, 0xf9, 0x5d,
	0x91, 0xcc, 0x82, 0xff, 0x93, 0x3c, 0x65, 0x1b, 0x02, 0xf9, 0x61, 0x85, 0x74, 0x1e, 0xd2, 0x7f,
	0xeb, 0xb2, 0xba, 0xc7, 0x4a, 0xd7, 0xd8, 0xb9, 0x32, 0x0e, 0xe0, 0xd3, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0xb6, 0x46, 0x87, 0xe9, 0x00, 0x00, 0x00,
}
