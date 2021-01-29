package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/controllers"
	gameCtrl "github.com/rcmendes/learnify/gameplay/pkg/game/controllers"
	gameStorage "github.com/rcmendes/learnify/gameplay/pkg/game/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

func main() {
	fmt.Println("Hello")
	app := fiber.New()

	app.Use(cors.New())

	categoryRepo := storage.NewCategoryPostgresRepository()
	quizRepo := storage.NewQuizPostgresRepository()
	imageRepo := storage.NewImageFSRepository("/home/rcmendes/git-projects/learnify/assets/images")
	audioRepo := storage.NewAudioFSRepository("/home/rcmendes/git-projects/learnify/assets/audios")
	controllers.Load(app, categoryRepo, quizRepo, imageRepo, audioRepo)

	gameRepo := gameStorage.NewGamePostgresRepository()
	quizSrv := services.NewQuizService(quizRepo, imageRepo, audioRepo)
	gameCtrl.Load(app, gameRepo, quizSrv)
	app.Listen(":8080")
}
