package repositories

import (
	"seeker/internal/domain/entities"

	"github.com/jackc/pgx"
)

type TalentRepository interface {
	CreateOne(tx *pgx.Tx, talent *entities.Talent) error
	FindByID(id string) (entities.Talent, error)
	GetOneByUserId(userId string) (entities.Talent, error)
	CreateProfile(tx *pgx.Tx, profile *entities.TalentProfile) error
	GetProfileByTalentId(talentId string) (entities.Talent, error)
	UpdateProfileByUserId(userId string, profile *entities.TalentProfile) error
}
