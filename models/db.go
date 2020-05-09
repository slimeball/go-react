package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
