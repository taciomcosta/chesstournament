package chessclub

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
	"github.com/taciomcosta/chesstournament/internal/repository"
)

var UnexistingClubError = errors.New("Chess Club was not found")

func NewService(r repository.ChessClub) *Service {
	return &Service{r}
}

type Service struct {
	r repository.ChessClub
}

func (s Service) GetClubById(id int) (*model.ChessClub, error) {
	club, err := s.r.GetById(id)
	if err != nil {
		return nil, UnexistingClubError
	}
	return club, nil
}
