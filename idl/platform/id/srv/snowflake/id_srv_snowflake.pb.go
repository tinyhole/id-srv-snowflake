// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform/id_srv_snowflake.proto

package snowflake

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetIDReq struct {
	Num                  int64    `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDReq) Reset()         { *m = GetIDReq{} }
func (m *GetIDReq) String() string { return proto.CompactTextString(m) }
func (*GetIDReq) ProtoMessage()    {}
func (*GetIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33a5768c42cac1, []int{0}
}

func (m *GetIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDReq.Unmarshal(m, b)
}
func (m *GetIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDReq.Marshal(b, m, deterministic)
}
func (m *GetIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDReq.Merge(m, src)
}
func (m *GetIDReq) XXX_Size() int {
	return xxx_messageInfo_GetIDReq.Size(m)
}
func (m *GetIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDReq proto.InternalMessageInfo

func (m *GetIDReq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GetIDRsp struct {
	IDs                  []int64  `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDRsp) Reset()         { *m = GetIDRsp{} }
func (m *GetIDRsp) String() string { return proto.CompactTextString(m) }
func (*GetIDRsp) ProtoMessage()    {}
func (*GetIDRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33a5768c42cac1, []int{1}
}

func (m *GetIDRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDRsp.Unmarshal(m, b)
}
func (m *GetIDRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDRsp.Marshal(b, m, deterministic)
}
func (m *GetIDRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDRsp.Merge(m, src)
}
func (m *GetIDRsp) XXX_Size() int {
	return xxx_messageInfo_GetIDRsp.Size(m)
}
func (m *GetIDRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDRsp proto.InternalMessageInfo

func (m *GetIDRsp) GetIDs() []int64 {
	if m != nil {
		return m.IDs
	}
	return nil
}

func init() {
	proto.RegisterType((*GetIDReq)(nil), "platform.id.srv.snowflake.GetIDReq")
	proto.RegisterType((*GetIDRsp)(nil), "platform.id.srv.snowflake.GetIDRsp")
}

func init() { proto.RegisterFile("platform/id_srv_snowflake.proto", fileDescriptor_cb33a5768c42cac1) }

var fileDescriptor_cb33a5768c42cac1 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xc8, 0x49, 0x2c,
	0x49, 0xcb, 0x2f, 0xca, 0xd5, 0xcf, 0x4c, 0x89, 0x2f, 0x2e, 0x2a, 0x8b, 0x2f, 0xce, 0xcb, 0x2f,
	0x4f, 0xcb, 0x49, 0xcc, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0x29, 0xd0,
	0xcb, 0x4c, 0xd1, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0x2b, 0x50, 0x92, 0xe1, 0xe2, 0x70, 0x4f, 0x2d,
	0xf1, 0x74, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0xf6, 0x2b, 0xcd, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x02, 0x31, 0x11, 0xb2, 0xc5, 0x05, 0x20, 0x59, 0x4f, 0x97, 0x62, 0x09, 0x46,
	0x05, 0x66, 0x90, 0xac, 0xa7, 0x4b, 0xb1, 0x51, 0x1c, 0x17, 0x67, 0x70, 0x5e, 0x7e, 0xb9, 0x1b,
	0xc8, 0x20, 0xa1, 0x40, 0x2e, 0x56, 0xb0, 0x52, 0x21, 0x65, 0x3d, 0x9c, 0xb6, 0xe9, 0xc1, 0xac,
	0x92, 0x22, 0xac, 0xa8, 0xb8, 0xc0, 0x49, 0x3a, 0x4a, 0x12, 0xc9, 0x67, 0xfa, 0xc5, 0x45, 0x65,
	0xfa, 0x70, 0x55, 0x49, 0x6c, 0x60, 0xaf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xc4,
	0xa0, 0xd1, 0xfd, 0x00, 0x00, 0x00,
}
