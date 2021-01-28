package controllers

import (
	"github.com/rcmendes/learnify/gameplay/pkg/game/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/game/ucs"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//GameController defines the services provided by the game API.
type GameController struct {
	uc ucs.GameUC
}

func NewGameController(gameRepo storage.GameRepository, quizSrv services.QuizService) *GameController {
	return &GameController{
		uc: ucs.NewGameUC(gameRepo, quizSrv),
	}

}
