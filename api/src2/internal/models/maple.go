package models

import (
	"time"

	"github.com/google/uuid"
)

type Maple struct {
	HabitID       int        `gorm:"primaryKey;autoIncrement" json:"habit_id"`
	UserID        uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	Emoji         string     `gorm:"type:varchar(31)" json:"emoji"`
	Title         string     `gorm:"type:varchar(255);not null" json:"title"`
	SmallestUnit  string     `gorm:"type:varchar(50);not null" json:"smallest_unit"`
	CurrStreak    int        `gorm:"default:0" json:"curr_streak"`
	HighestStreak int        `gorm:"default:0" json:"highest_streak"`
	MapleDays     []MapleDay `gorm:"foreignKey:HabitID" json:"maple_days"`
	CategoryID    int        `json:"category_id,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MapleDay struct {
	HabitID     int       `gorm:"not null" json:"habit_id"`
	PerformedAt time.Time `gorm:"not null" json:"performed_at"`
}
