package repositories

import (
	BFE "breakfast/errors"
	"breakfast/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func CreateYogurtTask(task *models.YogurtTask) error {
	query := `
    INSERT INTO yogurt
      (user_id, emoji, title, description, completed, task_size, difficulty, priority, category_id)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id
  `
	err := Instance.QueryRow(
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
	return nil
}

func GetAllTasks(user_id uuid.UUID) ([]models.YogurtTask, error) {
	query := "SELECT * FROM yogurt WHERE user_id = $1"

	rows, err := Instance.Query(query, user_id)
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

func GetTaskByID(id int, user_id uuid.UUID) (*models.YogurtTask, error) {
	query := `SELECT * FROM yogurt WHERE id = $1 AND user_id = $2`
	var task models.YogurtTask
	task.UserID = user_id
	task.TaskID = id
	err := Instance.QueryRow(query, id, user_id).Scan(
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
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find category with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	return &task, nil
}

func AlterTaskCompletedStatus(id int, user_id uuid.UUID, status bool) error {
	query := `
  UPDATE yogurt
  SET completed = $1
  WHERE id = $2 AND user_id = $3;
  `

	_, err := Instance.Exec(query, status, id, user_id)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}

func PatchTask(id int, user_id uuid.UUID, updates map[string]interface{}) error {
	validFields := map[string]bool{
		"emoji":       true,
		"title":       true,
		"description": true,
		"task_size":   true,
		"difficulty":  true,
		"priority":    true,
    "completed":   true,
		"category_id": true,
	}

	whereClause := "id = $1 AND user_id = $2"
	query, args, err := BuildUpdateQuery("yogurt", updates, validFields, whereClause, id, user_id)
	if err != nil {
		return err
	}

	_, execErr := Instance.Exec(query, args...)
	if execErr != nil {
		if execErr == sql.ErrNoRows {
			return BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find task with ID: %v", id))
		}
		return BFE.New(BFE.ErrDatabase, execErr)
	}

	return nil
}
