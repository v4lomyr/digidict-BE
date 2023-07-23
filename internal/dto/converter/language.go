package converter

import (
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func EntityToDTOLanguage(in *entity.Language) *dto.Language {
	out := &dto.Language{
		Id:   in.LanguageId,
		Name: in.Language,
	}

	return out
}
