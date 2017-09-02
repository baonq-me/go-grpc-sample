// Code generated by protoc-gen-go.
// source: sadmin/sadmin.proto
// DO NOT EDIT!

/*
Package sadmin is a generated protocol buffer package.

It is generated from these files:
	sadmin/sadmin.proto

It has these top-level messages:
	Empty
	VersionInfoResponse
	CreateAgencyStaffRequest
	CreateAgencyStaffResponse
	CreateServiceProviderRequest
	CreateServiceProviderResponse
	CreateServiceRequest
	CreateServiceResponse
	GenerateServiceProviderSecretRequest
	GenerateServiceProviderSecretResponse
	AddWebhookRequest
	AddWebhookResponse
	RemoveWebhookRequest
	RemoveWebhookResponse
*/
package sadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import agency "github.com/ng-vu/go-grpc-sample/pb/agency"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VersionInfoResponse struct {
	Service     string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	UpdatedTime int64  `protobuf:"varint,3,opt,name=updated_time,json=updatedTime" json:"updated_time,omitempty"`
}

func (m *VersionInfoResponse) Reset()                    { *m = VersionInfoResponse{} }
func (m *VersionInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionInfoResponse) ProtoMessage()               {}
func (*VersionInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionInfoResponse) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *VersionInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionInfoResponse) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type CreateAgencyStaffRequest struct {
	Info     *agency.AgencyStaff `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Password string              `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateAgencyStaffRequest) Reset()                    { *m = CreateAgencyStaffRequest{} }
