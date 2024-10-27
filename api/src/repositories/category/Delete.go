package category_repo

import (
	BFE "breakfast/errors"
	R "breakfast/repositories"
	"fmt"

	"github.com/google/uuid"
)

func DeleteCategory(id int, user_id uuid.UUID) error {
	query := `SELECT delete_category($1, $2);`
	var success bool
	err := R.Instance.QueryRow(query, user_id, id).Scan(&success)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	if !success {
		return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("could not find category with ID: %v", id))
	}

	return nil
}
