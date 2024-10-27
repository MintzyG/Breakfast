package yogurt_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func GetTaskByID(id int, user_id uuid.UUID) (*models.YogurtTask, error) {
	query := `SELECT * FROM yogurt WHERE id = $1 AND user_id = $2`
	var task models.YogurtTask
	task.UserID = user_id
	task.TaskID = id
	err := R.Instance.QueryRow(query, id, user_id).Scan(
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find task with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	return &task, nil
}
