package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage/postgres"
)

// CategoryList defines a list of categories.
type CategoryList []Category

//CategoryRepository defines the services that a Cetegory repository must provide.
type CategoryRepository interface {
	Insert(category Category) (*uuid.UUID, error)
	ListAll() (*CategoryList, error)
}

type categoryRepository struct {
	connectFunc func() *pg.DB
}

//NewCategoryPostgresRepository creates a Category repository for Postgres database.
func NewCategoryPostgresRepository() CategoryRepository {
	return &categoryRepository{
		connectFunc: postgres.Connect,
	}
}

//TODO Customize errors

func (repo *categoryRepository) Insert(category Category) (*uuid.UUID, error) {
	conn := repo.connectFunc()
	defer conn.Close()

	category.ID = uuid.New()

	result, error := conn.Model(&category).Insert()
	if error != nil {
		return nil, error
	}

	if result.RowsAffected() > 0 {
		return &category.ID, nil
	}

	return nil, nil
}

func (repo *categoryRepository) ListAll() (*CategoryList, error) {
	conn := repo.connectFunc()
	defer conn.Close()

	var categories CategoryList

	if err := conn.Model(&categories).Select(); err != nil {
		return nil, err
	}

	return &categories, nil
}
