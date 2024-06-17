package service

import (
	"context"

	"_gdc_/lib/log"
)

func Hello(ctx context.Context) (string, error) {
	log.Info(ctx, "hello")
	log.Error(ctx, "hello error")
	return "hello", nil
}
