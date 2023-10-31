package service

import (
	"ginchat/models"
	"ginchat/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	user := models.UserBasic{}
	name := c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	phone := c.PostForm("phone")

	hasPhone := models.SearchPhone(phone)
	if hasPhone {
		c.JSON(400, gin.H{
			"message": "手机号已注册",
			"code":    0,
		})
		return
	}
	if password != repassword {
		c.JSON(400, gin.H{
			"message": "两次密码不一致",
			"code":    0,
		})
		return
	}

	user.Name = name
	user.Password = utils.EnCodeMD5(password)
	user.Phone = phone

	err := models.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    0,
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "注册",
		"code":    1,
	})
}

func UserLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	if phone == "" || password == "" {
		c.JSON(400, gin.H{
			"message": "参数异常",
			"code":    0,
		})
		return
	}

	curUser := models.SearchUserByPhone(phone)
	if curUser.Id == 0 {
		c.JSON(400, gin.H{
			"message": "该用户不存在",
			"code":    0,
		})
		return
	}

	flag := utils.UnCodeMD5(password, curUser.Password)
	if !flag {
		c.JSON(400, gin.H{
			"message": "密码错误",
			"code":    0,
			"flag":    flag,
		})
		return
	}

	token := utils.CreateToken(map[string]interface{}{
		"phone": phone,
	})

	c.SetCookie("token", token, int(time.Hour*24*7), "/", "localhost", false, true)
	storeUserDataInRedis(token, phone, time.Hour*24*7)
	c.JSON(200, gin.H{
		"message": "登录成功",
		"code":    1,
		"token":   token,
	})
}

func GetUserInfo(c *gin.Context) {
	token, _ := c.Get("token")
	phone := utils.GetPhoneFromToken(token.(string))
	user := models.SearchUserByPhone(phone)
	result := map[string]interface{}{
		"name":   user.Name,
		"phone":  user.Phone,
		"avatar": user.Avatar,
	}
	c.JSON(200, gin.H{
		"code":    1,
		"data":    result,
		"message": "获取用户信息成功",
	})
}

func storeUserDataInRedis(key string, value string, expire time.Duration) {
	utils.RDB.Set(utils.RDB.Context(), key, value, expire)
}
