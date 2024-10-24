package categories

import (
	MW "breakfast/middleware"
	"fmt"
	"net/http"
)

func Run(mux *http.ServeMux) {
	fmt.Println("Connecting CategoryController")
	mux.Handle("DELETE /api/v1/categories/{id}", MW.AuthMiddleware(http.HandlerFunc(deleteCategory)))
	mux.Handle("PATCH /api/v1/categories/{id}", MW.AuthMiddleware(http.HandlerFunc(patchCategory)))
	mux.Handle("POST /api/v1/categories", MW.AuthMiddleware(http.HandlerFunc(createCategory)))
	mux.Handle("GET /api/v1/categories/{id}", MW.AuthMiddleware(http.HandlerFunc(getCategoryByID)))
	mux.Handle("GET /api/v1/categories", MW.AuthMiddleware(http.HandlerFunc(getAllCategories)))
}
