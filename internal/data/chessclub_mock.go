package data

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

var MockValidChessClub model.ChessClub = model.ChessClub{
	Id:      1,
	Name:    "QueenClub",
	Address: "Neverland",
}

var MockInvalidChessClub model.ChessClub

var mockClubs []model.ChessClub = []model.ChessClub{MockValidChessClub}

type MockChessClubRepository struct{}

func (r *MockChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	for _, club := range mockClubs {
		if club.Id == id {
			return &club, nil
		}
	}
	return nil, model.UnexistingError
}

func (r *MockChessClubRepository) ListClubs(f model.Filter) ([]model.ChessClub, error) {
	if f.OrderBy == "invalid" {
		return []model.ChessClub{}, errors.New("invalid query")

	}
	return mockClubs, nil
}

func (r *MockChessClubRepository) Add(c *model.ChessClub) (*model.ChessClub, error) {
	return c, nil
}

func (r *MockChessClubRepository) Remove(c *model.ChessClub) error {
	return nil
}
