package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
)

//Load sets all the endpoints on the application.
func Load(app *fiber.App, categoryRepo storage.CategoryRepository, quizRepo storage.QuizRepository, imageRepo storage.ImageRepository) {
	app.Get("/status", Status)

	loadCategoriesEndpoints(app, categoryRepo)
	loadQuizzesEndpoints(app, quizRepo, imageRepo)
}
