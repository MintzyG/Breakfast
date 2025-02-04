package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserRegister struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type User struct {
	ID             uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name           string    `gorm:"size:100;not null" json:"name"`
	Email          string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password       string    `gorm:"type:varchar(255);not null" json:"-"`
	ProfilePicture string    `gorm:"type:varchar(255)" json:"profile_picture,omitempty"`
	Bio            string    `gorm:"type:text" json:"bio,omitempty"`
	Location       string    `gorm:"size:100" json:"location,omitempty"`
	PhoneNumber    string    `gorm:"size:20" json:"phone_number,omitempty"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	UserLinks        []UserLink             `gorm:"foreignKey:UserID" json:"user_links,omitempty"`
	Toasts           []Toast                `gorm:"foreignKey:UserID" json:"toasts"`
	Omelettes        []OmeletteTable        `gorm:"foreignKey:UserID" json:"omelettes"`
	Pancakes         []Pancake              `gorm:"foreignKey:UserID" json:"pancakes"`
	CerealDays       []CerealDay            `gorm:"foreignKey:UserID" json:"cereal_days"`
	YogurtTasks      []Yogurt               `gorm:"foreignKey:UserID" json:"yogurt_tasks"`
	EspressoSessions []EspressoSession      `gorm:"foreignKey:UserID" json:"espresso_sessions"`
	Maples           []Maple                `gorm:"foreignKey:UserID" json:"maples"`
	Parfaits         []ParfaitEvent         `gorm:"foreignKey:UserID" json:"parfaits"`
	EspressoConfigs  []EspressoUserSettings `gorm:"foreignKey:UserID" json:"espresso_configs"`
	ParfaitReminders []ParfaitReminder      `gorm:"foreignKey:UserID" json:"parfait_reminders"`
}

type UserLink struct {
	UserID uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Link   string    `gorm:"type:varchar(255);not null" json:"link"`
}

type UserClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
