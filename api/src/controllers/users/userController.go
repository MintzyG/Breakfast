package users

import (
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

type user_claims struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting UserController")
	mux.HandleFunc("GET /greet/{id}", greetUserByID)
	mux.HandleFunc("POST /register", registerUser)
	mux.HandleFunc("POST /login", loginUser)
}
