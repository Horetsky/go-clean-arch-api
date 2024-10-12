package modules

import (
	"log"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/handlers"

	"github.com/julienschmidt/httprouter"
)

type AuthModule struct {
	Usecase usecases.AuthUsecase
}

func NewAuthModule(router *httprouter.Router, userModule *UserModule) *AuthModule {

	usecase := usecases.NewAuthUsecase(userModule.Repository)

	handlers.NewAuthHandler(usecase).Register(router)

	log.Println("AuthModule dependencies initialized")

	return &AuthModule{
		Usecase: usecase,
	}
}
