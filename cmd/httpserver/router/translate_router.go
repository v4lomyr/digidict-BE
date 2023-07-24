package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/handler"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/middleware"
	"github.com/v4lomyr/digidict-BE/internal/dto"
)

func translateRouter(router *gin.Engine, handler *handler.Handler) {
	router.POST(PublicRoute(TranslatePath), middleware.BindJsonBody[dto.TranslateJSONRequest](), handler.Translate)
}
