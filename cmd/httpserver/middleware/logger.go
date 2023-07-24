package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/internal/constants"
	"github.com/v4lomyr/digidict-BE/internal/logger"
)

func (m *middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()
		reqId, exists := c.Get(constants.ContextKeyRequestId)
		if !exists {
			reqId = ""
		}

		param := map[string]interface{}{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
			"request_id":  reqId,
		}

		if len(c.Errors) == 0 {
			logger.Log.WithFields(param).Info("incoming request")
		} else {
			errList := []error{}
			for _, err := range c.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				param["errors"] = errList
				logger.Log.WithFields(param).Error("got error")
			}
		}
	}
}
