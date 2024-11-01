package category_repo

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	R "breakfast/repositories"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func GetCategoryByID(id int, user_id uuid.UUID) (*models.Category, error) {
	query := `SELECT title, description, emoji, color, text_color FROM categories WHERE id = $1 AND user_id = $2`
	var c models.Category
	c.UserId = user_id
	c.ID = id
	err := R.Instance.QueryRow(query, id, user_id).Scan(&c.Title, &c.Description, &c.Emoji, &c.Color, &c.TextColor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.New(BFE.ErrResourceNotFound, fmt.Errorf("Could not find category with ID: %v", id))
		}
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	return &c, nil
}
