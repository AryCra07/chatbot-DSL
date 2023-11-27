package controller

import (
	"backend/consts"
	"backend/log"
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserTimer(c *gin.Context) {
	var request model.TimerRequest
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
	lastTime := request.Data.LastTime
	nowTime := request.Data.NowTime
	//nowTime_, err := strconv.Atoi(timeString)

	// get user id
	userIdValue := c.Value("UserId")
	if userIdValue == nil {
		log.Error(consts.Controller, "user id does not exist when timer")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id does not exist when timer",
			"data": nil,
		})
		return
	}
	userId, ok := userIdValue.(string)
	if !ok {
		log.Error(consts.Controller, "user id parse fail when timer")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id parse fail when timer",
			"data": nil,
		})
		return
	}

	// get response
	response, ok := service.Timer(userId, lastTime, nowTime)
	if response == nil || ok == false {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get Timer fail",
			"data": gin.H{
				"content": nil,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Get Timer success",
			"data": gin.H{
				"content": response.Answer,
				"exit":    response.IsExit,
				"reset":   response.Reset_,
			},
		})
	}
}
