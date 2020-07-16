package adminMiddleware

import (
	"net/http"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"

	"github.com/gin-gonic/gin"
)

// 验证是否为后台账号
// param query url "token" OR header key "Authorization"
func UserTypeIsAdmin() gin.HandlerFunc {
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
		user := model.User{
			Token: token,
		}
		if err := user.QueryByToken(); err != nil {
			message = "token 错误"
		}
		if user.Type != "1" {
			message = "没有权限访问请求资源"
		}
		if message != "success" {
			response := entity.ResponseData{
				Message: message,
			}
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		c.Next()
	}
}
