package repository

import (
	"context"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

func (r repository) FindManyIndonesianWord(ctx context.Context) ([]*entity.IndonesianWord, error) {
	indonesianWords := make([]*entity.IndonesianWord, 0)

	q := r.db.Start(ctx).Table("indonesian_words as iw").
		Select("iw.indonesian_word_id, w.word, w.word_class, w.word_description, w.example_phrase").
		Joins("INNER JOIN words as w ON iw.word_id = w.word_id")

	if err := q.Find(&indonesianWords).Error; err != nil {
		return nil, fmt.Errorf("LanguageRepository.FindManyIndonesianWord: %w", err)
	}

	return indonesianWords, nil
}
