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
	pancakeRepo := repositories.NewPancakeRepository(database)

	// Services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	dataService := services.NewDataService(userRepo)
	pancakeService := services.NewPancakeService(pancakeRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	dataHandler := handlers.NewDataHandler(dataService)
	pancakeHandler := handlers.NewPancakeHandler(pancakeService)

	mux := http.NewServeMux()

	// Authentication
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	// Pancake Endpoints
	mux.Handle("POST /pancake/create", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Create)))
	mux.Handle("GET /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.GetByID)))
	mux.Handle("GET /pancake", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.GetNotes)))
	mux.Handle("PATCH /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Update)))
	mux.Handle("DELETE /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Delete)))

	// Data Endpoints
	mux.Handle("GET /me", mw.AuthMiddleware(http.HandlerFunc(dataHandler.HelloMe)))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
