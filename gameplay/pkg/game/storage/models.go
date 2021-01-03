package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
)

type GameID = uuid.UUID
type QuizID = uuid.UUID
type PlayerID = uuid.UUID
type status string

var StatusCreated status = "C"
var StatusRunning status = "R"
var StatusFinished status = "F"

type Game struct {
	storage.Storable
	tableName struct{}    `pg:"game"`
	Player    PlayerID    `pg:"player_id"`
	Status    status      `pg:"status"`
	Quizzes   []*GameQuiz //`pg:"rel:has-many"`
}

type GameQuiz struct {
	tableName struct{}   `pg:"game_quiz"`
	GameID    GameID     `pg:"game_id"`
	QuizID    QuizID     `pg:"quiz_id"`
	Success   bool       `pg:"success"`
	UpdatedAt *time.Time `pg:"updated_at"`
}
