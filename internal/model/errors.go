package model

import (
	"errors"
)

var UnknownError = errors.New("An internal error has occurred")

var UnexistingError = errors.New("Resource was not found")

type InvalidModelError struct {
	Msg string
}

func (err InvalidModelError) Error() string {
	return err.Msg
}
