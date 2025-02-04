package category_repo

import (
	BFE "breakfast/_internal/errors"
	R "breakfast/repositories"
	"fmt"

	"github.com/google/uuid"
)

func DeleteCategory(id int, user_id uuid.UUID) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `SELECT delete_category($1, $2);`
	var success bool
	err = tx.QueryRow(query, user_id, id).Scan(&success)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	if !success {
		return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("could not find category with ID: %v", id))
	}

	err = tx.Commit()
	if err != nil {
		BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
