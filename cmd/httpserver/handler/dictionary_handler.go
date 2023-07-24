package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/helper/request"
	"github.com/v4lomyr/digidict-BE/internal/helper/response"
)

func (h Handler) GetAllIndonesianWords(c *gin.Context) {
	pathParam := request.GetPathParamRequest[dto.GetDictionaryPathParamRequest](c)

	ctx := c.Request.Context()

	var respData *dto.GetDictionaryResponse = nil
	var err error

	switch pathParam.Language {
	case "id":
		respData, err = h.usecase.GetIndonesianDictionary(ctx)
	default:
		respData, err = h.usecase.GetEnglishDictionary(ctx)
	}
	if err != nil {
		c.Error(err)
		return
	}

	response.ResponseOKData(c, respData)
}
