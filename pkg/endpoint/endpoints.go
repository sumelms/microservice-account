package endpoint

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/sumelms/microservice-user/pkg/domain"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		User *domain.User `json:"user"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		User *domain.User `json:"user"`
	}

	UpdateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserResponse struct {
		User *domain.User `json:"user"`
	}

	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Id string `json:"id"`
	}

	ListUsersRequest  struct{}
	ListUsersResponse struct {
		Users *[]domain.User `json:"users"`
	}
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	ListUsers  endpoint.Endpoint
}

func MakeEndpoints(s domain.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
		ListUsers:  makeDeleteUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		data, _ := json.Marshal(req)
		user := domain.User{}
		json.Unmarshal([]byte(data), &user)

		ok, err := s.CreateUser(ctx, &user)

		return CreateUserResponse{User: ok}, err
	}
}

func makeGetUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)

		user, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{User: user}, err
	}
}

func makeUpdateUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)

		data, _ := json.Marshal(req)
		user := domain.User{}
		json.Unmarshal([]byte(data), &user)

		updated, err := s.UpdateUser(ctx, &user)

		return UpdateUserResponse{User: updated}, err
	}
}

func makeDeleteUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)

		err := s.DeleteUser(ctx, req.Id)

		return DeleteUserResponse{Id: req.Id}, err
	}
}

func makeListUsers(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// @TODO Pagination and filters
		// req := request.(ListUsersRequest)

		users, err := s.ListUsers(ctx)

		return ListUsersResponse{Users: &users}, err
	}
}
