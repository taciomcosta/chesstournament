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
