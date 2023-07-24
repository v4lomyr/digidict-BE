package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/v4lomyr/digidict-BE/cmd/httpserver"
	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/logger"
	"github.com/v4lomyr/digidict-BE/internal/util/validator"
)

const httpServerMode = "http-server"

var mode string

func init() {
	flag.StringVar(&mode, "mode", httpServerMode, "service run mode")
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	config.InitConfig()

	logger.SetLogrusLogger(config.Get())

	validator.InitValidator()

	var stopFn func()
	switch mode {
	case httpServerMode:
		stopFn = httpserver.Start()
	}

	for {
		<-quit
		stopFn()
		os.Exit(0)
	}

}
