package model

import "errors"

type ChessClub struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var InvalidModelError error = errors.New("Model is invalid")

func (c *ChessClub) Validate() error {
	if c.Name == "" {
		return errors.New("invalid name")
	}
	if c.Address == "" {
		return errors.New("invalid address")
	}
	return nil
}
