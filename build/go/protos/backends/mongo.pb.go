// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongo.proto

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

type MongoConn struct {
	// Full DNS to connect to mongo
	// Ex: mongodb://user:pass@localhost:27017
	Dsn                  string   `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoConn) Reset()         { *m = MongoConn{} }
func (m *MongoConn) String() string { return proto.CompactTextString(m) }
func (*MongoConn) ProtoMessage()    {}
func (*MongoConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c3854ae3094fec9, []int{0}
}

func (m *MongoConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoConn.Unmarshal(m, b)
}
func (m *MongoConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoConn.Marshal(b, m, deterministic)
}
func (m *MongoConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoConn.Merge(m, src)
}
func (m *MongoConn) XXX_Size() int {
	return xxx_messageInfo_MongoConn.Size(m)
}
func (m *MongoConn) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoConn.DiscardUnknown(m)
}

var xxx_messageInfo_MongoConn proto.InternalMessageInfo

func (m *MongoConn) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

type MongoReadArgs struct {
	// Optional
	// Specify to limit CDC changes to a single database of the mongo instance
	// Leave empty to get CDC changes from all databases
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Optional
	// Specify to limit CDC changes to a single collection.
	// Leave blank to get changes from all collections in the specified database
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	// Include full document in update in update changes.
	// Default is to return deltas only
	IncludeFullDocument  bool     `protobuf:"varint,3,opt,name=include_full_document,json=includeFullDocument,proto3" json:"include_full_document,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoReadArgs) Reset()         { *m = MongoReadArgs{} }
func (m *MongoReadArgs) String() string { return proto.CompactTextString(m) }
func (*MongoReadArgs) ProtoMessage()    {}
func (*MongoReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c3854ae3094fec9, []int{1}
}

func (m *MongoReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoReadArgs.Unmarshal(m, b)
}
func (m *MongoReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoReadArgs.Marshal(b, m, deterministic)
}
func (m *MongoReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoReadArgs.Merge(m, src)
}
func (m *MongoReadArgs) XXX_Size() int {
	return xxx_messageInfo_MongoReadArgs.Size(m)
}
func (m *MongoReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MongoReadArgs proto.InternalMessageInfo

func (m *MongoReadArgs) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *MongoReadArgs) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *MongoReadArgs) GetIncludeFullDocument() bool {
	if m != nil {
		return m.IncludeFullDocument
	}
	return false
}

// TODO: Add correct kong tag to not allow this option
type MongoWriteArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoWriteArgs) Reset()         { *m = MongoWriteArgs{} }
func (m *MongoWriteArgs) String() string { return proto.CompactTextString(m) }
func (*MongoWriteArgs) ProtoMessage()    {}
func (*MongoWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c3854ae3094fec9, []int{2}
}

func (m *MongoWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoWriteArgs.Unmarshal(m, b)
}
func (m *MongoWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoWriteArgs.Marshal(b, m, deterministic)
}
func (m *MongoWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoWriteArgs.Merge(m, src)
}
func (m *MongoWriteArgs) XXX_Size() int {
	return xxx_messageInfo_MongoWriteArgs.Size(m)
}
func (m *MongoWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MongoWriteArgs proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MongoConn)(nil), "protos.backends.MongoConn")
	proto.RegisterType((*MongoReadArgs)(nil), "protos.backends.MongoReadArgs")
	proto.RegisterType((*MongoWriteArgs)(nil), "protos.backends.MongoWriteArgs")
}

func init() { proto.RegisterFile("mongo.proto", fileDescriptor_5c3854ae3094fec9) }

var fileDescriptor_5c3854ae3094fec9 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xbb, 0x4b, 0xc4, 0x40,
	0x10, 0xc6, 0x89, 0x07, 0x72, 0xb7, 0xa2, 0x1e, 0x2b, 0x42, 0x10, 0x94, 0x23, 0xd5, 0x35, 0x66,
	0x41, 0x6b, 0x11, 0x1f, 0xd8, 0xd9, 0xa4, 0x11, 0x6c, 0xc2, 0x3e, 0xc6, 0x64, 0x71, 0x76, 0x27,
	0xec, 0xa3, 0xf6, 0x5f, 0x97, 0xac, 0x51, 0xac, 0x66, 0xe6, 0xfb, 0x4d, 0xf1, 0xfd, 0xd8, 0x91,
	0x23, 0x3f, 0x50, 0x3b, 0x05, 0x4a, 0xc4, 0x4f, 0xcb, 0x88, 0xad, 0x92, 0xfa, 0x13, 0xbc, 0x89,
	0xcd, 0x25, 0xdb, 0xbc, 0xce, 0xfc, 0x89, 0xbc, 0xe7, 0x5b, 0xb6, 0x32, 0xd1, 0xd7, 0xd5, 0xae,
	0xda, 0x6f, 0xba, 0x79, 0x6d, 0xbe, 0xd8, 0x71, 0xc1, 0x1d, 0x48, 0xf3, 0x10, 0x86, 0xc8, 0x2f,
	0xd8, 0xda, 0xc8, 0x24, 0x95, 0x8c, 0xb0, 0xfc, 0xfd, 0xdd, 0xfc, 0x8a, 0x31, 0x4d, 0x88, 0xa0,
	0x93, 0x25, 0x5f, 0x1f, 0x14, 0xfa, 0x2f, 0xe1, 0x37, 0xec, 0xdc, 0x7a, 0x8d, 0xd9, 0x40, 0xff,
	0x91, 0x11, 0x7b, 0x43, 0x3a, 0x3b, 0xf0, 0xa9, 0x5e, 0xed, 0xaa, 0xfd, 0xba, 0x3b, 0x5b, 0xe0,
	0x4b, 0x46, 0x7c, 0x5e, 0x50, 0xb3, 0x65, 0x27, 0xa5, 0xc0, 0x5b, 0xb0, 0x09, 0xe6, 0x06, 0x8f,
	0xf7, 0xef, 0x77, 0x83, 0x4d, 0x63, 0x56, 0xad, 0x26, 0x27, 0x94, 0x4c, 0x7a, 0xd4, 0x14, 0x26,
	0x31, 0x61, 0x76, 0x0a, 0xc2, 0x75, 0xd4, 0x23, 0x38, 0x19, 0x85, 0xca, 0x16, 0x8d, 0x18, 0x48,
	0xfc, 0x28, 0x8b, 0x5f, 0x65, 0x75, 0x58, 0x82, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x7a, 0x17, 0xb6, 0x19, 0x01, 0x00, 0x00,
}
