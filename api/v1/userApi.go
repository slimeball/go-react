package v1

import (
	"bookstore/dao"
	"bookstore/model"

	"github.com/gin-gonic/gin"
)

// operate signup
func SignIn(c *gin.Context) {
	var userinfo model.Users
	c.BindJSON(&userinfo)
	result, err := dao.CheckUserExist(userinfo.Username)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "用戶名不存在",
			"success": false,
		})
	} else {
		if userinfo.Password == result.Password {
			c.JSON(200, gin.H{
				"code":     "1",
				"message":  "登錄成功",
				"userinfo": result,
				"success":  true,
			})
		} else {
			c.JSON(200, gin.H{
				"code":    "0",
				"message": "用戶名或密碼錯誤",
				"success": false,
			})
		}
	}
}

// operate signup
func SignUp(c *gin.Context) {
	var registerInfo model.Users
	c.Bind(&registerInfo)

	err := dao.UserRegister(&registerInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": registerInfo,
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "註冊成功",
			"success": true,
		})
	}
}
