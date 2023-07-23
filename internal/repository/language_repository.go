package repository

import (
	"context"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/database"
	"github.com/v4lomyr/digidict-BE/internal/entity"
)

type (
	LanguageRepository interface {
		GetManyLanguages(ctx context.Context) ([]*entity.Language, error)
	}

	languageRepository struct {
		db *database.GormWrapper
	}
)

func New(db *database.GormWrapper) LanguageRepository {
	return &languageRepository{
		db: db,
	}
}

func (lr languageRepository) GetManyLanguages(ctx context.Context) ([]*entity.Language, error) {
	languages := make([]*entity.Language, 0)

	q := lr.db.Start(ctx)

	if err := q.Find(&languages).Error; err != nil {
		return nil, fmt.Errorf("LanguageRepository.GetManyLanguages: %w", err)
	}

	return languages, nil
}
