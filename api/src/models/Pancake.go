package models

import (
	"time"

	"github.com/google/uuid"
)

type Pancake struct {
	NoteID     int       `json:"note_id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID int       `json:"category_id"`
}
