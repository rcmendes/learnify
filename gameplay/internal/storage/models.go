package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type storableID = uuid.UUID
type QuizID = storableID
type CategoryID = storableID

//Storable defines some attributes of a database entity.
type Storable struct {
	ID        storableID `pg:"id"`
	CreatedAt *time.Time `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
}

func (s *Storable) String() string {
	return fmt.Sprintf("Storable <ID=%d, createdAt=%q, updatedAt=%q>",
		s.ID, s.CreatedAt, s.UpdatedAt)
}

//Category defines the properties of a Category database entity.
type Category struct {
	Storable
	tableName   struct{} `pg:"categories"`
	Title       string   `pg:"title"`
	Description *string  `pg:"description"`
}

func (c *Category) String() string {
	return fmt.Sprintf("User<title=%s, description=%q, storable=(%s)>",
		c.Title, *c.Description, c.Storable.String())
}

//Quiz defines the properties of a Quiz database entity.
type Quiz struct {
	Storable
	tableName     struct{} `pg:"quizzes"`
	ImageFilename string   `pg:"image_filename"`
	Category      string   `pg:"category"`
	Palavra       string   `pg:"palavra"`
	Mot           string   `pg:"mot"`
	AudioFilename string   `pg:"audio_filename"`
}

func (q *Quiz) String() string {
	return fmt.Sprintf("Quiz <url=%s, category=%s, "+
		"palavra=%s, mot=%s, audio=%s, storable=(%s)>",
		q.ImageFilename, q.Category, q.Palavra, q.Mot, q.AudioFilename,
		q.Storable.String())
}
