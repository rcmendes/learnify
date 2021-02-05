package ucs

import (
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//CreateGame defines the base structure to the 'Create a Game' Use Case.
type CreateGame struct {
	gameRepo GameRepository
	quizSrv  services.QuizService
}

//Create is a method that creates a game.
func (uc *CreateGame) Create(newGame CreateGameRequest) (*CreateGameResponse, error) {
	// category := newGame.Category
	// quizzes := int(newGame.Quizzes)

	// quizList, err := uc.quizSrv.ListQuizzesByCategory(category)
	// if err != nil {
	// 	//TODO handle error
	// 	return nil, err
	// }

	// length := len(*quizList)
	// if length == 0 {
	// 	return nil, fmt.Errorf("no quiz found for this category")
	// }

	// gameQuizzes := make([]*game.GameQuiz, 0, quizzes)

	// for i := 0; i < length && len(gameQuizzes) < quizzes; i++ {
	// 	index := rand.Intn(length - 1)
	// 	quiz := (*quizList)[index]
	// 	if !containsQuiz(gameQuizzes, &quiz.ID) {
	// 		gameQuizzes = append(gameQuizzes, &game.GameQuiz{QuizID: quiz.ID})
	// 	}
	// }

	// playerID, _ := uuid.Parse("2f5b6610-b714-4258-8026-3e77ac1a30a4")

	// game := game.Game{
	// 	Player:  playerID,
	// 	Quizzes: gameQuizzes,
	// 	Status:  game.StatusCreated,
	// }

	// id, err := uc.gameRepo.Insert(game)

	// if err != nil {
	// 	//TODO handle error
	// 	return nil, err
	// }

	// response := &CreateGameResponse{
	// 	ID: *id,
	// }

	// return response, nil

	return nil, nil
}

// func containsQuiz(list []*game.GameQuiz, id *uuid.UUID) bool {
// 	for i := 0; i < len(list); i++ {
// 		gq := list[i]
// 		if gq.QuizID == *id {
// 			return true
// 		}
// 	}

// 	return false
// }
