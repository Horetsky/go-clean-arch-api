package repositories

import (
	"seeker/internal/domain/entities"

	"github.com/jackc/pgx"
)

type RecruiterRepository interface {
	CreateOne(tx *pgx.Tx, recruiter *entities.Recruiter) error
	GetOneByUserId(userId string) (entities.Recruiter, error)
	CreateProfile(tx *pgx.Tx, profile *entities.RecruiterProfile) error
	GetProfileByRecruiterId(recruiterId string) (entities.Recruiter, error)
	UpdateProfileByUserId(userId string, profile *entities.RecruiterProfile) error
}
