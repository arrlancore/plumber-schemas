// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: opts/ps_opts_manage.proto

package opts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ManageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"embed"
	GlobalOptions *GlobalManageOptions `protobuf:"bytes,1,opt,name=global_options,json=globalOptions,proto3" json:"global_options,omitempty" kong:"embed"`
	// @gotags: kong:"cmd,help='Get resource(s) from plumber server'"
	Get *GetOptions `protobuf:"bytes,2,opt,name=get,proto3" json:"get,omitempty" kong:"cmd,help='Get resource(s) from plumber server'"`
	// @gotags: kong:"cmd,help='Create resources in plumber server'"
	Create *CreateOptions `protobuf:"bytes,3,opt,name=create,proto3" json:"create,omitempty" kong:"cmd,help='Create resources in plumber server'"`
	// @gotags: kong:"cmd,help='Delete resources in plumber server'"
	Delete *DeleteOptions `protobuf:"bytes,5,opt,name=delete,proto3" json:"delete,omitempty" kong:"cmd,help='Delete resources in plumber server'"`
	// @gotags: kong:"cmd,help='Stop resources in plumber server'"
	Stop *StopOptions `protobuf:"bytes,6,opt,name=stop,proto3" json:"stop,omitempty" kong:"cmd,help='Stop resources in plumber server'"`
	// @gotags: kong:"cmd,help='Resume/Start resources in plumber server'"
	Resume *ResumeOptions `protobuf:"bytes,7,opt,name=resume,proto3" json:"resume,omitempty" kong:"cmd,help='Resume/Start resources in plumber server'"`
}

func (x *ManageOptions) Reset() {
	*x = ManageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageOptions) ProtoMessage() {}

func (x *ManageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageOptions.ProtoReflect.Descriptor instead.
func (*ManageOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{0}
}

func (x *ManageOptions) GetGlobalOptions() *GlobalManageOptions {
	if x != nil {
		return x.GlobalOptions
	}
	return nil
}

func (x *ManageOptions) GetGet() *GetOptions {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *ManageOptions) GetCreate() *CreateOptions {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *ManageOptions) GetDelete() *DeleteOptions {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *ManageOptions) GetStop() *StopOptions {
	if x != nil {
		return x.Stop
	}
	return nil
}

func (x *ManageOptions) GetResume() *ResumeOptions {
	if x != nil {
		return x.Resume
	}
	return nil
}

type GlobalManageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Plumber server gRPC API address',default='localhost:9090'"
	ManageAddress string `protobuf:"bytes,1,opt,name=manage_address,json=manageAddress,proto3" json:"manage_address,omitempty" kong:"help='Plumber server gRPC API address',default='localhost:9090'"`
	// @gotags: kong:"help='Plumber server auth token',default='plumber-token'"
	ManageToken string `protobuf:"bytes,2,opt,name=manage_token,json=manageToken,proto3" json:"manage_token,omitempty" kong:"help='Plumber server auth token',default='plumber-token'"`
	// @gotags: kong:"help='gRPC call timeout seconds',default=10"
	ManageTimeoutSeconds int64 `protobuf:"varint,3,opt,name=manage_timeout_seconds,json=manageTimeoutSeconds,proto3" json:"manage_timeout_seconds,omitempty" kong:"help='gRPC call timeout seconds',default=10"`
	// @gotags: kong:"help='Use TLS when talking to plumber server',default='false'"
	ManageUseTls bool `protobuf:"varint,4,opt,name=manage_use_tls,json=manageUseTls,proto3" json:"manage_use_tls,omitempty" kong:"help='Use TLS when talking to plumber server',default='false'"`
	// @gotags: kong:"help='Skip TLS server certificate verification',default='false'"
	ManageInsecureTls bool `protobuf:"varint,5,opt,name=manage_insecure_tls,json=manageInsecureTls,proto3" json:"manage_insecure_tls,omitempty" kong:"help='Skip TLS server certificate verification',default='false'"`
	// @gotags: kong:"help='TLS CA file'"
	ManageTlsCaFile string `protobuf:"bytes,6,opt,name=manage_tls_ca_file,json=manageTlsCaFile,proto3" json:"manage_tls_ca_file,omitempty" kong:"help='TLS CA file'"`
	// @gotags: kong:"help='TLS client cert file'"
	ManageTlsCertFile string `protobuf:"bytes,7,opt,name=manage_tls_cert_file,json=manageTlsCertFile,proto3" json:"manage_tls_cert_file,omitempty" kong:"help='TLS client cert file'"`
	// @gotags: kong:"help='TLS client key file'"
	ManageTlsKeyFile string `protobuf:"bytes,8,opt,name=manage_tls_key_file,json=manageTlsKeyFile,proto3" json:"manage_tls_key_file,omitempty" kong:"help='TLS client key file'"`
	// @gotags: kong:"help='Disable pretty/colorized output',default='false'"
	DisablePretty bool `protobuf:"varint,9,opt,name=disable_pretty,json=disablePretty,proto3" json:"disable_pretty,omitempty" kong:"help='Disable pretty/colorized output',default='false'"`
}

