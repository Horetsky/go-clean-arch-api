package handlers

import (
	"net/http"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/middlewares"
	"seeker/pkg/handler"
	"seeker/pkg/handler/response"

	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) handler.Handler {
	return &userHandler{
		usecase: usecase,
	}
}

const (
	users = "/users"
)

func (h *userHandler) Register(router *httprouter.Router) {
	router.GET(users, middlewares.WithAuth(h.handleFindUser))
}

func (h *userHandler) handleFindUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	users, err := h.usecase.FindUser(queryValues)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.JSON(w, users, http.StatusOK)
}
