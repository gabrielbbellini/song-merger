package entities

import "fmt"

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		Message: message,
	}
}
