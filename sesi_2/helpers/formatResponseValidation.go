package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Tag   string
	Value string
}

// https://stackoverflow.com/questions/70069834/return-custom-error-message-from-struct-tag-validation

func FormatValidationError(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{Field: fe.Field(), Tag: fe.Tag(), Value: fe.Param()}
			fmt.Println(fe)
		}
		return out
	}
	return nil
}
