package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetClubById(id int) (*model.ChessClub, error) {
	club, err := s.chessclubRepository.GetById(id)
	if err != nil {
		return nil, model.UnexistingClubError
	}
	return club, nil
}

func (s service) CreateChessclub(c *model.ChessClub) (*model.ChessClub, error) {
	if err := model.Validate(c); err != nil {
		return nil, err
	}
	return s.chessclubRepository.Add(c)
}

func (s service) EditChessclub(id int, c *model.ChessClub) error {
	if _, err := s.GetClubById(id); err != nil {
		return err
	}
	c.Id = id
	_, err := s.CreateChessclub(c)
	return err
}

func (s service) ListClubs(r model.Filter) ([]model.ChessClub, error) {
	cs, err := s.chessclubRepository.ListClubs(r)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (s service) DeleteClub(id int) (*model.ChessClub, error) {
	c, err := s.GetClubById(id)
	s.chessclubRepository.Remove(c)
	return c, err
}
