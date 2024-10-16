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
}

type recruiterUsecase struct {
	recruiterRepository repositories.RecruiterRepository
	pqClient            postgres.Client
}

func NewRecruiterUsecase(
	recruiterRepository repositories.RecruiterRepository,
	pqClient postgres.Client,
) RecruiterUsecase {
	return &recruiterUsecase{
		recruiterRepository: recruiterRepository,
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

	newRecruiter.Profile = newRecruiterProfile

	return newRecruiter, nil
}
