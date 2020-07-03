package model

import (
	"errors"
	"strings"

	v10Validator "github.com/go-playground/validator/v10"
)

func Validate(model interface{}) error {
	validate := v10Validator.New()
	if errs := validate.Struct(model); errs != nil {
		return formattedErrors(errs)
	}
	return nil
}

func formattedErrors(errs error) error {
	fs := []string{}
	for _, err := range errs.(v10Validator.ValidationErrors) {
		fs = append(fs, err.Field())
	}
	return errors.New("Invalid fields: " + strings.Join(fs, ","))
}
