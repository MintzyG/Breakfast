package toast_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetSessionByID(id int, user_id uuid.UUID) (*models.Toast, error) {
	query := `SELECT session_name, description, start_time, end_time, duration, category_id FROM toast WHERE id = $1 AND user_id = $2`
	var s models.Toast
	s.UserID = user_id
	s.SessionID = id
  var endTime sql.NullTime
	err := R.Instance.QueryRow(query, id, user_id).Scan(&s.SessionName, &s.Description, &s.StartTime, &endTime, &s.Duration, &s.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find session with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

  if !endTime.Valid {
    return nil, BFE.New(BFE.ErrDatabase, errors.New("Could not get endtime"))
  }

  s.EndTime = endTime.Time

	return &s, nil
}

func GetSessionByIDIncomplete(id int, user_id uuid.UUID) (*models.Toast, error) {
	query := `SELECT session_name, description, start_time, end_time, duration, category_id FROM toast WHERE id = $1 AND user_id = $2`
	var s models.Toast
	s.UserID = user_id
	s.SessionID = id
  var endTime sql.NullTime
	err := R.Instance.QueryRow(query, id, user_id).Scan(&s.SessionName, &s.Description, &s.StartTime, &endTime, &s.Duration, &s.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find session with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

  s.EndTime = time.Time{}
	return &s, nil
}
