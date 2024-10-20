package usecases

import (
	"errors"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/repositories"
	"seeker/pkg/db/postgres"

	"github.com/jackc/pgx"
)

type RecruiterUsecase interface {
	CreateProfile(input dto.CreateRecruiterProfileInput) (entities.Recruiter, error)
	PostJob(input dto.PostJobDTO) (entities.Job, error)
}

type recruiterUsecase struct {
	recruiterRepository repositories.RecruiterRepository
	jobRepository       repositories.JobRepository
	pqClient            postgres.Client
}

func NewRecruiterUsecase(
	recruiterRepository repositories.RecruiterRepository,
	jobRepository repositories.JobRepository,
	pqClient postgres.Client,
) RecruiterUsecase {
	return &recruiterUsecase{
		recruiterRepository: recruiterRepository,
		jobRepository:       jobRepository,
		pqClient:            pqClient,
	}
}

func (u *recruiterUsecase) CreateProfile(input dto.CreateRecruiterProfileInput) (entities.Recruiter, error) {
	dbRecruiter, err := u.recruiterRepository.GetOneByUserId(input.UserID)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return entities.Recruiter{}, err
		}
	}
	if dbRecruiter.ID != "" {
		return entities.Recruiter{}, errs.ErrRecruiterAlreadyExists
	}

	tx, err := u.pqClient.Begin()
	defer func() {
		if err != nil {
			postgres.HandleTxRollback(tx)
		} else {
			postgres.HandleTxCommit(tx)
		}
	}()

	newRecruiter := entities.Recruiter{
		UserID: input.UserID,
	}

	err = u.recruiterRepository.CreateOne(tx, &newRecruiter)
	if err != nil {
		return entities.Recruiter{}, errs.ErrFailedToCreateRecruiter
	}

	newRecruiterProfile := entities.RecruiterProfile{
		RecruiterID:       newRecruiter.ID,
		FirstName:         input.FirstName,
		LastName:          input.LastName,
		Phone:             input.Phone,
		LinkedInUrl:       input.LinkedInUrl,
		CompanyName:       input.CompanyName,
		CompanyWebsiteUrl: input.CompanyWebsiteUrl,
	}

	err = u.recruiterRepository.CreateProfile(tx, &newRecruiterProfile)
	if err != nil {
		return entities.Recruiter{}, errs.ErrFailedToCreateRecruiter
	}

	newRecruiter.Profile = &newRecruiterProfile

	return newRecruiter, nil
}

func (u *recruiterUsecase) PostJob(input dto.PostJobDTO) (entities.Job, error) {
	newJob := entities.Job{
		RecruiterID:  input.RecruiterID,
		Category:     input.Category,
		Title:        input.Title,
		Description:  input.Description,
		Requirements: input.Requirements,
	}

	err := u.jobRepository.CreateJob(&newJob)

	if err != nil {
		return entities.Job{}, errs.ErrFailedToPostJob
	}

	return newJob, nil
}
