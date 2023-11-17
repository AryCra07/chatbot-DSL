package controller

import (
	"backend/consts"
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
	name := request.Data.Name
	input := request.Data.Input

	// get response
	response, ok := service.GetMessage(name, input)
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
			"msg":  "Login success",
			"data": gin.H{
				"content": response.Answer,
				"timer":   response.Timer,
			},
		})
	}
}
