package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/helper/response"
)

func (h Handler) GetAllLanguage(c *gin.Context) {
	ctx := c.Request.Context()

	respData, err := h.usecase.GetAllLanguage(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, respData)
}
