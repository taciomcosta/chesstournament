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
	thenAssertValueIsNotNil(t, err)
}

func TestListClubsValidFilters(t *testing.T) {
	clubs, err := s.ListClubs(model.Filter{})
	thenAssertSliceLenIs(t, clubs, 2)
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

func TestDeleteClubWithPlayers(t *testing.T) {
	club, err := s.DeleteClub(2)
	thenAssertValueIsNil(t, club)
	thenAssertValueIsNotNil(t, err)
}

func TestEditClub(t *testing.T) {
	err := s.EditClub(1, &data.MockValidClub)
	thenAssertErrorIsNil(t, err)
}

func TestEditUnexistentClub(t *testing.T) {
	err := s.EditClub(-1, &data.MockValidClub)
	thenAssertValueIsNotNil(t, err)
}

func TestEditClubInvalidInput(t *testing.T) {
	err := s.EditClub(1, &data.MockInvalidClub)
	thenAssertValueIsNotNil(t, err)
}
