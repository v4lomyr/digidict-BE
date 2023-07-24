package request

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/constants"
)

func GetJSONBodyRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestJSON).(bodyRequestType)
}

func GetPathParamRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestPathParam).(bodyRequestType)
}

func GetQueryParamRequest[bodyRequestType any](c *gin.Context) bodyRequestType {
	return c.MustGet(constants.ContextRequestQueryParam).(bodyRequestType)
}
