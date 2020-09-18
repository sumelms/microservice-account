package main

import (
	"context"
	"fmt"
	endpoints "github.com/sumelms/microservice-user/pkg/endpoint"
	"github.com/sumelms/microservice-user/proto"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sumelms/microservice-user/pkg/config"
	grpctransport "github.com/sumelms/microservice-user/pkg/transport/grpc"
	httptransport "github.com/sumelms/microservice-user/pkg/transport/http"

	database "github.com/sumelms/microservice-user/pkg/adapter/database/gorm"
	"github.com/sumelms/microservice-user/pkg/adapter/logger"
	"github.com/sumelms/microservice-user/pkg/domain"

	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Logger
	logger := logger.NewLogger(cfg)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	ctx := context.Background()
	var srv domain.Service
	{
		repository := database.NewRepository(db, logger)
		srv = domain.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := endpoints.MakeEndpoints(srv)

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

		proto.RegisterUserServer(grpcServer, handler)
		reflection.Register(grpcServer)

		errs <- grpcServer.Serve(listener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
