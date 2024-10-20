package repositories

import (
  BFE "breakfast/errors"
	"breakfast/models"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := Instance.Exec(query, user.UserID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	return nil
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	var user models.User
	err := Instance.QueryRow(query, id).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.NewBFError(BFE.USER_NOT_FOUND_CODE, err.Error())
		}
		return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email = $1`
	var user models.User
	err := Instance.QueryRow(query, email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.NewBFError(BFE.USER_NOT_FOUND_CODE, err.Error())
		}
		return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	return &user, nil
}

func IsUserValid(id uuid.UUID) bool {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`

	var exists bool
	err := Instance.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
