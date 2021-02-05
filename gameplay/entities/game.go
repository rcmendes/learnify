package entities

import (
	"github.com/google/uuid"
)

//GameID defines the type of the ID of a Game
type GameID = uuid.UUID

//Game defines the structure of a Game.
type Game struct {
	ID      GameID
	Player  Player
	Status  statusType
	Quizzes []*GameQuiz
}

//GameQuiz defines the structure of Quizzes associated to a Game.
type GameQuiz struct {
	Quiz   Quiz
	Status statusType
}

type statusType = uint8

type gameStatusOptions struct {
	//Created represents the status of a game when it's created.
	Created statusType
	//Playing represents the status of a game when it's been played.
	Playing statusType
	//Finished represents the status of a game when it's finished.
	Finished statusType
}

// GameStatus defines the available status of a Game.
var GameStatus = gameStatusOptions{
	Created:  1,
	Playing:  2,
	Finished: 3,
}

type gameQuizStatusOptions struct {
	//NotPlayed represents the status of a game when it's created.
	NotPlayed statusType
	//Playing represents the status of a game when it's been played.
	Playing statusType
	//Finished represents the status of a game when it's finished.
	Finished statusType
}

// GameQuizStatus defines the available status of a GameQuiz.
var GameQuizStatus = gameQuizStatusOptions{
	NotPlayed: 1,
	Playing:   2,
	Finished:  3,
}
