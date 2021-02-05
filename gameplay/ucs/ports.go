package ucs

import (
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/entities"
)

//CreateGameRequest defines the contract of a request for creating a Game.
type CreateGameRequest struct {
	Category string    `json:"category"`
	PlayerID uuid.UUID `json:"player_id"`
	Quizzes  uint      `json:"quizzes"`
}

//CreateGameResponse defines the contract of the response of creating a Game.
type CreateGameResponse struct {
	ID uuid.UUID `json:"id"`
}

//GetGameByIDRequest defines the contract of the request for loading a Game.
type GetGameByIDRequest struct {
	ID uuid.UUID `json:"id"`
}

//GetGameByIDResponse defines the contract of the response of the loaded a Game.
type GetGameByIDResponse struct {
	Question string      `json:"question"`
	Images   []MediaLink `json:"images"`
	Audio    MediaLink   `json:"audio"`
	// Images   []uuid.UUID `json:"images"`
	// Audio    uuid.UUID   `json:"audio"`
}

type MediaLink struct {
	ID  uuid.UUID `json:"id"`
	URI string    `json:"uri"`
}

type GameRepository interface {
	Insert(game entities.Game) error
	GetByID(id entities.GameID) (*entities.Game, error)
	// Update(id GameID, game Game) error
}
