package user_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func GetUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email FROM users WHERE id = $1`
	var user models.User
	err := R.Instance.QueryRow(query, id).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrUserNotFound, errors.New("User not found."))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	return &user, nil
}
