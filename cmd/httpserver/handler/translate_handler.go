package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/helper/request"
	"github.com/v4lomyr/digidict-BE/internal/helper/response"
)

func (h Handler) Translate(c *gin.Context) {
	jsonBody := request.GetJSONBodyRequest[dto.TranslateJSONRequest](c)

	ctx := c.Request.Context()
	respData, err := h.usecase.Translate(ctx, &jsonBody)
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, respData)
}
