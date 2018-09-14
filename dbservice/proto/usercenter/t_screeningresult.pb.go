// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_screeningresult.proto

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

type TScreeningResult struct {
	ModelId              string               `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	UserNum              string               `protobuf:"bytes,3,opt,name=user_num,json=userNum,proto3" json:"user_num,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	LocalFilePath        string               `protobuf:"bytes,5,opt,name=local_file_path,json=localFilePath,proto3" json:"local_file_path,omitempty"`
	CdnUrl               string               `protobuf:"bytes,6,opt,name=cdn_url,json=cdnUrl,proto3" json:"cdn_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TScreeningResult) Reset()         { *m = TScreeningResult{} }
func (m *TScreeningResult) String() string { return proto.CompactTextString(m) }
func (*TScreeningResult) ProtoMessage()    {}
func (*TScreeningResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_screeningresult_2148ad864b25e86b, []int{0}
}
func (m *TScreeningResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TScreeningResult.Unmarshal(m, b)
}
func (m *TScreeningResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TScreeningResult.Marshal(b, m, deterministic)
}
func (dst *TScreeningResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TScreeningResult.Merge(dst, src)
}
func (m *TScreeningResult) XXX_Size() int {
	return xxx_messageInfo_TScreeningResult.Size(m)
}
func (m *TScreeningResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TScreeningResult.DiscardUnknown(m)
}

var xxx_messageInfo_TScreeningResult proto.InternalMessageInfo

func (m *TScreeningResult) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *TScreeningResult) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TScreeningResult) GetUserNum() string {
	if m != nil {
		return m.UserNum
	}
	return ""
}

func (m *TScreeningResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TScreeningResult) GetLocalFilePath() string {
	if m != nil {
		return m.LocalFilePath
	}
	return ""
}

func (m *TScreeningResult) GetCdnUrl() string {
	if m != nil {
		return m.CdnUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*TScreeningResult)(nil), "usercenter.TScreeningResult")
}

func init() {
	proto.RegisterFile("usercenter/t_screeningresult.proto", fileDescriptor_t_screeningresult_2148ad864b25e86b)
}

var fileDescriptor_t_screeningresult_2148ad864b25e86b = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xf7, 0xd5, 0x69, 0x4d, 0x11, 0x25, 0x0b, 0x8d, 0x5d, 0x95, 0x2e, 0xa4, 0xab,
	0x19, 0xd0, 0x65, 0x77, 0x22, 0x82, 0x1b, 0x91, 0x5a, 0x37, 0x6e, 0x42, 0x9a, 0xdc, 0x66, 0x02,
	0xf9, 0x18, 0x92, 0x9b, 0x85, 0x7f, 0xd6, 0xdf, 0x22, 0x93, 0x71, 0xe8, 0x2e, 0x27, 0xf7, 0xf0,
	0x3c, 0x87, 0xac, 0x73, 0x82, 0x28, 0xc1, 0x23, 0xc4, 0x16, 0x79, 0x92, 0x11, 0xc0, 0x1b, 0xaf,
	0x23, 0xa4, 0x6c, 0xb1, 0xe9, 0x63, 0xc0, 0x40, 0xc9, 0xa9, 0xb3, 0xdc, 0x6a, 0x83, 0x5d, 0x3e,
	0x34, 0x32, 0xb8, 0x56, 0x07, 0x2b, 0xbc, 0x6e, 0x4b, 0xe9, 0x90, 0x8f, 0x6d, 0x8f, 0xdf, 0x3d,
	0xa4, 0x16, 0x8d, 0x83, 0x84, 0xc2, 0xf5, 0xa7, 0xd7, 0x08, 0x5a, 0xff, 0x54, 0xe4, 0x7a, 0xff,
	0x31, 0x39, 0x76, 0xc5, 0x41, 0xef, 0xc8, 0xdc, 0x05, 0x05, 0x96, 0x1b, 0xc5, 0xaa, 0x55, 0xb5,
	0xb9, 0xd8, 0xcd, 0x4a, 0x7e, 0x55, 0x74, 0x4b, 0x16, 0x32, 0x82, 0x40, 0xe0, 0x4a, 0x20, 0xb0,
	0x7f, 0xab, 0x6a, 0xb3, 0x78, 0x58, 0x36, 0x3a, 0x04, 0x6d, 0xa1, 0x99, 0xbc, 0xcd, 0x7e, 0xd2,
	0xec, 0xc8, 0x58, 0x7f, 0x16, 0x08, 0x03, 0x77, 0xd8, 0xcd, 0x7d, 0x76, 0xec, 0xff, 0xc8, 0x1d,
	0xf2, 0x5b, 0x76, 0xf4, 0x86, 0xd4, 0x09, 0x05, 0xe6, 0xc4, 0xce, 0xca, 0xe1, 0x2f, 0xd1, 0x7b,
	0x72, 0x65, 0x83, 0x14, 0x96, 0x1f, 0x8d, 0x05, 0xde, 0x0b, 0xec, 0xd8, 0x79, 0x29, 0x5c, 0x96,
	0xef, 0x17, 0x63, 0xe1, 0x5d, 0x60, 0x47, 0x6f, 0xc9, 0x4c, 0x2a, 0xcf, 0x73, 0xb4, 0xac, 0x1e,
	0x01, 0x52, 0xf9, 0xcf, 0x68, 0x9f, 0xe6, 0x5f, 0x75, 0xd9, 0x9e, 0x0e, 0x75, 0x59, 0xf7, 0xf8,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x93, 0xe4, 0x8e, 0x60, 0x01, 0x00, 0x00,
}
