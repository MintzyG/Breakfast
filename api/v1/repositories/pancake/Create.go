package pancake_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
)

func CreateNote(p *models.Pancake) error {
	tx, err := R.BeginTransaction()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO pancake (user_id, title, content, category_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err = tx.QueryRow(query, p.UserID, p.Title, p.Content, p.CategoryID).Scan(&p.NoteID)
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	err = tx.Commit()
	if err != nil {
		return BFE.New(BFE.ErrDatabase, err)
	}

	return nil
}
