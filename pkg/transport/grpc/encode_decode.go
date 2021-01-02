package grpc

import (
	"context"
	"encoding/json"

	"github.com/sumelms/microservice-account/pkg/endpoint/user"
	protouser "github.com/sumelms/microservice-account/proto/user"
)

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
	return user.ListUsersRequest{}, nil
}
func encodeListUsersResponse(c context.Context, r interface{}) (interface{}, error) {
	res := r.(user.ListUsersResponse)

	var users []*protouser.UserModel
	data, _ := json.Marshal(res)
	json.Unmarshal([]byte(data), &users)

	return &protouser.ListUsersResponse{Users: users}, nil
}
