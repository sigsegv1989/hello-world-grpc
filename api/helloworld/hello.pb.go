// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/hello.proto

package helloworld

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

type BatchRequest struct {
	Requests             []*Message `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{0}
}

func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest.Unmarshal(m, b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
}
func (m *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(m, src)
}
func (m *BatchRequest) XXX_Size() int {
	return xxx_messageInfo_BatchRequest.Size(m)
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetRequests() []*Message {
	if m != nil {
		return m.Requests
	}
	return nil
}

type BatchResponse struct {
	Responses            []*Message `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BatchResponse) Reset()         { *m = BatchResponse{} }
func (m *BatchResponse) String() string { return proto.CompactTextString(m) }
func (*BatchResponse) ProtoMessage()    {}
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{1}
}

func (m *BatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchResponse.Unmarshal(m, b)
}
func (m *BatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchResponse.Marshal(b, m, deterministic)
}
func (m *BatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse.Merge(m, src)
}
func (m *BatchResponse) XXX_Size() int {
	return xxx_messageInfo_BatchResponse.Size(m)
}
func (m *BatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse proto.InternalMessageInfo

func (m *BatchResponse) GetResponses() []*Message {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Message struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_860c41a0e7a04411, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *Message) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*BatchRequest)(nil), "helloworld.BatchRequest")
	proto.RegisterType((*BatchResponse)(nil), "helloworld.BatchResponse")
	proto.RegisterType((*Message)(nil), "helloworld.Message")
}

func init() {
	proto.RegisterFile("api/hello.proto", fileDescriptor_860c41a0e7a04411)
}

var fileDescriptor_860c41a0e7a04411 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0xc8, 0xd4,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x73, 0xca,
	0xf3, 0x8b, 0x72, 0x52, 0x94, 0xec, 0xb9, 0x78, 0x9c, 0x12, 0x4b, 0x92, 0x33, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf4, 0xb9, 0x38, 0x8a, 0x20, 0xcc, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x61, 0x3d, 0x84, 0x72, 0x3d, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4, 0x20,
	0xb8, 0x22, 0x25, 0x27, 0x2e, 0x5e, 0xa8, 0x01, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x86,
	0x5c, 0x9c, 0x45, 0x50, 0x36, 0x5e, 0x23, 0x10, 0xaa, 0x94, 0xfc, 0xb9, 0xd8, 0xa1, 0xa2, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6,
	0x90, 0x14, 0x17, 0x47, 0x7a, 0x51, 0x6a, 0x6a, 0x49, 0x66, 0x5e, 0xba, 0x04, 0x13, 0x58, 0x1c,
	0xce, 0x17, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x56, 0x60, 0xd4, 0x60,
	0x0d, 0x82, 0x70, 0x8c, 0xbc, 0xb8, 0xd8, 0xdd, 0x41, 0x2a, 0x52, 0x8b, 0x84, 0xec, 0xb9, 0x38,
	0x82, 0x13, 0x2b, 0x3d, 0x40, 0xf6, 0x0b, 0x49, 0x20, 0xbb, 0x03, 0xd9, 0xdb, 0x52, 0x92, 0x58,
	0x64, 0x20, 0xae, 0x73, 0x12, 0x8a, 0x12, 0xd0, 0xd3, 0x87, 0x07, 0x21, 0x58, 0x45, 0x12, 0x1b,
	0x38, 0x20, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x08, 0xd8, 0xb9, 0x5b, 0x01, 0x00,
	0x00,
}
