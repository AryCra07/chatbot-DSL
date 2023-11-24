package controller

import (
	"backend/consts"
	"backend/log"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserChatHello user register
/*
 * @param c: gin context
 */
func UserChatHello(c *gin.Context) {
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

	helloWords := service.Hello(userId)
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
