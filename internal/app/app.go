package app

import (
	"seeker/internal/app/config"
	"seeker/internal/app/modules"
	"seeker/pkg/db/postgres"

	"github.com/jackc/pgx"
)

func Start() error {
	cfg := config.Load()
	server, router := NewHttpServer()

	postgresqlClient := postgres.NewClient(pgx.ConnConfig{
		Host:     cfg.DB.Host,
		Port:     uint16(cfg.DB.Port),
		Database: cfg.DB.Name,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
	})

	modules.NewUserModule(router, postgresqlClient)
	auth := modules.NewAuthModule(router, postgresqlClient)
	talent := modules.NewTalentModule(router, postgresqlClient, auth)
	recruiter := modules.NewRecruiterModule(router, postgresqlClient, auth)
	modules.NewJobModule(router, postgresqlClient, recruiter, talent)

	return server.Start(cfg.HTTP.Port)
}
