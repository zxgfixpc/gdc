package httpserver

import (
	"gdc/controller"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	testGroup := router.Group("test/")
	{
		testGroup.GET("hello", controller.Hello)
	}
}
