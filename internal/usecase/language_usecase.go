package usecase

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/dto/converter"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
)

func (u usecase) GetAllLanguage(ctx context.Context) (*dto.GetLanguageResponse, error) {
	languages, err := u.repo.FindManyLanguages(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	respData := &dto.GetLanguageResponse{
		Languages: make([]dto.Language, 0),
	}

	for _, language := range languages {
		respData.Languages = append(respData.Languages, *converter.EntityToDTOLanguage(language))
	}

	return respData, nil
}
