package modules

import (
	"log"
	"seeker/internal/domain/services"
	"seeker/internal/domain/usecases"
	"seeker/internal/infrastructure/emailSender"
	"seeker/internal/transport/handlers"

	"github.com/julienschmidt/httprouter"
)

type AuthModule struct {
	Usecase usecases.AuthUsecase
}

func NewAuthModule(router *httprouter.Router, userModule *UserModule) *AuthModule {
	jwtService := services.NewJWTService()
	emailSender := emailSender.NewSmtpSender()

	usecase := usecases.NewAuthUsecase(userModule.Repository, jwtService, emailSender)

	handlers.NewAuthHandler(usecase).Register(router)

	log.Println("AuthModule dependencies initialized")

	return &AuthModule{
		Usecase: usecase,
	}
}
