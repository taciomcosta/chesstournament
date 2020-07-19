package model

type ChessClub struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
}

type ChessClubRepository interface {
	GetById(int) (*ChessClub, error)
	Add(*ChessClub) (*ChessClub, error)
	ListClubs(Filter) ([]ChessClub, error)
	Remove(*ChessClub) error
}

type Repository interface {
	FindOne(v interface{}) (interface{}, error)
}
