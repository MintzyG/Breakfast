package main

import (
	"log"
	"net/http"
	"os"

	"breakfast/config"
	"breakfast/internal/db"
	"breakfast/internal/handlers"
	mw "breakfast/internal/middleware"
	"breakfast/internal/repositories"
	"breakfast/internal/services"

	"github.com/rs/cors"
)

func main() {
	cfg := config.LoadConfig()
	database := db.Connect(cfg.DSN)
	db.Migrate()

	// Repos
	userRepo := repositories.NewUserRepository(database)

	// Services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	dataService := services.NewDataService(userRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	dataHandler := handlers.NewDataHandler(dataService)

	mux := http.NewServeMux()

	// Authentication
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	// Data Endpoints
	mux.Handle("GET /me", mw.AuthMiddleware(http.HandlerFunc(dataHandler.HelloMe)))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
