package main

import (
	"context"

	"github.com/Snake1-1eyes/Yandex_Service/pkg/logger"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

}
