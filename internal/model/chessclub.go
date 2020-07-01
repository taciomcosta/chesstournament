package model

import "errors"

type ChessClub struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
}

var InvalidModelError error = errors.New("Model is invalid")
