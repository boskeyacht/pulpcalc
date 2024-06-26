// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: tree.proto

package proto

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

// CalcTreeClient is the client API for CalcTree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcTreeClient interface {
	GetCalcTree(ctx context.Context, in *CalcTreeRequest, opts ...grpc.CallOption) (*CalcTreeResponse, error)
}

type calcTreeClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcTreeClient(cc grpc.ClientConnInterface) CalcTreeClient {
	return &calcTreeClient{cc}
}

func (c *calcTreeClient) GetCalcTree(ctx context.Context, in *CalcTreeRequest, opts ...grpc.CallOption) (*CalcTreeResponse, error) {
	out := new(CalcTreeResponse)
	err := c.cc.Invoke(ctx, "/proto.CalcTree/GetCalcTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcTreeServer is the server API for CalcTree service.
// All implementations must embed UnimplementedCalcTreeServer
// for forward compatibility
type CalcTreeServer interface {
	GetCalcTree(context.Context, *CalcTreeRequest) (*CalcTreeResponse, error)
	mustEmbedUnimplementedCalcTreeServer()
}

// UnimplementedCalcTreeServer must be embedded to have forward compatible implementations.
type UnimplementedCalcTreeServer struct {
}

func (UnimplementedCalcTreeServer) GetCalcTree(context.Context, *CalcTreeRequest) (*CalcTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCalcTree not implemented")
}
func (UnimplementedCalcTreeServer) mustEmbedUnimplementedCalcTreeServer() {}

// UnsafeCalcTreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcTreeServer will
// result in compilation errors.
type UnsafeCalcTreeServer interface {
	mustEmbedUnimplementedCalcTreeServer()
}

func RegisterCalcTreeServer(s grpc.ServiceRegistrar, srv CalcTreeServer) {
	s.RegisterService(&CalcTree_ServiceDesc, srv)
}

func _CalcTree_GetCalcTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcTreeServer).GetCalcTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CalcTree/GetCalcTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcTreeServer).GetCalcTree(ctx, req.(*CalcTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcTree_ServiceDesc is the grpc.ServiceDesc for CalcTree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcTree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CalcTree",
	HandlerType: (*CalcTreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCalcTree",
			Handler:    _CalcTree_GetCalcTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tree.proto",
}