func (x *GlobalManageOptions) Reset() {
	*x = GlobalManageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalManageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalManageOptions) ProtoMessage() {}

func (x *GlobalManageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalManageOptions.ProtoReflect.Descriptor instead.
func (*GlobalManageOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalManageOptions) GetManageAddress() string {
	if x != nil {
		return x.ManageAddress
	}
	return ""
}

func (x *GlobalManageOptions) GetManageToken() string {
	if x != nil {
		return x.ManageToken
	}
	return ""
}

func (x *GlobalManageOptions) GetManageTimeoutSeconds() int64 {
	if x != nil {
		return x.ManageTimeoutSeconds
	}
	return 0
}

func (x *GlobalManageOptions) GetManageUseTls() bool {
	if x != nil {
		return x.ManageUseTls
	}
	return false
}

func (x *GlobalManageOptions) GetManageInsecureTls() bool {
	if x != nil {
		return x.ManageInsecureTls
	}
	return false
}

func (x *GlobalManageOptions) GetManageTlsCaFile() string {
	if x != nil {
		return x.ManageTlsCaFile
	}
	return ""
}

func (x *GlobalManageOptions) GetManageTlsCertFile() string {
	if x != nil {
		return x.ManageTlsCertFile
	}
	return ""
}

func (x *GlobalManageOptions) GetManageTlsKeyFile() string {
	if x != nil {
		return x.ManageTlsKeyFile
	}
	return ""
}

func (x *GlobalManageOptions) GetDisablePretty() bool {
	if x != nil {
		return x.DisablePretty
	}
	return false
}

type GetOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"cmd,help='Get connection(s) from plumber server'"
	Connection *GetConnectionOptions `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty" kong:"cmd,help='Get connection(s) from plumber server'"`
	// @gotags: kong:"cmd,help='Get relay(s) from plumber server'"
	Relay *GetRelayOptions `protobuf:"bytes,2,opt,name=relay,proto3" json:"relay,omitempty" kong:"cmd,help='Get relay(s) from plumber server'"`
	// @gotags: kong:"cmd,help='Get tunnel(s) from plumber server'"
	Tunnel *GetTunnelOptions `protobuf:"bytes,3,opt,name=tunnel,proto3" json:"tunnel,omitempty" kong:"cmd,help='Get tunnel(s) from plumber server'"`
}

func (x *GetOptions) Reset() {
	*x = GetOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptions) ProtoMessage() {}

func (x *GetOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptions.ProtoReflect.Descriptor instead.
func (*GetOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{2}
}

func (x *GetOptions) GetConnection() *GetConnectionOptions {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *GetOptions) GetRelay() *GetRelayOptions {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *GetOptions) GetTunnel() *GetTunnelOptions {
	if x != nil {
		return x.Tunnel
	}
	return nil
}

type CreateOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"cmd,help='Create connection in plumber server'"
	Connection *CreateConnectionOptions `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty" kong:"cmd,help='Create connection in plumber server'"`
	// @gotags: kong:"cmd,help='Create relay in plumber server'"
	Relay *CreateRelayOptions `protobuf:"bytes,2,opt,name=relay,proto3" json:"relay,omitempty" kong:"cmd,help='Create relay in plumber server'"`
	// @gotags: kong:"cmd,help='Create tunnel in plumber server'"
	Tunnel *CreateTunnelOptions `protobuf:"bytes,3,opt,name=tunnel,proto3" json:"tunnel,omitempty" kong:"cmd,help='Create tunnel in plumber server'"`
}

func (x *CreateOptions) Reset() {
	*x = CreateOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOptions) ProtoMessage() {}

func (x *CreateOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOptions.ProtoReflect.Descriptor instead.
func (*CreateOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOptions) GetConnection() *CreateConnectionOptions {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *CreateOptions) GetRelay() *CreateRelayOptions {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *CreateOptions) GetTunnel() *CreateTunnelOptions {
	if x != nil {
		return x.Tunnel
	}
	return nil
}

type DeleteOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"cmd,help='Delete connection in plumber server'"
	Connection *DeleteConnectionOptions `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty" kong:"cmd,help='Delete connection in plumber server'"`
	// @gotags: kong:"cmd,help='Delete relay in plumber server'"
	Relay *DeleteRelayOptions `protobuf:"bytes,2,opt,name=relay,proto3" json:"relay,omitempty" kong:"cmd,help='Delete relay in plumber server'"`
	// @gotags: kong:"cmd,help='Delete tunnel in plumber server'"
	Tunnel *DeleteTunnelOptions `protobuf:"bytes,3,opt,name=tunnel,proto3" json:"tunnel,omitempty" kong:"cmd,help='Delete tunnel in plumber server'"`
}

