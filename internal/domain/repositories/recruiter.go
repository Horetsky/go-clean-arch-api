package repositories

import (
	"seeker/internal/domain/entities"

	"github.com/jackc/pgx"
)

type RecruiterRepository interface {
	Create(tx *pgx.Tx, recruiter *entities.Recruiter) error
	FindAll() ([]entities.Recruiter, error)
	FindByUserID(userId string) (entities.Recruiter, error)
	CreateProfile(tx *pgx.Tx, profile *entities.RecruiterProfile) error
	FindProfileByRecruiterID(recruiterId string) (entities.Recruiter, error)
	UpdateProfile(userId string, profile *entities.RecruiterProfile) error
}
