package middleware

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AUTHMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		url := context.Request.URL.String()

		// Authentication excluded for register and login
		if strings.Contains(url, "login") || strings.Contains(url, "register") {
			context.Next()
			return
		}

		tokenParse, claims, err := utils.ParseToken(token)
		if err != nil || !tokenParse.Valid {
			/*
				to do
			*/
			context.Abort()
		} else {
			userId := claims.UserId
			println(userId) // fake code
		}
	}
}
