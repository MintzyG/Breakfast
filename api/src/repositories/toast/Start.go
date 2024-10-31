package toast_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"
)

func StartToastSession(t *models.Toast) (*models.Toast, error) {
  tx, err := R.BeginTransaction()
  if err != nil {
    return nil, BFE.New(BFE.ErrDatabase, err)
  }
  defer tx.Rollback()
  
	query := `INSERT INTO toast (user_id, session_name, description, start_time, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = tx.QueryRow(query, t.UserID, t.SessionName, t.Description, t.StartTime, t.CategoryID).Scan(&t.SessionID)
	if err != nil {
    return nil, BFE.New(BFE.ErrDatabase, err)
	}

  err = tx.Commit()
  if err != nil {
    return nil, BFE.New(BFE.ErrDatabase, err)
  }
 
	return t, nil
}
