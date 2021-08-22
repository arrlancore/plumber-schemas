// Code generated by protoc-gen-go. DO NOT EDIT.
// source: write.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	encoding "github.com/batchcorp/plumber-schemas/build/go/protos/encoding"
	records "github.com/batchcorp/plumber-schemas/build/go/protos/records"
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

type WriteRecord struct {
	// Types that are valid to be assigned to Records:
	//	*WriteRecord_Kafka
	Records              isWriteRecord_Records `protobuf_oneof:"Records"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WriteRecord) Reset()         { *m = WriteRecord{} }
func (m *WriteRecord) String() string { return proto.CompactTextString(m) }
func (*WriteRecord) ProtoMessage()    {}
func (*WriteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_67966b2b12a73214, []int{0}
}

func (m *WriteRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRecord.Unmarshal(m, b)
}
func (m *WriteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRecord.Marshal(b, m, deterministic)
}
func (m *WriteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRecord.Merge(m, src)
}
func (m *WriteRecord) XXX_Size() int {
	return xxx_messageInfo_WriteRecord.Size(m)
}
func (m *WriteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRecord proto.InternalMessageInfo

type isWriteRecord_Records interface {
	isWriteRecord_Records()
}

type WriteRecord_Kafka struct {
	Kafka *records.Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

func (*WriteRecord_Kafka) isWriteRecord_Records() {}

func (m *WriteRecord) GetRecords() isWriteRecord_Records {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *WriteRecord) GetKafka() *records.Kafka {
	if x, ok := m.GetRecords().(*WriteRecord_Kafka); ok {
		return x.Kafka
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WriteRecord) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WriteRecord_Kafka)(nil),
	}
}

type WriteRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Must point to an existing connection
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// We need to use a WriteRecord as a wrapper since there is no way to "repeated oneof ..."
	Records []*WriteRecord `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	// Specifying encoding options will cause the _value_ of the record to be
	// encoded and set in *plumber.types.WriteMessage.Value.
	EncodeOptions        *encoding.Options `protobuf:"bytes,3,opt,name=encode_options,json=encodeOptions,proto3" json:"encode_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67966b2b12a73214, []int{1}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *WriteRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *WriteRequest) GetRecords() []*WriteRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *WriteRequest) GetEncodeOptions() *encoding.Options {
	if m != nil {
		return m.EncodeOptions
	}
	return nil
}

type WriteResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67966b2b12a73214, []int{2}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteRecord)(nil), "protos.WriteRecord")
	proto.RegisterType((*WriteRequest)(nil), "protos.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "protos.WriteResponse")
}

func init() { proto.RegisterFile("write.proto", fileDescriptor_67966b2b12a73214) }

var fileDescriptor_67966b2b12a73214 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xad, 0xd3, 0x8d, 0x65, 0x9b, 0x60, 0x87, 0x12, 0x76, 0x1a, 0xf3, 0xb2, 0xcb, 0x1a,
	0x50, 0xf1, 0x3a, 0xdc, 0x45, 0x45, 0x50, 0x88, 0x07, 0xc1, 0xcb, 0x68, 0xd3, 0xb8, 0x96, 0xad,
	0xf9, 0x6a, 0x93, 0xe0, 0xcf, 0xf0, 0x97, 0xf9, 0x3f, 0xfc, 0x19, 0x92, 0xe4, 0x2b, 0xce, 0x53,
	0xe9, 0xfb, 0xbe, 0x79, 0xf2, 0x7e, 0x5f, 0xc8, 0xe0, 0xb3, 0x29, 0x8d, 0x4c, 0xea, 0x06, 0x0c,
	0xc4, 0x5d, 0xff, 0xd1, 0x93, 0x71, 0x23, 0x05, 0x34, 0xb9, 0x66, 0xdb, 0xf4, 0x7d, 0x9b, 0x06,
	0x73, 0x72, 0x2a, 0xa0, 0xaa, 0x40, 0xb1, 0xd4, 0x9a, 0x02, 0xa5, 0x31, 0x4a, 0xda, 0xa4, 0xc6,
	0x6a, 0x14, 0xcf, 0xa5, 0x12, 0x90, 0x97, 0x6a, 0xc3, 0xa0, 0x36, 0x25, 0x28, 0xd4, 0x67, 0x77,
	0x64, 0xf0, 0xea, 0xee, 0xe2, 0x9e, 0x1d, 0x2f, 0xc8, 0xb1, 0xa7, 0xd3, 0x7c, 0x1a, 0xcd, 0x07,
	0x97, 0x67, 0x21, 0xa5, 0x13, 0xbc, 0x3a, 0x79, 0x74, 0xe6, 0xfd, 0x01, 0x0f, 0xa9, 0x55, 0x9f,
	0xf4, 0xc2, 0x41, 0x3d, 0xfb, 0x8e, 0xc8, 0x10, 0x49, 0x1f, 0x56, 0x6a, 0x13, 0xcf, 0xc9, 0x91,
	0x2b, 0x45, 0xbf, 0x9e, 0x3c, 0x6a, 0xdc, 0xa2, 0x42, 0xbb, 0xe4, 0xd6, 0x9a, 0x82, 0xfb, 0x44,
	0x7c, 0x41, 0x46, 0x02, 0x94, 0x92, 0xc2, 0x15, 0x5b, 0x97, 0x39, 0x8d, 0xa6, 0xd1, 0xbc, 0xcf,
	0x87, 0x7f, 0xe2, 0x83, 0x6b, 0xd6, 0xc3, 0x12, 0xf4, 0x70, 0xda, 0xd9, 0x07, 0xee, 0xf5, 0xe7,
	0x6d, 0x26, 0x5e, 0x92, 0x13, 0x3f, 0xb1, 0x5c, 0xe3, 0xbc, 0xb4, 0xe3, 0x6b, 0xd0, 0xf6, 0x54,
	0xbb, 0x8f, 0xe4, 0x39, 0xf8, 0x7c, 0x14, 0xf2, 0xf8, 0x3b, 0x5b, 0x92, 0x11, 0x82, 0x75, 0x0d,
	0x4a, 0xcb, 0x38, 0x21, 0xdd, 0xb0, 0x51, 0xfa, 0xd3, 0xfb, 0xbf, 0x1c, 0x9c, 0xe8, 0xc5, 0xbb,
	0x1c, 0x53, 0xab, 0x9b, 0xb7, 0xeb, 0x4d, 0x69, 0x0a, 0x9b, 0x39, 0x9f, 0x65, 0xa9, 0x11, 0x85,
	0x80, 0xa6, 0x66, 0xf5, 0xce, 0x56, 0x99, 0x6c, 0x16, 0x5a, 0x14, 0xb2, 0x4a, 0x35, 0xcb, 0x6c,
	0xb9, 0xcb, 0xd9, 0x06, 0x58, 0xa0, 0x65, 0xe1, 0xb9, 0xaf, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0x80, 0x4f, 0xda, 0x04, 0x02, 0x00, 0x00,
}
