package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//CategoryController defines the endpoints of a Category controller.
type CategoryController interface {
	ListAll(c *fiber.Ctx) error

	FindOneByUUID(c *fiber.Ctx) error

	DeleteByUUID(c *fiber.Ctx) error

	Create(c *fiber.Ctx) error
}

type categoryController struct {
	categorySrv services.CategoryService
}

func loadCategoriesEndpoints(app *fiber.App, categoryRepo storage.CategoryRepository) {
	categorySrv := services.NewCategoryService(categoryRepo)
	categoriesController := NewCategoryController(categorySrv)
	categoriesGroup := app.Group("/categories")
	categoriesGroup.Post("/", categoriesController.Create)
	categoriesGroup.Get("/", categoriesController.ListAll)
	// categoriesGroup.Get("/:uuid", categoriesController.FindOneByUUID)
	// categoriesGroup.Delete("/:uuid", categoriesController.DeleteByUUID)
}

//TODO Handle errors

// NewCategoryController creates a Category controller.
func NewCategoryController(categorySrv services.CategoryService) CategoryController {
	return &categoryController{
		categorySrv,
	}
}

func (controller *categoryController) ListAll(c *fiber.Ctx) error {
	categories, err := controller.categorySrv.ListAll()
	if err != nil {
		return err
	}

	return c.JSON(categories)
}

func (controller *categoryController) FindOneByUUID(c *fiber.Ctx) error {

	return c.SendString("Find One Category by UUID")
}

func (controller *categoryController) DeleteByUUID(c *fiber.Ctx) error {
	return c.SendString("Delete a Category by UUID")
}

func (controller *categoryController) Create(c *fiber.Ctx) error {
	createCategoryInput := services.CreateCategoryInput{}

	if err := c.BodyParser(&createCategoryInput); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return err
	}

	output, err := controller.categorySrv.Create(createCategoryInput)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return err
	}

	return c.Status(http.StatusCreated).JSON(output)
}
