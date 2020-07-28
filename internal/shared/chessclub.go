package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetClubById(id int) (*model.Club, error) {
	return s.chessclubRepository.GetById(id)
}

func (s service) CreateChessclub(c *model.Club) (*model.Club, error) {
	if err := model.Validate(c); err != nil {
		return nil, err
	}
	return s.chessclubRepository.Add(c)
}

func (s service) EditChessclub(id int, c *model.Club) error {
	if _, err := s.GetClubById(id); err != nil {
		return err
	}
	c.Id = id
	_, err := s.CreateChessclub(c)
	return err
}

func (s service) ListClubs(r model.Filter) ([]model.Club, error) {
	cs, err := s.chessclubRepository.ListClubs(r)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (s service) DeleteClub(id int) (*model.Club, error) {
	c, err := s.GetClubById(id)
	s.chessclubRepository.Remove(c)
	return c, err
}
