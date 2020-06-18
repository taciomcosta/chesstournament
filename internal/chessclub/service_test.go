package chessclub

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s service

func TestNew(t *testing.T) {
	var newService Service = New()
	if _, ok := newService.(service); !ok {
		t.Error("it should instantiate a service")
	}
}

func TestGetClubById(t *testing.T) {
	chessclubRepository = &repository.MockChessClub{}
	t.Run("should return existing chess club", testGetExistingChessclubById)
	t.Run("should not retrieve unexistent chess club", testGetUnexistentChessclubById)
}

func testGetExistingChessclubById(t *testing.T) {
	club, _ := s.GetClubById(1)
	if *club != model.MockChessClub {
		t.Error("it should get chess club by id")
	}
}

func testGetUnexistentChessclubById(t *testing.T) {
	_, err := s.GetClubById(-1)
	if _, ok := err.(UnexistingClubErr); !ok {
		t.Error("it should return an UnexistingClubErr")
	}
}
