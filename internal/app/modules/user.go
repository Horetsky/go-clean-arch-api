package modules

import (
	"log"
	"seeker/internal/data/repositories"
	"seeker/internal/domain/storages"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type UserModule struct {
	Repository storages.UserStorage
	Usecase    usecases.UserUsecase
}

func NewUserModule(router *httprouter.Router, pqClient postgres.Client) *UserModule {

	repository := repositories.NewUserRepository(pqClient)
	usecase := usecases.NewUserUsecase(repository)

	handlers.NewUserHandler(usecase).Register(router)

	log.Println("UserModule dependencies initialized")

	return &UserModule{
		Repository: repository,
		Usecase:    usecase,
	}
}
