package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/sumelms/microservice-account/pkg/domain/user"
	"github.com/sumelms/microservice-account/pkg/validator"
)

type (
	CreateUserRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,alphanum,min=6"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}
	CreateUserResponse struct {
		ID string `json:"id"`
	}

	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	UpdateUserRequest struct {
		ID              string `json:"id"`
		Email           string `json:"email" validate:"email"`
		Password        string `json:"password" validate:"alphanum,min=6"`
		ConfirmPassword string `json:"confirm_password" validate:"required_with=Password,eqfield=Password"`
	}
	UpdateUserResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	DeleteUserRequest struct {
		ID string `json:"id"`
	}
	DeleteUserResponse struct {
		ID string `json:"id"`
	}

	ListUsersRequest  struct{}
	ListUsersResponse struct {
		Users *[]user.User `json:"users"`
	}
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	ListUsers  endpoint.Endpoint
}

func MakeEndpoints(s user.ServiceInterface) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
		ListUsers:  makeListUsersEndpoint(s),
	}
}

func makeCreateUserEndpoint(s user.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		u := user.User{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &u)
		if err != nil {
			return nil, err
		}

		ok, err := s.CreateUser(ctx, &u)
		if err != nil {
			return nil, err
		}

		return CreateUserResponse{ID: ok.ID.String()}, err
	}
}

func makeGetUserEndpoint(s user.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)

		if req.ID == "" {
			return nil, errors.New("bad request, missing param id")
		}

		u, err := s.GetUser(ctx, req.ID)

		return GetUserResponse{
			ID:    u.ID.String(),
			Email: u.Email,
		}, err
	}
}

func makeUpdateUserEndpoint(s user.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		u := user.User{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &u)
		if err != nil {
			return nil, err
		}

		updated, err := s.UpdateUser(ctx, &u)

		return UpdateUserResponse{
			ID:    updated.ID.String(),
			Email: updated.Email,
		}, err
	}
}

func makeDeleteUserEndpoint(s user.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)

		if req.ID == "" {
			return nil, errors.New("bad request, missing id param")
		}

		err := s.DeleteUser(ctx, req.ID)

		return request.(DeleteUserResponse), err
	}
}

func makeListUsersEndpoint(s user.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.ListUsers(ctx)

		return ListUsersResponse{Users: &users}, err
	}
}
