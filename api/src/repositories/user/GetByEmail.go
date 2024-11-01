package user_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"errors"
)

func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email = $1`
	var user models.User
	err := R.Instance.QueryRow(query, email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrUserNotFound, errors.New("User not found."))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	return &user, nil
}
