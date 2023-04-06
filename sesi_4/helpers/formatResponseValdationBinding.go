package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// https://stackoverflow.com/questions/70069834/return-custom-error-message-from-struct-tag-validation

func FormatValidationError(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{
				Field:   fe.Field(),
				Tag:     fe.Tag(),
				Value:   fe.Param(),
				Message: messageValidation(fe.Tag(), fe.Param()),
			}
			fmt.Println(fe.Value())
		}
		return out
	}
	return nil
}

func messageValidation(tag string, value string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("This field is required%s", value)
	case "gt":
		return fmt.Sprintf("This field is must be greater than %s", value)
	case "email":
		return "This field is must be format email"
	case "unique":
		return "This field is must unique"
	}
	return ""
}
