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
func Run(mux *http.ServeMux) {
	fmt.Println("Connecting YogurtController")
	mux.Handle("POST /api/v1/yogurt", MW.AuthMiddleware(http.HandlerFunc(greetUserByID)))
}
