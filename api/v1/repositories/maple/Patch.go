package maple_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"
)

func PatchHabit(habit models.Maple, updates map[string]bool) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	whereClause := "id = $1 AND user_id = $2"
	query, args, err := R.BuildUpdateQuery("maple", habit, updates, whereClause, habit.HabitID, habit.UserID)
	if err != nil {
		return err
	}

	_, execErr := tx.Exec(query, args...)
	if execErr != nil {
		if execErr == sql.ErrNoRows {
			return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find habit with ID: %v", habit.HabitID))
		}
		return BFE.New(BFE.ErrDatabase, execErr)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
