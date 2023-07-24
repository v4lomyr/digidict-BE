package middleware

import "github.com/gin-gonic/gin"

type (
	Middleware interface {
		RequestId() gin.HandlerFunc
		ErrorHandler() gin.HandlerFunc
		Logger() gin.HandlerFunc
	}

	middleware struct{}
)

func NewMiddleware() Middleware {
	return &middleware{}
}
