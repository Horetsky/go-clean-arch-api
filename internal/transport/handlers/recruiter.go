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

type recruiterHandler struct {
	usecase     usecases.RecruiterUsecase
	authUsecase usecases.AuthUsecase
}

func NewRecruiterHandler(
	usecase usecases.RecruiterUsecase,
	authUsecase usecases.AuthUsecase,
) handler.Handler {
	return &recruiterHandler{
		usecase:     usecase,
		authUsecase: authUsecase,
	}
}

const (
	recruiter = "/recruiter"
)

func (h *recruiterHandler) Register(router *httprouter.Router) {
	router.POST(recruiter, middlewares.WithAuth(h.handleCreateRecruiterProfile))
}

func (h *recruiterHandler) handleCreateRecruiterProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := request.GetSession(r)

	if err != nil {
		response.Error(w, errs.ErrUnauthorized, http.StatusForbidden)
		return
	}

	body := dto.CreateRecruiterProfileInput{
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
		Type:          session.User.Type,
		Picture:       session.User.Picture,
		EmailVerified: session.User.EmailVerified,
		Recruiter:     &profile,
	}

	tokens, _, err := h.authUsecase.GenerateSession(user)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.PrivateCookie(w, dto.AccessTokenCookieKey, tokens.AccessToken)
	response.PrivateCookie(w, dto.RefreshTokenCookieKey, tokens.RefreshToken)
	response.JSON(w, profile, http.StatusCreated)
}
