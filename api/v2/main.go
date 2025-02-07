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
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	database := db.Connect(cfg.DSN)
	db.Migrate()

	mux := intializeMux(database, cfg)

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

func intializeMux(database *gorm.DB, cfg *config.Config) *http.ServeMux {
	// Repos
	userRepo := repositories.NewUserRepository(database)
	pancakeRepo := repositories.NewPancakeRepository(database)
	yogurtRepo := repositories.NewYogurtRepository(database)
	mapleRepo := repositories.NewMapleRepository(database)
	espressoRepo := repositories.NewEspressoRepository(database)
	toastRepo := repositories.NewToastRepository(database)
	cerealRepo := repositories.NewCerealRepository(database)
	parfaitRepo := repositories.NewParfaitRepository(database)
	omeletteRepo := repositories.NewOmeletteRepository(database)

	// Services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	pancakeService := services.NewPancakeService(pancakeRepo)
	yogurtService := services.NewYogurtService(yogurtRepo)
	mapleService := services.NewMapleService(mapleRepo)
	espressoService := services.NewEspressoService(espressoRepo)
	toastService := services.NewToastService(toastRepo)
	cerealService := services.NewCerealService(cerealRepo)
	parfaitService := services.NewParfaitService(parfaitRepo)
	omeletteService := services.NewOmeletteService(omeletteRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	pancakeHandler := handlers.NewPancakeHandler(pancakeService)
	yogurtHandler := handlers.NewYogurtHandler(yogurtService)
	mapleHandler := handlers.NewMapleHandler(mapleService)
	espressoHandler := handlers.NewEspressoHandler(espressoService)
	toastHandler := handlers.NewToastHandler(toastService)
	cerealHandler := handlers.NewCerealHandler(cerealService)
	parfaitHandler := handlers.NewParfaitHandler(parfaitService)
	omeletteHandler := handlers.NewOmeletteHandler(omeletteService)

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
	mux.Handle("GET /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.GetByID)))
	mux.Handle("GET /maple", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.GetAll)))
	mux.Handle("PATCH /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.Update)))
	mux.Handle("DELETE /maple/{id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.Delete)))
	// MapleDays Endpoints
	mux.Handle("POST /maple/{id}/day", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.CreateDay)))
	mux.Handle("GET /maple/{id}/day/{day_id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.GetDay)))
	mux.Handle("PATCH /maple/{id}/day/{day_id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.UpdateDay)))
	mux.Handle("DELETE /maple/{id}/day/{day_id}", mw.AuthMiddleware(http.HandlerFunc(mapleHandler.DeleteDay)))

	// Espresso Endpoints
	mux.Handle("POST /espresso", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Create)))
	mux.Handle("GET /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.GetByID)))
	mux.Handle("GET /espresso", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.GetAll)))
	mux.Handle("PATCH /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Update)))
	mux.Handle("DELETE /espresso/{id}", mw.AuthMiddleware(http.HandlerFunc(espressoHandler.Delete)))

	// Toast Endpoints
	mux.Handle("POST /toast", mw.AuthMiddleware(http.HandlerFunc(toastHandler.Create)))
	mux.Handle("GET /toast/{id}", mw.AuthMiddleware(http.HandlerFunc(toastHandler.GetByID)))
	mux.Handle("GET /toast", mw.AuthMiddleware(http.HandlerFunc(toastHandler.GetAll)))
	mux.Handle("PATCH /toast/{id}", mw.AuthMiddleware(http.HandlerFunc(toastHandler.Update)))
	mux.Handle("DELETE /toast/{id}", mw.AuthMiddleware(http.HandlerFunc(toastHandler.Delete)))

	// Cereal Endpoints
	mux.Handle("POST /cereal", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.Create)))
	mux.Handle("GET /cereal/{id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.GetByID)))
	mux.Handle("GET /cereal", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.GetAll)))
	mux.Handle("PATCH /cereal/{id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.Update)))
	mux.Handle("DELETE /cereal/{id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.Delete)))
	// Cereal Activities Endpoints
	mux.Handle("POST /cereal/{id}/activity", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.CreateActivity)))
	mux.Handle("GET /cereal/{id}/activity/{activity_id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.GetActivity)))
	mux.Handle("PATCH /cereal/{id}/activity/{activity_id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.UpdateActivity)))
	mux.Handle("DELETE /cereal/{id}/activity/{activity_id}", mw.AuthMiddleware(http.HandlerFunc(cerealHandler.DeleteActivity)))

	// Parfait Endpoints
	mux.Handle("POST /parfait", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.Create)))
	mux.Handle("GET /parfait/{id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.GetByID)))
	mux.Handle("GET /parfait", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.GetAll)))
	mux.Handle("PATCH /parfait/{id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.Update)))
	mux.Handle("DELETE /parfait/{id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.Delete)))
	// Parfait Reminders Endpoint
	mux.Handle("POST /parfait/{id}/reminder", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.CreateReminder)))
	mux.Handle("GET /parfait/{id}/reminder/{reminder_id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.GetReminder)))
	mux.Handle("GET /reminders", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.GetAllReminders)))
	mux.Handle("PATCH /parfait/{id}/reminder/{reminder_id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.UpdateReminder)))
	mux.Handle("DELETE /parfait/{id}/reminder/{reminder_id}", mw.AuthMiddleware(http.HandlerFunc(parfaitHandler.DeleteReminder)))

	// Omelette Endpoints
	mux.Handle("POST /omelette", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.Create)))
	mux.Handle("GET /omelette/{id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetByID)))
	mux.Handle("GET /omelette", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetAll)))
	mux.Handle("PATCH /omelette/{id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.Update)))
	mux.Handle("DELETE /omelette/{id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.Delete)))
	// Omelette Lists Endpoints
	mux.Handle("POST /omelette/{id}/list", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.CreateList)))
	mux.Handle("GET /omelette/{id}/list/{list_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetListByID)))
	mux.Handle("GET /list", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetAllLists)))
	mux.Handle("PATCH /omelette/{id}/list/{list_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.UpdateList)))
	mux.Handle("DELETE /omelette/{id}/list/{list_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.DeleteList)))
	// Omelette Card Endpoints
	mux.Handle("POST /omelette/{id}/list/{list_id}/card", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.CreateCard)))
	mux.Handle("GET /omelette/{id}/list/{list_id}/card/{card_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetCardByID)))
	mux.Handle("GET /card", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.GetAllCards)))
	mux.Handle("PATCH /omelette/{id}/list/{list_id}/card/{card_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.UpdateCard)))
	mux.Handle("DELETE /omelette/{id}/list/{list_id}/card/{card_id}", mw.AuthMiddleware(http.HandlerFunc(omeletteHandler.DeleteCard)))

	return mux
}
