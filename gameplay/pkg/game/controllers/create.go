package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//PlayerID defines the ID of a player
type PlayerID = uuid.UUID

type createGameRequest struct {
	Category string `json:"category"`
	// PlayerID PlayerID

}

func (ctrl *GameController) CreateGame(c *fiber.Ctx) error {
	request := new(createGameRequest)

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
