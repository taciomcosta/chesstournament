package validator

import (
	"errors"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestValidate(t *testing.T) {
	want := errors.New("Invalid fields: Name,Address")
	got := Validate(model.ChessClub{})
	if want.Error() != got.Error() {
		t.Errorf("want %s, got %s", want, got)
	}
}
