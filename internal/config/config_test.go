package config

import (
	"testing"
)

func TestString(t *testing.T) {
	env = map[string]string{
		"key": "1",
	}
	got := String("key")
	if got != "1" {
		t.Errorf("it should return value as string, given a key")
	}
}
