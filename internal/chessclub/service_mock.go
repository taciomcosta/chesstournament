package chessclub

import (
	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

func Mock() Service {
	return mockService{}
}

type mockService struct{}

func (m mockService) GetClubById(id int) (*model.ChessClub, error) {
	repository := repository.MockChessClub{}
	return repository.GetById(id)
}
