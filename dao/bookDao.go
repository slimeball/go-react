package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func AddBookModel(book *model.Book) error {
	return utils.DB.Table("book").Create(&book).Error
}

func GetBookByIdDAO(id int) (*model.Book, error) {
	var book model.Book
	err := utils.DB.Table("book").Where("id=?", id).Find(&book).Error
	return &book, err
}

func GetBookListDAO(page *model.Page) (*model.Page, error) {
	var PageNo int64
	var PageSize int64
	var PageTotal int64
	var PageTotalNum int64
	PageNo = page.PageNo
	PageSize = page.PageSize

	utils.DB.Table("book").Count(&PageTotal)

	var book []*model.Book
	err := utils.DB.Table("book").Offset((PageNo - 1) * PageSize).Limit(PageSize).Find(&book).Error
	if PageTotal%PageSize == 0 {
		PageTotalNum = PageTotal / PageSize
	} else {
		PageTotalNum = PageTotal/PageSize + 1
	}
	pageStruct := &model.Page{
		Book:         book,
		PageNo:       PageNo,
		PageSize:     PageSize,
		PageTotalNum: PageTotalNum,
		PageTotal:    PageTotal,
	}
	return pageStruct, err
}

func UpdateBookDAO(book *model.Book) error {
	err := utils.DB.Table("book").Model(&book).Save(&book).Error
	return err
}

func DeleteBookDAO(book *model.Book) error {
	err := utils.DB.Table("book").Delete(&book).Error
	return err
}
