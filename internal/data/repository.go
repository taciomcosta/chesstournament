package data

type Repository struct{}

func (r Repository) FindOne(criteria interface{}) (interface{}, error) {
	if err := db.Select(criteria); err != nil {
		return nil, err
	}
	return criteria, nil
}
