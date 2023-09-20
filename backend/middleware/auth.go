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
			log.Error("Middleware", "Authentication failed with invalid token.")
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": consts.FAIL,
				"msg":  "Authentication failed with invalid token.",
				"data": nil,
			})
			context.Abort()
		} else {
			userId := claims.UserId
			auth := claims.Auth

			if userId == consts.NotExistId && auth < 1 {
				log.Error("Middleware", "Authentication fail with invalid uid")
				context.JSON(http.StatusUnauthorized, gin.H{
					"code": consts.SUCCESS,
					"msg":  "Authentication failed with nonexistent user",
					"data": nil,
				})
				context.Abort()
			} else {
				log.Info("Middleware", "Authentication successful, ID = "+fmt.Sprintf("%d", userId))
				context.Set("UserId", userId)
				context.Next()
			}
		}
	}
}
