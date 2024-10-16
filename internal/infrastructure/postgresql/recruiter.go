package postgresql

import (
	"fmt"
	"seeker/internal/domain/entities"
	"seeker/internal/domain/repositories"
	"seeker/pkg/db/postgres"
	"seeker/pkg/utils/str"

	"github.com/jackc/pgx"
)

type recruiterRepository struct {
	client postgres.Client
}

func NewRecruiterRepository(client postgres.Client) repositories.RecruiterRepository {
	return &recruiterRepository{
		client: client,
	}
}

func (r recruiterRepository) CreateOne(tx *pgx.Tx, recruiter *entities.Recruiter) error {
	query := `
		INSERT INTO recruiters (user_id)
		VALUES ($1)
		RETURNING id
	`

	row := tx.QueryRow(query, recruiter.UserID)

	if err := row.Scan(&recruiter.ID); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r recruiterRepository) GetOneByUserId(userId string) (entities.Recruiter, error) {
	query := `
		SELECT id, user_id FROM recruiters
		WHERE user_id = $1
	`

	row := r.client.QueryRow(query, userId)

	var recruiter entities.Recruiter

	if err := row.Scan(&recruiter.ID, &recruiter.UserID); err != nil {
		return recruiter, postgres.NewError(err)
	}

	return recruiter, nil
}

func (r recruiterRepository) CreateProfile(tx *pgx.Tx, profile *entities.RecruiterProfile) error {
	query := `
		INSERT INTO recruiter_profiles (recruiter_id, first_name, last_name, phone, linkedIn_url, company_name, company_website_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	row := tx.QueryRow(
		query,
		profile.RecruiterID,
		profile.FirstName,
		profile.LastName,
		profile.Phone,
		profile.LinkedInUrl,
		profile.CompanyName,
		profile.CompanyWebsiteUrl,
	)

	if err := row.Scan(&profile.ID); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r recruiterRepository) GetProfileByRecruiterId(recruiterId string) (entities.Recruiter, error) {
	query := `
		SELECT id, 
		       user_id,
		       profile.id,
		       profile.recruiter_id,
		       profile.first_name,
		       profile.last_name,
		       profile.phone,
		       profile.linkedIn_url,
		       profile.company_name,
		       profile.company_website_url
		FROM recruiters AS recruiter
		    JOIN recruiter_profiles AS profile ON recruiter.id = profile.recruiter_id
		WHERE id = $1
	`

	row := r.client.QueryRow(query, recruiterId)

	var recruiter entities.Recruiter

	if err := row.Scan(
		&recruiter.ID,
		&recruiter.UserID,
		&recruiter.Profile.RecruiterID,
		&recruiter.Profile.FirstName,
		&recruiter.Profile.LastName,
		&recruiter.Profile.Phone,
		&recruiter.Profile.LinkedInUrl,
		&recruiter.Profile.CompanyName,
		&recruiter.Profile.CompanyWebsiteUrl,
	); err != nil {
		return recruiter, postgres.NewError(err)
	}

	return recruiter, nil
}

func (r recruiterRepository) UpdateProfileByUserId(recruiterId string, profile *entities.RecruiterProfile) error {
	query := `
		UPDATE recruiter_profiles SET
	`

	str.ForEach(profile, func(key string, val any) {
		switch key {
		case "FirstName":
			query += fmt.Sprintf(" first_name = '%s',", val)
			break
		case "LastName":
			query += fmt.Sprintf(" last_name = '%s',", val)
			break
		case "Phone":
			query += fmt.Sprintf(" phone = '%s',", val)
			break
		case "LinkedInUrl":
			query += fmt.Sprintf(" linkedIn_url = '%s',", val)
			break
		case "CompanyName":
			query += fmt.Sprintf(" company_name = '%s',", val)
			break
		case "CompanyWebsiteUrl":
			query += fmt.Sprintf(" company_website_url = '%s',", val)
			break
		}
	})

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE recruiter_id = '%s'", recruiterId)

	row := r.client.QueryRow(query)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}
