package usecases

import (
	"seeker/internal/domain/dto"
	"seeker/internal/domain/entities"
	"seeker/internal/domain/repositories"
)

type JobUsecase interface {
	ListJob(input dto.ListJobDTO) ([]entities.JobWithRecruiter, error)
}

type jobUsecase struct {
	jobRepository repositories.JobRepository
}

func NewJobUsecase(
	jobRepository repositories.JobRepository,
) JobUsecase {
	return &jobUsecase{
		jobRepository: jobRepository,
	}
}

func (u *jobUsecase) ListJob(input dto.ListJobDTO) ([]entities.JobWithRecruiter, error) {
	if input.Category != "" {
		return u.listByCategory(input.Category)
	}

	return u.listAll()
}

func (u *jobUsecase) listByCategory(category string) ([]entities.JobWithRecruiter, error) {
	list, err := u.jobRepository.FindByCategory(category)
	if err != nil {
		return []entities.JobWithRecruiter{}, nil
	}

	return list, nil
}

func (u *jobUsecase) listAll() ([]entities.JobWithRecruiter, error) {
	list, err := u.jobRepository.FindAll()
	if err != nil {
		return []entities.JobWithRecruiter{}, nil
	}

	return list, nil
}
