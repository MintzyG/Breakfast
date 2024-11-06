package pancake_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

func GetAllNotes(user_id uuid.UUID) ([]models.Pancake, error) {
	query := "SELECT * FROM pancake WHERE user_id = $1"

	rows, err := R.Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer rows.Close()

	var notes []models.Pancake
	for rows.Next() {
		var note models.Pancake
		if err := rows.Scan(&note.NoteID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.CategoryID); err != nil {
			return nil, BFE.New(BFE.ErrDatabase, err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return notes, nil
}
