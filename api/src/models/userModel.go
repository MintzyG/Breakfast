package models

import (
  BFE "breakfast/errors"
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
		return BFE.NewBFError(BFE.PASSWORD_ERROR_CODE, err.Error())
	}
	return nil
}

func GetUserClaims(r *http.Request) (*UserClaims, error) {
    claims, ok := r.Context().Value("claims").(*UserClaims)
    if !ok {
        return nil, BFE.NewBFError(BFE.CLAIMS_ERROR_CODE, "Missing/malformed claims")
    }
    return claims, nil
}
