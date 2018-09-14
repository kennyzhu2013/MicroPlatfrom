// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_membergrade.proto

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

type TMemberGrade struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinVitality          string   `protobuf:"bytes,3,opt,name=min_vitality,json=minVitality,proto3" json:"min_vitality,omitempty"`
	Pid                  string   `protobuf:"bytes,4,opt,name=pid,proto3" json:"pid,omitempty"`
	TitleId              string   `protobuf:"bytes,5,opt,name=title_id,json=titleId,proto3" json:"title_id,omitempty"`
	MaxVitality          string   `protobuf:"bytes,6,opt,name=max_vitality,json=maxVitality,proto3" json:"max_vitality,omitempty"`
	AwardPoolId          string   `protobuf:"bytes,7,opt,name=award_pool_id,json=awardPoolId,proto3" json:"award_pool_id,omitempty"`
	MonthAwardDays       string   `protobuf:"bytes,8,opt,name=month_award_days,json=monthAwardDays,proto3" json:"month_award_days,omitempty"`
	MonthAwardPoolId     string   `protobuf:"bytes,9,opt,name=month_award_pool_id,json=monthAwardPoolId,proto3" json:"month_award_pool_id,omitempty"`
	Descs                string   `protobuf:"bytes,10,opt,name=descs,proto3" json:"descs,omitempty"`
	LogoUrl              string   `protobuf:"bytes,11,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	GradeId              string   `protobuf:"bytes,12,opt,name=grade_id,json=gradeId,proto3" json:"grade_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TMemberGrade) Reset()         { *m = TMemberGrade{} }
func (m *TMemberGrade) String() string { return proto.CompactTextString(m) }
func (*TMemberGrade) ProtoMessage()    {}
func (*TMemberGrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_membergrade_ec606a44f09e3c07, []int{0}
}
func (m *TMemberGrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberGrade.Unmarshal(m, b)
}
func (m *TMemberGrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberGrade.Marshal(b, m, deterministic)
}
func (dst *TMemberGrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberGrade.Merge(dst, src)
}
func (m *TMemberGrade) XXX_Size() int {
	return xxx_messageInfo_TMemberGrade.Size(m)
}
func (m *TMemberGrade) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberGrade.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberGrade proto.InternalMessageInfo

func (m *TMemberGrade) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMemberGrade) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TMemberGrade) GetMinVitality() string {
	if m != nil {
		return m.MinVitality
	}
	return ""
}

func (m *TMemberGrade) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *TMemberGrade) GetTitleId() string {
	if m != nil {
		return m.TitleId
	}
	return ""
}

func (m *TMemberGrade) GetMaxVitality() string {
	if m != nil {
		return m.MaxVitality
	}
	return ""
}

func (m *TMemberGrade) GetAwardPoolId() string {
	if m != nil {
		return m.AwardPoolId
	}
	return ""
}

func (m *TMemberGrade) GetMonthAwardDays() string {
	if m != nil {
		return m.MonthAwardDays
	}
	return ""
}

func (m *TMemberGrade) GetMonthAwardPoolId() string {
	if m != nil {
		return m.MonthAwardPoolId
	}
	return ""
}

func (m *TMemberGrade) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func (m *TMemberGrade) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *TMemberGrade) GetGradeId() string {
	if m != nil {
		return m.GradeId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberGrade)(nil), "usercenter.TMemberGrade")
}

func init() {
	proto.RegisterFile("usercenter/t_membergrade.proto", fileDescriptor_t_membergrade_ec606a44f09e3c07)
}

