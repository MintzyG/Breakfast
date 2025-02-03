package models

import (
	"time"

	"github.com/google/uuid"
)

type Toast struct {
	SessionID   int        `gorm:"primaryKey;autoIncrement" json:"session_id"`
	UserID      uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	Emoji       string     `gorm:"type:varchar(31)" json:"emoji"`
	SessionName string     `gorm:"type:varchar(255);not null" json:"session_name"`
	Description string     `gorm:"type:text" json:"description"`
	StartTime   time.Time  `gorm:"autoCreateTime" json:"start_time"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Duration    int64      `json:"duration"`
	Active      bool       `gorm:"default:true" json:"active"`
	CategoryID  *int       `json:"category_id,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
