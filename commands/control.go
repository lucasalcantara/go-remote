package commands

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
)

const (
	methodPath  = "/{0}/{1}"
	serviceName = "control.commands"
)

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0x6b, 0xe4, 0xc9, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a,
	0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa1, 0x87, 0xe4, 0x02, 0x64,
	0xcb, 0xa4, 0xc4, 0xb0, 0xc8, 0x00, 0xad, 0x50, 0x62, 0x70, 0x32, 0xe0, 0x92, 0xce, 0xcc, 0xd7,
	0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x46, 0x52, 0xeb,
	0xc4, 0x0f, 0x56, 0x1c, 0x0e, 0x62, 0x07, 0x80, 0xbc, 0x14, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x9b,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb7, 0xcd, 0xf2, 0xef, 0x00, 0x00, 0x00,
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: serviceName,
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MouseCommandEndPoint",
			Handler:    _Command_MouseCommandEndPoint_Handler,
		},
		{
			MethodName: "ClickCommandEndPoint",
			Handler:    _Command_ClickCommandEndPoint_Handler,
		},
		{
			MethodName: "KeyboardCommandEndPoint",
			Handler:    _Command_KeyboardCommandEndPoint_Handler,
		},
		{
			MethodName: "SpaceCommandEndPoint",
			Handler:    _Command_SpaceCommandEndPoint_Handler,
		},
		{
			MethodName: "EnterCommandEndPoint",
			Handler:    _Command_EnterCommandEndPoint_Handler,
		},
		{
			MethodName: "RightCommandEndPoint",
			Handler:    _Command_RightCommandEndPoint_Handler,
		},
		{
			MethodName: "LeftCommandEndPoint",
			Handler:    _Command_LeftCommandEndPoint_Handler,
		},
		{
			MethodName: "LowerAudioCommandEndPoint",
			Handler:    _Command_LowerAudioCommandEndPoint_Handler,
		},
		{
			MethodName: "UpCommandEndPoint",
			Handler:    _Command_UpCommandEndPoint_Handler,
		},
		{
			MethodName: "DownCommandEndPoint",
			Handler:    _Command_DownCommandEndPoint_Handler,
		},
		{
			MethodName: "IncreaseAudioCommandEndPoint",
			Handler:    _Command_IncreaseAudioCommandEndPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func _Command_MouseCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).MouseCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).MouseCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "MouseCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_ClickCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).ClickCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).ClickCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "ClickCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_KeyboardCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).KeyboardCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).KeyboardCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "KeyboardCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_SpaceCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).SpaceCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).SpaceCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "SpaceCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_EnterCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).EnterCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).EnterCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "EnterCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_LeftCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).LeftCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).LeftCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "LeftCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_RightCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).RightCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).RightCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "RightCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_LowerAudioCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).LowerAudioCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).LowerAudioCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "LowerAudioCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_UpCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).UpCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).UpCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "UpCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_DownCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).DownCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).DownCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "DownCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func _Command_IncreaseAudioCommandEndPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).IncreaseAudioCommandEndPoint(ctx, in)
	}

	endPoint := srv.(ControlServer).IncreaseAudioCommandEndPoint
	info, handler := fillUnaryServerInfoAndUndaryHandler(srv, "IncreaseAudioCommandEndPoint", endPoint)

	return interceptor(ctx, in, info, handler)
}

func fillUnaryServerInfoAndUndaryHandler(srv interface{}, method string, endPoint EndPoint) (info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) {
	info = &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.commands/IncreaseAudioCommandEndPoint",
	}

	handler = func(ctx context.Context, req interface{}) (interface{}, error) {
		return endPoint(ctx, req.(*CommandRequest))
	}

	return
}

func RegisterGreeterServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func init() { proto.RegisterFile("control.proto", fileDescriptor0) }
