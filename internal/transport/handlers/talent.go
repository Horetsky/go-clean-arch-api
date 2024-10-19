package handlers

import (
	"net/http"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/usecases"
	"seeker/internal/transport/middlewares"
	"seeker/pkg/handler"
	"seeker/pkg/handler/request"
	"seeker/pkg/handler/response"

	"github.com/julienschmidt/httprouter"
)

type talentHandler struct {
	usecase     usecases.TalentUsecase
	authUsecase usecases.AuthUsecase
}

func NewTalentHandler(
	usecase usecases.TalentUsecase,
	authUsecase usecases.AuthUsecase,
) handler.Handler {
	return &talentHandler{
		usecase:     usecase,
		authUsecase: authUsecase,
	}
}

const (
	talent = "/talent"
)

func (h *talentHandler) Register(router *httprouter.Router) {
	router.POST(talent, middlewares.WithAuth(h.handleCreateTalentProfile))
}

func (h *talentHandler) handleCreateTalentProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := request.GetSession(r)

	if err != nil {
		response.Error(w, errs.ErrUnauthorized, http.StatusForbidden)
		return
	}

	body := dto.CreateTalentProfileInput{
		UserID: session.User.ID,
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

	user := &entities.User{
		ID:            session.User.ID,
		Email:         session.User.Email,
		Picture:       session.User.Picture,
		EmailVerified: session.User.EmailVerified,
		Talent:        &profile,
	}

	tokens, s, err := h.authUsecase.GenerateSession(user)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.PrivateCookie(w, dto.AccessTokenCookieKey, tokens.AccessToken)
	response.PrivateCookie(w, dto.RefreshTokenCookieKey, tokens.RefreshToken)
	response.JSON(w, s, http.StatusCreated)
}
