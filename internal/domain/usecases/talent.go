package usecases

import (
	"errors"
	"fmt"
	"log"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	errs "seeker/internal/domain/errors"
	"seeker/internal/domain/repositories"
	"seeker/internal/domain/services"
	"seeker/pkg/db/postgres"

	"github.com/jackc/pgx"
)

type TalentUsecase interface {
	CreateProfile(input dto.CreateTalentProfileInput) (entities.Talent, error)
	ApplyJob(input dto.ApplyJobDTO) error
}

type talentUsecase struct {
	talentRepository    repositories.TalentRepository
	userRepository      repositories.UserRepository
	jobRepository       repositories.JobRepository
	recruiterRepository repositories.RecruiterRepository
	emailService        services.EmailService
	pqClient            postgres.Client
}

func NewTalentUsecase(
	talentRepository repositories.TalentRepository,
	userRepository repositories.UserRepository,
	recruiterRepository repositories.RecruiterRepository,
	jobRepository repositories.JobRepository,
	emailService services.EmailService,
	pqClient postgres.Client,
) TalentUsecase {
	return &talentUsecase{
		talentRepository:    talentRepository,
		recruiterRepository: recruiterRepository,
		userRepository:      userRepository,
		jobRepository:       jobRepository,
		emailService:        emailService,
		pqClient:            pqClient,
	}
}

func (u *talentUsecase) CreateProfile(input dto.CreateTalentProfileInput) (entities.Talent, error) {
	dbTalent, err := u.talentRepository.FindByUserID(input.UserID)

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

	err = u.talentRepository.Create(tx, &newTalent)
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

	newTalent.Profile = &newTalentProfile

	return newTalent, nil
}

func (u *talentUsecase) ApplyJob(input dto.ApplyJobDTO) error {

	dbApplication, err := u.jobRepository.FindApplication(input.TalentID, input.JobID)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return errs.ErrApplicationAlreadyExists
		}
	}

	if dbApplication.JobID != "" {
		return errs.ErrApplicationAlreadyExists
	}

	err = u.jobRepository.ApplyJob(input.TalentID, input.JobID)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return errs.ErrFailedToApplyJob
		}
	}

	go u.sendApplicationEmailNotification(input.TalentID, input.JobID)

	return nil
}

func (u *talentUsecase) sendApplicationEmailNotification(talentId string, jobId string) {
	job, err := u.jobRepository.FindByID(jobId)

	if err != nil {
		log.Println(errs.ErrFailedToSendApplicationEmail)
		return
	}

	talent, err := u.talentRepository.FindByID(talentId)

	if err != nil {
		log.Println(errs.ErrFailedToSendApplicationEmail)
		return
	}

	dbUser, err := u.userRepository.FindByID(job.Recruiter.UserID)

	if err != nil {
		log.Println(errs.ErrFailedToSendApplicationEmail)
		return
	}

	emailInput := dto.SendJobApplicationEmailDTO{
		JobTitle:      job.Title,
		RecruiterName: fmt.Sprintf("%s %s", job.Recruiter.Profile.FirstName, job.Recruiter.Profile.LastName),
		ApplicantName: fmt.Sprintf("%s %s", talent.Profile.FirstName, talent.Profile.LastName),
		CompanyName:   job.Recruiter.Profile.CompanyName,
	}

	err = u.emailService.SendJobApplicationEmail(dbUser.Email, emailInput)

	if err != nil {
		log.Println(errs.ErrFailedToSendApplicationEmail)
		return
	}
}
