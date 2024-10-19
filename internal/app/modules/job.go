package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

func NewJobModule(router *httprouter.Router, _ postgres.Client, recruiterUsecase usecases.RecruiterUsecase) usecases.JobUsecase {

	usecase := usecases.NewJobUsecase()
	handlers.NewJobHandler(recruiterUsecase).Register(router)

	log.Println("JobModule dependencies initialized")

	return usecase
}
