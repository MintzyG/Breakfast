package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CerealDay struct {
	DayID     int       `gorm:"primaryKey;autoIncrement" json:"day_id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:idx_user_date" json:"user_id"`
	Emoji     string    `gorm:"type:varchar(31)" json:"emoji"`
	Date      time.Time `gorm:"type:date;not null;uniqueIndex:idx_user_date" json:"date" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Activities []CerealActivity `gorm:"foreignKey:DayID;constraint:OnDelete:CASCADE" json:"activities"`
}

type CerealActivity struct {
	ActivityID int       `gorm:"primaryKey;autoIncrement" json:"activity_id"`
	DayID      int       `gorm:"not null" json:"day_id"`
	Title      string    `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Date       time.Time `gorm:"type:date;not null" json:"date" validate:"required"`
	StartTime  time.Time `gorm:"type:timestamp;not null" json:"start_time" validate:"required"`
	EndTime    time.Time `gorm:"type:timestamp;not null" json:"end_time" validate:"required"`
	Notify     bool      `gorm:"default:false" json:"notify"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *CerealDay) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
