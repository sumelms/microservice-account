package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/sumelms/microservice-user/pkg/endpoint"
	"github.com/sumelms/microservice-user/proto"
)

type server struct {
	createUser grpc.Handler
	getUser    grpc.Handler
	updateUser grpc.Handler
	deleteUser grpc.Handler
	listUsers  grpc.Handler
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

func (s server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.UpdateUserResponse), nil
}

func (s server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.DeleteUserResponse), nil
}

func (s server) ListUsers(ctx context.Context, req *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	_, resp, err := s.listUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.ListUsersResponse), nil
}

func NewGrpcServer(_ context.Context, endpoints endpoint.Endpoints) proto.UserServer {
	return &server{
		createUser: grpc.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse),
		getUser: grpc.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse),
		updateUser: grpc.NewServer(
			endpoints.UpdateUser,
			decodeUpdateUserRequest,
			encodeUpdateUserResponse),
		deleteUser: grpc.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUserRequest,
			encodeDeleteUserResponse),
		listUsers: grpc.NewServer(
			endpoints.ListUsers,
			decodeListUsersRequest,
			encodeListUsersResponse),
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
	return &proto.CreateUserResponse{}, nil
}

func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.GetUserRequest)
	return endpoint.GetUserRequest{Id: req.Id}, nil
}
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetUserResponse)
	return &proto.GetUserResponse{}, nil
}

func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*proto.UpdateUserRequest)
	return endpoint.UpdateUserRequest{}, nil
}
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.UpdateUserResponse)
	return &proto.UpdateUserResponse{}, nil
}

func decodeDeleteUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*proto.DeleteUserRequest)
	return endpoint.DeleteUserRequest{}, nil
}
func encodeDeleteUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.DeleteUserResponse)
	return &proto.DeleteUserResponse{}, nil
}

func decodeListUsersRequest(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*proto.ListUsersRequest)
	return endpoint.ListUsersRequest{}, nil
}
func encodeListUsersResponse(c context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ListUsersResponse)
	return &proto.ListUsersResponse{}, nil
}
