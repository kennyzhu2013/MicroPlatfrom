// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_gradetitle.proto

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

type TGradeTitle struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	CreateUser           string               `protobuf:"bytes,4,opt,name=create_user,json=createUser,proto3" json:"create_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TGradeTitle) Reset()         { *m = TGradeTitle{} }
func (m *TGradeTitle) String() string { return proto.CompactTextString(m) }
func (*TGradeTitle) ProtoMessage()    {}
func (*TGradeTitle) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_gradetitle_386f87f611bd8588, []int{0}
}
func (m *TGradeTitle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TGradeTitle.Unmarshal(m, b)
}
func (m *TGradeTitle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TGradeTitle.Marshal(b, m, deterministic)
}
func (dst *TGradeTitle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TGradeTitle.Merge(dst, src)
}
func (m *TGradeTitle) XXX_Size() int {
	return xxx_messageInfo_TGradeTitle.Size(m)
}
func (m *TGradeTitle) XXX_DiscardUnknown() {
	xxx_messageInfo_TGradeTitle.DiscardUnknown(m)
}

var xxx_messageInfo_TGradeTitle proto.InternalMessageInfo

func (m *TGradeTitle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TGradeTitle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TGradeTitle) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TGradeTitle) GetCreateUser() string {
	if m != nil {
		return m.CreateUser
	}
	return ""
}

func init() {
	proto.RegisterType((*TGradeTitle)(nil), "usercenter.TGradeTitle")
}

func init() {
	proto.RegisterFile("usercenter/t_gradetitle.proto", fileDescriptor_t_gradetitle_386f87f611bd8588)
}

var fileDescriptor_t_gradetitle_386f87f611bd8588 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xc9, 0x7a, 0x1c, 0x9a, 0x05, 0x8b, 0x54, 0xe1, 0x40, 0x3c, 0xac, 0xae, 0x4a, 0x40,
	0xcb, 0xeb, 0x44, 0xb0, 0x3f, 0xd6, 0xc6, 0xe6, 0x48, 0x36, 0xb3, 0x31, 0xb0, 0xd9, 0x84, 0x64,
	0x52, 0xf8, 0x14, 0xbe, 0xb2, 0x6c, 0xd6, 0x75, 0xbb, 0x9f, 0x99, 0x6f, 0xbe, 0x99, 0xa1, 0x0f,
	0x25, 0x43, 0xea, 0x61, 0x42, 0x48, 0x12, 0xaf, 0x36, 0x29, 0x03, 0xe8, 0x70, 0x04, 0x11, 0x53,
	0xc0, 0xc0, 0xe8, 0xd6, 0x3e, 0x9c, 0xad, 0xc3, 0xaf, 0xa2, 0x45, 0x1f, 0xbc, 0xb4, 0x61, 0x54,
	0x93, 0x95, 0x15, 0xd2, 0x65, 0x90, 0x11, 0xbf, 0x23, 0x64, 0x89, 0xce, 0x43, 0x46, 0xe5, 0xe3,
	0x96, 0x16, 0xd1, 0xd3, 0x0f, 0xa1, 0x6d, 0xf7, 0x3e, 0xeb, 0xbb, 0x59, 0xcf, 0xee, 0x69, 0xe3,
	0x0c, 0x27, 0x47, 0x72, 0xba, 0xbb, 0x34, 0xce, 0x30, 0x46, 0x77, 0x93, 0xf2, 0xc0, 0x9b, 0x5a,
	0xa9, 0x99, 0x9d, 0x69, 0xdb, 0x27, 0x50, 0x08, 0x57, 0xa3, 0x10, 0xf8, 0xcd, 0x91, 0x9c, 0xda,
	0xe7, 0x83, 0xb0, 0x21, 0xd8, 0xf5, 0x40, 0x5d, 0x06, 0xd1, 0xad, 0xab, 0x2e, 0x74, 0xc1, 0xdf,
	0x14, 0x02, 0x7b, 0xfc, 0x1f, 0x9e, 0x5f, 0xe0, 0xbb, 0xea, 0xfd, 0x03, 0x3e, 0x32, 0xa4, 0xd7,
	0xdb, 0xcf, 0xbd, 0x0f, 0x06, 0xc6, 0xac, 0xf7, 0x55, 0xf5, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x4d, 0x02, 0xbd, 0x98, 0x0c, 0x01, 0x00, 0x00,
}
