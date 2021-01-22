package ucs

import (
	"github.com/rcmendes/learnify/gameplay/pkg/game/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

// type ImageUUID = uuid.UUID
// type AudioURL = string

// type GameQuiz struct {
// 	Audio    AudioURL  `json:"audio_url"`
// 	Question string    `json:"question"`
// 	AnswerID ImageUUID `json:"answer"`
// 	Options  []uuid.UUID
// }

//GameUC defines  a Create Game Use Case handler.
type GameUC interface {
	Create(category string, quizzes int) (*GameID, error)
	// Load(id GameID) (*GameDTO, error)
	// Finish(uuid uuid.UUID) error
}

type gameUC struct {
	gameRepo storage.GameRepository
	quizSrv  services.QuizService
}

//NewCreateGameUC creates a Create Game Use Case handler.
func NewGameUC(gameRepo storage.GameRepository, quizSrv services.QuizService) GameUC {
	return &gameUC{
		gameRepo: gameRepo,
		quizSrv:  quizSrv,
	}
}
