package http

import (
	"context"
	"net/http"

	"github.com/sumelms/microservice-account/pkg/endpoint/user"
	"github.com/sumelms/microservice-account/pkg/middleware"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer creates http server router
func NewHTTPServer(ctx context.Context, endpoints user.Endpoints) http.Handler {
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
