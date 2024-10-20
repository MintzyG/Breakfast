package repositories

import (
	"breakfast/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func CreateYogurtTask(task *models.YogurtTask, user_id uuid.UUID) error {
	query := `
    INSERT INTO yogurt_task
      (user_id, emoji, title, description, size, difficulty, priority, category_id)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
  `
	err := Instance.QueryRow(
		query,
		user_id,
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
