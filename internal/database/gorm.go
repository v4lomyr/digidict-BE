package database

import (
	"context"
	"fmt"
	"time"

	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/constants"
	"github.com/v4lomyr/digidict-BE/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitGorm(cfg *config.Config) *gorm.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DBName,
		dbCfg.Port,
		dbCfg.SSLMode,
	)

	var logCfg gormlogger.Config

	if cfg.App.Environment == constants.AppEnvironmentProduction {
		logCfg = gormlogger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
			LogLevel:                  gormlogger.Warn,
		}
	} else {
		logCfg = gormlogger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
			LogLevel:                  gormlogger.Info,
		}
	}

	newLogger := gormlogger.New(
		logger.Log,
		logCfg,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Log.Fatal("error initializing database: ", err.Error())
	}

	dbInstance, err := db.DB()
	if err != nil {
		logger.Log.Fatal("error getting generic database instance: ", err.Error())
	}

	dbInstance.SetMaxIdleConns(dbCfg.MaxIdleConn)
	dbInstance.SetMaxOpenConns(dbCfg.MaxOpenConn)
	dbInstance.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db
}

type GormWrapper struct {
	db *gorm.DB
}

func NewGormWrapper(db *gorm.DB) *GormWrapper {
	return &GormWrapper{
		db: db,
	}
}

func (w *GormWrapper) Start(ctx context.Context) *gorm.DB {
	return w.db.WithContext(ctx)
}
