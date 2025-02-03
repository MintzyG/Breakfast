package models

import (
	"time"

	"github.com/google/uuid"
)

type OmeletteTable struct {
	TableID   int       `gorm:"primaryKey;autoIncrement" json:"table_id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Emoji     string    `gorm:"type:varchar(31)" json:"emoji"`
	TableName string    `gorm:"type:varchar(255);not null" json:"table_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Lists []OmeletteList `gorm:"foreignKey:TableID" json:"lists"`
}

type OmeletteList struct {
	ListID    int       `gorm:"primaryKey;autoIncrement" json:"list_id"`
	TableID   int       `gorm:"not null" json:"table_id"`
	ListName  string    `gorm:"type:varchar(255);not null" json:"list_name"`
	Position  int       `gorm:"not null" json:"position"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Cards []OmeletteCard `gorm:"foreignKey:ListID" json:"cards"`
}

type OmeletteCard struct {
	CardID    int       `gorm:"primaryKey;autoIncrement" json:"card_id"`
	ListID    int       `gorm:"not null" json:"list_id"`
	CardName  string    `gorm:"type:varchar(255);not null" json:"card_name"`
	Position  int       `gorm:"not null" json:"position"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
