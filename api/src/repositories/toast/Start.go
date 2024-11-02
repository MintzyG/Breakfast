package toast_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"fmt"
)

func StartToastSession(t *models.Toast) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO toast (user_id, session_name, description, start_time, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = tx.QueryRow(query, t.UserID, t.SessionName, t.Description, t.StartTime, t.CategoryID).Scan(&t.SessionID)
	if err != nil {
		if R.IsForeignKeyViolation(err) {
			return BFE.New(BFE.ErrDatabase, fmt.Errorf("foreign key violation: user_id %v may not exist", t.UserID))
		}
		return BFE.New(BFE.ErrDatabase, fmt.Errorf("failed to insert toast session: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
