package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify/gameplay/internal/storage/postgres"
)

// QuizList defines a list of quizzes.
type QuizList []Quiz

type quizRepository struct {
	connectFn func() *pg.DB
}

//QuizRepository defines the contract of a Quiz entity repository.
type QuizRepository interface {
	ListAll() (*QuizList, error)
	FindQuizzesByCategory(category string) (*QuizList, error)
	GetQuizByID(id QuizID) (*Quiz, error)
	FindQuizzesSameCategory(id QuizID) (*QuizList, error)
}

//TODO Customize errors

//NewQuizPostgresRepository creates a Quiz repository instance.
func NewQuizPostgresRepository() QuizRepository {
	return &quizRepository{
		connectFn: postgres.Connect,
	}
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

func (repo *quizRepository) FindQuizzesByCategory(category string) (*QuizList, error) {
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

func (repo *quizRepository) FindQuizzesSameCategory(id QuizID) (*QuizList, error) {
	conn := repo.connectFn()
	defer conn.Close()

	//TODO improve this code for better performance. 2 Queries => 1 query.

	// TODO refactor for handle error
	quiz, err := repo.GetQuizByID(id)
	if err != nil {
		return nil, err
	}

	var quizzes QuizList

	if err := conn.Model(&quizzes).
		Where("category = ?", quiz.Category).
		Select(); err != nil {

		return nil, err
	}

	return &quizzes, nil
}

func (repo *quizRepository) GetQuizByID(id QuizID) (*Quiz, error) {
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
