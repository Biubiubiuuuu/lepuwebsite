package jwtMiddleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/helper/jwtHelper"
	"github.com/Biubiubiuuuu/yuepuwebsite/model"

	"github.com/gin-gonic/gin"
)

// JWT中间件验证
// param query url "token" OR header key "Authorization"
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := entity.ResponseData{}
		token := c.Query("token")
		if token == "" {
			authToken := c.GetHeader("Authorization")
			if authToken == "" {
				res.Message = "Query not 'token' param OR header Authorization has not Bearer token"
			}
			token = strings.TrimSpace(authToken)
		}
		claims, err := jwtHelper.ParseToken(token)
		if err != nil {
			res.Message = err.Error()
		} else if time.Now().Unix() > claims.ExpiresAt {
			res.Message = "token 已过期"
		} else {
			res.Status = true
		}
		user := model.User{
			Token: token,
		}
		if token == "" {
			res.Message = "token不能为空"
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		if err := user.QueryByToken(); err != nil {
			res.Message = "token错误，请重新登录获取授权"
			res.Status = false
		}
		if user.IsEnable {
			res.Message = "用户已禁用,无权访问任何信息"
			c.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}
		if !res.Status {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		c.Next()
	}
}
