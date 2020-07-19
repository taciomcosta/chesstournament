package data

import "github.com/taciomcosta/chesstournament/internal/model"

type MockRepository struct{}

func (r MockRepository) FindOne(criteria interface{}) (interface{}, error) {
	criteriaWithId := criteria.(struct{ Id int })
	if criteriaWithId.Id == 1 {
		return criteriaWithId, nil
	}
	return nil, model.UnexistingError
}
