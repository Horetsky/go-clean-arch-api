package repositories

import "seeker/internal/domain/entities"

type JobRepository interface {
	CreateJob(job *entities.Job) error
	GetJobByID(id string) (entities.Job, error)
	GetJobByRecruiter(recruiterId string) (entities.Job, error)
	UpdateJob(job *entities.Job) error
	DeleteJob(jobId string) error
}
