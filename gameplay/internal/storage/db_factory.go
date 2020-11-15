package storage

import (
	"github.com/go-pg/pg/v10"
)

/*TODO Add parameters*/

// Connect connects to the postgres database.
func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "dev",
		Password: "dev",
		Database: "learnify_game_dev",
	})
}
