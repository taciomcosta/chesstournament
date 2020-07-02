package validator

import (
	"errors"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		want error
		c    model.ChessClub
	}{
		{errors.New("Invalid fields: Name,Address"), model.ChessClub{}},
		{nil, model.ChessClub{Name: "name", Address: "address"}},
	}

	for _, tt := range tests {
		got := Validate(tt.c)
		if tt.want == nil && got != nil {
			t.Errorf("want %s, got %s", tt.want, got)
		}
		if tt.want != nil && got == nil {
			t.Errorf("want %s, got %s", tt.want, got)
		}
	}
}
