package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type TalentModule struct {
	Usecase usecases.TalentUsecase
}

func NewTalentModule(router *httprouter.Router, pqClient postgres.Client, authUsecase usecases.AuthUsecase) usecases.TalentUsecase {

	repository := postgresql.NewTalentRepository(pqClient)
	usecase := usecases.NewTalentUsecase(repository, pqClient)

	handlers.NewTalentHandler(usecase, authUsecase).Register(router)

	log.Println("TalentModule dependencies initialized")

	return usecase
}
