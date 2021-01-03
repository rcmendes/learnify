package storage

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage/postgres"
)

// QuizList defines a list of quizzes.
type QuizList []Quiz

//QuizRepository defines the contract of a Quiz entity repository.
type QuizRepository interface {
	ListAll() (*QuizList, error)
	ListQuizzesByCategory(category string) (*QuizList, error)
	GetQuizByUUID(uuid uuid.UUID) (*Quiz, error)
}

//ImageRepository defines the contract of an Image Repository.
type ImageRepository interface {
	GetImageByID(id string) (*[]byte, error)
}

type quizRepository struct {
	connectFn func() *pg.DB
}

type imageFSRepository struct {
	basePath string
}

//TODO Customize errors

//NewQuizPostgresRepository creates a Quiz repository instance.
func NewQuizPostgresRepository() QuizRepository {
	return &quizRepository{
		connectFn: postgres.Connect,
	}
}

// NewImageFSRepository creates a File System based Image repository instance.
func NewImageFSRepository(basePath string) ImageRepository {
	return &imageFSRepository{basePath: basePath}
}

func (repo *quizRepository) ListAll() (*QuizList, error) {
	conn := repo.connectFn()
	defer conn.Close()

	var quizzes QuizList

	if err := conn.Model(&quizzes).Select(); err != nil {
		return nil, err
	}

	return &quizzes, nil
}

func (repo *quizRepository) ListQuizzesByCategory(category string) (*QuizList, error) {
	conn := repo.connectFn()
	defer conn.Close()

	var quizzes QuizList

	if err := conn.Model(&quizzes).
		Where("category = ?", category).
		Select(); err != nil {

		return nil, err
	}

	return &quizzes, nil
}

func (repo *quizRepository) GetQuizByUUID(id uuid.UUID) (*Quiz, error) {
	conn := repo.connectFn()
	defer conn.Close()

	var quiz Quiz

	if err := conn.Model(&quiz).
		Where("id = ?", id).
		First(); err != nil {

		return nil, err
	}

	return &quiz, nil
}

func (repo *imageFSRepository) GetImageByID(id string) (*[]byte, error) {
	filePath := filepath.Join(repo.basePath, id)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return &data, nil
}
