package chessclub

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

func Mock() Service {
	return mockService{}
}

type mockService struct{}

func (m mockService) GetClubById(id int) (*model.ChessClub, error) {
	return chessclubRepository.GetById(id)
}
