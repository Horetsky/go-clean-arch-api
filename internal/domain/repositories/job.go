package repositories

import "seeker/internal/domain/entities"

type JobRepository interface {
	CreateJob(job *entities.Job) error
	FindByID(id string) (entities.JobWithRecruiter, error)
	FindByRecruiter(recruiterId string) (entities.Job, error)
	ApplyJob(talentId, jobId string) error
	FindApplication(talentId, jobId string) (entities.JobApplication, error)
	UpdateJob(job *entities.Job) error
	DeleteJob(jobId string) error
}
