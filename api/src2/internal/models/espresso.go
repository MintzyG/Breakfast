package models

import (
  "time"

  "github.com/google/uuid"
)

type EspressoSession struct {
  SessionID   int       `gorm:"primaryKey;autoIncrement" json:"session_id"`
  UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
  Emoji       string    `gorm:"type:varchar(31)" json:"emoji"`
  SessionName string    `gorm:"type:varchar(255);not null" json:"session_name"`
  Duration    int       `gorm:"not null" json:"duration"`
  BreakTime   int       `gorm:"not null" json:"break_time"`
  FocusStart  time.Time `gorm:"not null" json:"focus_start"`
  FocusEnd    time.Time `gorm:"not null" json:"focus_end"`
  IsActive    bool      `gorm:"default:true" json:"is_active"`
  CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
  UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
  Rounds      int       `gorm:"not null;default:1" json:"rounds"`
  BigBreak    int       `gorm:"not null;default:15" json:"big_break"`
  Laps        int       `gorm:"not null;default:1" json:"laps"`
}

type EspressoUserSettings struct {
  SettingsID    int       `gorm:"primaryKey;autoIncrement" json:"settings_id"`
  UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
  FocusDuration int       `gorm:"not null" json:"focus_duration"`
  BreakDuration int       `gorm:"not null" json:"break_duration"`
  Rounds        int       `gorm:"not null;default:1" json:"rounds"`
  BigBreak      int       `gorm:"not null;default:15" json:"big_break"`
  Laps          int       `gorm:"not null;default:1" json:"laps"`
  CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
  UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

