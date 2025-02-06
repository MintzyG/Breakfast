package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Maple struct {
	HabitID         int        `gorm:"primaryKey;autoIncrement" json:"habit_id"`
	UserID          uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	Emoji           string     `gorm:"type:varchar(31)" json:"emoji"`
	Title           string     `gorm:"type:varchar(255);not null" json:"title"`
	SmallestUnit    int        `gorm:"default:1;not null" json:"smallest_unit" validate:"required"`
	SmallestMeasure string     `gorm:"type:varchar(15);not null" json:"smallest_measure" validate:"required,min=0"`
	CurrStreak      int        `gorm:"default:0" json:"curr_streak" validate:"eq=0"`
	HighestStreak   int        `gorm:"default:0" json:"highest_streak" validate:"eq=0"`
	MapleDays       []MapleDay `gorm:"foreignKey:HabitID;constraint:OnDelete:CASCADE" json:"maple_days"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MapleDay struct {
	HabitID   int  `gorm:"not null" json:"habit_id"`
  DayID     int  `gorm:"primaryKey;autoIncrement" json:"day_id"`
	UnitsDone int  `gorm:"default:0" json:"units_done" validate:"required,min=1"`
	Achieved  bool `gorm:"default:false" json:"achieved"`
  Date time.Time `gorm:"type:date;not null" json:"date"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *Maple) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func (t *MapleDay) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
