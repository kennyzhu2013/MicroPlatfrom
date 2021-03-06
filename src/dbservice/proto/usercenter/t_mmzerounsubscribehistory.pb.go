// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mmzerounsubscribehistory.proto

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

type TMmZeroUnsubscribeHistory struct {
	OrderNum             string               `protobuf:"bytes,1,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ReqCsspDate          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=req_cssp_date,json=reqCsspDate,proto3" json:"req_cssp_date,omitempty"`
	CallBackDate         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=call_back_date,json=callBackDate,proto3" json:"call_back_date,omitempty"`
	TransactionId        string               `protobuf:"bytes,7,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	RightsId             string               `protobuf:"bytes,8,opt,name=rights_id,json=rightsId,proto3" json:"rights_id,omitempty"`
	RightsType           string               `protobuf:"bytes,9,opt,name=rights_type,json=rightsType,proto3" json:"rights_type,omitempty"`
	PayCode              string               `protobuf:"bytes,10,opt,name=pay_code,json=payCode,proto3" json:"pay_code,omitempty"`
	PartitionId          string               `protobuf:"bytes,11,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMmZeroUnsubscribeHistory) Reset()         { *m = TMmZeroUnsubscribeHistory{} }
func (m *TMmZeroUnsubscribeHistory) String() string { return proto.CompactTextString(m) }
func (*TMmZeroUnsubscribeHistory) ProtoMessage()    {}
func (*TMmZeroUnsubscribeHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mmzerounsubscribehistory_231b26c0f55410a8, []int{0}
}
func (m *TMmZeroUnsubscribeHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMmZeroUnsubscribeHistory.Unmarshal(m, b)
}
func (m *TMmZeroUnsubscribeHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMmZeroUnsubscribeHistory.Marshal(b, m, deterministic)
}
func (dst *TMmZeroUnsubscribeHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMmZeroUnsubscribeHistory.Merge(dst, src)
}
func (m *TMmZeroUnsubscribeHistory) XXX_Size() int {
	return xxx_messageInfo_TMmZeroUnsubscribeHistory.Size(m)
}
func (m *TMmZeroUnsubscribeHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TMmZeroUnsubscribeHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TMmZeroUnsubscribeHistory proto.InternalMessageInfo

func (m *TMmZeroUnsubscribeHistory) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMmZeroUnsubscribeHistory) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetReqCsspDate() *timestamp.Timestamp {
	if m != nil {
		return m.ReqCsspDate
	}
	return nil
}

func (m *TMmZeroUnsubscribeHistory) GetCallBackDate() *timestamp.Timestamp {
	if m != nil {
		return m.CallBackDate
	}
	return nil
}

func (m *TMmZeroUnsubscribeHistory) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetRightsId() string {
	if m != nil {
		return m.RightsId
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetRightsType() string {
	if m != nil {
		return m.RightsType
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetPayCode() string {
	if m != nil {
		return m.PayCode
	}
	return ""
}

func (m *TMmZeroUnsubscribeHistory) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMmZeroUnsubscribeHistory)(nil), "usercenter.TMmZeroUnsubscribeHistory")
}

func init() {
	proto.RegisterFile("usercenter/t_mmzerounsubscribehistory.proto", fileDescriptor_t_mmzerounsubscribehistory_231b26c0f55410a8)
}

var fileDescriptor_t_mmzerounsubscribehistory_231b26c0f55410a8 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0xaf, 0xd3, 0x30,
	0x10, 0xc6, 0x55, 0x1e, 0xe4, 0xa5, 0xce, 0x6b, 0x07, 0x0f, 0x28, 0x2d, 0x03, 0x05, 0x09, 0xa9,
	0x12, 0x52, 0x22, 0xc1, 0x58, 0x09, 0xa1, 0x96, 0x81, 0x0e, 0x30, 0x54, 0x65, 0xe9, 0x12, 0x39,
	0xf6, 0x91, 0x5a, 0x8d, 0x63, 0xd7, 0x3e, 0x0f, 0xe1, 0x4f, 0xe4, 0xaf, 0x42, 0xb1, 0x5b, 0xfa,
	0xb6, 0x6e, 0xbe, 0xef, 0xfb, 0xfc, 0xbb, 0xf3, 0x99, 0x7c, 0xf4, 0x0e, 0x2c, 0x87, 0x0e, 0xc1,
	0x96, 0x58, 0x29, 0xf5, 0x07, 0xac, 0xf6, 0x9d, 0xf3, 0xb5, 0xe3, 0x56, 0xd6, 0x70, 0x94, 0x0e,
	0xb5, 0xed, 0x0b, 0x63, 0x35, 0x6a, 0x4a, 0x6e, 0xe1, 0xf9, 0xaa, 0x91, 0x78, 0xf4, 0x75, 0xc1,
	0xb5, 0x2a, 0x1b, 0xdd, 0xb2, 0xae, 0x29, 0x43, 0xa8, 0xf6, 0xbf, 0x4b, 0x83, 0xbd, 0x01, 0x57,
	0xa2, 0x54, 0xe0, 0x90, 0x29, 0x73, 0x3b, 0x45, 0xd0, 0xfb, 0xbf, 0x0f, 0x64, 0xb6, 0xff, 0xa1,
	0x0e, 0x60, 0xf5, 0xaf, 0x5b, 0xb3, 0xef, 0xb1, 0x19, 0x7d, 0x43, 0xc6, 0xda, 0x0a, 0xb0, 0x55,
	0xe7, 0x55, 0x3e, 0x5a, 0x8c, 0x96, 0xe3, 0x5d, 0x1a, 0x84, 0x9f, 0x5e, 0xd1, 0xd7, 0x24, 0x51,
	0xba, 0x96, 0x2d, 0xe4, 0x2f, 0x82, 0x73, 0xa9, 0xe8, 0x8a, 0x64, 0xdc, 0x02, 0x43, 0xa8, 0x04,
	0x43, 0xc8, 0x1f, 0x16, 0xa3, 0x65, 0xf6, 0x69, 0x5e, 0x34, 0x5a, 0x37, 0x2d, 0x14, 0xd7, 0xd1,
	0x8a, 0xfd, 0x75, 0x92, 0x1d, 0x89, 0xf1, 0x6f, 0x0c, 0x61, 0x80, 0x3a, 0x64, 0xe8, 0x5d, 0xfe,
	0x32, 0x42, 0x63, 0x45, 0xbf, 0x90, 0x89, 0x85, 0x73, 0xc5, 0x9d, 0x33, 0x11, 0xfb, 0xea, 0x2e,
	0x36, 0xb3, 0x70, 0xde, 0x38, 0x67, 0x02, 0xf7, 0x2b, 0x99, 0x72, 0xd6, 0xb6, 0x55, 0xcd, 0xf8,
	0x29, 0x02, 0x92, 0xbb, 0x80, 0xa7, 0xe1, 0xc6, 0x9a, 0xf1, 0x53, 0x20, 0x7c, 0x20, 0x53, 0xb4,
	0xac, 0x73, 0x8c, 0xa3, 0xd4, 0x5d, 0x25, 0x45, 0xfe, 0x18, 0x26, 0x9c, 0x3c, 0x53, 0xb7, 0x62,
	0x58, 0x99, 0x95, 0xcd, 0x11, 0xdd, 0x90, 0x48, 0xe3, 0xca, 0xa2, 0xb0, 0x15, 0xf4, 0x2d, 0xc9,
	0x2e, 0xe6, 0xf0, 0x2f, 0xf9, 0x38, 0xd8, 0x24, 0x4a, 0xfb, 0xde, 0x00, 0x9d, 0x91, 0xd4, 0xb0,
	0xbe, 0xe2, 0x5a, 0x40, 0x4e, 0x82, 0xfb, 0x68, 0x58, 0xbf, 0xd1, 0x02, 0xe8, 0x3b, 0xf2, 0x64,
	0x98, 0x45, 0x79, 0xed, 0x9e, 0x05, 0x3b, 0xfb, 0xaf, 0x6d, 0xc5, 0x3a, 0x3d, 0x24, 0x4a, 0x0b,
	0x68, 0x5d, 0x9d, 0x84, 0xe7, 0x7c, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x78, 0x37, 0xa2, 0x0d,
	0x55, 0x02, 0x00, 0x00,
}
