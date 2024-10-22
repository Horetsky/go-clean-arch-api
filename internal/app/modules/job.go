package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

func NewJobModule(
	router *httprouter.Router,
	pqClient postgres.Client,
	recruiterUsecase usecases.RecruiterUsecase,
	talentUsecase usecases.TalentUsecase,
) usecases.JobUsecase {

	jobRepository := postgresql.NewJobRepository(pqClient)
	usecase := usecases.NewJobUsecase(jobRepository)
	handlers.NewJobHandler(
		recruiterUsecase,
		talentUsecase,
		usecase,
	).Register(router)

	log.Println("JobModule dependencies initialized")

	return usecase
}
