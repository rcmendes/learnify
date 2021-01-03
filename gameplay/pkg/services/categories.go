package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
)

//CreateCategoryInput defines the basic data to create a Category entity.
type CreateCategoryInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

//CreateCategoryOutput defines the response data of a Category entity creation.
type CreateCategoryOutput struct {
	ID uuid.UUID `json:"id"`
}

//CategoryDTO defines the data returned when fetching a Category entity.
type CategoryDTO struct {
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
}

//CategoryService defines the contract of functions provided by a Category service.
type CategoryService interface {
	Create(category CreateCategoryInput) (CreateCategoryOutput, error)
	ListAll() ([]CategoryDTO, error)
}

type categoryService struct {
	categoryRepo storage.CategoryRepository
}

//TODO Handle errors or create custom ones.

func (svc *categoryService) Create(category CreateCategoryInput) (CreateCategoryOutput, error) {
	dbCategory := storage.Category{
		Title:       category.Title,
		Description: category.Description,
	}

	//TODO Validate if it already exists (by title).

	uuid, err := svc.categoryRepo.Insert(dbCategory)

	if err != nil {
		return CreateCategoryOutput{}, err
	}

	return CreateCategoryOutput{
		*uuid,
	}, nil
}

func (svc *categoryService) ListAll() ([]CategoryDTO, error) {
	list, err := svc.categoryRepo.ListAll()

	categories := make([]CategoryDTO, 0)

	if err != nil {
		return categories, err
	}

	for _, c := range *list {
		categories = append(categories, CategoryDTO{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}

	return categories, nil
}

//NewCategoryService creates a Category service instance.
func NewCategoryService(categoryRepo storage.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo,
	}
}
