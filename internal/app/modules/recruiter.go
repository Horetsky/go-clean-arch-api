package modules

import (
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type RecruiterModule struct {
	Repository repositories.RecruiterRepository
	Usecase    usecases.RecruiterUsecase
}

func NewRecruiterModule(router *httprouter.Router, pqClient postgres.Client) *RecruiterModule {

	repository := postgresql.NewRecruiterRepository(pqClient)
	usecase := usecases.NewRecruiterUsecase(repository, pqClient)

	handlers.NewRecruiterHandler(usecase).Register(router)

	return &RecruiterModule{
		Repository: repository,
		Usecase:    usecase,
	}
}
