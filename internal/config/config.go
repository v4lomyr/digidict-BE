package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App        AppConfig
		HttpServer HttpServerConfig
		Database   DatabaseConfig
	}

	AppConfig struct {
		Environment string
	}

	HttpServerConfig struct {
		Host        string
		Port        int
		GracePeriod int
	}

	DatabaseConfig struct {
		Host                  string
		Port                  int
		DBName                string
		Username              string
		Password              string
		SSLMode               string
		MaxIdleConn           int
		MaxOpenConn           int
		MaxConnLifetimeMinute int
	}
)

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		App:        initAppConfig(),
		HttpServer: initHttpServerConfig(),
		Database:   initDbConfig(),
	}
}

func initAppConfig() AppConfig {
	environment := os.Getenv("APP_ENVIRONMENT")

	return AppConfig{
		Environment: environment,
	}
}

func initHttpServerConfig() HttpServerConfig {
	host := os.Getenv("HTTP_SERVER_HOST")

	port, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_PORT")
	}

	gracePeriod, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_GRACE_PERIOD"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	return HttpServerConfig{
		Host:        host,
		Port:        int(port),
		GracePeriod: int(gracePeriod),
	}
}

func initDbConfig() DatabaseConfig {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_PORT")
	}

	maxIdleConn, err := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_IDLE_CONN")
	}

	maxOpenConn, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_OPEN_CONN")
	}

	connMaxLifetime, err := strconv.ParseInt(os.Getenv("DB_CONN_MAX_LIFETIME"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_CONN_MAX_LIFETIME")
	}

	return DatabaseConfig{
		Port:                  int(port),
		Host:                  host,
		DBName:                name,
		Username:              user,
		Password:              password,
		SSLMode:               sslMode,
		MaxIdleConn:           int(maxIdleConn),
		MaxOpenConn:           int(maxOpenConn),
		MaxConnLifetimeMinute: int(connMaxLifetime),
	}
}
