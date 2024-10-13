package repositories

import "seeker/internal/domain/entities"

type UserRepository interface {
	GetByEmail(email string) (entities.User, error)
	GetByID(id string) (entities.User, error)
	CreateOne(user *entities.User) error
	UpdateByEmail(email string, user *entities.User) error
}
