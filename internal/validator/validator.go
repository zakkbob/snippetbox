package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(field, error string) {
	if v.FieldErrors == nil {
		v.FieldErrors = map[string]string{}
	}

	if _, exists := v.FieldErrors[field]; !exists {
		v.FieldErrors[field] = error
	}
}

func (v *Validator) CheckField(valid bool, field, error string) {
	if !valid {
		v.AddFieldError(field, error)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func AllowedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
