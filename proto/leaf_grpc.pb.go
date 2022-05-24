// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// LeafClient is the client API for Leaf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeafClient interface {
	GetLeafID(ctx context.Context, in *GetLeafReq, opts ...grpc.CallOption) (*GetLeafResp, error)
	CreateLeaf(ctx context.Context, in *CreateLeafReq, opts ...grpc.CallOption) (*CreateLeafResq, error)
}

type leafClient struct {
	cc grpc.ClientConnInterface
}

func NewLeafClient(cc grpc.ClientConnInterface) LeafClient {
	return &leafClient{cc}
}

func (c *leafClient) GetLeafID(ctx context.Context, in *GetLeafReq, opts ...grpc.CallOption) (*GetLeafResp, error) {
	out := new(GetLeafResp)
	err := c.cc.Invoke(ctx, "/proto.Leaf/GetLeafID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leafClient) CreateLeaf(ctx context.Context, in *CreateLeafReq, opts ...grpc.CallOption) (*CreateLeafResq, error) {
	out := new(CreateLeafResq)
	err := c.cc.Invoke(ctx, "/proto.Leaf/CreateLeaf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeafServer is the server API for Leaf service.
// All implementations should embed UnimplementedLeafServer
// for forward compatibility
type LeafServer interface {
	GetLeafID(context.Context, *GetLeafReq) (*GetLeafResp, error)
	CreateLeaf(context.Context, *CreateLeafReq) (*CreateLeafResq, error)
}

// UnimplementedLeafServer should be embedded to have forward compatible implementations.
type UnimplementedLeafServer struct {
}

func (UnimplementedLeafServer) GetLeafID(context.Context, *GetLeafReq) (*GetLeafResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeafID not implemented")
}
func (UnimplementedLeafServer) CreateLeaf(context.Context, *CreateLeafReq) (*CreateLeafResq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeaf not implemented")
}

// UnsafeLeafServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeafServer will
// result in compilation errors.
type UnsafeLeafServer interface {
	mustEmbedUnimplementedLeafServer()
}

func RegisterLeafServer(s grpc.ServiceRegistrar, srv LeafServer) {
	s.RegisterService(&Leaf_ServiceDesc, srv)
}

func _Leaf_GetLeafID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeafReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeafServer).GetLeafID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Leaf/GetLeafID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeafServer).GetLeafID(ctx, req.(*GetLeafReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaf_CreateLeaf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeafReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeafServer).CreateLeaf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Leaf/CreateLeaf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeafServer).CreateLeaf(ctx, req.(*CreateLeafReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Leaf_ServiceDesc is the grpc.ServiceDesc for Leaf service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leaf_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Leaf",
	HandlerType: (*LeafServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeafID",
			Handler:    _Leaf_GetLeafID_Handler,
		},
		{
			MethodName: "CreateLeaf",
			Handler:    _Leaf_CreateLeaf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaf.proto",
}
