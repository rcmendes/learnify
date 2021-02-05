package entities

import (
	"fmt"

	"github.com/google/uuid"
)

//QuizID defines the type of the ID of a Quiz
type QuizID = uuid.UUID

//Quiz defines the properties of a Quiz database entity.
type Quiz struct {
	ID       QuizID
	Category string
	Palavra  string
	Mot      string
	Image    MediaInfo
	Audio    MediaInfo
}

//MediaInfo defines the data structure information of a Media type.
type MediaInfo struct {
	Name string
	Data *[]byte
}

func (q *Quiz) String() string {
	return fmt.Sprintf("Quiz <category=%s, "+
		"palavra=%s, mot=%s, image=%s, audio=%s>",
		q.Category, q.Palavra, q.Mot, q.Image.Name, q.Audio.Name)
}
