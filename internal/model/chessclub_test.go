package model

import (
	"testing"
)

func TestValidate(t *testing.T) {
	c := ChessClub{}
	err := c.Validate()
	if err.Error() != "invalid name" {
		t.Error("should have invalid name")
	}

	c.Name = "name"
	err = c.Validate()
	if err.Error() != "invalid address" {
		t.Error("should have invalid address")
	}

	c.Address = "address"
	err = c.Validate()
	if err != nil {
		t.Error("should be valid ChessClub")
	}
}
