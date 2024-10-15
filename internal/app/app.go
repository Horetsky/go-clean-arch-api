package app

import (
	"seeker/internal/app/config"
	"seeker/internal/app/modules"
	"seeker/pkg/db/postgres"

	"github.com/jackc/pgx"
)

func Start() error {
	config := config.Load()
	server, router := NewHttpServer()

	postgresqlClient := postgres.NewClient(pgx.ConnConfig{
		Host:     config.DB.Host,
		Port:     uint16(config.DB.Port),
		Database: config.DB.Name,
		User:     config.DB.User,
		Password: config.DB.Password,
	})

	userModule := modules.NewUserModule(router, postgresqlClient)
	modules.NewAuthModule(router, userModule)
	modules.NewTalentModule(router, postgresqlClient)

	return server.Start(config.HTTP.Port)
}
