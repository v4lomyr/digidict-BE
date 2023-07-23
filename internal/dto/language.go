package dto

type (
	Language struct {
		Id   uint64 `json:"language_id"`
		Name string `json:"language_name"`
	}

	GetLanguageResponse struct {
		Languages []Language `json:"languages"`
	}
)
