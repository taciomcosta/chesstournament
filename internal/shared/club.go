package shared

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetClubById(id int) (*model.Club, error) {
	return s.repository.Club.GetById(id)
}

func (s service) CreateClub(c *model.Club) (*model.Club, error) {
	if err := model.Validate(c); err != nil {
		return nil, err
	}
	return s.repository.Club.Add(c)
}

func (s service) EditClub(id int, c *model.Club) error {
	if _, err := s.GetClubById(id); err != nil {
		return err
	}
	c.Id = id
	_, err := s.CreateClub(c)
	return err
}

func (s service) ListClubs(r model.Filter) ([]model.Club, error) {
	cs, err := s.repository.Club.ListClubs(r)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (s service) DeleteClub(id int) (*model.Club, error) {
	club, err := s.GetClubById(id)
	if err != nil {
		return nil, err
	}
	if s.hasAssociatedPlayers(club) {
		return nil, errors.New("Cannot delete club with associated players")
	}
	s.repository.Club.Remove(club)
	return club, err
}

func (s service) hasAssociatedPlayers(club *model.Club) bool {
	criteria := &model.Player{ClubId: club.Id}
	return s.repository.Player.Count(criteria) > 0
}
