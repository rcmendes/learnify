package controllers

import (
	"github.com/gofiber/fiber/v2"
	uuidLib "github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/internal/storage"
	"github.com/rcmendes/learnify/gameplay/pkg/services"
)

//QuizController defines the endpoints of a Quiz controller.
type QuizController interface {
	ListAll(c *fiber.Ctx) error

	FindOneByUUID(c *fiber.Ctx) error

	DeleteByUUID(c *fiber.Ctx) error

	Create(c *fiber.Ctx) error

	GetImageByUUID(c *fiber.Ctx) error
}

type quizController struct {
	quizSrv services.QuizService
}

//TODO Handle errors

func loadQuizzesEndpoints(app *fiber.App, quizRepo storage.QuizRepository, imageRepo storage.ImageRepository) {
	quizSrv := services.NewQuizService(quizRepo, imageRepo)
	quizzesController := NewQuizController(quizSrv)
	quizzesGroup := app.Group("/quizzes")
	// quizzesGroup.Post("/", quizzesController.Create)
	quizzesGroup.Get("/", quizzesController.ListAll)
	quizzesGroup.Get("/:uuid", quizzesController.FindOneByUUID)
	quizzesGroup.Get("/:uuid/image", quizzesController.GetImageByUUID)
	// quizzesGroup.Delete("/:uuid", quizzesController.DeleteByUUID)
}

func contextPath(c *fiber.Ctx) string {
	baseURL := c.BaseURL()
	context := c.Route().Path

	return baseURL + context + "/"
}

func buildQuizDTOImageAndAudioURLs(quiz *services.QuizDTO, c *fiber.Ctx) {
	contextPath := contextPath(c)
	_ = quiz.BuildImageURL(contextPath)
	_ = quiz.BuildAudioURL(contextPath)
}

// NewQuizController creates a Quiz controller.
func NewQuizController(quizSrv services.QuizService) QuizController {
	return &quizController{
		quizSrv,
	}
}

func (controller *quizController) ListAll(c *fiber.Ctx) error {
	category := c.Query("category", "")

	var quizzes *[]services.QuizDTO
	var err error
	if category == "" {
		quizzes, err = controller.quizSrv.ListAll()
		if err != nil {
			return err
		}

	} else {
		quizzes, err = controller.quizSrv.ListQuizzesByCategory(category)
	}
	if err != nil {
		return err
	}
	for i := 0; i < len(*quizzes); i++ {
		buildQuizDTOImageAndAudioURLs(&(*quizzes)[i], c)
	}

	return c.JSON(quizzes)
}

func (controller *quizController) FindOneByUUID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	uuidParam := c.Params("uuid")

	uuid, err := uuidLib.FromBytes([]byte(uuidParam))
	if err != nil {
		return err
	}

	quiz, err := controller.quizSrv.GetQuizByUUID(uuid)
	if err != nil {
		//TODO handle error
		return err
	}

	buildQuizDTOImageAndAudioURLs(quiz, c)

	return c.JSON(quiz)
}

func (controller *quizController) DeleteByUUID(c *fiber.Ctx) error {
	return c.SendString("Delete a Quiz by UUID")
}

func (controller *quizController) Create(c *fiber.Ctx) error {
	// createQuizInput := services.CreateQuizInput{}

	// if err := c.BodyParser(&createQuizInput); err != nil {
	// 	c.SendStatus(http.StatusBadRequest)
	// 	return err
	// }

	// output, err := controller.quizSrv.Create(createQuizInput)
	// if err != nil {
	// 	c.SendStatus(http.StatusInternalServerError)
	// 	return err
	// }

	// return c.Status(http.StatusCreated).JSON(output)
	return c.SendString("Create a Quiz")
}

func (controller *quizController) GetImageByUUID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	uuidParam := c.Params("uuid")

	uuid, err := uuidLib.Parse(uuidParam)
	if err != nil {
		return err
	}

	image, err := controller.quizSrv.GetImageByUUID(uuid)
	if err != nil {
		//TODO handle error
		return err
	}

	if _, err := c.Write(*image); err != nil {
		return err
	}

	return nil
}
