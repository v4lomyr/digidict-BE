package main

import (
	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/httpserver"
	"github.com/v4lomyr/digidict-BE/internal/logger"
)

func main() {
	cfg := config.InitConfig()

	logger.SetLogrusLogger(cfg)

	httpserver.StartGinHTTPServer(cfg)
}
