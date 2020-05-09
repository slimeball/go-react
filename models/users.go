package models

type Users struct {
	Id       int
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
}

func CheckUserExist(username string) (*Users, error) {
	var users Users
	err = DB.Where("username=?", username).First(&users).Error
	return &users, err
}

func UserRegister(user *Users) error {
	return DB.Create(&user).Error
}

// func ValidateUsernameAndPassword(username string, password string) *User {
// 	sqlStr := "SELECT id, username, password, email FROM users WHERE username = ? and password = ?"
// 	row := DB.QueryRow(sqlStr, username, password)
// 	user := &User{}
// 	row.Scan(user.ID, user.Username, user.Password, user.Email)
// 	return user
// }
