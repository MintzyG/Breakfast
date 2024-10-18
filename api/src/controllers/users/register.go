package users

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	RSP "breakfast/response"
	"breakfast/models"
	"github.com/lib/pq"
	"github.com/google/uuid"
	DB "breakfast/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Empty name fields", "EMPTY_REQUIRED_FIELDS")
		return
	}

	if user.Email == "" {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Empty email field", "EMPTY_REQUIRED_FIELDS")
		return
	}

	if user.Password == "" {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Empty password field", "EMPTY_REQUIRED_FIELDS")
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, err.Error(), "PASSWORD_ERROR")
		return
	}
	user.Password = string(bytes)

	user.ID = uuid.New()
	err = DB.CreateUser(&user)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			RSP.SendErrorResponse(w, http.StatusConflict, "User with this email already exists", "USER_EXISTS")
			return
        }
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err.Error()), "DATABASE_ERROR")
        return
    }

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &user_claims {
		UserID:    user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error authenticating user: %v", err.Error()), "AUTH_ERROR")
		return
	}

	RSP.SendSuccessResponse(w, http.StatusCreated, tokenString)
}
