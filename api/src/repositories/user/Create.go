package user_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"errors"

	"github.com/lib/pq"
)

func CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := R.Instance.Exec(query, user.UserID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return BFE.New(BFE.ErrConflict, errors.New("Email already exists"))
			}
		}
		return BFE.New(BFE.ErrDatabase, err)
	}
	return nil
}
