package usecase

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/dto/converter"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"
)

func (u usecase) GetEnglishDictionary(ctx context.Context) (*dto.GetDictionaryResponse, error) {
	englishWords, err := u.repo.FindManyEnglishWord(ctx)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	respData := &dto.GetDictionaryResponse{
		Words: make([]dto.Word, 0),
	}

	for _, word := range englishWords {
		respData.Words = append(respData.Words, *converter.EntityToDTOEnglishWord(word))
	}

	return respData, nil
}
