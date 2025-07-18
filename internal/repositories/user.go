package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO 
			users(
				role_id, 
				first_name,
				last_name,
				username,
				email,
				password
			)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		user.RoleId,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT 
			id, role_id, first_name, last_name, username, email, created_at, updated_at 
		FROM
			users
		WHERE username = $1
	`

	err := r.DB.
		QueryRow(
			query,
			username,
		).
		Scan(&user.Id,
			&user.RoleId,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT 
			id, role_id, first_name, last_name, username, email, created_at, updated_at 
		FROM
			users
		WHERE email = $1
	`

	err := r.DB.
		QueryRow(
			query,
			email,
		).
		Scan(&user.Id,
			&user.RoleId,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
