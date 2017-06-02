// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commandProto.proto

/*
Package commandProto is a generated protocol buffer package.

It is generated from these files:
	commandProto.proto

It has these top-level messages:
	Empty
	AddRequest
	DelRequest
	ErrorReply
	VersionReply
	InfoReply
	AddReply
	DelReply
*/
package commandProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty for Null request
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AddRequest for ADD command input
type AddRequest struct {
	Version   string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Request   string `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=requestId" json:"requestId,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AddRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *AddRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// DelRequest for DEL command input
type DelRequest struct {
	Version   string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Request   string `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=requestId" json:"requestId,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DelRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DelRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *DelRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// ErrorReply for error message output
type ErrorReply struct {
	CdiVersion string `protobuf:"bytes,1,opt,name=cdiVersion" json:"cdiVersion,omitempty"`
	Code       int32  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Msg        string `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
	Details    string `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
}

func (m *ErrorReply) Reset()                    { *m = ErrorReply{} }
func (m *ErrorReply) String() string            { return proto.CompactTextString(m) }
func (*ErrorReply) ProtoMessage()               {}
func (*ErrorReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ErrorReply) GetCdiVersion() string {
	if m != nil {
		return m.CdiVersion
	}
	return ""
}

func (m *ErrorReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ErrorReply) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

// VersionReply for VERSION command output
type VersionReply struct {
	CdiVersion        string   `protobuf:"bytes,1,opt,name=cdiVersion" json:"cdiVersion,omitempty"`
	SupportedVersions []string `protobuf:"bytes,2,rep,name=supportedVersions" json:"supportedVersions,omitempty"`
}

func (m *VersionReply) Reset()                    { *m = VersionReply{} }
func (m *VersionReply) String() string            { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()               {}
func (*VersionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VersionReply) GetCdiVersion() string {
	if m != nil {
		return m.CdiVersion
	}
	return ""
}

func (m *VersionReply) GetSupportedVersions() []string {
	if m != nil {
		return m.SupportedVersions
	}
	return nil
}

// InfoReply for REPLY command output
type InfoReply struct {
	CdiVersion string      `protobuf:"bytes,1,opt,name=cdiVersion" json:"cdiVersion,omitempty"`
	Gpu        int32       `protobuf:"varint,2,opt,name=gpu" json:"gpu,omitempty"`
	Devices    []string    `protobuf:"bytes,3,rep,name=devices" json:"devices,omitempty"`
	InfoError  *ErrorReply `protobuf:"bytes,4,opt,name=infoError" json:"infoError,omitempty"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InfoReply) GetCdiVersion() string {
	if m != nil {
		return m.CdiVersion
	}
	return ""
}

func (m *InfoReply) GetGpu() int32 {
	if m != nil {
		return m.Gpu
	}
	return 0
}

func (m *InfoReply) GetDevices() []string {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *InfoReply) GetInfoError() *ErrorReply {
	if m != nil {
		return m.InfoError
	}
	return nil
}

// AddReply for ADD command output
type AddReply struct {
	CdiVersion string      `protobuf:"bytes,1,opt,name=cdiVersion" json:"cdiVersion,omitempty"`
	Devices    []string    `protobuf:"bytes,2,rep,name=devices" json:"devices,omitempty"`
	AddError   *ErrorReply `protobuf:"bytes,3,opt,name=addError" json:"addError,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddReply) GetCdiVersion() string {
	if m != nil {
		return m.CdiVersion
	}
	return ""
}

func (m *AddReply) GetDevices() []string {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *AddReply) GetAddError() *ErrorReply {
	if m != nil {
		return m.AddError
	}
	return nil
}

// DelReply for DEL command output
type DelReply struct {
	CdiVersion string      `protobuf:"bytes,1,opt,name=cdiVersion" json:"cdiVersion,omitempty"`
	DelError   *ErrorReply `protobuf:"bytes,2,opt,name=delError" json:"delError,omitempty"`
}

func (m *DelReply) Reset()                    { *m = DelReply{} }
func (m *DelReply) String() string            { return proto.CompactTextString(m) }
func (*DelReply) ProtoMessage()               {}
func (*DelReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DelReply) GetCdiVersion() string {
	if m != nil {
		return m.CdiVersion
	}
	return ""
}

func (m *DelReply) GetDelError() *ErrorReply {
	if m != nil {
		return m.DelError
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "commandProto.Empty")
	proto.RegisterType((*AddRequest)(nil), "commandProto.AddRequest")
	proto.RegisterType((*DelRequest)(nil), "commandProto.DelRequest")
	proto.RegisterType((*ErrorReply)(nil), "commandProto.ErrorReply")
	proto.RegisterType((*VersionReply)(nil), "commandProto.VersionReply")
	proto.RegisterType((*InfoReply)(nil), "commandProto.InfoReply")
	proto.RegisterType((*AddReply)(nil), "commandProto.AddReply")
	proto.RegisterType((*DelReply)(nil), "commandProto.DelReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CmdProto service

type CmdProtoClient interface {
	// Sends command reply
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionReply, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error)
}

type cmdProtoClient struct {
	cc *grpc.ClientConn
}

func NewCmdProtoClient(cc *grpc.ClientConn) CmdProtoClient {
	return &cmdProtoClient{cc}
}

func (c *cmdProtoClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := grpc.Invoke(ctx, "/commandProto.CmdProto/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdProtoClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/commandProto.CmdProto/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdProtoClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/commandProto.CmdProto/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdProtoClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelReply, error) {
	out := new(DelReply)
	err := grpc.Invoke(ctx, "/commandProto.CmdProto/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CmdProto service

type CmdProtoServer interface {
	// Sends command reply
	Version(context.Context, *Empty) (*VersionReply, error)
	Info(context.Context, *Empty) (*InfoReply, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
	Del(context.Context, *DelRequest) (*DelReply, error)
}

func RegisterCmdProtoServer(s *grpc.Server, srv CmdProtoServer) {
	s.RegisterService(&_CmdProto_serviceDesc, srv)
}

func _CmdProto_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdProtoServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commandProto.CmdProto/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdProtoServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdProto_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdProtoServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commandProto.CmdProto/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdProtoServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdProto_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdProtoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commandProto.CmdProto/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdProtoServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdProto_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdProtoServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commandProto.CmdProto/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdProtoServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmdProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commandProto.CmdProto",
	HandlerType: (*CmdProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _CmdProto_Version_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _CmdProto_Info_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _CmdProto_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _CmdProto_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commandProto.proto",
}

func init() { proto.RegisterFile("commandProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6b, 0xea, 0x40,
	0x10, 0x36, 0x89, 0x3e, 0x93, 0x79, 0x1e, 0x7c, 0xfb, 0xa0, 0x5d, 0xa4, 0x14, 0xd9, 0x93, 0x87,
	0xe2, 0xc1, 0x8a, 0x50, 0xe8, 0x45, 0xaa, 0x07, 0x6f, 0x25, 0x87, 0x9e, 0x4a, 0xa9, 0xcd, 0xae,
	0x12, 0x48, 0xb2, 0x69, 0x12, 0x05, 0xfb, 0x4f, 0xf4, 0xbf, 0xed, 0xb9, 0xec, 0x8f, 0x98, 0x1f,
	0x15, 0xcc, 0xa5, 0x97, 0x30, 0xf3, 0xcd, 0xcc, 0xf7, 0xcd, 0xee, 0xce, 0x04, 0x90, 0xc7, 0xc3,
	0x70, 0x1d, 0xd1, 0xc7, 0x84, 0x67, 0x7c, 0x1c, 0x8b, 0x2f, 0xea, 0x95, 0x31, 0xd2, 0x85, 0xce,
	0x32, 0x8c, 0xb3, 0x03, 0x79, 0x01, 0x98, 0x53, 0xea, 0xb2, 0xf7, 0x1d, 0x4b, 0x33, 0x84, 0xa1,
	0xbb, 0x67, 0x49, 0xea, 0xf3, 0x08, 0x1b, 0x43, 0x63, 0xe4, 0xb8, 0xb9, 0x2b, 0x22, 0x89, 0x4a,
	0xc2, 0xa6, 0x8a, 0x68, 0x17, 0x5d, 0x81, 0xa3, 0xcd, 0x15, 0xc5, 0x96, 0x8c, 0x15, 0x80, 0xe0,
	0x5f, 0xb0, 0xe0, 0xf7, 0xf8, 0x03, 0x80, 0x65, 0x92, 0xf0, 0xc4, 0x65, 0x71, 0x70, 0x40, 0xd7,
	0x00, 0x1e, 0xf5, 0x9f, 0x2a, 0x12, 0x25, 0x04, 0x21, 0x68, 0x7b, 0x9c, 0x32, 0x29, 0xd1, 0x71,
	0xa5, 0x8d, 0xfa, 0x60, 0x85, 0xe9, 0x56, 0x33, 0x0b, 0x53, 0xf4, 0x42, 0x59, 0xb6, 0xf6, 0x83,
	0x14, 0xb7, 0x55, 0x2f, 0xda, 0x25, 0xcf, 0xd0, 0xd3, 0x54, 0xcd, 0xf4, 0x6e, 0xe0, 0x5f, 0xba,
	0x8b, 0x63, 0x9e, 0x64, 0x8c, 0x6a, 0x2c, 0xc5, 0xe6, 0xd0, 0x1a, 0x39, 0xee, 0xcf, 0x00, 0xf9,
	0x34, 0xc0, 0x59, 0x45, 0x1b, 0xde, 0x8c, 0xbb, 0x0f, 0xd6, 0x36, 0xde, 0xe9, 0xa3, 0x08, 0x53,
	0xf5, 0xbd, 0xf7, 0x3d, 0x96, 0x62, 0x4b, 0x6a, 0xe4, 0x2e, 0x9a, 0x81, 0xe3, 0x47, 0x1b, 0x2e,
	0x6f, 0x4a, 0x9e, 0xe9, 0xef, 0x04, 0x8f, 0x2b, 0x43, 0x52, 0x5c, 0xa2, 0x5b, 0xa4, 0x92, 0x0f,
	0xb0, 0xe5, 0x74, 0x34, 0xe9, 0xa7, 0xa4, 0x6e, 0x56, 0xd5, 0xa7, 0x60, 0xaf, 0x29, 0x55, 0xe2,
	0xd6, 0x19, 0xf1, 0x63, 0x26, 0x79, 0x05, 0x5b, 0x4e, 0x4e, 0x13, 0xed, 0x29, 0xd8, 0x94, 0x05,
	0x4a, 0xc1, 0x3c, 0xa7, 0x90, 0x67, 0x4e, 0xbe, 0x0c, 0xb0, 0x1f, 0x42, 0x95, 0x81, 0xee, 0xa1,
	0x9b, 0xb3, 0xfd, 0xaf, 0xd5, 0x8a, 0x45, 0x19, 0x0c, 0xaa, 0x60, 0x79, 0x0c, 0x48, 0x0b, 0xcd,
	0xa0, 0x2d, 0x5e, 0xee, 0x74, 0xe9, 0x65, 0x15, 0x3c, 0x3e, 0x31, 0x69, 0xa1, 0x3b, 0xb0, 0xe6,
	0x94, 0xa2, 0x5a, 0xb7, 0xc5, 0x46, 0x0e, 0x2e, 0x4e, 0x44, 0x8e, 0xa5, 0x0b, 0x16, 0xd4, 0x4b,
	0x8b, 0x65, 0xab, 0x97, 0xe6, 0x97, 0x49, 0x5a, 0x6f, 0x7f, 0xe4, 0x2f, 0xe1, 0xf6, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xa6, 0x10, 0x4c, 0x28, 0x04, 0x00, 0x00,
}
