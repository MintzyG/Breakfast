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

	UserLinks        []UserLink             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user_links,omitempty"`
	Toasts           []ToastSession         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"toasts"`
	Omelettes        []OmeletteTable        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"omelettes"`
	Pancakes         []Pancake              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"pancakes"`
	CerealDays       []CerealDay            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"cereal_days"`
	YogurtTasks      []Yogurt               `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"yogurt_tasks"`
	EspressoSessions []EspressoSession      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"espresso_sessions"`
	Maples           []Maple                `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"maples"`
	Parfaits         []ParfaitEvent         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"parfaits"`
	EspressoConfigs  []EspressoUserSettings `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"espresso_configs"`
	ParfaitReminders []ParfaitReminder      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"parfait_reminders"`
}

type UserLink struct {
	UserID uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Link   string    `gorm:"type:varchar(255);not null" json:"link"`
}

type UserClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
