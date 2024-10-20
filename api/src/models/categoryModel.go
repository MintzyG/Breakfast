package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Category struct {
	ID          int       `json:"CategoryId"`
	UserId      uuid.UUID `json:"UserId"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Emoji       string    `json:"Emoji"`
	Color       string    `json:"Color"`
	TextColor   string    `json:"TextColor"`
}

func (c Category) String() string {
	return fmt.Sprintf("%v - %v", c.Emoji, c.Title)
}