var fileDescriptor_t_membergrade_ec606a44f09e3c07 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xd5, 0xf4, 0xff, 0xb6, 0x54, 0x95, 0x61, 0x30, 0x0c, 0x08, 0x3a, 0x75, 0xa1, 0x19,
	0x18, 0x99, 0x40, 0x48, 0xa8, 0x03, 0x12, 0x42, 0xc0, 0xc0, 0x12, 0x39, 0xb5, 0x49, 0x2d, 0xd9,
	0x71, 0x64, 0x3b, 0xd0, 0xbc, 0x34, 0xcf, 0x80, 0x7c, 0x1d, 0x08, 0xdb, 0x3d, 0xe7, 0x7c, 0xfa,
	0x24, 0xcb, 0x70, 0x5e, 0x3b, 0x61, 0x77, 0xa2, 0xf4, 0xc2, 0xa6, 0x3e, 0xd3, 0x42, 0xe7, 0xc2,
	0x16, 0x96, 0x71, 0xb1, 0xa9, 0xac, 0xf1, 0x86, 0x40, 0xb7, 0x9f, 0xdd, 0x14, 0xd2, 0xef, 0xeb,
	0x7c, 0xb3, 0x33, 0x3a, 0x2d, 0x8c, 0x62, 0x65, 0x91, 0x22, 0x94, 0xd7, 0x1f, 0x69, 0xe5, 0x9b,
	0x4a, 0xb8, 0xd4, 0x4b, 0x2d, 0x9c, 0x67, 0xba, 0xea, 0xae, 0x28, 0x5a, 0x7d, 0x27, 0x30, 0x7f,
	0x79, 0x44, 0xff, 0x43, 0xf0, 0x93, 0x05, 0x24, 0x92, 0xd3, 0xde, 0x45, 0x6f, 0x3d, 0x7d, 0x4e,
	0x24, 0x27, 0x04, 0x06, 0x25, 0xd3, 0x82, 0x26, 0xd8, 0xe0, 0x4d, 0x2e, 0x61, 0xae, 0x65, 0x99,
	0x7d, 0x4a, 0xcf, 0x94, 0xf4, 0x0d, 0xed, 0xe3, 0x36, 0xd3, 0xb2, 0x7c, 0x6b, 0x2b, 0xb2, 0x84,
	0x7e, 0x25, 0x39, 0x1d, 0xe0, 0x12, 0x4e, 0x72, 0x0a, 0x13, 0x2f, 0xbd, 0x12, 0x99, 0xe4, 0x74,
	0x88, 0xf5, 0x18, 0xf3, 0x96, 0xa3, 0x8f, 0x1d, 0x3a, 0xdf, 0xa8, 0xf5, 0xb1, 0xc3, 0x9f, 0x6f,
	0x05, 0x47, 0xec, 0x8b, 0x59, 0x9e, 0x55, 0xc6, 0xa8, 0xa0, 0x18, 0x47, 0x06, 0xcb, 0x27, 0x63,
	0xd4, 0x96, 0x93, 0x35, 0x2c, 0xb5, 0x29, 0xfd, 0x3e, 0x8b, 0x24, 0x67, 0x8d, 0xa3, 0x13, 0xc4,
	0x16, 0xd8, 0xdf, 0x86, 0xfa, 0x9e, 0x35, 0x8e, 0x5c, 0xc1, 0xf1, 0x7f, 0xf2, 0xd7, 0x39, 0x45,
	0x78, 0xd9, 0xc1, 0xad, 0xf8, 0x04, 0x86, 0x5c, 0xb8, 0x9d, 0xa3, 0x80, 0x40, 0x0c, 0xe1, 0x41,
	0xca, 0x14, 0x26, 0xab, 0xad, 0xa2, 0xb3, 0xf8, 0xa0, 0x90, 0x5f, 0xad, 0x0a, 0x13, 0xfe, 0x56,
	0x90, 0xce, 0xe3, 0x84, 0x79, 0xcb, 0xef, 0x26, 0xef, 0x23, 0x6d, 0xb8, 0x50, 0x2e, 0x1f, 0xe1,
	0x0f, 0x5c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x30, 0xb5, 0x28, 0xec, 0x01, 0x00, 0x00,
}
