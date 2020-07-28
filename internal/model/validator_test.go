package model

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		want error
		c    Club
	}{
		{errors.New("Invalid fields: Name,Address"), Club{}},
		{nil, Club{Name: "name", Address: "address"}},
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
