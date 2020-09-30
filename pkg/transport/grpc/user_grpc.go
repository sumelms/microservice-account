package grpc

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/sumelms/microservice-account/pkg/endpoint/user"
	protouser "github.com/sumelms/microservice-account/proto/user"
)

func (s server) CreateUser(ctx context.Context, req *protouser.CreateUserRequest) (*protouser.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*protouser.CreateUserResponse), nil
}

func (s server) GetUser(ctx context.Context, req *protouser.GetUserRequest) (*protouser.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*protouser.GetUserResponse), nil
}

func (s server) UpdateUser(ctx context.Context, req *protouser.UpdateUserRequest) (*protouser.UpdateUserResponse, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*protouser.UpdateUserResponse), nil
}

func (s server) DeleteUser(ctx context.Context, req *protouser.DeleteUserRequest) (*protouser.DeleteUserResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*protouser.DeleteUserResponse), nil
}

func (s server) ListUsers(ctx context.Context, req *protouser.ListUsersRequest) (*protouser.ListUsersResponse, error) {
	_, resp, err := s.listUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*protouser.ListUsersResponse), nil
}

func NewGrpcServer(_ context.Context, endpoints user.Endpoints) protouser.UserServer {
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
	req := r.(*protouser.CreateUserRequest)
	return user.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password}, nil
}
func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(user.CreateUserResponse)

	var user protouser.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &protouser.CreateUserResponse{User: &user}, nil
}

func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*protouser.GetUserRequest)
	return user.GetUserRequest{Id: req.Id}, nil
}
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(user.GetUserResponse)

	var user protouser.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &protouser.GetUserResponse{User: &user}, nil
}

func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*protouser.UpdateUserRequest)

	return user.UpdateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(user.UpdateUserResponse)

	var user protouser.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &user)

	return &protouser.UpdateUserResponse{User: &user}, nil
}

func decodeDeleteUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*protouser.DeleteUserRequest)
	return user.DeleteUserRequest{Id: req.Id}, nil
}
func encodeDeleteUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(user.DeleteUserResponse)
	return &protouser.DeleteUserResponse{Id: res.Id}, nil
}

func decodeListUsersRequest(_ context.Context, r interface{}) (interface{}, error) {
	// @TODO Pagination and Filter
	//req := r.(*proto.ListUsersRequest)
	return user.ListUsersRequest{}, nil
}
func encodeListUsersResponse(c context.Context, r interface{}) (interface{}, error) {
	res := r.(user.ListUsersResponse)

	var users []*protouser.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &users)

	return &protouser.ListUsersResponse{Users: users}, nil
}

