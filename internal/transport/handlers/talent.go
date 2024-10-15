package handlers

import (
	"net/http"
	"seeker/internal/domain/dto"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/middlewares"
	"seeker/pkg/handler"
	"seeker/pkg/handler/request"
	"seeker/pkg/handler/response"

	"github.com/julienschmidt/httprouter"
)

type talentHandler struct {
	usecase usecases.TalentUsecase
}

func NewTalentHandler(
	usecase usecases.TalentUsecase,
) handler.Handler {
	return &talentHandler{
		usecase: usecase,
	}
}

const (
	talent = "/talent"
)

func (h *talentHandler) Register(router *httprouter.Router) {
	router.POST(talent, middlewares.WithAuth(h.handleCreateTalentProfile))
}

func (h *talentHandler) handleCreateTalentProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session, err := request.GetSession(r)

	if err != nil {
		response.Error(w, errs.ErrUnauthorized, http.StatusForbidden)
		return
	}

	body := dto.CreateTalentProfileInput{
		UserID: session.ID,
	}

	if err := request.ReadBody(r, &body); err != nil {
		response.Error(w, nil, http.StatusBadRequest)
		return
	}

	profile, err := h.usecase.CreateProfile(body)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.JSON(w, profile, http.StatusCreated)
}
