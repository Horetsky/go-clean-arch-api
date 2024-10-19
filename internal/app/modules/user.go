package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type UserModule struct {
	Usecase usecases.UserUsecase
}

func NewUserModule(router *httprouter.Router, pqClient postgres.Client) usecases.UserUsecase {

	repository := postgresql.NewUserRepository(pqClient)
	usecase := usecases.NewUserUsecase(repository)

	handlers.NewUserHandler(usecase).Register(router)

	log.Println("UserModule dependencies initialized")

	return usecase
}
