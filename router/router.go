package router

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1Router := router.Group("/v1")
	{
		InitUserRouter(v1Router)
		InitBookRouter(v1Router)
	}
}
