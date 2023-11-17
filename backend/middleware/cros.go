package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CROSMiddleware is a middleware for Cross-Origin Resource Sharing
/*
	1. Set Access-Control-Allow-Origin to *
	2. Set Access-Control-Allow-Methods to OPTIONS, POST, GET, PUT, PATCH, DELETE
	3. Set Access-Control-Allow-Headers to Content-Type, Authorization
	4. If request method is OPTIONS, return status 204
	5. If request method is not OPTIONS, next
*/
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
