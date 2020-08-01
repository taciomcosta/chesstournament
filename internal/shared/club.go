package shared

import (
	"errors"

	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) GetClubById(id int) (*model.Club, error) {
	return s.chessclubRepository.GetById(id)
}

func (s service) CreateClub(c *model.Club) (*model.Club, error) {
	if err := model.Validate(c); err != nil {
		return nil, err
	}
	return s.chessclubRepository.Add(c)
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
	cs, err := s.chessclubRepository.ListClubs(r)
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
	s.chessclubRepository.Remove(club)
	return club, err
}

func (s service) hasAssociatedPlayers(club *model.Club) bool {
	playersCount := s.playerRepository.Count(&model.Player{ClubId: club.Id})
	return playersCount > 0
}
