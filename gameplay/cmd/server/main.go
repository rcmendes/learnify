package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rcmendes/learnify/gameplay/adapters/database/postgres"
	"github.com/rcmendes/learnify/gameplay/config/routes"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/controllers"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

func main() {
	fmt.Println("Starting the server.")
	app := fiber.New()

	app.Use(cors.New())

	categoryRepo := storage.NewCategoryPostgresRepository()
	quizRepo := storage.NewQuizPostgresRepository()
	imageRepo := storage.NewImageFSRepository("/home/rcmendes/git-projects/learnify/assets/images")
	audioRepo := storage.NewAudioFSRepository("/home/rcmendes/git-projects/learnify/assets/audios")
	controllers.Load(app, categoryRepo, quizRepo, imageRepo, audioRepo)

	gameRepo := postgres.NewGamePostgresRepository()
	quizSrv := services.NewQuizService(quizRepo, imageRepo, audioRepo)
	routes.LoadGameRoutes(app, gameRepo, quizSrv)
	app.Listen(":8080")
}
