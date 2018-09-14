// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_memberactionlimit.proto

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

type TMemberActionLimit struct {
	Mobile               string               `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AccessDate           string               `protobuf:"bytes,2,opt,name=access_date,json=accessDate,proto3" json:"access_date,omitempty"`
	SvcType              string               `protobuf:"bytes,3,opt,name=svc_type,json=svcType,proto3" json:"svc_type,omitempty"`
	AccessSeq            string               `protobuf:"bytes,4,opt,name=access_seq,json=accessSeq,proto3" json:"access_seq,omitempty"`
	SignDate             string               `protobuf:"bytes,5,opt,name=sign_date,json=signDate,proto3" json:"sign_date,omitempty"`
	ActionType           string               `protobuf:"bytes,6,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	PartitionId          string               `protobuf:"bytes,8,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberActionLimit) Reset()         { *m = TMemberActionLimit{} }
func (m *TMemberActionLimit) String() string { return proto.CompactTextString(m) }
func (*TMemberActionLimit) ProtoMessage()    {}
func (*TMemberActionLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_memberactionlimit_c4a9c7264127e749, []int{0}
}
func (m *TMemberActionLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberActionLimit.Unmarshal(m, b)
}
func (m *TMemberActionLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberActionLimit.Marshal(b, m, deterministic)
}
func (dst *TMemberActionLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberActionLimit.Merge(dst, src)
}
func (m *TMemberActionLimit) XXX_Size() int {
	return xxx_messageInfo_TMemberActionLimit.Size(m)
}
func (m *TMemberActionLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberActionLimit.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberActionLimit proto.InternalMessageInfo

func (m *TMemberActionLimit) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberActionLimit) GetAccessDate() string {
	if m != nil {
		return m.AccessDate
	}
	return ""
}

func (m *TMemberActionLimit) GetSvcType() string {
	if m != nil {
		return m.SvcType
	}
	return ""
}

func (m *TMemberActionLimit) GetAccessSeq() string {
	if m != nil {
		return m.AccessSeq
	}
	return ""
}

func (m *TMemberActionLimit) GetSignDate() string {
	if m != nil {
		return m.SignDate
	}
	return ""
}

func (m *TMemberActionLimit) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *TMemberActionLimit) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberActionLimit) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberActionLimit)(nil), "usercenter.TMemberActionLimit")
}

func init() {
	proto.RegisterFile("usercenter/t_memberactionlimit.proto", fileDescriptor_t_memberactionlimit_c4a9c7264127e749)
}

var fileDescriptor_t_memberactionlimit_c4a9c7264127e749 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xd5, 0xfe, 0x3f, 0x69, 0x7a, 0xc3, 0xe4, 0x01, 0x85, 0x22, 0x44, 0x41, 0x0c, 0x9d,
	0x12, 0x09, 0xc6, 0x4e, 0x20, 0x16, 0x24, 0x58, 0x4a, 0x27, 0x96, 0xc8, 0x71, 0x2e, 0xc1, 0x52,
	0x1c, 0xbb, 0xf6, 0x6d, 0xa5, 0x3e, 0x10, 0xef, 0x89, 0x72, 0xdd, 0xd2, 0x2d, 0x39, 0xe7, 0xe4,
	0xfb, 0x62, 0xc3, 0xfd, 0x36, 0xa0, 0x57, 0xd8, 0x13, 0xfa, 0x92, 0x2a, 0x83, 0xa6, 0x46, 0x2f,
	0x15, 0x69, 0xdb, 0x77, 0xda, 0x68, 0x2a, 0x9c, 0xb7, 0x64, 0x05, 0x9c, 0x56, 0xb3, 0x65, 0xab,
	0xe9, 0x7b, 0x5b, 0x17, 0xca, 0x9a, 0xb2, 0xb5, 0x9d, 0xec, 0xdb, 0x92, 0x47, 0xf5, 0xf6, 0xab,
	0x74, 0xb4, 0x77, 0x18, 0x4a, 0xd2, 0x06, 0x03, 0x49, 0xe3, 0x4e, 0x4f, 0x11, 0x74, 0xf7, 0x33,
	0x06, 0xb1, 0x7e, 0x67, 0xcb, 0x13, 0x5b, 0xde, 0x06, 0x8b, 0xb8, 0x80, 0xc4, 0xd8, 0x5a, 0x77,
	0x98, 0x8f, 0xe6, 0xa3, 0xc5, 0x74, 0x75, 0x78, 0x13, 0x37, 0x90, 0x49, 0xa5, 0x30, 0x84, 0xaa,
	0x91, 0x84, 0xf9, 0x98, 0x4b, 0x88, 0xd1, 0x8b, 0x24, 0x14, 0x97, 0x90, 0x86, 0x9d, 0xaa, 0x06,
	0x6d, 0xfe, 0x8f, 0xdb, 0x49, 0xd8, 0xa9, 0xf5, 0xde, 0xa1, 0xb8, 0x86, 0xc3, 0xb0, 0x0a, 0xb8,
	0xc9, 0xff, 0x73, 0x39, 0x8d, 0xc9, 0x07, 0x6e, 0xc4, 0x15, 0x4c, 0x83, 0x6e, 0xfb, 0x08, 0x3e,
	0xe3, 0x36, 0x1d, 0x02, 0xc6, 0xb2, 0x77, 0xf8, 0xbd, 0x48, 0x4e, 0x8e, 0xde, 0x21, 0x62, 0xf8,
	0x12, 0x32, 0xe5, 0x51, 0x12, 0xc6, 0xef, 0x27, 0xf3, 0xd1, 0x22, 0x7b, 0x98, 0x15, 0xad, 0xb5,
	0x6d, 0x87, 0xc5, 0xf1, 0x3e, 0x8a, 0xf5, 0xf1, 0xf8, 0x2b, 0x88, 0x73, 0xa6, 0xdf, 0xc2, 0xb9,
	0x93, 0x9e, 0x34, 0x0b, 0x74, 0x93, 0xa7, 0x8c, 0xcf, 0xfe, 0xb2, 0xd7, 0xe6, 0x39, 0xfd, 0x4c,
	0x8c, 0x6d, 0xb0, 0x0b, 0x75, 0xc2, 0xb0, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0xe5,
	0x0d, 0xf5, 0xa9, 0x01, 0x00, 0x00,
}
