package models

import "github.com/google/uuid"

type YogurtTask struct {
	TaskID      int       `json:"task_id"`
	UserID      uuid.UUID `json:"user_id"`
	Emoji       string    `json:"emoji"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TaskSize    int       `json:"task_size"`
	Difficulty  int       `json:"difficulty"`
	Priority    int       `json:"priority"`
	CategoryID  int       `json:"category_id"`
}
