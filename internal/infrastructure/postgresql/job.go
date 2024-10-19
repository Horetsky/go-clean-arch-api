package postgresql

import (
	"fmt"
	"seeker/internal/domain/entities"
	"seeker/internal/domain/repositories"
	"seeker/pkg/db/postgres"
	"seeker/pkg/utils/str"
)

type jobRepository struct {
	client postgres.Client
}

func NewJobRepository(
	client postgres.Client,
) repositories.JobRepository {
	return &jobRepository{
		client: client,
	}
}

func (r *jobRepository) CreateJob(job *entities.Job) error {
	query := `
		INSERT INTO jobs (recruiter_id, category, title, description, requirements)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	row := r.client.QueryRow(query, job.RecruiterID, job.Category, job.Title, job.Description, job.Requirements)

	if err := row.Scan(&job.ID); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *jobRepository) GetJobByID(id string) (entities.Job, error) {
	//TODO implement me
	panic("implement me")
}

func (r *jobRepository) GetJobByRecruiter(recruiterId string) (entities.Job, error) {
	//TODO implement me
	panic("implement me")
}

func (r *jobRepository) UpdateJob(job *entities.Job) error {
	query := `
		UPDATE jobs SET
	`

	str.ForEach(job, func(key string, val any) {
		switch key {
		case "Category":
			query += fmt.Sprintf(" category = '%s',", val)
			break
		case "Title":
			query += fmt.Sprintf(" title = '%s',", val)
			break
		case "Description":
			query += fmt.Sprintf(" description = '%s',", val)
			break
		case "Requirements":
			query += fmt.Sprintf(" requirements = '%s',", val)
			break
		}
	})

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id = '%s'", job.ID)

	row := r.client.QueryRow(query)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *jobRepository) DeleteJob(jobId string) error {
	query := `
		DELETE FROM jobs
		WHERE id = $1
	`

	row := r.client.QueryRow(query, jobId)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}
