// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka.proto

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

type KafkaConn_SASLType int32

const (
	KafkaConn_NONE  KafkaConn_SASLType = 0
	KafkaConn_PLAIN KafkaConn_SASLType = 1
	KafkaConn_SCRAM KafkaConn_SASLType = 2
)

var KafkaConn_SASLType_name = map[int32]string{
	0: "NONE",
	1: "PLAIN",
	2: "SCRAM",
}

var KafkaConn_SASLType_value = map[string]int32{
	"NONE":  0,
	"PLAIN": 1,
	"SCRAM": 2,
}

func (x KafkaConn_SASLType) String() string {
	return proto.EnumName(KafkaConn_SASLType_name, int32(x))
}

func (KafkaConn_SASLType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{0, 0}
}

type KafkaConn struct {
	// Required
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
	// Optional; how long to attempt connecting for (default: 10s)
	TimeoutSeconds int32 `protobuf:"varint,2,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	// Optional
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// Optional
	InsecureTls bool `protobuf:"varint,4,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty"`
	// Optional
	SaslType KafkaConn_SASLType `protobuf:"varint,5,opt,name=sasl_type,json=saslType,proto3,enum=protos.backends.KafkaConn_SASLType" json:"sasl_type,omitempty"`
	// Required if sasl_type is not NONE
	SaslUsername string `protobuf:"bytes,6,opt,name=sasl_username,json=saslUsername,proto3" json:"sasl_username,omitempty"`
	// Required if sasl_type is not NONE
	SaslPassword         string   `protobuf:"bytes,7,opt,name=sasl_password,json=saslPassword,proto3" json:"sasl_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaConn) Reset()         { *m = KafkaConn{} }
func (m *KafkaConn) String() string { return proto.CompactTextString(m) }
func (*KafkaConn) ProtoMessage()    {}
func (*KafkaConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{0}
}

func (m *KafkaConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaConn.Unmarshal(m, b)
}
func (m *KafkaConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaConn.Marshal(b, m, deterministic)
}
func (m *KafkaConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaConn.Merge(m, src)
}
func (m *KafkaConn) XXX_Size() int {
	return xxx_messageInfo_KafkaConn.Size(m)
}
func (m *KafkaConn) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaConn.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaConn proto.InternalMessageInfo

func (m *KafkaConn) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *KafkaConn) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *KafkaConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *KafkaConn) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *KafkaConn) GetSaslType() KafkaConn_SASLType {
	if m != nil {
		return m.SaslType
	}
	return KafkaConn_NONE
}

func (m *KafkaConn) GetSaslUsername() string {
	if m != nil {
		return m.SaslUsername
	}
	return ""
}

func (m *KafkaConn) GetSaslPassword() string {
	if m != nil {
		return m.SaslPassword
	}
	return ""
}

