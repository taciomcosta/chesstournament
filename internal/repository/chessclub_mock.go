package repository

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

type MockChessClub struct{}

func (repository *MockChessClub) GetById(id int) (*model.ChessClub, error) {
	if id == model.MockChessClub.Id {
		return &model.MockChessClub, nil
	}
	return &model.ChessClub{}, errors.New("Non-existing resource")
}
