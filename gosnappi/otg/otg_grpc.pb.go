// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package otg

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OpenapiClient is the client API for Openapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenapiClient interface {
	// Sets configuration resources on the traffic generator.
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	// Description missing in models
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	// The Response.Warnings in the Success response is available for implementers to disclose
	// additional information about a state change including any implicit changes that are
	// outside the scope of the state change.
	SetTransmitState(ctx context.Context, in *SetTransmitStateRequest, opts ...grpc.CallOption) (*SetTransmitStateResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetLinkState(ctx context.Context, in *SetLinkStateRequest, opts ...grpc.CallOption) (*SetLinkStateResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetCaptureState(ctx context.Context, in *SetCaptureStateRequest, opts ...grpc.CallOption) (*SetCaptureStateResponse, error)
	// Updates flow properties without disruption of transmit state.
	UpdateFlows(ctx context.Context, in *UpdateFlowsRequest, opts ...grpc.CallOption) (*UpdateFlowsResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetRouteState(ctx context.Context, in *SetRouteStateRequest, opts ...grpc.CallOption) (*SetRouteStateResponse, error)
	// API to send an IPv4 and/or IPv6 ICMP Echo Request(s) between endpoints. For each
	// endpoint 1 ping packet will be sent and API shall wait for ping response to either
	// be successful or timeout. The API wait timeout for each request is 300ms.
	SendPing(ctx context.Context, in *SendPingRequest, opts ...grpc.CallOption) (*SendPingResponse, error)
	// Sets all configured protocols to `start` or `stop` state.
	SetProtocolState(ctx context.Context, in *SetProtocolStateRequest, opts ...grpc.CallOption) (*SetProtocolStateResponse, error)
	// Set specific state/actions on device configuration resources on the traffic generator.
	SetDeviceState(ctx context.Context, in *SetDeviceStateRequest, opts ...grpc.CallOption) (*SetDeviceStateResponse, error)
	// Description missing in models
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Description missing in models
	GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error)
	// Description missing in models
	GetCapture(ctx context.Context, in *GetCaptureRequest, opts ...grpc.CallOption) (*GetCaptureResponse, error)
}

type openapiClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenapiClient(cc grpc.ClientConnInterface) OpenapiClient {
	return &openapiClient{cc}
}

func (c *openapiClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetTransmitState(ctx context.Context, in *SetTransmitStateRequest, opts ...grpc.CallOption) (*SetTransmitStateResponse, error) {
	out := new(SetTransmitStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetTransmitState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetLinkState(ctx context.Context, in *SetLinkStateRequest, opts ...grpc.CallOption) (*SetLinkStateResponse, error) {
	out := new(SetLinkStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetLinkState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetCaptureState(ctx context.Context, in *SetCaptureStateRequest, opts ...grpc.CallOption) (*SetCaptureStateResponse, error) {
	out := new(SetCaptureStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetCaptureState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) UpdateFlows(ctx context.Context, in *UpdateFlowsRequest, opts ...grpc.CallOption) (*UpdateFlowsResponse, error) {
	out := new(UpdateFlowsResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/UpdateFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetRouteState(ctx context.Context, in *SetRouteStateRequest, opts ...grpc.CallOption) (*SetRouteStateResponse, error) {
	out := new(SetRouteStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetRouteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SendPing(ctx context.Context, in *SendPingRequest, opts ...grpc.CallOption) (*SendPingResponse, error) {
	out := new(SendPingResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SendPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetProtocolState(ctx context.Context, in *SetProtocolStateRequest, opts ...grpc.CallOption) (*SetProtocolStateResponse, error) {
	out := new(SetProtocolStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetProtocolState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) SetDeviceState(ctx context.Context, in *SetDeviceStateRequest, opts ...grpc.CallOption) (*SetDeviceStateResponse, error) {
	out := new(SetDeviceStateResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/SetDeviceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error) {
	out := new(GetStatesResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/GetStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) GetCapture(ctx context.Context, in *GetCaptureRequest, opts ...grpc.CallOption) (*GetCaptureResponse, error) {
	out := new(GetCaptureResponse)
	err := c.cc.Invoke(ctx, "/otg.Openapi/GetCapture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenapiServer is the server API for Openapi service.
// All implementations must embed UnimplementedOpenapiServer
// for forward compatibility
type OpenapiServer interface {
	// Sets configuration resources on the traffic generator.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	// Description missing in models
	GetConfig(context.Context, *emptypb.Empty) (*GetConfigResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	// The Response.Warnings in the Success response is available for implementers to disclose
	// additional information about a state change including any implicit changes that are
	// outside the scope of the state change.
	SetTransmitState(context.Context, *SetTransmitStateRequest) (*SetTransmitStateResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetLinkState(context.Context, *SetLinkStateRequest) (*SetLinkStateResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetCaptureState(context.Context, *SetCaptureStateRequest) (*SetCaptureStateResponse, error)
	// Updates flow properties without disruption of transmit state.
	UpdateFlows(context.Context, *UpdateFlowsRequest) (*UpdateFlowsResponse, error)
	// Updates the state of configuration resources on the traffic generator.
	SetRouteState(context.Context, *SetRouteStateRequest) (*SetRouteStateResponse, error)
	// API to send an IPv4 and/or IPv6 ICMP Echo Request(s) between endpoints. For each
	// endpoint 1 ping packet will be sent and API shall wait for ping response to either
	// be successful or timeout. The API wait timeout for each request is 300ms.
	SendPing(context.Context, *SendPingRequest) (*SendPingResponse, error)
	// Sets all configured protocols to `start` or `stop` state.
	SetProtocolState(context.Context, *SetProtocolStateRequest) (*SetProtocolStateResponse, error)
	// Set specific state/actions on device configuration resources on the traffic generator.
	SetDeviceState(context.Context, *SetDeviceStateRequest) (*SetDeviceStateResponse, error)
	// Description missing in models
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Description missing in models
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)
	// Description missing in models
	GetCapture(context.Context, *GetCaptureRequest) (*GetCaptureResponse, error)
	mustEmbedUnimplementedOpenapiServer()
}

// UnimplementedOpenapiServer must be embedded to have forward compatible implementations.
type UnimplementedOpenapiServer struct {
}

func (UnimplementedOpenapiServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedOpenapiServer) GetConfig(context.Context, *emptypb.Empty) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedOpenapiServer) SetTransmitState(context.Context, *SetTransmitStateRequest) (*SetTransmitStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTransmitState not implemented")
}
func (UnimplementedOpenapiServer) SetLinkState(context.Context, *SetLinkStateRequest) (*SetLinkStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLinkState not implemented")
}
func (UnimplementedOpenapiServer) SetCaptureState(context.Context, *SetCaptureStateRequest) (*SetCaptureStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCaptureState not implemented")
}
func (UnimplementedOpenapiServer) UpdateFlows(context.Context, *UpdateFlowsRequest) (*UpdateFlowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlows not implemented")
}
func (UnimplementedOpenapiServer) SetRouteState(context.Context, *SetRouteStateRequest) (*SetRouteStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRouteState not implemented")
}
func (UnimplementedOpenapiServer) SendPing(context.Context, *SendPingRequest) (*SendPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPing not implemented")
}
func (UnimplementedOpenapiServer) SetProtocolState(context.Context, *SetProtocolStateRequest) (*SetProtocolStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProtocolState not implemented")
}
func (UnimplementedOpenapiServer) SetDeviceState(context.Context, *SetDeviceStateRequest) (*SetDeviceStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceState not implemented")
}
func (UnimplementedOpenapiServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedOpenapiServer) GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStates not implemented")
}
func (UnimplementedOpenapiServer) GetCapture(context.Context, *GetCaptureRequest) (*GetCaptureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapture not implemented")
}
func (UnimplementedOpenapiServer) mustEmbedUnimplementedOpenapiServer() {}

// UnsafeOpenapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenapiServer will
// result in compilation errors.
type UnsafeOpenapiServer interface {
	mustEmbedUnimplementedOpenapiServer()
}

func RegisterOpenapiServer(s grpc.ServiceRegistrar, srv OpenapiServer) {
	s.RegisterService(&Openapi_ServiceDesc, srv)
}

func _Openapi_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetTransmitState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTransmitStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetTransmitState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetTransmitState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetTransmitState(ctx, req.(*SetTransmitStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetLinkState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLinkStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetLinkState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetLinkState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetLinkState(ctx, req.(*SetLinkStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetCaptureState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCaptureStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetCaptureState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetCaptureState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetCaptureState(ctx, req.(*SetCaptureStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_UpdateFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).UpdateFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/UpdateFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).UpdateFlows(ctx, req.(*UpdateFlowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetRouteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRouteStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetRouteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetRouteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetRouteState(ctx, req.(*SetRouteStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SendPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SendPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SendPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SendPing(ctx, req.(*SendPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetProtocolState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProtocolStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetProtocolState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetProtocolState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetProtocolState(ctx, req.(*SetProtocolStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_SetDeviceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).SetDeviceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/SetDeviceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).SetDeviceState(ctx, req.(*SetDeviceStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_GetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).GetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/GetStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).GetStates(ctx, req.(*GetStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_GetCapture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).GetCapture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otg.Openapi/GetCapture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).GetCapture(ctx, req.(*GetCaptureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Openapi_ServiceDesc is the grpc.ServiceDesc for Openapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Openapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otg.Openapi",
	HandlerType: (*OpenapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConfig",
			Handler:    _Openapi_SetConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Openapi_GetConfig_Handler,
		},
		{
			MethodName: "SetTransmitState",
			Handler:    _Openapi_SetTransmitState_Handler,
		},
		{
			MethodName: "SetLinkState",
			Handler:    _Openapi_SetLinkState_Handler,
		},
		{
			MethodName: "SetCaptureState",
			Handler:    _Openapi_SetCaptureState_Handler,
		},
		{
			MethodName: "UpdateFlows",
			Handler:    _Openapi_UpdateFlows_Handler,
		},
		{
			MethodName: "SetRouteState",
			Handler:    _Openapi_SetRouteState_Handler,
		},
		{
			MethodName: "SendPing",
			Handler:    _Openapi_SendPing_Handler,
		},
		{
			MethodName: "SetProtocolState",
			Handler:    _Openapi_SetProtocolState_Handler,
		},
		{
			MethodName: "SetDeviceState",
			Handler:    _Openapi_SetDeviceState_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _Openapi_GetMetrics_Handler,
		},
		{
			MethodName: "GetStates",
			Handler:    _Openapi_GetStates_Handler,
		},
		{
			MethodName: "GetCapture",
			Handler:    _Openapi_GetCapture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "otg.proto",
}
