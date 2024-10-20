package repositories

import (
	BFE "breakfast/errors"
	"breakfast/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateCategory(c *models.Category) error {
	query := `INSERT INTO categories (user_id, title, description, emoji, color, text_color) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := Instance.QueryRow(query, c.UserId, c.Title, c.Description, c.Emoji, c.Color, c.TextColor).Scan(&c.ID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return BFE.NewBFError(BFE.CONFLICT_ERROR_CODE, "Category with this title already exists")
			}
		}
		return BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	return nil
}

func GetCategoryByID(id int, user_id uuid.UUID) (*models.Category, error) {
	query := `SELECT title, description, emoji, color, text_color FROM categories WHERE id = $1 AND user_id = $2`
	var c models.Category
	c.UserId = user_id
	c.ID = id
	err := Instance.QueryRow(query, id, user_id).Scan(&c.Title, &c.Description, &c.Emoji, &c.Color, &c.TextColor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, BFE.NewBFError(BFE.RESOURCE_NOT_FOUND_CODE, fmt.Sprintf("Could not find category with ID: %v", id))
		}
		return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	return &c, nil
}

func GetAllCategories(user_id uuid.UUID) ([]models.Category, error) {
	query := "SELECT * FROM categories WHERE user_id = $1"

	rows, err := Instance.Query(query, user_id)
	if err != nil {
		return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.UserId, &category.Title, &category.Description, &category.Emoji, &category.Color, &category.TextColor); err != nil {
			return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, BFE.NewBFError(BFE.DATABASE_ERROR_CODE, err.Error())
	}

	return categories, nil
}
