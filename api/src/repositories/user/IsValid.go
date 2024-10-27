package user_repo

import (
	BFE "breakfast/errors"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

func IsUserValid(id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	var exists bool
	err := R.Instance.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, BFE.New(BFE.ErrDatabase, err)
	}
	return exists, nil
}
