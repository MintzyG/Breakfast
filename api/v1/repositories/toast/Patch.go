package toast_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"
)

func PatchSession(session models.Toast, updates map[string]bool) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	whereClause := "id = $1 AND user_id = $2"
	query, args, err := R.BuildUpdateQuery("toast", session, updates, whereClause, session.SessionID, session.UserID)
	if err != nil {
		return err
	}

	_, execErr := tx.Exec(query, args...)
	if execErr != nil {
		if execErr == sql.ErrNoRows {
			return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find session with ID: %v", session.SessionID))
		}
		return BFE.New(BFE.ErrDatabase, execErr)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
