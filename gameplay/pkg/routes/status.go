package routes

import "github.com/gofiber/fiber/v2"

func Status(c *fiber.Ctx) error {
	return c.SendString("Ok")
}
