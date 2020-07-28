package data

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

var MockValidChessClub model.Club = model.Club{
	Id:      1,
	Name:    "QueenClub",
	Address: "Neverland",
}

var MockInvalidChessClub model.Club

var mockClubs []model.Club = []model.Club{MockValidChessClub}

type MockChessClubRepository struct{}

func (r *MockChessClubRepository) GetById(id int) (*model.Club, error) {
	for _, club := range mockClubs {
		if club.Id == id {
			return &club, nil
		}
	}
	return nil, model.UnexistingError
}

func (r *MockChessClubRepository) ListClubs(f model.Filter) ([]model.Club, error) {
	if f.OrderBy == "invalid" {
		return []model.Club{}, errors.New("invalid query")

	}
	return mockClubs, nil
}

func (r *MockChessClubRepository) Add(c *model.Club) (*model.Club, error) {
	return c, nil
}

func (r *MockChessClubRepository) Remove(c *model.Club) error {
	return nil
}
