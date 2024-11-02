package toast_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
)

func CreateToastSession(t *models.Toast) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO toast (user_id, session_name, description, start_time, end_time, duration, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = tx.QueryRow(query, t.UserID, t.SessionName, t.Description, t.StartTime, t.EndTime, t.Duration, t.CategoryID).Scan(&t.SessionID)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
