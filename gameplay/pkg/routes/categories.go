package routes

import "github.com/gofiber/fiber/v2"

//CategoryController defines the endpoints of a Category controller.
type CategoryController interface {
	ListAll(c *fiber.Ctx) error

	FindOneByUUID(c *fiber.Ctx) error

	DeleteByUUID(c *fiber.Ctx) error

	Create(c *fiber.Ctx) error
}

type categoryController struct {
}

// NewCategoryController creates a Category controller.
func NewCategoryController() CategoryController {
	return &categoryController{}
}

func (controler *categoryController) ListAll(c *fiber.Ctx) error {
	return c.SendString("List All Categories")
}

func (controler *categoryController) FindOneByUUID(c *fiber.Ctx) error {
	return c.SendString("Find One Category by UUID")
}

func (controler *categoryController) DeleteByUUID(c *fiber.Ctx) error {
	return c.SendString("Delete a Category by UUID")
}

func (controler *categoryController) Create(c *fiber.Ctx) error {
	return c.SendString("Create a  Category")
}
