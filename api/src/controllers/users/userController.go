package users

import (
	MW "breakfast/middleware"
	"breakfast/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func generateJWTToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.UserClaims{
		UserID:    user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return tokenString, err
}

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting UserController")
	mux.Handle("GET /greet/{id}", MW.AuthMiddleware(http.HandlerFunc(greetUserByID)))
	mux.HandleFunc("POST /register", registerUser)
	mux.HandleFunc("POST /login", loginUser)
}
