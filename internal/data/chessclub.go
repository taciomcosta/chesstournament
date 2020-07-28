package data

import (
	"github.com/taciomcosta/chesstournament/internal/model"
)

type ChessClubRepository struct{}

func (r ChessClubRepository) GetById(id int) (*model.Club, error) {
	club := &model.Club{Id: id}
	if err := db.Select(club); err != nil {
		return nil, err
	}
	return club, nil
}

func (r ChessClubRepository) Add(c *model.Club) (*model.Club, error) {
	_, err := db.Model(c).OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, model.UnknownError
	}
	return c, nil
}

func (r ChessClubRepository) ListClubs(lr model.Filter) ([]model.Club, error) {
	cs := make([]model.Club, 0)
	if err := lr.Validate(); err != nil {
		return cs, err
	}
	err := db.Model(&cs).OrderExpr(lr.OrderBy).Offset(lr.Offset).Limit(lr.Limit).Select()
	return cs, err
}

func (r ChessClubRepository) Remove(c *model.Club) error {
	return db.Delete(c)
}