func (m *CreateAgencyStaffRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateAgencyStaffRequest) ProtoMessage()               {}
func (*CreateAgencyStaffRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateAgencyStaffRequest) GetInfo() *agency.AgencyStaff {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CreateAgencyStaffRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateAgencyStaffResponse struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *CreateAgencyStaffResponse) Reset()                    { *m = CreateAgencyStaffResponse{} }
func (m *CreateAgencyStaffResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateAgencyStaffResponse) ProtoMessage()               {}
func (*CreateAgencyStaffResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateAgencyStaffResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CreateServiceProviderRequest struct {
	Codename string `protobuf:"bytes,1,opt,name=codename" json:"codename,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CreateServiceProviderRequest) Reset()                    { *m = CreateServiceProviderRequest{} }
func (m *CreateServiceProviderRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceProviderRequest) ProtoMessage()               {}
func (*CreateServiceProviderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateServiceProviderRequest) GetCodename() string {
	if m != nil {
		return m.Codename
	}
	return ""
}

func (m *CreateServiceProviderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateServiceProviderResponse struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
}

func (m *CreateServiceProviderResponse) Reset()                    { *m = CreateServiceProviderResponse{} }
func (m *CreateServiceProviderResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceProviderResponse) ProtoMessage()               {}
func (*CreateServiceProviderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateServiceProviderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateServiceProviderResponse) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

type CreateServiceRequest struct {
}

func (m *CreateServiceRequest) Reset()                    { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()               {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type CreateServiceResponse struct {
}

func (m *CreateServiceResponse) Reset()                    { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()               {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GenerateServiceProviderSecretRequest struct {
	Codename string `protobuf:"bytes,1,opt,name=codename" json:"codename,omitempty"`
}

func (m *GenerateServiceProviderSecretRequest) Reset()         { *m = GenerateServiceProviderSecretRequest{} }
func (m *GenerateServiceProviderSecretRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateServiceProviderSecretRequest) ProtoMessage()    {}
func (*GenerateServiceProviderSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8}
}

func (m *GenerateServiceProviderSecretRequest) GetCodename() string {
	if m != nil {
		return m.Codename
	}
	return ""
}

type GenerateServiceProviderSecretResponse struct {
}

func (m *GenerateServiceProviderSecretResponse) Reset()         { *m = GenerateServiceProviderSecretResponse{} }
func (m *GenerateServiceProviderSecretResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateServiceProviderSecretResponse) ProtoMessage()    {}
func (*GenerateServiceProviderSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9}
}

type AddWebhookRequest struct {
}

func (m *AddWebhookRequest) Reset()                    { *m = AddWebhookRequest{} }
func (m *AddWebhookRequest) String() string            { return proto.CompactTextString(m) }
func (*AddWebhookRequest) ProtoMessage()               {}
func (*AddWebhookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type AddWebhookResponse struct {
}

func (m *AddWebhookResponse) Reset()                    { *m = AddWebhookResponse{} }
func (m *AddWebhookResponse) String() string            { return proto.CompactTextString(m) }
func (*AddWebhookResponse) ProtoMessage()               {}
func (*AddWebhookResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type RemoveWebhookRequest struct {
}

func (m *RemoveWebhookRequest) Reset()                    { *m = RemoveWebhookRequest{} }
func (m *RemoveWebhookRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveWebhookRequest) ProtoMessage()               {}
func (*RemoveWebhookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type RemoveWebhookResponse struct {
}

func (m *RemoveWebhookResponse) Reset()                    { *m = RemoveWebhookResponse{} }
func (m *RemoveWebhookResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveWebhookResponse) ProtoMessage()               {}
func (*RemoveWebhookResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Empty)(nil), "sadmin.Empty")
	proto.RegisterType((*VersionInfoResponse)(nil), "sadmin.VersionInfoResponse")
	proto.RegisterType((*CreateAgencyStaffRequest)(nil), "sadmin.CreateAgencyStaffRequest")
	proto.RegisterType((*CreateAgencyStaffResponse)(nil), "sadmin.CreateAgencyStaffResponse")
	proto.RegisterType((*CreateServiceProviderRequest)(nil), "sadmin.CreateServiceProviderRequest")
	proto.RegisterType((*CreateServiceProviderResponse)(nil), "sadmin.CreateServiceProviderResponse")
	proto.RegisterType((*CreateServiceRequest)(nil), "sadmin.CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "sadmin.CreateServiceResponse")
	proto.RegisterType((*GenerateServiceProviderSecretRequest)(nil), "sadmin.GenerateServiceProviderSecretRequest")
	proto.RegisterType((*GenerateServiceProviderSecretResponse)(nil), "sadmin.GenerateServiceProviderSecretResponse")
	proto.RegisterType((*AddWebhookRequest)(nil), "sadmin.AddWebhookRequest")
	proto.RegisterType((*AddWebhookResponse)(nil), "sadmin.AddWebhookResponse")
	proto.RegisterType((*RemoveWebhookRequest)(nil), "sadmin.RemoveWebhookRequest")
	proto.RegisterType((*RemoveWebhookResponse)(nil), "sadmin.RemoveWebhookResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlueSAdmin service

type BlueSAdminClient interface {
	VersionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfoResponse, error)
	CreateAgencyStaff(ctx context.Context, in *CreateAgencyStaffRequest, opts ...grpc.CallOption) (*CreateAgencyStaffResponse, error)
	CreateServiceProvider(ctx context.Context, in *CreateServiceProviderRequest, opts ...grpc.CallOption) (*CreateServiceProviderResponse, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	GenerateServiceProviderSecret(ctx context.Context, in *GenerateServiceProviderSecretRequest, opts ...grpc.CallOption) (*GenerateServiceProviderSecretResponse, error)
	AddWebhook(ctx context.Context, in *AddWebhookRequest, opts ...grpc.CallOption) (*AddWebhookResponse, error)
	RemoveWebhook(ctx context.Context, in *RemoveWebhookRequest, opts ...grpc.CallOption) (*RemoveWebhookResponse, error)
}

type blueSAdminClient struct {
	cc *grpc.ClientConn
}

func NewBlueSAdminClient(cc *grpc.ClientConn) BlueSAdminClient {
	return &blueSAdminClient{cc}
}

func (c *blueSAdminClient) VersionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionInfoResponse, error) {
	out := new(VersionInfoResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/VersionInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) CreateAgencyStaff(ctx context.Context, in *CreateAgencyStaffRequest, opts ...grpc.CallOption) (*CreateAgencyStaffResponse, error) {
	out := new(CreateAgencyStaffResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/CreateAgencyStaff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) CreateServiceProvider(ctx context.Context, in *CreateServiceProviderRequest, opts ...grpc.CallOption) (*CreateServiceProviderResponse, error) {
	out := new(CreateServiceProviderResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/CreateServiceProvider", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/CreateService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) GenerateServiceProviderSecret(ctx context.Context, in *GenerateServiceProviderSecretRequest, opts ...grpc.CallOption) (*GenerateServiceProviderSecretResponse, error) {
	out := new(GenerateServiceProviderSecretResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/GenerateServiceProviderSecret", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) AddWebhook(ctx context.Context, in *AddWebhookRequest, opts ...grpc.CallOption) (*AddWebhookResponse, error) {
	out := new(AddWebhookResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/AddWebhook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSAdminClient) RemoveWebhook(ctx context.Context, in *RemoveWebhookRequest, opts ...grpc.CallOption) (*RemoveWebhookResponse, error) {
	out := new(RemoveWebhookResponse)
	err := grpc.Invoke(ctx, "/sadmin.BlueSAdmin/RemoveWebhook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlueSAdmin service

type BlueSAdminServer interface {
	VersionInfo(context.Context, *Empty) (*VersionInfoResponse, error)
	CreateAgencyStaff(context.Context, *CreateAgencyStaffRequest) (*CreateAgencyStaffResponse, error)
	CreateServiceProvider(context.Context, *CreateServiceProviderRequest) (*CreateServiceProviderResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	GenerateServiceProviderSecret(context.Context, *GenerateServiceProviderSecretRequest) (*GenerateServiceProviderSecretResponse, error)
	AddWebhook(context.Context, *AddWebhookRequest) (*AddWebhookResponse, error)
	RemoveWebhook(context.Context, *RemoveWebhookRequest) (*RemoveWebhookResponse, error)
}

func RegisterBlueSAdminServer(s *grpc.Server, srv BlueSAdminServer) {
	s.RegisterService(&_BlueSAdmin_serviceDesc, srv)
}

func _BlueSAdmin_VersionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).VersionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/VersionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).VersionInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_CreateAgencyStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgencyStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).CreateAgencyStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/CreateAgencyStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).CreateAgencyStaff(ctx, req.(*CreateAgencyStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_CreateServiceProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).CreateServiceProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/CreateServiceProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).CreateServiceProvider(ctx, req.(*CreateServiceProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_GenerateServiceProviderSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateServiceProviderSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).GenerateServiceProviderSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/GenerateServiceProviderSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).GenerateServiceProviderSecret(ctx, req.(*GenerateServiceProviderSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_AddWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).AddWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/AddWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).AddWebhook(ctx, req.(*AddWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSAdmin_RemoveWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSAdminServer).RemoveWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sadmin.BlueSAdmin/RemoveWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSAdminServer).RemoveWebhook(ctx, req.(*RemoveWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlueSAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sadmin.BlueSAdmin",
	HandlerType: (*BlueSAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VersionInfo",
			Handler:    _BlueSAdmin_VersionInfo_Handler,
		},
		{
			MethodName: "CreateAgencyStaff",
			Handler:    _BlueSAdmin_CreateAgencyStaff_Handler,
		},
		{
			MethodName: "CreateServiceProvider",
			Handler:    _BlueSAdmin_CreateServiceProvider_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _BlueSAdmin_CreateService_Handler,
		},
		{
			MethodName: "GenerateServiceProviderSecret",
			Handler:    _BlueSAdmin_GenerateServiceProviderSecret_Handler,
		},
		{
			MethodName: "AddWebhook",
			Handler:    _BlueSAdmin_AddWebhook_Handler,
		},
		{
			MethodName: "RemoveWebhook",
			Handler:    _BlueSAdmin_RemoveWebhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sadmin/sadmin.proto",
}

func init() { proto.RegisterFile("sadmin/sadmin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0xd3, 0xfe, 0x92, 0x76, 0xfa, 0x2b, 0xa8, 0x9b, 0xb6, 0x49, 0x4d, 0x83, 0xd2, 0x55,
	0xab, 0x96, 0x42, 0x62, 0x54, 0x38, 0xf5, 0xd6, 0x22, 0x04, 0x05, 0x09, 0xa1, 0x04, 0x81, 0xc4,
	0x25, 0xda, 0xc4, 0x13, 0x77, 0xd5, 0x78, 0xd7, 0xd8, 0x4e, 0x50, 0x2e, 0x1c, 0x78, 0x00, 0x2e,
	0xbc, 0x07, 0x2f, 0xc3, 0x9d, 0x13, 0x0f, 0x82, 0xbc, 0x7f, 0x9a, 0x26, 0x71, 0xff, 0x9c, 0x92,
	0xd9, 0xf9, 0x76, 0xbe, 0x6f, 0x3c, 0xdf, 0x68, 0xa1, 0x9c, 0x30, 0x3f, 0xe4, 0xc2, 0xd3, 0x3f,
	0xcd, 0x28, 0x96, 0xa9, 0x24, 0x45, 0x1d, 0xb9, 0x84, 0x05, 0x28, 0x7a, 0x63, 0x2f, 0x94, 0x3e,
	0x0e, 0x74, 0xce, 0xdd, 0x0e, 0xa4, 0x0c, 0x06, 0xe8, 0xb1, 0x88, 0x7b, 0x4c, 0x08, 0x99, 0xb2,
	0x94, 0x4b, 0x91, 0xe8, 0x2c, 0x2d, 0xc1, 0x7f, 0x2f, 0xc3, 0x28, 0x1d, 0xd3, 0x01, 0x94, 0x3f,
	0x62, 0x9c, 0x70, 0x29, 0xce, 0x44, 0x5f, 0xb6, 0x30, 0x89, 0xa4, 0x48, 0x90, 0x54, 0xa1, 0x94,
	0x60, 0x3c, 0xe2, 0x3d, 0xac, 0x3a, 0x75, 0xe7, 0x60, 0xb9, 0x65, 0xc3, 0x2c, 0x33, 0xd2, 0x17,
	0xaa, 0x05, 0x9d, 0x31, 0x21, 0xd9, 0x81, 0xff, 0x87, 0x91, 0xcf, 0x52, 0xf4, 0x3b, 0x29, 0x0f,
	0xb1, 0xba, 0x50, 0x77, 0x0e, 0x16, 0x5a, 0x2b, 0xe6, 0xec, 0x03, 0x0f, 0x91, 0x76, 0xa0, 0xfa,
	0x22, 0x46, 0x96, 0xe2, 0x89, 0x12, 0xdc, 0x4e, 0x59, 0xbf, 0xdf, 0xc2, 0x2f, 0x43, 0x4c, 0x52,
	0xb2, 0x0f, 0x8b, 0x5c, 0xf4, 0xa5, 0xe2, 0x5b, 0x39, 0x2a, 0x37, 0x75, 0x4f, 0xcd, 0xab, 0x48,
	0x05, 0x20, 0x2e, 0x2c, 0x45, 0x2c, 0x49, 0xbe, 0xca, 0xd8, 0x37, 0x12, 0x2e, 0x63, 0xfa, 0x1c,
	0xb6, 0x72, 0x08, 0x4c, 0x53, 0x15, 0x28, 0x0d, 0x13, 0x8c, 0x3b, 0xdc, 0x37, 0x4d, 0x15, 0xb3,
	0xf0, 0xcc, 0xa7, 0xef, 0x60, 0x5b, 0xdf, 0x6a, 0xeb, 0x26, 0xdf, 0xc7, 0x72, 0xc4, 0x7d, 0x8c,
	0xad, 0x34, 0x17, 0x96, 0x7a, 0xd2, 0x47, 0xc1, 0x42, 0xfb, 0x39, 0x2e, 0x63, 0x42, 0x60, 0x51,
	0x9d, 0x6b, 0x25, 0xea, 0x3f, 0x7d, 0x0d, 0xb5, 0x6b, 0xea, 0x19, 0x25, 0xf7, 0xa0, 0x70, 0x29,
	0xa2, 0xc0, 0xfd, 0x4c, 0x19, 0x8b, 0x78, 0xe7, 0x02, 0xc7, 0xa6, 0x4e, 0x91, 0x45, 0xfc, 0x2d,
	0x8e, 0xe9, 0x26, 0xac, 0x4f, 0x55, 0x32, 0x8a, 0x68, 0x05, 0x36, 0x66, 0xce, 0x75, 0x65, 0x7a,
	0x0a, 0xbb, 0xaf, 0x50, 0x60, 0x3c, 0x4f, 0xde, 0xc6, 0x5e, 0x8c, 0xe9, 0x1d, 0x5a, 0xa2, 0xfb,
	0xb0, 0x77, 0x4b, 0x0d, 0x43, 0x56, 0x86, 0xb5, 0x13, 0xdf, 0xff, 0x84, 0xdd, 0x73, 0x29, 0x2f,
	0xac, 0xb4, 0x75, 0x20, 0x57, 0x0f, 0x0d, 0x74, 0x13, 0xd6, 0x5b, 0x18, 0xca, 0x11, 0xce, 0xa0,
	0x2b, 0xb0, 0x31, 0x73, 0xae, 0x2f, 0x1c, 0xfd, 0x29, 0x02, 0x9c, 0x0e, 0x86, 0xd8, 0x3e, 0xc9,
	0x2c, 0x4e, 0xde, 0xc0, 0xca, 0x15, 0x9f, 0x92, 0xd5, 0xa6, 0x59, 0x04, 0xe5, 0x62, 0xf7, 0x81,
	0x0d, 0x73, 0xbc, 0x4c, 0xef, 0x7f, 0xff, 0xfd, 0xf7, 0x67, 0x61, 0x99, 0x94, 0xd4, 0x2e, 0x74,
	0x3a, 0xe4, 0x1b, 0xac, 0xcd, 0x99, 0x84, 0xd4, 0x6d, 0x89, 0xeb, 0x0c, 0xea, 0xee, 0xdc, 0x80,
	0x30, 0x54, 0xbb, 0x8a, 0xea, 0x21, 0xdd, 0x32, 0x54, 0xde, 0x1c, 0xf4, 0xd8, 0x39, 0x24, 0x3f,
	0x9c, 0x99, 0xe9, 0xd9, 0xcf, 0x4b, 0x76, 0xa7, 0x29, 0xf2, 0xed, 0xe8, 0xee, 0xdd, 0x82, 0x32,
	0x62, 0x0e, 0x94, 0x18, 0x4a, 0x6b, 0xd3, 0x62, 0x66, 0xe0, 0x99, 0x20, 0x01, 0xab, 0x53, 0x39,
	0xb2, 0x9d, 0xcb, 0x60, 0xf9, 0x6b, 0xd7, 0x64, 0x0d, 0x6f, 0x5d, 0xf1, 0xba, 0x74, 0x23, 0x97,
	0x37, 0xe3, 0xfb, 0xe5, 0x40, 0xed, 0x46, 0x87, 0x91, 0x27, 0x96, 0xe2, 0x2e, 0x66, 0x76, 0x1b,
	0x77, 0x44, 0x1b, 0x81, 0x4f, 0x95, 0xc0, 0x43, 0xba, 0x67, 0x05, 0xde, 0x78, 0x2d, 0x13, 0xdc,
	0x05, 0x98, 0x78, 0x9a, 0x6c, 0x59, 0xba, 0x39, 0xf3, 0xbb, 0x6e, 0x5e, 0xca, 0xd0, 0xd6, 0x14,
	0x6d, 0x85, 0x12, 0x4b, 0x3b, 0xc1, 0x98, 0x21, 0x4c, 0x6d, 0xc2, 0x64, 0x08, 0x79, 0x8b, 0x33,
	0x19, 0x42, 0xee, 0xfa, 0xcc, 0x0f, 0x61, 0x0a, 0x76, 0xec, 0x1c, 0x9e, 0x3e, 0xfe, 0xfc, 0x28,
	0xe0, 0xe9, 0xf9, 0xb0, 0xdb, 0xec, 0xc9, 0xd0, 0x13, 0x41, 0x63, 0x34, 0xf4, 0x02, 0xd9, 0x08,
	0xe2, 0xa8, 0xd7, 0x48, 0x58, 0x18, 0x0d, 0xd0, 0x8b, 0xba, 0xe6, 0xbd, 0xe9, 0x16, 0xd5, 0xb3,
	0xf1, 0xec, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xbe, 0x0f, 0x3a, 0x87, 0x06, 0x00, 0x00,
}