package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Yogurt struct {
	TaskID      int       `gorm:"primaryKey;autoIncrement" json:"task_id"`
	UserID      uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Emoji       string    `gorm:"type:varchar(31)" json:"emoji"`
	Color       string    `gorm:"type:varchar(6);default:4CAF50" json:"color"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	TaskSize    int       `gorm:"default:1" json:"task_size" validate:"min=1,max=3,required"`
	Difficulty  int       `gorm:"default:1" json:"difficulty" validate:"min=1,max=3,required"`
	Priority    int       `gorm:"default:1" json:"priority" validate:"min=1,max=3,required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *Yogurt) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
