package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Category struct {
	ID          int       `json:"category_id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Emoji       string    `json:"emoji"`
	Color       string    `json:"color"`
	TextColor   string    `json:"text_color"`
}

func (c Category) String() string {
	return fmt.Sprintf("%v - %v", c.Emoji, c.Title)
}
