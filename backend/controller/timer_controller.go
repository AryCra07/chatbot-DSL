package controller

import (
	"backend/consts"
	"backend/log"
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	timeString := request.Data.NowTime
	nowTime_, err := strconv.Atoi(timeString)
	if err != nil {
		log.Error(consts.Controller, "time parse fail when timer")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "time parse fail when timer",
			"data": nil,
		})
		return
	}
	nowTime := int32(nowTime_)

	// get user id
	userIdValue := c.Value("UserId")
	if userIdValue == nil {
		log.Error(consts.Controller, "user id does not exist when hello")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id does not exist when hello",
			"data": nil,
		})
		return
	}
	userId, ok := userIdValue.(string)
	if !ok {
		log.Error(consts.Controller, "user id parse fail when hello")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "user id parse fail when hello",
			"data": nil,
		})
		return
	}

	// get response
	response, ok := service.Timer(userId, nowTime)
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
				"exit":    response.IsExit,
				"reset":   response.Reset_,
			},
		})
	}
}
