package router

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := &controllers.UserController{}
	router.POST("/signin", userController.SignIn)
	router.POST("/signup", userController.SignUp)
	bookController := &controllers.BookController{}
	router.POST("/addbook", bookController.AddBook)
}
