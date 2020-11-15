package storage

import "github.com/google/uuid"

//CategoryRepository defines the services that a Cetegory repository must provide.
type CategoryRepository interface {
	Insert(category Category) (*uuid.UUID, error)
	ListAll() (*[]Category, error)
}

type categoryRepository struct{}

//NewCategoryPostgresRepository creates a Category repository for Postgres database.
func NewCategoryPostgresRepository() CategoryRepository {
	return &categoryRepository{}
}

func (repo *categoryRepository) Insert(category Category) (*uuid.UUID, error) {
	return nil, nil
}

func (repo *categoryRepository) ListAll() (*[]Category, error) {
	return nil, nil
}
