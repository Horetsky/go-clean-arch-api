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

type jobHandler struct {
	recruiter usecases.RecruiterUsecase
}

func NewJobHandler(
	recruiter usecases.RecruiterUsecase,
) handler.Handler {
	return &jobHandler{
		recruiter: recruiter,
	}
}

const (
	job = "/job"
)

func (h *jobHandler) Register(router *httprouter.Router) {
	router.POST(job, middlewares.WithAuth(h.handlePostJob))
}

func (h *jobHandler) handlePostJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var body dto.PostJobDTO

	session, err := request.GetSession(r)

	if err != nil {
		response.Error(w, errs.ErrUnauthorized, http.StatusForbidden)
		return
	}

	if err := request.ReadBody(r, &body); err != nil {
		response.Error(w, nil, http.StatusBadRequest)
		return
	}

	body.RecruiterID = session.User.RecruiterID

	newJob, err := h.recruiter.PostJob(body)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.JSON(w, newJob, http.StatusOK)
}
