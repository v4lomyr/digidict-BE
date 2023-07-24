package repository

import (
	"context"

	"github.com/v4lomyr/digidict-BE/internal/entity"
)

type Repository interface {
	FindManyLanguages(ctx context.Context) ([]*entity.Language, error)
}
