package grpc

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/sumelms/microservice-account/pkg/endpoint/user"
	protouser "github.com/sumelms/microservice-account/proto/user"
)

type server struct {
	createUser grpc.Handler
	getUser    grpc.Handler
	updateUser grpc.Handler
	deleteUser grpc.Handler
	listUsers  grpc.Handler
	// Embed the unimplemented server
	protouser.UnimplementedUserServer
}

// NewGRPCServer reates grpc server
func NewGRPCServer(_ context.Context, endpoints user.Endpoints) protouser.UserServer {
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
