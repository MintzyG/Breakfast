package maple_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func GetHabitByID(id int, user_id uuid.UUID) (*models.Maple, error) {
	query := `
    SELECT title, smallest_unit, curr_streak, highest_streak, days_performed, category_id
    FROM maple
    WHERE id = $1 AND user_id = $2
  `

	var h models.Maple
	h.UserID = user_id
	h.HabitID = id
	err := R.Instance.QueryRow(query, id, user_id).Scan(
    &h.Title,
    &h.SmallestUnit,
    &h.CurrStreak,
    &h.HighestStreak,
    &h.DaysPerformed,
    &h.CategoryID,
  )
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find habit with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return &h, nil
}
