package middleware

import (
	"backend/consts"
	"backend/log"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AUTHMiddleware is a middleware for authentication
/*
	1. Get token from header
	2. Parse token
	3. If token is invalid, return error
	4. If token is valid, set UserId in context
*/
func AUTHMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		url := context.Request.URL.String()

		// Authentication excluded for register and login
		if strings.Contains(url, "login") {
			context.Next()
			return
		}

		tokenParse, claims, err := utils.ParseToken(token)
		if err != nil || !tokenParse.Valid {
			log.Error("Middleware", "Authentication failed with invalid token.")
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": consts.InvalidToken,
				"msg":  "Authentication failed with invalid token!",
				"data": nil,
			})
			context.Abort()
		} else {
			userId := claims.UserId
			log.Info("Middleware", "Authentication successful, ID = "+fmt.Sprintf("%d", userId))
			context.Set("UserId", userId)
			context.Next()

		}
	}
}
