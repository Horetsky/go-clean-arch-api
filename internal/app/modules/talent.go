package modules

import (
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

type TalentModule struct {
	Repository repositories.TalentRepository
	Usecase    usecases.TalentUsecase
}

func NewTalentModule(router *httprouter.Router, pqClient postgres.Client) *TalentModule {

	repository := postgresql.NewTalentRepository(pqClient)
	usecase := usecases.NewTalentUsecase(repository, pqClient)

	handlers.NewTalentHandler(usecase).Register(router)

	return &TalentModule{
		Repository: repository,
		Usecase:    usecase,
	}
}
