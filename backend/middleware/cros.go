package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CROSMiddleware() gin.HandlerFunc {
	const (
		methodList string = "OPTIONS, POST, GET, PUT, PATCH, DELETE"
		headerList string = "Content-Type, Authorization"
	)

	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", methodList)
		context.Writer.Header().Set("Access-Control-Allow-Headers", headerList)
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	}
}
