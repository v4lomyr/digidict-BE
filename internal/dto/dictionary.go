package dto

type (
	Word struct {
		WordId          uint64 `json:"word_id"`
		Word            string `json:"word"`
		WordClass       string `json:"word_class"`
		WordDescription string `json:"word_description"`
		ExamplePhrase   string `json:"example_phrase"`
	}

	GetDictionaryResponse struct {
		Words []Word `json:"words"`
	}
)

type (
	GetDictionaryPathParamRequest struct {
		Language string `uri:"language" validate:"language"`
	}
)
