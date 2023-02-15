package entities

import "fmt"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message,
	}
}
