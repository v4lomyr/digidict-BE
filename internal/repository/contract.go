package repository

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

type Repository interface {
	FindManyLanguages(ctx context.Context) ([]*entity.Language, error)

	FindManyIndonesianWord(ctx context.Context) ([]*entity.IndonesianWord, error)
	FindManyEnglishWord(ctx context.Context) ([]*entity.EnglishWord, error)
}
