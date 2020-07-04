package model

import (
	"errors"
)

var UnknownError = errors.New("An internal error has occurred")

var UnexistingClubError = errors.New("Chess Club was not found")

var InvalidModelError error = errors.New("Model is invalid")
