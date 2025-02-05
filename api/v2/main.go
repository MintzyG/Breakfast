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
	yogurtRepo := repositories.NewYogurtRepository(database)
	mapleRepo := repositories.NewMapleRepository(database)
	espressoRepo := repositories.NewEspressoRepository(database)

	// Services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	pancakeService := services.NewPancakeService(pancakeRepo)
	yogurtService := services.NewYogurtService(yogurtRepo)
	mapleService := services.NewMapleService(mapleRepo)
	espressoService := services.NewEspressoService(espressoRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	pancakeHandler := handlers.NewPancakeHandler(pancakeService)
	yogurtHandler := handlers.NewYogurtHandler(yogurtService)
	mapleHandler := handlers.NewMapleHandler(mapleService)
	espressoHandler := handlers.NewEspressoHandler(espressoService)

	mux := http.NewServeMux()

	// Authentication
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	// Pancake Endpoints
	mux.Handle("POST /pancake", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Create)))
	mux.Handle("GET /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.GetByID)))
	mux.Handle("GET /pancake", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.GetAll)))
	mux.Handle("PATCH /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Update)))
	mux.Handle("DELETE /pancake/{id}", mw.AuthMiddleware(http.HandlerFunc(pancakeHandler.Delete)))

	// Yogurt Endpoints
	mux.Handle("POST /yogurt", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.Create)))
	mux.Handle("GET /yogurt/{id}", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.GetByID)))
	mux.Handle("GET /yogurt", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.GetAll)))
	mux.Handle("PATCH /yogurt/{id}", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.Update)))
	mux.Handle("PATCH /yogurt/{id}/completed", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.UpdateCompleted)))
	mux.Handle("DELETE /yogurt/{id}", mw.AuthMiddleware(http.HandlerFunc(yogurtHandler.Delete)))

	// Maple Endpoints
	mux.Handle("POST /maple", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.Create)))
	mux.Handle("POST /maple/{id}/day", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.CreateDay)))
	mux.Handle("GET /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.GetByID)))
	mux.Handle("GET /maple", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.GetAll)))
	mux.Handle("PATCH /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.Update)))
	mux.Handle("DELETE /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.Delete)))

  // Espresso Endpoints
  mux.Handle("POST /espresso", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Create)))
  mux.Handle("GET /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.GetByID)))
  mux.Handle("GET /espresso", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.GetAll)))
  mux.Handle("PATCH /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Update)))
  mux.Handle("DELETE /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Delete)))

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
