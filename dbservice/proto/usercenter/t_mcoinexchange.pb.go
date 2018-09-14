// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mcoinexchange.proto

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

type TMcoinExchange struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AwardTypeId          string               `protobuf:"bytes,3,opt,name=award_type_id,json=awardTypeId,proto3" json:"award_type_id,omitempty"`
	PersonLimitSeq       string               `protobuf:"bytes,4,opt,name=person_limit_seq,json=personLimitSeq,proto3" json:"person_limit_seq,omitempty"`
	ExchangeDate         string               `protobuf:"bytes,5,opt,name=exchange_date,json=exchangeDate,proto3" json:"exchange_date,omitempty"`
	DayLimitSeq          string               `protobuf:"bytes,6,opt,name=day_limit_seq,json=dayLimitSeq,proto3" json:"day_limit_seq,omitempty"`
	GoodsName            string               `protobuf:"bytes,7,opt,name=goods_name,json=goodsName,proto3" json:"goods_name,omitempty"`
	GoodsType            string               `protobuf:"bytes,8,opt,name=goods_type,json=goodsType,proto3" json:"goods_type,omitempty"`
	ExchangeNum          string               `protobuf:"bytes,9,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num,omitempty"`
	EqualMcoin           string               `protobuf:"bytes,10,opt,name=equal_mcoin,json=equalMcoin,proto3" json:"equal_mcoin,omitempty"`
	EqualScore           string               `protobuf:"bytes,11,opt,name=equal_score,json=equalScore,proto3" json:"equal_score,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,12,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Descs                string               `protobuf:"bytes,13,opt,name=descs,proto3" json:"descs,omitempty"`
	PartitionId          string               `protobuf:"bytes,14,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	PropCode             string               `protobuf:"bytes,15,opt,name=prop_code,json=propCode,proto3" json:"prop_code,omitempty"`
	UnionPrice           string               `protobuf:"bytes,16,opt,name=union_price,json=unionPrice,proto3" json:"union_price,omitempty"`
	ExchangeType         string               `protobuf:"bytes,17,opt,name=exchange_type,json=exchangeType,proto3" json:"exchange_type,omitempty"`
	Channel              string               `protobuf:"bytes,18,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMcoinExchange) Reset()         { *m = TMcoinExchange{} }
func (m *TMcoinExchange) String() string { return proto.CompactTextString(m) }
func (*TMcoinExchange) ProtoMessage()    {}
func (*TMcoinExchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mcoinexchange_49c14e6a34deb46d, []int{0}
}
func (m *TMcoinExchange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMcoinExchange.Unmarshal(m, b)
}
func (m *TMcoinExchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMcoinExchange.Marshal(b, m, deterministic)
}
func (dst *TMcoinExchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMcoinExchange.Merge(dst, src)
}
func (m *TMcoinExchange) XXX_Size() int {
	return xxx_messageInfo_TMcoinExchange.Size(m)
}
func (m *TMcoinExchange) XXX_DiscardUnknown() {
	xxx_messageInfo_TMcoinExchange.DiscardUnknown(m)
}

var xxx_messageInfo_TMcoinExchange proto.InternalMessageInfo

func (m *TMcoinExchange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMcoinExchange) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMcoinExchange) GetAwardTypeId() string {
	if m != nil {
		return m.AwardTypeId
	}
	return ""
}

func (m *TMcoinExchange) GetPersonLimitSeq() string {
	if m != nil {
		return m.PersonLimitSeq
	}
	return ""
}

func (m *TMcoinExchange) GetExchangeDate() string {
	if m != nil {
		return m.ExchangeDate
	}
	return ""
}

func (m *TMcoinExchange) GetDayLimitSeq() string {
	if m != nil {
		return m.DayLimitSeq
	}
	return ""
}

func (m *TMcoinExchange) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *TMcoinExchange) GetGoodsType() string {
	if m != nil {
		return m.GoodsType
	}
	return ""
}

func (m *TMcoinExchange) GetExchangeNum() string {
	if m != nil {
		return m.ExchangeNum
	}
	return ""
}

func (m *TMcoinExchange) GetEqualMcoin() string {
	if m != nil {
		return m.EqualMcoin
	}
	return ""
}

func (m *TMcoinExchange) GetEqualScore() string {
	if m != nil {
		return m.EqualScore
	}
	return ""
}

func (m *TMcoinExchange) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMcoinExchange) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func (m *TMcoinExchange) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func (m *TMcoinExchange) GetPropCode() string {
	if m != nil {
		return m.PropCode
	}
	return ""
}

func (m *TMcoinExchange) GetUnionPrice() string {
	if m != nil {
		return m.UnionPrice
	}
	return ""
}

func (m *TMcoinExchange) GetExchangeType() string {
	if m != nil {
		return m.ExchangeType
	}
	return ""
}

func (m *TMcoinExchange) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*TMcoinExchange)(nil), "usercenter.TMcoinExchange")
}

func init() {
	proto.RegisterFile("usercenter/t_mcoinexchange.proto", fileDescriptor_t_mcoinexchange_49c14e6a34deb46d)
}

var fileDescriptor_t_mcoinexchange_49c14e6a34deb46d = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0xd5, 0xd0, 0xa6, 0xc9, 0x6c, 0x13, 0x8a, 0x85, 0x90, 0x55, 0x84, 0x08, 0xe5, 0x92,
	0x53, 0x56, 0x82, 0x63, 0x6f, 0xfc, 0x39, 0x54, 0x82, 0x0a, 0xb5, 0x39, 0x71, 0x59, 0x79, 0xd7,
	0xd3, 0xad, 0xa5, 0xf5, 0x9f, 0xd8, 0x5e, 0x41, 0x3e, 0x30, 0xdf, 0x03, 0x79, 0x9c, 0x4d, 0xd2,
	0xdb, 0xce, 0x6f, 0xde, 0xce, 0x1b, 0xcf, 0x83, 0x45, 0x1f, 0xd0, 0x37, 0x68, 0x22, 0xfa, 0x32,
	0x56, 0xba, 0xb1, 0xca, 0xe0, 0xdf, 0xe6, 0x49, 0x98, 0x16, 0x57, 0xce, 0xdb, 0x68, 0x19, 0x1c,
	0x14, 0x57, 0x37, 0xad, 0x8a, 0x4f, 0x7d, 0xbd, 0x6a, 0xac, 0x2e, 0x5b, 0xdb, 0x09, 0xd3, 0x96,
	0x24, 0xaa, 0xfb, 0xc7, 0xd2, 0xc5, 0xad, 0xc3, 0x50, 0x46, 0xa5, 0x31, 0x44, 0xa1, 0xdd, 0xe1,
	0x2b, 0x0f, 0xba, 0xfe, 0x77, 0x0a, 0xf3, 0xf5, 0xcf, 0xe4, 0xf0, 0x7d, 0xe7, 0xc0, 0xe6, 0x30,
	0x52, 0x92, 0x9f, 0x2c, 0x4e, 0x96, 0xd3, 0xfb, 0x91, 0x92, 0xec, 0x0d, 0x8c, 0xb5, 0xad, 0x55,
	0x87, 0x7c, 0x44, 0x6c, 0x57, 0xb1, 0x6b, 0x98, 0x89, 0x3f, 0xc2, 0xcb, 0x2a, 0x79, 0x54, 0x4a,
	0xf2, 0x17, 0xd4, 0x2e, 0x08, 0xae, 0xb7, 0x0e, 0x6f, 0x25, 0x5b, 0xc2, 0xa5, 0x43, 0x1f, 0xac,
	0xa9, 0x3a, 0xa5, 0x55, 0xac, 0x02, 0x6e, 0xf8, 0x29, 0xc9, 0xe6, 0x99, 0xff, 0x48, 0xf8, 0x01,
	0x37, 0xec, 0x23, 0xcc, 0x86, 0x37, 0x56, 0x52, 0x44, 0xe4, 0x67, 0x24, 0xbb, 0x18, 0xe0, 0x37,
	0x11, 0xc9, 0x52, 0x8a, 0xed, 0xd1, 0xac, 0x71, 0xb6, 0x94, 0x62, 0xbb, 0x1f, 0xf4, 0x0e, 0xa0,
	0xb5, 0x56, 0x86, 0xca, 0x08, 0x8d, 0xfc, 0x9c, 0x04, 0x53, 0x22, 0x77, 0x42, 0xe3, 0xa1, 0x9d,
	0xb6, 0xe6, 0x93, 0xa3, 0x76, 0x5a, 0x99, 0x7d, 0x80, 0xbd, 0x63, 0x65, 0x7a, 0xcd, 0xa7, 0xd9,
	0x60, 0x60, 0x77, 0xbd, 0x66, 0xef, 0xa1, 0xc0, 0x4d, 0x2f, 0xba, 0x1c, 0x0c, 0x07, 0x52, 0x00,
	0x21, 0x3a, 0xe4, 0x41, 0x10, 0x1a, 0xeb, 0x91, 0x17, 0x47, 0x82, 0x87, 0x44, 0xd8, 0x0d, 0x14,
	0x8d, 0x47, 0x11, 0x77, 0x2f, 0xbd, 0x58, 0x9c, 0x2c, 0x8b, 0x4f, 0x57, 0xab, 0xd6, 0xda, 0xb6,
	0xdb, 0x25, 0x5c, 0xf7, 0x8f, 0xab, 0xf5, 0x90, 0xd5, 0x3d, 0x64, 0x39, 0xdd, 0xe0, 0x35, 0x9c,
	0x49, 0x0c, 0x4d, 0xe0, 0x33, 0x9a, 0x9b, 0x8b, 0xb4, 0xb7, 0x13, 0x3e, 0xaa, 0xa8, 0xac, 0x49,
	0x59, 0xcc, 0xf3, 0xde, 0x7b, 0x76, 0x2b, 0xd9, 0x5b, 0x98, 0x3a, 0x6f, 0x5d, 0xd5, 0x58, 0x89,
	0xfc, 0x25, 0xf5, 0x27, 0x09, 0x7c, 0xb5, 0x12, 0xd3, 0xce, 0xbd, 0x49, 0xff, 0x3a, 0xaf, 0x1a,
	0xe4, 0x97, 0x79, 0x67, 0x42, 0xbf, 0x12, 0x79, 0x96, 0x0f, 0x9d, 0xee, 0xd5, 0xf3, 0x7c, 0xe8,
	0x7a, 0x1c, 0xce, 0x53, 0x65, 0xb0, 0xe3, 0x8c, 0xda, 0x43, 0xf9, 0x65, 0xf2, 0x7b, 0xac, 0xad,
	0xc4, 0x2e, 0xd4, 0x63, 0x7a, 0xdf, 0xe7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xcd, 0xcd,
	0x96, 0xe5, 0x02, 0x00, 0x00,
}
