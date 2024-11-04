package toast_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

// Will break if a session has no endtime
func GetAllSessions(user_id uuid.UUID) ([]models.Toast, error) {
	query := "SELECT * FROM toast WHERE user_id = $1"

	rows, err := R.Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer rows.Close()

	var sessions []models.Toast
	for rows.Next() {
		var session models.Toast
		if err := rows.Scan(&session.SessionID, &session.UserID, &session.SessionName, &session.Description, &session.StartTime, &session.EndTime, &session.Duration, &session.Active, &session.CategoryID); err != nil {
			return nil, BFE.New(BFE.ErrDatabase, err)
		}
		sessions = append(sessions, session)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return sessions, nil
}
