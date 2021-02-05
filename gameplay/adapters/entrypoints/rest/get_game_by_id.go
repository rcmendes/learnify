package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/ucs"
)

func (ctrl *GameController) GetGameByID(c *fiber.Ctx) error {
	// gameID := c.Params("id")

	ids := []string{"44e420ad-c5fd-4d26-9ed5-2838f2450147",
		"93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0",
		"32ed9040-1c14-4875-9ab3-c504dc7a962c",
		"5509e717-a4de-4457-b3b9-5b1676e591e4"}

	options := make([]ucs.MediaLink, 0, 4)

	for _, id := range ids {
		media := ucs.MediaLink{
			ID:  uuid.MustParse(id),
			URI: fmt.Sprintf("%s/quizzes/%s/image", c.BaseURL(), id),
		}
		options = append(options, media)
	}

	audio := ucs.MediaLink{
		ID:  uuid.MustParse("44e420ad-c5fd-4d26-9ed5-2838f2450147"),
		URI: fmt.Sprintf("%s/quizzes/%s/audio", c.BaseURL(), "44e420ad-c5fd-4d26-9ed5-2838f2450147"),
	}

	response := ucs.GetGameByIDResponse{
		Question: "What is this animal?",
		Audio:    audio,
		Images:   options,
	}

	return c.JSON(response)

}
