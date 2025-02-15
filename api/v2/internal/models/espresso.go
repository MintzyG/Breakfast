package models

import (
	"time"

	"github.com/google/uuid"
)

type EspressoSession struct {
	SessionID   int       `gorm:"primaryKey;autoIncrement" json:"session_id"`
	UserID      uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Emoji       string    `gorm:"type:varchar(31)" json:"emoji"`
	Color       string    `gorm:"type:varchar(7);default:4CAF50" json:"color"`
	SessionName string    `gorm:"type:varchar(255);not null" json:"session_name" validate:"required"`
	Duration    int64     `gorm:"not null" json:"duration"`
	BreakTime   int64       `gorm:"not null" json:"break_time"`
	Distractions   int    `gorm:"default:0" json:"distractions"`
	FocusStart  time.Time `gorm:"type:timestamp;not null" json:"focus_start" validate:"required"`
	FocusEnd    time.Time `gorm:"type:timestamp;not null" json:"focus_end" validate:"required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
