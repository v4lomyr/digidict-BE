package validator

import (
	"reflect"
	"strings"

	"github.com/v4lomyr/digidict-BE/internal/dto"
	apperror "github.com/v4lomyr/digidict-BE/internal/error"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type (
	Validator interface {
		Validate(obj interface{}) error
	}

	customValidator struct {
		validate   *validator.Validate
		translator ut.Translator
	}
)

var vldtr *customValidator

func InitValidator() {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	validate.RegisterValidation("language", validateLanguage)

	en_translations.RegisterDefaultTranslations(validate, trans)
	vldtr = &customValidator{
		validate:   validate,
		translator: trans,
	}
}

func GetValidator() *customValidator {
	return vldtr
}

func (vldtr customValidator) Validate(obj interface{}) error {
	err := vldtr.validate.Struct(obj)
	errsData := make([]dto.ValidationErrorResponse, 0)
	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		for _, e := range validatorErrs {
			translatedErr := e.Translate(vldtr.translator)

			switch e.ActualTag() {
			case "language":
				translatedErr = "currently only en or id is available"
			}

			errData := dto.ValidationErrorResponse{
				Field:   e.Field(),
				Message: translatedErr,
			}
			errsData = append(errsData, errData)
		}
	}

	if len(errsData) > 0 {
		validationErr := apperror.NewValidationError("input validation error", errsData)
		return validationErr
	}

	return nil
}
