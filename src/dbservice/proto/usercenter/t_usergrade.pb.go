// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_usergrade.proto

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

type TUserGrade struct {
	Msisdn               string               `protobuf:"bytes,1,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Usergrade            string               `protobuf:"bytes,2,opt,name=usergrade,proto3" json:"usergrade,omitempty"`
	Vipstatus            string               `protobuf:"bytes,3,opt,name=vipstatus,proto3" json:"vipstatus,omitempty"`
	Viptype              string               `protobuf:"bytes,4,opt,name=viptype,proto3" json:"viptype,omitempty"`
	Vipstartdate         *timestamp.Timestamp `protobuf:"bytes,5,opt,name=vipstartdate,proto3" json:"vipstartdate,omitempty"`
	Vipenddate           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=vipenddate,proto3" json:"vipenddate,omitempty"`
	Changetime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=changetime,proto3" json:"changetime,omitempty"`
	Partitionid          string               `protobuf:"bytes,8,opt,name=partitionid,proto3" json:"partitionid,omitempty"`
	Gradedesc            string               `protobuf:"bytes,9,opt,name=gradedesc,proto3" json:"gradedesc,omitempty"`
	Status1              string               `protobuf:"bytes,10,opt,name=status1,proto3" json:"status1,omitempty"`
	Status2              string               `protobuf:"bytes,11,opt,name=status2,proto3" json:"status2,omitempty"`
	Validtime1           *timestamp.Timestamp `protobuf:"bytes,12,opt,name=validtime1,proto3" json:"validtime1,omitempty"`
	Validtime2           *timestamp.Timestamp `protobuf:"bytes,13,opt,name=validtime2,proto3" json:"validtime2,omitempty"`
	Mbiquto              string               `protobuf:"bytes,14,opt,name=mbiquto,proto3" json:"mbiquto,omitempty"`
	Unsubsixtime         *timestamp.Timestamp `protobuf:"bytes,15,opt,name=unsubsixtime,proto3" json:"unsubsixtime,omitempty"`
	Mcoin                string               `protobuf:"bytes,16,opt,name=mcoin,proto3" json:"mcoin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TUserGrade) Reset()         { *m = TUserGrade{} }
func (m *TUserGrade) String() string { return proto.CompactTextString(m) }
func (*TUserGrade) ProtoMessage()    {}
func (*TUserGrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_usergrade_0be7cb7f5da9e38a, []int{0}
}
func (m *TUserGrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TUserGrade.Unmarshal(m, b)
}
func (m *TUserGrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TUserGrade.Marshal(b, m, deterministic)
}
func (dst *TUserGrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TUserGrade.Merge(dst, src)
}
func (m *TUserGrade) XXX_Size() int {
	return xxx_messageInfo_TUserGrade.Size(m)
}
func (m *TUserGrade) XXX_DiscardUnknown() {
	xxx_messageInfo_TUserGrade.DiscardUnknown(m)
}

var xxx_messageInfo_TUserGrade proto.InternalMessageInfo

func (m *TUserGrade) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TUserGrade) GetUsergrade() string {
	if m != nil {
		return m.Usergrade
	}
	return ""
}

func (m *TUserGrade) GetVipstatus() string {
	if m != nil {
		return m.Vipstatus
	}
	return ""
}

func (m *TUserGrade) GetViptype() string {
	if m != nil {
		return m.Viptype
	}
	return ""
}

func (m *TUserGrade) GetVipstartdate() *timestamp.Timestamp {
	if m != nil {
		return m.Vipstartdate
	}
	return nil
}

func (m *TUserGrade) GetVipenddate() *timestamp.Timestamp {
	if m != nil {
		return m.Vipenddate
	}
	return nil
}

func (m *TUserGrade) GetChangetime() *timestamp.Timestamp {
	if m != nil {
		return m.Changetime
	}
	return nil
}

func (m *TUserGrade) GetPartitionid() string {
	if m != nil {
		return m.Partitionid
	}
	return ""
}

func (m *TUserGrade) GetGradedesc() string {
	if m != nil {
		return m.Gradedesc
	}
	return ""
}

func (m *TUserGrade) GetStatus1() string {
	if m != nil {
		return m.Status1
	}
	return ""
}

func (m *TUserGrade) GetStatus2() string {
	if m != nil {
		return m.Status2
	}
	return ""
}

func (m *TUserGrade) GetValidtime1() *timestamp.Timestamp {
	if m != nil {
		return m.Validtime1
	}
	return nil
}

func (m *TUserGrade) GetValidtime2() *timestamp.Timestamp {
	if m != nil {
		return m.Validtime2
	}
	return nil
}

func (m *TUserGrade) GetMbiquto() string {
	if m != nil {
		return m.Mbiquto
	}
	return ""
}

func (m *TUserGrade) GetUnsubsixtime() *timestamp.Timestamp {
	if m != nil {
		return m.Unsubsixtime
	}
	return nil
}

func (m *TUserGrade) GetMcoin() string {
	if m != nil {
		return m.Mcoin
	}
	return ""
}

func init() {
	proto.RegisterType((*TUserGrade)(nil), "usercenter.TUserGrade")
}

func init() {
	proto.RegisterFile("usercenter/t_usergrade.proto", fileDescriptor_t_usergrade_0be7cb7f5da9e38a)
}

var fileDescriptor_t_usergrade_0be7cb7f5da9e38a = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3d, 0x4f, 0xeb, 0x30,
	0x14, 0x55, 0xdf, 0x7b, 0x4d, 0x5b, 0xb7, 0x0f, 0x90, 0x85, 0xd0, 0x55, 0xd5, 0xa1, 0x62, 0xea,
	0x94, 0xa8, 0x61, 0x03, 0x89, 0x81, 0x85, 0xbd, 0x2a, 0x0b, 0x0b, 0x72, 0x62, 0x37, 0xb5, 0x94,
	0xd8, 0xc1, 0x1f, 0x15, 0xfc, 0x09, 0x7e, 0x33, 0xb2, 0x9d, 0x26, 0xe9, 0x54, 0xd8, 0x7c, 0xee,
	0x39, 0xe7, 0xea, 0xdc, 0x63, 0xb4, 0xb0, 0x9a, 0xa9, 0x9c, 0x09, 0xc3, 0x54, 0x62, 0xde, 0x1c,
	0x28, 0x14, 0xa1, 0x2c, 0xae, 0x95, 0x34, 0x12, 0xa3, 0x8e, 0x9d, 0x3f, 0x14, 0xdc, 0xec, 0x6d,
	0x16, 0xe7, 0xb2, 0x4a, 0x0a, 0x59, 0x12, 0x51, 0x24, 0x5e, 0x94, 0xd9, 0x5d, 0x52, 0x9b, 0xcf,
	0x9a, 0xe9, 0xc4, 0xf0, 0x8a, 0x69, 0x43, 0xaa, 0xba, 0x7b, 0x85, 0x45, 0xb7, 0x5f, 0x43, 0x84,
	0xb6, 0x2f, 0x9a, 0xa9, 0x67, 0xb7, 0x1d, 0xdf, 0xa0, 0xa8, 0xd2, 0x5c, 0x53, 0x01, 0x83, 0xe5,
	0x60, 0x35, 0xd9, 0x34, 0x08, 0x2f, 0xd0, 0xa4, 0x8d, 0x00, 0x7f, 0x3c, 0xd5, 0x0d, 0x1c, 0x7b,
	0xe0, 0xb5, 0x36, 0xc4, 0x58, 0x0d, 0x7f, 0x03, 0xdb, 0x0e, 0x30, 0xa0, 0xd1, 0x81, 0xfb, 0x20,
	0xf0, 0xcf, 0x73, 0x47, 0x88, 0x1f, 0xd1, 0x2c, 0xc8, 0x94, 0xa1, 0xc4, 0x30, 0x18, 0x2e, 0x07,
	0xab, 0x69, 0x3a, 0x8f, 0x0b, 0x29, 0x8b, 0xb2, 0x39, 0x35, 0xb3, 0xbb, 0x78, 0x7b, 0x0c, 0xbd,
	0x39, 0xd1, 0xe3, 0x7b, 0x84, 0x0e, 0xbc, 0x66, 0x82, 0x7a, 0x77, 0x74, 0xd6, 0xdd, 0x53, 0x3b,
	0x6f, 0xbe, 0x27, 0xa2, 0x60, 0xae, 0x11, 0x18, 0x9d, 0xf7, 0x76, 0x6a, 0xbc, 0x44, 0xd3, 0x9a,
	0x28, 0xc3, 0x0d, 0x97, 0x82, 0x53, 0x18, 0xfb, 0xab, 0xfa, 0x23, 0xd7, 0x88, 0xaf, 0x86, 0x32,
	0x9d, 0xc3, 0x24, 0x34, 0xd2, 0x0e, 0x5c, 0x23, 0xa1, 0x9b, 0x35, 0xa0, 0xd0, 0x48, 0x03, 0x3b,
	0x26, 0x85, 0x69, 0x9f, 0x49, 0xfd, 0xad, 0xa4, 0xe4, 0xd4, 0x05, 0x58, 0xc3, 0xec, 0x07, 0xb7,
	0xb6, 0xea, 0x13, 0x6f, 0x0a, 0xff, 0x7f, 0xe1, 0x4d, 0x5d, 0xa2, 0x2a, 0xe3, 0xef, 0xd6, 0x48,
	0xb8, 0x08, 0x89, 0x1a, 0xe8, 0x7e, 0xcf, 0x0a, 0x6d, 0x33, 0xcd, 0x3f, 0x7c, 0x87, 0x97, 0xe7,
	0x7f, 0xaf, 0xaf, 0xc7, 0xd7, 0x68, 0x58, 0xe5, 0x92, 0x0b, 0xb8, 0xf2, 0x7b, 0x03, 0x78, 0x1a,
	0xbf, 0x46, 0x95, 0xa4, 0xac, 0xd4, 0x59, 0xe4, 0x37, 0xdc, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x61, 0x01, 0x73, 0x79, 0x0a, 0x03, 0x00, 0x00,
}
