package repository

import (
	"github.com/v4lomyr/digidict-BE/internal/config"
	"github.com/v4lomyr/digidict-BE/internal/database"
)

type repository struct {
	db *database.GormWrapper
}

func NewRepository(db *database.GormWrapper) Repository {
	return &repository{
		db: db,
	}
}

func InitDependencies() Repository {
	cfg := config.Get()

	db := database.InitGorm(cfg)

	gormWrapper := database.NewGormWrapper(db)

	return NewRepository(gormWrapper)
}
