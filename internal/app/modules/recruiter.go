package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type RecruiterModule struct {
	Usecase usecases.RecruiterUsecase
}

func NewRecruiterModule(router *httprouter.Router, pqClient postgres.Client, authUsecase usecases.AuthUsecase) usecases.RecruiterUsecase {
	repository := postgresql.NewRecruiterRepository(pqClient)
	jobRepository := postgresql.NewJobRepository(pqClient)

	usecase := usecases.NewRecruiterUsecase(repository, jobRepository, pqClient)

	handlers.NewRecruiterHandler(usecase, authUsecase).Register(router)

	log.Println("RecruiterModule dependencies initialized")

	return usecase
}
