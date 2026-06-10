package main

import (
	"log"
	"net/http"
	"pancake/internal/db"

	"pancake/internal/notes"
	"pancake/internal/scim"
	"pancake/internal/ws"

	"github.com/MintzyG/fun"
	"github.com/go-chi/chi/v5"
)

func init() {
	fun.SetPathParamFunc(chi.URLParam)
}

func main() {
	mongo, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	notesRepo := notes.NewRepository(mongo.Collection("notes"))
	notesHandler := notes.NewHandler(notesRepo)
	scimRepo := scim.NewRepository(mongo.Collection("scim_users"))
	scimHandler := scim.NewHandler(scimRepo)
	wsHandler := ws.NewHandler()

	r := chi.NewRouter()

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", notesHandler.List)
		r.Post("/", notesHandler.Create)
		r.Get("/{id}", notesHandler.Get)
		r.Patch("/{id}", notesHandler.Patch)
		r.Delete("/{id}", notesHandler.Delete)
	})

	r.Route("/scim/v2/Users", func(r chi.Router) {
		r.Get("/", scimHandler.ListUsers)
		r.Post("/", scimHandler.CreateUser)
		r.Get("/{id}", scimHandler.GetUser)
		r.Put("/{id}", scimHandler.ReplaceUser)
		r.Patch("/{id}", scimHandler.PatchUser)
		r.Delete("/{id}", scimHandler.DeleteUser)
	})

	r.Get("/ws/{docId}", wsHandler.Connect)

	log.Println("pancake running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
