package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/v4lomyr/digidict-BE/cmd/httpserver/handler"
	"github.com/v4lomyr/digidict-BE/cmd/httpserver/router"
	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/logger"
	"github.com/v4lomyr/digidict-BE/internal/usecase"
)

var srv *http.Server

func Start() (stopFn func()) {
	usecase := usecase.InitDependencies()
	handler := handler.NewHandler(usecase)

	router := router.Init(handler)

	srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Get().HttpServer.Host, config.Get().HttpServer.Port),
		Handler: router,
	}

	go func() {
		logger.Log.Info("running server on port :", config.Get().HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatal("error while server listen and serve: ", err)
			}
		}
		logger.Log.Info("server is not receiving new requests...")
	}()

	return func() {
		GracefulShutdown()
	}
}

func GracefulShutdown() (err error) {
	graceDuration := time.Duration(config.Get().HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	logger.Log.Info("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("error shutting down server: ", err)
	}

	logger.Log.Info("http server is shutting down gracefully")
	return
}
