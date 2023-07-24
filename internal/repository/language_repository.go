package repository

import (
	"context"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func (r repository) FindManyLanguages(ctx context.Context) ([]*entity.Language, error) {
	languages := make([]*entity.Language, 0)

	q := r.db.Start(ctx)

	if err := q.Find(&languages).Error; err != nil {
		return nil, fmt.Errorf("LanguageRepository.GetManyLanguages: %w", err)
	}

	return languages, nil
}
