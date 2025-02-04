package maple_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
)

func CreateHabit(h *models.Maple) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `
    INSERT INTO maple (user_id, title, smallest_unit, curr_streak, highest_streak, days_performed, category_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id
  `

	err = tx.QueryRow(query, h.UserID, h.Title, h.SmallestUnit, h.CurrStreak, h.HighestStreak, h.DaysPerformed, h.CategoryID).Scan(&h.HabitID)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
