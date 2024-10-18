package users

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"breakfast/models"
	DB "breakfast/repositories"
	"github.com/golang-jwt/jwt/v5"
)

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
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

	DB.CheckUserPassword(user)

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
