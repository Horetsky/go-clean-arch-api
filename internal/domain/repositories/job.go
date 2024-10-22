package repositories

import "seeker/internal/domain/entities"

type JobRepository interface {
	Create(job *entities.Job) error
	FindByID(id string) (entities.JobWithRecruiter, error)
	FindAll() ([]entities.JobWithRecruiter, error)
	FindByCategory(category string) ([]entities.JobWithRecruiter, error)
	ApplyJob(talentId, jobId string) error
	FindApplication(talentId, jobId string) (entities.JobApplication, error)
}
