// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: proto/staff/staff.proto

package staff

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

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	CreateStaff(ctx context.Context, in *CreateStaffRequset, opts ...grpc.CallOption) (*CreateStaffResponse, error)
	StaffLogin(ctx context.Context, in *StaffLoginRequest, opts ...grpc.CallOption) (*StaffLoginResponse, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) CreateStaff(ctx context.Context, in *CreateStaffRequset, opts ...grpc.CallOption) (*CreateStaffResponse, error) {
	out := new(CreateStaffResponse)
	err := c.cc.Invoke(ctx, "/staff.StaffService/CreateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffLogin(ctx context.Context, in *StaffLoginRequest, opts ...grpc.CallOption) (*StaffLoginResponse, error) {
	out := new(StaffLoginResponse)
	err := c.cc.Invoke(ctx, "/staff.StaffService/StaffLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations should embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	CreateStaff(context.Context, *CreateStaffRequset) (*CreateStaffResponse, error)
	StaffLogin(context.Context, *StaffLoginRequest) (*StaffLoginResponse, error)
}

// UnimplementedStaffServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) CreateStaff(context.Context, *CreateStaffRequset) (*CreateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaff not implemented")
}
func (UnimplementedStaffServiceServer) StaffLogin(context.Context, *StaffLoginRequest) (*StaffLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffLogin not implemented")
}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_CreateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).CreateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.StaffService/CreateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).CreateStaff(ctx, req.(*CreateStaffRequset))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.StaffService/StaffLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffLogin(ctx, req.(*StaffLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStaff",
			Handler:    _StaffService_CreateStaff_Handler,
		},
		{
			MethodName: "StaffLogin",
			Handler:    _StaffService_StaffLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/staff/staff.proto",
}
