// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_memberscoredetail.proto

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

type TMemberScoreDetail struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Action               string               `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	ActionName           string               `protobuf:"bytes,4,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	ScoreNum             string               `protobuf:"bytes,5,opt,name=score_num,json=scoreNum,proto3" json:"score_num,omitempty"`
	Source               string               `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	ActionDate           string               `protobuf:"bytes,7,opt,name=action_date,json=actionDate,proto3" json:"action_date,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Descs                string               `protobuf:"bytes,9,opt,name=descs,proto3" json:"descs,omitempty"`
	PartitionId          string               `protobuf:"bytes,10,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	GoodsId              string               `protobuf:"bytes,11,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	GoodsNum             string               `protobuf:"bytes,12,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
	ActionCode           string               `protobuf:"bytes,13,opt,name=action_code,json=actionCode,proto3" json:"action_code,omitempty"`
	ActivityId           string               `protobuf:"bytes,14,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberScoreDetail) Reset()         { *m = TMemberScoreDetail{} }
func (m *TMemberScoreDetail) String() string { return proto.CompactTextString(m) }
func (*TMemberScoreDetail) ProtoMessage()    {}
func (*TMemberScoreDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_memberscoredetail_e72a953c92727ac8, []int{0}
}
func (m *TMemberScoreDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberScoreDetail.Unmarshal(m, b)
}
func (m *TMemberScoreDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberScoreDetail.Marshal(b, m, deterministic)
}
func (dst *TMemberScoreDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberScoreDetail.Merge(dst, src)
}
func (m *TMemberScoreDetail) XXX_Size() int {
	return xxx_messageInfo_TMemberScoreDetail.Size(m)
}
func (m *TMemberScoreDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberScoreDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberScoreDetail proto.InternalMessageInfo

func (m *TMemberScoreDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMemberScoreDetail) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberScoreDetail) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TMemberScoreDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *TMemberScoreDetail) GetScoreNum() string {
	if m != nil {
		return m.ScoreNum
	}
	return ""
}

func (m *TMemberScoreDetail) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *TMemberScoreDetail) GetActionDate() string {
	if m != nil {
		return m.ActionDate
	}
	return ""
}

func (m *TMemberScoreDetail) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberScoreDetail) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func (m *TMemberScoreDetail) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func (m *TMemberScoreDetail) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *TMemberScoreDetail) GetGoodsNum() string {
	if m != nil {
		return m.GoodsNum
	}
	return ""
}

func (m *TMemberScoreDetail) GetActionCode() string {
	if m != nil {
		return m.ActionCode
	}
	return ""
}

func (m *TMemberScoreDetail) GetActivityId() string {
	if m != nil {
		return m.ActivityId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberScoreDetail)(nil), "usercenter.TMemberScoreDetail")
}

func init() {
	proto.RegisterFile("usercenter/t_memberscoredetail.proto", fileDescriptor_t_memberscoredetail_e72a953c92727ac8)
}

var fileDescriptor_t_memberscoredetail_e72a953c92727ac8 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x6f, 0xe2, 0x30,
	0x18, 0xc6, 0x05, 0x1c, 0x21, 0x38, 0x1c, 0x83, 0x75, 0x3a, 0xf9, 0xb8, 0xa1, 0xb4, 0xea, 0xc0,
	0x94, 0x48, 0xed, 0xc8, 0xd6, 0xb2, 0x30, 0x94, 0x81, 0x32, 0x75, 0x41, 0x8e, 0xfd, 0x36, 0xb5,
	0x14, 0xc7, 0x91, 0xed, 0x54, 0xe2, 0xb3, 0xf5, 0xcb, 0x55, 0x7e, 0x5d, 0x08, 0x5b, 0x9e, 0x3f,
	0x79, 0x5e, 0xfd, 0x4c, 0xee, 0x3b, 0x07, 0x56, 0x40, 0xe3, 0xc1, 0x16, 0xfe, 0xa8, 0x41, 0x97,
	0x60, 0x9d, 0x30, 0x16, 0x24, 0x78, 0xae, 0xea, 0xbc, 0xb5, 0xc6, 0x1b, 0x4a, 0xfa, 0xd6, 0x62,
	0x5d, 0x29, 0xff, 0xd1, 0x95, 0xb9, 0x30, 0xba, 0xa8, 0x4c, 0xcd, 0x9b, 0xaa, 0xc0, 0x52, 0xd9,
	0xbd, 0x17, 0xad, 0x3f, 0xb5, 0xe0, 0x0a, 0xaf, 0x34, 0x38, 0xcf, 0x75, 0xdb, 0x7f, 0xc5, 0xa1,
	0xbb, 0xaf, 0x11, 0xa1, 0x87, 0x17, 0xbc, 0xf2, 0x1a, 0xae, 0x6c, 0xf0, 0x0a, 0x9d, 0x93, 0xa1,
	0x92, 0x6c, 0xb0, 0x1c, 0xac, 0xa6, 0xfb, 0xa1, 0x92, 0xf4, 0x2f, 0x49, 0xb4, 0x29, 0x55, 0x0d,
	0x6c, 0x88, 0xde, 0x8f, 0x0a, 0x3e, 0x17, 0x5e, 0x99, 0x86, 0x8d, 0xa2, 0x1f, 0x15, 0xbd, 0x21,
	0x59, 0xfc, 0x3a, 0x36, 0x5c, 0x03, 0xfb, 0x85, 0x21, 0x89, 0xd6, 0x8e, 0x6b, 0xa0, 0xff, 0xc9,
	0x14, 0xa9, 0x8e, 0x4d, 0xa7, 0xd9, 0x18, 0xe3, 0x14, 0x8d, 0x5d, 0xa7, 0xc3, 0xaa, 0x33, 0x9d,
	0x15, 0xc0, 0x92, 0xb8, 0x1a, 0xd5, 0xd5, 0xaa, 0xe4, 0x1e, 0xd8, 0xe4, 0x7a, 0x75, 0xc3, 0x3d,
	0xd0, 0x35, 0xc9, 0x84, 0x05, 0xee, 0x21, 0x16, 0xd2, 0xe5, 0x60, 0x95, 0x3d, 0x2c, 0xf2, 0xca,
	0x98, 0xaa, 0x86, 0xfc, 0xfc, 0x2a, 0xf9, 0xe1, 0xfc, 0x08, 0x7b, 0x12, 0xeb, 0xf8, 0xf3, 0x1f,
	0x32, 0x96, 0xe0, 0x84, 0x63, 0x53, 0xdc, 0x8d, 0x82, 0xde, 0x92, 0x59, 0xcb, 0xad, 0x57, 0x78,
	0x56, 0x49, 0x46, 0x30, 0xcc, 0x2e, 0xde, 0x56, 0xd2, 0x7f, 0x24, 0xad, 0x8c, 0x91, 0x2e, 0xc4,
	0x19, 0xc6, 0x13, 0xd4, 0x5b, 0x19, 0x30, 0x63, 0x14, 0x30, 0x67, 0x11, 0x13, 0x8d, 0x80, 0xd9,
	0xe3, 0x08, 0x23, 0x81, 0xfd, 0xbe, 0xc6, 0x79, 0x36, 0xf2, 0xc2, 0xfb, 0xa9, 0xfc, 0x29, 0x6c,
	0xcf, 0xfb, 0x42, 0xb0, 0xb6, 0xf2, 0x29, 0x7d, 0x4b, 0xb4, 0x91, 0x50, 0xbb, 0x32, 0x41, 0xb8,
	0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x51, 0x6c, 0x79, 0x3f, 0x02, 0x00, 0x00,
}