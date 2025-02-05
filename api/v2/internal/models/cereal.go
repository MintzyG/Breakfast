package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

type CerealDay struct {
	DayID     int        `gorm:"primaryKey;autoIncrement" json:"day_id"`
	UserID    uuid.UUID  `gorm:"type:char(36);not null;uniqueIndex:idx_user_date" json:"user_id"`
	Emoji     string     `gorm:"type:varchar(31)" json:"emoji"`
  Date      time.Time  `gorm:"type:date;not null;uniqueIndex:idx_user_date" json:"date" validate:"required"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	Activities []CerealActivity `gorm:"foreignKey:DayID;constraint:OnDelete:CASCADE" json:"activities"`
}

type CerealActivity struct {
	ActivityID int        `gorm:"primaryKey;autoIncrement" json:"activity_id"`
	DayID      int        `gorm:"not null" json:"day_id"`
	Title      string     `gorm:"type:varchar(255);not null" json:"title"`
  Date       time.Time  `gorm:"type:date;not null" json:"date"`
  StartTime  time.Time  `gorm:"type:timestamp;not null" json:"start_time"`
  EndTime    time.Time  `gorm:"type:timestamp;not null" json:"end_time"`
	Notify     bool       `gorm:"default:false" json:"notify"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *CerealDay) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
