package controllers

import (
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

func (ctrl *BookController) AddBook(c *gin.Context) {
	var bookInfo models.Book
	c.BindJSON(&bookInfo)
	err := models.AddBookModel(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "添加失敗",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "添加成功",
			"success": true,
		})
	}
	return
}
