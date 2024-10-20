package repositories

import (
	"breakfast/models"
	_ "github.com/lib/pq"
)

func CreateYogurtTask(task *models.YogurtTask) error {
	query := `
    INSERT INTO yogurt_task
      (user_id, emoji, title, description, task_size, difficulty, priority, category_id)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
  `
	err := Instance.QueryRow(
		query,
		task.UserID,
		task.Emoji,
		task.Title,
		task.Description,
		task.Difficulty,
		task.TaskSize,
		task.Priority,
		task.CategoryID,
	).Scan(&task.TaskID)
	if err != nil {
		return err
	}
	return nil
}
