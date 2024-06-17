package service

import (
	"context"

	"_gdc_/lib/logger"
)

func Hello(ctx context.Context) (string, error) {
	logger.Info(ctx, "hello")
	logger.Error(ctx, "hello error")
	return "hello", nil
}
