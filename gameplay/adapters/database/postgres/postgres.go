package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify/gameplay/adapters/database/postgres/models"
	"github.com/rcmendes/learnify/gameplay/entities"
	"github.com/rcmendes/learnify/gameplay/internal/storage/postgres"
	"github.com/rcmendes/learnify/gameplay/ucs"
)

type gamePostgresRepository struct {
	connectFunc func() *pg.DB
}

func NewGamePostgresRepository() ucs.GameRepository {
	return &gamePostgresRepository{
		connectFunc: postgres.Connect,
	}
}

func (repo *gamePostgresRepository) Insert(game entities.Game) error {
	conn := repo.connectFunc()
	defer conn.Close()

	model := models.Game{}
	model.Load(game)

	_, err := conn.Conn().Model(&model).Insert()
	if err != nil {
		//TODO Handle error
		return err
	}

	for _, quiz := range model.Quizzes {
		quiz.GameID = model.ID
		_, err := conn.Model(quiz).Insert()
		if err != nil {
			//TODO Handle error
			return err
		}
	}

	return nil
}

func (repo *gamePostgresRepository) GetByID(id entities.GameID) (*entities.Game, error) {
	conn := repo.connectFunc()
	defer conn.Close()

	var model models.Game

	if err := conn.Conn().Model(&model).Where("id=?", id).First(); err != nil {
		//TODO Handle error
		return nil, err
	}

	game := model.To()
	return &game, nil
}
