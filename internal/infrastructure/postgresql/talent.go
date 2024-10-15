package postgresql

import (
	"fmt"
	"seeker/internal/domain/entities"
	"seeker/internal/domain/repositories"
	"seeker/pkg/db/postgres"
	"seeker/pkg/utils/str"

	"github.com/jackc/pgx"
)

type talentRepository struct {
	client postgres.Client
}

func NewTalentRepository(client postgres.Client) repositories.TalentRepository {
	return &talentRepository{
		client: client,
	}
}

func (r *talentRepository) CreateOne(tx *pgx.Tx, talent *entities.Talent) error {
	query := `
		INSERT INTO "talents" (user_id)
		VALUES ($1)
		RETURNING id
	`

	row := tx.QueryRow(query, talent.UserID)

	if err := row.Scan(&talent.ID); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *talentRepository) GetOneByUserId(userId string) (entities.Talent, error) {
	query := `
		SELECT id, user_id FROM talents
		WHERE user_id = $1
	`

	row := r.client.QueryRow(query, userId)

	var talent entities.Talent

	if err := row.Scan(&talent.ID, &talent.UserID); err != nil {
		return talent, postgres.NewError(err)
	}

	return talent, nil
}

func (r *talentRepository) CreateProfile(tx *pgx.Tx, profile *entities.TalentProfile) error {
	query := `
		INSERT INTO "talent_profiles" (talent_id, category, first_name, last_name, phone, linkedin_url, resume_url, photo)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	row := tx.QueryRow(
		query,
		profile.TalentId,
		profile.Category,
		profile.FirstName,
		profile.LastName,
		profile.Phone,
		profile.LinkedInUrl,
		profile.ResumeUrl,
		profile.Photo,
	)

	if err := row.Scan(&profile.ID); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *talentRepository) GetProfileByTalentId(talentId string) (entities.Talent, error) {
	query := `
		SELECT id, 
		       user_id,
		       profile.id,
		       profile.talent_id,
		       profile.category,
		       profile.first_name,
		       profile.last_name,
		       profile.phone,
		       profile.linkedin_url,
		       profile.resume_url,
		       profile.photo
		FROM talents AS talent
		    JOIN talent_profiles AS profile ON talent.id = profile.talent_id
		WHERE id = $1
	`

	row := r.client.QueryRow(query, talentId)

	var talent entities.Talent

	if err := row.Scan(
		&talent.ID,
		&talent.UserID,
		&talent.Profile.TalentId,
		&talent.Profile.Category,
		&talent.Profile.FirstName,
		&talent.Profile.LastName,
		&talent.Profile.Phone,
		&talent.Profile.LinkedInUrl,
		&talent.Profile.ResumeUrl,
		&talent.Profile.Photo,
	); err != nil {
		return talent, postgres.NewError(err)
	}

	return talent, nil
}

func (r *talentRepository) UpdateProfileByUserId(userId string, profile *entities.TalentProfile) error {
	query := `
		UPDATE "talent_profiles" SET
	`

	str.ForEach(profile, func(key string, val any) {
		switch key {
		case "CategoryId":
			query += fmt.Sprintf(" category_id = '%s',", val)
			break
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
		case "ResumeUrl":
			query += fmt.Sprintf(" resume_url = '%s',", val)
			break
		case "Photo":
			query += fmt.Sprintf(" photo = '%s',", val)
			break
		}
	})

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE user_id = '%s'", userId)

	row := r.client.QueryRow(query)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}
