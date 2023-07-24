package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/dto"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
)

const (
	MessageInternalServerError        = "currently our server is facing unexpected error, please try again later"
	MessageValidationError            = "input validation error"
	MessageInvalidJsonValueTypeError  = "invalid value for %s"
	MessageInvalidJsonUnmarshallError = "invalid JSON format"
	MessageJsonSyntaxError            = "invalid JSON syntax"
)

func (m *middleware) ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 1 {
			//usually not a good idea, but since Gin already throws panic if we give nil to c.Error, it is probably safe
			var err error
			err = c.Errors[0]

			var serverError *apperror.ServerError
			var clientError *apperror.ClientError
			var jutErr *json.UnmarshalTypeError
			var sErr *json.SyntaxError
			var vldtnErr *apperror.ValidationError

			isClientError := false
			if errors.As(err, &clientError) {
				isClientError = true
				err = clientError.Unwrap()
			}

			switch {
			case errors.As(err, &sErr):
				handleJsonSyntaxError(c, sErr)
				return
			case errors.As(err, &jutErr):
				handleJsonUnmarshalTypeError(c, jutErr)
				return
			case isClientError:
				c.AbortWithStatusJSON(clientError.GetCode(), dto.ErrorResponse{
					Message: clientError.Error(),
					Code:    clientError.GetCode(),
				})
				return
			case errors.As(err, &serverError):
				c.AbortWithStatusJSON(serverError.GetCode(), dto.ErrorResponse{
					Message: MessageInternalServerError,
					Code:    serverError.GetCode(),
				})
				return
			case errors.As(err, &vldtnErr):
				c.AbortWithStatusJSON(vldtnErr.GetCode(), dto.ErrorResponse{
					Message: vldtnErr.Error(),
					Data:    vldtnErr.GetDetails(),
					Code:    vldtnErr.GetCode(),
				})
				return

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
					Message: MessageInternalServerError,
					Code:    http.StatusInternalServerError,
				})
				return
			}
		} else if len(c.Errors) > 1 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Message: MessageInternalServerError,
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}
}

func handleJsonSyntaxError(c *gin.Context, err *json.SyntaxError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: MessageJsonSyntaxError,
		Code:    http.StatusBadRequest,
	})
}

func handleJsonUnmarshalTypeError(c *gin.Context, err *json.UnmarshalTypeError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: fmt.Sprintf(MessageInvalidJsonValueTypeError, err.Field),
		Code:    http.StatusBadRequest,
	})
}
