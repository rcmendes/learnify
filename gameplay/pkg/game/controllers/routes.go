package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/pkg/game/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

func Load(app *fiber.App, gameRepo storage.GameRepository, quizSrv services.QuizService) {
	ctrl := NewGameController(gameRepo, quizSrv)

	gameGroup := app.Group("/game")
	gameGroup.Post("/", ctrl.CreateGame)
	// gameGroup.Get("/:id", gameController.ListAll)
	// categoriesGroup.Get("/:uuid", categoriesController.FindOneByUUID)
	// categoriesGroup.Delete("/:uuid", categoriesController.DeleteByUUID)
}
