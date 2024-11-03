package toast

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting ToastController")
	mux.Handle("DELETE /api/v1/toast/{id}", MW.AuthMiddleware(http.HandlerFunc(deleteSession)))
	// mux.Handle("PATCH /api/v1/toast/{id}", MW.AuthMiddleware(http.HandlerFunc(patchSession)))
	mux.Handle("POST /api/v1/toast", MW.AuthMiddleware(http.HandlerFunc(createSession)))
	mux.Handle("POST /api/v1/toast/start", MW.AuthMiddleware(http.HandlerFunc(startSession)))
	mux.Handle("POST /api/v1/toast/stop", MW.AuthMiddleware(http.HandlerFunc(stopSession)))
	// mux.Handle("GET /api/v1/toast", MW.AuthMiddleware(http.HandlerFunc(getAllSessions)))
	mux.Handle("GET /api/v1/toast/{id}", MW.AuthMiddleware(http.HandlerFunc(getSessionByID)))
}
