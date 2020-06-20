package chessclub

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s Service

func TestNewService(t *testing.T) {
	service := NewService(&repository.MockChessClub{})
	if service == nil {
		t.Error("it should return a *chessclub.Service")
	}
}

func TestGetClubById(t *testing.T) {
	s = Service{&repository.MockChessClub{}}
	t.Run("should return existing chess club", testGetExistingChessclubById)
	t.Run("should not retrieve unexistent chess club", testGetUnexistentChessclubById)
}

func testGetExistingChessclubById(t *testing.T) {
	club, _ := s.GetClubById(1)
	if club == nil {
		t.Error("it should get chess club by id")
	}
}

func testGetUnexistentChessclubById(t *testing.T) {
	_, err := s.GetClubById(-1)
	if err != UnexistingClubError {
		t.Error("it should return an UnexistingClubError")
	}
}
