package repository

import (
	"context"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func (r repository) FindManyEnglishWord(ctx context.Context) ([]*entity.EnglishWord, error) {
	englishWords := make([]*entity.EnglishWord, 0)

	q := r.db.Start(ctx).Table("english_words as ew").
		Select("ew.english_word_id, w.word, w.word_class, w.word_description, w.example_phrase").
		Joins("INNER JOIN words as w ON ew.word_id = w.word_id")

	if err := q.Find(&englishWords).Error; err != nil {
		return nil, fmt.Errorf("LanguageRepository.FindManyEnglishLanguage: %w", err)
	}

	return englishWords, nil
}
