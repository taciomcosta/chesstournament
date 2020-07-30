package shared

import (
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func TestGetClubById(t *testing.T) {
	club, err := s.GetClubById(1)
	thenAssertValueIsNotNil(t, club)
	thenAssertErrorIsNil(t, err)
}

func TestGetClubUnexistent(t *testing.T) {
	club, err := s.GetClubById(-1)
	thenAssertValueIsNil(t, club)
	thenAssertValueIsNotNil(t, err)
}

func TestCreateClub(t *testing.T) {
	club, err := s.CreateClub(&data.MockValidClub)
	thenAssertValueIs(t, *club, data.MockValidClub)
	thenAssertErrorIsNil(t, err)
}

func TestCreateInvalidClub(t *testing.T) {
	club, err := s.CreateClub(&data.MockInvalidClub)
	thenAssertValueIsNil(t, club)
	thenAssertErrorIs(t, err, model.InvalidModelError{Msg: "Invalid fields: Name,Address"})
}

func TestListClubsValidFilters(t *testing.T) {
	clubs, err := s.ListClubs(model.Filter{})
	thenAssertSliceLenIs(t, clubs, 1)
	thenAssertErrorIsNil(t, err)
}

func TestClubsInvalidFilters(t *testing.T) {
	clubs, err := s.ListClubs(model.Filter{OrderBy: "invalid"})
	thenAssertSliceLenIs(t, clubs, 0)
	thenAssertValueIsNotNil(t, err)
}

func TestDeleteClub(t *testing.T) {
	club, err := s.DeleteClub(1)
	thenAssertValueIs(t, *club, data.MockValidClub)
	thenAssertErrorIsNil(t, err)
}

func TestDeleteUnexistentClub(t *testing.T) {
	club, err := s.DeleteClub(-1)
	thenAssertValueIsNil(t, club)
	thenAssertErrorIs(t, err, model.UnexistingError)
}

func TestEditClub(t *testing.T) {
	tests := []struct {
		id          int
		club        *model.Club
		expectsErr  bool
		description string
	}{
		{
			id:          1,
			club:        &model.Club{Name: "name", Address: "address"},
			expectsErr:  false,
			description: "should edit chess club without errors",
		},
		{
			id:          -1,
			club:        &model.Club{Name: "name", Address: "address"},
			expectsErr:  true,
			description: "should not edit non-existing chessclub",
		},
		{
			id:          1,
			club:        &model.Club{},
			expectsErr:  true,
			description: "should not edit club with invalid/empty paramters",
		},
	}

	for _, tt := range tests {
		err := s.EditClub(tt.id, tt.club)

		if tt.expectsErr && err == nil {
			t.Error(tt.description)
		}
	}
}
