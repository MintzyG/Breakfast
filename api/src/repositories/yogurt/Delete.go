package yogurt_repo

import (
	BFE "breakfast/errors"
	R "breakfast/repositories"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func DeleteTask(id int, user_id uuid.UUID) error {
  tx, err := R.BeginTransaction()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }
  defer tx.Rollback()

	query := `DELETE FROM yogurt WHERE id = $1 AND user_id = $2`

  result, err := tx.Exec(query, id, user_id)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	if rowsAffected == 0 {
		return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find task with ID: %v", id))
	}

  err = tx.Commit()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }

	return nil
}
