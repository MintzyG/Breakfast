package models

import (
    "time"

    "github.com/google/uuid"
)

type Pancake struct {
    NoteID    int       `gorm:"primaryKey;autoIncrement" json:"note_id"`
    UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    Emoji     string    `gorm:"type:varchar(31)" json:"emoji"`
    Title     string    `gorm:"type:varchar(255);not null" json:"title"`
    Content   string    `gorm:"type:text;not null" json:"content"`
    CategoryID int      `json:"category_id,omitempty"`

    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
