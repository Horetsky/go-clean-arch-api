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
	jwtService := services.NewJWTService()
	sender := emailSender.NewSmtpSender()

	usecase := usecases.NewAuthUsecase(userRepository, jwtService, sender)

	handlers.NewAuthHandler(usecase).Register(router)

	log.Println("AuthModule dependencies initialized")

	return usecase
}
