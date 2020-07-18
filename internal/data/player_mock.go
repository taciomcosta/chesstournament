package data

import "github.com/taciomcosta/chesstournament/internal/model"

var MockValidPlayer model.Player = model.Player{
	Id:        1,
	ClubId:    1,
	Ranking:   model.RankingMaster,
	FirstName: "Magnus",
	LastName:  "Carlsen",
	Address:   "Somewhere",
	Phone:     "12345678",
	Email:     "magnus.carlsen@gmail.com",
}

var MockNoClubPlayer model.Player = model.Player{
	Id:        2,
	ClubId:    2,
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
	if p.FirstName == "invalid" {
		return nil, model.UnknownError
	}
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
