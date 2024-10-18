package models

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func CheckUserPassword(password_a string, password_b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password_a), []byte(password_b))
	if err != nil {
		return err
	}
	return nil
}
