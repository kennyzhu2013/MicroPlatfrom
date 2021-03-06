// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_memberunsignactionlimit.proto

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

type TMemberUnsignActionLimit struct {
	Mobile               string               `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AccessDate           string               `protobuf:"bytes,2,opt,name=access_date,json=accessDate,proto3" json:"access_date,omitempty"`
	SvcType              string               `protobuf:"bytes,3,opt,name=svc_type,json=svcType,proto3" json:"svc_type,omitempty"`
	AccessSeq            string               `protobuf:"bytes,4,opt,name=access_seq,json=accessSeq,proto3" json:"access_seq,omitempty"`
	SignDate             string               `protobuf:"bytes,5,opt,name=sign_date,json=signDate,proto3" json:"sign_date,omitempty"`
	ActionType           string               `protobuf:"bytes,6,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	CurrentDay           string               `protobuf:"bytes,8,opt,name=current_day,json=currentDay,proto3" json:"current_day,omitempty"`
	CurrentWeek          string               `protobuf:"bytes,9,opt,name=current_week,json=currentWeek,proto3" json:"current_week,omitempty"`
	CurrentMonth         string               `protobuf:"bytes,10,opt,name=current_month,json=currentMonth,proto3" json:"current_month,omitempty"`
	CurrentDaySeq        string               `protobuf:"bytes,11,opt,name=current_day_seq,json=currentDaySeq,proto3" json:"current_day_seq,omitempty"`
	CurrentWeekSeq       string               `protobuf:"bytes,12,opt,name=current_week_seq,json=currentWeekSeq,proto3" json:"current_week_seq,omitempty"`
	CurrentMonthSeq      string               `protobuf:"bytes,13,opt,name=current_month_seq,json=currentMonthSeq,proto3" json:"current_month_seq,omitempty"`
	PartitionId          string               `protobuf:"bytes,14,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberUnsignActionLimit) Reset()         { *m = TMemberUnsignActionLimit{} }
func (m *TMemberUnsignActionLimit) String() string { return proto.CompactTextString(m) }
func (*TMemberUnsignActionLimit) ProtoMessage()    {}
func (*TMemberUnsignActionLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_memberunsignactionlimit_00250584e3ae0757, []int{0}
}
func (m *TMemberUnsignActionLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberUnsignActionLimit.Unmarshal(m, b)
}
func (m *TMemberUnsignActionLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberUnsignActionLimit.Marshal(b, m, deterministic)
}
func (dst *TMemberUnsignActionLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberUnsignActionLimit.Merge(dst, src)
}
func (m *TMemberUnsignActionLimit) XXX_Size() int {
	return xxx_messageInfo_TMemberUnsignActionLimit.Size(m)
}
func (m *TMemberUnsignActionLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberUnsignActionLimit.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberUnsignActionLimit proto.InternalMessageInfo

func (m *TMemberUnsignActionLimit) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetAccessDate() string {
	if m != nil {
		return m.AccessDate
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetSvcType() string {
	if m != nil {
		return m.SvcType
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetAccessSeq() string {
	if m != nil {
		return m.AccessSeq
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetSignDate() string {
	if m != nil {
		return m.SignDate
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberUnsignActionLimit) GetCurrentDay() string {
	if m != nil {
		return m.CurrentDay
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCurrentWeek() string {
	if m != nil {
		return m.CurrentWeek
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCurrentMonth() string {
	if m != nil {
		return m.CurrentMonth
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCurrentDaySeq() string {
	if m != nil {
		return m.CurrentDaySeq
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCurrentWeekSeq() string {
	if m != nil {
		return m.CurrentWeekSeq
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetCurrentMonthSeq() string {
	if m != nil {
		return m.CurrentMonthSeq
	}
	return ""
}

func (m *TMemberUnsignActionLimit) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberUnsignActionLimit)(nil), "usercenter.TMemberUnsignActionLimit")
}

func init() {
	proto.RegisterFile("usercenter/t_memberunsignactionlimit.proto", fileDescriptor_t_memberunsignactionlimit_00250584e3ae0757)
}

var fileDescriptor_t_memberunsignactionlimit_00250584e3ae0757 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x55, 0x18, 0x5d, 0xfb, 0xd2, 0x6e, 0xe0, 0x03, 0x32, 0x43, 0x88, 0x0d, 0x24, 0x54,
	0xed, 0x90, 0x48, 0x70, 0xdc, 0x09, 0xb4, 0x0b, 0x12, 0xbb, 0x8c, 0x22, 0x24, 0x2e, 0x91, 0xe3,
	0x3c, 0x32, 0x6b, 0x71, 0x9c, 0xd9, 0x4e, 0x51, 0xbe, 0x06, 0x9f, 0x18, 0xf9, 0xb9, 0xa6, 0xd9,
	0xad, 0xfd, 0xbf, 0xdf, 0xfb, 0xff, 0x9c, 0x07, 0x97, 0x83, 0x43, 0x2b, 0xb1, 0xf3, 0x68, 0x0b,
	0x5f, 0x6a, 0xd4, 0x15, 0xda, 0xa1, 0x73, 0xaa, 0xe9, 0x84, 0xf4, 0xca, 0x74, 0xad, 0xd2, 0xca,
	0xe7, 0xbd, 0x35, 0xde, 0x30, 0x38, 0xb0, 0x67, 0x57, 0x8d, 0xf2, 0x77, 0x43, 0x95, 0x4b, 0xa3,
	0x8b, 0xc6, 0xb4, 0xa2, 0x6b, 0x0a, 0x82, 0xaa, 0xe1, 0x77, 0xd1, 0xfb, 0xb1, 0x47, 0x57, 0x78,
	0xa5, 0xd1, 0x79, 0xa1, 0xfb, 0xc3, 0xaf, 0x58, 0xf4, 0xee, 0xef, 0x11, 0xf0, 0xed, 0x0d, 0xb9,
	0x7e, 0x90, 0xeb, 0x33, 0xb9, 0xbe, 0x05, 0x17, 0x7b, 0x09, 0x73, 0x6d, 0x2a, 0xd5, 0x22, 0x9f,
	0x9d, 0xcf, 0x36, 0xcb, 0xdb, 0xfd, 0x3f, 0xf6, 0x16, 0x32, 0x21, 0x25, 0x3a, 0x57, 0xd6, 0xc2,
	0x23, 0x7f, 0x42, 0x43, 0x88, 0xd1, 0xb5, 0xf0, 0xc8, 0x5e, 0xc1, 0xc2, 0xed, 0x64, 0x19, 0xe4,
	0xfc, 0x29, 0x4d, 0x8f, 0xdd, 0x4e, 0x6e, 0xc7, 0x1e, 0xd9, 0x1b, 0xd8, 0x83, 0xa5, 0xc3, 0x07,
	0x7e, 0x44, 0xc3, 0x65, 0x4c, 0xbe, 0xe3, 0x03, 0x7b, 0x0d, 0xcb, 0xf0, 0x8a, 0x58, 0xfc, 0x8c,
	0xa6, 0x8b, 0x10, 0x50, 0x2d, 0x79, 0xc3, 0xf3, 0x62, 0xf3, 0x3c, 0x79, 0x43, 0x44, 0xe5, 0x57,
	0x90, 0x49, 0x8b, 0xc2, 0x63, 0xdc, 0x3f, 0x3e, 0x9f, 0x6d, 0xb2, 0x8f, 0x67, 0x79, 0x63, 0x4c,
	0xd3, 0x62, 0x9e, 0xae, 0x92, 0x6f, 0xd3, 0x11, 0x6e, 0x21, 0xe2, 0xa9, 0x5d, 0x0e, 0xd6, 0x62,
	0xe7, 0xcb, 0x5a, 0x8c, 0x7c, 0x11, 0xdb, 0xf7, 0xd1, 0xb5, 0x18, 0xd9, 0x05, 0xac, 0x12, 0xf0,
	0x07, 0xf1, 0x9e, 0x2f, 0x89, 0x48, 0x4b, 0x3f, 0x11, 0xef, 0xd9, 0x7b, 0x58, 0x27, 0x44, 0x9b,
	0xce, 0xdf, 0x71, 0x20, 0x26, 0xed, 0xdd, 0x84, 0x8c, 0x7d, 0x80, 0xd3, 0x89, 0x88, 0xee, 0x90,
	0x11, 0xb6, 0x3e, 0xc8, 0xc2, 0x2d, 0x36, 0xf0, 0x7c, 0xea, 0x23, 0x70, 0x45, 0xe0, 0xc9, 0xc4,
	0x19, 0xc8, 0x4b, 0x78, 0xf1, 0x48, 0x4b, 0xe8, 0x9a, 0xd0, 0xd3, 0xa9, 0x3a, 0xb0, 0x17, 0xb0,
	0xea, 0x85, 0xf5, 0x8a, 0xee, 0xa8, 0x6a, 0x7e, 0x12, 0xbf, 0xe2, 0x7f, 0xf6, 0xb5, 0xfe, 0xb2,
	0xf8, 0x35, 0xd7, 0xa6, 0xc6, 0xd6, 0x55, 0x73, 0xba, 0xd9, 0xa7, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x94, 0x0f, 0xe8, 0x9c, 0x02, 0x00, 0x00,
}
