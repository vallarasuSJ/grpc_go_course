// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: calculator.proto

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

const (
	CalculatorService_Sum_FullMethodName           = "/calculator.CalculatorService/Sum"
	CalculatorService_Multiply_FullMethodName      = "/calculator.CalculatorService/Multiply"
	CalculatorService_Primes_FullMethodName        = "/calculator.CalculatorService/Primes"
	CalculatorService_Average_FullMethodName       = "/calculator.CalculatorService/Average"
	CalculatorService_Max_FullMethodName           = "/calculator.CalculatorService/Max"
	CalculatorService_Subtract_FullMethodName      = "/calculator.CalculatorService/Subtract"
	CalculatorService_Fibonacci_FullMethodName     = "/calculator.CalculatorService/Fibonacci"
	CalculatorService_MaxEvenNumber_FullMethodName = "/calculator.CalculatorService/MaxEvenNumber"
	CalculatorService_VowelsCount_FullMethodName   = "/calculator.CalculatorService/VowelsCount"
	CalculatorService_Sqrt_FullMethodName          = "/calculator.CalculatorService/Sqrt"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
	Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimesClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error)
	Subtract(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error)
	Fibonacci(ctx context.Context, in *FibRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error)
	MaxEvenNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxEvenNumberClient, error)
	VowelsCount(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_VowelsCountClient, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Sum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Multiply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_Primes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], CalculatorService_Average_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], CalculatorService_Max_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxClient{stream}
	return x, nil
}

type CalculatorService_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Subtract(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error) {
	out := new(SubResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Subtract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Fibonacci(ctx context.Context, in *FibRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[3], CalculatorService_Fibonacci_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_FibonacciClient interface {
	Recv() (*FibResponse, error)
	grpc.ClientStream
}

type calculatorServiceFibonacciClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFibonacciClient) Recv() (*FibResponse, error) {
	m := new(FibResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) MaxEvenNumber(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxEvenNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[4], CalculatorService_MaxEvenNumber_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxEvenNumberClient{stream}
	return x, nil
}

type CalculatorService_MaxEvenNumberClient interface {
	Send(*EvenRequest) error
	CloseAndRecv() (*EvenResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaxEvenNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxEvenNumberClient) Send(m *EvenRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxEvenNumberClient) CloseAndRecv() (*EvenResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EvenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) VowelsCount(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_VowelsCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[5], CalculatorService_VowelsCount_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceVowelsCountClient{stream}
	return x, nil
}

type CalculatorService_VowelsCountClient interface {
	Send(*VowelRequest) error
	Recv() (*VowelResponse, error)
	grpc.ClientStream
}

type calculatorServiceVowelsCountClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceVowelsCountClient) Send(m *VowelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceVowelsCountClient) Recv() (*VowelResponse, error) {
	m := new(VowelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Sqrt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	Primes(*PrimeRequest, CalculatorService_PrimesServer) error
	Average(CalculatorService_AverageServer) error
	Max(CalculatorService_MaxServer) error
	Subtract(context.Context, *SubRequest) (*SubResponse, error)
	Fibonacci(*FibRequest, CalculatorService_FibonacciServer) error
	MaxEvenNumber(CalculatorService_MaxEvenNumberServer) error
	VowelsCount(CalculatorService_VowelsCountServer) error
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorServiceServer) Primes(*PrimeRequest, CalculatorService_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCalculatorServiceServer) Average(CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServiceServer) Max(CalculatorService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCalculatorServiceServer) Subtract(context.Context, *SubRequest) (*SubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedCalculatorServiceServer) Fibonacci(*FibRequest, CalculatorService_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedCalculatorServiceServer) MaxEvenNumber(CalculatorService_MaxEvenNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method MaxEvenNumber not implemented")
}
func (UnimplementedCalculatorServiceServer) VowelsCount(CalculatorService_VowelsCountServer) error {
	return status.Errorf(codes.Unimplemented, "method VowelsCount not implemented")
}
func (UnimplementedCalculatorServiceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Primes(m, &calculatorServicePrimesServer{stream})
}

type CalculatorService_PrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServicePrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Max(&calculatorServiceMaxServer{stream})
}

type CalculatorService_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Subtract(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Fibonacci(m, &calculatorServiceFibonacciServer{stream})
}

type CalculatorService_FibonacciServer interface {
	Send(*FibResponse) error
	grpc.ServerStream
}

type calculatorServiceFibonacciServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFibonacciServer) Send(m *FibResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_MaxEvenNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).MaxEvenNumber(&calculatorServiceMaxEvenNumberServer{stream})
}

type CalculatorService_MaxEvenNumberServer interface {
	SendAndClose(*EvenResponse) error
	Recv() (*EvenRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxEvenNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxEvenNumberServer) SendAndClose(m *EvenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxEvenNumberServer) Recv() (*EvenRequest, error) {
	m := new(EvenRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_VowelsCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).VowelsCount(&calculatorServiceVowelsCountServer{stream})
}

type CalculatorService_VowelsCountServer interface {
	Send(*VowelResponse) error
	Recv() (*VowelRequest, error)
	grpc.ServerStream
}

type calculatorServiceVowelsCountServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceVowelsCountServer) Send(m *VowelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceVowelsCountServer) Recv() (*VowelRequest, error) {
	m := new(VowelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Sqrt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalculatorService_Subtract_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalculatorService_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _CalculatorService_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CalculatorService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Fibonacci",
			Handler:       _CalculatorService_Fibonacci_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MaxEvenNumber",
			Handler:       _CalculatorService_MaxEvenNumber_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "VowelsCount",
			Handler:       _CalculatorService_VowelsCount_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
