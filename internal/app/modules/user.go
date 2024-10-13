package modules

import (
	"log"
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type UserModule struct {
	Repository repositories.UserRepository
	Usecase    usecases.UserUsecase
}

func NewUserModule(router *httprouter.Router, pqClient postgres.Client) *UserModule {

	repository := postgresql.NewUserRepository(pqClient)
	usecase := usecases.NewUserUsecase(repository)

	handlers.NewUserHandler(usecase).Register(router)

	log.Println("UserModule dependencies initialized")

	return &UserModule{
		Repository: repository,
		Usecase:    usecase,
	}
}
