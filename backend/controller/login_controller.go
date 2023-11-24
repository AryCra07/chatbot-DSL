package controller

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/model"
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserLoginPost user register
/*
 * @param c: gin context
 */
func UserLoginPost(c *gin.Context) {

	// get request and bind to model
	var request model.RegisterLoginRequest
	if err := c.BindJSON(&request); err != nil {
		log.Error(consts.Controller, "Body invalid")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": consts.FAIL,
			"msg":  consts.InvalidRequest,
			"data": nil,
		})
		return
	}
	// get parameters
	username := request.Data.Name
	password := request.Data.Password

	// check parameters
	if username == "" || password == "" {
		log.Error(consts.Controller, "Username or password is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": consts.FAIL,
			"msg":  consts.InvalidRequest,
			"data": "wrong",
		})
		return
	}

	// login
	log.Info(consts.Controller, "Login with name="+username+", "+"password="+password)
	id, ok := service.Login(username, password)

	// check and send response
	if ok {
		// release token
		token, _ := utils.GenerateToken(id, password)

		// reset user state
		err := dao.UpdateUserState(id, 0)
		if err != nil {
			log.Error(consts.Controller, "Update user state fail")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Update user state fail",
				"data": gin.H{
					"token": nil,
				},
			})
			return
		}

		// success
		log.Info(consts.Controller, "Login success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Login success",
			"data": gin.H{
				"token": token,
			},
		})
	} else {
		// fail
		log.Error(consts.Controller, "Login fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Login fail, maybe username or password is wrong",
			"data": gin.H{
				"token": nil,
			},
		})
	}
}
