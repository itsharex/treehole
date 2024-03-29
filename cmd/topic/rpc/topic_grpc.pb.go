// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: topic.proto

package rpc

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
	TopicService_CreateTopic_FullMethodName    = "/proto.TopicService/CreateTopic"
	TopicService_GetTopic_FullMethodName       = "/proto.TopicService/GetTopic"
	TopicService_PutStar_FullMethodName        = "/proto.TopicService/PutStar"
	TopicService_GetStarList_FullMethodName    = "/proto.TopicService/GetStarList"
	TopicService_GetCommentList_FullMethodName = "/proto.TopicService/GetCommentList"
	TopicService_AddComment_FullMethodName     = "/proto.TopicService/AddComment"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error)
	PutStar(ctx context.Context, in *PutStarReq, opts ...grpc.CallOption) (*PutStarResp, error)
	GetStarList(ctx context.Context, in *GetStarListReq, opts ...grpc.CallOption) (*GetStarListResp, error)
	GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListResp, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_CreateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicResponse, error) {
	out := new(GetTopicResponse)
	err := c.cc.Invoke(ctx, TopicService_GetTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) PutStar(ctx context.Context, in *PutStarReq, opts ...grpc.CallOption) (*PutStarResp, error) {
	out := new(PutStarResp)
	err := c.cc.Invoke(ctx, TopicService_PutStar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetStarList(ctx context.Context, in *GetStarListReq, opts ...grpc.CallOption) (*GetStarListResp, error) {
	out := new(GetStarListResp)
	err := c.cc.Invoke(ctx, TopicService_GetStarList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListResp, error) {
	out := new(GetCommentListResp)
	err := c.cc.Invoke(ctx, TopicService_GetCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, TopicService_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations should embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
	GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error)
	PutStar(context.Context, *PutStarReq) (*PutStarResp, error)
	GetStarList(context.Context, *GetStarListReq) (*GetStarListResp, error)
	GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListResp, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
}

// UnimplementedTopicServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedTopicServiceServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedTopicServiceServer) PutStar(context.Context, *PutStarReq) (*PutStarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStar not implemented")
}
func (UnimplementedTopicServiceServer) GetStarList(context.Context, *GetStarListReq) (*GetStarListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStarList not implemented")
}
func (UnimplementedTopicServiceServer) GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedTopicServiceServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_PutStar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutStarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).PutStar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_PutStar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).PutStar(ctx, req.(*PutStarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetStarList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStarListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetStarList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetStarList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetStarList(ctx, req.(*GetStarListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_GetCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetCommentList(ctx, req.(*GetCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "PutStar",
			Handler:    _TopicService_PutStar_Handler,
		},
		{
			MethodName: "GetStarList",
			Handler:    _TopicService_GetStarList_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _TopicService_GetCommentList_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _TopicService_AddComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topic.proto",
}
