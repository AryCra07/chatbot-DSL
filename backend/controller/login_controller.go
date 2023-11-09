package controller

import (
	"backend/consts"
	"backend/log"
	"backend/model"
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginPost(c *gin.Context) {
	// get request
	var request model.RegisterLoginRequest
	if err := c.BindJSON(&request); err != nil {
		log.Error(consts.Controller, "Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  consts.InvalidRequest,
			"data": nil,
		})
		return
	}

	// get parameters
	username := request.Name
	password := request.Password

	// login
	log.Info(consts.Controller, "Login with name="+username+", "+"password="+password)
	id, ok := service.Login(username, password)

	// send response
	if ok {
		// release token
		token, _ := utils.GenerateToken(id, password)
		log.Info(consts.Controller, "Login success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Login success",
			"data": gin.H{
				"token": token,
			},
		})
	} else {
		log.Error(consts.Controller, "Login fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Login fail",
			"data": gin.H{
				"token": nil,
			},
		})
	}
}
