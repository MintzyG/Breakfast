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
	StartTime   time.Time `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime     time.Time `gorm:"type:timestamp;not null" json:"end_time"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Reminders []ParfaitReminder `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"reminders"`
}

type ParfaitReminder struct {
	ReminderID       int       `gorm:"primaryKey;autoIncrement" json:"reminder_id"`
	EventID          int       `gorm:"not null" json:"event_id"`
	UserID           uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	RemindAt         time.Time `gorm:"type:timestamp;" json:"remind_at" validate:"required"`
	Title            string    `gorm:"type:varchar(255);not null" json:"title"`
	Description      string    `gorm:"type:text" json:"description"`
	NotificationSent bool      `gorm:"default:false" json:"notification_sent"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
