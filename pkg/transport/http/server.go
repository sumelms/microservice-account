package http

import (
	"context"
	"encoding/json"
	"github.com/sumelms/microservice-account/pkg/adapter/middleware"
	"github.com/sumelms/microservice-account/pkg/endpoint"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.JsonEncodeMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/user").Handler(httptransport.NewServer(
		endpoints.ListUsers,
		decodeListUsersRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateUserRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUser,
		decodeDeleteUserRequest,
		encodeResponse,
	))

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetUserRequest
	vars := mux.Vars(r)

	req = endpoint.GetUserRequest{Id: vars["id"]}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	req = endpoint.UpdateUserRequest{
		Id:              vars["id"],
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.DeleteUserRequest
	vars := mux.Vars(r)

	req = endpoint.DeleteUserRequest{Id: vars["id"]}

	return req, nil
}

func decodeListUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ListUsersRequest{}
	return req, nil
}
