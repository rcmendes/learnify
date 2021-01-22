package ucs

func (uc *gameUC) GetGameByID(id GameID) (*GameDTO, error) {
	game, err := uc.gameRepo.GetByID(id)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	quizzesDTOs := make([]GameQuizDTO, 0)

	// for _, q := range game.Quizzes {
	// 	quizzes, err := uc.quizSrv.FindQuizzesSameCategory(q.QuizID)
	// 	if err != nil {
	// 		//TODO handle error

	// 	}

	// 	dto := GameQuizDTO {}

	// 	for _, similar := range *quizzes {
	// 		if similar.ID = q.QuizID {
	// 			dto.Question = similar.Palavra
	// 		}

	// 	}

	// }

	dto := GameDTO{
		ID:       game.ID,
		PlayerID: game.Player,
		Quizzes:  quizzesDTOs,
	}

	return &dto, nil
}
