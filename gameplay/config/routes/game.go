package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
	"github.com/rcmendes/learnify/gameplay/ucs"
)

func LoadGameRoutes(app *fiber.App, gameRepo ucs.GameRepository, quizSrv services.QuizService) {

	ctrl := rest.NewGameController(nil)

	gameGroup := app.Group("/game")
	gameGroup.Post("/", ctrl.CreateGame)
	gameGroup.Get("/:id", ctrl.GetGameByID)
	// gameGroup.Get("/:id", gameController.ListAll)
	// categoriesGroup.Get("/:uuid", categoriesController.FindOneByUUID)
	// categoriesGroup.Delete("/:uuid", categoriesController.DeleteByUUID)
}
