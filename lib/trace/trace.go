package trace

import "github.com/gin-gonic/gin"

const (
	CtxKeyTraceID = "_ctx_key_trace_id"
)

func WithTraceID(c *gin.Context) {

}

func GetTraceID(c *gin.Context) string {
	return ""
}
