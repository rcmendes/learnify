package ucs

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/pkg/game/storage"
)

func (uc *gameUC) Create(category string, quizzesNumber int) (*GameID, error) {

	quizzes, err := uc.quizSrv.ListQuizzesByCategory(category)
	if err != nil {
		//TODO handle error
		return nil, err
	}

	length := len(*quizzes)
	if length == 0 {
		return nil, fmt.Errorf("no quiz found for this category")
	}

	gameQuizzes := make([]*storage.GameQuiz, 0, quizzesNumber)

	for i := 0; i < length && len(gameQuizzes) < quizzesNumber; i++ {
		index := rand.Intn(length - 1)
		quiz := (*quizzes)[index]
		if !containsQuiz(gameQuizzes, &quiz.ID) {
			gameQuizzes = append(gameQuizzes, &storage.GameQuiz{QuizID: quiz.ID})
		}
	}

	playerID, _ := uuid.Parse("2f5b6610-b714-4258-8026-3e77ac1a30a4")

	game := storage.Game{
		Player:  playerID,
		Quizzes: gameQuizzes,
		Status:  storage.StatusCreated,
	}

	id, err := uc.gameRepo.Insert(game)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	game.ID = *id

	return id, nil
}

func containsQuiz(list []*storage.GameQuiz, id *uuid.UUID) bool {
	for i := 0; i < len(list); i++ {
		gq := list[i]
		if gq.QuizID == *id {
			return true
		}
	}

	return false
}
