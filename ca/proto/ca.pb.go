// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ca.proto

It has these top-level messages:
	IssueCertificateRequest
	Certificate
	GenerateOCSPRequest
	OCSPResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IssueCertificateRequest struct {
	Csr              []byte `protobuf:"bytes,1,opt,name=csr" json:"csr,omitempty"`
	RegistrationID   *int64 `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IssueCertificateRequest) Reset()                    { *m = IssueCertificateRequest{} }
func (m *IssueCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*IssueCertificateRequest) ProtoMessage()               {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IssueCertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *IssueCertificateRequest) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

type Certificate struct {
	RegistrationID   *int64  `protobuf:"varint,1,opt,name=registrationID" json:"registrationID,omitempty"`
	Serial           *string `protobuf:"bytes,2,opt,name=serial" json:"serial,omitempty"`
	Digest           *string `protobuf:"bytes,3,opt,name=digest" json:"digest,omitempty"`
	Der              []byte  `protobuf:"bytes,4,opt,name=der" json:"der,omitempty"`
	Issued           *int64  `protobuf:"varint,5,opt,name=issued" json:"issued,omitempty"`
	Expires          *int64  `protobuf:"varint,6,opt,name=expires" json:"expires,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto1.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Certificate) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

func (m *Certificate) GetSerial() string {
	if m != nil && m.Serial != nil {
		return *m.Serial
	}
	return ""
}

func (m *Certificate) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

func (m *Certificate) GetDer() []byte {
	if m != nil {
		return m.Der
	}
	return nil
}

func (m *Certificate) GetIssued() int64 {
	if m != nil && m.Issued != nil {
		return *m.Issued
	}
	return 0
}

func (m *Certificate) GetExpires() int64 {
	if m != nil && m.Expires != nil {
		return *m.Expires
	}
	return 0
}

type GenerateOCSPRequest struct {
	CertDER          []byte  `protobuf:"bytes,1,opt,name=certDER" json:"certDER,omitempty"`
	Status           *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Reason           *int32  `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	RevokedAt        *int64  `protobuf:"varint,4,opt,name=revokedAt" json:"revokedAt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenerateOCSPRequest) Reset()                    { *m = GenerateOCSPRequest{} }
func (m *GenerateOCSPRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateOCSPRequest) ProtoMessage()               {}
func (*GenerateOCSPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GenerateOCSPRequest) GetCertDER() []byte {
	if m != nil {
		return m.CertDER
	}
	return nil
}

func (m *GenerateOCSPRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *GenerateOCSPRequest) GetReason() int32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func (m *GenerateOCSPRequest) GetRevokedAt() int64 {
	if m != nil && m.RevokedAt != nil {
		return *m.RevokedAt
	}
	return 0
}

type OCSPResponse struct {
	Response         []byte `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OCSPResponse) Reset()                    { *m = OCSPResponse{} }
func (m *OCSPResponse) String() string            { return proto1.CompactTextString(m) }
func (*OCSPResponse) ProtoMessage()               {}
func (*OCSPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OCSPResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto1.RegisterType((*IssueCertificateRequest)(nil), "ca.IssueCertificateRequest")
	proto1.RegisterType((*Certificate)(nil), "ca.Certificate")
	proto1.RegisterType((*GenerateOCSPRequest)(nil), "ca.GenerateOCSPRequest")
	proto1.RegisterType((*OCSPResponse)(nil), "ca.OCSPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CertificateAuthority service

type CertificateAuthorityClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*Certificate, error)
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type certificateAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityClient(cc *grpc.ClientConn) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*Certificate, error) {
	out := new(Certificate)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/GenerateOCSP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthority service

type CertificateAuthorityServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*Certificate, error)
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateAuthority_IssueCertificate_Handler,
		},
		{
			MethodName: "GenerateOCSP",
			Handler:    _CertificateAuthority_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0xdb, 0xe6, 0xf4, 0x36, 0x8d, 0xda, 0x1c, 0x83, 0xa8, 0x55, 0x36, 0x55, 0x56, 0xac,
	0xba, 0x60, 0x8b, 0x84, 0xd4, 0x0b, 0x42, 0x5d, 0x81, 0xca, 0x0a, 0x76, 0x56, 0x32, 0x14, 0x0b,
	0x14, 0x87, 0x19, 0x07, 0xc1, 0x93, 0xf0, 0xba, 0xd8, 0x6e, 0x11, 0x11, 0x2a, 0x3b, 0xcf, 0xe5,
	0x9f, 0xef, 0x9f, 0x31, 0xf4, 0x32, 0x35, 0x2b, 0xc9, 0x58, 0x23, 0x5a, 0x99, 0x4a, 0x2f, 0x61,
	0xbc, 0x66, 0xae, 0x70, 0x89, 0x64, 0xf5, 0xa3, 0xce, 0x94, 0xc5, 0x0d, 0xbe, 0x56, 0xc8, 0x56,
	0x0c, 0x20, 0xca, 0x98, 0x64, 0x73, 0xda, 0x3c, 0x8b, 0xc5, 0x09, 0x0c, 0x09, 0xb7, 0x9a, 0x2d,
	0x29, 0xab, 0x4d, 0xb1, 0x5e, 0xc9, 0x96, 0xcb, 0x47, 0x29, 0xc3, 0xa0, 0x26, 0x3d, 0xd0, 0xe6,
	0xe5, 0x91, 0x18, 0x42, 0x87, 0x91, 0xb4, 0x7a, 0x09, 0xb2, 0xbe, 0x8f, 0x73, 0xbd, 0x75, 0x14,
	0x19, 0x85, 0xd8, 0xb1, 0x72, 0x24, 0xf9, 0x2f, 0xb0, 0x5c, 0x51, 0x7b, 0x4f, 0xb9, 0x6c, 0x07,
	0xf1, 0x08, 0xba, 0xf8, 0x5e, 0x6a, 0x42, 0x96, 0x9d, 0x00, 0xbd, 0x87, 0xa3, 0x6b, 0x2c, 0xd0,
	0x31, 0xf0, 0x66, 0x79, 0x77, 0xfb, 0x6d, 0xd8, 0xf5, 0x65, 0xce, 0xcb, 0xea, 0x6a, 0xb3, 0x37,
	0xed, 0xa9, 0x56, 0xd9, 0x8a, 0x7f, 0xa8, 0x84, 0x8a, 0x4d, 0x11, 0xa8, 0x6d, 0xf1, 0x1f, 0xfa,
	0x84, 0x6f, 0xe6, 0x19, 0xf3, 0xb9, 0x0d, 0xec, 0x28, 0x9d, 0x42, 0xbc, 0x1b, 0xc9, 0xa5, 0x29,
	0x18, 0x45, 0x02, 0x3d, 0xda, 0xbf, 0x77, 0x43, 0xcf, 0x3f, 0x9b, 0x70, 0x5c, 0x5b, 0x79, 0x5e,
	0xd9, 0x27, 0x43, 0xda, 0x7e, 0x88, 0x05, 0x24, 0xbf, 0x4f, 0x29, 0x4e, 0x67, 0xee, 0xda, 0x7f,
	0x1c, 0x78, 0x32, 0xf2, 0xc5, 0x5a, 0x3e, 0x6d, 0x88, 0x0b, 0x88, 0xeb, 0x9b, 0x89, 0xb1, 0x6f,
	0x39, 0xb0, 0xeb, 0x24, 0xf1, 0x85, 0xba, 0xd3, 0xb4, 0xb1, 0xe8, 0x3e, 0xb4, 0xc3, 0xc7, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x25, 0xda, 0x9b, 0x25, 0xe3, 0x01, 0x00, 0x00,
}
