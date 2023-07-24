package usecase

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/dto/converter"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
)

func (u usecase) GetIndonesianDictionary(ctx context.Context) (*dto.GetDictionaryResponse, error) {
	indonesianWords, err := u.repo.FindManyIndonesianWord(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	respData := &dto.GetDictionaryResponse{
		Words: make([]dto.Word, 0),
	}

	for _, word := range indonesianWords {
		respData.Words = append(respData.Words, *converter.EntityToDTOIndonesianWord(word))
	}

	return respData, nil
}
