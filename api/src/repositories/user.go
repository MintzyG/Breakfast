package repositories

import (
	"breakfast/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := Instance.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	var user models.User
	err := Instance.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email = $1`
	var user models.User
	err := Instance.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}
