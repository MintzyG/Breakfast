package category_repo

import (
	BFE "breakfast/errors"
	"breakfast/models"
	R "breakfast/repositories"

	"github.com/google/uuid"
)

func GetAllCategories(user_id uuid.UUID) ([]models.Category, error) {
	query := "SELECT * FROM categories WHERE user_id = $1"

	rows, err := R.Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.UserId, &category.Title, &category.Description, &category.Emoji, &category.Color, &category.TextColor); err != nil {
			return nil, BFE.New(BFE.ErrDatabase, err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.New(BFE.ErrDatabase, err)
	}

	return categories, nil
}
