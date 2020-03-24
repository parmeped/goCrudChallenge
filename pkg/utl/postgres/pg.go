package postgres

import (
	"github.com/jinzhu/gorm"
	// DB dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbLogger struct{}

// New creates new database connection to a postgres database
func New(psn string) (*gorm.DB, error) {

	db, err := gorm.Open("postgres", psn)
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
