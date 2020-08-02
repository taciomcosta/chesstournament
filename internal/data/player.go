package data

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type PlayerRepository struct {
	base baseRepository
}

func (r PlayerRepository) Add(p *model.Player) (*model.Player, error) {
	return p, r.base.Add(p)
}

func (r PlayerRepository) FindOne(criteria *model.Player) (*model.Player, error) {
	return criteria, r.base.FindOne(criteria)
}

func (r PlayerRepository) Remove(player *model.Player) error {
	return r.base.Remove(player)
}

func (r PlayerRepository) Count(criteria *model.Player) int {
	count, err := db.Model(criteria).Where("club_id = ?", criteria.ClubId).Count()
	if err != nil {
		return 0
	}
	return count
}
