package validator

import "github.com/go-playground/validator/v10"

func validateLanguage(fl validator.FieldLevel) bool {
	language := fl.Field().String()

	switch language {
	case "id":
	case "en":
	default:
		return false
	}

	return true
}
