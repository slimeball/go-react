package v1

import (
	"bookstore/dao"
	"bookstore/model"

	"github.com/gin-contrib/sessions"
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
			"message": "username doesn't exsit.",
			"success": false,
		})
	} else {
		if userinfo.Password == result.Password {
			c.JSON(200, gin.H{
				"code":     "1",
				"message":  "sign in success.",
				"userinfo": result,
				"success":  true,
			})
			session := sessions.Default(c)
			session.Set("loginuser", 123)
			session.Save()

		} else {
			c.JSON(200, gin.H{
				"code":    "0",
				"message": "username or password incorrect.",
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
			"message": "sign up success.",
			"success": true,
		})
		// uid, uiderr := uuid.NewUUID()
		// if uiderr != nil {
		// 	return
		// }
		// fmt.Println(uid)
		// session := sessions.Default(c)
		// sess := &model.SessionInfo{
		// 	SessionId: uid.String(),
		// 	Username:  registerInfo.Username,
		// 	UserId:    registerInfo.Id,
		// }
		// err := dao.CreateSession(sess)
		// store := cookie.NewStore([]byte("secret"))
	}
}
