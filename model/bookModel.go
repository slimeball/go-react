package model

type Book struct {
	Id        int     `form:"id" json:"id"`
	Bookname  string  `form:"bookname" json:"bookname"`
	Author    string  `form:"author" json:"author"`
	Price     float64 `form:"price" json:"price"`
	Sales     int     `form:"sales" json:"sales"`
	Inventory int     `form:"inventory" json:"inventory"`
	ImgPath   string  `form:"img_path" json:"img_path"`
}
