package models

import (
	BFE "breakfast/_internal/errors"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
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

func CheckUserPassword(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return BFE.New(BFE.ErrPassword, err)
	}
	return nil
}

func GetUserClaims(r *http.Request) (*UserClaims, error) {
	claims, ok := r.Context().Value(ClaimsContext).(*UserClaims)
	if !ok {
		return nil, BFE.New(BFE.ErrClaims, errors.New("Missing/malformed claims"))
	}
	return claims, nil
}

func GetUserID(r *http.Request) (uuid.UUID, error) {
	claims, err := GetUserClaims(r)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claims.UserID)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
