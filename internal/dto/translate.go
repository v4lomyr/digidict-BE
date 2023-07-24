package dto

type (
	TranslateResponse struct {
		TranslateResult Word `json:"translation_result"`
	}
)

type (
	TranslateJSONRequest struct {
		SourceLanguageId      uint64 `json:"source_language_id"`
		TranslationLanguageId uint64 `json:"translation_language_id"`
		Word                  string `json:"word"`
	}
)
