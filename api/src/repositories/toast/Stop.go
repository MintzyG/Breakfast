package toast_repo

import (
	"errors"
	"fmt"
	"time"

	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

func StopToastSession(id int, user_id uuid.UUID) (*models.Toast, error) {
	tx, err := R.BeginTransaction()
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	session, err := GetSessionByID(id, user_id)
	if err != nil || !session.Active {
		return nil, err
	}

  session.EndTime = time.Now()
  session.Active = false

	if session.EndTime.Before(session.StartTime) {
		return nil, BFE.New(BFE.ErrUnprocessable, errors.New("EndTime can't be before StartTime"))
	}

	session.Duration = int64(session.EndTime.Sub(session.StartTime).Seconds())

	query := `
		UPDATE toast
		SET end_time = $1, duration = $2, description = $3, active = $4
		WHERE id = $5 AND user_id = $6
		RETURNING id, user_id, session_name, description, start_time, end_time, duration, active, category_id
	`

	err = tx.QueryRow(query, session.EndTime, session.Duration, session.Description, session.Active, session.SessionID, session.UserID).Scan(
		&session.SessionID,
		&session.UserID,
		&session.SessionName,
		&session.Description,
		&session.StartTime,
		&session.EndTime,
		&session.Duration,
		&session.Active,
		&session.CategoryID,
	)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, fmt.Errorf("failed to stop toast session: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return session, nil
}
