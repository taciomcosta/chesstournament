package chessclub

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s service

func TestMain(m *testing.M) {
	chessclubRepository = &repository.MockChessClub{}
}

func TestNew(t *testing.T) {
	var newService Service = New()
	if _, ok := newService.(service); !ok {
		t.Error("it should instantiate a service")
	}
}

func TestGetClubById(t *testing.T) {
	club, _ := s.GetClubById(1)
	if *club != model.MockChessClub {
		t.Error("it should get chess club by id")
	}
}

func TestGetClubByIdUnexistent(t *testing.T) {
	club, err := s.GetClubById(-1)
	if err == nil || club != nil {
		t.Error("it should not retrieve unexistent chess club")
	}
}
