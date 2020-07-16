package shared

import (
	"os"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
)

var s Service

func TestMain(m *testing.M) {
	s = service{
		chessclubRepository: &data.MockChessClubRepository{},
		playerRepository:    &data.MockPlayerRepository{},
	}
	os.Exit(m.Run())
}

func TestNewService(t *testing.T) {
	service := NewService(&data.MockChessClubRepository{}, &data.MockPlayerRepository{})
	if service == nil {
		t.Error("it should return a shared.Service")
	}
}
