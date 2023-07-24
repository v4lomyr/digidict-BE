package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/handler"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/middleware"
)

func Init(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	middleware := middleware.NewMiddleware()
	router.Use(
		middleware.RequestId(),
		middleware.ErrorHandler(),
		middleware.Logger(),
	)

	languageRouter(router, h)
	dictionaryRouter(router, h)
	translateRouter(router, h)

	return router
}
