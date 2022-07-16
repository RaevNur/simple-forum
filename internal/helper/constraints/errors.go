package constraints

import "fmt"

type ValidateError struct {
	Field       string
	Description string
}

func (e *ValidateError) Error() string {
	return fmt.Sprintf("validate: wrong %s - %s", e.Field, e.Description)
}

type ExistsError struct {
	Title       string
	Description string
}

func (e *ExistsError) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Description)
}

type ImportError struct {
	Field       string
	Description string
}

func (e *ImportError) Error() string {
	return fmt.Sprintf("cant import %s: %s", e.Field, e.Description)
}
