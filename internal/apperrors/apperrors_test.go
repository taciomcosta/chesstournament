package apperrors

import "testing"

func TestInternalErr(t *testing.T) {
	want := "An internal error has occurred"
	got := InternalErr{}.Error()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
