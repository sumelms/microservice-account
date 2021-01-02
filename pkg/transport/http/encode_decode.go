package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-account/pkg/endpoint/user"
)

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.GetUserRequest
	vars := mux.Vars(r)

	req = user.GetUserRequest{Id: vars["id"]}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	req = user.UpdateUserRequest{
		Id:              vars["id"],
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.DeleteUserRequest
	vars := mux.Vars(r)

	req = user.DeleteUserRequest{Id: vars["id"]}

	return req, nil
}

func decodeListUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := user.ListUsersRequest{}
	return req, nil
}
