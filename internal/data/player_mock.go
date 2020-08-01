package data

import "github.com/taciomcosta/chesstournament/internal/model"

var MockValidPlayer model.Player = model.Player{
	Id:        1,
	Ranking:   model.RankingMaster,
	FirstName: "Magnus",
	LastName:  "Carlsen",
	Address:   "Somewhere",
	Phone:     "12345678",
	Email:     "magnus.carlsen@gmail.com",
}

var MockNoClubPlayer model.Player = model.Player{
	Id:        2,
	Ranking:   model.RankingMaster,
	FirstName: "No",
	LastName:  "Club",
	Address:   "Address2",
	Phone:     "00000000",
	Email:     "no.club@gmail.com",
}

var MockInvalidPlayer model.Player = model.Player{}

var mockPlayers []model.Player = []model.Player{MockValidPlayer}

type MockPlayerRepository struct{}

func (r MockPlayerRepository) Add(p *model.Player) (*model.Player, error) {
	p.Id = 1
	return p, nil
}

func (r MockPlayerRepository) FindOne(criteria *model.Player) (*model.Player, error) {

	for _, p := range mockPlayers {
		if p.Id == criteria.Id {
			return &p, nil
		}
	}
	return nil, model.UnexistingError
}

func (r MockPlayerRepository) Remove(p *model.Player) error {
	return nil
}

func (r MockPlayerRepository) Count(criteria *model.Player) int {
	if criteria.ClubId == MockValidClubWithPlayer.Id {
		return 1
	}
	return 0
}
