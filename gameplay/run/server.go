package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/pkg/routes"
)

func main() {
	fmt.Println("Hello")
	app := fiber.New()

	routes.Load(app)

	app.Listen(":8080")
}
