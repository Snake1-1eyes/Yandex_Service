package service

import test "github.com/Snake1-1eyes/Yandex_Service/pkg/api/test/api"

type Service struct {
	test.OrderServiceServer
}

func New() *Service {
	return &Service{}
}
