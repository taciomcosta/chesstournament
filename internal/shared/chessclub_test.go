package shared

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestGetClubId(t *testing.T) {
	var tests = []struct {
		id          int
		expectsClub bool
		expectsErr  bool
	}{
		{1, true, false},
		{-1, false, true},
	}

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
	tests := []struct {
		c           *model.ChessClub
		expectsClub bool
		expectsErr  bool
	}{
		{&data.MockValidChessClub, true, false},
		{&data.MockInvalidChessClub, false, true},
	}

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
		r            model.Filter
		expectsClubs bool
		expectsErr   bool
	}{
		{model.Filter{}, true, false},
		{model.Filter{OrderBy: "invalid"}, false, true},
	}

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

func TestDeleteClub(t *testing.T) {
	tests := []struct {
		clubId      int
		clubExists  bool
		expectsErr  bool
		description string
	}{
		{
			clubId:      1,
			clubExists:  true,
			expectsErr:  false,
			description: "should return delete chess club",
		},
		{
			clubId:      -1,
			clubExists:  false,
			expectsErr:  true,
			description: "should return error for non-existing chessclub",
		},
	}

	for _, tt := range tests {
		c, err := s.DeleteClub(tt.clubId)
		if tt.clubExists && c == nil {
			t.Error(tt.description)
		}

		if tt.expectsErr && err == nil {
			t.Error(tt.description)
		}
	}
}

func TestEditChessclub(t *testing.T) {
	tests := []struct {
		id          int
		c           *model.ChessClub
		expectsErr  bool
		description string
	}{
		{
			id:          1,
			c:           &model.ChessClub{Name: "name", Address: "address"},
			expectsErr:  false,
			description: "should edit chess club without errors",
		},
		{
			id:          -1,
			c:           &model.ChessClub{Name: "name", Address: "address"},
			expectsErr:  true,
			description: "should not edit non-existing chessclub",
		},
		{
			id:          1,
			c:           &model.ChessClub{},
			expectsErr:  true,
			description: "should not edit club with invalid/empty paramters",
		},
	}

	for _, tt := range tests {
		err := s.EditChessclub(tt.id, tt.c)

		if tt.expectsErr && err == nil {
			t.Error(tt.description)
		}
	}
}
