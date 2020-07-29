package model

type Club struct {
	Id      int    `json:"id" pg:",pk"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
}

type ClubRepository interface {
	GetById(int) (*Club, error)
	Add(*Club) (*Club, error)
	ListClubs(Filter) ([]Club, error)
	Remove(*Club) error
}

type Repository interface {
	FindOne(v interface{}) (interface{}, error)
}
