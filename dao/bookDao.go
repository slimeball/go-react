package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddBookModel(book *model.Book) error {
	return utils.DB.Table("book").Create(&book).Error
}

func GetBookByIdModel(id int) (model.Book, error) {
	var book model.Book
	err := utils.DB.Table("book").Where("id=?", id).Find(&book).Error
	return book, err
}

func GetBookListModel() ([]*model.Book, error) {
	var book []*model.Book
	err := utils.DB.Table("book").Find(&book).Error
	return book, err
}

func UpdateBookModel(book *model.Book) error {
	err := utils.DB.Table("book").Model(&book).Save(&book).Error
	return err
}

func DeleteBookModel(book *model.Book) error {
	err := utils.DB.Table("book").Delete(&book).Error
	return err
}
