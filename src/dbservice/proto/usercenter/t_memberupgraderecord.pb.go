// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_memberupgraderecord.proto

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

type TMemberUpgradeRecord struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	PreGrade             string               `protobuf:"bytes,3,opt,name=pre_grade,json=preGrade,proto3" json:"pre_grade,omitempty"`
	ChangedGrade         string               `protobuf:"bytes,4,opt,name=changed_grade,json=changedGrade,proto3" json:"changed_grade,omitempty"`
	Reason               string               `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Type                 string               `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Mcoin                string               `protobuf:"bytes,8,opt,name=mcoin,proto3" json:"mcoin,omitempty"`
	Status               string               `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	FetchDate            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=fetch_date,json=fetchDate,proto3" json:"fetch_date,omitempty"`
	ExpiredDate          *timestamp.Timestamp `protobuf:"bytes,11,opt,name=expired_date,json=expiredDate,proto3" json:"expired_date,omitempty"`
	PartitionId          string               `protobuf:"bytes,12,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberUpgradeRecord) Reset()         { *m = TMemberUpgradeRecord{} }
func (m *TMemberUpgradeRecord) String() string { return proto.CompactTextString(m) }
func (*TMemberUpgradeRecord) ProtoMessage()    {}
func (*TMemberUpgradeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_memberupgraderecord_b0958c9f121deacd, []int{0}
}
func (m *TMemberUpgradeRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberUpgradeRecord.Unmarshal(m, b)
}
func (m *TMemberUpgradeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberUpgradeRecord.Marshal(b, m, deterministic)
}
func (dst *TMemberUpgradeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberUpgradeRecord.Merge(dst, src)
}
func (m *TMemberUpgradeRecord) XXX_Size() int {
	return xxx_messageInfo_TMemberUpgradeRecord.Size(m)
}
func (m *TMemberUpgradeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberUpgradeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberUpgradeRecord proto.InternalMessageInfo

func (m *TMemberUpgradeRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetPreGrade() string {
	if m != nil {
		return m.PreGrade
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetChangedGrade() string {
	if m != nil {
		return m.ChangedGrade
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberUpgradeRecord) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetMcoin() string {
	if m != nil {
		return m.Mcoin
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TMemberUpgradeRecord) GetFetchDate() *timestamp.Timestamp {
	if m != nil {
		return m.FetchDate
	}
	return nil
}

func (m *TMemberUpgradeRecord) GetExpiredDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredDate
	}
	return nil
}

func (m *TMemberUpgradeRecord) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberUpgradeRecord)(nil), "usercenter.TMemberUpgradeRecord")
}

func init() {
	proto.RegisterFile("usercenter/t_memberupgraderecord.proto", fileDescriptor_t_memberupgraderecord_b0958c9f121deacd)
}

var fileDescriptor_t_memberupgraderecord_b0958c9f121deacd = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x86, 0xc9, 0x97, 0x37, 0x1e, 0x7b, 0xf7, 0x20, 0x42, 0x10, 0xd9, 0x4b, 0x76, 0x17, 0x96,
	0x9c, 0x6c, 0x68, 0x4f, 0x25, 0xf4, 0x52, 0x0a, 0xa5, 0x87, 0x5e, 0x42, 0x7a, 0xe9, 0xc5, 0xc8,
	0xd6, 0xc4, 0x11, 0x58, 0x96, 0x90, 0x65, 0x68, 0xff, 0x57, 0x7f, 0x60, 0xc9, 0xc8, 0x69, 0x8e,
	0xb9, 0xe9, 0x1d, 0x3f, 0xf3, 0x3e, 0x1e, 0xf8, 0xdf, 0x77, 0xe8, 0x2a, 0x6c, 0x3d, 0xba, 0xdc,
	0x17, 0x1a, 0x75, 0x89, 0xae, 0xb7, 0xb5, 0x13, 0x12, 0x1d, 0x56, 0xc6, 0xc9, 0xcc, 0x3a, 0xe3,
	0x0d, 0x83, 0x0b, 0xb7, 0xda, 0xd6, 0xca, 0x1f, 0xfb, 0x32, 0xab, 0x8c, 0xce, 0x6b, 0xd3, 0x88,
	0xb6, 0xce, 0x09, 0x2a, 0xfb, 0x43, 0x6e, 0xfd, 0x87, 0xc5, 0x2e, 0xf7, 0x4a, 0x63, 0xe7, 0x85,
	0xb6, 0x97, 0x57, 0x28, 0xfa, 0xfb, 0x39, 0x81, 0xc5, 0xfe, 0x85, 0x3c, 0xaf, 0xc1, 0xb3, 0x23,
	0x0f, 0xfb, 0x05, 0x63, 0x25, 0xf9, 0x68, 0x3d, 0xda, 0xc4, 0xbb, 0xb1, 0x92, 0x6c, 0x09, 0x91,
	0x36, 0xa5, 0x6a, 0x90, 0x8f, 0x69, 0x36, 0x24, 0xf6, 0x1b, 0x62, 0xeb, 0xb0, 0xa0, 0x55, 0x3e,
	0xa1, 0x4f, 0x73, 0xeb, 0xf0, 0xe9, 0x94, 0xd9, 0x3f, 0xf8, 0x59, 0x1d, 0x45, 0x5b, 0xa3, 0x1c,
	0x80, 0x29, 0x01, 0xe9, 0x30, 0x0c, 0xd0, 0x12, 0x22, 0x87, 0xa2, 0x33, 0x2d, 0x9f, 0x85, 0xe6,
	0x90, 0xd8, 0x16, 0x92, 0xca, 0xa1, 0xf0, 0x58, 0x48, 0xe1, 0x91, 0x47, 0xeb, 0xd1, 0x26, 0xb9,
	0x59, 0x65, 0xb5, 0x31, 0x75, 0x83, 0xd9, 0xf9, 0xc4, 0x6c, 0x7f, 0xbe, 0x68, 0x07, 0x01, 0x7f,
	0x14, 0x1e, 0x19, 0x83, 0xe9, 0xe9, 0x74, 0xfe, 0x83, 0x2a, 0xe9, 0xcd, 0x16, 0x30, 0xd3, 0x95,
	0x51, 0x2d, 0x9f, 0xd3, 0x30, 0x84, 0x93, 0xbe, 0xf3, 0xc2, 0xf7, 0x1d, 0x8f, 0x83, 0x3e, 0x24,
	0x76, 0x07, 0x70, 0x40, 0x5f, 0x1d, 0x83, 0x1d, 0xae, 0xda, 0x63, 0xa2, 0x49, 0x7e, 0x0f, 0x29,
	0xbe, 0x5b, 0xe5, 0x50, 0x86, 0xe5, 0xe4, 0xea, 0x72, 0x32, 0xf0, 0xb4, 0xfe, 0x07, 0x52, 0x2b,
	0x9c, 0x57, 0x5e, 0x99, 0xb6, 0x50, 0x92, 0xa7, 0xf4, 0x5f, 0xc9, 0xf7, 0xec, 0x59, 0x3e, 0xcc,
	0xdf, 0x22, 0x6d, 0x24, 0x36, 0x5d, 0x19, 0x51, 0xdb, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x94, 0x5d, 0x49, 0x27, 0x3a, 0x02, 0x00, 0x00,
}