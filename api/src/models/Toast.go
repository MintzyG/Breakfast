package models

import (
	"time"

	"github.com/google/uuid"
)

type Toast struct {
	SessionID   int       `json:"session_id"`
	UserID      uuid.UUID `json:"user_id"`
	SessionName string    `json:"session_name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time,omitempty"`
	Duration    int64     `json:"duration"`
	CategoryID  *int      `json:"category_id"`
}
