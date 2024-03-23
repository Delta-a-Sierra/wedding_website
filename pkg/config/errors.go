package config

import "fmt"

type ValidationError struct {
	reason string
}

func NewValidationError(reason string) ValidationError {
	return ValidationError{reason: reason}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("failed validation checks, reason: %s", v.Error())
}
