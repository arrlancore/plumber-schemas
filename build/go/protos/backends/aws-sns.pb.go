// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aws-sns.proto

package backends

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

type AWSSNSConn struct {
	// Required
	AwsRegion string `protobuf:"bytes,1,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty"`
	// Required
	AwsAccessKeyId string `protobuf:"bytes,2,opt,name=aws_access_key_id,json=awsAccessKeyId,proto3" json:"aws_access_key_id,omitempty"`
	// Required
	AwsSecretAccessKey   string   `protobuf:"bytes,3,opt,name=aws_secret_access_key,json=awsSecretAccessKey,proto3" json:"aws_secret_access_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSNSConn) Reset()         { *m = AWSSNSConn{} }
func (m *AWSSNSConn) String() string { return proto.CompactTextString(m) }
func (*AWSSNSConn) ProtoMessage()    {}
func (*AWSSNSConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac516fe8f86ec289, []int{0}
}

func (m *AWSSNSConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSNSConn.Unmarshal(m, b)
}
func (m *AWSSNSConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSNSConn.Marshal(b, m, deterministic)
}
func (m *AWSSNSConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSNSConn.Merge(m, src)
}
func (m *AWSSNSConn) XXX_Size() int {
	return xxx_messageInfo_AWSSNSConn.Size(m)
}
func (m *AWSSNSConn) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSNSConn.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSNSConn proto.InternalMessageInfo

func (m *AWSSNSConn) GetAwsRegion() string {
	if m != nil {
		return m.AwsRegion
	}
	return ""
}

func (m *AWSSNSConn) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *AWSSNSConn) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

type AWSSNSReadArgs struct {
	// Topic ARN to write to
	TopicArn             string   `protobuf:"bytes,1,opt,name=topic_arn,json=topicArn,proto3" json:"topic_arn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSNSReadArgs) Reset()         { *m = AWSSNSReadArgs{} }
func (m *AWSSNSReadArgs) String() string { return proto.CompactTextString(m) }
func (*AWSSNSReadArgs) ProtoMessage()    {}
func (*AWSSNSReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac516fe8f86ec289, []int{1}
}

func (m *AWSSNSReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSNSReadArgs.Unmarshal(m, b)
}
func (m *AWSSNSReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSNSReadArgs.Marshal(b, m, deterministic)
}
func (m *AWSSNSReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSNSReadArgs.Merge(m, src)
}
func (m *AWSSNSReadArgs) XXX_Size() int {
	return xxx_messageInfo_AWSSNSReadArgs.Size(m)
}
func (m *AWSSNSReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSNSReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSNSReadArgs proto.InternalMessageInfo

func (m *AWSSNSReadArgs) GetTopicArn() string {
	if m != nil {
		return m.TopicArn
	}
	return ""
}

type AWSSNSWriteArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSNSWriteArgs) Reset()         { *m = AWSSNSWriteArgs{} }
func (m *AWSSNSWriteArgs) String() string { return proto.CompactTextString(m) }
func (*AWSSNSWriteArgs) ProtoMessage()    {}
func (*AWSSNSWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac516fe8f86ec289, []int{2}
}

func (m *AWSSNSWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSNSWriteArgs.Unmarshal(m, b)
}
func (m *AWSSNSWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSNSWriteArgs.Marshal(b, m, deterministic)
}
func (m *AWSSNSWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSNSWriteArgs.Merge(m, src)
}
func (m *AWSSNSWriteArgs) XXX_Size() int {
	return xxx_messageInfo_AWSSNSWriteArgs.Size(m)
}
func (m *AWSSNSWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSNSWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSNSWriteArgs proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AWSSNSConn)(nil), "protos.backends.AWSSNSConn")
	proto.RegisterType((*AWSSNSReadArgs)(nil), "protos.backends.AWSSNSReadArgs")
	proto.RegisterType((*AWSSNSWriteArgs)(nil), "protos.backends.AWSSNSWriteArgs")
}

func init() { proto.RegisterFile("aws-sns.proto", fileDescriptor_ac516fe8f86ec289) }

var fileDescriptor_ac516fe8f86ec289 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x86, 0xe9, 0xf7, 0x81, 0xd8, 0x80, 0xbb, 0x6c, 0x40, 0x28, 0x88, 0x20, 0x3d, 0xe9, 0xa1,
	0x0d, 0xe2, 0x59, 0xa4, 0x7a, 0x12, 0xc1, 0x43, 0x7b, 0x58, 0xf0, 0x52, 0x26, 0xe9, 0xd0, 0x86,
	0xdd, 0x26, 0x25, 0x93, 0x12, 0xf6, 0x27, 0xf8, 0xaf, 0xa5, 0x29, 0x2b, 0x9e, 0x06, 0x9e, 0x79,
	0x5e, 0x78, 0x79, 0xd9, 0x15, 0x04, 0x2a, 0xc8, 0x50, 0x39, 0x39, 0xeb, 0x2d, 0xdf, 0xc6, 0x43,
	0xa5, 0x04, 0x75, 0x40, 0xd3, 0x51, 0xfe, 0x9d, 0x30, 0x56, 0xed, 0x9b, 0xe6, 0xb3, 0x79, 0xb3,
	0xc6, 0xf0, 0x5b, 0xc6, 0x20, 0x50, 0xeb, 0xb0, 0xd7, 0xd6, 0x64, 0xc9, 0x5d, 0x72, 0x9f, 0xd6,
	0x29, 0x04, 0xaa, 0x23, 0xe0, 0x0f, 0x6c, 0xb7, 0xbc, 0x41, 0x29, 0x24, 0x6a, 0x0f, 0x78, 0x6a,
	0x75, 0x97, 0xfd, 0x8b, 0xd6, 0x06, 0x02, 0x55, 0x91, 0x7f, 0xe0, 0xe9, 0xbd, 0xe3, 0x8f, 0xec,
	0x7a, 0x51, 0x09, 0x95, 0x43, 0xff, 0x27, 0x91, 0xfd, 0x8f, 0x3a, 0x87, 0x40, 0x4d, 0xfc, 0xfd,
	0x86, 0xf2, 0x82, 0x6d, 0xd6, 0x2a, 0x35, 0x42, 0x57, 0xb9, 0x9e, 0xf8, 0x0d, 0x4b, 0xbd, 0x9d,
	0xb4, 0x6a, 0xc1, 0x9d, 0xdb, 0x5c, 0x46, 0x50, 0x39, 0x93, 0xef, 0xd8, 0x76, 0xd5, 0xf7, 0x4e,
	0x7b, 0x5c, 0xfc, 0xd7, 0x97, 0xaf, 0xe7, 0x5e, 0xfb, 0x61, 0x96, 0xa5, 0xb2, 0xa3, 0x90, 0xe0,
	0xd5, 0xa0, 0xac, 0x9b, 0xc4, 0x74, 0x9c, 0x47, 0x89, 0xae, 0x20, 0x35, 0xe0, 0x08, 0x24, 0xe4,
	0xac, 0x8f, 0x9d, 0xe8, 0xad, 0x58, 0xe7, 0x10, 0xe7, 0x39, 0xe4, 0x45, 0x04, 0x4f, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x64, 0xed, 0x2c, 0x37, 0x01, 0x00, 0x00,
}
