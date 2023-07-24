package middleware

import (
	"github.com/v4lomyr/digidict-BE/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *middleware) RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New().String()
		c.Set(constants.ContextKeyRequestId, uuid)
		c.Next()
	}
}