type KafkaReadArgs struct {
	// Required
	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	// Optional; specify what offset the consumer should read from (only works if '--use-consumer-group' is false)
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty"`
	// Optional: (default: true)
	UseConsumerGroup bool `protobuf:"varint,3,opt,name=use_consumer_group,json=useConsumerGroup,proto3" json:"use_consumer_group,omitempty"`
	// Optional; used only if "use_consumer_group" is true
	ConsumerGroupName string `protobuf:"bytes,4,opt,name=consumer_group_name,json=consumerGroupName,proto3" json:"consumer_group_name,omitempty"`
	// Optional; how long to wait for new data when reading batches of messages (default: 1s)
	MaxWaitSeconds int32 `protobuf:"varint,5,opt,name=max_wait_seconds,json=maxWaitSeconds,proto3" json:"max_wait_seconds,omitempty"`
	// Optional; minimum number of bytes to fetch in a single kafka request (throughput optimization)
	MinBytes int32 `protobuf:"varint,6,opt,name=min_bytes,json=minBytes,proto3" json:"min_bytes,omitempty"`
	// Optional; maximum number of bytes to fetch in a single kafka request (throughput optimization)
	MaxBytes int32 `protobuf:"varint,7,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Optional; how often to commit offsets to broker (default: 0 == synchronous)
	// NOTE: Used only-if "use_consumer_group" is true
	CommitIntervalSeconds int32 `protobuf:"varint,8,opt,name=commit_interval_seconds,json=commitIntervalSeconds,proto3" json:"commit_interval_seconds,omitempty"`
	// Optional; how long a coordinator will wait for member joins as part of a rebalance (default: 5s)
	RebalanceTimeoutSeconds int32    `protobuf:"varint,9,opt,name=rebalance_timeout_seconds,json=rebalanceTimeoutSeconds,proto3" json:"rebalance_timeout_seconds,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *KafkaReadArgs) Reset()         { *m = KafkaReadArgs{} }
func (m *KafkaReadArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaReadArgs) ProtoMessage()    {}
func (*KafkaReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{1}
}

func (m *KafkaReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaReadArgs.Unmarshal(m, b)
}
func (m *KafkaReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaReadArgs.Marshal(b, m, deterministic)
}
func (m *KafkaReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaReadArgs.Merge(m, src)
}
func (m *KafkaReadArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaReadArgs.Size(m)
}
func (m *KafkaReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaReadArgs proto.InternalMessageInfo

func (m *KafkaReadArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *KafkaReadArgs) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *KafkaReadArgs) GetUseConsumerGroup() bool {
	if m != nil {
		return m.UseConsumerGroup
	}
	return false
}

func (m *KafkaReadArgs) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *KafkaReadArgs) GetMaxWaitSeconds() int32 {
	if m != nil {
		return m.MaxWaitSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetMinBytes() int32 {
	if m != nil {
		return m.MinBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetCommitIntervalSeconds() int32 {
	if m != nil {
		return m.CommitIntervalSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetRebalanceTimeoutSeconds() int32 {
	if m != nil {
		return m.RebalanceTimeoutSeconds
	}
	return 0
}

type KafkaWriteArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaWriteArgs) Reset()         { *m = KafkaWriteArgs{} }
func (m *KafkaWriteArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaWriteArgs) ProtoMessage()    {}
func (*KafkaWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{2}
}

func (m *KafkaWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaWriteArgs.Unmarshal(m, b)
}
func (m *KafkaWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaWriteArgs.Marshal(b, m, deterministic)
}
func (m *KafkaWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaWriteArgs.Merge(m, src)
}
func (m *KafkaWriteArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaWriteArgs.Size(m)
}
func (m *KafkaWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaWriteArgs proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("protos.backends.KafkaConn_SASLType", KafkaConn_SASLType_name, KafkaConn_SASLType_value)
	proto.RegisterType((*KafkaConn)(nil), "protos.backends.KafkaConn")
	proto.RegisterType((*KafkaReadArgs)(nil), "protos.backends.KafkaReadArgs")
	proto.RegisterType((*KafkaWriteArgs)(nil), "protos.backends.KafkaWriteArgs")
}

func init() { proto.RegisterFile("kafka.proto", fileDescriptor_68928ed13de9fb92) }

var fileDescriptor_68928ed13de9fb92 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0xfd, 0x4c, 0xdc, 0xad, 0x0b, 0x46, 0xd0, 0x20, 0x1e, 0x28, 0xdd, 0x03, 0x11, 0x82,
	0x44, 0x02, 0x89, 0x07, 0x24, 0x04, 0x5d, 0x85, 0xd0, 0xc4, 0xe8, 0xa6, 0xb4, 0x68, 0x12, 0x2f,
	0x91, 0xe3, 0xb8, 0xad, 0xd5, 0x38, 0x8e, 0x6c, 0x87, 0xb5, 0x7f, 0x8f, 0x1f, 0x86, 0x90, 0x9d,
	0xa4, 0xfb, 0x78, 0xb2, 0xef, 0x39, 0xc7, 0xbe, 0xf7, 0x9e, 0x7b, 0xc1, 0x60, 0x8b, 0x56, 0x5b,
	0x14, 0x14, 0x82, 0x2b, 0x0e, 0x4f, 0xcc, 0x21, 0x83, 0x04, 0xe1, 0x2d, 0xc9, 0x53, 0x39, 0xf9,
	0xdb, 0x02, 0xce, 0x0f, 0x2d, 0x98, 0xf1, 0x3c, 0x87, 0x1e, 0xe8, 0xa3, 0x34, 0x15, 0x44, 0x4a,
	0xcf, 0x1a, 0xb7, 0x7d, 0x27, 0x6a, 0x42, 0xf8, 0x1a, 0x9c, 0x28, 0xca, 0x08, 0x2f, 0x55, 0x2c,
	0x09, 0xe6, 0x79, 0x2a, 0xbd, 0xd6, 0xd8, 0xf2, 0xbb, 0xd1, 0xb0, 0x86, 0x17, 0x15, 0x0a, 0x47,
	0xa0, 0x5f, 0x4a, 0x12, 0xab, 0x4c, 0x7a, 0xed, 0xb1, 0xe5, 0xdb, 0x51, 0xaf, 0x94, 0x64, 0x99,
	0x49, 0xf8, 0x0a, 0x1c, 0xd1, 0x5c, 0x12, 0x5c, 0x8a, 0x8a, 0xed, 0x18, 0x76, 0xd0, 0x60, 0x5a,
	0xf2, 0x15, 0x38, 0x12, 0xc9, 0x2c, 0x56, 0xfb, 0x82, 0x78, 0xdd, 0xb1, 0xe5, 0x0f, 0xdf, 0x9f,
	0x06, 0x0f, 0x2a, 0x0e, 0x0e, 0xd5, 0x06, 0x8b, 0xe9, 0xe2, 0x62, 0xb9, 0x2f, 0x48, 0x64, 0xeb,
	0x57, 0xfa, 0x06, 0x4f, 0xc1, 0xb1, 0xf9, 0xa1, 0x94, 0x44, 0xe4, 0x88, 0x11, 0xaf, 0x37, 0xb6,
	0x7c, 0x27, 0x3a, 0xd2, 0xe0, 0xaf, 0x1a, 0x3b, 0x88, 0x0a, 0x24, 0xe5, 0x0d, 0x17, 0xa9, 0xd7,
	0xbf, 0x15, 0x5d, 0xd5, 0xd8, 0xe4, 0x0d, 0xb0, 0x9b, 0xff, 0xa1, 0x0d, 0x3a, 0xf3, 0xcb, 0xf9,
	0x37, 0xf7, 0x11, 0x74, 0x40, 0xf7, 0xea, 0x62, 0x7a, 0x3e, 0x77, 0x2d, 0x7d, 0x5d, 0xcc, 0xa2,
	0xe9, 0x4f, 0xb7, 0x35, 0xf9, 0xd7, 0x02, 0xc7, 0xa6, 0xac, 0x88, 0xa0, 0x74, 0x2a, 0xd6, 0x12,
	0x3e, 0x03, 0x3d, 0xc5, 0x0b, 0x8a, 0x1b, 0x1f, 0xeb, 0x08, 0xbe, 0x04, 0x03, 0x41, 0x50, 0x1a,
	0xf3, 0xd5, 0x4a, 0x12, 0x65, 0x2c, 0x6c, 0x47, 0x40, 0x43, 0x97, 0x06, 0x81, 0x6f, 0x01, 0xd4,
	0xf6, 0x61, 0x9e, 0xcb, 0x92, 0x11, 0x11, 0xaf, 0x05, 0x2f, 0x8b, 0xda, 0x49, 0xb7, 0x94, 0x64,
	0x56, 0x13, 0xdf, 0x35, 0x0e, 0x03, 0xf0, 0xe4, 0xbe, 0x32, 0x36, 0x4d, 0x77, 0x4c, 0x3f, 0x8f,
	0xf1, 0x5d, 0xed, 0x5c, 0x77, 0xee, 0x03, 0x97, 0xa1, 0x5d, 0x7c, 0x83, 0xe8, 0xed, 0x18, 0xbb,
	0xd5, 0x18, 0x19, 0xda, 0x5d, 0x23, 0x7a, 0x18, 0xe3, 0x0b, 0xe0, 0x30, 0x9a, 0xc7, 0xc9, 0x5e,
	0x11, 0x69, 0x4c, 0xec, 0x46, 0x36, 0xa3, 0xf9, 0x99, 0x8e, 0x0d, 0x89, 0x76, 0x35, 0xd9, 0xaf,
	0x49, 0xb4, 0xab, 0xc8, 0x8f, 0x60, 0x84, 0x39, 0x63, 0x54, 0xc5, 0x34, 0x57, 0x44, 0xfc, 0x41,
	0xd9, 0x21, 0x95, 0x6d, 0xa4, 0x4f, 0x2b, 0xfa, 0xbc, 0x66, 0x9b, 0x8c, 0x9f, 0xc0, 0x73, 0x41,
	0x12, 0x94, 0xa1, 0x1c, 0x93, 0xf8, 0xe1, 0xae, 0x39, 0xe6, 0xe5, 0xe8, 0x20, 0x58, 0xde, 0x5b,
	0xba, 0x89, 0x0b, 0x86, 0xc6, 0xff, 0x6b, 0x41, 0x15, 0xd1, 0x03, 0x38, 0xfb, 0xf2, 0xfb, 0xf3,
	0x9a, 0xaa, 0x4d, 0x99, 0x04, 0x98, 0xb3, 0x30, 0x41, 0x0a, 0x6f, 0x30, 0x17, 0x45, 0x58, 0x64,
	0x25, 0x4b, 0x88, 0x78, 0x27, 0xf1, 0x86, 0x30, 0x24, 0xc3, 0xa4, 0xa4, 0x59, 0x1a, 0xae, 0x79,
	0x58, 0xad, 0x59, 0xd8, 0xac, 0x59, 0xd2, 0x33, 0xc0, 0x87, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0x27, 0xa4, 0x61, 0x3f, 0x03, 0x00, 0x00,
}
