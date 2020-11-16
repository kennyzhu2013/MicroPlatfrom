// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mcouponsinfohis.proto

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

type TMcouponsinfoHis struct {
	Recordseq            string               `protobuf:"bytes,1,opt,name=recordseq,proto3" json:"recordseq,omitempty"`
	Msisdn               string               `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Mcouponsvalue        string               `protobuf:"bytes,3,opt,name=mcouponsvalue,proto3" json:"mcouponsvalue,omitempty"`
	Modtype              string               `protobuf:"bytes,4,opt,name=modtype,proto3" json:"modtype,omitempty"`
	Actiontime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=actiontime,proto3" json:"actiontime,omitempty"`
	Gifttype             string               `protobuf:"bytes,6,opt,name=gifttype,proto3" json:"gifttype,omitempty"`
	Appname              string               `protobuf:"bytes,7,opt,name=appname,proto3" json:"appname,omitempty"`
	Orderid              string               `protobuf:"bytes,8,opt,name=orderid,proto3" json:"orderid,omitempty"`
	Appid                string               `protobuf:"bytes,9,opt,name=appid,proto3" json:"appid,omitempty"`
	Description          string               `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Updatetime           *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	BatchId              string               `protobuf:"bytes,12,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	UserId               string               `protobuf:"bytes,13,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsBatch              string               `protobuf:"bytes,14,opt,name=is_batch,json=isBatch,proto3" json:"is_batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMcouponsinfoHis) Reset()         { *m = TMcouponsinfoHis{} }
func (m *TMcouponsinfoHis) String() string { return proto.CompactTextString(m) }
func (*TMcouponsinfoHis) ProtoMessage()    {}
func (*TMcouponsinfoHis) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mcouponsinfohis_548ed0cff6bbd132, []int{0}
}
func (m *TMcouponsinfoHis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMcouponsinfoHis.Unmarshal(m, b)
}
func (m *TMcouponsinfoHis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMcouponsinfoHis.Marshal(b, m, deterministic)
}
func (dst *TMcouponsinfoHis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMcouponsinfoHis.Merge(dst, src)
}
func (m *TMcouponsinfoHis) XXX_Size() int {
	return xxx_messageInfo_TMcouponsinfoHis.Size(m)
}
func (m *TMcouponsinfoHis) XXX_DiscardUnknown() {
	xxx_messageInfo_TMcouponsinfoHis.DiscardUnknown(m)
}

var xxx_messageInfo_TMcouponsinfoHis proto.InternalMessageInfo

func (m *TMcouponsinfoHis) GetRecordseq() string {
	if m != nil {
		return m.Recordseq
	}
	return ""
}

func (m *TMcouponsinfoHis) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TMcouponsinfoHis) GetMcouponsvalue() string {
	if m != nil {
		return m.Mcouponsvalue
	}
	return ""
}

func (m *TMcouponsinfoHis) GetModtype() string {
	if m != nil {
		return m.Modtype
	}
	return ""
}

func (m *TMcouponsinfoHis) GetActiontime() *timestamp.Timestamp {
	if m != nil {
		return m.Actiontime
	}
	return nil
}

func (m *TMcouponsinfoHis) GetGifttype() string {
	if m != nil {
		return m.Gifttype
	}
	return ""
}

func (m *TMcouponsinfoHis) GetAppname() string {
	if m != nil {
		return m.Appname
	}
	return ""
}

func (m *TMcouponsinfoHis) GetOrderid() string {
	if m != nil {
		return m.Orderid
	}
	return ""
}

func (m *TMcouponsinfoHis) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *TMcouponsinfoHis) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TMcouponsinfoHis) GetUpdatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Updatetime
	}
	return nil
}

func (m *TMcouponsinfoHis) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

func (m *TMcouponsinfoHis) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TMcouponsinfoHis) GetIsBatch() string {
	if m != nil {
		return m.IsBatch
	}
	return ""
}

func init() {
	proto.RegisterType((*TMcouponsinfoHis)(nil), "usercenter.TMcouponsinfoHis")
}

func init() {
	proto.RegisterFile("usercenter/t_mcouponsinfohis.proto", fileDescriptor_t_mcouponsinfohis_548ed0cff6bbd132)
}

var fileDescriptor_t_mcouponsinfohis_548ed0cff6bbd132 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4f, 0xeb, 0x30,
	0x14, 0x85, 0xd5, 0xd7, 0xd7, 0x34, 0xbd, 0x7d, 0x7d, 0x42, 0x16, 0x02, 0x53, 0x31, 0x54, 0x15,
	0x43, 0xa7, 0x44, 0x82, 0x0d, 0xb6, 0x4e, 0x74, 0x60, 0xa9, 0x3a, 0xb1, 0x54, 0x4e, 0xec, 0xa6,
	0x96, 0xe2, 0x5c, 0x63, 0x3b, 0x48, 0xfc, 0x52, 0xfe, 0x0e, 0xb2, 0xdd, 0x90, 0x32, 0xb1, 0xf5,
	0xbb, 0xe7, 0xf8, 0x9e, 0x7b, 0x1a, 0x58, 0xb6, 0x56, 0x98, 0x52, 0x34, 0x4e, 0x98, 0xdc, 0xed,
	0x55, 0x89, 0xad, 0xc6, 0xc6, 0xca, 0xe6, 0x80, 0x47, 0x69, 0x33, 0x6d, 0xd0, 0x21, 0x81, 0xde,
	0x33, 0x7f, 0xaa, 0xa4, 0x3b, 0xb6, 0x45, 0x56, 0xa2, 0xca, 0x2b, 0xac, 0x59, 0x53, 0xe5, 0xc1,
	0x54, 0xb4, 0x87, 0x5c, 0xbb, 0x0f, 0x2d, 0x6c, 0xee, 0xa4, 0x12, 0xd6, 0x31, 0xa5, 0xfb, 0x5f,
	0x71, 0xd1, 0xf2, 0x73, 0x08, 0x17, 0xbb, 0x97, 0xb3, 0x8c, 0x67, 0x69, 0xc9, 0x2d, 0x4c, 0x8c,
	0x28, 0xd1, 0x70, 0x2b, 0xde, 0xe8, 0x60, 0x31, 0x58, 0x4d, 0xb6, 0xfd, 0x80, 0x5c, 0x41, 0xa2,
	0xac, 0xb4, 0xbc, 0xa1, 0x7f, 0x82, 0x74, 0x22, 0x72, 0x07, 0xb3, 0xee, 0xd8, 0x77, 0x56, 0xb7,
	0x82, 0x0e, 0x83, 0xfc, 0x73, 0x48, 0x28, 0x8c, 0x15, 0x72, 0x7f, 0x16, 0xfd, 0x1b, 0xf4, 0x0e,
	0xc9, 0x23, 0x00, 0x2b, 0x9d, 0xc4, 0xc6, 0xdf, 0x48, 0x47, 0x8b, 0xc1, 0x6a, 0x7a, 0x3f, 0xcf,
	0x2a, 0xc4, 0xaa, 0x16, 0x59, 0xd7, 0x28, 0xdb, 0x75, 0x05, 0xb6, 0x67, 0x6e, 0x32, 0x87, 0xb4,
	0x92, 0x07, 0x17, 0xd6, 0x26, 0x61, 0xed, 0x37, 0xfb, 0x44, 0xa6, 0x75, 0xc3, 0x94, 0xa0, 0xe3,
	0x98, 0x78, 0x42, 0xaf, 0xa0, 0xe1, 0xc2, 0x48, 0x4e, 0xd3, 0xa8, 0x9c, 0x90, 0x5c, 0xc2, 0x88,
	0x69, 0x2d, 0x39, 0x9d, 0x84, 0x79, 0x04, 0xb2, 0x80, 0x29, 0x17, 0xb6, 0x34, 0x52, 0xfb, 0x60,
	0x0a, 0x41, 0x3b, 0x1f, 0xf9, 0x0e, 0xad, 0xe6, 0xcc, 0x89, 0xd0, 0x61, 0xfa, 0x7b, 0x87, 0xde,
	0x4d, 0x6e, 0x20, 0x2d, 0x98, 0x2b, 0x8f, 0x7b, 0xc9, 0xe9, 0xbf, 0x78, 0x4e, 0xe0, 0x0d, 0x27,
	0xd7, 0x30, 0xf6, 0x1f, 0xdc, 0x2b, 0xb3, 0xf8, 0x9f, 0x7b, 0xdc, 0x70, 0xff, 0x46, 0xda, 0x7d,
	0xb0, 0xd1, 0xff, 0xf1, 0x8d, 0xb4, 0x6b, 0x8f, 0xeb, 0xf4, 0x35, 0x51, 0xc8, 0x45, 0x6d, 0x8b,
	0x24, 0x04, 0x3f, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x44, 0x19, 0x25, 0x59, 0x02, 0x00,
	0x00,
}