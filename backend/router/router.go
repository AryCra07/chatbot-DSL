package router

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServer() {
	// create router engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default();
	router.Use(middleware.CROSMiddleware(), middleware.AUTHMiddleware())

	http.ListenAndServe(":8848", router)
}
