package controller

import (
	"gdc/lib/ginsugar"
	"gdc/service"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	rlt, err := service.Hello(ginsugar.Context(c))
	if err != nil {
		ginsugar.Fail(c, rlt, err)
		return
	}
	ginsugar.Success(c, rlt)
}
