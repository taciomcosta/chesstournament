package repository

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

type MockChessClubRepository struct{}

func (r *MockChessClubRepository) GetById(id int) (*model.ChessClub, error) {
	club := mockClubs()
	for _, c := range club {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, errors.New("Non-existing resource")
}

func (r *MockChessClubRepository) ListClubs(f model.Filter) ([]model.ChessClub, error) {
	if f.OrderBy == "invalid" {
		return []model.ChessClub{}, errors.New("invalid query")

	}
	return mockClubs(), nil
}

func (r *MockChessClubRepository) Add(c *model.ChessClub) (*model.ChessClub, error) {
	return c, nil
}

func mockClubs() []model.ChessClub {
	return []model.ChessClub{
		{
			Id:      1,
			Name:    "QueenClub",
			Address: "Neverland",
		},
	}
}

func (r *MockChessClubRepository) Remove(c *model.ChessClub) error {
	return nil
}
