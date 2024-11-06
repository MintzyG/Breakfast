package pancake

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting PancakeController")
	mux.Handle("DELETE /api/v1/pancake/{id}", MW.AuthMiddleware(http.HandlerFunc(deleteNote)))
	mux.Handle("PATCH /api/v1/pancake/{id}", MW.AuthMiddleware(http.HandlerFunc(patchNote)))
	mux.Handle("POST /api/v1/pancake", MW.AuthMiddleware(http.HandlerFunc(createNote)))
	mux.Handle("GET /api/v1/pancake/{id}", MW.AuthMiddleware(http.HandlerFunc(getNoteByID)))
	mux.Handle("GET /api/v1/pancake", MW.AuthMiddleware(http.HandlerFunc(getAllNotes)))
}
