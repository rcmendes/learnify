package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/pkg/game/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/game/ucs"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//GameController defines the services provided by the game API.
type GameController struct {
	uc ucs.CreateGameUC
}

func NewCreateGameController(gameRepo storage.GameRepository, quizSrv services.QuizService) *GameController {
	return &GameController{
		uc: ucs.NewCreateGameUC(gameRepo, quizSrv),
	}

}

type createGameRequest struct {
	Category string `json:"category"`
	Quizzes  int    `json:"quizzes"`
	// player string

}

func (ctrl *GameController) CreateGame(c *fiber.Ctx) error {
	request := new(createGameRequest)

	if err := c.BodyParser(request); err != nil {
		//TODO handle error
		return err
	}

	if request.Category == "" {
		return fiber.ErrBadRequest
	}

	id, err := ctrl.uc.Create(request.Category, request.Quizzes)

	if err != nil {
		//TODO Handle Error
		return fiber.ErrInternalServerError
	}

	//TODO replace for a struct and a 201
	return c.SendString((*id).String())
}
