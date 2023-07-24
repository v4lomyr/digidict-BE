package repository

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

type Repository interface {
	FindManyLanguages(ctx context.Context) ([]*entity.Language, error)
	FindOneLanguageById(ctx context.Context, languageId uint64) (*entity.Language, error)

	FindManyIndonesianWord(ctx context.Context) ([]*entity.IndonesianWord, error)
	FindManyEnglishWord(ctx context.Context) ([]*entity.EnglishWord, error)

	FindOneTranslation(ctx context.Context, sourceLanguageId, translationLanguageId uint64, word string) (*entity.Translation, error)
}
