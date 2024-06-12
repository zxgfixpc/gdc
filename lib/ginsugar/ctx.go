package ginsugar

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Context(c *gin.Context) context.Context {
	return c.Request.Context()
}
