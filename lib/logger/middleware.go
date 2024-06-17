package logger

import (
	"time"

	"_gdc_/lib/trace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      // 请求路径 eg: /test
		query := c.Request.URL.RawQuery //query类型的请求参数：?name=1&password=2
		trace.WithTraceID(c)            // 设置traceID

		c.Next()

		traceID := trace.GetTraceID(c)
		cost := time.Since(start)
		status := c.Writer.Status()

		zapLogger.Info(path,
			zap.Int("status", status),              // 状态码 eg: 200
			zap.String("method", c.Request.Method), // 请求方法类型 eg: GET
			zap.String("path", path),               // 请求路径 eg: /test
			zap.String("query", query),             // 请求参数 eg: name=1&password=2
			zap.String("ip", c.ClientIP()),         // 返回真实的客户端IP eg: ::1（这个就是本机IP，ipv6地址）
			zap.Duration("cost", cost),             // 返回花费时间
			zap.String("trace_id", traceID),        // traceID
		)
	}
}
