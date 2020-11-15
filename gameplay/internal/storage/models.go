package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// type Storable struct {
// 	ID   int64
// 	UUID uuid.UUID
// }

//Category defines the properties of a Category.
type Category struct {
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	ID          int64
	UUID        uuid.UUID
	Title       string
	Description *string
}

func (c *Category) String() string {
	return fmt.Sprintf("User<title=%s, description=%q, id=%d, UUID=%s, createdAt=%q, updatedAt=%q>",
		c.Title, *c.Description, c.ID, c.UUID, c.CreatedAt, c.UpdatedAt)
}
