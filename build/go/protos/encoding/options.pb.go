// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package encoding

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

type Type int32

const (
	Type_NONE        Type = 0
	Type_JSON        Type = 1
	Type_JSON_SCHEMA Type = 2
	Type_PROTOBUF    Type = 3
	Type_AVRO        Type = 4
)

var Type_name = map[int32]string{
	0: "NONE",
	1: "JSON",
	2: "JSON_SCHEMA",
	3: "PROTOBUF",
	4: "AVRO",
}

var Type_value = map[string]int32{
	"NONE":        0,
	"JSON":        1,
	"JSON_SCHEMA": 2,
	"PROTOBUF":    3,
	"AVRO":        4,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

type Protobuf struct {
	RootType             string   `protobuf:"bytes,1,opt,name=root_type,json=rootType,proto3" json:"root_type,omitempty"`
	ZipArchive           []byte   `protobuf:"bytes,2,opt,name=zip_archive,json=zipArchive,proto3" json:"zip_archive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Protobuf) Reset()         { *m = Protobuf{} }
func (m *Protobuf) String() string { return proto.CompactTextString(m) }
func (*Protobuf) ProtoMessage()    {}
func (*Protobuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *Protobuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Protobuf.Unmarshal(m, b)
}
func (m *Protobuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Protobuf.Marshal(b, m, deterministic)
}
func (m *Protobuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Protobuf.Merge(m, src)
}
func (m *Protobuf) XXX_Size() int {
	return xxx_messageInfo_Protobuf.Size(m)
}
func (m *Protobuf) XXX_DiscardUnknown() {
	xxx_messageInfo_Protobuf.DiscardUnknown(m)
}

var xxx_messageInfo_Protobuf proto.InternalMessageInfo

func (m *Protobuf) GetRootType() string {
	if m != nil {
		return m.RootType
	}
	return ""
}

func (m *Protobuf) GetZipArchive() []byte {
	if m != nil {
		return m.ZipArchive
	}
	return nil
}

type JSONSchema struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONSchema) Reset()         { *m = JSONSchema{} }
func (m *JSONSchema) String() string { return proto.CompactTextString(m) }
func (*JSONSchema) ProtoMessage()    {}
func (*JSONSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *JSONSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONSchema.Unmarshal(m, b)
}
func (m *JSONSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONSchema.Marshal(b, m, deterministic)
}
func (m *JSONSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONSchema.Merge(m, src)
}
func (m *JSONSchema) XXX_Size() int {
	return xxx_messageInfo_JSONSchema.Size(m)
}
func (m *JSONSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONSchema.DiscardUnknown(m)
}

var xxx_messageInfo_JSONSchema proto.InternalMessageInfo

type Avro struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Avro) Reset()         { *m = Avro{} }
func (m *Avro) String() string { return proto.CompactTextString(m) }
func (*Avro) ProtoMessage()    {}
func (*Avro) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *Avro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Avro.Unmarshal(m, b)
}
func (m *Avro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Avro.Marshal(b, m, deterministic)
}
func (m *Avro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Avro.Merge(m, src)
}
func (m *Avro) XXX_Size() int {
	return xxx_messageInfo_Avro.Size(m)
}
func (m *Avro) XXX_DiscardUnknown() {
	xxx_messageInfo_Avro.DiscardUnknown(m)
}

var xxx_messageInfo_Avro proto.InternalMessageInfo

type Options struct {
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.encoding.Type" json:"type,omitempty"`
	// Only filled out if "type" is not NONE or JSON
	//
	// Types that are valid to be assigned to Encoding:
	//	*Options_Protobuf
	//	*Options_Avro
	//	*Options_JsonSchema
	Encoding             isOptions_Encoding `protobuf_oneof:"Encoding"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{3}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

type isOptions_Encoding interface {
	isOptions_Encoding()
}

type Options_Protobuf struct {
	Protobuf *Protobuf `protobuf:"bytes,100,opt,name=protobuf,proto3,oneof"`
}

type Options_Avro struct {
	Avro *Avro `protobuf:"bytes,102,opt,name=avro,proto3,oneof"`
}

type Options_JsonSchema struct {
	JsonSchema *JSONSchema `protobuf:"bytes,101,opt,name=json_schema,json=jsonSchema,proto3,oneof"`
}

func (*Options_Protobuf) isOptions_Encoding() {}

func (*Options_Avro) isOptions_Encoding() {}

func (*Options_JsonSchema) isOptions_Encoding() {}

func (m *Options) GetEncoding() isOptions_Encoding {
	if m != nil {
		return m.Encoding
	}
	return nil
}

func (m *Options) GetProtobuf() *Protobuf {
	if x, ok := m.GetEncoding().(*Options_Protobuf); ok {
		return x.Protobuf
	}
	return nil
}

func (m *Options) GetAvro() *Avro {
	if x, ok := m.GetEncoding().(*Options_Avro); ok {
		return x.Avro
	}
	return nil
}

func (m *Options) GetJsonSchema() *JSONSchema {
	if x, ok := m.GetEncoding().(*Options_JsonSchema); ok {
		return x.JsonSchema
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Options) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Options_Protobuf)(nil),
		(*Options_Avro)(nil),
		(*Options_JsonSchema)(nil),
	}
}

func init() {
	proto.RegisterEnum("protos.encoding.Type", Type_name, Type_value)
	proto.RegisterType((*Protobuf)(nil), "protos.encoding.Protobuf")
	proto.RegisterType((*JSONSchema)(nil), "protos.encoding.JSONSchema")
	proto.RegisterType((*Avro)(nil), "protos.encoding.Avro")
	proto.RegisterType((*Options)(nil), "protos.encoding.Options")
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x6b, 0xe2, 0x40,
	0x18, 0x86, 0x13, 0x37, 0xb8, 0xf1, 0x8b, 0xbb, 0x86, 0x81, 0x85, 0x2c, 0x1e, 0x2a, 0x39, 0xd9,
	0x96, 0x26, 0x60, 0x0f, 0x3d, 0xb5, 0x25, 0x8a, 0x25, 0x14, 0x6a, 0x64, 0xb4, 0x3d, 0xf4, 0x12,
	0x92, 0x18, 0xcd, 0x14, 0xcd, 0x0c, 0x93, 0x44, 0xd0, 0x1f, 0xdc, 0xdf, 0x51, 0x66, 0xa2, 0x16,
	0xea, 0x29, 0x5f, 0x5e, 0x9e, 0x77, 0xbe, 0x67, 0x06, 0xfe, 0x50, 0x56, 0x12, 0x9a, 0x17, 0x0e,
	0xe3, 0xb4, 0xa4, 0xa8, 0x23, 0x3f, 0x85, 0x93, 0xe6, 0x09, 0x5d, 0x90, 0x7c, 0x65, 0xfb, 0xa0,
	0x4f, 0x45, 0x14, 0x57, 0x4b, 0xd4, 0x85, 0x16, 0xa7, 0xb4, 0x0c, 0xcb, 0x1d, 0x4b, 0x2d, 0xb5,
	0xa7, 0xf6, 0x5b, 0x58, 0x17, 0xc1, 0x7c, 0xc7, 0x52, 0x74, 0x01, 0xc6, 0x9e, 0xb0, 0x30, 0xe2,
	0x49, 0x46, 0xb6, 0xa9, 0xd5, 0xe8, 0xa9, 0xfd, 0x36, 0x86, 0x3d, 0x61, 0x5e, 0x9d, 0xd8, 0x6d,
	0x80, 0xe7, 0x59, 0x30, 0x99, 0x25, 0x59, 0xba, 0x89, 0xec, 0x26, 0x68, 0xde, 0x96, 0x53, 0xfb,
	0x53, 0x85, 0xdf, 0x41, 0xad, 0x80, 0x2e, 0x41, 0x3b, 0x1d, 0xfd, 0x77, 0xf0, 0xcf, 0xf9, 0xe1,
	0xe2, 0x88, 0x3d, 0x58, 0x22, 0xe8, 0x0e, 0x74, 0x76, 0xd0, 0xb2, 0x16, 0x3d, 0xb5, 0x6f, 0x0c,
	0xfe, 0x9f, 0xe1, 0x47, 0x6f, 0x5f, 0xc1, 0x27, 0x18, 0x5d, 0x83, 0x16, 0x6d, 0x39, 0xb5, 0x96,
	0xb2, 0x74, 0xbe, 0x43, 0x48, 0xf9, 0x0a, 0x96, 0x10, 0x7a, 0x00, 0xe3, 0xa3, 0xa0, 0x79, 0x58,
	0x48, 0x67, 0x2b, 0x95, 0x9d, 0xee, 0x59, 0xe7, 0xfb, 0x5a, 0xbe, 0x82, 0x41, 0x34, 0xea, 0xbf,
	0x21, 0x80, 0x3e, 0x3e, 0x40, 0x57, 0x23, 0xd0, 0xe4, 0x3b, 0xe9, 0xa0, 0x4d, 0x82, 0xc9, 0xd8,
	0x54, 0xc4, 0x24, 0x9a, 0xa6, 0x8a, 0x3a, 0x60, 0x88, 0x29, 0x9c, 0x8d, 0xfc, 0xf1, 0x8b, 0x67,
	0x36, 0x50, 0x1b, 0xf4, 0x29, 0x0e, 0xe6, 0xc1, 0xf0, 0xf5, 0xc9, 0xfc, 0x25, 0x40, 0xef, 0x0d,
	0x07, 0xa6, 0x36, 0x7c, 0x7c, 0xbf, 0x5f, 0x91, 0x32, 0xab, 0x62, 0x27, 0xa1, 0x1b, 0x37, 0x8e,
	0xca, 0x24, 0x4b, 0x28, 0x67, 0x2e, 0x5b, 0x57, 0x9b, 0x38, 0xe5, 0x37, 0xb5, 0x68, 0xe1, 0xc6,
	0x15, 0x59, 0x2f, 0xdc, 0x15, 0x75, 0x6b, 0x55, 0xf7, 0xa8, 0x1a, 0x37, 0x65, 0x70, 0xfb, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x39, 0xf8, 0x9d, 0x3f, 0xf7, 0x01, 0x00, 0x00,
}