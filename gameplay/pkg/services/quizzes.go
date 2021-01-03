package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
)

//QuizDTO defines the data returned when fetching a Quiz entity.
type QuizDTO struct {
	ID            uuid.UUID `json:"uuid"`
	Category      string    `json:"category"`
	Palavra       string    `json:"palavra"`
	Mot           string    `json:"mot"`
	ImageURL      string    `json:"image_url"`
	AudioURL      string    `json:"audio_url"`
	audioFilename string
	imageFilename string
}

//QuizService defines the contract of functions provided by a Quiz service.
type QuizService interface {
	ListAll() (*[]QuizDTO, error)
	ListQuizzesByCategory(category string) (*[]QuizDTO, error)
	GetQuizByUUID(uuid uuid.UUID) (*QuizDTO, error)
	GetImageByUUID(uuid uuid.UUID) (*[]byte, error)
}

type quizService struct {
	quizRepo  storage.QuizRepository
	imageRepo storage.ImageRepository
}

func buildQuizDTOFrom(quiz storage.Quiz) QuizDTO {
	return QuizDTO{
		ID:            quiz.ID,
		imageFilename: quiz.ImageFilename,
		Category:      quiz.Category,
		Palavra:       quiz.Palavra,
		Mot:           quiz.Mot,
		audioFilename: quiz.AudioFilename,
	}
}

//BuildImageURL build the QuizDTO image URL from the contextPath.
func (dto *QuizDTO) BuildImageURL(contextPath string) error {
	//TODO create a typed error
	if dto.imageFilename == "" {
		return errors.New("Image was not defined")
	}
	dto.ImageURL = contextPath + dto.ID.String() + "/image"

	return nil
}

//BuildAudioURL build the QuizDTO audio URL from the contextPath.
func (dto *QuizDTO) BuildAudioURL(contextPath string) error {
	//TODO create a typed error
	if dto.audioFilename == "" {
		return errors.New("Audio was not defined")
	}
	dto.AudioURL = contextPath + dto.ID.String() + "/image"

	return nil
}

//TODO Handle errors or create custom ones.

//NewQuizService creates a Quiz service instance.
func NewQuizService(quizRepo storage.QuizRepository, imageRepo storage.ImageRepository) QuizService {
	return &quizService{
		quizRepo,
		imageRepo,
	}
}

func (svc *quizService) ListAll() (*[]QuizDTO, error) {
	list, err := svc.quizRepo.ListAll()

	if err != nil {
		return nil, err
	}

	quizzes := convertQuizListToQuizDTOList(*list)

	return &quizzes, nil
}

func (svc *quizService) ListQuizzesByCategory(category string) (*[]QuizDTO, error) {
	list, err := svc.quizRepo.ListQuizzesByCategory(category)

	if err != nil {
		return nil, err
	}

	quizzes := convertQuizListToQuizDTOList(*list)

	return &quizzes, nil
}

func (svc *quizService) GetQuizByUUID(id uuid.UUID) (*QuizDTO, error) {
	quiz, err := svc.quizRepo.GetQuizByUUID(id)

	if err != nil {
		return nil, err
	}

	quizDTO := buildQuizDTOFrom(*quiz)

	return &quizDTO, nil
}

func convertQuizListToQuizDTOList(list storage.QuizList) []QuizDTO {
	quizzes := make([]QuizDTO, 0, len(list))
	for _, quiz := range list {
		quizDTO := buildQuizDTOFrom(quiz)
		quizzes = append(quizzes, quizDTO)
	}

	return quizzes
}

func (svc *quizService) GetImageByUUID(id uuid.UUID) (*[]byte, error) {
	quiz, err := svc.quizRepo.GetQuizByUUID(id)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	image, err := svc.imageRepo.GetImageByID(quiz.ImageFilename)

	if err != nil {
		//TODO Handle error (wrap???)
		return nil, err
	}

	return image, nil
}
