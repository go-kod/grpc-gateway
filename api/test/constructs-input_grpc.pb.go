// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: test/constructs-input.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Constructs_Scalars__FullMethodName   = "/constructs.Constructs/Scalars_"
	Constructs_Repeated__FullMethodName  = "/constructs.Constructs/Repeated_"
	Constructs_Maps__FullMethodName      = "/constructs.Constructs/Maps_"
	Constructs_Any__FullMethodName       = "/constructs.Constructs/Any_"
	Constructs_Empty__FullMethodName     = "/constructs.Constructs/Empty_"
	Constructs_Empty2__FullMethodName    = "/constructs.Constructs/Empty2_"
	Constructs_Empty3__FullMethodName    = "/constructs.Constructs/Empty3_"
	Constructs_Ref__FullMethodName       = "/constructs.Constructs/Ref_"
	Constructs_Oneof__FullMethodName     = "/constructs.Constructs/Oneof_"
	Constructs_CallWithId_FullMethodName = "/constructs.Constructs/CallWithId"
	Constructs_Anyway__FullMethodName    = "/constructs.Constructs/Anyway_"
)

// ConstructsClient is the client API for Constructs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConstructsClient interface {
	Scalars_(ctx context.Context, in *Scalars, opts ...grpc.CallOption) (*Scalars, error)
	Repeated_(ctx context.Context, in *Repeated, opts ...grpc.CallOption) (*Repeated, error)
	Maps_(ctx context.Context, in *Maps, opts ...grpc.CallOption) (*Maps, error)
	Any_(ctx context.Context, in *anypb.Any, opts ...grpc.CallOption) (*anypb.Any, error)
	Empty_(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Empty, error)
	Empty2_(ctx context.Context, in *EmptyRecursive, opts ...grpc.CallOption) (*EmptyNested, error)
	Empty3_(ctx context.Context, in *Empty3, opts ...grpc.CallOption) (*Empty3, error)
	Ref_(ctx context.Context, in *Ref, opts ...grpc.CallOption) (*Ref, error)
	Oneof_(ctx context.Context, in *Oneof, opts ...grpc.CallOption) (*Oneof, error)
	CallWithId(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Anyway_(ctx context.Context, in *Any, opts ...grpc.CallOption) (*AnyInput, error)
}

type constructsClient struct {
	cc grpc.ClientConnInterface
}

func NewConstructsClient(cc grpc.ClientConnInterface) ConstructsClient {
	return &constructsClient{cc}
}

