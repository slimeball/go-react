package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func CheckUserExist(username string) (*model.Users, error) {
	var users model.Users
	err := utils.DB.Where("username=?", username).First(&users).Error
	return &users, err
}

func UserRegister(user *model.Users) error {
	return utils.DB.Create(&user).Error
}

// func ValidateUsernameAndPassword(username string, password string) *User {
// 	sqlStr := "SELECT id, username, password, email FROM users WHERE username = ? and password = ?"
// 	row := DB.QueryRow(sqlStr, username, password)
// 	user := &User{}
// 	row.Scan(user.ID, user.Username, user.Password, user.Email)
// 	return user
// }
