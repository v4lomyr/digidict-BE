package middleware

import "github.com/gin-gonic/gin"

type (
	Middleware interface {
		ErrorHandler() gin.HandlerFunc
	}

	middleware struct{}
)

func NewMiddleware() Middleware {
	return &middleware{}
}
