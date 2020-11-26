package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jotyy/go-crud-example/auth"
	"github.com/jotyy/go-crud-example/model"
	"github.com/jotyy/go-crud-example/response"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)

	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", user.UserName).Count(&count)
	if count > 0 {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "用户名已注册",
		})
		return
	}

	if err := model.DB.Create(&user).Error; err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "注册失败",
		})
	} else {
		c.JSON(200, response.Response{
			Code: 0,
			Msg:  "注册成功",
			Data: user,
		})
	}
}

func Login(c *gin.Context) {
	var user model.User
	username := c.PostForm("user_name")
	password := c.PostForm("password")

	if err := model.DB.Model(model.User{}).Where("user_name = ?", username).First(&user).Error; err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "账号或密码错误",
		})
		return
	} else if err := model.VerifyPassword(user.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "账号或密码错误",
		})
		return
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "token生成失败",
		})
		return
	}
	user.SetToken(token)

	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "登录成功",
		Data: user,
	})
}

func UnRegister(c *gin.Context) {
	var user model.User
	username := c.Param("user_name")

	if err := model.DB.Model(model.User{}).Where("user_name = ?", username).First(&user).Error; err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "账号或密码错误",
		})
		return
	}
	model.DB.Model(model.User{}).Delete(&user)
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "注销成功",
	})
}
