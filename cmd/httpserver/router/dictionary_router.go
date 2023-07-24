package router

import (
	"github.com/gin-gonic/gin"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/handler"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/middleware"
	"github.com/v4lomyr/digidict-BE/internal/dto"
)

func dictionaryRouter(router *gin.Engine, handler *handler.Handler) {
	router.GET(PublicRoute(DictionaryPath), middleware.BindPathParam[dto.GetDictionaryPathParamRequest](), handler.GetAllIndonesianWords)
}
