package main

import (
	"errors"
	"net/http"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Error   string            `json:"error"`
	Details []ValidationError `json:"details"`
}

func (rd *ValidationErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ValidationsErrorResponse(v []ValidationError) *ValidationErrorResponse {
	resp := &ValidationErrorResponse{Error: "Validation failed", Details: v}

	return resp
}

var ErrValidation = errors.New("validation error found")

func appendValidationError(v []ValidationError, field string, message string) []ValidationError {
	v = append(v, ValidationError{
		Field:   field,
		Message: message,
	})
	return v
}
