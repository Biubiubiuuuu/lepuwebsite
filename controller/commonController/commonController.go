package commonController

import (
	"net/http"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"github.com/Biubiubiuuuu/yuepuwebsite/service/commonService"

	"github.com/gin-gonic/gin"
)

// @Summary 获取验证码
// @tags 公共接口
// @Accept  application/json
// @Produce  json
// @Param tel query string true "手机号码"
// @Success 200 {object} entity.ResponseData "desc"
// @Router /api/v1/common/verificationcode [GET]
func VerificationCode(c *gin.Context) {
	tel := c.Query("tel")
	res := commonService.VerificationCode(tel)
	c.JSON(http.StatusOK, res)
}

func GetToken(c *gin.Context) (token string, res entity.ResponseData) {
	token = c.Query("token")
	if token == "" {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			res.Message = "Query not 'token' param OR header Authorization has not Bearer token"
			return
		}
		token = strings.TrimSpace(authToken)
	}
	res.Status = true
	return
}
