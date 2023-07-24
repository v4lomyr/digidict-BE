package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/constants"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
	"github.com/v4lomyr/digidict-BE/internal/util/validator"
)

func bind[requestStruct any](c *gin.Context, contextName string, binder func(interface{}) error) {
	var body requestStruct

	if err := binder(&body); err != nil {
		c.Error(apperror.NewClientError(err))
		c.Abort()
		return
	}

	if err := validator.GetValidator().Validate(body); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Set(contextName, body)
	c.Next()
}

func BindJsonBody[requestBodyStruct any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		bind[requestBodyStruct](c, constants.ContextRequestJSON, c.ShouldBindJSON)
	}
}

func BindPathParam[requestPathParamStruct any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		bind[requestPathParamStruct](c, constants.ContextRequestPathParam, c.ShouldBindUri)
	}
}

func BindQueryParam[requestQueryParamStruct any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		bind[requestQueryParamStruct](c, constants.ContextRequestQueryParam, c.ShouldBindQuery)
	}
}