func (c *constructsClient) Scalars_(ctx context.Context, in *Scalars, opts ...grpc.CallOption) (*Scalars, error) {
	out := new(Scalars)
	err := c.cc.Invoke(ctx, Constructs_Scalars__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Repeated_(ctx context.Context, in *Repeated, opts ...grpc.CallOption) (*Repeated, error) {
	out := new(Repeated)
	err := c.cc.Invoke(ctx, Constructs_Repeated__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Maps_(ctx context.Context, in *Maps, opts ...grpc.CallOption) (*Maps, error) {
	out := new(Maps)
	err := c.cc.Invoke(ctx, Constructs_Maps__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Any_(ctx context.Context, in *anypb.Any, opts ...grpc.CallOption) (*anypb.Any, error) {
	out := new(anypb.Any)
	err := c.cc.Invoke(ctx, Constructs_Any__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Empty_(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Constructs_Empty__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Empty2_(ctx context.Context, in *EmptyRecursive, opts ...grpc.CallOption) (*EmptyNested, error) {
	out := new(EmptyNested)
	err := c.cc.Invoke(ctx, Constructs_Empty2__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Empty3_(ctx context.Context, in *Empty3, opts ...grpc.CallOption) (*Empty3, error) {
	out := new(Empty3)
	err := c.cc.Invoke(ctx, Constructs_Empty3__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Ref_(ctx context.Context, in *Ref, opts ...grpc.CallOption) (*Ref, error) {
	out := new(Ref)
	err := c.cc.Invoke(ctx, Constructs_Ref__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Oneof_(ctx context.Context, in *Oneof, opts ...grpc.CallOption) (*Oneof, error) {
	out := new(Oneof)
	err := c.cc.Invoke(ctx, Constructs_Oneof__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) CallWithId(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Constructs_CallWithId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *constructsClient) Anyway_(ctx context.Context, in *Any, opts ...grpc.CallOption) (*AnyInput, error) {
	out := new(AnyInput)
	err := c.cc.Invoke(ctx, Constructs_Anyway__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConstructsServer is the server API for Constructs service.
// All implementations should embed UnimplementedConstructsServer
// for forward compatibility
type ConstructsServer interface {
	Scalars_(context.Context, *Scalars) (*Scalars, error)
	Repeated_(context.Context, *Repeated) (*Repeated, error)
	Maps_(context.Context, *Maps) (*Maps, error)
	Any_(context.Context, *anypb.Any) (*anypb.Any, error)
	Empty_(context.Context, *emptypb.Empty) (*Empty, error)
	Empty2_(context.Context, *EmptyRecursive) (*EmptyNested, error)
	Empty3_(context.Context, *Empty3) (*Empty3, error)
	Ref_(context.Context, *Ref) (*Ref, error)
	Oneof_(context.Context, *Oneof) (*Oneof, error)
	CallWithId(context.Context, *Empty) (*Empty, error)
	Anyway_(context.Context, *Any) (*AnyInput, error)
}

// UnimplementedConstructsServer should be embedded to have forward compatible implementations.
type UnimplementedConstructsServer struct {
}

func (UnimplementedConstructsServer) Scalars_(context.Context, *Scalars) (*Scalars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scalars_ not implemented")
}
func (UnimplementedConstructsServer) Repeated_(context.Context, *Repeated) (*Repeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Repeated_ not implemented")
}
func (UnimplementedConstructsServer) Maps_(context.Context, *Maps) (*Maps, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Maps_ not implemented")
}
func (UnimplementedConstructsServer) Any_(context.Context, *anypb.Any) (*anypb.Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Any_ not implemented")
}
func (UnimplementedConstructsServer) Empty_(context.Context, *emptypb.Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty_ not implemented")
}
func (UnimplementedConstructsServer) Empty2_(context.Context, *EmptyRecursive) (*EmptyNested, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty2_ not implemented")
}
func (UnimplementedConstructsServer) Empty3_(context.Context, *Empty3) (*Empty3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty3_ not implemented")
}
func (UnimplementedConstructsServer) Ref_(context.Context, *Ref) (*Ref, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ref_ not implemented")
}
func (UnimplementedConstructsServer) Oneof_(context.Context, *Oneof) (*Oneof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Oneof_ not implemented")
}
func (UnimplementedConstructsServer) CallWithId(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallWithId not implemented")
}
func (UnimplementedConstructsServer) Anyway_(context.Context, *Any) (*AnyInput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Anyway_ not implemented")
}

// UnsafeConstructsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConstructsServer will
// result in compilation errors.
type UnsafeConstructsServer interface {
	mustEmbedUnimplementedConstructsServer()
}

func RegisterConstructsServer(s grpc.ServiceRegistrar, srv ConstructsServer) {
	s.RegisterService(&Constructs_ServiceDesc, srv)
}

func _Constructs_Scalars__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scalars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Scalars_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Scalars__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Scalars_(ctx, req.(*Scalars))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Repeated__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repeated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Repeated_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Repeated__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Repeated_(ctx, req.(*Repeated))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Maps__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Maps)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Maps_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Maps__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Maps_(ctx, req.(*Maps))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Any__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(anypb.Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Any_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Any__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Any_(ctx, req.(*anypb.Any))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Empty__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Empty_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Empty__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Empty_(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Empty2__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRecursive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Empty2_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Empty2__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Empty2_(ctx, req.(*EmptyRecursive))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Empty3__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty3)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Empty3_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Empty3__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Empty3_(ctx, req.(*Empty3))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Ref__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ref)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Ref_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Ref__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Ref_(ctx, req.(*Ref))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Oneof__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oneof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Oneof_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Oneof__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Oneof_(ctx, req.(*Oneof))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_CallWithId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).CallWithId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_CallWithId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).CallWithId(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Constructs_Anyway__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConstructsServer).Anyway_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Constructs_Anyway__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConstructsServer).Anyway_(ctx, req.(*Any))
	}
	return interceptor(ctx, in, info, handler)
}

// Constructs_ServiceDesc is the grpc.ServiceDesc for Constructs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Constructs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "constructs.Constructs",
	HandlerType: (*ConstructsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scalars_",
			Handler:    _Constructs_Scalars__Handler,
		},
		{
			MethodName: "Repeated_",
			Handler:    _Constructs_Repeated__Handler,
		},
		{
			MethodName: "Maps_",
			Handler:    _Constructs_Maps__Handler,
		},
		{
			MethodName: "Any_",
			Handler:    _Constructs_Any__Handler,
		},
		{
			MethodName: "Empty_",
			Handler:    _Constructs_Empty__Handler,
		},
		{
			MethodName: "Empty2_",
			Handler:    _Constructs_Empty2__Handler,
		},
		{
			MethodName: "Empty3_",
			Handler:    _Constructs_Empty3__Handler,
		},
		{
			MethodName: "Ref_",
			Handler:    _Constructs_Ref__Handler,
		},
		{
			MethodName: "Oneof_",
			Handler:    _Constructs_Oneof__Handler,
		},
		{
			MethodName: "CallWithId",
			Handler:    _Constructs_CallWithId_Handler,
		},
		{
			MethodName: "Anyway_",
			Handler:    _Constructs_Anyway__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/constructs-input.proto",
}
