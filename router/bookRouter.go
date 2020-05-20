package router

import (
	v1 "bookstore/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBookRouter(Router *gin.RouterGroup) {
	Router.POST("/addbook", v1.AddBook)
	Router.GET("/getbookbyid", v1.GetBookById)
	Router.GET("/getbooklist", v1.GetBookList)
	Router.POST("/updatebook", v1.UpdateBook)
	Router.POST("/deletebook", v1.DeleteBook)
}
