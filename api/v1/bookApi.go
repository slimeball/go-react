package v1

import (
	"bookstore/dao"
	"bookstore/model"

	"github.com/gin-gonic/gin"
)

// add new book
func AddBook(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	err := dao.AddBookModel(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "add failed.",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "add success.",
			"success": true,
		})
	}
}

// get one book
func GetBookById(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	book, err := dao.GetBookByIdDAO(bookInfo.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "search failed.",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "search success.",
			"data":    book,
			"success": true,
		})
	}
}

// receive book list
func GetBookList(c *gin.Context) {
	var pageInfo model.Page
	c.Bind(&pageInfo)

	list, err := dao.GetBookListDAO(&pageInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "search failed.",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "search success.",
			"data":    list,
			"success": true,
		})
	}
}

// update book
func UpdateBook(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	err := dao.UpdateBookDAO(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "update failed.",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "update success.",
			"data":    bookInfo,
			"success": true,
		})
	}
}

// delete book
func DeleteBook(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	err := dao.DeleteBookDAO(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "delete failed.",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "delete success",
			"success": true,
		})
	}
}
