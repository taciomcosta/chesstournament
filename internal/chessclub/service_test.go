package chessclub

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var s Service

func TestNewService(t *testing.T) {
	service := NewService(&repository.MockChessClub{})
	if service == nil {
		t.Error("it should return a *chessclub.Service")
	}
}

func TestGetClubId(t *testing.T) {
	var tests = []struct {
		id          int
		expectsClub bool
		expectsErr  bool
	}{
		{1, true, false},
		{-1, false, true},
	}
	s := NewService(&repository.MockChessClub{})

	for _, tt := range tests {
		c, err := s.GetClubById(tt.id)
		if tt.expectsClub && c == nil {
			t.Error("it should return a Chess Club")
		}
		if tt.expectsErr && err == nil {
			t.Error("it should return an error")
		}
	}
}

func TestCreateChessclub(t *testing.T) {
	var tests = []struct {
		c           *model.ChessClub
		expectsClub bool
		expectsErr  bool
	}{
		{&model.ChessClub{Id: 0, Name: "name", Address: "address"}, true, false},
		{&model.ChessClub{}, false, true},
	}

	s := NewService(&repository.MockChessClub{})

	for _, tt := range tests {
		c, err := s.CreateChessclub(tt.c)
		if tt.expectsClub && c == nil {
			t.Error("it should return created chessclub")
		}
		if tt.expectsErr && err == nil {
			t.Error("it should return an error")
		}
	}
}

func TestListClubs(t *testing.T) {
	tests := []struct {
		r            repository.ListFilter
		expectsClubs bool
		expectsErr   bool
	}{
		{repository.ListFilter{}, true, false},
		{repository.ListFilter{"invalid"}, false, true},
	}

	s := NewService(&repository.MockChessClub{})

	for _, tt := range tests {
		cs, err := s.ListClubs(tt.r)

		if tt.expectsClubs && len(cs) == 0 {
			t.Error("it should list Chess Clubs")
		}

		if tt.expectsErr && err == nil {
			t.Error("it should return an error")
		}
	}
}
