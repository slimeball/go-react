package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func CreateSession(sess *model.Book) error {
	return utils.DB.Table("sessions").Create(&sess).Error
}

func DelecteSession(sess *model.SessionInfo) error {
	return utils.DB.Table("sessions").Delete(&sess).Error
}
