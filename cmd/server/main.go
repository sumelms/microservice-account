package main

import (
	"context"
	"fmt"
	user "github.com/sumelms/microservice-account/pkg/database/gorm/user"
	userendpoint "github.com/sumelms/microservice-account/pkg/endpoint/user"
	protouser "github.com/sumelms/microservice-account/proto/user"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sumelms/microservice-account/pkg/config"
	grpctransport "github.com/sumelms/microservice-account/pkg/transport/grpc"
	httptransport "github.com/sumelms/microservice-account/pkg/transport/http"

	database "github.com/sumelms/microservice-account/pkg/database/gorm"
	userdomain "github.com/sumelms/microservice-account/pkg/domain/user"
	"github.com/sumelms/microservice-account/pkg/logger"

	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// @TODO Split this file into http and grpc files
func main() {
	// Logger
	logger := logger.NewLogger()

	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	ctx := context.Background()
	repository := user.NewRepository(db, logger)
	srv := userdomain.NewService(repository, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := userendpoint.MakeEndpoints(srv)

	// HTTP Server
	go func() {
		fmt.Println("HTTP Server Listening on", cfg.Server.Http.Host)

		httpServer := httptransport.NewHttpServer(ctx, endpoints)

		errs <- http.ListenAndServe(cfg.Server.Http.Host, httpServer)
	}()

	// gRPC Server
	go func() {
		listener, err := net.Listen("tcp", cfg.Server.Grpc.Host)
		if err != nil {
			errs <- err
			return
		}

		fmt.Println("gRPC Server Listening on", cfg.Server.Grpc.Host)

		handler := grpctransport.NewGrpcServer(ctx, endpoints)
		grpcServer := grpc.NewServer()

		protouser.RegisterUserServer(grpcServer, handler)
		reflection.Register(grpcServer)

		errs <- grpcServer.Serve(listener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
