package trace

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	CtxKeyTraceID = "_ctx_key_trace_id"
)

func WithTraceID(c *gin.Context) {

}

func GetTraceID(c *gin.Context) string {
	return ""
}

func GetTraceIDByCtx(ctx context.Context) string {
	return ""
}
