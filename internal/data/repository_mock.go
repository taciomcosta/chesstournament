package data

import "github.com/taciomcosta/chesstournament/internal/model"

func MockRepository() model.Repository {
	return model.Repository{
		Club:   MockClubRepository{},
		Player: MockPlayerRepository{},
	}
}
