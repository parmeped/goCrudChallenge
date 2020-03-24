package api

import (
	cs "github.com/goCrudChallenge/pkg/api/contact"
	ct "github.com/goCrudChallenge/pkg/api/contact/transport"
	"github.com/goCrudChallenge/pkg/utl/config"
	"github.com/goCrudChallenge/pkg/utl/postgres"
	"github.com/goCrudChallenge/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {

	// Tries to connect to the DB
	db, err := postgres.New(cfg.DB.PSN)
	if err != nil {
		return err
	}

	// Initializes server
	e := server.New()

	v1 := e.Group("/v1")

	// Initializes contact service
	ct.NewHTTP(cs.Initialize(db), v1)

	// Starts server
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
