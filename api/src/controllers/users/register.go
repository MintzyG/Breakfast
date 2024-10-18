package users

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"breakfast/models"
	"github.com/google/uuid"
	DB "breakfast/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" || user.LastName == "" {
		http.Error(w, "Missing required name fields", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Missing required email field", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Missing required password field", http.StatusBadRequest)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error hashing password: %v", err.Error()), http.StatusBadRequest)
	}
	user.Password = string(bytes)

	user.ID = uuid.New()
	err = DB.CreateUser(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err.Error()), http.StatusInternalServerError)
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
		http.Error(w, fmt.Sprintf("Error creating auth token: %v", err.Error()), http.StatusInternalServerError)
	}


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(tokenString); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
    }
}
