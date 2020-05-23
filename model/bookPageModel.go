package model

// 前端傳過來的就是int64 定義int需要強轉
type Page struct {
	Book         []*Book `form:"book" json:"book"`
	PageNo       int64   `form:"pageno" json:"pageno"`
	PageSize     int64   `form:"pagesize" json:"pagesize"`
	PageTotal    int64   `form:"pagetotal" json:"pagetotal"`
	PageTotalNum int64   `form:"pagetotalnum" json:"pagetotalnum"`
}
