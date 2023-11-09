package router

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServer() {
	// create router engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(middleware.CROSMiddleware(), middleware.AUTHMiddleware())

	// user
	user := router.Group("/user")
	// login page
	user.POST("/sign", controller.UserLoginPost)

	// index page
	user.POST("/hello", controller.UserChatHello)
	user.POST("/message", controller.UserChatMessage)

	// port
	err := http.ListenAndServe(":8848", router)
	if err != nil {
		return
	}
}
