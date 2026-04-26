package errors

import "fmt"

type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s field: %s", e.Field, e.Msg)
}

type DuplicateError struct {
	Resource string
	Field    string
	Value    string
}

func (e *DuplicateError) Error() string {
	return fmt.Sprintf("duplicate %s found for %s with value %s", e.Resource, e.Field, e.Value)
}
