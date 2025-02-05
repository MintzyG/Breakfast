package models

import (
	"time"

	"github.com/google/uuid"
)

type ToastSession struct {
	SessionID   int        `gorm:"primaryKey;autoIncrement" json:"session_id"`
	UserID      uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	Emoji       string     `gorm:"type:varchar(31)" json:"emoji"`
  SessionName string     `gorm:"type:varchar(255);not null" json:"session_name" validate:"required"`
	Description string     `gorm:"type:text" json:"description"`
  StartTime   time.Time  `gorm:"type:timestamp;not null" json:"start_time" validate:"required"`
  EndTime     *time.Time `gorm:"type:timestamp;not null" json:"end_time" validate:"required"`
  Duration    int64      `gorm:"not null" json:"duration"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
