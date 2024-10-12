package storages

import "seeker/internal/domain/entities"

type UserStorage interface {
	GetByEmail(email string) (entities.User, error)
	GetByID(id string) (entities.User, error)
	CreateOne(user *entities.User) error
}
