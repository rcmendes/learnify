package ucs

import (
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//GetGameByID defines the base structure to the 'Load a Game' Use Case.
type GetGameByID struct {
	gameRepo GameRepository
	quizSrv  services.QuizService
}

//GetByID is a method that loads a game data.
func (uc *GetGameByID) GetByID(request GetGameByIDRequest) (*GetGameByIDResponse, error) {
	_, err := uc.gameRepo.GetByID(request.ID)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	response := &GetGameByIDResponse{
		// Question: game.,
	}

	return response, nil
}
