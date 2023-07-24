package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/v4lomyr/digidict-BE/internal/entity"
	"gorm.io/gorm"
)

func (r repository) FindOneTranslation(
	ctx context.Context,
	sourceLanguageId,
	translationLanguageId uint64,
	word string,
) (*entity.Translation, error) {
	translation := &entity.Translation{}

	q := r.db.Start(ctx).Table("translations AS t").
		Select(
			"t.translation_id",
			"tw.word",
			"tw.word_class",
			"tw.word_description",
			"tw.example_phrase",
			"t.created_at",
			"t.updated_at",
			"t.deleted_at").
		Joins("INNER JOIN words tw on t.translation_word_id = tw.word_id").
		Joins("INNER JOIN words sw on t.source_word_id = sw.word_id").
		Where("t.source_language_id = ?", sourceLanguageId).
		Where("t.translation_language_id = ?", translationLanguageId).
		Where("sw.word LIKE ?", word)

	if err := q.First(&translation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("translationRepository.FindOneTranslation: %w", err)
	}

	return translation, nil
}
