package models

import (
	"time"

	"github.com/google/uuid"
)

type ParfaitEvent struct {
	EventID     int       `gorm:"primaryKey;autoIncrement" json:"event_id"`
	UserID      uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	StartTime   time.Time `gorm:"not null" json:"start_time"`
	EndTime     time.Time `gorm:"not null" json:"end_time"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
	CategoryID  *int      `json:"category_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ParfaitReminder struct {
	ReminderID       int       `gorm:"primaryKey;autoIncrement" json:"reminder_id"`
	EventID          int       `gorm:"not null" json:"event_id"`
	UserID           uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	ReminderAt       time.Time `gorm:"not null" json:"reminder_at"`
	NotificationSent bool      `gorm:"default:false" json:"notification_sent"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
