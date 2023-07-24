package usecase

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/dto"
)

type Usecase interface {
	GetAllLanguage(ctx context.Context) (*dto.GetLanguageResponse, error)

	GetIndonesianDictionary(ctx context.Context) (*dto.GetDictionaryResponse, error)
	GetEnglishDictionary(ctx context.Context) (*dto.GetDictionaryResponse, error)
}
