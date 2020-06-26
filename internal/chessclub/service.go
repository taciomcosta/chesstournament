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

func (s Service) CreateChessclub(c *model.ChessClub) (*model.ChessClub, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return s.r.Create(c)
}

func (s Service) ListClubs(r repository.Filter) ([]model.ChessClub, error) {
	cs, err := s.r.ListClubs(r)
	if err != nil {
		return []model.ChessClub{}, err
	}
	return cs, nil
}

func (s Service) DeleteClub(id int) (*model.ChessClub, error) {
	c, err := s.GetClubById(id)
	s.r.Remove(c)
	return c, err
}
