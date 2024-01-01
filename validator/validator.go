package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct(payload any) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			// element.Field = err.StructNamespace()
			element.Field = strings.ToLower(strings.Split(err.StructNamespace(), ".")[1])
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}

	return errors
}
