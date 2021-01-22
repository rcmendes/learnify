package ucs

import (
	"github.com/google/uuid"
)

type GameID = uuid.UUID
type ImageID = uuid.UUID
type AudioID = uuid.UUID
type PlayerID = uuid.UUID

// type QuizDTO struct {
// 	ID uuid.UUID `json:"id"`
// }

// type CreateGameResponse struct {
// 	ID      uuid.UUID `json:"id"`
// 	Quizzes []QuizDTO `json:"quizzes"`
// }

type QuizImage struct {
	ID ImageID
}

type GameQuizDTO struct {
	Question string
	Images   []ImageID
}

type GameDTO struct {
	ID       GameID
	PlayerID PlayerID
	Quizzes  []GameQuizDTO
}

// func (dto *GameDTO) Load(game storage.Game, quizzes []*intStorage.Quiz) {
// 	dto.ID = game.ID
// 	dto.PlayerID = game.Player
// 	for _, q := range quizzes {
// 		quizDTO := GameQuizDTO{
// 			Question: q.Palavra,
// 			Images: q.,
// 		}

// 	}
// }
