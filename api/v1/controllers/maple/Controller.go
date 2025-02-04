package maple

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting MapleController")
	mux.Handle("DELETE /api/v1/maple/{id}", MW.AuthMiddleware(http.HandlerFunc(deleteHabit)))
	mux.Handle("PATCH /api/v1/maple/{id}", MW.AuthMiddleware(http.HandlerFunc(patchHabit)))
	mux.Handle("POST /api/v1/maple/{id}", MW.AuthMiddleware(http.HandlerFunc(markDay)))
	mux.Handle("POST /api/v1/maple", MW.AuthMiddleware(http.HandlerFunc(createHabit)))
	mux.Handle("GET /api/v1/maple/{id}", MW.AuthMiddleware(http.HandlerFunc(getHabitByID)))
	mux.Handle("GET /api/v1/maple", MW.AuthMiddleware(http.HandlerFunc(getAllHabits)))
}
