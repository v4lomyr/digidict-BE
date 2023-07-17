package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/database"
	"github.com/v4lomyr/digidict-BE/internal/logger"

	"github.com/gin-gonic/gin"
)

func StartGinHTTPServer(cfg *config.Config) {
	db := database.InitGorm(cfg)

	_ = database.NewGormWrapper(db)

	r := gin.New()
	r.ContextWithFallback = true

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	go func() {
		logger.Log.Info("running server on port :", cfg.HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatal("error while server listen and serve: ", err)
			}
		}
		logger.Log.Info("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(cfg.HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	logger.Log.Info("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("error shutting down server: ", err)
	}

	logger.Log.Info("http server is shutting down gracefully")
}
