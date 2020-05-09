package controllers

import (
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// operate signup
func (ctrl *UserController) SignIn(c *gin.Context) {
	var userinfo models.Users
	c.BindJSON(&userinfo)
	result, err := models.CheckUserExist(userinfo.Username)
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
	return
}

// operate signup
func (ctrl *UserController) SignUp(c *gin.Context) {
	var registerInfo models.Users
	c.BindJSON(&registerInfo)

	err := models.UserRegister(&registerInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "用戶名已存在",
			"success": false,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "註冊成功",
			"success": true,
		})
	}
	return
}
