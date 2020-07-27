package model

import "testing"

func TestInvalidModelError(t *testing.T) {
	err := InvalidModelError{"some error message"}
	if err.Error() != "some error message" {
		t.Error("InvalidModelError should contain error message")
	}
}
