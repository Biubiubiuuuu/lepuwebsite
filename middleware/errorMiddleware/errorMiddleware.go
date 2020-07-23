package errorMiddleware

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 404
func NotFound(c *gin.Context) {
	response := entity.ResponseData{
		Message: "404 Not Found",
	}
	c.JSON(http.StatusNotFound, response)
}
