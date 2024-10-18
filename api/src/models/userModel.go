package models

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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

type UserClaims struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
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
