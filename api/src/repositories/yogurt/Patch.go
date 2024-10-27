package yogurt_repo

import (
	BFE "breakfast/errors"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func PatchTask(id int, user_id uuid.UUID, updates map[string]interface{}) error {
	validFields := map[string]bool{
		"emoji":       true,
		"title":       true,
		"description": true,
		"task_size":   true,
		"difficulty":  true,
		"priority":    true,
		"completed":   true,
		"category_id": true,
	}

	whereClause := "id = $1 AND user_id = $2"
	query, args, err := R.BuildUpdateQuery("yogurt", updates, validFields, whereClause, id, user_id)
	if err != nil {
		return err
	}

	_, execErr := R.Instance.Exec(query, args...)
	if execErr != nil {
		if execErr == sql.ErrNoRows {
			return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find task with ID: %v", id))
		}
		return BFE.New(BFE.ErrDatabase, execErr)
	}

	return nil
}
