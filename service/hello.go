package service

import "context"

func Hello(ctx context.Context) (string, error) {
	return "hello", nil
}
