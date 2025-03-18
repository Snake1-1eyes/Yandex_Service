package service

import (
	"context"

	test "github.com/Snake1-1eyes/Yandex_Service/pkg/api/test/api"
	"github.com/Snake1-1eyes/Yandex_Service/pkg/logger"
)

type Service struct {
	test.OrderServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	return &test.CreateOrderResponse{}, nil
}

func (s *Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	return &test.GetOrderResponse{}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	return &test.UpdateOrderResponse{}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	return &test.DeleteOrderResponse{}, nil
}

func (s *Service) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	logger.GetLoggerFromCtx(ctx).Info(ctx, "list ordeers")
	return &test.ListOrdersResponse{}, nil
}
