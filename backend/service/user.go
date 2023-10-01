package service

import (
	"backend/dao"
	"backend/model"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

// Login 用户登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// TODO 查询数据库
	user, flag, err := dao.Authentication(username, utils.GetMd5(password))
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "查询数据库错误",
		})
		return
	}
	if !flag {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名或密码错误",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "登录成功",
			"data": gin.H{
				"token": utils.GenerateToken(user.Id, user.Password),
			},
		})
		return
	}
}

// Register 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名或密码不能为空",
		})
		return
	}

	flag, _ := dao.IsUsernameExist(username)
	if flag {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名已存在",
		})
		return
	}

	user := model.User{
		Id:       utils.GetUUID(),
		Name:     username,
		Password: utils.GetMd5(password),
	}

	err := dao.InsertUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "插入数据库错误",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "注册成功",
			"data": gin.H{
				"token": utils.GenerateToken(user.Id, user.Password),
			},
		})
		return
	}
}
