package maple_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

func GetAllHabits(user_id uuid.UUID) ([]models.Maple, error) {
	query := "SELECT * FROM maple WHERE user_id = $1"

	rows, err := R.Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer rows.Close()

	var habits []models.Maple
	for rows.Next() {
		var habit models.Maple
		if err := rows.Scan(
      &habit.HabitID,
      &habit.UserID,
      &habit.Title,
      &habit.SmallestUnit,
      &habit.CurrStreak,
      &habit.HighestStreak,
      &habit.DaysPerformed,
      &habit.CategoryID,
    ); err != nil {
			return nil, BFE.New(BFE.ErrDatabase, err)
		}
		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return habits, nil
}
