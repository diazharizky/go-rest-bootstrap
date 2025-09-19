package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/go-playground/validator/v10"
)

func NewInputValidator() *validator.Validate {
	var inputValidator = validator.New()

	inputValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return inputValidator
}

var inputErrorTagMessageMap = map[string]string{
	"required": "Field is required",
	"email":    "Not a valid email",
	"min":      "Value must contain at least %s characters",
	"max":      "Value must be at most %s characters long",
}

type XValidationErrors validator.ValidationErrors

func (errs XValidationErrors) ResolveInputErrors() (inputErrs []apiresp.InputError) {
	inputErrs = make([]apiresp.InputError, len(errs))
	for i, err := range errs {
		errDescription, exists := inputErrorTagMessageMap[err.Tag()]
		if !exists {
			errDescription = err.Error()
		} else if err.Param() != "" {
			errDescription = fmt.Sprintf(errDescription, err.Param())
		}

		inputErrs[i] = apiresp.InputError{
			Field:       err.Field(),
			Description: errDescription,
		}
	}
	return
}
