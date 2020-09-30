package grpc

import (
	"github.com/go-kit/kit/transport/grpc"
)

type server struct {
	createUser grpc.Handler
	getUser    grpc.Handler
	updateUser grpc.Handler
	deleteUser grpc.Handler
	listUsers  grpc.Handler
}
