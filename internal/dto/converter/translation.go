package converter

import (
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func EntityToDTOTranslation(in *entity.Translation) *dto.Word {
	out := &dto.Word{
		WordId:          in.TranslationId,
		Word:            in.Word.Word,
		WordClass:       in.WordClass,
		WordDescription: in.WordDescription,
		ExamplePhrase:   in.ExamplePhrase,
	}

	return out
}
