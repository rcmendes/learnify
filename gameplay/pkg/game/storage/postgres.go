package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage/postgres"
)

type GameRepository interface {
	Insert(game Game) (*GameID, error)
	GetByID(id GameID) (*Game, error)
	// Update(id GameID, game Game) error
}

type gamePostgresRepository struct {
	connectFunc func() *pg.DB
}

func NewGamePostgresRepository() GameRepository {
	return &gamePostgresRepository{
		connectFunc: postgres.Connect,
	}
}

func (repo *gamePostgresRepository) Insert(game Game) (*GameID, error) {
	conn := repo.connectFunc()
	defer conn.Close()

	game.ID = uuid.New()

	_, err := conn.Conn().Model(&game).Insert()
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	for _, quiz := range game.Quizzes {
		quiz.GameID = game.ID
		_, err := conn.Model(quiz).Insert()
		if err != nil {
			//TODO Handle error
			return nil, err
		}
	}

	return &game.ID, nil
}

func (repo *gamePostgresRepository) GetByID(id GameID) (*Game, error) {
	conn := repo.connectFunc()
	defer conn.Close()

	var game Game

	if err := conn.Conn().Model(&game).Where("id=?", id).First(); err != nil {
		//TODO Handle error
		return nil, err
	}

	return &game, nil
}
