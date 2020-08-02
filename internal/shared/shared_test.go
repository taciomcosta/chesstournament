package shared

import (
	"os"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
)

var s Service

func TestMain(m *testing.M) {
	s = NewService(data.MockRepository())
	os.Exit(m.Run())
}

func TestNewService(t *testing.T) {
	service := NewService(data.MockRepository())
	if service == nil {
		t.Error("it should return a shared.Service")
	}
}
