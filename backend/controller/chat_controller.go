package controller

import (
	"backend/consts"
	"backend/log"
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserChatMessage is a controller for user register
/*
 * @param c: gin context
 */
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
	input := request.Data.Input
	userIdValue := c.Value("UserId")
	if userIdValue == nil {
		log.Error(consts.Controller, "user id does not exist when chat")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id does not exist when chat",
			"data": nil,
		})
		return
	}
	userId, ok := userIdValue.(string)
	if !ok {
		log.Error(consts.Controller, "user id parse fail when chat")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id parse fail when chat",
			"data": nil,
		})
		return
	}

	// get response
	response, ok := service.ChatResponse(userId, input)
	if response == nil || ok == false {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get Answer fail",
			"data": gin.H{
				"content": nil,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Get Answer success",
			"data": gin.H{
				"content": response.Answer,
			},
		})
	}
}
