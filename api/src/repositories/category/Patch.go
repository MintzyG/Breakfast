package category_repo

import (
	BFE "breakfast/_internal/errors"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func PatchCategory(id int, user_id uuid.UUID, updates map[string]interface{}) error {
  tx, err := R.BeginTransaction()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }
  defer tx.Rollback()

	validFields := map[string]bool{
		"title":       true,
		"description": true,
		"emoji":       true,
		"color":       true,
		"text_color":  true,
	}

	whereClause := "id = $1 AND user_id = $2"
	query, args, err := R.BuildUpdateQuery("categories", updates, validFields, whereClause, id, user_id)
	if err != nil {
		return err
	}

	_, execErr := tx.Exec(query, args...)
	if execErr != nil {
		if execErr == sql.ErrNoRows {
			return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find category with ID: %v", id))
		}
		return BFE.New(BFE.ErrDatabase, execErr)
	}

  err = tx.Commit()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }

	return nil
}
