package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Snake1-1eyes/Yandex_Service/internal/service"
	test "github.com/Snake1-1eyes/Yandex_Service/pkg/api/test/api"
	"github.com/Snake1-1eyes/Yandex_Service/pkg/logger"
	"github.com/Snake1-1eyes/Yandex_Service/pkg/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	pgConfig := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "1234",
		Database: "postgres",
	}

	_, err := postgres.New(&pgConfig)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect to database: %w", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer()
	test.RegisterOrderServiceServer(server, srv)

	if err := server.Serve(lis); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve: %w", zap.Error(err))
	}
}
