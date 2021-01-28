package services

import (
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
)

//QuizDTO defines the data returned when fetching a Quiz entity.
type QuizDTO struct {
	ID       uuid.UUID `json:"uuid"`
	Category string    `json:"category"`
	Palavra  string    `json:"palavra"`
	Mot      string    `json:"mot"`
}

type ImageKind byte

const UnknownImageKind ImageKind = 0
const PngImageKind ImageKind = 1
const JpegImageKind ImageKind = 2
const GifImageKind ImageKind = 3

type ImageData struct {
	Data []byte
	Kind ImageKind
}

//QuizID defines the type of the ID of a Quiz.
type QuizID = uuid.UUID

//QuizService defines the contract of functions provided by a Quiz service.
type QuizService interface {
	ListAll() (*[]QuizDTO, error)
	ListQuizzesByCategory(category string) (*[]QuizDTO, error)
	GetQuizByID(id QuizID) (*QuizDTO, error)
	GetQuizImageByID(id QuizID) (*ImageData, error)
	FindQuizzesSameCategory(id QuizID) (*[]QuizDTO, error)
}

type quizService struct {
	quizRepo  storage.QuizRepository
	imageRepo storage.ImageRepository
}

func buildQuizDTOFrom(quiz storage.Quiz) QuizDTO {
	return QuizDTO{
		ID:       quiz.ID,
		Category: quiz.Category,
		Palavra:  quiz.Palavra,
		Mot:      quiz.Mot,
	}
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
	list, err := svc.quizRepo.FindQuizzesByCategory(category)

	if err != nil {
		return nil, err
	}

	quizzes := convertQuizListToQuizDTOList(*list)

	return &quizzes, nil
}

func (svc *quizService) GetQuizByID(id QuizID) (*QuizDTO, error) {
	quiz, err := svc.quizRepo.GetQuizByID(id)

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

func (svc *quizService) FindQuizzesSameCategory(id QuizID) (*[]QuizDTO, error) {
	list, err := svc.quizRepo.FindQuizzesSameCategory(id)

	if err != nil {
		return nil, err
	}

	quizzes := convertQuizListToQuizDTOList(*list)

	return &quizzes, nil
}

func (svc *quizService) GetQuizImageByID(id QuizID) (*ImageData, error) {
	quiz, err := svc.quizRepo.GetQuizByID(id)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	image, err := svc.imageRepo.GetImageByFilename(quiz.ImageFilename)

	extension := filepath.Ext(quiz.ImageFilename)
	imageKind := imageKindFromExtension(extension)

	if err != nil {
		//TODO handle error
		return nil, err
	}

	return &ImageData{
		Data: *image,
		Kind: imageKind,
	}, nil
}

func imageKindFromExtension(extension string) ImageKind {
	if strings.Index(extension, ".") == 0 {
		extension = extension[1:]
	}

	extension = strings.ToLower(extension)

	if extension == "png" {
		return PngImageKind
	}

	if extension == "jpg" || extension == "jpeg" {
		return JpegImageKind
	}

	if extension == "gif" {
		return GifImageKind
	}

	return UnknownImageKind
}
