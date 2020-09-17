package endpoint

import (
	"context"
	"github.com/sumelms/microservice-user/pkg/domain"

	"github.com/go-kit/kit/endpoint"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s domain.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)

		email, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email: email,
		}, err
	}
}
