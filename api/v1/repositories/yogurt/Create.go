package yogurt_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"

	_ "github.com/lib/pq"
)

func CreateYogurtTask(task *models.YogurtTask) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `
    INSERT INTO yogurt
      (user_id, emoji, title, description, completed, task_size, difficulty, priority, category_id)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id
  `
	err = tx.QueryRow(
		query,
		task.UserID,
		task.Emoji,
		task.Title,
		task.Description,
		task.Completed,
		task.Difficulty,
		task.TaskSize,
		task.Priority,
		task.CategoryID,
	).Scan(&task.TaskID)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
