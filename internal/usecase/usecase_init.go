package usecase

import "github.com/v4lomyr/digidict-BE/internal/repository"

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func InitDependencies() Usecase {
	repository := repository.InitDependencies()

	return NewUsecase(repository)
}
