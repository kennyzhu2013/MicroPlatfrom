// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_membertask.proto

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

type TMemberTask struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskName             string               `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	CreateUser           string               `protobuf:"bytes,6,opt,name=create_user,json=createUser,proto3" json:"create_user,omitempty"`
	Lupddate             *timestamp.Timestamp `protobuf:"bytes,7,opt,name=lupddate,proto3" json:"lupddate,omitempty"`
	Remark               string               `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberTask) Reset()         { *m = TMemberTask{} }
func (m *TMemberTask) String() string { return proto.CompactTextString(m) }
func (*TMemberTask) ProtoMessage()    {}
func (*TMemberTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_membertask_c20d7b329200e35b, []int{0}
}
func (m *TMemberTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberTask.Unmarshal(m, b)
}
func (m *TMemberTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberTask.Marshal(b, m, deterministic)
}
func (dst *TMemberTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberTask.Merge(dst, src)
}
func (m *TMemberTask) XXX_Size() int {
	return xxx_messageInfo_TMemberTask.Size(m)
}
func (m *TMemberTask) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberTask.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberTask proto.InternalMessageInfo

func (m *TMemberTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMemberTask) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *TMemberTask) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TMemberTask) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TMemberTask) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberTask) GetCreateUser() string {
	if m != nil {
		return m.CreateUser
	}
	return ""
}

func (m *TMemberTask) GetLupddate() *timestamp.Timestamp {
	if m != nil {
		return m.Lupddate
	}
	return nil
}

func (m *TMemberTask) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberTask)(nil), "usercenter.TMemberTask")
}

func init() {
	proto.RegisterFile("usercenter/t_membertask.proto", fileDescriptor_t_membertask_c20d7b329200e35b)
}

var fileDescriptor_t_membertask_c20d7b329200e35b = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4f, 0xb4, 0x40,
	0x10, 0xc6, 0x03, 0xef, 0x2b, 0xc2, 0x90, 0x58, 0x6c, 0x61, 0x36, 0x67, 0x8c, 0x17, 0xab, 0xab,
	0x20, 0xd1, 0xc4, 0xc4, 0x5c, 0x67, 0x6c, 0xb5, 0xb8, 0x60, 0x63, 0x43, 0x06, 0x18, 0x91, 0xc0,
	0xb2, 0x64, 0x77, 0x28, 0xfc, 0x6c, 0x7e, 0x39, 0xb3, 0xcb, 0xfd, 0x29, 0xaf, 0xe3, 0x19, 0x7e,
	0xf3, 0x9b, 0x27, 0x0b, 0xb7, 0xb3, 0x25, 0x53, 0xd3, 0xc8, 0x64, 0x72, 0x2e, 0x15, 0xa9, 0x8a,
	0x0c, 0xa3, 0xed, 0xb3, 0xc9, 0x68, 0xd6, 0x02, 0x4e, 0xbf, 0x57, 0xdb, 0xb6, 0xe3, 0xef, 0xb9,
	0xca, 0x6a, 0xad, 0xf2, 0x56, 0x0f, 0x38, 0xb6, 0xb9, 0x87, 0xaa, 0xf9, 0x2b, 0x9f, 0xf8, 0x67,
	0x22, 0x9b, 0x73, 0xa7, 0xc8, 0x32, 0xaa, 0xe9, 0xf4, 0xb5, 0x88, 0xee, 0x7f, 0x43, 0x48, 0x8b,
	0x37, 0xaf, 0x2f, 0xd0, 0xf6, 0xe2, 0x0a, 0xc2, 0xae, 0x91, 0xc1, 0x3a, 0xd8, 0x24, 0xbb, 0xb0,
	0x6b, 0xc4, 0x0d, 0x24, 0xee, 0x6c, 0x39, 0xa2, 0x22, 0x19, 0xfa, 0x71, 0xec, 0x06, 0xef, 0xa8,
	0x48, 0x3c, 0x03, 0x58, 0x46, 0xc3, 0xa5, 0xb3, 0xca, 0x7f, 0xeb, 0x60, 0x93, 0x3e, 0xac, 0xb2,
	0x56, 0xeb, 0x76, 0xa0, 0xec, 0xd0, 0x21, 0x2b, 0x0e, 0x27, 0x77, 0x89, 0xa7, 0x5d, 0x16, 0xd7,
	0x10, 0x59, 0x46, 0x9e, 0xad, 0xfc, 0xef, 0xa5, 0xfb, 0x24, 0xb6, 0x90, 0xd6, 0x86, 0x90, 0xa9,
	0x6c, 0x90, 0x49, 0x5e, 0x9c, 0x75, 0xc2, 0x82, 0xbf, 0x22, 0x93, 0xb8, 0x3b, 0x2e, 0xbb, 0xe7,
	0x91, 0x91, 0x37, 0xef, 0x81, 0x0f, 0x4b, 0x46, 0x3c, 0x41, 0x3c, 0xcc, 0x53, 0xe3, 0xd5, 0x97,
	0x67, 0xd5, 0x47, 0xd6, 0xb5, 0x35, 0xa4, 0xd0, 0xf4, 0x32, 0x5e, 0xda, 0x2e, 0xe9, 0x25, 0xfe,
	0x8c, 0x94, 0x6e, 0x68, 0xb0, 0x55, 0xe4, 0xf7, 0x1f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0xe0, 0x8c, 0x04, 0xb8, 0x01, 0x00, 0x00,
}
