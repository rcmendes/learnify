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
	FindOneByID(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	GetImageByID(c *fiber.Ctx) error
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
	quizzesGroup.Get("/:uuid", quizzesController.FindOneByID)
	quizzesGroup.Get("/:uuid/image", quizzesController.GetImageByID)
	// quizzesGroup.Delete("/:uuid", quizzesController.DeleteByUUID)
}

func contextPath(c *fiber.Ctx) string {
	baseURL := c.BaseURL()
	context := c.Route().Path

	return baseURL + context + "/"
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

	return c.JSON(quizzes)
}

func (controller *quizController) FindOneByID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	uuidParam := c.Params("uuid")

	uuid, err := uuidLib.FromBytes([]byte(uuidParam))
	if err != nil {
		return err
	}

	quiz, err := controller.quizSrv.GetQuizByID(uuid)
	if err != nil {
		//TODO handle error
		return err
	}

	return c.JSON(quiz)
}

func (controller *quizController) DeleteByID(c *fiber.Ctx) error {
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

func (controller *quizController) GetImageByID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	uuidParam := c.Params("uuid")

	uuid, err := uuidLib.Parse(uuidParam)
	if err != nil {
		return err
	}

	image, err := controller.quizSrv.GetQuizImageByID(uuid)
	if err != nil {
		//TODO handle error
		return err
	}

	contentType := contentTypeFromImageKind(image.Kind)
	c.Response().Header.Add("Content-type", contentType)

	if _, err := c.Write(*&image.Data); err != nil {
		return err
	}

	return nil
}

func contentTypeFromImageKind(king services.ImageKind) string {
	if king == services.PngImageKind {
		return "image/png"
	}

	if king == services.JpegImageKind {
		return "image/jpeg"
	}

	if king == services.GifImageKind {
		return "image/gif"
	}

	return "application/octet-stream"
}