func (x *DeleteOptions) Reset() {
	*x = DeleteOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOptions) ProtoMessage() {}

func (x *DeleteOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOptions.ProtoReflect.Descriptor instead.
func (*DeleteOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOptions) GetConnection() *DeleteConnectionOptions {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *DeleteOptions) GetRelay() *DeleteRelayOptions {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *DeleteOptions) GetTunnel() *DeleteTunnelOptions {
	if x != nil {
		return x.Tunnel
	}
	return nil
}

type StopOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"cmd,help='Stop relay in plumber server'"
	Relay *StopRelayOptions `protobuf:"bytes,1,opt,name=relay,proto3" json:"relay,omitempty" kong:"cmd,help='Stop relay in plumber server'"`
	// @gotags: kong:"cmd,help='Stop tunnel in plumber server'"
	Tunnel *StopTunnelOptions `protobuf:"bytes,2,opt,name=tunnel,proto3" json:"tunnel,omitempty" kong:"cmd,help='Stop tunnel in plumber server'"`
}

func (x *StopOptions) Reset() {
	*x = StopOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOptions) ProtoMessage() {}

func (x *StopOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOptions.ProtoReflect.Descriptor instead.
func (*StopOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{5}
}

func (x *StopOptions) GetRelay() *StopRelayOptions {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *StopOptions) GetTunnel() *StopTunnelOptions {
	if x != nil {
		return x.Tunnel
	}
	return nil
}

type ResumeOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"cmd,help='Resume/Start relay in plumber server'"
	Relay *ResumeRelayOptions `protobuf:"bytes,1,opt,name=relay,proto3" json:"relay,omitempty" kong:"cmd,help='Resume/Start relay in plumber server'"`
	// @gotags: kong:"cmd,help='Resume/Start tunnel in plumber server'"
	Tunnel *ResumeTunnelOptions `protobuf:"bytes,2,opt,name=tunnel,proto3" json:"tunnel,omitempty" kong:"cmd,help='Resume/Start tunnel in plumber server'"`
}

func (x *ResumeOptions) Reset() {
	*x = ResumeOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_ps_opts_manage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeOptions) ProtoMessage() {}

func (x *ResumeOptions) ProtoReflect() protoreflect.Message {
	mi := &file_opts_ps_opts_manage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeOptions.ProtoReflect.Descriptor instead.
func (*ResumeOptions) Descriptor() ([]byte, []int) {
	return file_opts_ps_opts_manage_proto_rawDescGZIP(), []int{6}
}

func (x *ResumeOptions) GetRelay() *ResumeRelayOptions {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *ResumeOptions) GetTunnel() *ResumeTunnelOptions {
	if x != nil {
		return x.Tunnel
	}
	return nil
}

var File_opts_ps_opts_manage_proto protoreflect.FileDescriptor

var file_opts_ps_opts_manage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x1a, 0x24, 0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70,
	0x73, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x03,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x32, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x9f, 0x03, 0x0a, 0x13, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6c, 0x73, 0x5f, 0x63, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x54, 0x6c, 0x73, 0x43, 0x61, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6c, 0x73,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6c,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x54, 0x6c, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x65, 0x74, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x22, 0xba, 0x01, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x05, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x35, 0x0a, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35,
	0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x70, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0xc6, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x44, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x70, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x70, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x38,
	0x0a, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x7a, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x70, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x36, 0x0a, 0x06,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x70, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x38, 0x0a,
	0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x06, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6f, 0x70, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opts_ps_opts_manage_proto_rawDescOnce sync.Once
	file_opts_ps_opts_manage_proto_rawDescData = file_opts_ps_opts_manage_proto_rawDesc
)

func file_opts_ps_opts_manage_proto_rawDescGZIP() []byte {
	file_opts_ps_opts_manage_proto_rawDescOnce.Do(func() {
		file_opts_ps_opts_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_opts_ps_opts_manage_proto_rawDescData)
	})
	return file_opts_ps_opts_manage_proto_rawDescData
}

