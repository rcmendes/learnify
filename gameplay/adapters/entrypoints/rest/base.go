package rest

import "github.com/rcmendes/learnify/gameplay/ucs"

//GameController defines the services provided by the game API.
type GameController struct {
	createGame  ucs.CreateGame
	getGameByID ucs.GetGameByID
}

func NewGameController(gameRepo ucs.GameRepository) *GameController {
	// createGame := ucs.CreateGame{
	// 	gameRepo,
	// }
	// ctrl := GameController{
	// 	createGame,
	// 	getGameByID,
	// }

	return nil
}
