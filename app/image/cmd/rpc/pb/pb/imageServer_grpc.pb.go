// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: imageServer.proto

package pb

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

const (
	Image_GetMyImage_FullMethodName   = "/pb.image/getMyImage"
	Image_GetUserImage_FullMethodName = "/pb.image/getUserImage"
	Image_GetAllImage_FullMethodName  = "/pb.image/getAllImage"
)

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	GetMyImage(ctx context.Context, in *GetMyImageReq, opts ...grpc.CallOption) (*GetMyImageResp, error)
	GetUserImage(ctx context.Context, in *GetUserImageReq, opts ...grpc.CallOption) (*GetUserImageResp, error)
	GetAllImage(ctx context.Context, in *GetAllImageReq, opts ...grpc.CallOption) (*GetAllImageResp, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) GetMyImage(ctx context.Context, in *GetMyImageReq, opts ...grpc.CallOption) (*GetMyImageResp, error) {
	out := new(GetMyImageResp)
	err := c.cc.Invoke(ctx, Image_GetMyImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetUserImage(ctx context.Context, in *GetUserImageReq, opts ...grpc.CallOption) (*GetUserImageResp, error) {
	out := new(GetUserImageResp)
	err := c.cc.Invoke(ctx, Image_GetUserImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetAllImage(ctx context.Context, in *GetAllImageReq, opts ...grpc.CallOption) (*GetAllImageResp, error) {
	out := new(GetAllImageResp)
	err := c.cc.Invoke(ctx, Image_GetAllImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	GetMyImage(context.Context, *GetMyImageReq) (*GetMyImageResp, error)
	GetUserImage(context.Context, *GetUserImageReq) (*GetUserImageResp, error)
	GetAllImage(context.Context, *GetAllImageReq) (*GetAllImageResp, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) GetMyImage(context.Context, *GetMyImageReq) (*GetMyImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyImage not implemented")
}
func (UnimplementedImageServer) GetUserImage(context.Context, *GetUserImageReq) (*GetUserImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserImage not implemented")
}
func (UnimplementedImageServer) GetAllImage(context.Context, *GetAllImageReq) (*GetAllImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllImage not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_GetMyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetMyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_GetMyImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetMyImage(ctx, req.(*GetMyImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetUserImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetUserImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_GetUserImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetUserImage(ctx, req.(*GetUserImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetAllImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetAllImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_GetAllImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetAllImage(ctx, req.(*GetAllImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMyImage",
			Handler:    _Image_GetMyImage_Handler,
		},
		{
			MethodName: "getUserImage",
			Handler:    _Image_GetUserImage_Handler,
		},
		{
			MethodName: "getAllImage",
			Handler:    _Image_GetAllImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imageServer.proto",
}
