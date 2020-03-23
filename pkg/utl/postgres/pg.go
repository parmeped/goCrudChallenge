package postgres

import (
	"time"

	"github.com/go-pg/pg"
	// DB adapter
	_ "github.com/lib/pq"
)

type dbLogger struct{}

// New creates new database connection to a postgres database
func New(psn string, timeout int) (*pg.DB, error) {
	u, err := pg.ParseURL(psn)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(u)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if timeout > 0 {
		db.WithTimeout(time.Second * time.Duration(timeout))
	}

	return db, nil
}
