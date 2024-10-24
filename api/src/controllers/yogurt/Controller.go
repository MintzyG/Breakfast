package yogurt

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting YogurtController")
	mux.Handle("PATCH /api/v1/yogurt/{id}", MW.AuthMiddleware(http.HandlerFunc(patchTask)))
	mux.Handle("POST /api/v1/yogurt", MW.AuthMiddleware(http.HandlerFunc(createTask)))
	mux.Handle("GET /api/v1/yogurt", MW.AuthMiddleware(http.HandlerFunc(getAllTasks)))
	mux.Handle("GET /api/v1/yogurt/{id}", MW.AuthMiddleware(http.HandlerFunc(getTaskByID)))
}
