package response

import (
	"net/http"

	"github.com/v4lomyr/digidict-BE/internal/dto"

	"github.com/gin-gonic/gin"
)

func ResponseOKPlain(c *gin.Context) {
	ResponseOKData(c, nil)
}

func ResponseOKData(c *gin.Context, data interface{}) {
	ResponseOK(c, "success", data)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	ResponseSuccessJSON(c, http.StatusOK, message, data)
}

func ResponseSuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse{
		Message: message,
		Data:    data,
		Code:    statusCode,
	})
}
