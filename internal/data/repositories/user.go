package repositories

import (
	"seeker/internal/domain/entities"
	"seeker/internal/domain/storages"
	"seeker/pkg/db/postgres"
)

type userRepository struct {
	client postgres.Client
}

func NewUserRepository(client postgres.Client) storages.UserStorage {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) GetByEmail(email string) (entities.User, error) {
	query := `
		SELECT id, email, picture, password FROM "users"
		WHERE email = $1
	`

	row := r.client.QueryRow(query, email)
	var user entities.User

	if err := row.Scan(&user.ID, &user.Email, &user.Picture, &user.Password); err != nil {
		return user, postgres.NewError(err)
	}

	return user, nil
}

func (r *userRepository) GetByID(id string) (entities.User, error) {
	query := `
		SELECT id, email, picture, password FROM "users"
		WHERE id = $1
	`

	row := r.client.QueryRow(query, id)
	var user entities.User

	if err := row.Scan(&user.ID, &user.Email, &user.Picture, &user.Password); err != nil {
		return user, postgres.NewError(err)
	}

	return user, nil
}

func (r *userRepository) CreateOne(user *entities.User) error {
	query := `
		INSERT INTO "users" (email, password) 
		VALUES ($1, $2)
		RETURNING id, email, picture
	`

	row := r.client.QueryRow(query, user.Email, user.Password)

	if err := row.Scan(&user.ID, &user.Email, &user.Picture); err != nil {
		return postgres.NewError(err)
	}

	return nil
}
