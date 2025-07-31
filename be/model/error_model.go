package model

import "fmt"

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	ValidationErrors []string `json:"validationErrors"`
}

type CustomValidationError struct {
	FieldName string `json:"fieldName"`
	Reason    string `json:"reason"`
}

func (e CustomValidationError) Error() string {
	return fmt.Sprintf("invalid input for field %v: %v", e.FieldName, e.Reason)
}
