// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_cpbillcycleprocessstatus.proto

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

type TCpBillCycleProcessStatus struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastFullSyncDate     string   `protobuf:"bytes,2,opt,name=last_full_sync_date,json=lastFullSyncDate,proto3" json:"last_full_sync_date,omitempty"`
	FileNameProcessing   string   `protobuf:"bytes,3,opt,name=file_name_processing,json=fileNameProcessing,proto3" json:"file_name_processing,omitempty"`
	FileProcessResult    string   `protobuf:"bytes,4,opt,name=file_process_result,json=fileProcessResult,proto3" json:"file_process_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TCpBillCycleProcessStatus) Reset()         { *m = TCpBillCycleProcessStatus{} }
func (m *TCpBillCycleProcessStatus) String() string { return proto.CompactTextString(m) }
func (*TCpBillCycleProcessStatus) ProtoMessage()    {}
func (*TCpBillCycleProcessStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_cpbillcycleprocessstatus_aefcf449db122a7e, []int{0}
}
func (m *TCpBillCycleProcessStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TCpBillCycleProcessStatus.Unmarshal(m, b)
}
func (m *TCpBillCycleProcessStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TCpBillCycleProcessStatus.Marshal(b, m, deterministic)
}
func (dst *TCpBillCycleProcessStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TCpBillCycleProcessStatus.Merge(dst, src)
}
func (m *TCpBillCycleProcessStatus) XXX_Size() int {
	return xxx_messageInfo_TCpBillCycleProcessStatus.Size(m)
}
func (m *TCpBillCycleProcessStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TCpBillCycleProcessStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TCpBillCycleProcessStatus proto.InternalMessageInfo

func (m *TCpBillCycleProcessStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TCpBillCycleProcessStatus) GetLastFullSyncDate() string {
	if m != nil {
		return m.LastFullSyncDate
	}
	return ""
}

func (m *TCpBillCycleProcessStatus) GetFileNameProcessing() string {
	if m != nil {
		return m.FileNameProcessing
	}
	return ""
}

func (m *TCpBillCycleProcessStatus) GetFileProcessResult() string {
	if m != nil {
		return m.FileProcessResult
	}
	return ""
}

func init() {
	proto.RegisterType((*TCpBillCycleProcessStatus)(nil), "usercenter.TCpBillCycleProcessStatus")
}

func init() {
	proto.RegisterFile("usercenter/t_cpbillcycleprocessstatus.proto", fileDescriptor_t_cpbillcycleprocessstatus_aefcf449db122a7e)
}

var fileDescriptor_t_cpbillcycleprocessstatus_aefcf449db122a7e = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0x47, 0x49, 0x94, 0x43, 0xb7, 0x10, 0xdd, 0xb3, 0x88, 0x56, 0x62, 0x25, 0x88, 0x89, 0x60,
	0x69, 0x77, 0x27, 0x96, 0x22, 0x77, 0x56, 0x36, 0xcb, 0x66, 0x33, 0x89, 0x0b, 0xb3, 0x7f, 0xc8,
	0xcc, 0x16, 0xf9, 0x6e, 0x7e, 0x38, 0xc9, 0x5e, 0xe4, 0xba, 0x85, 0xf7, 0x66, 0x1f, 0x3f, 0xf1,
	0x98, 0x08, 0x46, 0x03, 0x9e, 0x61, 0x6c, 0x58, 0x99, 0xd8, 0x5a, 0x44, 0x33, 0x19, 0x84, 0x38,
	0x06, 0x03, 0x44, 0xc4, 0x9a, 0x13, 0xd5, 0x71, 0x0c, 0x1c, 0xa4, 0x38, 0xca, 0xb7, 0xaf, 0x83,
	0xe5, 0x9f, 0xd4, 0xd6, 0x26, 0xb8, 0x66, 0x08, 0xa8, 0xfd, 0xd0, 0x64, 0xa9, 0x4d, 0x7d, 0x13,
	0x79, 0x8a, 0x40, 0x0d, 0x5b, 0x07, 0xc4, 0xda, 0xc5, 0xe3, 0xeb, 0xf0, 0xd1, 0xfd, 0x6f, 0x21,
	0x6e, 0xbe, 0xb6, 0x71, 0x63, 0x11, 0xb7, 0x73, 0xec, 0xf3, 0x10, 0xdb, 0xe7, 0x98, 0xbc, 0x10,
	0xa5, 0xed, 0xaa, 0xe2, 0xae, 0x78, 0x38, 0xdf, 0x95, 0xb6, 0x93, 0x4f, 0x62, 0x8d, 0x9a, 0x58,
	0xf5, 0x09, 0x51, 0xd1, 0xe4, 0x8d, 0xea, 0x34, 0x43, 0x55, 0x66, 0xe1, 0x72, 0x46, 0xef, 0x09,
	0x71, 0x3f, 0x79, 0xf3, 0xa6, 0x19, 0xe4, 0xb3, 0xb8, 0xee, 0x2d, 0x82, 0xf2, 0xda, 0x81, 0x5a,
	0x66, 0x58, 0x3f, 0x54, 0x27, 0xd9, 0x97, 0x33, 0xfb, 0xd0, 0xee, 0xbf, 0x69, 0xfd, 0x20, 0x6b,
	0xb1, 0xce, 0x17, 0x8b, 0xac, 0x46, 0xa0, 0x84, 0x5c, 0x9d, 0xe6, 0x83, 0xab, 0x19, 0x2d, 0xf2,
	0x2e, 0x83, 0xcd, 0xd9, 0xf7, 0xca, 0x85, 0x0e, 0x90, 0xda, 0x55, 0xde, 0xf3, 0xf2, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x71, 0x73, 0xd9, 0x27, 0x47, 0x01, 0x00, 0x00,
}