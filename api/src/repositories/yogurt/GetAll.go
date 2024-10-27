package yogurt_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func GetAllTasks(user_id uuid.UUID) ([]models.YogurtTask, error) {
	query := "SELECT * FROM yogurt WHERE user_id = $1"

	rows, err := R.Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer rows.Close()

	var tasks []models.YogurtTask
	for rows.Next() {
		var task models.YogurtTask
		if err := rows.Scan(
			&task.TaskID,
			&task.UserID,
			&task.Emoji,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Difficulty,
			&task.TaskSize,
			&task.Priority,
			&task.CategoryID,
		); err != nil {
			return nil, BFE.New(BFE.ErrDatabase, err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return tasks, nil
}
