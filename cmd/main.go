package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/Snake1-1eyes/Yandex_Service/internal/config"
	"github.com/Snake1-1eyes/Yandex_Service/internal/service"
	test "github.com/Snake1-1eyes/Yandex_Service/pkg/api/test/api"
	"github.com/Snake1-1eyes/Yandex_Service/pkg/logger"
	"github.com/Snake1-1eyes/Yandex_Service/pkg/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	ctx, _ = logger.New(ctx)

	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to load config", zap.Error(err))
	}

	pool, err := postgres.New(ctx, &cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect to database: %w", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer(grpc.UnaryInterceptor(logger.InterceptorWithLogger(ctx, logger.GetLoggerFromCtx(ctx))))
	test.RegisterOrderServiceServer(server, srv)

	rt := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = test.RegisterOrderServiceHandlerFromEndpoint(ctx, rt, "localhost:"+strconv.Itoa(cfg.GRPCPort), opts)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to register handler server", zap.Error(err))
	}

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.RestPORT), rt); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve: %w", zap.Error(err))
		}
	}()

	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve: %w", zap.Error(err))
		}
	}()

	select {
	case <-ctx.Done():
		server.GracefulStop()
		pool.Close()
		logger.GetLoggerFromCtx(ctx).Info(ctx, "Server stoped")
	}
}
