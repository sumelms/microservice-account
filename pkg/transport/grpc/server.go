package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/sumelms/sumelms/microservice-user/pkg/endpoint"
	"github.com/sumelms/sumelms/microservice-user/proto"
)

type server struct {
	createUser grpc.Handler
	getUser    grpc.Handler
}

func (s server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CreateUserResponse), nil
}

func (s server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.GetUserResponse), nil
}

func NewGrpcServer(ctx context.Context, endpoints endpoint.Endpoints) proto.UserServer {
	return &server{
		createUser: grpc.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse),
		getUser: grpc.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse),
	}
}

func decodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.CreateUserRequest)
	return endpoint.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password}, nil
}

func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.CreateUserResponse)
	return &proto.CreateUserResponse{Ok: res.Ok}, nil
}

func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.GetUserRequest)
	return endpoint.GetUserRequest{Id: req.Id}, nil
}

func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetUserResponse)
	return &proto.GetUserResponse{Email: res.Email}, nil
}
