// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: grpcApi/service.proto

package grpcApi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	EncryptData(ctx context.Context, in *EncryptDecryptRequest, opts ...grpc.CallOption) (*EncryptDecryptResponse, error)
	DecryptData(ctx context.Context, in *EncryptDecryptRequest, opts ...grpc.CallOption) (*EncryptDecryptResponse, error)
	EncryptFileFromComputer(ctx context.Context, in *EncryptDecryptFileRequest, opts ...grpc.CallOption) (*EncryptDecryptFileResponse, error)
	DecryptFileFromComputer(ctx context.Context, in *EncryptDecryptFileRequest, opts ...grpc.CallOption) (*EncryptDecryptFileResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) EncryptData(ctx context.Context, in *EncryptDecryptRequest, opts ...grpc.CallOption) (*EncryptDecryptResponse, error) {
	out := new(EncryptDecryptResponse)
	err := c.cc.Invoke(ctx, "/service.CryptoService/EncryptData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DecryptData(ctx context.Context, in *EncryptDecryptRequest, opts ...grpc.CallOption) (*EncryptDecryptResponse, error) {
	out := new(EncryptDecryptResponse)
	err := c.cc.Invoke(ctx, "/service.CryptoService/DecryptData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) EncryptFileFromComputer(ctx context.Context, in *EncryptDecryptFileRequest, opts ...grpc.CallOption) (*EncryptDecryptFileResponse, error) {
	out := new(EncryptDecryptFileResponse)
	err := c.cc.Invoke(ctx, "/service.CryptoService/EncryptFileFromComputer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DecryptFileFromComputer(ctx context.Context, in *EncryptDecryptFileRequest, opts ...grpc.CallOption) (*EncryptDecryptFileResponse, error) {
	out := new(EncryptDecryptFileResponse)
	err := c.cc.Invoke(ctx, "/service.CryptoService/DecryptFileFromComputer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	EncryptData(context.Context, *EncryptDecryptRequest) (*EncryptDecryptResponse, error)
	DecryptData(context.Context, *EncryptDecryptRequest) (*EncryptDecryptResponse, error)
	EncryptFileFromComputer(context.Context, *EncryptDecryptFileRequest) (*EncryptDecryptFileResponse, error)
	DecryptFileFromComputer(context.Context, *EncryptDecryptFileRequest) (*EncryptDecryptFileResponse, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) EncryptData(context.Context, *EncryptDecryptRequest) (*EncryptDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptData not implemented")
}
func (UnimplementedCryptoServiceServer) DecryptData(context.Context, *EncryptDecryptRequest) (*EncryptDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptData not implemented")
}
func (UnimplementedCryptoServiceServer) EncryptFileFromComputer(context.Context, *EncryptDecryptFileRequest) (*EncryptDecryptFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptFileFromComputer not implemented")
}
func (UnimplementedCryptoServiceServer) DecryptFileFromComputer(context.Context, *EncryptDecryptFileRequest) (*EncryptDecryptFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptFileFromComputer not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_EncryptData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).EncryptData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CryptoService/EncryptData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).EncryptData(ctx, req.(*EncryptDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DecryptData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DecryptData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CryptoService/DecryptData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DecryptData(ctx, req.(*EncryptDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_EncryptFileFromComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptDecryptFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).EncryptFileFromComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CryptoService/EncryptFileFromComputer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).EncryptFileFromComputer(ctx, req.(*EncryptDecryptFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DecryptFileFromComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptDecryptFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DecryptFileFromComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CryptoService/DecryptFileFromComputer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DecryptFileFromComputer(ctx, req.(*EncryptDecryptFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncryptData",
			Handler:    _CryptoService_EncryptData_Handler,
		},
		{
			MethodName: "DecryptData",
			Handler:    _CryptoService_DecryptData_Handler,
		},
		{
			MethodName: "EncryptFileFromComputer",
			Handler:    _CryptoService_EncryptFileFromComputer_Handler,
		},
		{
			MethodName: "DecryptFileFromComputer",
			Handler:    _CryptoService_DecryptFileFromComputer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcApi/service.proto",
}
