package routes

import "github.com/gofiber/fiber/v2"

//Load sets all the endpoints on the application.
func Load(app *fiber.App) {
	app.Get("/status", Status)

	categoriesControler := NewCategoryController()
	categoriesGroup := app.Group("/categories")
	categoriesGroup.Post("/", categoriesControler.Create)
	categoriesGroup.Get("/", categoriesControler.ListAll)
	categoriesGroup.Get("/:uuid", categoriesControler.FindOneByUUID)
	categoriesGroup.Delete("/:uuid", categoriesControler.DeleteByUUID)
}
