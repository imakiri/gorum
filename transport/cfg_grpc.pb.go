// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transport

import (
	context "context"
	core "github.com/imakiri/playground/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CfgClient is the client API for Cfg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CfgClient interface {
	Api(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgApi, error)
	App(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgApp, error)
	Auth(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgAuth, error)
	Data(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgData, error)
	Gate(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgGate, error)
	Web(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgWeb, error)
}

type cfgClient struct {
	cc grpc.ClientConnInterface
}

func NewCfgClient(cc grpc.ClientConnInterface) CfgClient {
	return &cfgClient{cc}
}

func (c *cfgClient) Api(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgApi, error) {
	out := new(core.CfgApi)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/Api", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfgClient) App(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgApp, error) {
	out := new(core.CfgApp)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/App", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfgClient) Auth(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgAuth, error) {
	out := new(core.CfgAuth)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfgClient) Data(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgData, error) {
	out := new(core.CfgData)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/Data", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfgClient) Gate(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgGate, error) {
	out := new(core.CfgGate)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/Gate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfgClient) Web(ctx context.Context, in *core.Request, opts ...grpc.CallOption) (*core.CfgWeb, error) {
	out := new(core.CfgWeb)
	err := c.cc.Invoke(ctx, "/cfg.Cfg/Web", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CfgServer is the server API for Cfg service.
// All implementations must embed UnimplementedCfgServer
// for forward compatibility
type CfgServer interface {
	Api(context.Context, *core.Request) (*core.CfgApi, error)
	App(context.Context, *core.Request) (*core.CfgApp, error)
	Auth(context.Context, *core.Request) (*core.CfgAuth, error)
	Data(context.Context, *core.Request) (*core.CfgData, error)
	Gate(context.Context, *core.Request) (*core.CfgGate, error)
	Web(context.Context, *core.Request) (*core.CfgWeb, error)
	mustEmbedUnimplementedCfgServer()
}

// UnimplementedCfgServer must be embedded to have forward compatible implementations.
type UnimplementedCfgServer struct {
}

func (UnimplementedCfgServer) Api(context.Context, *core.Request) (*core.CfgApi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Api not implemented")
}
func (UnimplementedCfgServer) App(context.Context, *core.Request) (*core.CfgApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (UnimplementedCfgServer) Auth(context.Context, *core.Request) (*core.CfgAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedCfgServer) Data(context.Context, *core.Request) (*core.CfgData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Data not implemented")
}
func (UnimplementedCfgServer) Gate(context.Context, *core.Request) (*core.CfgGate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gate not implemented")
}
func (UnimplementedCfgServer) Web(context.Context, *core.Request) (*core.CfgWeb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Web not implemented")
}
func (UnimplementedCfgServer) mustEmbedUnimplementedCfgServer() {}

// UnsafeCfgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CfgServer will
// result in compilation errors.
type UnsafeCfgServer interface {
	mustEmbedUnimplementedCfgServer()
}

func RegisterCfgServer(s grpc.ServiceRegistrar, srv CfgServer) {
	s.RegisterService(&Cfg_ServiceDesc, srv)
}

func _Cfg_Api_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).Api(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/Api",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).Api(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cfg_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/App",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).App(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cfg_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).Auth(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cfg_Data_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).Data(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/Data",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).Data(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cfg_Gate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).Gate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/Gate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).Gate(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cfg_Web_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CfgServer).Web(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cfg.Cfg/Web",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CfgServer).Web(ctx, req.(*core.Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Cfg_ServiceDesc is the grpc.ServiceDesc for Cfg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cfg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cfg.Cfg",
	HandlerType: (*CfgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Api",
			Handler:    _Cfg_Api_Handler,
		},
		{
			MethodName: "App",
			Handler:    _Cfg_App_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Cfg_Auth_Handler,
		},
		{
			MethodName: "Data",
			Handler:    _Cfg_Data_Handler,
		},
		{
			MethodName: "Gate",
			Handler:    _Cfg_Gate_Handler,
		},
		{
			MethodName: "Web",
			Handler:    _Cfg_Web_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cfg.proto",
}
