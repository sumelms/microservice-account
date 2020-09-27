package grpc

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/sumelms/microservice-account/pkg/endpoint"
	"github.com/sumelms/microservice-account/proto"
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

	var user proto.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &proto.CreateUserResponse{User: &user}, nil
}

func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.GetUserRequest)
	return endpoint.GetUserRequest{Id: req.Id}, nil
}
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.GetUserResponse)

	var user proto.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &proto.GetUserResponse{User: &user}, nil
}

func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.UpdateUserRequest)

	return endpoint.UpdateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.UpdateUserResponse)

	var user proto.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &proto.UpdateUserResponse{User: &user}, nil
}

func decodeDeleteUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.DeleteUserRequest)
	return endpoint.DeleteUserRequest{Id: req.Id}, nil
}
func encodeDeleteUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.DeleteUserResponse)
	return &proto.DeleteUserResponse{Id: res.Id}, nil
}

func decodeListUsersRequest(_ context.Context, r interface{}) (interface{}, error) {
	// @TODO Pagination and Filter
	//req := r.(*proto.ListUsersRequest)
	return endpoint.ListUsersRequest{}, nil
}
func encodeListUsersResponse(c context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ListUsersResponse)

	var users []*proto.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &users)

	return &proto.ListUsersResponse{Users: users}, nil
}