var file_opts_ps_opts_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_opts_ps_opts_manage_proto_goTypes = []interface{}{
	(*ManageOptions)(nil),           // 0: protos.opts.ManageOptions
	(*GlobalManageOptions)(nil),     // 1: protos.opts.GlobalManageOptions
	(*GetOptions)(nil),              // 2: protos.opts.GetOptions
	(*CreateOptions)(nil),           // 3: protos.opts.CreateOptions
	(*DeleteOptions)(nil),           // 4: protos.opts.DeleteOptions
	(*StopOptions)(nil),             // 5: protos.opts.StopOptions
	(*ResumeOptions)(nil),           // 6: protos.opts.ResumeOptions
	(*GetConnectionOptions)(nil),    // 7: protos.opts.GetConnectionOptions
	(*GetRelayOptions)(nil),         // 8: protos.opts.GetRelayOptions
	(*GetTunnelOptions)(nil),        // 9: protos.opts.GetTunnelOptions
	(*CreateConnectionOptions)(nil), // 10: protos.opts.CreateConnectionOptions
	(*CreateRelayOptions)(nil),      // 11: protos.opts.CreateRelayOptions
	(*CreateTunnelOptions)(nil),     // 12: protos.opts.CreateTunnelOptions
	(*DeleteConnectionOptions)(nil), // 13: protos.opts.DeleteConnectionOptions
	(*DeleteRelayOptions)(nil),      // 14: protos.opts.DeleteRelayOptions
	(*DeleteTunnelOptions)(nil),     // 15: protos.opts.DeleteTunnelOptions
	(*StopRelayOptions)(nil),        // 16: protos.opts.StopRelayOptions
	(*StopTunnelOptions)(nil),       // 17: protos.opts.StopTunnelOptions
	(*ResumeRelayOptions)(nil),      // 18: protos.opts.ResumeRelayOptions
	(*ResumeTunnelOptions)(nil),     // 19: protos.opts.ResumeTunnelOptions
}
var file_opts_ps_opts_manage_proto_depIdxs = []int32{
	1,  // 0: protos.opts.ManageOptions.global_options:type_name -> protos.opts.GlobalManageOptions
	2,  // 1: protos.opts.ManageOptions.get:type_name -> protos.opts.GetOptions
	3,  // 2: protos.opts.ManageOptions.create:type_name -> protos.opts.CreateOptions
	4,  // 3: protos.opts.ManageOptions.delete:type_name -> protos.opts.DeleteOptions
	5,  // 4: protos.opts.ManageOptions.stop:type_name -> protos.opts.StopOptions
	6,  // 5: protos.opts.ManageOptions.resume:type_name -> protos.opts.ResumeOptions
	7,  // 6: protos.opts.GetOptions.connection:type_name -> protos.opts.GetConnectionOptions
	8,  // 7: protos.opts.GetOptions.relay:type_name -> protos.opts.GetRelayOptions
	9,  // 8: protos.opts.GetOptions.tunnel:type_name -> protos.opts.GetTunnelOptions
	10, // 9: protos.opts.CreateOptions.connection:type_name -> protos.opts.CreateConnectionOptions
	11, // 10: protos.opts.CreateOptions.relay:type_name -> protos.opts.CreateRelayOptions
	12, // 11: protos.opts.CreateOptions.tunnel:type_name -> protos.opts.CreateTunnelOptions
	13, // 12: protos.opts.DeleteOptions.connection:type_name -> protos.opts.DeleteConnectionOptions
	14, // 13: protos.opts.DeleteOptions.relay:type_name -> protos.opts.DeleteRelayOptions
	15, // 14: protos.opts.DeleteOptions.tunnel:type_name -> protos.opts.DeleteTunnelOptions
	16, // 15: protos.opts.StopOptions.relay:type_name -> protos.opts.StopRelayOptions
	17, // 16: protos.opts.StopOptions.tunnel:type_name -> protos.opts.StopTunnelOptions
	18, // 17: protos.opts.ResumeOptions.relay:type_name -> protos.opts.ResumeRelayOptions
	19, // 18: protos.opts.ResumeOptions.tunnel:type_name -> protos.opts.ResumeTunnelOptions
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_opts_ps_opts_manage_proto_init() }
func file_opts_ps_opts_manage_proto_init() {
	if File_opts_ps_opts_manage_proto != nil {
		return
	}
	file_opts_ps_opts_manage_connection_proto_init()
	file_opts_ps_opts_manage_relay_proto_init()
	file_opts_ps_opts_manage_tunnel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_opts_ps_opts_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalManageOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_opts_ps_opts_manage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opts_ps_opts_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opts_ps_opts_manage_proto_goTypes,
		DependencyIndexes: file_opts_ps_opts_manage_proto_depIdxs,
		MessageInfos:      file_opts_ps_opts_manage_proto_msgTypes,
	}.Build()
	File_opts_ps_opts_manage_proto = out.File
	file_opts_ps_opts_manage_proto_rawDesc = nil
	file_opts_ps_opts_manage_proto_goTypes = nil
	file_opts_ps_opts_manage_proto_depIdxs = nil
}
