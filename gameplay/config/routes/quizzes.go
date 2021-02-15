package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/ucs/ports"
)

func LoadQuizzesRoutes(
	app *fiber.App,
	findAllQuizzesUC ports.FindAllQuizzes,
	findQuizByCategoryNameUC ports.FindQuizByCategoryName,
	findQuizzesSameCategoryUC ports.FindQuizzesSameCategory) {

	ctrl := rest.NewQuizController(
		findAllQuizzesUC,
		findQuizByCategoryNameUC,
		findQuizzesSameCategoryUC)

	group := app.Group("/quizzes")
	group.Get("/", ctrl.ListAll)
	group.Get("/:uuid", ctrl.FindOneByID)
	group.Get("/:uuid/image", ctrl.GetImageByID)
	group.Get("/:uuid/audio", ctrl.GetAudioByID)
	// quizzesGroup.Delete("/:uuid", ctrl.DeleteByUUID)

}
