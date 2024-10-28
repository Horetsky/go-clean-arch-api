package postgresql

import (
	"fmt"
	"seeker/internal/domain/entities"
	"seeker/internal/domain/repositories"
	"seeker/pkg/db/postgres"
	"seeker/pkg/utils/str"
)

type userRepository struct {
	client postgres.Client
}

func NewUserRepository(client postgres.Client) repositories.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Create(user *entities.User) error {
	query := `
		INSERT INTO "users" (email, password, type) 
		VALUES ($1, $2, $3)
		RETURNING id, email, picture, email_verified
	`

	row := r.client.QueryRow(query, user.Email, user.Password, user.Type)

	if err := row.Scan(&user.ID, &user.Email, &user.Picture, &user.EmailVerified); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	query := `
		SELECT id, email, picture, email_verified, password, type FROM users
		WHERE email = $1
	`

	row := r.client.QueryRow(query, email)
	var user entities.User

	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Picture,
		&user.EmailVerified,
		&user.Password,
		&user.Type,
	); err != nil {
		return user, postgres.NewError(err)
	}

	return user, nil
}

func (r *userRepository) FindByID(id string) (entities.User, error) {
	query := `
		SELECT id, email, picture, email_verified, password FROM "users"
		WHERE id = $1
	`

	row := r.client.QueryRow(query, id)
	var user entities.User

	if err := row.Scan(&user.ID, &user.Email, &user.Picture, &user.EmailVerified, &user.Password); err != nil {
		return user, postgres.NewError(err)
	}

	return user, nil
}

func (r *userRepository) DeleteByEmail(email string) error {
	query := `
		DELETE FROM users
		WHERE email = $1
	`

	row := r.client.QueryRow(query, email)

	if err := row.Scan(); err != nil {
		return postgres.NewError(err)
	}

	return nil
}

func (r *userRepository) UpdateByEmail(email string, user *entities.User) error {
	query := `
		UPDATE "users" SET
	`

	str.ForEach(user, func(key string, val any) {
		switch key {
		case "Email":
			query += fmt.Sprintf(" email = '%s',", user.Email)
			break
		case "Password":
			query += fmt.Sprintf(" password = '%s',", user.Password)
			break
		case "Picture":
			query += fmt.Sprintf(" picture = '%s',", user.Password)
			break
		case "EmailVerified":
			query += fmt.Sprintf(" email_verified = '%v',", user.EmailVerified)
			break
		}
	})

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE email = '%s'", email)
	query += " RETURNING id, email, picture, email_verified, type"

	rows := r.client.QueryRow(query)

	if err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Picture,
		&user.EmailVerified,
		&user.Type,
	); err != nil {
		return postgres.NewError(err)
	}

	return nil
}
