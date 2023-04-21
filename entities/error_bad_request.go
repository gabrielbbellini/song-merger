package entities

import "fmt"

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		Message: message,
	}
}
