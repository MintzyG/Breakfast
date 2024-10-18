package models

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u User) String() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}
