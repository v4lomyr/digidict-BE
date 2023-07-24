package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/entity"
	"gorm.io/gorm"
)

func (r repository) FindManyLanguages(ctx context.Context) ([]*entity.Language, error) {
	languages := make([]*entity.Language, 0)

	q := r.db.Start(ctx)

	if err := q.Find(&languages).Error; err != nil {
		return nil, fmt.Errorf("LanguageRepository.GetManyLanguages: %w", err)
	}

	return languages, nil
}

func (r repository) FindOneLanguageById(ctx context.Context, languageId uint64) (*entity.Language, error) {
	language := &entity.Language{}

	q := r.db.Start(ctx).Where("language_id = ?", languageId)

	if err := q.First(&language).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("LanguageRepository.FindOneLanguageById: %w", err)
	}

	return language, nil
}
