package grpc

import (
	"context"

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
