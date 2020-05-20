package router

import (
	v1 "bookstore/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	Router.POST("/signin", v1.SignIn)
	Router.POST("/signup", v1.SignUp)
}
