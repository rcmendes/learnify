package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/ucs"
)

func (ctrl *GameController) CreateGame(c *fiber.Ctx) error {
	request := new(ucs.CreateGameRequest)

	if err := c.BodyParser(request); err != nil {
		//TODO handle error
		return err
	}

	if request.Category == "" {
		return fiber.ErrBadRequest
	}

	// id, err := ctrl.uc.Create(request.Category, request.Quizzes)

	// if err != nil {
	// 	//TODO Handle Error
	// 	return fiber.ErrInternalServerError
	// }

	//TODO replace for a struct and a 201
	return c.SendString(request.Category)
}
