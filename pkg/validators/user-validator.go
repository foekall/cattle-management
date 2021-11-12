package validators

import (
	"github.com/go-playground/validator/v10"
)

func ValidateEmail(field validator.FieldLevel) bool {
	return false //strings.Contains((field.Field().String(), "abc")
}
