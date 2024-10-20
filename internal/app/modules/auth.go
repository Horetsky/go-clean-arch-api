package modules

import (
	"log"
	"seeker/internal/domain/services"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/emailSender"
	"seeker/internal/infrastructure/postgresql"
	"seeker/internal/transport/handlers"
	"seeker/pkg/db/postgres"

	"github.com/julienschmidt/httprouter"
)

func NewAuthModule(router *httprouter.Router, pqClient postgres.Client) usecases.AuthUsecase {

	userRepository := postgresql.NewUserRepository(pqClient)
	talentRepository := postgresql.NewTalentRepository(pqClient)
	recruiterRepository := postgresql.NewRecruiterRepository(pqClient)
	jwtService := services.NewJWTService()
	emailService := emailSender.NewSmtpSender()

	usecase := usecases.NewAuthUsecase(
		userRepository,
		talentRepository,
		recruiterRepository,
		jwtService,
		emailService,
	)

	handlers.NewAuthHandler(usecase).Register(router)

	log.Println("AuthModule dependencies initialized")

	return usecase
}
