package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Snake1-1eyes/Yandex_Service/internal/config"
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

	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to load config", zap.Error(err))
	}

	_, err = postgres.New(&cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect to database: %w", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
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
