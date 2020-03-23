package api

import (
	"github.com/goCrudChallenge/pkg/utl/zlog"
	"github.com/gorsk/pkg/api/user"

	"github.com/goCrudChallenge/pkg/utl/config"
	"github.com/goCrudChallenge/pkg/utl/postgres"
	"github.com/goCrudChallenge/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {

	// Tries to connect to the DB
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	// Initializes Log
	log := zlog.New()

	// Initializes server
	e := server.New()
	// TODO: Swagger
	//e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	v1 := e.Group("/v1")

	// Initializes contact service
	ut.NewHTTP(cl.New(user.Initialize(db, rbac, sec), log), v1)

	// Starts server
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
