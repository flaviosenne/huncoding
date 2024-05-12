package validation

import (
	"encoding/json"
	"errors"

	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("campo com tipo inválido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}
		for _, elem := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: elem.Translate(transl),
				Field:   elem.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValiationError("Alguns campos estão inválidos", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Erro em tentar converter campos")
	}
}
