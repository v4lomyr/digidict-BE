package converter

import (
	"github.com/v4lomyr/digidict-BE/internal/dto"
	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func EntityToDTOIndonesianWord(in *entity.IndonesianWord) *dto.Word {
	out := &dto.Word{
		WordId:          in.IndonesianWordId,
		Word:            in.Word.Word,
		WordClass:       in.WordClass,
		WordDescription: in.WordDescription,
		ExamplePhrase:   in.ExamplePhrase,
	}

	return out
}

func EntityToDTOEnglishWord(in *entity.EnglishWord) *dto.Word {
	out := &dto.Word{
		WordId:          in.EnglishWordId,
		Word:            in.Word.Word,
		WordClass:       in.WordClass,
		WordDescription: in.WordDescription,
		ExamplePhrase:   in.ExamplePhrase,
	}

	return out
}
