package shared

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

func (s service) CreatePlayer(p *model.Player) (*model.Player, error) {
	if err := model.Validate(p); err != nil {
		return nil, err
	}
	if _, err := s.GetClubById(p.ClubId); err != nil {
		return nil, err
	}
	return s.playerRepository.Add(p)
}

func (s service) GetPlayerById(id int) (*model.Player, error) {
	p, err := s.playerRepository.FindOne(&model.Player{Id: id})
	if err != nil {
		return nil, model.UnexistingError
	}
	return p, nil
}

func (s service) DeletePlayer(id int) (*model.Player, error) {
	p, err := s.GetPlayerById(id)
	s.playerRepository.Remove(p)
	return p, err
}
