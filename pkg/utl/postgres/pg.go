package postgres

import (
	"github.com/jinzhu/gorm"
	// DB dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbLogger struct{}

// New creates new database connection to a postgres database
func New(psn string, timeout int) (*gorm.DB, error) {

	db, err := gorm.Open("postgres", psn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	if timeout > 0 {
		db.InstantSet("db:lock_timeout", timeout)
	}

	return db, nil
}
