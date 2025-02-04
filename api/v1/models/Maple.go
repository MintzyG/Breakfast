package models

import (
	"time"

	"github.com/google/uuid"
)

type Maple struct {
	HabitID       int         `json:"habit_id"`
	UserID        uuid.UUID   `json:"user_id"`
	Title         string      `json:"title"`
	SmallestUnit  string      `json:"smallest_unit"`
	CurrStreak    int         `json:"curr_streak"`
	HighestStreak int         `json:"highest_streak"`
	DaysPerformed []time.Time `json:"days_performed"`
	CategoryID    int         `json:"category_id"`
}
