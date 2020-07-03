package model

type InternalErr struct{}

func (e InternalErr) Error() string {
	return "An internal error has occurred"
}
