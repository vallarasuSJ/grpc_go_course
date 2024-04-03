// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: student.proto

package proto

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

const (
	StudentService_CreateStudent_FullMethodName = "/student.StudentService/CreateStudent"
	StudentService_ReadStudent_FullMethodName   = "/student.StudentService/ReadStudent"
	StudentService_UpdateStudent_FullMethodName = "/student.StudentService/UpdateStudent"
	StudentService_ListStudents_FullMethodName  = "/student.StudentService/ListStudents"
	StudentService_DeleteStudent_FullMethodName = "/student.StudentService/DeleteStudent"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *StudentDetails, opts ...grpc.CallOption) (*StudentId, error)
	ReadStudent(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentDetails, error)
	UpdateStudent(ctx context.Context, in *StudentDetails, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListStudents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StudentService_ListStudentsClient, error)
	DeleteStudent(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *StudentDetails, opts ...grpc.CallOption) (*StudentId, error) {
	out := new(StudentId)
	err := c.cc.Invoke(ctx, StudentService_CreateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ReadStudent(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentDetails, error) {
	out := new(StudentDetails)
	err := c.cc.Invoke(ctx, StudentService_ReadStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *StudentDetails, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, StudentService_UpdateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StudentService_ListStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[0], StudentService_ListStudents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceListStudentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentService_ListStudentsClient interface {
	Recv() (*StudentDetails, error)
	grpc.ClientStream
}

type studentServiceListStudentsClient struct {
	grpc.ClientStream
}

func (x *studentServiceListStudentsClient) Recv() (*StudentDetails, error) {
	m := new(StudentDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, StudentService_DeleteStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	CreateStudent(context.Context, *StudentDetails) (*StudentId, error)
	ReadStudent(context.Context, *StudentId) (*StudentDetails, error)
	UpdateStudent(context.Context, *StudentDetails) (*emptypb.Empty, error)
	ListStudents(*emptypb.Empty, StudentService_ListStudentsServer) error
	DeleteStudent(context.Context, *StudentId) (*emptypb.Empty, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *StudentDetails) (*StudentId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) ReadStudent(context.Context, *StudentId) (*StudentDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadStudent not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudent(context.Context, *StudentDetails) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(*emptypb.Empty, StudentService_ListStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudent(context.Context, *StudentId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_CreateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*StudentDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ReadStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ReadStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ReadStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ReadStudent(ctx, req.(*StudentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_UpdateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudent(ctx, req.(*StudentDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).ListStudents(m, &studentServiceListStudentsServer{stream})
}

type StudentService_ListStudentsServer interface {
	Send(*StudentDetails) error
	grpc.ServerStream
}

type studentServiceListStudentsServer struct {
	grpc.ServerStream
}

func (x *studentServiceListStudentsServer) Send(m *StudentDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _StudentService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_DeleteStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudent(ctx, req.(*StudentId))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "ReadStudent",
			Handler:    _StudentService_ReadStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentService_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentService_DeleteStudent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStudents",
			Handler:       _StudentService_ListStudents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "student.proto",
}