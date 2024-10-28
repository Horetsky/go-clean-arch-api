package repositories

import "seeker/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (entities.User, error)
	FindByID(id string) (entities.User, error)
	UpdateByEmail(email string, user *entities.User) error
	DeleteByEmail(email string) error
}
