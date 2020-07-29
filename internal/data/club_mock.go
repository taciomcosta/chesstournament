package data

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

var MockValidClub model.Club = model.Club{
	Id:      1,
	Name:    "QueenClub",
	Address: "Neverland",
}

var MockInvalidClub model.Club

var mockClubs []model.Club = []model.Club{MockValidClub}

type MockClubRepository struct{}

func (r *MockClubRepository) GetById(id int) (*model.Club, error) {
	for _, club := range mockClubs {
		if club.Id == id {
			return &club, nil
		}
	}
	return nil, model.UnexistingError
}

func (r *MockClubRepository) ListClubs(f model.Filter) ([]model.Club, error) {
	if f.OrderBy == "invalid" {
		return []model.Club{}, errors.New("invalid query")

	}
	return mockClubs, nil
}

func (r *MockClubRepository) Add(c *model.Club) (*model.Club, error) {
	return c, nil
}

func (r *MockClubRepository) Remove(c *model.Club) error {
	return nil
}
