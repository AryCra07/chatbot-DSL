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

// UserChatHello user register
/*
 * @param c: gin context
 */
func UserChatHello(c *gin.Context) {
	var request model.HelloRequest
	if err := c.BindJSON(&request); err != nil {
		log.Error(consts.Controller, "Hello Request Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  consts.InvalidRequest,
			"data": nil,
		})
		return
	}

	// get parameters
	name := request.Data.Name
	userInfo, flag := dao.GetUserInfo(name)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get hello message fail -- user not exist",
			"data": gin.H{
				"content": nil,
			},
		})
		return
	}
	helloWords := service.GetHello(name, 0, map[string]int32{"balance": userInfo.Balance, "bill": userInfo.Bill})
	if helloWords == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get hello message fail",
			"data": gin.H{
				"content": nil,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Get hello message successfully",
			"data": gin.H{
				"content": helloWords,
			},
		})
	}
}
