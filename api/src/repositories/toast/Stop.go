package toast_repo

import (
	"errors"
	"fmt"

	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
)

func StopToastSession(t *models.Toast) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	session, err := GetSessionByIDIncomplete(t.SessionID, t.UserID)
	if err != nil {
		return err
	}

	if t.EndTime.Before(session.StartTime) {
		return BFE.New(BFE.ErrUnprocessable, errors.New("EndTime can't be before StartTime"))
	}

	t.Duration = int64(t.EndTime.Sub(session.StartTime).Seconds())

	query := `
		UPDATE toast
		SET end_time = $1, duration = $2, description = $3
		WHERE id = $4 AND user_id = $5
		RETURNING id, user_id, session_name, description, start_time, end_time, duration, category_id
	`

	err = tx.QueryRow(query, t.EndTime, t.Duration, t.Description, t.SessionID, t.UserID).Scan(
		&t.SessionID,
		&t.UserID,
		&t.SessionName,
		&t.Description,
		&t.StartTime,
		&t.EndTime,
		&t.Duration,
		&t.CategoryID,
	)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, fmt.Errorf("failed to stop toast session: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
