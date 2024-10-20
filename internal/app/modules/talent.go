package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/emailSender"
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
	userRepository := postgresql.NewUserRepository(pqClient)
	recruiterRepository := postgresql.NewRecruiterRepository(pqClient)
	jobRepository := postgresql.NewJobRepository(pqClient)
	emailService := emailSender.NewSmtpSender()

	usecase := usecases.NewTalentUsecase(
		repository,
		userRepository,
		recruiterRepository,
		jobRepository,
		emailService,
		pqClient,
	)

	handlers.NewTalentHandler(usecase, authUsecase).Register(router)

	log.Println("TalentModule dependencies initialized")

	return usecase
}
