package shared

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestGetClubById(t *testing.T) {
	f := func(id int) (interface{}, error) { return s.GetClubById(id) }
	testGetById(f, t)
}

func TestCreateChessclub2(t *testing.T) {
	club, err := s.CreateChessclub(&data.MockValidChessClub)
	thenAssertValueIs(t, *club, data.MockValidChessClub)
	thenAssertErrorIsNil(t, err)
}

func TestCreateInvalidChessclub(t *testing.T) {
	club, err := s.CreateChessclub(&data.MockInvalidChessClub)
	thenAssertValueIsNil(t, club)
	thenAssertErrorIs(t, err, model.InvalidModelError{Msg: "Invalid fields: Name,Address"})
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
	club, err := s.DeleteClub(1)
	thenAssertValueIs(t, *club, data.MockValidChessClub)
	thenAssertErrorIsNil(t, err)
}

func TestDeleteUnexistentClub(t *testing.T) {
	club, err := s.DeleteClub(-1)
	thenAssertValueIsNil(t, club)
	thenAssertErrorIs(t, err, model.UnexistingError)
}

func TestEditChessclub(t *testing.T) {
	tests := []struct {
		id          int
		c           *model.Club
		expectsErr  bool
		description string
	}{
		{
			id:          1,
			c:           &model.Club{Name: "name", Address: "address"},
			expectsErr:  false,
			description: "should edit chess club without errors",
		},
		{
			id:          -1,
			c:           &model.Club{Name: "name", Address: "address"},
			expectsErr:  true,
			description: "should not edit non-existing chessclub",
		},
		{
			id:          1,
			c:           &model.Club{},
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
