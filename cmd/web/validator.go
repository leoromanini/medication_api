package main

import (
	"errors"
	"net/http"
	"strings"
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

func extractReadableUnmarshalError(err error) string {
	if readableError := strings.Split(err.Error(), "."); len(readableError) > 1 {
		return readableError[1]
	} else {
		return err.Error()
	}
}

func (m *MedicationsRequest) Bind(r *http.Request) error {
	if m.Medications == nil {
		return errors.New("missing Medications fields")
	}

	if m.Medications.Name == "" {
		m.validationsErrors = appendValidationError(m.validationsErrors, "name", "Name is a required field")
	}

	if m.Medications.Dosage == "" {
		m.validationsErrors = appendValidationError(m.validationsErrors, "dosage", "Dosage is a required field")
	}

	if m.Medications.Form == "" {
		m.validationsErrors = appendValidationError(m.validationsErrors, "form", "Form is a required field")
	}

	if len(m.Medications.Name) > 100 {
		m.validationsErrors = appendValidationError(m.validationsErrors, "name", "Name cannot exceed 100 characters")
	}

	if len(m.Medications.Dosage) > 20 {
		m.validationsErrors = appendValidationError(m.validationsErrors, "dosage", "Name cannot exceed 20 characters")
	}

	if len(m.Medications.Form) > 20 {
		m.validationsErrors = appendValidationError(m.validationsErrors, "form", "Name cannot exceed 20 characters")
	}

	if len(m.validationsErrors) > 0 {
		return ErrValidation
	}

	// TODO: Additional business logic validations would be placed here.

	return nil
}
