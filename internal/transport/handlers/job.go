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
	talent    usecases.TalentUsecase
	usecase   usecases.JobUsecase
}

func NewJobHandler(
	recruiter usecases.RecruiterUsecase,
	talent usecases.TalentUsecase,
	usecase usecases.JobUsecase,
) handler.Handler {
	return &jobHandler{
		recruiter: recruiter,
		talent:    talent,
		usecase:   usecase,
	}
}

const (
	job      = "/job"
	apply    = job + "/apply"
	listJobs = job + "/list"
)

func (h *jobHandler) Register(router *httprouter.Router) {
	router.POST(job, middlewares.WithAuth(h.handlePostJob))
	router.POST(apply, middlewares.WithAuth(h.handleApplyJob))
	router.GET(listJobs, middlewares.WithAuth(h.handleListJobs))
}

func (h *jobHandler) handlePostJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *jobHandler) handleApplyJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body dto.ApplyJobDTO

	session, err := request.GetSession(r)

	if err != nil {
		response.Error(w, errs.ErrUnauthorized, http.StatusForbidden)
		return
	}

	if err := request.ReadBody(r, &body); err != nil {
		response.Error(w, nil, http.StatusBadRequest)
		return
	}

	body.TalentID = session.User.TalentID

	err = h.talent.ApplyJob(body)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	response.JSON(w, "OK", http.StatusOK)
}

func (h *jobHandler) handleListJobs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()

	category := queryValues.Get("category")

	input := dto.ListJobDTO{
		Category: category,
	}

	list, err := h.usecase.ListJob(input)

	if err != nil {
		response.Error(w, nil, http.StatusBadRequest)
		return
	}

	response.JSON(w, list, http.StatusOK)
}
