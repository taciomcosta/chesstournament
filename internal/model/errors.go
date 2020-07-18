package model

import (
	"errors"
)

var UnknownError = errors.New("An internal error has occurred")

var UnexistingError = errors.New("Resource was not found")

var InvalidModelError error = errors.New("Model is invalid")
