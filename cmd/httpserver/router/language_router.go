package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/handler"
)

func languageRouter(router *gin.Engine, handler *handler.Handler) {
	router.GET(PublicRoute(LanguagePath), handler.GetAllLanguage)
}
