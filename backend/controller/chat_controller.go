package controller

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserChatMessage(c *gin.Context) {
	var request model.MessageRequest
	if err := c.BindJSON(&request); err != nil {
		log.Error(consts.Controller, "Message Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  consts.InvalidRequest,
			"data": nil,
		})
		return
	}

	// get parameters
	name := request.Name
	input := request.Input
	user, flag := dao.GetUserInfo(name)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get message fail -- user not exist",
			"data": gin.H{
				"content": nil,
			},
		})
		return
	}
	answer := service.GetMessage(name, user.State, input)
	if answer == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Login fail",
			"data": gin.H{
				"content": nil,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Login success",
			"data": gin.H{
				"content": answer,
			},
		})
	}
}
