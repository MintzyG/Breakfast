package pancake_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func GetNoteByID(id int, user_id uuid.UUID) (*models.Pancake, error) {
	query := `SELECT title, content, created_at, updated_at, category_id FROM pancake WHERE id = $1 AND user_id = $2`
	var p models.Pancake
	p.UserID = user_id
	p.NoteID = id
	err := R.Instance.QueryRow(query, id, user_id).Scan(&p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find note with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return &p, nil
}
