package handler

import "github.com/v4lomyr/digidict-BE/internal/usecase"

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
