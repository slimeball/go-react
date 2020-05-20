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
}

// get one book
func GetBookById(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	book, err := dao.GetBookByIdModel(bookInfo.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "查找失敗",
			"data":    book,
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "查找成功",
			"data":    book,
			"success": true,
		})
	}
}

// receive book list
func GetBookList(c *gin.Context) {
	// var bookInfo model.Book
	// c.BindJSON(&bookInfo)
	list, err := dao.GetBookListModel()
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "添加失敗",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "查詢成功",
			"data":    list,
			"success": true,
		})
	}
}

// update book
func UpdateBook(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	err := dao.UpdateBookModel(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "修改失敗",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "修改成功",
			"data":    bookInfo,
			"success": true,
		})
	}
}

// delete book
func DeleteBook(c *gin.Context) {
	var bookInfo model.Book
	c.Bind(&bookInfo)
	err := dao.DeleteBookModel(&bookInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "刪除失敗",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "刪除成功",
			"success": true,
		})
	}
}
