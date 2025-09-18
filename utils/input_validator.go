package utils

import (
	"reflect"
	"strings"

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
