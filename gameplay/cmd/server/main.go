package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rcmendes/learnify/gameplay/adapters/database/postgres"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/adapters/filesystem"
	"github.com/rcmendes/learnify/gameplay/config/routes"
	"github.com/rcmendes/learnify/gameplay/ucs"
)

func main() {
	fmt.Println("Starting the server.")
	app := fiber.New()

	app.Use(cors.New())

	//Repositories
	categoryRepo := postgres.NewCategoryPostgresRepository()
	quizRepo := postgres.NewQuizPostgresRepository()
	gameRepo := postgres.NewGamePostgresRepository()
	imageRepo := filesystem.NewImageFSRepository("/home/rcmendes/git-projects/learnify/assets/images")
	audioRepo := filesystem.NewAudioFSRepository("/home/rcmendes/git-projects/learnify/assets/audios")

	// Use Cases: Category
	createCategoryUC := ucs.MakeCreateCategory(categoryRepo)
	findAllCategoriesUC := ucs.MakeFindAllCategories(categoryRepo)

	// Use Cases: Quiz
	findAllQuizzesUC := ucs.MakeFindAllQuizzes(quizRepo)
	findQuizByCategoryName := ucs.MakeFindQuizByCategoryName(quizRepo)
	// findQuizzesSameCategory := ucs.MakeFindQuizByCategoryName(quizRepo)
	findQuizzesSameCategory
	// Use Cases: Game
	createGameUC := ucs.MakeCreateGame(gameRepo, quizRepo)
	validateAnswerGameQuizUC := ucs.MakeValidateAnswerGameQuiz(gameRepo)
	findOneNotPlayedGameQuizUC := ucs.MakeFindOneNotPlayedGameQuiz(gameRepo)

	// Routes
	app.Get("/status", rest.Status)
	routes.LoadCategoriesRoutes(app, createCategoryUC, findAllCategoriesUC)
	routes.LoadQuizzesRoutes(app, findAllQuizzesUC, findQuizByCategoryName, find)
	routes.LoadGameRoutes(app, createGameUC, validateAnswerGameQuizUC, findOneNotPlayedGameQuizUC)
	app.Listen(":8080")
}
