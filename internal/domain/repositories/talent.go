package repositories

import (
	"seeker/internal/domain/entities"

	"github.com/jackc/pgx"
)

type TalentRepository interface {
	Create(tx *pgx.Tx, talent *entities.Talent) error
	FindByID(id string) (entities.Talent, error)
	FindByUserID(userId string) (entities.Talent, error)
	CreateProfile(tx *pgx.Tx, profile *entities.TalentProfile) error
	FindProfileByTalentID(talentId string) (entities.Talent, error)
	UpdateProfile(userId string, profile *entities.TalentProfile) error
}
