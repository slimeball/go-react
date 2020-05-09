package models

type Book struct {
	Id        int
	Bookname  string  `form:"bookname" json:"bookname" binding:"required"`
	Author    string  `form:"author" json:"author" binding:"required"`
	Price     float64 `form:"price" json:"price" binding:"required"`
	Sales     int     `form:"sales" json:"sales" binding:"required"`
	Inventory int     `form:"inventory" json:"inventory" binding:"required"`
	ImgPath   string  `form:"img_path" json:"img_path" binding:"required"`
}

func AddBookModel(book *Book) error {
	return DB.Table("book").Create(&book).Error
}
