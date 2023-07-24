package usecase

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/constants"
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/dto/converter"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
)

func (u usecase) Translate(ctx context.Context, payload *dto.TranslateJSONRequest) (*dto.TranslateResponse, error) {
	language, err := u.repo.FindOneLanguageById(ctx, payload.SourceLanguageId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if language == nil {
		return nil, apperror.NewDataNotFoundError(constants.LanugageEntity, payload.SourceLanguageId)
	}

	language, err = u.repo.FindOneLanguageById(ctx, payload.TranslationLanguageId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if language == nil {
		return nil, apperror.NewDataNotFoundError(constants.LanugageEntity, payload.TranslationLanguageId)
	}

	translation, err := u.repo.FindOneTranslation(ctx, payload.SourceLanguageId, payload.TranslationLanguageId, payload.Word)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	if translation == nil {
		return nil, apperror.NewDataNotFoundError(constants.TranslationEntity, payload.Word)
	}

	respData := &dto.TranslateResponse{
		TranslateResult: *converter.EntityToDTOTranslation(translation),
	}

	return respData, nil
}
