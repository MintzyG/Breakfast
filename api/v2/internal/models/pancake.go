package models

import (
	"time"

	"github.com/google/uuid"
)

type Pancake struct {
	NoteID  int       `gorm:"primaryKey;autoIncrement" json:"note_id"`
	UserID  uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Emoji   string    `gorm:"type:varchar(31)" json:"emoji"`
	Title   string    `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Content string    `gorm:"type:text;not null" json:"content" validate:"required"`
  Color   string    `gorm:"type:varchar(6);default:4CAF50" json:"color"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (p *Pancake) Validate() error {
	return validate.Struct(p)
}
