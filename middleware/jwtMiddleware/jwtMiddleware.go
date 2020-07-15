package jwtMiddleware

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"
	"net/http"
	"strings"
	"time"

	"github.com/Biubiubiuuuu/yuepuwebsite/helper/jwtHelper"
	"github.com/gin-gonic/gin"
)

// JWT中间件验证
// param query url "token" OR header key "Authorization"
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := "success"
		token := c.Query("token")
		if token == "" {
			authToken := c.GetHeader("Authorization")
			if authToken == "" {
				message = "Query not 'token' param OR header Authorization has not Bearer token"
			}
			token = strings.TrimSpace(authToken)
		}
		claims, err := jwtHelper.ParseToken(token)
		if err != nil {
			user := model.User{
				Token: token,
			}
			if err := user.QueryByToken(); err != nil {
				message = "token 错误"
			}
		} else if time.Now().Unix() > claims.ExpiresAt {
			message = "token 已过期"
		}
		if message != "success" {
			response := entity.ResponseData{
				Message: message,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Next()
	}
}
