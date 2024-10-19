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

type TalentUsecase interface {
	CreateProfile(input dto.CreateTalentProfileInput) (entities.Talent, error)
}

type talentUsecase struct {
	talentRepository repositories.TalentRepository
	pqClient         postgres.Client
}

func NewTalentUsecase(
	talentRepository repositories.TalentRepository,
	pqClient postgres.Client,
) TalentUsecase {
	return &talentUsecase{
		talentRepository: talentRepository,
		pqClient:         pqClient,
	}
}

func (u *talentUsecase) CreateProfile(input dto.CreateTalentProfileInput) (entities.Talent, error) {

	dbTalent, err := u.talentRepository.GetOneByUserId(input.UserID)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return entities.Talent{}, err
		}
	}
	if dbTalent.ID != "" {
		return entities.Talent{}, errs.ErrTalentAlreadyExists
	}

	tx, err := u.pqClient.Begin()
	defer func() {
		if err != nil {
			postgres.HandleTxRollback(tx)
		} else {
			postgres.HandleTxCommit(tx)
		}
	}()

	newTalent := entities.Talent{
		UserID: input.UserID,
	}

	err = u.talentRepository.CreateOne(tx, &newTalent)
	if err != nil {
		return entities.Talent{}, errs.ErrFailedToCreateTalent
	}

	newTalentProfile := entities.TalentProfile{
		TalentId:    newTalent.ID,
		Category:    input.Category,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		LinkedInUrl: input.LinkedInUrl,
		ResumeUrl:   input.ResumeUrl,
		Photo:       input.Photo,
	}

	err = u.talentRepository.CreateProfile(tx, &newTalentProfile)
	if err != nil {
		return entities.Talent{}, errs.ErrFailedToCreateTalent
	}

	newTalent.Profile = newTalentProfile

	return newTalent, nil
}
