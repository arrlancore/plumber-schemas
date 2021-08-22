// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nsq.proto

package conns

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

type NSQ struct {
	// Required for writes
	// Required for reads if lookupd_address is not specified
	NsqdAddress string `protobuf:"bytes,1,opt,name=nsqd_address,json=nsqdAddress,proto3" json:"nsqd_address,omitempty"`
	// Required for reads if nsqd_address is not specified
	// Ignored for writes
	LookupdAddress string `protobuf:"bytes,2,opt,name=lookupd_address,json=lookupdAddress,proto3" json:"lookupd_address,omitempty"`
	// Connect using TLS
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// Specify to not verify server's TLS certificate
	InsecureTls bool `protobuf:"varint,4,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty"`
	// Optional
	TlsCaCert []byte `protobuf:"bytes,5,opt,name=tls_ca_cert,json=tlsCaCert,proto3" json:"tls_ca_cert,omitempty"`
	// Optional
	TlsClientCert []byte `protobuf:"bytes,6,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty"`
	// Optional
	TlsClientKey []byte `protobuf:"bytes,7,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty"`
	// Optional
	AuthSecret string `protobuf:"bytes,8,opt,name=auth_secret,json=authSecret,proto3" json:"auth_secret,omitempty"`
	// Required
	ClientId             string   `protobuf:"bytes,9,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NSQ) Reset()         { *m = NSQ{} }
func (m *NSQ) String() string { return proto.CompactTextString(m) }
func (*NSQ) ProtoMessage()    {}
func (*NSQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_f83e7c7ac3c992fd, []int{0}
}

func (m *NSQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSQ.Unmarshal(m, b)
}
func (m *NSQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSQ.Marshal(b, m, deterministic)
}
func (m *NSQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSQ.Merge(m, src)
}
func (m *NSQ) XXX_Size() int {
	return xxx_messageInfo_NSQ.Size(m)
}
func (m *NSQ) XXX_DiscardUnknown() {
	xxx_messageInfo_NSQ.DiscardUnknown(m)
}

var xxx_messageInfo_NSQ proto.InternalMessageInfo

func (m *NSQ) GetNsqdAddress() string {
	if m != nil {
		return m.NsqdAddress
	}
	return ""
}

func (m *NSQ) GetLookupdAddress() string {
	if m != nil {
		return m.LookupdAddress
	}
	return ""
}

func (m *NSQ) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *NSQ) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *NSQ) GetTlsCaCert() []byte {
	if m != nil {
		return m.TlsCaCert
	}
	return nil
}

func (m *NSQ) GetTlsClientCert() []byte {
	if m != nil {
		return m.TlsClientCert
	}
	return nil
}

func (m *NSQ) GetTlsClientKey() []byte {
	if m != nil {
		return m.TlsClientKey
	}
	return nil
}

func (m *NSQ) GetAuthSecret() string {
	if m != nil {
		return m.AuthSecret
	}
	return ""
}

func (m *NSQ) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*NSQ)(nil), "protos.conns.NSQ")
}

func init() { proto.RegisterFile("nsq.proto", fileDescriptor_f83e7c7ac3c992fd) }

var fileDescriptor_f83e7c7ac3c992fd = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x71, 0xba, 0xfd, 0xff, 0xdb, 0x9a, 0xd6, 0x0d, 0x7a, 0xb1, 0x20, 0xe8, 0x14, 0xd1,
	0x5d, 0x5c, 0x0e, 0xde, 0xc4, 0x8b, 0xee, 0x24, 0x82, 0xe0, 0xe6, 0xc9, 0x4b, 0x49, 0x93, 0x87,
	0xb5, 0x2c, 0x4b, 0xba, 0x3c, 0xc9, 0x61, 0x2f, 0xcd, 0x77, 0x27, 0x7d, 0xba, 0xa9, 0xa7, 0xd2,
	0xef, 0xef, 0x43, 0x4b, 0xc2, 0x62, 0x83, 0xbb, 0x79, 0xe3, 0xac, 0xb7, 0x59, 0x4a, 0x0f, 0x9c,
	0x4b, 0x6b, 0x0c, 0x5e, 0x7d, 0xf5, 0x58, 0xff, 0x6d, 0xf5, 0x9e, 0x5d, 0xb2, 0xd4, 0xe0, 0x4e,
	0x15, 0x42, 0x29, 0x07, 0x88, 0x79, 0x34, 0x8d, 0x66, 0xf1, 0x32, 0x69, 0xdb, 0x53, 0x97, 0xb2,
	0x5b, 0x36, 0xd1, 0xd6, 0x6e, 0x42, 0xf3, 0xab, 0x7a, 0xa4, 0xc6, 0x87, 0x7c, 0x84, 0xa7, 0x6c,
	0x18, 0x10, 0x0a, 0xaf, 0x31, 0xef, 0x4f, 0xa3, 0xd9, 0x68, 0x39, 0x08, 0x08, 0x1f, 0x1a, 0xdb,
	0x9f, 0xd4, 0x06, 0x41, 0x06, 0xd7, 0xad, 0xff, 0x68, 0x4d, 0x8e, 0xad, 0x25, 0xe7, 0x2c, 0xf1,
	0x1a, 0x0b, 0x29, 0x0a, 0x09, 0xce, 0xe7, 0xff, 0xa7, 0xd1, 0x2c, 0x5d, 0xc6, 0x5e, 0xe3, 0x42,
	0x2c, 0xc0, 0xf9, 0xec, 0x86, 0x4d, 0x68, 0xd7, 0x35, 0x18, 0xdf, 0x99, 0x01, 0x99, 0x93, 0xd6,
	0x50, 0x25, 0x77, 0xcd, 0xc6, 0x7f, 0xdc, 0x06, 0xf6, 0xf9, 0x90, 0x58, 0xfa, 0xc3, 0x5e, 0x61,
	0x9f, 0x5d, 0xb0, 0x44, 0x04, 0x5f, 0x15, 0x08, 0xd2, 0x81, 0xcf, 0x47, 0x74, 0x1c, 0xd6, 0xa6,
	0x15, 0x95, 0xec, 0x8c, 0xc5, 0x87, 0x4f, 0xd4, 0x2a, 0x8f, 0x69, 0x1e, 0x75, 0xe1, 0x45, 0x3d,
	0x3f, 0x7e, 0x3e, 0xac, 0x6b, 0x5f, 0x85, 0x72, 0x2e, 0xed, 0x96, 0x97, 0xc2, 0xcb, 0x4a, 0x5a,
	0xd7, 0xf0, 0x46, 0x87, 0x6d, 0x09, 0xee, 0x0e, 0x65, 0x05, 0x5b, 0x81, 0xbc, 0x0c, 0xb5, 0x56,
	0x7c, 0x6d, 0x79, 0x77, 0xf3, 0x9c, 0x6e, 0xbe, 0x1c, 0xd0, 0xdb, 0xfd, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x92, 0xc0, 0x32, 0x9b, 0x01, 0x00, 0x00,
}