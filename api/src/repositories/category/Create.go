package category_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"errors"

	"github.com/lib/pq"
)

func CreateCategory(c *models.Category) error {
  tx, err := R.BeginTransaction()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }
  defer tx.Rollback()

	query := `INSERT INTO categories (user_id, title, description, emoji, color, text_color) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = tx.QueryRow(query, c.UserId, c.Title, c.Description, c.Emoji, c.Color, c.TextColor).Scan(&c.ID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return BFE.New(BFE.ErrConflict, errors.New("Category with this title already exists"))
			}
		}
		return BFE.New(BFE.ErrDatabase, err)
	}

  err = tx.Commit()
  if err != nil {
    return BFE.New(BFE.ErrDatabase, err)
  }
 
	return nil
}
