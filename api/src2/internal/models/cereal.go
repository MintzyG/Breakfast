package models

import (
    "time"
    "github.com/google/uuid"
)

type CerealDay struct {
    DayID     int       `gorm:"primaryKey;autoIncrement" json:"day_id"`
    UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    Date      time.Time `gorm:"type:date;not null" json:"date"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    Activities []CerealActivity `gorm:"foreignKey:DayID" json:"activities"`
}

type CerealActivity struct {
    ActivityID int       `gorm:"primaryKey;autoIncrement" json:"activity_id"`
    DayID       int       `gorm:"not null" json:"day_id"`
    Title       string    `gorm:"type:varchar(255);not null" json:"title"`
    StartTime   time.Time `gorm:"not null" json:"start_time"`
    EndTime     time.Time `gorm:"not null" json:"end_time"`
    Category    string    `gorm:"type:varchar(100);not null" json:"category"`
    Notify      bool      `gorm:"default:false" json:"notify"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

