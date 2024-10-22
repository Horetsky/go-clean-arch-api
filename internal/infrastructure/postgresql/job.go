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

func (r *jobRepository) Create(job *entities.Job) error {
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

func (r *jobRepository) FindAll() ([]entities.JobWithRecruiter, error) {
	query := `
		SELECT 
		    job.id,
		    job.recruiter_id,
		    job.category,
		    job.title,
		    job.description,
		    job.requirements,
		    recruiter.id,
		    recruiter.user_id,
		    profile.first_name,
			profile.last_name,
			profile.phone,
			profile.linkedIn_url,
			profile.company_name,
			profile.company_website_url
		FROM jobs AS job
		JOIN recruiters AS recruiter ON job.recruiter_id = recruiter.id
		JOIN recruiter_profiles AS profile ON recruiter.id = profile.recruiter_id
	`

	rows, err := r.client.Query(query)

	if err != nil {
		return nil, postgres.NewError(err)
	}

	defer rows.Close()

	var jobs []entities.JobWithRecruiter

	for rows.Next() {
		var job entities.JobWithRecruiter
		job.Recruiter = entities.Recruiter{}
		job.Recruiter.Profile = &entities.RecruiterProfile{}

		if err := rows.Scan(
			&job.ID,
			&job.RecruiterID,
			&job.Category,
			&job.Title,
			&job.Description,
			&job.Requirements,
			&job.Recruiter.ID,
			&job.Recruiter.UserID,
			&job.Recruiter.Profile.FirstName,
			&job.Recruiter.Profile.LastName,
			&job.Recruiter.Profile.Phone,
			&job.Recruiter.Profile.LinkedInUrl,
			&job.Recruiter.Profile.CompanyName,
			&job.Recruiter.Profile.CompanyWebsiteUrl,
		); err != nil {
			return nil, postgres.NewError(err)
		}

		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, postgres.NewError(err)
	}

	return jobs, nil
}

func (r *jobRepository) FindByID(id string) (entities.JobWithRecruiter, error) {
	query := `
		SELECT 
		    job.id,
		    job.recruiter_id,
		    job.category,
		    job.title,
		    job.description,
		    job.requirements,
		    recruiter.id,
		    recruiter.user_id,
		    profile.first_name,
			profile.last_name,
			profile.phone,
			profile.linkedIn_url,
			profile.company_name,
			profile.company_website_url
		FROM jobs AS job
		JOIN recruiters AS recruiter ON job.recruiter_id = recruiter.id
		JOIN recruiter_profiles AS profile ON recruiter.id = profile.recruiter_id
		WHERE job.id = $1
	`

	row := r.client.QueryRow(query, id)

	var job entities.JobWithRecruiter
	job.Recruiter = entities.Recruiter{}
	job.Recruiter.Profile = &entities.RecruiterProfile{}

	if err := row.Scan(
		&job.ID,
		&job.RecruiterID,
		&job.Category,
		&job.Title,
		&job.Description,
		&job.Requirements,
		&job.Recruiter.ID,
		&job.Recruiter.UserID,
		&job.Recruiter.Profile.FirstName,
		&job.Recruiter.Profile.LastName,
		&job.Recruiter.Profile.Phone,
		&job.Recruiter.Profile.LinkedInUrl,
		&job.Recruiter.Profile.CompanyName,
		&job.Recruiter.Profile.CompanyWebsiteUrl,
	); err != nil {
		return job, postgres.NewError(err)
	}

	return job, nil
}

func (r *jobRepository) FindByCategory(category string) ([]entities.JobWithRecruiter, error) {
	query := `
			SELECT 
		    job.id,
		    job.recruiter_id,
		    job.category,
		    job.title,
		    job.description,
		    job.requirements,
		    recruiter.id,
		    recruiter.user_id,
		    profile.first_name,
			profile.last_name,
			profile.phone,
			profile.linkedIn_url,
			profile.company_name,
			profile.company_website_url
		FROM jobs AS job
		JOIN recruiters AS recruiter ON job.recruiter_id = recruiter.id
		JOIN recruiter_profiles AS profile ON recruiter.id = profile.recruiter_id
		WHERE job.category = $1
	`

	rows, err := r.client.Query(query, category)

	if err != nil {
		return nil, postgres.NewError(err)
	}

	defer rows.Close()

	var jobs []entities.JobWithRecruiter

	for rows.Next() {
		var job entities.JobWithRecruiter
		job.Recruiter = entities.Recruiter{}
		job.Recruiter.Profile = &entities.RecruiterProfile{}

		if err := rows.Scan(
			&job.ID,
			&job.RecruiterID,
			&job.Category,
			&job.Title,
			&job.Description,
			&job.Requirements,
			&job.Recruiter.ID,
			&job.Recruiter.UserID,
			&job.Recruiter.Profile.FirstName,
			&job.Recruiter.Profile.LastName,
			&job.Recruiter.Profile.Phone,
			&job.Recruiter.Profile.LinkedInUrl,
			&job.Recruiter.Profile.CompanyName,
			&job.Recruiter.Profile.CompanyWebsiteUrl,
		); err != nil {
			return nil, postgres.NewError(err)
		}

		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, postgres.NewError(err)
	}

	return jobs, nil
}

func (r *jobRepository) ApplyJob(talentId, jobId string) error {
	query := `
		INSERT INTO job_applications (talent_id, job_id)
		VALUES ($1, $2)
	`

	row := r.client.QueryRow(query, talentId, jobId)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *jobRepository) FindApplication(talentId, jobId string) (entities.JobApplication, error) {
	query := `
		SELECT talent_id, job_id FROM job_applications
		WHERE talent_id = $1 AND job_id = $2
	`

	row := r.client.QueryRow(query, talentId, jobId)

	var application entities.JobApplication

	if err := row.Scan(
		&application.TalentID,
		&application.JobID,
	); err != nil {
		return entities.JobApplication{}, postgres.NewError(err)
	}

	return application, nil
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
